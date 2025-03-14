package widgets

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

/**
 * Create an animation image objects
 * @param parent pointer to an object, it will be the parent of the new button
 * @return pointer to the created animation image object
 */
// lv_obj_t * lv_animimg_create(lv_obj_t * parent);

//go:linkname LvAnimimgCreate C.lv_animimg_create
func LvAnimimgCreate(parent *c.Void) *c.Void

/*=====================
 * Setter functions
 *====================*/

/**
 * Set the image animation images source.
 * @param img   pointer to an animation image object
 * @param dsc   pointer to a series images
 * @param num   images' number
 */

// void lv_animimg_set_src(lv_obj_t * img, const void * dsc[], size_t num);

//go:linkname LvAnimimgSetSrc C.lv_animimg_set_src
func LvAnimimgSetSrc(img *c.Void, dsc []c.Void, num c.Int)

/**
 * Startup the image animation.
 * @param obj   pointer to an animation image object
 */

// void lv_animimg_start(lv_obj_t * obj);

//go:linkname LvAnimimgStart C.lv_animimg_start
func LvAnimimgStart(obj *c.Void)

/**
 * Set the image animation duration time. unit:ms
 * @param img       pointer to an animation image object
 * @param duration  the duration in milliseconds
 */

// void lv_animimg_set_duration(lv_obj_t * img, uint32_t duration);

//go:linkname LvAnimimgSetDuration C.lv_animimg_set_duration
func LvAnimimgSetDuration(img *c.Void, duration c.Uint)

/**
 * Set the image animation repeatedly play times.
 * @param img       pointer to an animation image object
 * @param count     the number of times to repeat the animation
 */

// void lv_animimg_set_repeat_count(lv_obj_t * img, uint32_t count);

//go:linkname LvAnimimgSetRepeatCount C.lv_animimg_set_repeat_count
func LvAnimimgSetRepeatCount(img *c.Void, count c.Uint)

/*=====================
 * Getter functions
 *====================*/

/**
 * Get the image animation images source.
 * @param img   pointer to an animation image object
 * @return a    pointer that will point to a series images
 */

// const void ** lv_animimg_get_src(lv_obj_t * img);

//go:linkname LvAnimimgGetSrc C.lv_animimg_get_src
func LvAnimimgGetSrc(img *c.Void) *c.Void

/**
 * Get the image animation images source.
 * @param img   pointer to an animation image object
 * @return      the number of source images
 */

// uint8_t lv_animimg_get_src_count(lv_obj_t * img);

//go:linkname LvAnimimgGetSrcCount C.lv_animimg_get_src_count
func LvAnimimgGetSrcCount(img *c.Void) c.Uint

/**
 * Get the image animation duration time. unit:ms
 * @param img   pointer to an animation image object
 * @return      the animation duration time
 */

// uint32_t lv_animimg_get_duration(lv_obj_t * img);

//go:linkname LvAnimimgGetDuration C.lv_animimg_get_duration
func LvAnimimgGetDuration(img *c.Void) c.Uint

/**
 * Get the image animation repeat play times.
 * @param img   pointer to an animation image object
 * @return      the repeat count
 */

// uint32_t lv_animimg_get_repeat_count(lv_obj_t * img);

//go:linkname LvAnimimgGetRepeatCount C.lv_animimg_get_repeat_count
func LvAnimimgGetRepeatCount(img *c.Void) c.Uint

/**
 * Get the image animation underlying animation.
 * @param img   pointer to an animation image object
 * @return      the animation reference
 */

// lv_anim_t * lv_animimg_get_anim(lv_obj_t * img);

//go:linkname LvAnimimgGetAnim C.lv_animimg_get_anim
func LvAnimimgGetAnim(img *c.Void) *c.Void
