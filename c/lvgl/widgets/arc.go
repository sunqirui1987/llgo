package widgets

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

type lv_value_precise_t = c.Float

type lv_arc_mode_t = c.Int

const (
	LV_ARC_MODE_NORMAL lv_arc_mode_t = iota
	LV_ARC_MODE_SYMMETRICAL
	LV_ARC_MODE_REVERSE
)

/**
 * Create an arc object
 * @param parent    pointer to an object, it will be the parent of the new arc
 * @return          pointer to the created arc
 */
// lv_obj_t * lv_arc_create(lv_obj_t * parent);

//go:linkname LvArcCreate C.lv_arc_create
func LvArcCreate(parent *c.Void) *c.Void

/*======================
 * Add/remove functions
 *=====================*/

/*=====================
 * Setter functions
 *====================*/

/**
 * Set the start angle of an arc. 0 deg: right, 90 bottom, etc.
 * @param obj       pointer to an arc object
 * @param start     the start angle. (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 */
// void lv_arc_set_start_angle(lv_obj_t * obj, lv_value_precise_t start);

//go:linkname LvArcSetStartAngle C.lv_arc_set_start_angle
func LvArcSetStartAngle(obj *c.Void, start lv_value_precise_t)

/**
 * Set the end angle of an arc. 0 deg: right, 90 bottom, etc.
 * @param obj       pointer to an arc object
 * @param end       the end angle  (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 */
// void lv_arc_set_end_angle(lv_obj_t * obj, lv_value_precise_t end);

//go:linkname LvArcSetEndAngle C.lv_arc_set_end_angle
func LvArcSetEndAngle(obj *c.Void, end lv_value_precise_t)

/**
 * Set the start and end angles
 * @param obj       pointer to an arc object
 * @param start     the start angle  (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 * @param end       the end angle  (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 */
// void lv_arc_set_angles(lv_obj_t * obj, lv_value_precise_t start, lv_value_precise_t end);

//go:linkname LvArcSetAngles C.lv_arc_set_angles
func LvArcSetAngles(obj *c.Void, start lv_value_precise_t, end lv_value_precise_t)

/**
 * Set the start angle of an arc background. 0 deg: right, 90 bottom, etc.
 * @param obj       pointer to an arc object
 * @param start     the start angle  (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 */
// void lv_arc_set_bg_start_angle(lv_obj_t * obj, lv_value_precise_t start);

//go:linkname LvArcSetBgStartAngle C.lv_arc_set_bg_start_angle
func LvArcSetBgStartAngle(obj *c.Void, start lv_value_precise_t)

/**
 * Set the start angle of an arc background. 0 deg: right, 90 bottom etc.
 * @param obj       pointer to an arc object
 * @param end       the end angle  (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 */
// void lv_arc_set_bg_end_angle(lv_obj_t * obj, lv_value_precise_t end);

//go:linkname LvArcSetBgEndAngle C.lv_arc_set_bg_end_angle
func LvArcSetBgEndAngle(obj *c.Void, end lv_value_precise_t)

/**
 * Set the start and end angles of the arc background
 * @param obj       pointer to an arc object
 * @param start     the start angle  (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 * @param end       the end angle  (if `LV_USE_FLOAT` is enabled it can be fractional too.)
 */
// void lv_arc_set_bg_angles(lv_obj_t * obj, lv_value_precise_t start, lv_value_precise_t end);

//go:linkname LvArcSetBgAngles C.lv_arc_set_bg_angles
func LvArcSetBgAngles(obj *c.Void, start lv_value_precise_t, end lv_value_precise_t)

/**
 * Set the rotation for the whole arc
 * @param obj           pointer to an arc object
 * @param rotation      rotation angle
 */
// void lv_arc_set_rotation(lv_obj_t * obj, int32_t rotation);

//go:linkname LvArcSetRotation C.lv_arc_set_rotation
func LvArcSetRotation(obj *c.Void, rotation c.Int)

/**
 * Set the type of arc.
 * @param obj       pointer to arc object
 * @param type      arc's mode
 */
// void lv_arc_set_mode(lv_obj_t * obj, lv_arc_mode_t type);

//go:linkname LvArcSetMode C.lv_arc_set_mode
func LvArcSetMode(obj *c.Void, _type lv_arc_mode_t)

/**
 * Set a new value on the arc
 * @param obj       pointer to an arc object
 * @param value     new value
 */
// void lv_arc_set_value(lv_obj_t * obj, int32_t value);

//go:linkname LvArcSetValue C.lv_arc_set_value
func LvArcSetValue(obj *c.Void, value c.Int)

/**
 * Set minimum and the maximum values of an arc
 * @param obj       pointer to the arc object
 * @param min       minimum value
 * @param max       maximum value
 */
// void lv_arc_set_range(lv_obj_t * obj, int32_t min, int32_t max);

//go:linkname LvArcSetRange C.lv_arc_set_range
func LvArcSetRange(obj *c.Void, min c.Int, max c.Int)

/**
 * Set a change rate to limit the speed how fast the arc should reach the pressed point.
 * @param obj       pointer to an arc object
 * @param rate      the change rate
 */
// void lv_arc_set_change_rate(lv_obj_t * obj, uint32_t rate);

//go:linkname LvArcSetChangeRate C.lv_arc_set_change_rate
func LvArcSetChangeRate(obj *c.Void, rate c.Uint)

/**
 * Set an offset angle for the knob
 * @param obj       pointer to an arc object
 * @param offset    knob offset from main arc in degrees
 */
// void lv_arc_set_knob_offset(lv_obj_t * obj, int32_t offset);

//go:linkname LvArcSetKnobOffset C.lv_arc_set_knob_offset
func LvArcSetKnobOffset(obj *c.Void, offset c.Int)
