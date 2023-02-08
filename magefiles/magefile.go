//go:build mage
// +build mage

package main

import (
	"fmt"
	"go/build"
	"os"
	"path/filepath"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
	"github.com/magefile/mage/sh"
	"github.com/magefile/mage/target"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

// A build step that requires additional params, or platform specific steps for example
// func Build() error {
// 	mg.Deps(InstallDeps)
// 	fmt.Println("Building...")
// 	cmd := exec.Command("go", "build", "-o", "MyApp", ".")
// 	return cmd.Run()
// }

// A custom install step if you need your bin someplace other than go/bin
// func Install() error {
// 	mg.Deps(Build)
// 	fmt.Println("Installing...")
// 	return os.Rename("./MyApp", "/usr/bin/MyApp")
// }

// template.qtpl.go: template.qtpl
//
//	go generate *.go
func TemplateQtplGo() error {
	if updated, err := target.Path("template.qtpl.go", "template.qtpl"); err != nil {
		return err
	} else if updated {
		return sh.RunV("go", "generate", "-v", ".")
	}
	return nil
}

// parser/parse_string.go: parser/parse.go
//
//	go generate ./parser/parse.go
func ParseStringerGo() error {
	src := "./parser/parse.go"
	if updated, err := target.Path("parser/parse_string.go", src); err != nil {
		return err
	} else if updated {
		return sh.RunV("go", "generate", src)
	}
	return nil
}

// vet: template.qtpl.go *.go parser/parse_string.go
//
//	@# adjust path produced in error meesage
//	go vet .
func Vet() error {
	mg.Deps(TemplateQtplGo, ParseStringerGo)
	return sh.RunV("go", "vet", ".")
}

// capi:
//
//	go run . -pkgdir ../cefingo
func Capi() error {
	return sh.RunV("go", "run", ".", "-pkgdir", "../cefingo")
}

func Fmt() error {
	if files, err := filepath.Glob("*.go"); err == nil {
		fmtArgs := append([]string{"fmt"}, files...)
		if err0 := sh.RunV("go", fmtArgs...); err0 != nil {
			return err0
		}
	} else {
		return err
	}

	if files, err := filepath.Glob("./parser/*.go"); err == nil {
		fmtArgs := append([]string{"fmt"}, files...)
		return sh.RunV("go", fmtArgs...)
	} else {
		return err
	}
}

//go:generate go run github.com/valyala/quicktemplate/qtc@latest -file=pkgconfig.qtpl
//go:generate sed -i -e "3i //go:build mage" *.qtpl.go

func genPkgConfig(out *os.File, cefBinaryPath string) error {
	var binDir string
	var err error
	if binDir, err = filepath.Abs(cefBinaryPath); err != nil {
		return err
	}
	// fmt.Println("T51:", binDir)
	WriteGenPkgConfig(out, filepath.ToSlash(binDir))
	return nil
}

func Pkgconfig(cefBinaryPath string) error {
	return genPkgConfig(os.Stdout, cefBinaryPath)
}

func CheckPkgConfig() error {
	return sh.RunV("pkgconf", "cefingo", "--variable=includedir")
}

func InstallPkgConfig(cefBinaryPath string) error {
	configPath := filepath.SplitList(os.Getenv("PKG_CONFIG_PATH"))
	if len(configPath) == 0 {
		return fmt.Errorf("PKG_CONFIG_PATH no defined")
	}
	configName := filepath.Join(configPath[0], "cefingo.pc")
	out, err := os.Create(configName)
	if err != nil {
		return err
	}
	defer out.Close()
	return genPkgConfig(out, cefBinaryPath)
}

func InstallCefLibrary(cefBinaryPath string) error {
	if err := InstallPkgConfig(cefBinaryPath); err != nil {
		return err
	}

	if binDir, err := filepath.Abs(cefBinaryPath); err == nil {
		targetDir := filepath.SplitList(build.Default.GOPATH)[0]
		if err := sh.RunV("cp", "-r", filepath.ToSlash(filepath.Join(binDir, "Release", "*")), filepath.ToSlash(filepath.Join(targetDir, "bin"))); err != nil {
			return err
		}
		if err := sh.RunV("cp", "-r", filepath.ToSlash(filepath.Join(binDir, "Resources", "*")), filepath.ToSlash(filepath.Join(targetDir, "bin"))); err != nil {
			return err
		}
	} else {
		return err
	}
	return nil
}
