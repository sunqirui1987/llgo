package core

import (
	"unsafe"
	_ "unsafe"
)

//go:linkname LvObjCreate C.lv_obj_create
func LvObjCreate(parent *lv_obj_t) *lv_obj_t

//go:linkname LvObjAddFlag C.lv_obj_add_flag
func LvObjAddFlag(obj *lv_obj_t, f lv_obj_flag_t)

//go:linkname LvObjRemoveFlag C.lv_obj_remove_flag
func LvObjRemoveFlag(obj *lv_obj_t, f lv_obj_flag_t)

//go:linkname LvObjUpdateFlag C.lv_obj_update_flag
func LvObjUpdateFlag(obj *lv_obj_t, f lv_obj_flag_t, v bool)

//go:linkname LvObjAddState C.lv_obj_add_state
func LvObjAddState(obj *lv_obj_t, state lv_state_t)

//go:linkname LvObjRemoveState C.lv_obj_remove_state
func LvObjRemoveState(obj *lv_obj_t, state lv_state_t)

//go:linkname LvObjSetState C.lv_obj_set_state
func LvObjSetState(obj *lv_obj_t, state lv_state_t, v bool)

//go:linkname LvObjSetUserData C.lv_obj_set_user_data
func LvObjSetUserData(obj *lv_obj_t, user_data unsafe.Pointer)

//go:linkname LvObjHasFlag C.lv_obj_has_flag
func LvObjHasFlag(obj *lv_obj_t, f lv_obj_flag_t) bool

//go:linkname LvObjHasFlagAny C.lv_obj_has_flag_any
func LvObjHasFlagAny(obj *lv_obj_t, f lv_obj_flag_t) bool

//go:linkname LvObjGetState C.lv_obj_get_state
func LvObjGetState(obj *lv_obj_t) lv_state_t

//go:linkname LvObjHasState C.lv_obj_has_state
func LvObjHasState(obj *lv_obj_t, state lv_state_t) bool

//go:linkname LvObjGetGroup C.lv_obj_get_group
func LvObjGetGroup(obj *lv_obj_t) *lv_group_t

//go:linkname LvObjGetUserData C.lv_obj_get_user_data
func LvObjGetUserData(obj *lv_obj_t) unsafe.Pointer

//go:linkname LvObjAllocateSpecAttr C.lv_obj_allocate_spec_attr
func LvObjAllocateSpecAttr(obj *lv_obj_t)

//go:linkname LvObjCheckType C.lv_obj_check_type
func LvObjCheckType(obj *lv_obj_t, class_p *lv_obj_class_t) bool

//go:linkname LvObjHasClass C.lv_obj_has_class
func LvObjHasClass(obj *lv_obj_t, class_p *lv_obj_class_t) bool

//go:linkname LvObjGetClass C.lv_obj_get_class
func LvObjGetClass(obj *lv_obj_t) *lv_obj_class_t

//go:linkname LvObjIsValid C.lv_obj_is_valid
func LvObjIsValid(obj *lv_obj_t) bool

//go:linkname LvObjNullOnDelete C.lv_obj_null_on_delete
func LvObjNullOnDelete(obj_ptr **lv_obj_t)
