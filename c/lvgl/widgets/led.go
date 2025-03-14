package widgets

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

/**
 * Create a led object
 * @param parent    pointer to an object, it will be the parent of the new led
 * @return          pointer to the created led
 */
// lv_obj_t * lv_led_create(lv_obj_t * parent);

//go:linkname LvLedCreate C.lv_led_create
func LvLedCreate(parent *c.Void) *c.Void

/**
 * Set the color of the LED
 * @param led       pointer to a LED object
 * @param color     the color of the LED
 */
// void lv_led_set_color(lv_obj_t * led, lv_color_t color);

//go:linkname LvLedSetColor C.lv_led_set_color
func LvLedSetColor(led *c.Void, color *c.Void)

/**
 * Set the brightness of a LED object
 * @param led       pointer to a LED object
 * @param bright    LV_LED_BRIGHT_MIN (max. dark) ... LV_LED_BRIGHT_MAX (max. light)
 */
// void lv_led_set_brightness(lv_obj_t * led, uint8_t bright);

//go:linkname LvLedSetBrightness C.lv_led_set_brightness
func LvLedSetBrightness(led *c.Void, bright c.Uint)

/**
 * Light on a LED
 * @param led       pointer to a LED object
 */
// void lv_led_on(lv_obj_t * led);

//go:linkname LvLedOn C.lv_led_on
func LvLedOn(led *c.Void)

/**
 * Light off a LED
 * @param led       pointer to a LED object
 */
// void lv_led_off(lv_obj_t * led);

//go:linkname LvLedOff C.lv_led_off
func LvLedOff(led *c.Void)

/**
 * Toggle the state of a LED
 * @param led       pointer to a LED object
 */
// void lv_led_toggle(lv_obj_t * led);

//go:linkname LvLedToggle C.lv_led_toggle
func LvLedToggle(led *c.Void)

/**
 * Get the brightness of a LED object
 * @param obj       pointer to LED object
 * @return bright   0 (max. dark) ... 255 (max. light)
 */
// uint8_t lv_led_get_brightness(const lv_obj_t * obj);

//go:linkname LvLedGetBrightness C.lv_led_get_brightness
func LvLedGetBrightness(obj *c.Void) c.Uint
