package cefingo

import (
	"unsafe"
)

// #include "cefingo.h"
import "C"

func ValueCreate() *CValueT {
	return (*CValueT)(C.cef_value_create())
}

func BinaryValueCreate(data []byte) *CBinaryValueT {
	return (*CBinaryValueT)(C.cef_binary_value_create(unsafe.Pointer(&data[0]), C.size_t(len(data))))
}

func DictionaryValueCreate() *CDictionaryValueT {
	return (*CDictionaryValueT)(C.cef_dictionary_value_create())
}

func ListValueCreate() *CListValueT {
	return (*CListValueT)(C.cef_list_value_create())
}

///
// Returns true (1) if this object is valid. This object may become invalid if
// the underlying data is owned by another object (e.g. list or dictionary)
// and that other object is then modified or destroyed. Do not call any other
// functions if this function returns false (0).
///
func (l *CListValueT) IsValid() bool {
	status := C.cefingo_list_value_is_valid((*C.cef_list_value_t)(l))
	return status == 1
}

///
// Returns true (1) if this object is currently owned by another object.
///
func (l *CListValueT) IsOwned() bool {
	status := C.cefingo_list_value_is_owned((*C.cef_list_value_t)(l))
	return status == 1
}

///
// Returns true (1) if the values of this object are read-only. Some APIs may
// expose read-only objects.
///
func (l *CListValueT) IsReadOnly() bool {
	status := C.cefingo_list_value_is_read_only((*C.cef_list_value_t)(l))
	return status == 1
}

///
// Returns true (1) if this object and |that| object have the same underlying
// data. If true (1) modifications to this object will also affect |that|
// object and vice-versa.
///
func (l *CListValueT) IsSame(that *CListValueT) bool {
	status := C.cefingo_list_value_is_same((*C.cef_list_value_t)(l),
		(*C.cef_list_value_t)(that))
	return status == 1
}

///
// Returns true (1) if this object and |that| object have an equivalent
// underlying value but are not necessarily the same object.
///
func (l *CListValueT) IsEqual(that *CListValueT) bool {
	status := C.cefingo_list_value_is_equal((*C.cef_list_value_t)(l),
		(*C.cef_list_value_t)(that))
	return status == 1
}

///
// Returns a writable copy of this object.
///
func (l *CListValueT) Copy() (d *CListValueT) {
	return (*CListValueT)(C.cefingo_list_value_copy((*C.cef_list_value_t)(l)))
}

///
// Sets the number of values. If the number of values is expanded all new
// value slots will default to type null. Returns true (1) on success.
///
func (l *CListValueT) SetSize(size int) bool {
	status := C.cefingo_list_value_set_size((*C.cef_list_value_t)(l),
		C.size_t(size))
	return status == 1
}

///
// Returns the number of values.
///
func (l *CListValueT) GetSize() int {
	return int(C.cefingo_list_value_get_size((*C.cef_list_value_t)(l)))

}

///
// Removes all values. Returns true (1) on success.
///
func (l *CListValueT) Clear() bool {
	status := C.cefingo_list_value_clear((*C.cef_list_value_t)(l))
	return status == 1
}

///
// Removes the value at the specified index.
///
func (l *CListValueT) Remove(index int) bool {
	status := C.cefingo_list_value_remove((*C.cef_list_value_t)(l), C.size_t(index))
	return status == 1
}

///
// Returns the value type at the specified index.
///
func (l *CListValueT) GetType(index int) CValueTypeT {
	return CValueTypeT(C.cefingo_list_value_get_type(
		(*C.cef_list_value_t)(l), C.size_t(index)))
}

///
// Returns the value at the specified index. For simple types the returned
// value will copy existing data and modifications to the value will not
// modify this object. For complex types (binary, dictionary and list) the
// returned value will reference existing data and modifications to the value
// will modify this object.
///
func (l *CListValueT) GetValue(index int) (v *CValueT) {
	v = (*CValueT)(C.cefingo_list_value_get_value(
		(*C.cef_list_value_t)(l), C.size_t(index)))
	BaseAddRef(v)
	return v
}

///
// Returns the value at the specified index as type bool.
///
func (l *CListValueT) GetBool(index int) bool {
	b := C.cefingo_list_value_get_bool((*C.cef_list_value_t)(l), C.size_t(index))
	return b == 1
}

///
// Returns the value at the specified index as type int.
///
func (l *CListValueT) GetInt(index int) int {
	return int(C.cefingo_list_value_get_int(
		(*C.cef_list_value_t)(l), (C.size_t)(index)))
}

///
// Returns the value at the specified index as type double.
///
func (l *CListValueT) GetDouble(index int) float64 {
	return float64(C.cefingo_list_value_get_double(
		(*C.cef_list_value_t)(l), C.size_t(index),
	))

}

///
// Returns the value at the specified index as type string.
///
// The resulting string must be freed by calling cef_string_userfree_free().
func (l *CListValueT) GetString(index int) (s string) {
	usfs := C.cefingo_list_value_get_string((*C.cef_list_value_t)(l), C.size_t(index))
	s = string_from_cef_string((*C.cef_string_t)(usfs))
	C.cef_string_userfree_free(usfs)
	return s
}

///
// Returns the value at the specified index as type binary. The returned value
// will reference existing data.
///
func (l *CListValueT) GetBinary(index int) *CBinaryValueT {
	b := (*CBinaryValueT)(C.cefingo_list_value_get_binary(
		(*C.cef_list_value_t)(l), C.size_t(index)))
	BaseAddRef(b)
	return b
}

///
// Returns the value at the specified index as type dictionary. The returned
// value will reference existing data and modifications to the value will
// modify this object.
///
func (l *CListValueT) GetDictionary(index int) *CDictionaryValueT {
	d := (*CDictionaryValueT)(C.cefingo_list_value_get_dictionary(
		(*C.cef_list_value_t)(l), C.size_t(index),
	))
	BaseAddRef(d)
	return d
}

///
// Returns the value at the specified index as type list. The returned value
// will reference existing data and modifications to the value will modify
// this object.
///
func (l *CListValueT) GetList(index int) *CListValueT {
	list := (*CListValueT)(C.cefingo_list_value_get_list(
		(*C.cef_list_value_t)(l), C.size_t(index),
	))
	BaseAddRef(list)
	return list
}

///
// Sets the value at the specified index. Returns true (1) if the value was
// set successfully. If |value| represents simple data then the underlying
// data will be copied and modifications to |value| will not modify this
// object. If |value| represents complex data (binary, dictionary or list)
// then the underlying data will be referenced and modifications to |value|
// will modify this object.
///
func (l *CListValueT) SetValue(index int, value *CValueT) bool {
	BaseAddRef(value)
	status := C.cefingo_list_value_set_value(
		(*C.cef_list_value_t)(l), C.size_t(index), (*C.cef_value_t)(value),
	)
	return status == 1
}

///
// Sets the value at the specified index as type null. Returns true (1) if the
// value was set successfully.
///
func (l *CListValueT) SetNull(index int) bool {
	status := C.cefingo_list_value_set_null(
		(*C.cef_list_value_t)(l), C.size_t(index),
	)
	return status == 1
}

///
// Sets the value at the specified index as type bool. Returns true (1) if the
// value was set successfully.
///
func (l *CListValueT) SetBool(index int, b bool) bool {
	var i int
	if b {
		i = 1
	}
	status := C.cefingo_list_value_set_bool((*C.cef_list_value_t)(l), C.size_t(index), C.int(i))
	return status == 1
}

///
// Sets the value at the specified index as type int. Returns true (1) if the
// value was set successfully.
///
func (l *CListValueT) SetInt(index int, value int) bool {
	status := C.cefingo_list_value_set_int(
		(*C.cef_list_value_t)(l), C.size_t(index), C.int(value))
	return status == 1
}

///
// Sets the value at the specified index as type double. Returns true (1) if
// the value was set successfully.
///
func (l *CListValueT) SetDouble(index int, value float64) bool {
	status := C.cefingo_list_value_set_double(
		(*C.cef_list_value_t)(l), C.size_t(index), C.double(value),
	)
	return status == 1
}

///
// Sets the value at the specified index as type string. Returns true (1) if
// the value was set successfully.
///
func (l *CListValueT) SetString(index int, s string) bool {
	value := create_cef_string(s)
	defer clear_cef_string(value)

	status := C.cefingo_list_value_set_string(
		(*C.cef_list_value_t)(l), C.size_t(index), value,
	)
	return status == 1
}

///
// Sets the value at the specified index as type binary. Returns true (1) if
// the value was set successfully. If |value| is currently owned by another
// object then the value will be copied and the |value| reference will not
// change. Otherwise, ownership will be transferred to this object and the
// |value| reference will be invalidated.
///
func (l *CListValueT) SetBinary(index int, value *CBinaryValueT) bool {
	BaseAddRef(value)
	status := C.cefingo_list_value_set_binary(
		(*C.cef_list_value_t)(l), C.size_t(index), (*C.cef_binary_value_t)(value),
	)
	return status == 1
}

///
// Sets the value at the specified index as type dict. Returns true (1) if the
// value was set successfully. If |value| is currently owned by another object
// then the value will be copied and the |value| reference will not change.
// Otherwise, ownership will be transferred to this object and the |value|
// reference will be invalidated.
///
func (l *CListValueT) SetDictionary(index int, value *CDictionaryValueT) bool {
	BaseAddRef(value)
	status := C.cefingo_list_value_set_dictionary(
		(*C.cef_list_value_t)(l), C.size_t(index),
		(*C.cef_dictionary_value_t)(value),
	)
	return status == 1
}

///
// Sets the value at the specified index as type list. Returns true (1) if the
// value was set successfully. If |value| is currently owned by another object
// then the value will be copied and the |value| reference will not change.
// Otherwise, ownership will be transferred to this object and the |value|
// reference will be invalidated.
///
func (l *CListValueT) SetList(index int, value *CListValueT) bool {
	BaseAddRef(value)
	status := C.cefingo_list_value_set_list(
		(*C.cef_list_value_t)(l), C.size_t(index),
		(*C.cef_list_value_t)(value))
	return status == 1
}
