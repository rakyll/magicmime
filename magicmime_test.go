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

// +build linux darwin

package magicmime

import (
	"encoding/base64"
	"testing"
)

// Tests a gif file.
func TestGifFile(t *testing.T) {
	testFile(t, UNSYNCHRONIZED, "./testdata/sample.gif", "image/gif")
	testFile(t, SYNCHRONIZED, "./testdata/sample.gif", "image/gif")
}

// Tests a jpeg file.
func TestJpegFile(t *testing.T) {
	testFile(t, UNSYNCHRONIZED, "./testdata/sample.jpg", "image/jpeg")
	testFile(t, SYNCHRONIZED, "./testdata/sample.jpg", "image/jpeg")
}

// Tests a png file.
func TestPngFile(t *testing.T) {
	testFile(t, UNSYNCHRONIZED, "./testdata/sample.png", "image/png")
	testFile(t, SYNCHRONIZED, "./testdata/sample.png", "image/png")
}

// Tests a pdf file.
func TestPdfFile(t *testing.T) {
	testFile(t, UNSYNCHRONIZED, "./testdata/sample.pdf", "application/pdf")
	testFile(t, SYNCHRONIZED, "./testdata/sample.pdf", "application/pdf")
}

// Tests a plain text file.
func TestTextFile(t *testing.T) {
	testFile(t, UNSYNCHRONIZED, "./testdata/sample.txt", "text/plain")
	testFile(t, SYNCHRONIZED, "./testdata/sample.txt", "text/plain")
}

// Tests a gzipped tar file.
func TestGzippedTarFile(t *testing.T) {
	testFile(t, UNSYNCHRONIZED, "./testdata/sample.tar.gz", "application/x-gzip")
	testFile(t, SYNCHRONIZED, "./testdata/sample.tar.gz", "application/x-gzip")
}

// Tests a zip file.
func TestZipFile(t *testing.T) {
	testFile(t, UNSYNCHRONIZED, "./testdata/sample.zip", "application/zip")
	testFile(t, SYNCHRONIZED, "./testdata/sample.zip", "application/zip")
}

// Tests a gif buffer with synchronized access
func TestGifBufferSYNCHRONIZED(t *testing.T) {
	db, err := Open(SYNCHRONIZED, MAGIC_MIME_TYPE|MAGIC_SYMLINK|MAGIC_ERROR)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	b64Gif := "R0lGODlhAQABAIAAAAUEBAAAACwAAAAAAQABAAACAkQBADs="
	expected := "image/gif"
	gif, err := base64.StdEncoding.DecodeString(b64Gif)
	if err != nil {
		panic(err)
	}
	mimetype, err := db.TypeByBuffer(gif)
	if err != nil {
		panic(err)
	}
	if mimetype != expected {
		t.Errorf("expected %s; got %s.", expected, mimetype)
	}
}

// Tests a gif buffer with unsynchronized access
func TestGifBufferUNSYNCHRONIZED(t *testing.T) {
	db, err := Open(UNSYNCHRONIZED, MAGIC_MIME_TYPE|MAGIC_SYMLINK|MAGIC_ERROR)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	b64Gif := "R0lGODlhAQABAIAAAAUEBAAAACwAAAAAAQABAAACAkQBADs="
	expected := "image/gif"
	gif, err := base64.StdEncoding.DecodeString(b64Gif)
	if err != nil {
		panic(err)
	}
	mimetype, err := db.TypeByBuffer(gif)
	if err != nil {
		panic(err)
	}
	if mimetype != expected {
		t.Errorf("expected %s; got %s.", expected, mimetype)
	}
}

func testFile(tb testing.TB, syncMode SyncMode, path string, expected string) {
	db, err := Open(syncMode, MAGIC_MIME_TYPE|MAGIC_SYMLINK|MAGIC_ERROR)
	if err != nil {
		tb.Fatal(err)
	}
	defer db.Close()

	mimetype, err := db.TypeByFile(path)
	if err != nil {
		panic(err)
	}
	if mimetype != expected {
		tb.Errorf("expected %s; got %s.", expected, mimetype)
	}
}

func TestMissingFile(t *testing.T) {
	db, err := Open(UNSYNCHRONIZED, MAGIC_MIME_TYPE|MAGIC_SYMLINK|MAGIC_ERROR)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	_, err = db.TypeByFile("missingFile.txt")
	if err == nil {
		t.Error("no error for missing file")
	}
}

func BenchmarkZipFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testFile(b, UNSYNCHRONIZED, "./testdata/sample.zip", "application/zip")
	}
}

func BenchmarkSYNCHRONIZEDZipFile(b *testing.B) {
	db, err := Open(SYNCHRONIZED, MAGIC_MIME_TYPE|MAGIC_SYMLINK|MAGIC_ERROR)
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	for i := 0; i < b.N; i++ {
		mimetype, err := db.TypeByFile("./testdata/sample.zip")
		if err != nil {
			panic(err)
		}
		if mimetype != "application/zip" {
			b.Errorf("expected %s; got %s.", "application/zip", mimetype)
		}
	}
}
