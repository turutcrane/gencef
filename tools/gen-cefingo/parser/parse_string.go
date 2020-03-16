// Code generated by "stringer -output parse_string.go -type IdentKind,Ty,StructType,DefKind,TypeQualifier parse.go"; DO NOT EDIT.

package parser

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[IkNone-0]
	_ = x[IkName-1]
	_ = x[IkFunc-2]
	_ = x[IkArray-3]
	_ = x[IkStructTag-4]
}

const _IdentKind_name = "IkNoneIkNameIkFuncIkArrayIkStructTag"

var _IdentKind_index = [...]uint8{0, 6, 12, 18, 25, 36}

func (i IdentKind) String() string {
	if i < 0 || i >= IdentKind(len(_IdentKind_index)-1) {
		return "IdentKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _IdentKind_name[_IdentKind_index[i]:_IdentKind_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TyUnknown-0]
	_ = x[TyVoid-1]
	_ = x[TyChar-2]
	_ = x[TyUnsigned-3]
	_ = x[TyInt-4]
	_ = x[TyLong-5]
	_ = x[TyLongLong-6]
	_ = x[TyULong-7]
	_ = x[TyULongLong-8]
	_ = x[TyFloat-9]
	_ = x[TyDouble-10]
	_ = x[TySizeT-11]
	_ = x[TyHWND-12]
	_ = x[TyStructUnhandled-13]
	_ = x[TyStructSimple-14]
	_ = x[TyStructRefCounted-15]
	_ = x[TyStructScoped-16]
	_ = x[TyStructNotDefined-17]
	_ = x[TyInt16-18]
	_ = x[TyInt32-19]
	_ = x[TyInt64-20]
	_ = x[TyUint16-21]
	_ = x[TyUint32-22]
	_ = x[TyUint64-23]
	_ = x[TyTimeT-24]
	_ = x[TyStringT-25]
	_ = x[TyStringUserfreeT-26]
	_ = x[TyOther-27]
	_ = x[TyEnum-28]
	_ = x[TySimple-29]
	_ = x[TyMSG-30]
	_ = x[TyHCURSOR-31]
	_ = x[TyHINSTANCE-32]
	_ = x[TyDWORD-33]
	_ = x[TyHMENU-34]
}

const _Ty_name = "TyUnknownTyVoidTyCharTyUnsignedTyIntTyLongTyLongLongTyULongTyULongLongTyFloatTyDoubleTySizeTTyHWNDTyStructUnhandledTyStructSimpleTyStructRefCountedTyStructScopedTyStructNotDefinedTyInt16TyInt32TyInt64TyUint16TyUint32TyUint64TyTimeTTyStringTTyStringUserfreeTTyOtherTyEnumTySimpleTyMSGTyHCURSORTyHINSTANCETyDWORDTyHMENU"

var _Ty_index = [...]uint16{0, 9, 15, 21, 31, 36, 42, 52, 59, 70, 77, 85, 92, 98, 115, 129, 147, 161, 179, 186, 193, 200, 208, 216, 224, 231, 240, 257, 264, 270, 278, 283, 292, 303, 310, 317}

func (i Ty) String() string {
	if i < 0 || i >= Ty(len(_Ty_index)-1) {
		return "Ty(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Ty_name[_Ty_index[i]:_Ty_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[StUnknown-0]
	_ = x[StRefCounted-1]
	_ = x[StScoped-2]
	_ = x[StYetNotDefined-3]
}

const _StructType_name = "StUnknownStRefCountedStScopedStYetNotDefined"

var _StructType_index = [...]uint8{0, 9, 21, 29, 44}

func (i StructType) String() string {
	if i < 0 || i >= StructType(len(_StructType_index)-1) {
		return "StructType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _StructType_name[_StructType_index[i]:_StructType_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DkUnknown-0]
	_ = x[DkUnhandled-1]
	_ = x[DkSimple-2]
	_ = x[DkEnum-3]
	_ = x[DkCefClass-4]
	_ = x[DkFunc-5]
	_ = x[DkStruct-6]
}

const _DefKind_name = "DkUnknownDkUnhandledDkSimpleDkEnumDkCefClassDkFuncDkStruct"

var _DefKind_index = [...]uint8{0, 9, 20, 28, 34, 44, 50, 58}

func (i DefKind) String() string {
	if i < 0 || i >= DefKind(len(_DefKind_index)-1) {
		return "DefKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DefKind_name[_DefKind_index[i]:_DefKind_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TqUnknown-0]
	_ = x[TqConst-1]
	_ = x[TqPointer-2]
}

const _TypeQualifier_name = "TqUnknownTqConstTqPointer"

var _TypeQualifier_index = [...]uint8{0, 9, 16, 25}

func (i TypeQualifier) String() string {
	if i < 0 || i >= TypeQualifier(len(_TypeQualifier_index)-1) {
		return "TypeQualifier(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TypeQualifier_name[_TypeQualifier_index[i]:_TypeQualifier_index[i+1]]
}
