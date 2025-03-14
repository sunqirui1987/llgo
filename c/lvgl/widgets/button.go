package widgets

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

/**
 * Create a button object
 * @param parent    pointer to an object, it will be the parent of the new button
 * @return          pointer to the created button
 */
// lv_obj_t * lv_button_create(lv_obj_t * parent);

//go:linkname LvButtonCreate C.lv_button_create
func LvButtonCreate(parent *c.Void) *c.Void
