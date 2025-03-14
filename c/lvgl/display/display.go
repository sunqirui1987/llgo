package display

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

/**
 * Get the active screen of the default display
 * @return          pointer to the active screen
 */
//lv_obj_t * lv_screen_active(void);
//
//go:linkname LvScreenActive C.lv_screen_active
func LvScreenActive() *c.Void
