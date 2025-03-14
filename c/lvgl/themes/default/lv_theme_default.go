package default_theme

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llgo/c/lvgl/core"
	font "github.com/goplus/llgo/c/lvgl/font"
	themes "github.com/goplus/llgo/c/lvgl/themes"
)

//go:linkname LvThemeDefaultInit C.lv_theme_default_init
func LvThemeDefaultInit(disp *core.LvDisplayT, colorPrimary core.LvColorT, colorSecondary core.LvColorT, dark c.Char, font *font.LvFontT) *themes.LvThemeT

//go:linkname LvThemeDefaultGet C.lv_theme_default_get
func LvThemeDefaultGet() *themes.LvThemeT

//go:linkname LvThemeDefaultIsInited C.lv_theme_default_is_inited
func LvThemeDefaultIsInited() c.Char

//go:linkname LvThemeDefaultDeinit C.lv_theme_default_deinit
func LvThemeDefaultDeinit()
