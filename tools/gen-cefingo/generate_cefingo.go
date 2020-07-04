package main

//go:generate go run github.com/valyala/quicktemplate/qtc

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/turutcrane/cefingo/tools/gen-cefingo/parser"
)

var (
	goFormat     bool = true
	fileComments map[int][]string
	capiDir      = flag.String("capidir", "capi", "outpu directory")
	logTags      LogTag
)

func main() {
	flag.Parse()
	tus := parser.Parse()
	logTagFile := filepath.Join(*capiDir, "logtag.txt")
	logTags.ReadTags(logTagFile)
	defer logTags.WriteToFile(logTagFile)
	log.Println("T37:", logTags)

	for i, tu := range tus {
		parser.ExternalDeclaration(i, tu.ExternalDeclaration)
	}

	goTypeFile := newFileGo("cefingo_types_gen.go", false)
	defer outFileGo(goTypeFile)

	goTypeWinFile := newFileGo("cefingo_types_win_gen.go", false)
	defer outFileGo(goTypeWinFile)

	goExportFile := newFileGo("cefingo_exports_gen.go", false)
	defer outFileGo(goExportFile)

	cFile, hFile := newFileCH()
	defer outFileCH(cFile, hFile)

	var goFile *Generator
	fn := []string{}
	for fname, _ := range parser.FileDefs {
		fn = append(fn, fname)
	}
	sort.Strings(fn)

	goGenFile := newFileGo("cefingo_gen.go", true)
	defer outFileGo(goGenFile)

	for _, fname := range fn {
		log.Printf("T35: Process %s\n", fname)
		if strings.HasSuffix(fname, "_capi.h") {
			goFile = goGenFile // newFileGo(makeGoFileName(fname), parser.HasHandlerClass(fname))
		} else if fname == "cef_types_win.h" {
			goFile = goTypeWinFile
		} else {
			goFile = goTypeFile
		}
		log.Printf("T13: %s\n", fname)
		a := parser.FileDefs[fname]
		first := true
		for _, d := range a {
			token := d.Common().Token()
			if first {
				getComments(token.AbsFilename())
				goFile.Printf("\n// %s, %s, \n", fname, token.FilePos())
				first = false
			}
			d.SetComment(fileComments)
			switch d.Common().Dk {
			case parser.DkSimple:
				s := d.(*parser.SimpleDecl)
				outSimple(goFile, s)
			case parser.DkStruct:
				if s, ok := d.(*parser.StructDecl); ok {
					outStruct(goFile, s)
				} else if s, ok := d.(*parser.SimpleDecl); ok {
					outSimple(goFile, s)
				} else {
					log.Panicf("T84: Can not handle: %v\n", d)
				}
			case parser.DkEnum:
				e := d.(*parser.EnumDecl)
				outEnum(goFile, e)
			case parser.DkCefClass:
				s := d.(*parser.CefClassDecl)
				if parser.IsHandlerClass(s.Token()) {
					outHandlerClass(goFile, goExportFile, cFile, hFile, s)
				} else {
					outCefObjectClass(goFile, cFile, hFile, s)
				}
			case parser.DkFunc:
				f := d.(*parser.FuncDecl)
				outGoFunction(goFile, f)
			default:
				log.Printf("T31: %s: %s\n", d.Common().Dk, token.Name())
			}
		}
		// if strings.HasSuffix(fname, "_capi.h") {
		// 	log.Printf("T63: %s\n", fname)
		// 	outFileGo(goFile)
		// }
	}

	log.Printf("T10: End: %d Translation Units", len(tus))
}

func newFileGo(fname string, hasCallbackClass bool) (gf *Generator) {
	gf = &Generator{}
	gf.fname = fname

	var imports []string
	var defIfMutex bool
	var winOnly bool
	if fname == "cefingo_types_gen.go" {
		imports = []string{"sync", "time", "unsafe"}
		defIfMutex = true
	} else if fname == "cefingo_exports.go" {
		imports = []string{"unsafe", "runtime"}
	} else if fname == "cefingo_types_win_gen.go" {
		winOnly = true
	} else {
		imports = []string{"unsafe", "runtime"}
	}

	WriteGoHead(gf, imports, defIfMutex, winOnly)
	return gf
}

func makeGoFileName(capiH string) (gfile string) {
	if strings.HasSuffix(capiH, "_capi.h") {
		name := strings.Replace(capiH, "cef_", "", 1)
		name = strings.Replace(name, "_capi.h", "", 1)
		gfile = name + "_gen.go"
	} else {
		log.Panicf("T117: Wrong file: %s\n", capiH)
	}

	return gfile
}

func outFileGo(gf *Generator) {
	// Format the output.
	var src []byte
	if goFormat {
		src = gf.format()
	} else {
		gf.Printf("// Not Formatted.!")
		src = gf.buf.Bytes()
	}

	// Write to file.
	outputName := filepath.Join(*capiDir, gf.fname)

	err := ioutil.WriteFile(outputName, src, 0644)
	if err != nil {
		log.Panicf("writing output: %s", err)
	}
}

func newFileCH() (cFile, hFile *Generator) {
	cFile = &Generator{}
	cFile.Printf(`// Code generated by "genCefingo.go" DO NOT EDIT.
#include "cefingo.h"
#include "_cgo_export.h"
`)

	hFile = &Generator{}
	hFile.Printf(`// Code generated by "genCefingo.go" DO NOT EDIT.
#ifndef CEFINGO_GEN_H_
#define CEFINGO_GEN_H_
`)
	return cFile, hFile
}

func outFileCH(cFile, hFile *Generator) {

	// Write to C file.
	outputCName := filepath.Join(*capiDir, "cefingo_gen.c")
	err := ioutil.WriteFile(outputCName, cFile.buf.Bytes(), 0644)
	if err != nil {
		log.Panicf("writing output: %s", err)
	}

	// Write to H file.
	hFile.Printf("#endif //CEFINGO_GEN_H_")
	outputHName := filepath.Join(*capiDir, "cefingo_gen.h")
	err = ioutil.WriteFile(outputHName, hFile.buf.Bytes(), 0644)
	if err != nil {
		log.Panicf("writing output: %s", err)
	}
}

func outComment(gf *Generator, line int) {
	if comments, ok := fileComments[line]; ok {
		gf.Printf("\n")
		for _, c := range comments {
			gf.Printf("%s\n", c)
		}
	}
}

func outSimple(gf *Generator, d *parser.SimpleDecl) {
	line := d.LineOfTypedef()
	outComment(gf, line)
	// fmt.Fprintf(fileGoTypes, "// Pos: %d\n", d.Line())
	gf.Printf("type %s C.%s\n", d.GoName(), d.CefName())
}

func outStruct(gf *Generator, d *parser.StructDecl) {
	line := d.LineOfTypedef()
	outComment(gf, line)
	// fmt.Fprintf(fileGoTypes, "// Pos: %d\n", d.Line())
	gf.Printf("type %s C.%s\n", d.GoName(), d.CefName())
	WriteNewStruct(gf, d)
	for i, m := range d.Members {
		if !(i == 0 && m.Type().Ty == parser.TySizeT && m.Name() == "size") {
			WriteMemberAccessor(gf, d, m)
		}
	}
}

func outEnum(gf *Generator, d *parser.EnumDecl) {
	outComment(gf, d.LineOfTypedef())
	gf.Printf("type %s C.%s\n", d.GoName(), d.CefName())
	gf.Printf("const (\n")
	for _, e := range d.Enums {
		outComment(gf, e.Line())
		gf.Printf("\t%s %s = C.%s\n", e.TitleCase(), d.GoName(), e.Name())
	}
	gf.Printf(")\n")
}

func getComments(fname string) {
	fileComments = map[int][]string{}
	f, err := os.Open(fname)
	if err != nil {
		log.Panicf("T186: %s, %v", fname, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	line := 0
	comment := false
	cRegex := regexp.MustCompile("^[[:blank:]]*//")
	var block []string
	for scanner.Scan() {
		line++
		b := scanner.Bytes()
		if cRegex.Match(b) {
			if comment {
				block = append(block, strings.TrimLeft(scanner.Text(), " "))
			} else {
				comment = true
				block = []string{strings.TrimLeft(scanner.Text(), " ")}
			}
		} else {
			if comment {
				fileComments[line] = block
				// fmt.Printf("T111: %d: %s\n", line, scanner.Text())
			}
			comment = false
		}
	}
}

func outHandlerClass(gf, expf, cf, hf *Generator, d *parser.CefClassDecl) {
	log.Printf("T226: Callback class: %s\n", d.Token())
	logTags.ResetTag(d.Token().GoName())

	outComment(gf, d.LineOfTypedef())
	WriteGoType(gf, d, &logTags)
	genGoInterface(gf, d)
	genGoAlloc(gf, d)
	genBindFunc(gf, d)
	genGoGetFunc(gf, expf, d)
	genGoCallbackFunc(expf, d)
	genChConstructFunc(cf, hf, d)
}

func genGoInterface(gf *Generator, st *parser.CefClassDecl) {
	for _, m := range st.Methods {
		if !m.IsGetFunc() {
			WriteGoIface(gf, m)
		}
	}
	WriteIfaceStruct(gf, st)
}

func genGoAlloc(gf *Generator, st *parser.CefClassDecl) {
	WriteGoAllocFunc(gf, st, &logTags)
}

func genBindFunc(gf *Generator, st *parser.CefClassDecl) {
	WriteGoBindFunc(gf, st, &logTags)
}

func genGoGetFunc(gf, expf *Generator, st *parser.CefClassDecl) {
	goName := st.GoName()
	baseName := st.BaseName()
	cName := st.CefName()
	for c := st; c != nil; c = c.GetBase() {
		for _, m := range c.Methods {
			if m.IsGetFunc() {
				WriteGoGetFunc(expf, m, goName, baseName, cName, &logTags)
				WriteAssocGetFunc(gf, m, goName, baseName, cName, &logTags)
			}
		}
	}
}

func genGoCallbackFunc(gf *Generator, st *parser.CefClassDecl) {
	for c := st; c != nil; c = c.GetBase() {
		for _, m := range c.Methods {
			if !m.IsGetFunc() {
				WriteGoCallback(gf, m, st, &logTags)
			}
		}
	}
}

func ConvToGoTypeExp(t parser.Type, val string) (exp string) {
	switch t.Ty {
	case parser.TyInt, parser.TyInt32, parser.TyInt64, parser.TyUint32, parser.TyUint64,
		parser.TyFloat, parser.TyDouble,
		parser.TyStructSimple, parser.TySimple, parser.TySizeT, parser.TyHWND, parser.TyHCURSOR:
		exp = "(" + t.GoType() + ")(" + val + ")"
	case parser.TyStructRefCounted:
		if t.Pointer == 1 {
			exp = "new" + t.Deref().GoType() + "(" + val + ")"
		} else {
			log.Panicf("T448: %s: %v, %t, %v\n", val, t.Ty, t.IsRefCountedClass(), t.GoType())
		}
	case parser.TyStructScoped:
		if t.Pointer == 1 {
			exp = "new" + t.Deref().GoType() + "(" + val + ")"
		} else {
			log.Panicf("T454: %s: %v\n", val, t)
		}
	case parser.TyStringT:
		if t.Pointer == 1 {
			exp = "string_from_cef_string(" + val + ")"
		} else {
			log.Panicf("T460: %s: %v\n", val, t)
		}
	case parser.TyEnum:
		exp = t.GoType() + "(" + val + ")"
	case parser.TyVoid:
		switch t.Pointer {
		case 1:
			exp = "unsafe.Pointer(" + val + ")"
		default:
			log.Panicf("T474: %s: %v\n", val, t)
		}
	case parser.TyMSG:
		if t.Pointer == 1 {
			exp = t.Deref().GoType() + "(" + val + ")"
		} else {
			log.Panicf("E326: %v\n", t)
		}
	default:
		log.Panicf("T456: %s: %v\n", val, t)
	}
	return exp
}

func genChConstructFunc(cf, hf *Generator, st *parser.CefClassDecl) {
	WriteCConstruct(cf, st, &logTags)

	WriteHCallback(hf, st, &logTags)
}

func outCefObjectClass(gf, cf, hf *Generator, d *parser.CefClassDecl) {
	log.Printf("T541: Cef object: %s\n", d.Token())
	logTags.ResetTag(d.Token().GoName())

	outComment(gf, d.LineOfTypedef())
	WriteGoType(gf, d, &logTags)
	outCefObjectMethod(gf, cf, hf, d)
}

func outGoFunction(gf *Generator, f *parser.FuncDecl) {
	log.Printf("T350: Cef Functin : %s\n", f.Token())
	logTags.ResetTag(f.Token().GoName())

	outComment(gf, f.LineOfTypedef())
	WriteGoFunction(gf, f, &logTags)
}

func outCefObjectMethod(gf, cf, hf *Generator, d *parser.CefClassDecl) {
	for _, m := range d.Methods {
		WriteCefObjectMethod(gf, m, &logTags)
		WriteCefObjectMethodC(cf, m)
		WriteCefObjectMethodH(hf, m)
	}
}

// Generator holds the state of the analysis. Primarily used to buffer
// the output for format.Source.
type Generator struct {
	fname string
	buf   bytes.Buffer // Accumulated output.
}

// Printf output fromated string
func (g *Generator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
}

func (g *Generator) Write(p []byte) (n int, err error) {
	return g.buf.Write(p)
}

// func (g *Generator) Lines() (line int) {
// 	return bytes.Count(g.buf.Bytes(), []byte("\n"))
// }

// format returns the gofmt-ed contents of the Generator's buffer.
func (g *Generator) format() []byte {
	src, err := format.Source(g.buf.Bytes())
	if err != nil {
		// Should never happen, but can arise when developing this code.
		// The user can compile the output to see the error.
		log.Printf("warning: internal error: invalid Go generated: %s", err)
		log.Printf("warning: compile the package to analyze the error")
		return g.buf.Bytes()
	}
	return src
}

// LogTag magages Tag number
type LogTag struct {
	TagMap map[string]int
	MaxTag int
	curTag int
	subTag int
}

// NextTag returns next tag string
func (logTag *LogTag) NextTag() string {
	logTag.subTag++
	return fmt.Sprintf("T%d.%d", logTag.curTag, logTag.subTag)
}

// ResetTag set curTag and reset subTag
func (logTag *LogTag) ResetTag(key string) {
	if t, ok := logTag.TagMap[key]; ok {
		logTag.curTag = t
	} else {
		logTag.MaxTag++
		logTag.TagMap[key] = logTag.MaxTag
	}
	logTag.subTag = 0
}

// ReadTags reads tag number from file
func (logTag *LogTag) ReadTags(fname string) {
	logTag.TagMap = map[string]int{}
	logTag.MaxTag = 100
	if b, err := ioutil.ReadFile(fname); err == nil {
		if err := json.Unmarshal(b, logTag); err != nil {
			log.Panicln("T465:", err)
		}
	}
}

// WriteWriteToFile writes tag number to file
func (logTag *LogTag) WriteToFile(fname string) {
	b, err := json.Marshal(logTag)
	if err != nil {
		log.Panicln("T475:", err)
	}
	err = ioutil.WriteFile(fname, b, 0644)
	if err != nil {
		log.Panicln("T475:", err)
	}
}
