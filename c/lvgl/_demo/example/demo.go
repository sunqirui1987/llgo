package main

import (
	"log"
	"time"

	"github.com/goplus/llgo/c"
	"github.com/goplus/llgo/c/lvgl"
	"github.com/goplus/llgo/c/lvgl/drivers/sdl"
)

func main() {
	// Initialize LVGL
	lvgl.LvglInit()

	// Create a display
	disp := sdl.LvSdlWindowCreate(480, 320)
	sdl.LvSdlWindowSetZoom(disp, 1.0)
	sdl.LvSdlWindowSetTitle(disp, c.Str("Profile Page"))
	log.Println("disp", disp)

	lv_example_style_1()

	// Main event loop
	for {
		lvgl.LvTimerHandler()
		time.Sleep(time.Millisecond * 5)
	}
}
