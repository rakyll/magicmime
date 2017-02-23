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

// +build linux darwin freebsd

package magicmime_test

import (
	"log"

	"github.com/rakyll/magicmime"
)

func ExampleTypeByFile() {
	if err := magicmime.Open(magicmime.MAGIC_MIME_TYPE | magicmime.MAGIC_SYMLINK | magicmime.MAGIC_ERROR); err != nil {
		log.Fatal(err)
	}
	defer magicmime.Close()

	mimetype, err := magicmime.TypeByFile("/path/to/file")
	if err != nil {
		log.Fatalf("error occured during type lookup: %v", err)
	}

	log.Printf("mime-type: %v", mimetype)
}
