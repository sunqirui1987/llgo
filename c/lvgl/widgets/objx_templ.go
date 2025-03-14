package widgets

import (
	_ "unsafe"

	core "github.com/goplus/llgo/c/lvgl/core"
)

//go:linkname LvTemplCreate C.lv_templ_create
func LvTemplCreate(parent *core.LvObjT) *core.LvObjT
