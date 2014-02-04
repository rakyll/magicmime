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

// Package magicmime detects mimetypes using libmagic.
package magicmime

// #cgo LDFLAGS: -lmagic
// #include <stdlib.h>
// #include <magic.h>
import "C"

import (
	"errors"
	"unsafe"
)

type Magic struct {
	db C.magic_t
}

func New() (*Magic, error) {
	db := C.magic_open(C.int(0))
	C.magic_setflags(db, C.int(C.MAGIC_MIME_TYPE|C.MAGIC_SYMLINK|C.MAGIC_ERROR))
	if code := C.magic_load(db, nil); code != 0 {
		return nil, errors.New(C.GoString(C.magic_error(db)))
	}
	return &Magic{db}, nil
}

// TypeByFile looks up for a file's mimetype by its content.
// It uses a magic number database which is described in magic(5).
func (m *Magic) TypeByFile(filePath string) (string, error) {
	path := C.CString(filePath)
	defer C.free(unsafe.Pointer(path))
	out := C.magic_file(m.db, path)
	if out == nil {
		return "", errors.New(C.GoString(C.magic_error(m.db)))
	}
	return C.GoString(out), nil
}

// TypeByBuffer looks up for a blob's mimetype by its contents.
// It uses a magic number database which is described in magic(5).
func (m *Magic) TypeByBuffer(blob []byte) (string, error) {
	bytes := unsafe.Pointer(&blob[0])
	out := C.magic_buffer(m.db, bytes, C.size_t(len(blob)))
	if out == nil {
		return "", errors.New(C.GoString(C.magic_error(m.db)))
	}
	return C.GoString(out), nil
}

func (m *Magic) Close() {
	C.magic_close(m.db)
}
