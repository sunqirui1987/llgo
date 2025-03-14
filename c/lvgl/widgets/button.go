package widgets

import (
	_ "unsafe"

	core "github.com/goplus/llgo/c/lvgl/core"
)

//go:linkname LvButtonCreate C.lv_button_create
func LvButtonCreate(parent *core.LvObjT) *core.LvObjT
