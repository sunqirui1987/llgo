package mono_theme

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
	core "github.com/goplus/llgo/c/lvgl/core"
	font "github.com/goplus/llgo/c/lvgl/font"
	themes "github.com/goplus/llgo/c/lvgl/themes"
)

//go:linkname LvThemeMonoInit C.lv_theme_mono_init
func LvThemeMonoInit(disp *core.LvDisplayT, dark c.Char, font *font.LvFontT) *themes.LvThemeT

//go:linkname LvThemeMonoGet C.lv_theme_mono_get
func LvThemeMonoGet() *themes.LvThemeT

//go:linkname LvThemeMonoIsInited C.lv_theme_mono_is_inited
func LvThemeMonoIsInited() c.Char

//go:linkname LvThemeMonoDeinit C.lv_theme_mono_deinit
func LvThemeMonoDeinit()
