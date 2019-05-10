package capi

// #include "cefingo.h"
import "C"

type StringMap struct {
	Key   string
	Value string
}

///
// Set the response error code. This can be used by custom scheme handlers to
// return errors during initial request processing.
///
func (self *CResponseT) SetError(
	errorCode CErrorcodeT,
) {
	C.cefingo_response_set_error(self.p_response, C.cef_errorcode_t(errorCode))
}

///
// Set the response status code.
///
func (self *CResponseT) SetStatus(
	status int,
) {
	C.cefingo_response_set_status(self.p_response, C.int(status))
}

///
// Set the response status text.
///
func (self *CResponseT) SetStatusText(
	statusText string,
) {
	s := create_cef_string(statusText)
	defer clear_cef_string(s)
	C.cefingo_response_set_status_text(self.p_response, s)
}

///
// Set the response mime type.
///
func (self *CResponseT) SetMimeType(mimeType string) {
	s := create_cef_string(mimeType)
	defer clear_cef_string(s)
	C.cefingo_response_set_mime_type(self.p_response, s)
}

func (self *CResponseT) SetHeaderMap(headers []StringMap) {
	m := C.cef_string_multimap_alloc()
	defer C.cef_string_multimap_free(m)

	key := &C.cef_string_t{}
	value := &C.cef_string_t{}
	for _, v := range headers {
		set_cef_string(key, v.Key)
		set_cef_string(value, v.Value)
		C.cef_string_multimap_append(m, key, value)
	}
	defer clear_cef_string(key)
	defer clear_cef_string(value)

	C.cefingo_response_set_header_map(self.p_response, m)
}

func (self *CResponseT) DumpHeaders() {
	m := C.cef_string_multimap_alloc()
	defer C.cef_string_multimap_free(m)

	C.cefingo_response_get_header_map(self.p_response, m)
	size := C.cef_string_multimap_size(m)
	Logf("T49: size:%d", size)
	for i := C.size_t(0); i < size; i++ {
		k := C.cef_string_t{}
		C.cef_string_multimap_key(m, i, &k)
		key := string_from_cef_string(&k)
		Logf("T53: %d key:%s", i, key)
	}

}
