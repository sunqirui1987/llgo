package core

import (
	"unsafe"
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

//go:linkname LvObjSendEvent C.lv_obj_send_event
func LvObjSendEvent(obj *lv_obj_t, event_code lv_event_code_t, param unsafe.Pointer) lv_result_t

//go:linkname LvObjEventBase C.lv_obj_event_base
func LvObjEventBase(class_p *lv_obj_class_t, e *lv_event_t) lv_result_t

//go:linkname LvEventGetCurrentTargetObj C.lv_event_get_current_target_obj
func LvEventGetCurrentTargetObj(e *lv_event_t) *lv_obj_t

//go:linkname LvEventGetTargetObj C.lv_event_get_target_obj
func LvEventGetTargetObj(e *lv_event_t) *lv_obj_t

//go:linkname LvObjAddEventCb C.lv_obj_add_event_cb
func LvObjAddEventCb(obj *lv_obj_t, event_cb lv_event_cb_t, filter lv_event_code_t, user_data unsafe.Pointer) *lv_event_dsc_t

//go:linkname LvObjGetEventCount C.lv_obj_get_event_count
func LvObjGetEventCount(obj *lv_obj_t) c.Uint

//go:linkname LvObjGetEventDsc C.lv_obj_get_event_dsc
func LvObjGetEventDsc(obj *lv_obj_t, index c.Uint) *lv_event_dsc_t

//go:linkname LvObjRemoveEvent C.lv_obj_remove_event
func LvObjRemoveEvent(obj *lv_obj_t, index c.Uint) bool

//go:linkname LvObjRemoveEventCb C.lv_obj_remove_event_cb
func LvObjRemoveEventCb(obj *lv_obj_t, event_cb lv_event_cb_t) bool

//go:linkname LvObjRemoveEventDsc C.lv_obj_remove_event_dsc
func LvObjRemoveEventDsc(obj *lv_obj_t, dsc *lv_event_dsc_t) bool

//go:linkname LvObjRemoveEventCbWithUserData C.lv_obj_remove_event_cb_with_user_data
func LvObjRemoveEventCbWithUserData(obj *lv_obj_t, event_cb lv_event_cb_t, user_data unsafe.Pointer) c.Uint

//go:linkname LvEventGetIndev C.lv_event_get_indev
func LvEventGetIndev(e *lv_event_t) *lv_indev_t

//go:linkname LvEventGetLayer C.lv_event_get_layer
func LvEventGetLayer(e *lv_event_t) *lv_layer_t

//go:linkname LvEventGetOldSize C.lv_event_get_old_size
func LvEventGetOldSize(e *lv_event_t) *lv_area_t

//go:linkname LvEventGetKey C.lv_event_get_key
func LvEventGetKey(e *lv_event_t) c.Uint

//go:linkname LvEventGetRotaryDiff C.lv_event_get_rotary_diff
func LvEventGetRotaryDiff(e *lv_event_t) c.Int

//go:linkname LvEventGetScrollAnim C.lv_event_get_scroll_anim
func LvEventGetScrollAnim(e *lv_event_t) *lv_anim_t

//go:linkname LvEventSetExtDrawSize C.lv_event_set_ext_draw_size
func LvEventSetExtDrawSize(e *lv_event_t, size c.Int)

//go:linkname LvEventGetSelfSizeInfo C.lv_event_get_self_size_info
func LvEventGetSelfSizeInfo(e *lv_event_t) *lv_point_t

//go:linkname LvEventGetHitTestInfo C.lv_event_get_hit_test_info
func LvEventGetHitTestInfo(e *lv_event_t) *lv_hit_test_info_t

//go:linkname LvEventGetCoverArea C.lv_event_get_cover_area
func LvEventGetCoverArea(e *lv_event_t) *lv_area_t

//go:linkname LvEventSetCoverRes C.lv_event_set_cover_res
func LvEventSetCoverRes(e *lv_event_t, res lv_cover_res_t)

//go:linkname LvEventGetDrawTask C.lv_event_get_draw_task
func LvEventGetDrawTask(e *lv_event_t) *lv_draw_task_t
