package core

import (
	_ "unsafe"

	"github.com/goplus/llgo/c"
)

//go:linkname LvColorHex C.lv_color_hex
func LvColorHex(value c.Int) LvColorT
