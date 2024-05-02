/*
 * Copyright (c) 2024 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package runtime

import (
	"unsafe"

	"github.com/goplus/llgo/internal/abi"
	"github.com/goplus/llgo/internal/runtime/c"
)

// Alloc allocates memory.
func Alloc(size uintptr) unsafe.Pointer {
	return c.Malloc(size)
}

// TracePanic prints panic message.
func TracePanic(v Interface) {
	kind := abi.Kind(v.tab._type.Kind_)
	switch {
	case kind == abi.String:
		stringTracef(c.Stderr, c.Str("panic: %s\n"), *(*String)(v.data))
	}
	// TODO(xsw): other message type
}

func stringTracef(fp c.FilePtr, format *c.Char, s String) {
	cs := c.Alloca(uintptr(s.len) + 1)
	c.Fprintf(fp, format, CStrCopy(cs, s))
}
