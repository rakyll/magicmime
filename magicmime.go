// Copyright 2013 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package rrqueue provides a container for priority queues
// and a simple round-robin scheduled consumer.

package magicmime

// #cgo LDFLAGS: -lmagic
// #include <stdlib.h>
// #include <magic.h>
import "C"

import (
	"errors"
	"unsafe"
)

var (
	ErrNoDB   = errors.New("no magic db to load")
	ErrLookup = errors.New("error during type lookup")
)

// TypeByContent looks up for a file's mimetype by its content.
// It uses a magic number database which is described in magic(5).
func TypeByContent(filePath string) (string, error) {
	// TODO: load db once, use for many lookups
	cookie := C.magic_open(C.int(0))
	defer C.magic_close(cookie)

	C.magic_setflags(cookie, C.int(C.MAGIC_MIME_TYPE))
	if code := C.magic_load(cookie, nil); code != 0 {
		return "", ErrNoDB
	}

	path := C.CString(filePath)
	defer C.free(unsafe.Pointer(path))
	out := C.magic_file(cookie, path)
	if out == nil {
		return "", ErrLookup // TODO: respond with what magic_error returns
	}
	return C.GoString(out), nil
}

// TODO: // add mime discovery from a byte buffer
