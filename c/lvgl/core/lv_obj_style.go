package core

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

//go:linkname LvObjAddStyle C.lv_obj_add_style
func LvObjAddStyle(obj *lv_obj_t, style *lv_style_t, selector lv_style_selector_t)

//go:linkname LvObjReplaceStyle C.lv_obj_replace_style
func LvObjReplaceStyle(obj *lv_obj_t, old_style *lv_style_t, new_style *lv_style_t, selector lv_style_selector_t) bool

//go:linkname LvObjRemoveStyle C.lv_obj_remove_style
func LvObjRemoveStyle(obj *lv_obj_t, style *lv_style_t, selector lv_style_selector_t)

//go:linkname LvObjRemoveStyleAll C.lv_obj_remove_style_all
func LvObjRemoveStyleAll(obj *lv_obj_t)

//go:linkname LvObjReportStyleChange C.lv_obj_report_style_change
func LvObjReportStyleChange(style *lv_style_t)

//go:linkname LvObjRefreshStyle C.lv_obj_refresh_style
func LvObjRefreshStyle(obj *lv_obj_t, part lv_part_t, prop lv_style_prop_t)

//go:linkname LvObjEnableStyleRefresh C.lv_obj_enable_style_refresh
func LvObjEnableStyleRefresh(en bool)

//go:linkname LvObjGetStyleProp C.lv_obj_get_style_prop
func LvObjGetStyleProp(obj *lv_obj_t, part lv_part_t, prop lv_style_prop_t) lv_style_value_t

//go:linkname LvObjHasStyleProp C.lv_obj_has_style_prop
func LvObjHasStyleProp(obj *lv_obj_t, selector lv_style_selector_t, prop lv_style_prop_t) bool

//go:linkname LvObjSetLocalStyleProp C.lv_obj_set_local_style_prop
func LvObjSetLocalStyleProp(obj *lv_obj_t, prop lv_style_prop_t, value lv_style_value_t, selector lv_style_selector_t)

//go:linkname LvObjGetLocalStyleProp C.lv_obj_get_local_style_prop
func LvObjGetLocalStyleProp(obj *lv_obj_t, prop lv_style_prop_t, value *lv_style_value_t, selector lv_style_selector_t) lv_style_res_t

//go:linkname LvObjRemoveLocalStyleProp C.lv_obj_remove_local_style_prop
func LvObjRemoveLocalStyleProp(obj *lv_obj_t, prop lv_style_prop_t, selector lv_style_selector_t) bool

//go:linkname LvObjStyleApplyColorFilter C.lv_obj_style_apply_color_filter
func LvObjStyleApplyColorFilter(obj *lv_obj_t, part lv_part_t, v lv_style_value_t) lv_style_value_t

//go:linkname LvObjFadeIn C.lv_obj_fade_in
func LvObjFadeIn(obj *lv_obj_t, time c.Uint32T, delay c.Uint32T)

//go:linkname LvObjFadeOut C.lv_obj_fade_out
func LvObjFadeOut(obj *lv_obj_t, time c.Uint32T, delay c.Uint32T)
