package lvgl

import (
	_ "unsafe"
)

const (
	LLGoPackage = "link: $(pkg-config --libs sdl2 lvgl); -llvgl -lSDL2"
)

const (
	LV_ALIGN_DEFAULT      = 0
	LV_ALIGN_TOP_LEFT     = 1
	LV_ALIGN_TOP_MID      = 2
	LV_ALIGN_TOP_RIGHT    = 3
	LV_ALIGN_BOTTOM_LEFT  = 4
	LV_ALIGN_BOTTOM_MID   = 5
	LV_ALIGN_BOTTOM_RIGHT = 6
	LV_ALIGN_LEFT_MID     = 7
	LV_ALIGN_RIGHT_MID    = 8
	LV_ALIGN_CENTER       = 9

	LV_ALIGN_OUT_TOP_LEFT     = 10
	LV_ALIGN_OUT_TOP_MID      = 11
	LV_ALIGN_OUT_TOP_RIGHT    = 12
	LV_ALIGN_OUT_BOTTOM_LEFT  = 13
	LV_ALIGN_OUT_BOTTOM_MID   = 14
	LV_ALIGN_OUT_BOTTOM_RIGHT = 15
	LV_ALIGN_OUT_LEFT_TOP     = 16
	LV_ALIGN_OUT_LEFT_MID     = 17
	LV_ALIGN_OUT_LEFT_BOTTOM  = 18
	LV_ALIGN_OUT_RIGHT_TOP    = 19
	LV_ALIGN_OUT_RIGHT_MID    = 20
	LV_ALIGN_OUT_RIGHT_BOTTOM = 21
)

//go:linkname LvglInit C.lv_init
func LvglInit()

//go:linkname LvglDeinit C.lv_deinit
func LvglDeinit()

//go:linkname LvglRun C.lv_run
func LvglRun()
