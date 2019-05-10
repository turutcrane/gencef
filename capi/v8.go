package capi

import (
	"runtime"
	"sync"
	"unsafe"

	"github.com/pkg/errors"
)

// #include "cefingo.h"
// #include "cefingo_v8.h"
import "C"

type V8handler interface {
	///
	// Handle execution of the function identified by |name|. |object| is the
	// receiver ('this' object) of the function. |arguments| is the list of
	// arguments passed to the function. If execution succeeds set |retval| to the
	// function return value. If execution fails set |exception| to the exception
	// that will be thrown. Return true (1) if execution was handled.
	// https://github.com/chromiumembedded/cef/blob/3497/include/capi/cef_v8_capi.h#L158-L183
	///
	Execute(self *CV8handlerT,
		name string,
		object *CV8valueT,
		argumentsCount int,
		arguments []*CV8valueT,
		retval **CV8valueT,
		exception *string) bool
}

var v8handlers = struct {
	m         sync.Mutex
	v8handler map[*C.cef_v8handler_t]V8handler
}{
	sync.Mutex{},
	map[*C.cef_v8handler_t]V8handler{},
}

// AllocCV8arrayBufferReleaseCallbackT allocates CV8arrayBufferReleaseCallbackT and construct it
func AllocCV8arrayBufferReleaseCallbackT() *CV8arrayBufferReleaseCallbackT {
	// TODO Bind 関数 を書く??
	up := c_calloc(1, C.sizeof_cefingo_v8array_buffer_release_callback_wrapper_t, "L65:")

	cefp := C.cefingo_construct_v8array_buffer_release_callback(
		(*C.cefingo_v8array_buffer_release_callback_wrapper_t)(up))

	return newCV8arrayBufferReleaseCallbackT(cefp)
}

//export cefingo_v8array_buffer_release_callback_release_buffer
func cefingo_v8array_buffer_release_callback_release_buffer(self *C.cef_v8array_buffer_release_callback_t, buffer unsafe.Pointer) {
	Tracef(unsafe.Pointer(buffer), "L46:")
	// C.free(buffer) // ??
}

func V8valueCreateUndefined() *CV8valueT {
	v := C.cef_v8value_create_undefined()
	return newCV8valueT(v)
}

func V8valueCreateNull() *CV8valueT {
	v := C.cef_v8value_create_null()
	return newCV8valueT(v)
}

func V8valueCreateBool(b bool) *CV8valueT {
	var i int
	if b {
		i = 1
	}
	v := C.cef_v8value_create_bool(C.int(i))
	return newCV8valueT(v)
}

func V8valueCreateInt(i int) *CV8valueT {
	v := C.cef_v8value_create_int(C.int32(i))
	return newCV8valueT(v)
}

func V8valueCreateUint(u uint) *CV8valueT {
	v := C.cef_v8value_create_uint(C.uint32(u))
	return newCV8valueT(v)
}

func V8valueCreateDouble(f float64) *CV8valueT {
	v := C.cef_v8value_create_double(C.double(f))
	return newCV8valueT(v)
}

func V8valueCreateDate(d CTimeT) *CV8valueT {
	v := C.cef_v8value_create_date((*C.cef_time_t)(&d))
	return newCV8valueT(v)
}

func V8valueCreateString(s string) (val *CV8valueT) {
	cef_string := create_cef_string(s)
	defer clear_cef_string(cef_string)

	v := C.cef_v8value_create_string(cef_string)
	return newCV8valueT(v)
}

func V8valueCreateStringFromByteArray(b []byte) (val *CV8valueT) {
	cef_string := create_cef_string_from_byte_array(b)
	defer clear_cef_string(cef_string)

	v := C.cef_v8value_create_string(cef_string)
	return newCV8valueT(v)
}

func V8valueCreateObject(accessor *CV8accessorT, interceptor *CV8interceptorT) *CV8valueT {
	o := C.cef_v8value_create_object(
		(*C.cef_v8accessor_t)(accessor), (*C.cef_v8interceptor_t)(interceptor))
	return newCV8valueT(o)
}

func V8valueCreateArrayBuffer(buffer []byte) *CV8valueT {
	release_callback := AllocCV8arrayBufferReleaseCallbackT()

	// buf := [100]byte{}
	cbytes := C.CBytes(buffer[:])
	buffer_len := (C.size_t)(len(buffer[:]))
	BaseAddRef(release_callback.p_v8array_buffer_release_callback)
	ab := C.cef_v8value_create_array_buffer(
		cbytes,
		buffer_len,
		release_callback.p_v8array_buffer_release_callback,
	)
	return newCV8valueT(ab)
}

func (self *CV8valueT) IsValid() bool {
	status := C.cefingo_v8value_is_valid(self.p_v8value)
	return status == 1
}

func (self *CV8valueT) IsUndefined() bool {
	status := C.cefingo_v8value_is_undefined(self.p_v8value)
	return status == 1
}

func (self *CV8valueT) IsNull() bool {
	status := C.cefingo_v8value_is_null(self.p_v8value)
	return status == 1
}

func (self *CV8valueT) IsBool() bool {
	status := C.cefingo_v8value_is_bool(self.p_v8value)
	return status == 1
}

func (self *CV8valueT) IsInt() bool {
	status := C.cefingo_v8value_is_int(self.p_v8value)
	return status == 1
}

func (self *CV8valueT) IsUint() bool {
	status := C.cefingo_v8value_is_uint(self.p_v8value)
	return status == 1
}

func (self *CV8valueT) IsDouble() bool {
	status := C.cefingo_v8value_is_double(self.p_v8value)
	return status == 1
}

func (self *CV8valueT) IsDate() bool {
	status := C.cefingo_v8value_is_date(self.p_v8value)
	return status == 1
}

func (self *CV8valueT) IsString() bool {
	status := C.cefingo_v8value_is_string(self.p_v8value)
	return status == 1
}

func (self *CV8valueT) IsObject() bool {
	status := C.cefingo_v8value_is_object(self.p_v8value)
	return status == 1
}

func (self *CV8valueT) IsArray() bool {
	status := C.cefingo_v8value_is_array(self.p_v8value)
	return status == 1
}

func (self *CV8valueT) IsArrayBuffer() bool {
	status := C.cefingo_v8value_is_array_buffer(self.p_v8value)
	return status == 1
}

func (self *CV8valueT) IsFunction() bool {
	status := C.cefingo_v8value_is_function(self.p_v8value)
	return status == 1
}

func (self *CV8valueT) IsSame(that *CV8valueT) bool {
	BaseAddRef(that.p_v8value)
	status := C.cefingo_v8value_is_same(self.p_v8value, that.p_v8value)
	return status == 1
}

func (self *CV8valueT) GetBoolValue() (v bool) {
	b := C.cefingo_v8value_get_bool_value(self.p_v8value)
	if b != 0 {
		v = true
	}
	return v
}

func (self *CV8valueT) GetIntValue() (v int32) {
	v = int32(C.cefingo_v8value_get_int_value(self.p_v8value))
	return v
}

func (self *CV8valueT) GetUintValue() (v uint32) {
	v = uint32(C.cefingo_v8value_get_uint_value(self.p_v8value))
	return v
}

func (self *CV8valueT) GetDoubleValue() (v float64) {
	v = float64(C.cefingo_v8value_get_double_value(self.p_v8value))
	return v
}

func (self *CV8valueT) GetDateValue() (v CTimeT) {
	v = CTimeT(C.cefingo_v8value_get_date_value(self.p_v8value))
	return v
}

func (self *CV8valueT) GetStringValue() (s string) {
	usfs := C.cefingo_v8value_get_string_value(self.p_v8value)
	s = string_from_cef_string((*C.cef_string_t)(usfs))
	if usfs != nil {
		C.cef_string_userfree_free(usfs)
	}
	return s
}

func (self *CV8valueT) HasException() (ret bool) {
	v := C.cefingo_v8value_has_exception(self.p_v8value)
	return v == 1
}

func (self *CV8valueT) GetException() *CV8exceptionT {
	e := C.cefingo_v8value_get_exception(self.p_v8value)
	return newCV8exceptionT(e)
}

func (self *CV8valueT) ClearException() (ret bool) {
	v := C.cefingo_v8value_clear_exception(self.p_v8value)
	return v == 1
}

func (self *CV8valueT) HasValueBykey(key string) bool {
	key_string := create_cef_string(key)
	defer clear_cef_string(key_string)

	status := C.cefingo_v8value_has_value_bykey(self.p_v8value, key_string)
	return (status == 1)
}

func (self *CV8valueT) DeleteValueBykey(key string) bool {
	key_string := create_cef_string(key)
	defer clear_cef_string(key_string)

	status := C.cefingo_v8value_delete_value_bykey(self.p_v8value, key_string)
	return (status == 1)
}

func (self *CV8valueT) GetValueBykey(key string) (value *CV8valueT) {
	key_string := create_cef_string(key)
	defer clear_cef_string(key_string)

	v := C.cefingo_v8value_get_value_bykey(self.p_v8value, key_string)
	if v != nil {
		value = newCV8valueT(v)
	}
	return value
}

func (self *CV8valueT) SetValueBykey(key string, value *CV8valueT) bool {
	key_string := create_cef_string(key)
	defer clear_cef_string(key_string)

	BaseAddRef(value.p_v8value)
	status := C.cefingo_v8context_set_value_bykey(self.p_v8value,
		key_string, value.p_v8value, C.V8_PROPERTY_ATTRIBUTE_NONE)
	return status == 1
}

func (self *CV8valueT) GetFunctionName() (s string) {
	if self.IsFunction() {
		usfs := C.cefingo_v8value_get_function_name(self.p_v8value)
		if usfs != nil {
			s = string_from_cef_string((*C.cef_string_t)(usfs))
			C.cef_string_userfree_free(usfs)
		}
	}
	return s
}

func (self *CV8valueT) ExecuteFunction(
	this *CV8valueT,
	argumentsCount int,
	arguments []*CV8valueT,
) (val *CV8valueT, err error) {

	if !self.IsFunction() {
		cause := errors.Errorf("Object is Not Function")
		return nil, cause
	}
	cargs := C.calloc((C.size_t)(argumentsCount), (C.size_t)(unsafe.Sizeof(this)))
	slice := (*[1 << 30]*C.cef_v8value_t)(cargs)[:argumentsCount:argumentsCount]

	BaseAddRef(this.p_v8value)
	for i, v := range arguments {
		BaseAddRef(v.p_v8value)
		slice[i] = v.p_v8value
	}
	v := C.cefingo_v8value_execute_function(
		self.p_v8value,
		this.p_v8value,
		(C.size_t)(argumentsCount),
		(**C.cef_v8value_t)(cargs))
	if v == nil {
		name := self.GetFunctionName()
		if this != nil && this.HasException() {
			e := this.GetException()
			m := e.GetMessage()
			err = errors.Errorf("%s returns NULL and has Exception: %s", name, m)
		} else {
			err = errors.Errorf("%s returns NULL", name)
		}
	} else if status := C.cefingo_v8value_is_valid(v); status == 1 {
		val = newCV8valueT(v)
	}
	return val, err
}

// V8valueCreateFunction create V8 function
///
// Create a new cef_v8value_t object of type function. This function should only
// be called from within the scope of a cef_render_process_handler_t,
// cef_v8handler_t or cef_v8accessor_t callback, or in combination with calling
// enter() and exit() on a stored cef_v8context_t reference.
// https://github.com/chromiumembedded/cef/blob/3497/include/capi/cef_v8_capi.h#L812-L819
///
func V8valueCreateFunction(name string, handler *CV8handlerT) (function *CV8valueT) {
	cef_name := create_cef_string(name)
	defer clear_cef_string(cef_name)

	BaseAddRef(handler.p_v8handler)
	return newCV8valueT(C.cef_v8value_create_function(cef_name, handler.p_v8handler))
}

// AllocCV8handlerT allocates CV8handlerT and construct it
func AllocCV8handlerT() *CV8handlerT {
	up := c_calloc(1, C.sizeof_cefingo_v8handler_wrapper_t, "L393:")
	cefp := C.cefingo_construct_v8handler((*C.cefingo_v8handler_wrapper_t)(up))

	registerDeassocer(up, DeassocFunc(func() {
		// Do not have reference to cApp itself in DeassocFunc,
		// or app is never GCed.
		Tracef(up, "L67:")
		v8handlers.m.Lock()
		defer v8handlers.m.Unlock()

		delete(v8handlers.v8handler, cefp)
	}))

	return newCV8handlerT(cefp)
}

func (v8handler *CV8handlerT) Bind(handler V8handler) *CV8handlerT {
	v8hp := v8handler.p_v8handler
	v8handlers.m.Lock()
	v8handlers.v8handler[v8hp] = handler
	v8handlers.m.Unlock()

	if accessor, ok := handler.(CV8handlerTAccessor); ok {
		accessor.SetCV8handlerT(v8handler)
		Logf("L109:")
	}

	return v8handler
}

// Excute Handler
//export cefingo_v8handler_execute
func cefingo_v8handler_execute(self *C.cef_v8handler_t,
	name *C.cef_string_t,
	object *C.cef_v8value_t,
	argumentsCount C.size_t,
	arguments **C.cef_v8value_t,
	retval **C.cef_v8value_t,
	exception *C.cef_string_t,
) (ret C.int) {
	goname := string_from_cef_string(name)
	v8handlers.m.Lock()
	handler := v8handlers.v8handler[self]
	v8handlers.m.Unlock()
	if handler == nil {
		Logf("L121: No V8 Execute Handler")
		ret = 0
	} else {
		var slice []*CV8valueT
		if arguments != nil {
			s := (*[1 << 30]*C.cef_v8value_t)(unsafe.Pointer(arguments))[:argumentsCount:argumentsCount]
			for _, v := range s {
				slice = append(slice, newCV8valueT(v))
			}
		}

		runtime.LockOSThread()
		defer runtime.UnlockOSThread()

		var exc string
		var v8v *CV8valueT
		if handler.Execute(newCV8handlerT(self), goname,
			newCV8valueT(object),
			(int)(argumentsCount), slice, &v8v, &exc) {
			// Is required release of member of arguments ?
			if v8v != nil {
				*retval = v8v.p_v8value
				BaseAddRef(*retval)
			}
			set_cef_string(exception, exc)
			ret = 1
		} else {
			ret = 0
		}
	}
	return ret
}

func V8contextInContext() bool {
	// Returns true (1) if V8 is currently inside a context.
	inContext := C.cef_v8context_in_context()
	Logf("L150: %d", inContext)
	return (inContext == 1)
}

func V8contextGetEnterdContext() (context *CV8contextT) {
	c := C.cef_v8context_get_entered_context()
	context = newCV8contextT(c)
	return context
}

func (self *CV8contextT) IsValid() bool {
	status := C.cefingo_v8context_is_valid(self.p_v8context)
	return status == 1
}

func (self *CV8contextT) GetBrowser() *CBrowserT {
	b := C.cefingo_v8context_get_browser(self.p_v8context)
	return newCBrowserT(b)
}

func (self *CV8contextT) GetGlobal() *CV8valueT {
	g := C.cefingo_v8context_get_global(self.p_v8context)
	return newCV8valueT(g)
}

func (self *CV8contextT) GetFrame() *CFrameT {
	f := C.cefingo_v8context_get_frame(self.p_v8context)
	return newCFrameT(f)
}

func (self *CV8contextT) Enter() bool {
	runtime.LockOSThread()
	c := C.cefingo_v8context_enter(self.p_v8context)
	return (c == 1)
}

func (self *CV8contextT) Exit() bool {
	c := C.cefingo_v8context_exit(self.p_v8context)
	runtime.UnlockOSThread()
	return (c == 1)
}

func (self *CV8contextT) IsSame(that *CV8contextT) bool {
	BaseAddRef(that.p_v8context)
	s := C.cefingo_v8context_is_same(
		self.p_v8context,
		that.p_v8context,
	)
	return s == 1
}

func (self *CV8contextT) Eval(code string, retval **CV8valueT, e **CV8exceptionT) (ret bool) {
	s := create_cef_string(code)
	defer clear_cef_string(s)
	var r *C.cef_v8value_t
	var exc *C.cef_v8exception_t
	status := C.cefingo_v8context_eval(
		self.p_v8context, s, nil, 0,
		&r, &exc)
	ret = (status == 1)
	if ret {
		*retval = newCV8valueT(r)
	} else {
		// if exc != nil {
		*e = newCV8exceptionT(exc)
		// BaseAddRef(*e)
		// }
	}

	return ret
}

func (self *CV8exceptionT) GetMessage() (msg string) {
	usfs := C.cefingo_v8exception_get_message(self.p_v8exception)
	msg = string_from_cef_string((*C.cef_string_t)(usfs))
	C.cef_string_userfree_free(usfs)
	return msg
}
