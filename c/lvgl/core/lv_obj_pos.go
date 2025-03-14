package core

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

//go:linkname LvObjSetPos C.lv_obj_set_pos
func LvObjSetPos(obj *lv_obj_t, x, y c.Int32T)

//go:linkname LvObjSetX C.lv_obj_set_x
func LvObjSetX(obj *lv_obj_t, x c.Int32T)

//go:linkname LvObjSetY C.lv_obj_set_y
func LvObjSetY(obj *lv_obj_t, y c.Int32T)

//go:linkname LvObjSetSize C.lv_obj_set_size
func LvObjSetSize(obj *lv_obj_t, w, h c.Int32T)

//go:linkname LvObjRefrSize C.lv_obj_refr_size
func LvObjRefrSize(obj *lv_obj_t) bool

//go:linkname LvObjSetWidth C.lv_obj_set_width
func LvObjSetWidth(obj *lv_obj_t, w c.Int32T)

//go:linkname LvObjSetHeight C.lv_obj_set_height
func LvObjSetHeight(obj *lv_obj_t, h c.Int32T)

//go:linkname LvObjSetContentWidth C.lv_obj_set_content_width
func LvObjSetContentWidth(obj *lv_obj_t, w c.Int32T)

//go:linkname LvObjSetContentHeight C.lv_obj_set_content_height
func LvObjSetContentHeight(obj *lv_obj_t, h c.Int32T)

//go:linkname LvObjSetLayout C.lv_obj_set_layout
func LvObjSetLayout(obj *lv_obj_t, layout c.Uint32T)

//go:linkname LvObjIsLayoutPositioned C.lv_obj_is_layout_positioned
func LvObjIsLayoutPositioned(obj *lv_obj_t) bool

//go:linkname LvObjMarkLayoutAsDirty C.lv_obj_mark_layout_as_dirty
func LvObjMarkLayoutAsDirty(obj *lv_obj_t)

//go:linkname LvObjUpdateLayout C.lv_obj_update_layout
func LvObjUpdateLayout(obj *lv_obj_t)

//go:linkname LvObjSetAlign C.lv_obj_set_align
func LvObjSetAlign(obj *lv_obj_t, align lv_align_t)

//go:linkname LvObjAlign C.lv_obj_align
func LvObjAlign(obj *lv_obj_t, align lv_align_t, x_ofs, y_ofs c.Int32T)

//go:linkname LvObjAlignTo C.lv_obj_align_to
func LvObjAlignTo(obj *lv_obj_t, base *lv_obj_t, align lv_align_t, x_ofs, y_ofs c.Int32T)

//go:linkname LvObjCenter C.lv_obj_center
func LvObjCenter(obj *lv_obj_t)

//go:linkname LvObjSetTransform C.lv_obj_set_transform
func LvObjSetTransform(obj *lv_obj_t, matrix *lv_matrix_t)

//go:linkname LvObjResetTransform C.lv_obj_reset_transform
func LvObjResetTransform(obj *lv_obj_t)

//go:linkname LvObjGetCoords C.lv_obj_get_coords
func LvObjGetCoords(obj *lv_obj_t, coords *lv_area_t)

//go:linkname LvObjGetX C.lv_obj_get_x
func LvObjGetX(obj *lv_obj_t) c.Int32T

//go:linkname LvObjGetX2 C.lv_obj_get_x2
func LvObjGetX2(obj *lv_obj_t) c.Int32T

//go:linkname LvObjGetY C.lv_obj_get_y
func LvObjGetY(obj *lv_obj_t) c.Int32T

//go:linkname LvObjGetY2 C.lv_obj_get_y2
func LvObjGetY2(obj *lv_obj_t) c.Int32T

//go:linkname LvObjGetXAligned C.lv_obj_get_x_aligned
func LvObjGetXAligned(obj *lv_obj_t) c.Int32T

//go:linkname LvObjGetYAligned C.lv_obj_get_y_aligned
func LvObjGetYAligned(obj *lv_obj_t) c.Int32T

//go:linkname LvObjGetWidth C.lv_obj_get_width
func LvObjGetWidth(obj *lv_obj_t) c.Int32T

//go:linkname LvObjGetHeight C.lv_obj_get_height
func LvObjGetHeight(obj *lv_obj_t) c.Int32T

//go:linkname LvObjGetContentWidth C.lv_obj_get_content_width
func LvObjGetContentWidth(obj *lv_obj_t) c.Int32T

//go:linkname LvObjGetContentHeight C.lv_obj_get_content_height
func LvObjGetContentHeight(obj *lv_obj_t) c.Int32T

//go:linkname LvObjGetContentCoords C.lv_obj_get_content_coords
func LvObjGetContentCoords(obj *lv_obj_t, area *lv_area_t)

//go:linkname LvObjGetSelfWidth C.lv_obj_get_self_width
func LvObjGetSelfWidth(obj *lv_obj_t) c.Int32T

//go:linkname LvObjGetSelfHeight C.lv_obj_get_self_height
func LvObjGetSelfHeight(obj *lv_obj_t) c.Int32T

//go:linkname LvObjRefreshSelfSize C.lv_obj_refresh_self_size
func LvObjRefreshSelfSize(obj *lv_obj_t) bool

//go:linkname LvObjRefrPos C.lv_obj_refr_pos
func LvObjRefrPos(obj *lv_obj_t)

//go:linkname LvObjMoveTo C.lv_obj_move_to
func LvObjMoveTo(obj *lv_obj_t, x, y c.Int32T)

//go:linkname LvObjMoveChildrenBy C.lv_obj_move_children_by
func LvObjMoveChildrenBy(obj *lv_obj_t, x_diff, y_diff c.Int32T, ignore_floating bool)

//go:linkname LvObjGetTransform C.lv_obj_get_transform
func LvObjGetTransform(obj *lv_obj_t) *lv_matrix_t

//go:linkname LvObjTransformPoint C.lv_obj_transform_point
func LvObjTransformPoint(obj *lv_obj_t, p *lv_point_t, flags lv_obj_point_transform_flag_t)

//go:linkname LvObjTransformPointArray C.lv_obj_transform_point_array
func LvObjTransformPointArray(obj *lv_obj_t, points *lv_point_t, count c.SizeT, flags lv_obj_point_transform_flag_t)

//go:linkname LvObjGetTransformedArea C.lv_obj_get_transformed_area
func LvObjGetTransformedArea(obj *lv_obj_t, area *lv_area_t, flags lv_obj_point_transform_flag_t)

//go:linkname LvObjInvalidateArea C.lv_obj_invalidate_area
func LvObjInvalidateArea(obj *lv_obj_t, area *lv_area_t)

//go:linkname LvObjInvalidate C.lv_obj_invalidate
func LvObjInvalidate(obj *lv_obj_t)

//go:linkname LvObjAreaIsVisible C.lv_obj_area_is_visible
func LvObjAreaIsVisible(obj *lv_obj_t, area *lv_area_t) bool

//go:linkname LvObjIsVisible C.lv_obj_is_visible
func LvObjIsVisible(obj *lv_obj_t) bool

//go:linkname LvObjSetExtClickArea C.lv_obj_set_ext_click_area
func LvObjSetExtClickArea(obj *lv_obj_t, size c.Int32T)

//go:linkname LvObjGetClickArea C.lv_obj_get_click_area
func LvObjGetClickArea(obj *lv_obj_t, area *lv_area_t)

//go:linkname LvObjHitTest C.lv_obj_hit_test
func LvObjHitTest(obj *lv_obj_t, point *lv_point_t) bool

//go:linkname LvClampWidth C.lv_clamp_width
func LvClampWidth(width, min_width, max_width, ref_width c.Int32T) c.Int32T

//go:linkname LvClampHeight C.lv_clamp_height
func LvClampHeight(height, min_height, max_height, ref_height c.Int32T) c.Int32T
