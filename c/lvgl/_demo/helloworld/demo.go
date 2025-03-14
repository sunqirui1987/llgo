package main

import (
	"time"

	"github.com/goplus/llgo/c"
	"github.com/goplus/llgo/c/lvgl"
	"github.com/goplus/llgo/c/lvgl/display"
	"github.com/goplus/llgo/c/lvgl/drivers/sdl"
	"github.com/goplus/llgo/c/lvgl/widgets"
)

func main() {
	lvgl.LvglInit()

	disp := sdl.LvSdlWindowCreate(480, 320)
	sdl.LvSdlWindowSetZoom(disp, 1.0)
	sdl.LvSdlWindowSetTitle(disp, c.Str("Hello World"))

	label := widgets.LvLabelCreate(display.LvScreenActive())
	widgets.LvLabelSetText(label, c.Str("Hello World"))
	//widgets.LvObjAlign(label, lvgl.LV_ALIGN_CENTER, 0, 0)

	for {
		lvgl.LvTimerHandler()
		time.Sleep(time.Millisecond * 5)
	}

}
