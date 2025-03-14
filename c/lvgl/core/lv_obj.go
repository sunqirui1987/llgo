package core

import (
	"unsafe"
	_ "unsafe"
)

//go:linkname LvObjCreate C.lv_obj_create
func LvObjCreate(parent *LvObjT) *LvObjT

//go:linkname LvObjAddFlag C.lv_obj_add_flag
func LvObjAddFlag(obj *LvObjT, f lv_obj_flag_t)

//go:linkname LvObjRemoveFlag C.lv_obj_remove_flag
func LvObjRemoveFlag(obj *LvObjT, f lv_obj_flag_t)

//go:linkname LvObjUpdateFlag C.lv_obj_update_flag
func LvObjUpdateFlag(obj *LvObjT, f lv_obj_flag_t, v bool)

//go:linkname LvObjAddState C.lv_obj_add_state
func LvObjAddState(obj *LvObjT, state lv_state_t)

//go:linkname LvObjRemoveState C.lv_obj_remove_state
func LvObjRemoveState(obj *LvObjT, state lv_state_t)

//go:linkname LvObjSetState C.lv_obj_set_state
func LvObjSetState(obj *LvObjT, state lv_state_t, v bool)

//go:linkname LvObjSetUserData C.lv_obj_set_user_data
func LvObjSetUserData(obj *LvObjT, user_data unsafe.Pointer)

//go:linkname LvObjHasFlag C.lv_obj_has_flag
func LvObjHasFlag(obj *LvObjT, f lv_obj_flag_t) bool

//go:linkname LvObjHasFlagAny C.lv_obj_has_flag_any
func LvObjHasFlagAny(obj *LvObjT, f lv_obj_flag_t) bool

//go:linkname LvObjGetState C.lv_obj_get_state
func LvObjGetState(obj *LvObjT) lv_state_t

//go:linkname LvObjHasState C.lv_obj_has_state
func LvObjHasState(obj *LvObjT, state lv_state_t) bool

//go:linkname LvObjGetGroup C.lv_obj_get_group
func LvObjGetGroup(obj *LvObjT) *lv_group_t

//go:linkname LvObjGetUserData C.lv_obj_get_user_data
func LvObjGetUserData(obj *LvObjT) unsafe.Pointer

//go:linkname LvObjAllocateSpecAttr C.lv_obj_allocate_spec_attr
func LvObjAllocateSpecAttr(obj *LvObjT)

//go:linkname LvObjCheckType C.lv_obj_check_type
func LvObjCheckType(obj *LvObjT, class_p *lv_obj_class_t) bool

//go:linkname LvObjHasClass C.lv_obj_has_class
func LvObjHasClass(obj *LvObjT, class_p *lv_obj_class_t) bool

//go:linkname LvObjGetClass C.lv_obj_get_class
func LvObjGetClass(obj *LvObjT) *lv_obj_class_t

//go:linkname LvObjIsValid C.lv_obj_is_valid
func LvObjIsValid(obj *LvObjT) bool

//go:linkname LvObjNullOnDelete C.lv_obj_null_on_delete
func LvObjNullOnDelete(obj_ptr **LvObjT)
