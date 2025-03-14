package main

import (
	"time"

	"github.com/goplus/llgo/c"
	"github.com/goplus/llgo/c/lvgl"
	"github.com/goplus/llgo/c/lvgl/core"
	"github.com/goplus/llgo/c/lvgl/display"
	"github.com/goplus/llgo/c/lvgl/drivers/sdl"
	"github.com/goplus/llgo/c/lvgl/widgets"
)

func main() {
	// Initialize LVGL
	lvgl.LvglInit()

	// Create a display
	disp := sdl.LvSdlWindowCreate(480, 320)
	sdl.LvSdlWindowSetZoom(disp, 1.0)
	sdl.LvSdlWindowSetTitle(disp, c.Str("Profile Page"))

	// Apply a theme
	// primaryColor := core.LvColorHex(0x343247)
	// secondaryColor := core.LvColorHex(0x343247)
	// defaulttheme.LvThemeDefaultInit(display.LvDisplayGetDefault(), (core.LvColorT)(primaryColor), (core.LvColorT)(secondaryColor),
	// 	0, nil)

	// Get the active screen
	screen := display.LvScreenActive()

	// Create a tab view for the top navigation
	tabview := widgets.LvTabviewCreate(screen)
	core.LvObjSetSize(tabview, c.Int32T(480), c.Int32T(320))
	core.LvObjSetPos(tabview, c.Int32T(0), c.Int32T(0))

	// Add tabs
	_ = widgets.LvTabviewAddTab(tabview, c.Str("Profile"))
	_ = widgets.LvTabviewAddTab(tabview, c.Str("Analytics"))
	_ = widgets.LvTabviewAddTab(tabview, c.Str("Shop"))

	//Create a floating action button
	fab := widgets.LvButtonCreate(screen)
	core.LvObjSetSize(fab, 50, 50)
	core.LvObjSetPos(fab, 410, 250)

	fabLabel := widgets.LvLabelCreate(fab)
	widgets.LvLabelSetText(fabLabel, c.Str("hello"))
	core.LvObjCenter(fabLabel)

	// img := widgets.LvImageCreate(screen)
	// imgSrc := c.Str("A:src/test_assets/test_img_lvgl_logo_with_exif_orientation_180.jpg")
	// widgets.LvImageSetSrc(img, imgSrc)

	chart := widgets.LvChartCreate(screen)
	widgets.LvChartSetType(chart, widgets.LV_CHART_TYPE_LINE) /*Show lines and points too*/
	core.LvObjSetSize(chart, 200, 150)
	core.LvObjCenter(chart)

	/*Create a LED and switch it OFF*/
	led1 := widgets.LvLedCreate(screen)
	core.LvObjAlign(led1, core.LV_ALIGN_CENTER, -80, 0)
	widgets.LvLedOff(led1)

	/*Copy the previous LED and set a brightness*/
	led2 := widgets.LvLedCreate(screen)
	core.LvObjAlign(led2, core.LV_ALIGN_CENTER, 0, 0)
	widgets.LvLedSetBrightness(led2, 150)
	widgets.LvLedSetColor(led2, core.LvColorHex(0x0000FF))

	/*Copy the previous LED and switch it ON*/
	led3 := widgets.LvLedCreate(screen)
	core.LvObjAlign(led3, core.LV_ALIGN_CENTER, -80, 0)
	widgets.LvLedOff(led3)

	// Main event loop
	for {
		lvgl.LvTimerHandler()
		time.Sleep(time.Millisecond * 5)
	}
}
