package core

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

// #if LV_USE_FLOAT
//type lv_value_precise_t c.Float

// #else
type lv_value_precise_t c.Int

// #endif

/**
 * LVGL error codes.
 */

type lv_result_t c.Int

const (
	LV_RESULT_INVALID lv_result_t = 0 /*Typically indicates that the object is deleted (become invalid) in the action
	  function or an operation was failed*/
	LV_RESULT_OK lv_result_t = 1 /*The object is valid (no deleted) after the action*/
)

type lv_opa_t c.Int

const (
	LV_OPA_TRANSP lv_opa_t = 0
	LV_OPA_0      lv_opa_t = 0
	LV_OPA_10     lv_opa_t = 25
	LV_OPA_20     lv_opa_t = 51
	LV_OPA_30     lv_opa_t = 76
	LV_OPA_40     lv_opa_t = 102
	LV_OPA_50     lv_opa_t = 127
	LV_OPA_60     lv_opa_t = 153
	LV_OPA_70     lv_opa_t = 178
	LV_OPA_80     lv_opa_t = 204
	LV_OPA_90     lv_opa_t = 229
	LV_OPA_100    lv_opa_t = 255
	LV_OPA_COVER  lv_opa_t = 255
)

const (
	LV_OPA_MIN lv_opa_t = 2
	LV_OPA_MAX lv_opa_t = 253
)

type lv_radius_t c.Int

const (
	LV_RADIUS_CIRCLE lv_radius_t = 0x7FFF
)

//llgo:type C
type lv_style_transition_dsc_t struct {
	Unused [0]byte
}

//llgo:type C
type lv_anim_t struct {
	Unused [0]byte
}

//llgo:type C
type lv_color_filter_dsc_t struct {
	Unused [0]byte
}

//llgo:type C
type lv_grad_dsc_t struct {
	Unused [0]byte
}

//llgo:type C
type lv_font_t struct {
	Unused [0]byte
}

//llgo:type C
type lv_obj_t struct {
	Unused [0]byte
}

//llgo:type C
type lv_obj_flag_t struct {
	Unused [0]byte
}

//llgo:type C
type lv_obj_class_t struct {
	Unused [0]byte
}

//llgo:type C
type lv_state_t struct {
	Unused [0]byte
}

//llgo:type C
type lv_group_t struct {
	Unused [0]byte
}

type lv_cover_res_t c.Int

const (
	LV_COVER_RES_COVER     lv_cover_res_t = 0
	LV_COVER_RES_NOT_COVER lv_cover_res_t = 1
	LV_COVER_RES_MASKED    lv_cover_res_t = 2
)

//llgo:type C
type lv_event_t struct {
	Unused [0]byte
}

//llgo:type C
type lv_event_dsc_t struct {
	Unused [0]byte
}

//llgo:type C
type lv_event_code_t struct {
	Unused [0]byte
}

//llgo:type C
type lv_event_cb_t struct {
	Unused [0]byte
}

//llgo:type C
type lv_indev_t struct {
	Unused [0]byte
}

//llgo:type C
type lv_layer_t struct {
	Unused [0]byte
}

//llgo:type C
type lv_area_t struct {
	Unused [0]byte
}

//llgo:type C
type lv_point_t struct {
	Unused [0]byte
}

//llgo:type C
type lv_hit_test_info_t struct {
	Unused [0]byte
}

//llgo:type C
type lv_draw_task_t struct {
	Unused [0]byte
}

type lv_obj_point_transform_flag_t c.Int

const (
	LV_OBJ_POINT_TRANSFORM_FLAG_NONE              lv_obj_point_transform_flag_t = 0x00
	LV_OBJ_POINT_TRANSFORM_FLAG_RECURSIVE         lv_obj_point_transform_flag_t = 0x01
	LV_OBJ_POINT_TRANSFORM_FLAG_INVERSE           lv_obj_point_transform_flag_t = 0x02
	LV_OBJ_POINT_TRANSFORM_FLAG_INVERSE_RECURSIVE lv_obj_point_transform_flag_t = 0x03
)

//llgo:type C
type lv_matrix_t struct {
	Unused [0]byte
}

//llgo:type
type lv_style_t struct {
	Unused [0]byte
}

type lv_style_state_cmp_t int

const (
	LV_STYLE_STATE_CMP_SAME lv_style_state_cmp_t = iota
	LV_STYLE_STATE_CMP_DIFF_REDRAW
	LV_STYLE_STATE_CMP_DIFF_DRAW_PAD
	LV_STYLE_STATE_CMP_DIFF_LAYOUT
)

type lv_style_selector_t c.Uint

type lv_align_t c.Int

const (
	LV_ALIGN_DEFAULT lv_align_t = iota
	LV_ALIGN_TOP_LEFT
	LV_ALIGN_TOP_MID
	LV_ALIGN_TOP_RIGHT
	LV_ALIGN_BOTTOM_LEFT
	LV_ALIGN_BOTTOM_MID
	LV_ALIGN_BOTTOM_RIGHT
	LV_ALIGN_LEFT_MID
	LV_ALIGN_RIGHT_MID
	LV_ALIGN_CENTER
	LV_ALIGN_OUT_TOP_LEFT
	LV_ALIGN_OUT_TOP_MID
	LV_ALIGN_OUT_TOP_RIGHT
	LV_ALIGN_OUT_BOTTOM_LEFT
	LV_ALIGN_OUT_BOTTOM_MID
	LV_ALIGN_OUT_BOTTOM_RIGHT
	LV_ALIGN_OUT_LEFT_TOP
	LV_ALIGN_OUT_LEFT_MID
	LV_ALIGN_OUT_LEFT_BOTTOM
	LV_ALIGN_OUT_RIGHT_TOP
	LV_ALIGN_OUT_RIGHT_MID
	LV_ALIGN_OUT_RIGHT_BOTTOM
)

// llgo:type C
type lv_color_t struct {
	Blue  c.Uint
	Green c.Uint
	Red   c.Uint
}

// lv_color16_t  LVGL（LittleVGL），lv_color16_t 就是用来存储 RGB565 颜色的结构体，
//
//llgo:type C
type lv_color16_t struct {
	Unused [0]byte
}

//llgo:type C
type lv_color32_t struct {
	Blue  c.Uint
	Green c.Uint
	Red   c.Uint
	Alpha c.Uint
}

//llgo:type C
type lv_color_hsv_t struct {
	Hue        c.Uint
	Saturation c.Uint
	Value      c.Uint
}

//llgo:type C
type lv_color16a_t struct {
	Luminance c.Uint
	Alpha     c.Uint
}

type lv_grad_dir_t c.Int

const (
	LV_GRAD_DIR_NONE    lv_grad_dir_t = iota /**< No gradient (the `grad_color` property is ignored)*/
	LV_GRAD_DIR_VER                          /**< Simple vertical (top to bottom) gradient*/
	LV_GRAD_DIR_HOR                          /**< Simple horizontal (left to right) gradient*/
	LV_GRAD_DIR_LINEAR                       /**< Linear gradient defined by start and end points. Can be at any angle.*/
	LV_GRAD_DIR_RADIAL                       /**< Radial gradient defined by start and end circles*/
	LV_GRAD_DIR_CONICAL                      /**< Conical gradient defined by center point, start and end angles*/
)

type lv_border_side_t c.Int

/**
 * Selects on which sides border should be drawn
 * 'OR'ed values can be used.
 */
const (
	LV_BORDER_SIDE_NONE     lv_border_side_t = 0x00
	LV_BORDER_SIDE_BOTTOM   lv_border_side_t = 0x01
	LV_BORDER_SIDE_TOP      lv_border_side_t = 0x02
	LV_BORDER_SIDE_LEFT     lv_border_side_t = 0x04
	LV_BORDER_SIDE_RIGHT    lv_border_side_t = 0x08
	LV_BORDER_SIDE_FULL     lv_border_side_t = 0x0F
	LV_BORDER_SIDE_INTERNAL lv_border_side_t = 0x10 /**< FOR matrix-like objects (e.g. Button matrix)*/
)

type lv_text_decor_t c.Int

/**
 * Some options to apply decorations on texts.
 * 'OR'ed values can be used.
 */
const (
	LV_TEXT_DECOR_NONE          lv_text_decor_t = 0x00
	LV_TEXT_DECOR_UNDERLINE     lv_text_decor_t = 0x01
	LV_TEXT_DECOR_STRIKETHROUGH lv_text_decor_t = 0x02
)

type lv_text_align_t c.Int

/** Label align policy*/
const (
	LV_TEXT_ALIGN_AUTO   lv_text_align_t = 0x00
	LV_TEXT_ALIGN_LEFT   lv_text_align_t = 0x01
	LV_TEXT_ALIGN_CENTER lv_text_align_t = 0x02
	LV_TEXT_ALIGN_RIGHT  lv_text_align_t = 0x03
)

type lv_blend_mode_t c.Int

/**
 * Possible options for blending opaque drawings
 */
const (
	LV_BLEND_MODE_NORMAL      lv_blend_mode_t = 0x00
	LV_BLEND_MODE_ADDITIVE    lv_blend_mode_t = 0x01
	LV_BLEND_MODE_SUBTRACTIVE lv_blend_mode_t = 0x02
	LV_BLEND_MODE_MULTIPLY    lv_blend_mode_t = 0x03
	LV_BLEND_MODE_DIFFERENCE  lv_blend_mode_t = 0x04
)

type lv_base_dir_t c.Int

const (
	LV_BASE_DIR_LTR  lv_base_dir_t = 0x00
	LV_BASE_DIR_RTL  lv_base_dir_t = 0x01
	LV_BASE_DIR_AUTO lv_base_dir_t = 0x02

	LV_BASE_DIR_NEUTRAL lv_base_dir_t = 0x20
	LV_BASE_DIR_WEAK    lv_base_dir_t = 0x21
)

//llgo:type C
type lv_style_value_t struct {
	Unused [0]byte
}

//llgo:type C
type lv_style_prop_t struct {
	Unused [0]byte
}

//llgo:type C
type lv_part_t struct {
	Unused [0]byte
}

//llgo:type C
type lv_style_res_t struct {
	Unused [0]byte
}

//llgo:type C
type lv_draw_buf_t struct {
	Unused [0]byte
}

type lv_color_format_t c.Int

const (
	LV_COLOR_FORMAT_UNKNOWN lv_color_format_t = 0

	LV_COLOR_FORMAT_RAW       lv_color_format_t = 0x01
	LV_COLOR_FORMAT_RAW_ALPHA lv_color_format_t = 0x02

	/*<=1 byte (+alpha) formats*/
	LV_COLOR_FORMAT_L8 lv_color_format_t = 0x06
	LV_COLOR_FORMAT_I1 lv_color_format_t = 0x07
	LV_COLOR_FORMAT_I2 lv_color_format_t = 0x08
	LV_COLOR_FORMAT_I4 lv_color_format_t = 0x09
	LV_COLOR_FORMAT_I8 lv_color_format_t = 0x0A
	LV_COLOR_FORMAT_A8 lv_color_format_t = 0x0E

	/*2 byte (+alpha) formats*/
	LV_COLOR_FORMAT_RGB565   lv_color_format_t = 0x12
	LV_COLOR_FORMAT_ARGB8565 lv_color_format_t = 0x13
	LV_COLOR_FORMAT_RGB565A8 lv_color_format_t = 0x14
	LV_COLOR_FORMAT_AL88     lv_color_format_t = 0x15

	/*3 byte (+alpha) formats*/
	LV_COLOR_FORMAT_RGB888   lv_color_format_t = 0x0F
	LV_COLOR_FORMAT_ARGB8888 lv_color_format_t = 0x10
	LV_COLOR_FORMAT_XRGB8888 lv_color_format_t = 0x11

	/*Formats not supported by software renderer but kept here so GPU can use it*/
	LV_COLOR_FORMAT_A1       lv_color_format_t = 0x0B
	LV_COLOR_FORMAT_A2       lv_color_format_t = 0x0C
	LV_COLOR_FORMAT_A4       lv_color_format_t = 0x0D
	LV_COLOR_FORMAT_ARGB1555 lv_color_format_t = 0x16
	LV_COLOR_FORMAT_ARGB4444 lv_color_format_t = 0x17
	LV_COLOR_FORMAT_ARGB2222 lv_color_format_t = 0x18

	/* reference to https://wiki.videolan.org/YUV/ */
	/*YUV planar formats*/
	LV_COLOR_FORMAT_YUV_START lv_color_format_t = 0x20
	LV_COLOR_FORMAT_I420      lv_color_format_t = LV_COLOR_FORMAT_YUV_START /*YUV420 planar(3 plane)*/
	LV_COLOR_FORMAT_I422      lv_color_format_t = 0x21                      /*YUV422 planar(3 plane)*/
	LV_COLOR_FORMAT_I444      lv_color_format_t = 0x22                      /*YUV444 planar(3 plane)*/
	LV_COLOR_FORMAT_I400      lv_color_format_t = 0x23                      /*YUV400 no chroma channel*/
	LV_COLOR_FORMAT_NV21      lv_color_format_t = 0x24                      /*YUV420 planar(2 plane), UV plane in 'V, U, V, U'*/
	LV_COLOR_FORMAT_NV12      lv_color_format_t = 0x25                      /*YUV420 planar(2 plane), UV plane in 'U, V, U, V'*/

	/*YUV packed formats*/
	LV_COLOR_FORMAT_YUY2 lv_color_format_t = 0x26 /*YUV422 packed like 'Y U Y V'*/
	LV_COLOR_FORMAT_UYVY lv_color_format_t = 0x27 /*YUV422 packed like 'U Y V Y'*/

	LV_COLOR_FORMAT_YUV_END lv_color_format_t = LV_COLOR_FORMAT_UYVY

	LV_COLOR_FORMAT_PROPRIETARY_START lv_color_format_t = 0x30

	LV_COLOR_FORMAT_NEMA_TSC_START lv_color_format_t = LV_COLOR_FORMAT_PROPRIETARY_START
	LV_COLOR_FORMAT_NEMA_TSC4      lv_color_format_t = LV_COLOR_FORMAT_NEMA_TSC_START
	LV_COLOR_FORMAT_NEMA_TSC6      lv_color_format_t = 0x31
	LV_COLOR_FORMAT_NEMA_TSC6A     lv_color_format_t = 0x32
	LV_COLOR_FORMAT_NEMA_TSC6AP    lv_color_format_t = 0x33
	LV_COLOR_FORMAT_NEMA_TSC12     lv_color_format_t = 0x34
	LV_COLOR_FORMAT_NEMA_TSC12A    lv_color_format_t = 0x35
	LV_COLOR_FORMAT_NEMA_TSC_END   lv_color_format_t = LV_COLOR_FORMAT_NEMA_TSC12A

	//     /*Color formats in which LVGL can render*/
	// #if LV_COLOR_DEPTH == 1
	// LV_COLOR_FORMAT_NATIVE            lv_color_format_t = LV_COLOR_FORMAT_I1
	// LV_COLOR_FORMAT_NATIVE_WITH_ALPHA lv_color_format_t = LV_COLOR_FORMAT_I1
	// // #elif LV_COLOR_DEPTH == 8
	// LV_COLOR_FORMAT_NATIVE            lv_color_format_t = LV_COLOR_FORMAT_L8
	// LV_COLOR_FORMAT_NATIVE_WITH_ALPHA lv_color_format_t = LV_COLOR_FORMAT_AL88
	// // #elif LV_COLOR_DEPTH == 16
	LV_COLOR_FORMAT_NATIVE            lv_color_format_t = LV_COLOR_FORMAT_RGB565
	LV_COLOR_FORMAT_NATIVE_WITH_ALPHA lv_color_format_t = LV_COLOR_FORMAT_RGB565A8
	// #elif LV_COLOR_DEPTH == 24
	//     LV_COLOR_FORMAT_NATIVE            = LV_COLOR_FORMAT_RGB888,
	//     LV_COLOR_FORMAT_NATIVE_WITH_ALPHA = LV_COLOR_FORMAT_ARGB8888,
// // #elif LV_COLOR_DEPTH == 32
// 	LV_COLOR_FORMAT_NATIVE            lv_color_format_t = LV_COLOR_FORMAT_XRGB8888
// 	LV_COLOR_FORMAT_NATIVE_WITH_ALPHA lv_color_format_t = LV_COLOR_FORMAT_ARGB8888
// #else
// #error "LV_COLOR_DEPTH should be 1, 8, 16, 24 or 32"
// #endif

)

//llgo:type C
type lv_image_dsc_t struct {
	Unused [0]byte
}

//llgo:type C
type lv_chart_series_t struct {
	Unused [0]byte
}

//llgo:type C
type lv_chart_cursor_t struct {
	Unused [0]byte
}

type lv_dir_t c.Int

const (
	LV_DIR_NONE   lv_dir_t = 0x00
	LV_DIR_LEFT   lv_dir_t = (1 << 0)
	LV_DIR_RIGHT  lv_dir_t = (1 << 1)
	LV_DIR_TOP    lv_dir_t = (1 << 2)
	LV_DIR_BOTTOM lv_dir_t = (1 << 3)
	LV_DIR_HOR    lv_dir_t = LV_DIR_LEFT | LV_DIR_RIGHT
	LV_DIR_VER    lv_dir_t = LV_DIR_TOP | LV_DIR_BOTTOM
	LV_DIR_ALL    lv_dir_t = LV_DIR_HOR | LV_DIR_VER
)

type lv_anim_enable_t c.Char

//llgo:type C
type lv_point_precise_t struct {
	X lv_value_precise_t
	Y lv_value_precise_t
}

//llgo:type C
type lv_display_t struct {
	Unused [0]byte
}

//llgo:type C
type lv_theme_t struct {
	Unused [0]byte
}

//llgo:type C
type lv_timer_t struct {
	Unused [0]byte
}
