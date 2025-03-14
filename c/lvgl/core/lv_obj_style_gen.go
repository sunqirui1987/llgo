package core

import (
	"unsafe"
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

//go:linkname LvStyleSetWidth C.lv_style_set_width
func LvStyleSetWidth(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetBgOpa C.lv_style_set_bg_opa
func LvStyleSetBgOpa(style *lv_style_t, value lv_opa_t)

//go:linkname LvStyleSetBgColor C.lv_style_set_bg_color
func LvStyleSetBgColor(style *lv_style_t, value lv_color_t)

//go:linkname LvStyleSetRadius C.lv_style_set_radius
func LvStyleSetRadius(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetPadRight C.lv_style_set_pad_right
func LvStyleSetPadRight(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetMinWidth C.lv_style_set_min_width
func LvStyleSetMinWidth(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetMaxWidth C.lv_style_set_max_width
func LvStyleSetMaxWidth(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetHeight C.lv_style_set_height
func LvStyleSetHeight(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetMinHeight C.lv_style_set_min_height
func LvStyleSetMinHeight(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetMaxHeight C.lv_style_set_max_height
func LvStyleSetMaxHeight(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetLength C.lv_style_set_length
func LvStyleSetLength(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetX C.lv_style_set_x
func LvStyleSetX(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetY C.lv_style_set_y
func LvStyleSetY(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetAlign C.lv_style_set_align
func LvStyleSetAlign(style *lv_style_t, value lv_align_t)

//go:linkname LvStyleSetTransformWidth C.lv_style_set_transform_width
func LvStyleSetTransformWidth(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetTransformHeight C.lv_style_set_transform_height
func LvStyleSetTransformHeight(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetTranslateX C.lv_style_set_translate_x
func LvStyleSetTranslateX(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetTranslateY C.lv_style_set_translate_y
func LvStyleSetTranslateY(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetTranslateRadial C.lv_style_set_translate_radial
func LvStyleSetTranslateRadial(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetTransformScaleX C.lv_style_set_transform_scale_x
func LvStyleSetTransformScaleX(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetTransformScaleY C.lv_style_set_transform_scale_y
func LvStyleSetTransformScaleY(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetTransformRotation C.lv_style_set_transform_rotation
func LvStyleSetTransformRotation(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetTransformPivotX C.lv_style_set_transform_pivot_x
func LvStyleSetTransformPivotX(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetTransformPivotY C.lv_style_set_transform_pivot_y
func LvStyleSetTransformPivotY(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetTransformSkewX C.lv_style_set_transform_skew_x
func LvStyleSetTransformSkewX(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetTransformSkewY C.lv_style_set_transform_skew_y
func LvStyleSetTransformSkewY(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetPadTop C.lv_style_set_pad_top
func LvStyleSetPadTop(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetPadBottom C.lv_style_set_pad_bottom
func LvStyleSetPadBottom(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetPadLeft C.lv_style_set_pad_left
func LvStyleSetPadLeft(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetPadRow C.lv_style_set_pad_row
func LvStyleSetPadRow(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetPadColumn C.lv_style_set_pad_column
func LvStyleSetPadColumn(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetPadRadial C.lv_style_set_pad_radial
func LvStyleSetPadRadial(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetMarginTop C.lv_style_set_margin_top
func LvStyleSetMarginTop(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetMarginBottom C.lv_style_set_margin_bottom
func LvStyleSetMarginBottom(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetMarginLeft C.lv_style_set_margin_left
func LvStyleSetMarginLeft(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetMarginRight C.lv_style_set_margin_right
func LvStyleSetMarginRight(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetBgGradColor C.lv_style_set_bg_grad_color
func LvStyleSetBgGradColor(style *lv_style_t, value lv_color_t)

//go:linkname LvStyleSetBgGradDir C.lv_style_set_bg_grad_dir
func LvStyleSetBgGradDir(style *lv_style_t, value lv_grad_dir_t)

//go:linkname LvStyleSetBgMainStop C.lv_style_set_bg_main_stop
func LvStyleSetBgMainStop(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetBgGradStop C.lv_style_set_bg_grad_stop
func LvStyleSetBgGradStop(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetBgMainOpa C.lv_style_set_bg_main_opa
func LvStyleSetBgMainOpa(style *lv_style_t, value lv_opa_t)

//go:linkname LvStyleSetBgGradOpa C.lv_style_set_bg_grad_opa
func LvStyleSetBgGradOpa(style *lv_style_t, value lv_opa_t)

//go:linkname LvStyleSetBgGrad C.lv_style_set_bg_grad
func LvStyleSetBgGrad(style *lv_style_t, value *lv_grad_dsc_t)

//go:linkname LvStyleSetBgImageSrc C.lv_style_set_bg_image_src
func LvStyleSetBgImageSrc(style *lv_style_t, value unsafe.Pointer)

//go:linkname LvStyleSetBgImageOpa C.lv_style_set_bg_image_opa
func LvStyleSetBgImageOpa(style *lv_style_t, value lv_opa_t)

//go:linkname LvStyleSetBgImageRecolor C.lv_style_set_bg_image_recolor
func LvStyleSetBgImageRecolor(style *lv_style_t, value lv_color_t)

//go:linkname LvStyleSetBgImageRecolorOpa C.lv_style_set_bg_image_recolor_opa
func LvStyleSetBgImageRecolorOpa(style *lv_style_t, value lv_opa_t)

//go:linkname LvStyleSetBgImageTiled C.lv_style_set_bg_image_tiled
func LvStyleSetBgImageTiled(style *lv_style_t, value bool)

//go:linkname LvStyleSetBorderColor C.lv_style_set_border_color
func LvStyleSetBorderColor(style *lv_style_t, value lv_color_t)

//go:linkname LvStyleSetBorderOpa C.lv_style_set_border_opa
func LvStyleSetBorderOpa(style *lv_style_t, value lv_opa_t)

//go:linkname LvStyleSetBorderWidth C.lv_style_set_border_width
func LvStyleSetBorderWidth(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetBorderSide C.lv_style_set_border_side
func LvStyleSetBorderSide(style *lv_style_t, value lv_border_side_t)

//go:linkname LvStyleSetBorderPost C.lv_style_set_border_post
func LvStyleSetBorderPost(style *lv_style_t, value bool)

//go:linkname LvStyleSetOutlineWidth C.lv_style_set_outline_width
func LvStyleSetOutlineWidth(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetOutlineColor C.lv_style_set_outline_color
func LvStyleSetOutlineColor(style *lv_style_t, value lv_color_t)

//go:linkname LvStyleSetOutlineOpa C.lv_style_set_outline_opa
func LvStyleSetOutlineOpa(style *lv_style_t, value lv_opa_t)

//go:linkname LvStyleSetOutlinePad C.lv_style_set_outline_pad
func LvStyleSetOutlinePad(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetShadowWidth C.lv_style_set_shadow_width
func LvStyleSetShadowWidth(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetShadowOffsetX C.lv_style_set_shadow_offset_x
func LvStyleSetShadowOffsetX(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetShadowOffsetY C.lv_style_set_shadow_offset_y
func LvStyleSetShadowOffsetY(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetShadowSpread C.lv_style_set_shadow_spread
func LvStyleSetShadowSpread(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetShadowColor C.lv_style_set_shadow_color
func LvStyleSetShadowColor(style *lv_style_t, value lv_color_t)

//go:linkname LvStyleSetShadowOpa C.lv_style_set_shadow_opa
func LvStyleSetShadowOpa(style *lv_style_t, value lv_opa_t)

//go:linkname LvStyleSetImageOpa C.lv_style_set_image_opa
func LvStyleSetImageOpa(style *lv_style_t, value lv_opa_t)

//go:linkname LvStyleSetImageRecolor C.lv_style_set_image_recolor
func LvStyleSetImageRecolor(style *lv_style_t, value lv_color_t)

//go:linkname LvStyleSetImageRecolorOpa C.lv_style_set_image_recolor_opa
func LvStyleSetImageRecolorOpa(style *lv_style_t, value lv_opa_t)

//go:linkname LvStyleSetLineWidth C.lv_style_set_line_width
func LvStyleSetLineWidth(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetLineDashWidth C.lv_style_set_line_dash_width
func LvStyleSetLineDashWidth(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetLineDashGap C.lv_style_set_line_dash_gap
func LvStyleSetLineDashGap(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetLineRounded C.lv_style_set_line_rounded
func LvStyleSetLineRounded(style *lv_style_t, value bool)

//go:linkname LvStyleSetLineColor C.lv_style_set_line_color
func LvStyleSetLineColor(style *lv_style_t, value lv_color_t)

//go:linkname LvStyleSetLineOpa C.lv_style_set_line_opa
func LvStyleSetLineOpa(style *lv_style_t, value lv_opa_t)

//go:linkname LvStyleSetArcWidth C.lv_style_set_arc_width
func LvStyleSetArcWidth(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetArcRounded C.lv_style_set_arc_rounded
func LvStyleSetArcRounded(style *lv_style_t, value bool)

//go:linkname LvStyleSetArcColor C.lv_style_set_arc_color
func LvStyleSetArcColor(style *lv_style_t, value lv_color_t)

//go:linkname LvStyleSetArcOpa C.lv_style_set_arc_opa
func LvStyleSetArcOpa(style *lv_style_t, value lv_opa_t)

//go:linkname LvStyleSetArcImageSrc C.lv_style_set_arc_image_src
func LvStyleSetArcImageSrc(style *lv_style_t, value unsafe.Pointer)

//go:linkname LvStyleSetTextColor C.lv_style_set_text_color
func LvStyleSetTextColor(style *lv_style_t, value lv_color_t)

//go:linkname LvStyleSetTextOpa C.lv_style_set_text_opa
func LvStyleSetTextOpa(style *lv_style_t, value lv_opa_t)

//go:linkname LvStyleSetTextFont C.lv_style_set_text_font
func LvStyleSetTextFont(style *lv_style_t, value *lv_font_t)

//go:linkname LvStyleSetTextLetterSpace C.lv_style_set_text_letter_space
func LvStyleSetTextLetterSpace(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetTextLineSpace C.lv_style_set_text_line_space
func LvStyleSetTextLineSpace(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetTextDecor C.lv_style_set_text_decor
func LvStyleSetTextDecor(style *lv_style_t, value lv_text_decor_t)

//go:linkname LvStyleSetTextAlign C.lv_style_set_text_align
func LvStyleSetTextAlign(style *lv_style_t, value lv_text_align_t)

//go:linkname LvStyleSetRadialOffset C.lv_style_set_radial_offset
func LvStyleSetRadialOffset(style *lv_style_t, value c.Int)

//go:linkname LvStyleSetClipCorner C.lv_style_set_clip_corner
func LvStyleSetClipCorner(style *lv_style_t, value bool)

//go:linkname LvStyleSetOpa C.lv_style_set_opa
func LvStyleSetOpa(style *lv_style_t, value lv_opa_t)

//go:linkname LvStyleSetOpaLayered C.lv_style_set_opa_layered
func LvStyleSetOpaLayered(style *lv_style_t, value lv_opa_t)

//go:linkname LvStyleSetColorFilterDsc C.lv_style_set_color_filter_dsc
func LvStyleSetColorFilterDsc(style *lv_style_t, value *lv_color_filter_dsc_t)

//go:linkname LvStyleSetColorFilterOpa C.lv_style_set_color_filter_opa
func LvStyleSetColorFilterOpa(style *lv_style_t, value lv_opa_t)

//go:linkname LvStyleSetAnim C.lv_style_set_anim
func LvStyleSetAnim(style *lv_style_t, value *lv_anim_t)

//go:linkname LvStyleSetAnimDuration C.lv_style_set_anim_duration
func LvStyleSetAnimDuration(style *lv_style_t, value c.Uint)

//go:linkname LvStyleSetTransition C.lv_style_set_transition
func LvStyleSetTransition(style *lv_style_t, value *lv_style_transition_dsc_t)

//go:linkname LvStyleSetBlendMode C.lv_style_set_blend_mode
func LvStyleSetBlendMode(style *lv_style_t, value lv_blend_mode_t)

//go:linkname LvStyleSetLayout C.lv_style_set_layout
func LvStyleSetLayout(style *lv_style_t, value c.Int16T)

//go:linkname LvStyleSetBaseDir C.lv_style_set_base_dir
func LvStyleSetBaseDir(style *lv_style_t, value lv_base_dir_t)

//go:linkname LvStyleSetBitmapMaskSrc C.lv_style_set_bitmap_mask_src
func LvStyleSetBitmapMaskSrc(style *lv_style_t, value *c.Void)

//go:linkname LvStyleSetRotarySensitivity C.lv_style_set_rotary_sensitivity
func LvStyleSetRotarySensitivity(style *lv_style_t, value c.Int16T)
