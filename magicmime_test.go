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

import (
	"encoding/base64"
	"testing"
)

var b64Gif = "R0lGODlhAQABAIAAAAUEBAAAACwAAAAAAQABAAACAkQBADs="

func TestGifContents(t *testing.T) {
	expected := "image/gif"
	gif, err := base64.StdEncoding.DecodeString(b64Gif)
	if err != nil {
		panic(err)
	}
	mimetype, err := TypeByContents(gif)
	if err != nil {
		panic(err)
	}
	if mimetype != expected {
		t.Errorf("expected %s; got %s.", expected, mimetype)
	}
}
