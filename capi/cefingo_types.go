// Code generated by "gen_cef_types.go" DO NOT EDIT.

package capi

import (
	"runtime"
	"unsafe"
)

// #include "cefingo.h"
import "C"

// Go type for cef_app_t
type CAppT struct {
	p_app *C.cef_app_t
}

type RefToCAppT struct {
	p_app *CAppT
}

type CAppTAccessor interface {
	GetCAppT() *CAppT
	SetCAppT(*CAppT)
}

func (r RefToCAppT) GetCAppT() *CAppT {
	return r.p_app
}

func (r *RefToCAppT) SetCAppT(p *CAppT) {
	r.p_app = p
}

// Go type CAppT wraps cef type *C.cef_app_t
func newCAppT(p *C.cef_app_t) *CAppT {
	Tracef(unsafe.Pointer(p), "T37:")
	BaseAddRef(p)
	go_app := CAppT{p}
	runtime.SetFinalizer(&go_app, func(g *CAppT) {
		Tracef(unsafe.Pointer(g.p_app), "T41:")
		BaseRelease(g.p_app)
	})
	return &go_app
}

// *C.cef_app_t has refCounted interface
func (p *C.cef_app_t) cast_to_p_base_ref_counted_t() *C.cef_base_ref_counted_t {
	return (*C.cef_base_ref_counted_t)(unsafe.Pointer(p))
}

// Go type for cef_binary_value_t
type CBinaryValueT struct {
	p_binary_value *C.cef_binary_value_t
}

type RefToCBinaryValueT struct {
	p_binary_value *CBinaryValueT
}

type CBinaryValueTAccessor interface {
	GetCBinaryValueT() *CBinaryValueT
	SetCBinaryValueT(*CBinaryValueT)
}

func (r RefToCBinaryValueT) GetCBinaryValueT() *CBinaryValueT {
	return r.p_binary_value
}

func (r *RefToCBinaryValueT) SetCBinaryValueT(p *CBinaryValueT) {
	r.p_binary_value = p
}

// Go type CBinaryValueT wraps cef type *C.cef_binary_value_t
func newCBinaryValueT(p *C.cef_binary_value_t) *CBinaryValueT {
	Tracef(unsafe.Pointer(p), "T76:")
	BaseAddRef(p)
	go_binary_value := CBinaryValueT{p}
	runtime.SetFinalizer(&go_binary_value, func(g *CBinaryValueT) {
		Tracef(unsafe.Pointer(g.p_binary_value), "T80:")
		BaseRelease(g.p_binary_value)
	})
	return &go_binary_value
}

// *C.cef_binary_value_t has refCounted interface
func (p *C.cef_binary_value_t) cast_to_p_base_ref_counted_t() *C.cef_base_ref_counted_t {
	return (*C.cef_base_ref_counted_t)(unsafe.Pointer(p))
}

// Go type for cef_browser_t
type CBrowserT struct {
	p_browser *C.cef_browser_t
}

type RefToCBrowserT struct {
	p_browser *CBrowserT
}

type CBrowserTAccessor interface {
	GetCBrowserT() *CBrowserT
	SetCBrowserT(*CBrowserT)
}

func (r RefToCBrowserT) GetCBrowserT() *CBrowserT {
	return r.p_browser
}

func (r *RefToCBrowserT) SetCBrowserT(p *CBrowserT) {
	r.p_browser = p
}

// Go type CBrowserT wraps cef type *C.cef_browser_t
func newCBrowserT(p *C.cef_browser_t) *CBrowserT {
	Tracef(unsafe.Pointer(p), "T115:")
	BaseAddRef(p)
	go_browser := CBrowserT{p}
	runtime.SetFinalizer(&go_browser, func(g *CBrowserT) {
		Tracef(unsafe.Pointer(g.p_browser), "T119:")
		BaseRelease(g.p_browser)
	})
	return &go_browser
}

// *C.cef_browser_t has refCounted interface
func (p *C.cef_browser_t) cast_to_p_base_ref_counted_t() *C.cef_base_ref_counted_t {
	return (*C.cef_base_ref_counted_t)(unsafe.Pointer(p))
}

// Go type for cef_browser_host_t
type CBrowserHostT struct {
	p_browser_host *C.cef_browser_host_t
}

type RefToCBrowserHostT struct {
	p_browser_host *CBrowserHostT
}

type CBrowserHostTAccessor interface {
	GetCBrowserHostT() *CBrowserHostT
	SetCBrowserHostT(*CBrowserHostT)
}

func (r RefToCBrowserHostT) GetCBrowserHostT() *CBrowserHostT {
	return r.p_browser_host
}

func (r *RefToCBrowserHostT) SetCBrowserHostT(p *CBrowserHostT) {
	r.p_browser_host = p
}

// Go type CBrowserHostT wraps cef type *C.cef_browser_host_t
func newCBrowserHostT(p *C.cef_browser_host_t) *CBrowserHostT {
	Tracef(unsafe.Pointer(p), "T154:")
	BaseAddRef(p)
	go_browser_host := CBrowserHostT{p}
	runtime.SetFinalizer(&go_browser_host, func(g *CBrowserHostT) {
		Tracef(unsafe.Pointer(g.p_browser_host), "T158:")
		BaseRelease(g.p_browser_host)
	})
	return &go_browser_host
}

// *C.cef_browser_host_t has refCounted interface
func (p *C.cef_browser_host_t) cast_to_p_base_ref_counted_t() *C.cef_base_ref_counted_t {
	return (*C.cef_base_ref_counted_t)(unsafe.Pointer(p))
}

// Go type for cef_browser_process_handler_t
type CBrowserProcessHandlerT struct {
	p_browser_process_handler *C.cef_browser_process_handler_t
}

type RefToCBrowserProcessHandlerT struct {
	p_browser_process_handler *CBrowserProcessHandlerT
}

type CBrowserProcessHandlerTAccessor interface {
	GetCBrowserProcessHandlerT() *CBrowserProcessHandlerT
	SetCBrowserProcessHandlerT(*CBrowserProcessHandlerT)
}

func (r RefToCBrowserProcessHandlerT) GetCBrowserProcessHandlerT() *CBrowserProcessHandlerT {
	return r.p_browser_process_handler
}

func (r *RefToCBrowserProcessHandlerT) SetCBrowserProcessHandlerT(p *CBrowserProcessHandlerT) {
	r.p_browser_process_handler = p
}

// Go type CBrowserProcessHandlerT wraps cef type *C.cef_browser_process_handler_t
func newCBrowserProcessHandlerT(p *C.cef_browser_process_handler_t) *CBrowserProcessHandlerT {
	Tracef(unsafe.Pointer(p), "T193:")
	BaseAddRef(p)
	go_browser_process_handler := CBrowserProcessHandlerT{p}
	runtime.SetFinalizer(&go_browser_process_handler, func(g *CBrowserProcessHandlerT) {
		Tracef(unsafe.Pointer(g.p_browser_process_handler), "T197:")
		BaseRelease(g.p_browser_process_handler)
	})
	return &go_browser_process_handler
}

// *C.cef_browser_process_handler_t has refCounted interface
func (p *C.cef_browser_process_handler_t) cast_to_p_base_ref_counted_t() *C.cef_base_ref_counted_t {
	return (*C.cef_base_ref_counted_t)(unsafe.Pointer(p))
}

// Go type for cef_callback_t
type CCallbackT struct {
	p_callback *C.cef_callback_t
}

type RefToCCallbackT struct {
	p_callback *CCallbackT
}

type CCallbackTAccessor interface {
	GetCCallbackT() *CCallbackT
	SetCCallbackT(*CCallbackT)
}

func (r RefToCCallbackT) GetCCallbackT() *CCallbackT {
	return r.p_callback
}

func (r *RefToCCallbackT) SetCCallbackT(p *CCallbackT) {
	r.p_callback = p
}

// Go type CCallbackT wraps cef type *C.cef_callback_t
func newCCallbackT(p *C.cef_callback_t) *CCallbackT {
	Tracef(unsafe.Pointer(p), "T232:")
	BaseAddRef(p)
	go_callback := CCallbackT{p}
	runtime.SetFinalizer(&go_callback, func(g *CCallbackT) {
		Tracef(unsafe.Pointer(g.p_callback), "T236:")
		BaseRelease(g.p_callback)
	})
	return &go_callback
}

// *C.cef_callback_t has refCounted interface
func (p *C.cef_callback_t) cast_to_p_base_ref_counted_t() *C.cef_base_ref_counted_t {
	return (*C.cef_base_ref_counted_t)(unsafe.Pointer(p))
}

// Go type for cef_client_t
type CClientT struct {
	p_client *C.cef_client_t
}

type RefToCClientT struct {
	p_client *CClientT
}

type CClientTAccessor interface {
	GetCClientT() *CClientT
	SetCClientT(*CClientT)
}

func (r RefToCClientT) GetCClientT() *CClientT {
	return r.p_client
}

func (r *RefToCClientT) SetCClientT(p *CClientT) {
	r.p_client = p
}

// Go type CClientT wraps cef type *C.cef_client_t
func newCClientT(p *C.cef_client_t) *CClientT {
	Tracef(unsafe.Pointer(p), "T271:")
	BaseAddRef(p)
	go_client := CClientT{p}
	runtime.SetFinalizer(&go_client, func(g *CClientT) {
		Tracef(unsafe.Pointer(g.p_client), "T275:")
		BaseRelease(g.p_client)
	})
	return &go_client
}

// *C.cef_client_t has refCounted interface
func (p *C.cef_client_t) cast_to_p_base_ref_counted_t() *C.cef_base_ref_counted_t {
	return (*C.cef_base_ref_counted_t)(unsafe.Pointer(p))
}

// Go type for cef_dictionary_value_t
type CDictionaryValueT struct {
	p_dictionary_value *C.cef_dictionary_value_t
}

type RefToCDictionaryValueT struct {
	p_dictionary_value *CDictionaryValueT
}

type CDictionaryValueTAccessor interface {
	GetCDictionaryValueT() *CDictionaryValueT
	SetCDictionaryValueT(*CDictionaryValueT)
}

func (r RefToCDictionaryValueT) GetCDictionaryValueT() *CDictionaryValueT {
	return r.p_dictionary_value
}

func (r *RefToCDictionaryValueT) SetCDictionaryValueT(p *CDictionaryValueT) {
	r.p_dictionary_value = p
}

// Go type CDictionaryValueT wraps cef type *C.cef_dictionary_value_t
func newCDictionaryValueT(p *C.cef_dictionary_value_t) *CDictionaryValueT {
	Tracef(unsafe.Pointer(p), "T310:")
	BaseAddRef(p)
	go_dictionary_value := CDictionaryValueT{p}
	runtime.SetFinalizer(&go_dictionary_value, func(g *CDictionaryValueT) {
		Tracef(unsafe.Pointer(g.p_dictionary_value), "T314:")
		BaseRelease(g.p_dictionary_value)
	})
	return &go_dictionary_value
}

// *C.cef_dictionary_value_t has refCounted interface
func (p *C.cef_dictionary_value_t) cast_to_p_base_ref_counted_t() *C.cef_base_ref_counted_t {
	return (*C.cef_base_ref_counted_t)(unsafe.Pointer(p))
}

// Go type for cef_frame_t
type CFrameT struct {
	p_frame *C.cef_frame_t
}

type RefToCFrameT struct {
	p_frame *CFrameT
}

type CFrameTAccessor interface {
	GetCFrameT() *CFrameT
	SetCFrameT(*CFrameT)
}

func (r RefToCFrameT) GetCFrameT() *CFrameT {
	return r.p_frame
}

func (r *RefToCFrameT) SetCFrameT(p *CFrameT) {
	r.p_frame = p
}

// Go type CFrameT wraps cef type *C.cef_frame_t
func newCFrameT(p *C.cef_frame_t) *CFrameT {
	Tracef(unsafe.Pointer(p), "T349:")
	BaseAddRef(p)
	go_frame := CFrameT{p}
	runtime.SetFinalizer(&go_frame, func(g *CFrameT) {
		Tracef(unsafe.Pointer(g.p_frame), "T353:")
		BaseRelease(g.p_frame)
	})
	return &go_frame
}

// *C.cef_frame_t has refCounted interface
func (p *C.cef_frame_t) cast_to_p_base_ref_counted_t() *C.cef_base_ref_counted_t {
	return (*C.cef_base_ref_counted_t)(unsafe.Pointer(p))
}

// Go type for cef_life_span_handler_t
type CLifeSpanHandlerT struct {
	p_life_span_handler *C.cef_life_span_handler_t
}

type RefToCLifeSpanHandlerT struct {
	p_life_span_handler *CLifeSpanHandlerT
}

type CLifeSpanHandlerTAccessor interface {
	GetCLifeSpanHandlerT() *CLifeSpanHandlerT
	SetCLifeSpanHandlerT(*CLifeSpanHandlerT)
}

func (r RefToCLifeSpanHandlerT) GetCLifeSpanHandlerT() *CLifeSpanHandlerT {
	return r.p_life_span_handler
}

func (r *RefToCLifeSpanHandlerT) SetCLifeSpanHandlerT(p *CLifeSpanHandlerT) {
	r.p_life_span_handler = p
}

// Go type CLifeSpanHandlerT wraps cef type *C.cef_life_span_handler_t
func newCLifeSpanHandlerT(p *C.cef_life_span_handler_t) *CLifeSpanHandlerT {
	Tracef(unsafe.Pointer(p), "T388:")
	BaseAddRef(p)
	go_life_span_handler := CLifeSpanHandlerT{p}
	runtime.SetFinalizer(&go_life_span_handler, func(g *CLifeSpanHandlerT) {
		Tracef(unsafe.Pointer(g.p_life_span_handler), "T392:")
		BaseRelease(g.p_life_span_handler)
	})
	return &go_life_span_handler
}

// *C.cef_life_span_handler_t has refCounted interface
func (p *C.cef_life_span_handler_t) cast_to_p_base_ref_counted_t() *C.cef_base_ref_counted_t {
	return (*C.cef_base_ref_counted_t)(unsafe.Pointer(p))
}

// Go type for cef_list_value_t
type CListValueT struct {
	p_list_value *C.cef_list_value_t
}

type RefToCListValueT struct {
	p_list_value *CListValueT
}

type CListValueTAccessor interface {
	GetCListValueT() *CListValueT
	SetCListValueT(*CListValueT)
}

func (r RefToCListValueT) GetCListValueT() *CListValueT {
	return r.p_list_value
}

func (r *RefToCListValueT) SetCListValueT(p *CListValueT) {
	r.p_list_value = p
}

// Go type CListValueT wraps cef type *C.cef_list_value_t
func newCListValueT(p *C.cef_list_value_t) *CListValueT {
	Tracef(unsafe.Pointer(p), "T427:")
	BaseAddRef(p)
	go_list_value := CListValueT{p}
	runtime.SetFinalizer(&go_list_value, func(g *CListValueT) {
		Tracef(unsafe.Pointer(g.p_list_value), "T431:")
		BaseRelease(g.p_list_value)
	})
	return &go_list_value
}

// *C.cef_list_value_t has refCounted interface
func (p *C.cef_list_value_t) cast_to_p_base_ref_counted_t() *C.cef_base_ref_counted_t {
	return (*C.cef_base_ref_counted_t)(unsafe.Pointer(p))
}

// Go type for cef_load_handler_t
type CLoadHandlerT struct {
	p_load_handler *C.cef_load_handler_t
}

type RefToCLoadHandlerT struct {
	p_load_handler *CLoadHandlerT
}

type CLoadHandlerTAccessor interface {
	GetCLoadHandlerT() *CLoadHandlerT
	SetCLoadHandlerT(*CLoadHandlerT)
}

func (r RefToCLoadHandlerT) GetCLoadHandlerT() *CLoadHandlerT {
	return r.p_load_handler
}

func (r *RefToCLoadHandlerT) SetCLoadHandlerT(p *CLoadHandlerT) {
	r.p_load_handler = p
}

// Go type CLoadHandlerT wraps cef type *C.cef_load_handler_t
func newCLoadHandlerT(p *C.cef_load_handler_t) *CLoadHandlerT {
	Tracef(unsafe.Pointer(p), "T466:")
	BaseAddRef(p)
	go_load_handler := CLoadHandlerT{p}
	runtime.SetFinalizer(&go_load_handler, func(g *CLoadHandlerT) {
		Tracef(unsafe.Pointer(g.p_load_handler), "T470:")
		BaseRelease(g.p_load_handler)
	})
	return &go_load_handler
}

// *C.cef_load_handler_t has refCounted interface
func (p *C.cef_load_handler_t) cast_to_p_base_ref_counted_t() *C.cef_base_ref_counted_t {
	return (*C.cef_base_ref_counted_t)(unsafe.Pointer(p))
}

// Go type for cef_process_message_t
type CProcessMessageT struct {
	p_process_message *C.cef_process_message_t
}

type RefToCProcessMessageT struct {
	p_process_message *CProcessMessageT
}

type CProcessMessageTAccessor interface {
	GetCProcessMessageT() *CProcessMessageT
	SetCProcessMessageT(*CProcessMessageT)
}

func (r RefToCProcessMessageT) GetCProcessMessageT() *CProcessMessageT {
	return r.p_process_message
}

func (r *RefToCProcessMessageT) SetCProcessMessageT(p *CProcessMessageT) {
	r.p_process_message = p
}

// Go type CProcessMessageT wraps cef type *C.cef_process_message_t
func newCProcessMessageT(p *C.cef_process_message_t) *CProcessMessageT {
	Tracef(unsafe.Pointer(p), "T505:")
	BaseAddRef(p)
	go_process_message := CProcessMessageT{p}
	runtime.SetFinalizer(&go_process_message, func(g *CProcessMessageT) {
		Tracef(unsafe.Pointer(g.p_process_message), "T509:")
		BaseRelease(g.p_process_message)
	})
	return &go_process_message
}

// *C.cef_process_message_t has refCounted interface
func (p *C.cef_process_message_t) cast_to_p_base_ref_counted_t() *C.cef_base_ref_counted_t {
	return (*C.cef_base_ref_counted_t)(unsafe.Pointer(p))
}

// Go type for cef_render_process_handler_t
type CRenderProcessHandlerT struct {
	p_render_process_handler *C.cef_render_process_handler_t
}

type RefToCRenderProcessHandlerT struct {
	p_render_process_handler *CRenderProcessHandlerT
}

type CRenderProcessHandlerTAccessor interface {
	GetCRenderProcessHandlerT() *CRenderProcessHandlerT
	SetCRenderProcessHandlerT(*CRenderProcessHandlerT)
}

func (r RefToCRenderProcessHandlerT) GetCRenderProcessHandlerT() *CRenderProcessHandlerT {
	return r.p_render_process_handler
}

func (r *RefToCRenderProcessHandlerT) SetCRenderProcessHandlerT(p *CRenderProcessHandlerT) {
	r.p_render_process_handler = p
}

// Go type CRenderProcessHandlerT wraps cef type *C.cef_render_process_handler_t
func newCRenderProcessHandlerT(p *C.cef_render_process_handler_t) *CRenderProcessHandlerT {
	Tracef(unsafe.Pointer(p), "T544:")
	BaseAddRef(p)
	go_render_process_handler := CRenderProcessHandlerT{p}
	runtime.SetFinalizer(&go_render_process_handler, func(g *CRenderProcessHandlerT) {
		Tracef(unsafe.Pointer(g.p_render_process_handler), "T548:")
		BaseRelease(g.p_render_process_handler)
	})
	return &go_render_process_handler
}

// *C.cef_render_process_handler_t has refCounted interface
func (p *C.cef_render_process_handler_t) cast_to_p_base_ref_counted_t() *C.cef_base_ref_counted_t {
	return (*C.cef_base_ref_counted_t)(unsafe.Pointer(p))
}

// Go type for cef_request_t
type CRequestT struct {
	p_request *C.cef_request_t
}

type RefToCRequestT struct {
	p_request *CRequestT
}

type CRequestTAccessor interface {
	GetCRequestT() *CRequestT
	SetCRequestT(*CRequestT)
}

func (r RefToCRequestT) GetCRequestT() *CRequestT {
	return r.p_request
}

func (r *RefToCRequestT) SetCRequestT(p *CRequestT) {
	r.p_request = p
}

// Go type CRequestT wraps cef type *C.cef_request_t
func newCRequestT(p *C.cef_request_t) *CRequestT {
	Tracef(unsafe.Pointer(p), "T583:")
	BaseAddRef(p)
	go_request := CRequestT{p}
	runtime.SetFinalizer(&go_request, func(g *CRequestT) {
		Tracef(unsafe.Pointer(g.p_request), "T587:")
		BaseRelease(g.p_request)
	})
	return &go_request
}

// *C.cef_request_t has refCounted interface
func (p *C.cef_request_t) cast_to_p_base_ref_counted_t() *C.cef_base_ref_counted_t {
	return (*C.cef_base_ref_counted_t)(unsafe.Pointer(p))
}

// Go type for cef_resource_handler_t
type CResourceHandlerT struct {
	p_resource_handler *C.cef_resource_handler_t
}

type RefToCResourceHandlerT struct {
	p_resource_handler *CResourceHandlerT
}

type CResourceHandlerTAccessor interface {
	GetCResourceHandlerT() *CResourceHandlerT
	SetCResourceHandlerT(*CResourceHandlerT)
}

func (r RefToCResourceHandlerT) GetCResourceHandlerT() *CResourceHandlerT {
	return r.p_resource_handler
}

func (r *RefToCResourceHandlerT) SetCResourceHandlerT(p *CResourceHandlerT) {
	r.p_resource_handler = p
}

// Go type CResourceHandlerT wraps cef type *C.cef_resource_handler_t
func newCResourceHandlerT(p *C.cef_resource_handler_t) *CResourceHandlerT {
	Tracef(unsafe.Pointer(p), "T622:")
	BaseAddRef(p)
	go_resource_handler := CResourceHandlerT{p}
	runtime.SetFinalizer(&go_resource_handler, func(g *CResourceHandlerT) {
		Tracef(unsafe.Pointer(g.p_resource_handler), "T626:")
		BaseRelease(g.p_resource_handler)
	})
	return &go_resource_handler
}

// *C.cef_resource_handler_t has refCounted interface
func (p *C.cef_resource_handler_t) cast_to_p_base_ref_counted_t() *C.cef_base_ref_counted_t {
	return (*C.cef_base_ref_counted_t)(unsafe.Pointer(p))
}

// Go type for cef_response_t
type CResponseT struct {
	p_response *C.cef_response_t
}

type RefToCResponseT struct {
	p_response *CResponseT
}

type CResponseTAccessor interface {
	GetCResponseT() *CResponseT
	SetCResponseT(*CResponseT)
}

func (r RefToCResponseT) GetCResponseT() *CResponseT {
	return r.p_response
}

func (r *RefToCResponseT) SetCResponseT(p *CResponseT) {
	r.p_response = p
}

// Go type CResponseT wraps cef type *C.cef_response_t
func newCResponseT(p *C.cef_response_t) *CResponseT {
	Tracef(unsafe.Pointer(p), "T661:")
	BaseAddRef(p)
	go_response := CResponseT{p}
	runtime.SetFinalizer(&go_response, func(g *CResponseT) {
		Tracef(unsafe.Pointer(g.p_response), "T665:")
		BaseRelease(g.p_response)
	})
	return &go_response
}

// *C.cef_response_t has refCounted interface
func (p *C.cef_response_t) cast_to_p_base_ref_counted_t() *C.cef_base_ref_counted_t {
	return (*C.cef_base_ref_counted_t)(unsafe.Pointer(p))
}

// Go type for cef_scheme_handler_factory_t
type CSchemeHandlerFactoryT struct {
	p_scheme_handler_factory *C.cef_scheme_handler_factory_t
}

type RefToCSchemeHandlerFactoryT struct {
	p_scheme_handler_factory *CSchemeHandlerFactoryT
}

type CSchemeHandlerFactoryTAccessor interface {
	GetCSchemeHandlerFactoryT() *CSchemeHandlerFactoryT
	SetCSchemeHandlerFactoryT(*CSchemeHandlerFactoryT)
}

func (r RefToCSchemeHandlerFactoryT) GetCSchemeHandlerFactoryT() *CSchemeHandlerFactoryT {
	return r.p_scheme_handler_factory
}

func (r *RefToCSchemeHandlerFactoryT) SetCSchemeHandlerFactoryT(p *CSchemeHandlerFactoryT) {
	r.p_scheme_handler_factory = p
}

// Go type CSchemeHandlerFactoryT wraps cef type *C.cef_scheme_handler_factory_t
func newCSchemeHandlerFactoryT(p *C.cef_scheme_handler_factory_t) *CSchemeHandlerFactoryT {
	Tracef(unsafe.Pointer(p), "T700:")
	BaseAddRef(p)
	go_scheme_handler_factory := CSchemeHandlerFactoryT{p}
	runtime.SetFinalizer(&go_scheme_handler_factory, func(g *CSchemeHandlerFactoryT) {
		Tracef(unsafe.Pointer(g.p_scheme_handler_factory), "T704:")
		BaseRelease(g.p_scheme_handler_factory)
	})
	return &go_scheme_handler_factory
}

// *C.cef_scheme_handler_factory_t has refCounted interface
func (p *C.cef_scheme_handler_factory_t) cast_to_p_base_ref_counted_t() *C.cef_base_ref_counted_t {
	return (*C.cef_base_ref_counted_t)(unsafe.Pointer(p))
}

// Go type for cef_run_file_dialog_callback_t
type CRunFileDialogCallbackT struct {
	p_run_file_dialog_callback *C.cef_run_file_dialog_callback_t
}

type RefToCRunFileDialogCallbackT struct {
	p_run_file_dialog_callback *CRunFileDialogCallbackT
}

type CRunFileDialogCallbackTAccessor interface {
	GetCRunFileDialogCallbackT() *CRunFileDialogCallbackT
	SetCRunFileDialogCallbackT(*CRunFileDialogCallbackT)
}

func (r RefToCRunFileDialogCallbackT) GetCRunFileDialogCallbackT() *CRunFileDialogCallbackT {
	return r.p_run_file_dialog_callback
}

func (r *RefToCRunFileDialogCallbackT) SetCRunFileDialogCallbackT(p *CRunFileDialogCallbackT) {
	r.p_run_file_dialog_callback = p
}

// Go type CRunFileDialogCallbackT wraps cef type *C.cef_run_file_dialog_callback_t
func newCRunFileDialogCallbackT(p *C.cef_run_file_dialog_callback_t) *CRunFileDialogCallbackT {
	Tracef(unsafe.Pointer(p), "T739:")
	BaseAddRef(p)
	go_run_file_dialog_callback := CRunFileDialogCallbackT{p}
	runtime.SetFinalizer(&go_run_file_dialog_callback, func(g *CRunFileDialogCallbackT) {
		Tracef(unsafe.Pointer(g.p_run_file_dialog_callback), "T743:")
		BaseRelease(g.p_run_file_dialog_callback)
	})
	return &go_run_file_dialog_callback
}

// *C.cef_run_file_dialog_callback_t has refCounted interface
func (p *C.cef_run_file_dialog_callback_t) cast_to_p_base_ref_counted_t() *C.cef_base_ref_counted_t {
	return (*C.cef_base_ref_counted_t)(unsafe.Pointer(p))
}

// Go type for cef_v8array_buffer_release_callback_t
type CV8arrayBufferReleaseCallbackT struct {
	p_v8array_buffer_release_callback *C.cef_v8array_buffer_release_callback_t
}

type RefToCV8arrayBufferReleaseCallbackT struct {
	p_v8array_buffer_release_callback *CV8arrayBufferReleaseCallbackT
}

type CV8arrayBufferReleaseCallbackTAccessor interface {
	GetCV8arrayBufferReleaseCallbackT() *CV8arrayBufferReleaseCallbackT
	SetCV8arrayBufferReleaseCallbackT(*CV8arrayBufferReleaseCallbackT)
}

func (r RefToCV8arrayBufferReleaseCallbackT) GetCV8arrayBufferReleaseCallbackT() *CV8arrayBufferReleaseCallbackT {
	return r.p_v8array_buffer_release_callback
}

func (r *RefToCV8arrayBufferReleaseCallbackT) SetCV8arrayBufferReleaseCallbackT(p *CV8arrayBufferReleaseCallbackT) {
	r.p_v8array_buffer_release_callback = p
}

// Go type CV8arrayBufferReleaseCallbackT wraps cef type *C.cef_v8array_buffer_release_callback_t
func newCV8arrayBufferReleaseCallbackT(p *C.cef_v8array_buffer_release_callback_t) *CV8arrayBufferReleaseCallbackT {
	Tracef(unsafe.Pointer(p), "T778:")
	BaseAddRef(p)
	go_v8array_buffer_release_callback := CV8arrayBufferReleaseCallbackT{p}
	runtime.SetFinalizer(&go_v8array_buffer_release_callback, func(g *CV8arrayBufferReleaseCallbackT) {
		Tracef(unsafe.Pointer(g.p_v8array_buffer_release_callback), "T782:")
		BaseRelease(g.p_v8array_buffer_release_callback)
	})
	return &go_v8array_buffer_release_callback
}

// *C.cef_v8array_buffer_release_callback_t has refCounted interface
func (p *C.cef_v8array_buffer_release_callback_t) cast_to_p_base_ref_counted_t() *C.cef_base_ref_counted_t {
	return (*C.cef_base_ref_counted_t)(unsafe.Pointer(p))
}

// Go type for cef_v8context_t
type CV8contextT struct {
	p_v8context *C.cef_v8context_t
}

type RefToCV8contextT struct {
	p_v8context *CV8contextT
}

type CV8contextTAccessor interface {
	GetCV8contextT() *CV8contextT
	SetCV8contextT(*CV8contextT)
}

func (r RefToCV8contextT) GetCV8contextT() *CV8contextT {
	return r.p_v8context
}

func (r *RefToCV8contextT) SetCV8contextT(p *CV8contextT) {
	r.p_v8context = p
}

// Go type CV8contextT wraps cef type *C.cef_v8context_t
func newCV8contextT(p *C.cef_v8context_t) *CV8contextT {
	Tracef(unsafe.Pointer(p), "T817:")
	BaseAddRef(p)
	go_v8context := CV8contextT{p}
	runtime.SetFinalizer(&go_v8context, func(g *CV8contextT) {
		Tracef(unsafe.Pointer(g.p_v8context), "T821:")
		BaseRelease(g.p_v8context)
	})
	return &go_v8context
}

// *C.cef_v8context_t has refCounted interface
func (p *C.cef_v8context_t) cast_to_p_base_ref_counted_t() *C.cef_base_ref_counted_t {
	return (*C.cef_base_ref_counted_t)(unsafe.Pointer(p))
}

// Go type for cef_v8exception_t
type CV8exceptionT struct {
	p_v8exception *C.cef_v8exception_t
}

type RefToCV8exceptionT struct {
	p_v8exception *CV8exceptionT
}

type CV8exceptionTAccessor interface {
	GetCV8exceptionT() *CV8exceptionT
	SetCV8exceptionT(*CV8exceptionT)
}

func (r RefToCV8exceptionT) GetCV8exceptionT() *CV8exceptionT {
	return r.p_v8exception
}

func (r *RefToCV8exceptionT) SetCV8exceptionT(p *CV8exceptionT) {
	r.p_v8exception = p
}

// Go type CV8exceptionT wraps cef type *C.cef_v8exception_t
func newCV8exceptionT(p *C.cef_v8exception_t) *CV8exceptionT {
	Tracef(unsafe.Pointer(p), "T856:")
	BaseAddRef(p)
	go_v8exception := CV8exceptionT{p}
	runtime.SetFinalizer(&go_v8exception, func(g *CV8exceptionT) {
		Tracef(unsafe.Pointer(g.p_v8exception), "T860:")
		BaseRelease(g.p_v8exception)
	})
	return &go_v8exception
}

// *C.cef_v8exception_t has refCounted interface
func (p *C.cef_v8exception_t) cast_to_p_base_ref_counted_t() *C.cef_base_ref_counted_t {
	return (*C.cef_base_ref_counted_t)(unsafe.Pointer(p))
}

// Go type for cef_v8handler_t
type CV8handlerT struct {
	p_v8handler *C.cef_v8handler_t
}

type RefToCV8handlerT struct {
	p_v8handler *CV8handlerT
}

type CV8handlerTAccessor interface {
	GetCV8handlerT() *CV8handlerT
	SetCV8handlerT(*CV8handlerT)
}

func (r RefToCV8handlerT) GetCV8handlerT() *CV8handlerT {
	return r.p_v8handler
}

func (r *RefToCV8handlerT) SetCV8handlerT(p *CV8handlerT) {
	r.p_v8handler = p
}

// Go type CV8handlerT wraps cef type *C.cef_v8handler_t
func newCV8handlerT(p *C.cef_v8handler_t) *CV8handlerT {
	Tracef(unsafe.Pointer(p), "T895:")
	BaseAddRef(p)
	go_v8handler := CV8handlerT{p}
	runtime.SetFinalizer(&go_v8handler, func(g *CV8handlerT) {
		Tracef(unsafe.Pointer(g.p_v8handler), "T899:")
		BaseRelease(g.p_v8handler)
	})
	return &go_v8handler
}

// *C.cef_v8handler_t has refCounted interface
func (p *C.cef_v8handler_t) cast_to_p_base_ref_counted_t() *C.cef_base_ref_counted_t {
	return (*C.cef_base_ref_counted_t)(unsafe.Pointer(p))
}

// Go type for cef_v8value_t
type CV8valueT struct {
	p_v8value *C.cef_v8value_t
}

type RefToCV8valueT struct {
	p_v8value *CV8valueT
}

type CV8valueTAccessor interface {
	GetCV8valueT() *CV8valueT
	SetCV8valueT(*CV8valueT)
}

func (r RefToCV8valueT) GetCV8valueT() *CV8valueT {
	return r.p_v8value
}

func (r *RefToCV8valueT) SetCV8valueT(p *CV8valueT) {
	r.p_v8value = p
}

// Go type CV8valueT wraps cef type *C.cef_v8value_t
func newCV8valueT(p *C.cef_v8value_t) *CV8valueT {
	Tracef(unsafe.Pointer(p), "T934:")
	BaseAddRef(p)
	go_v8value := CV8valueT{p}
	runtime.SetFinalizer(&go_v8value, func(g *CV8valueT) {
		Tracef(unsafe.Pointer(g.p_v8value), "T938:")
		BaseRelease(g.p_v8value)
	})
	return &go_v8value
}

// *C.cef_v8value_t has refCounted interface
func (p *C.cef_v8value_t) cast_to_p_base_ref_counted_t() *C.cef_base_ref_counted_t {
	return (*C.cef_base_ref_counted_t)(unsafe.Pointer(p))
}

// Go type for cef_v8stack_trace_t
type CV8stackTraceT struct {
	p_v8stack_trace *C.cef_v8stack_trace_t
}

type RefToCV8stackTraceT struct {
	p_v8stack_trace *CV8stackTraceT
}

type CV8stackTraceTAccessor interface {
	GetCV8stackTraceT() *CV8stackTraceT
	SetCV8stackTraceT(*CV8stackTraceT)
}

func (r RefToCV8stackTraceT) GetCV8stackTraceT() *CV8stackTraceT {
	return r.p_v8stack_trace
}

func (r *RefToCV8stackTraceT) SetCV8stackTraceT(p *CV8stackTraceT) {
	r.p_v8stack_trace = p
}

// Go type CV8stackTraceT wraps cef type *C.cef_v8stack_trace_t
func newCV8stackTraceT(p *C.cef_v8stack_trace_t) *CV8stackTraceT {
	Tracef(unsafe.Pointer(p), "T973:")
	BaseAddRef(p)
	go_v8stack_trace := CV8stackTraceT{p}
	runtime.SetFinalizer(&go_v8stack_trace, func(g *CV8stackTraceT) {
		Tracef(unsafe.Pointer(g.p_v8stack_trace), "T977:")
		BaseRelease(g.p_v8stack_trace)
	})
	return &go_v8stack_trace
}

// *C.cef_v8stack_trace_t has refCounted interface
func (p *C.cef_v8stack_trace_t) cast_to_p_base_ref_counted_t() *C.cef_base_ref_counted_t {
	return (*C.cef_base_ref_counted_t)(unsafe.Pointer(p))
}

// Go type for cef_value_t
type CValueT struct {
	p_value *C.cef_value_t
}

type RefToCValueT struct {
	p_value *CValueT
}

type CValueTAccessor interface {
	GetCValueT() *CValueT
	SetCValueT(*CValueT)
}

func (r RefToCValueT) GetCValueT() *CValueT {
	return r.p_value
}

func (r *RefToCValueT) SetCValueT(p *CValueT) {
	r.p_value = p
}

// Go type CValueT wraps cef type *C.cef_value_t
func newCValueT(p *C.cef_value_t) *CValueT {
	Tracef(unsafe.Pointer(p), "T1012:")
	BaseAddRef(p)
	go_value := CValueT{p}
	runtime.SetFinalizer(&go_value, func(g *CValueT) {
		Tracef(unsafe.Pointer(g.p_value), "T1016:")
		BaseRelease(g.p_value)
	})
	return &go_value
}

// *C.cef_value_t has refCounted interface
func (p *C.cef_value_t) cast_to_p_base_ref_counted_t() *C.cef_base_ref_counted_t {
	return (*C.cef_base_ref_counted_t)(unsafe.Pointer(p))
}
