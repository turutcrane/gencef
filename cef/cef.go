package cef

import (
	"os"
	"unsafe"

	"github.com/turutcrane/cefingo/capi"
)

// #include <stdlib.h>
import "C"

type noCopy struct{}

func (*noCopy) Lock()   {}
func (*noCopy) UNlock() {}

func ExecuteProcess(mainArgs *capi.CMainArgsT, app *capi.CAppT) {

	var sandBoxInfo unsafe.Pointer

	code := capi.ExecuteProcess(mainArgs, app, sandBoxInfo)
	if code >= 0 {
		os.Exit(int(code))
	}
}

func Initialize(mainArgs *capi.CMainArgsT, settings *capi.CSettingsT, app *capi.CAppT) {

	capi.Initialize(mainArgs, settings, app, nil)
}

func PostElementGetBytes(e *capi.CPostDataElementT) (bytes []byte) {
	if e.GetType() != capi.PdeTypeBytes {
		return bytes
	}
	n := e.GetBytesCount()
	cb := C.malloc(C.size_t(n))
	defer C.free(cb)

	bSize := e.GetBytes(n, cb)
	return C.GoBytes(cb, C.int(bSize))
}

type TaskFunc func()

var _ capi.CTaskTExecuteHandler = TaskFunc(func() {})

func (f TaskFunc) Execute(self *capi.CTaskT) {
	f()
}

func NewTask(f TaskFunc) *capi.CTaskT {
	task := capi.AllocCTaskT().Bind(f)
	return task
}
