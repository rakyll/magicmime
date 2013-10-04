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

// newMagic creates a magic_t handle and loads database
func newMagic() (C.magic_t, error) {
	cookie := C.magic_open(C.int(0))
	C.magic_setflags(cookie, C.int(C.MAGIC_MIME_TYPE|C.MAGIC_SYMLINK))
	if code := C.magic_load(cookie, nil); code != 0 {
		return nil, ErrNoDB
	}
	return cookie, nil
}

// TypeByFile looks up for a file's mimetype by its content.
// It uses a magic number database which is described in magic(5).
func TypeByFile(filePath string) (string, error) {
	// TODO: load db once, use for many lookups
	cookie, err := newMagic()
	defer C.magic_close(cookie)

	if err != nil {
		return "", err
	}

	path := C.CString(filePath)
	defer C.free(unsafe.Pointer(path))
	out := C.magic_file(cookie, path)
	if out == nil {
		return "", ErrLookup // TODO: respond with what magic_error returns
	}
	return C.GoString(out), nil
}

// TypeByBuffer looks up for a blob's mimetype by its contents.
// It uses a magic number database which is described in magic(5).
func TypeByBuffer(blob []byte) (string, error) {
	// TODO: load db once, use for many lookups
	cookie, err := newMagic()
	defer C.magic_close(cookie)

	if err != nil {
		return "", err
	}

	bytes := unsafe.Pointer(&blob[0])
	out := C.magic_buffer(cookie, bytes, C.size_t(len(blob)))
	if out == nil {
		return "", ErrLookup // TODO: respond with what magic_error returns
	}
	return C.GoString(out), nil
}
