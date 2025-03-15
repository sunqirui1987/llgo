package core

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

/**
 * Initialize a style
 * @param style pointer to a style to initialize
 * @note Do not call `lv_style_init` on styles that already have some properties
 *       because this function won't free the used memory, just sets a default state for the style.
 *       In other words be sure to initialize styles only once!
 */
//go:linkname LvStyleInit C.lv_style_init
func LvStyleInit(style *LvStyleT)

/**
 * Clear all properties from a style and free all allocated memories.
 * @param style pointer to a style
 */
//go:linkname LvStyleReset C.lv_style_reset
func LvStyleReset(style *LvStyleT)

//go:linkname LvStyleCopy C.lv_style_copy
func LvStyleCopy(dst *LvStyleT, src *LvStyleT)

//go:linkname LvStyleSetSize C.lv_style_set_size
func LvStyleSetSize(style *LvStyleT, width c.Int32T, height c.Int32T)

//go:linkname LvStyleSetPadAll C.lv_style_set_pad_all
func LvStyleSetPadAll(style *LvStyleT, value c.Int32T)

//go:linkname LvStyleSetPadHor C.lv_style_set_pad_hor
func LvStyleSetPadHor(style *LvStyleT, value c.Int32T)

//go:linkname LvStyleSetPadVer C.lv_style_set_pad_ver
func LvStyleSetPadVer(style *LvStyleT, value c.Int32T)

//go:linkname LvStyleSetPadGap C.lv_style_set_pad_gap
func LvStyleSetPadGap(style *LvStyleT, value c.Int32T)

//go:linkname LvStyleSetMarginHor C.lv_style_set_margin_hor
func LvStyleSetMarginHor(style *LvStyleT, value c.Int32T)

//go:linkname LvStyleSetMarginVer C.lv_style_set_margin_ver
func LvStyleSetMarginVer(style *LvStyleT, value c.Int32T)

//go:linkname LvStyleSetMarginAll C.lv_style_set_margin_all
func LvStyleSetMarginAll(style *LvStyleT, value c.Int32T)

//go:linkname LvStyleSetTransformScale C.lv_style_set_transform_scale
func LvStyleSetTransformScale(style *LvStyleT, value c.Int32T)

//go:linkname LvStyleSetTransformScaleAll C.lv_style_set_transform_scale_all
func LvStyleSetTransformScaleAll(style *LvStyleT, value c.Int32T)

/**
 * @brief Check if the style property has a specified behavioral flag.
 *
 * Do not pass multiple flags to this function as backwards-compatibility is not guaranteed
 * for that.
 *
 * @param prop Property ID
 * @param flag Flag
 * @return true if the flag is set for this property
 */
//go:linkname LvStylePropHasFlag C.lv_style_prop_has_flag
func LvStylePropHasFlag(prop LvStylePropT, flag c.Uint8T) bool
