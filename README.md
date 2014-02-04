# magicmime

`magicmime` is a Go package which allows you to discover a file's mimetype by looking for magic numbers in its content. It could be used as a supplementary for Go's [`mime`](http://golang.org/pkg/mime/) package which only interprets the file extension to detect mimetypes. Internally, it implements [libmagic(3)](http://linux.die.net/man/3/libmagic) bindings.

Tested on Linux and Mac OS X, should be working on BSD. You could be able to build and make it working with Cygwin on Windows.

## Usage
In order to start, go get this repository:
~~~ go
go get github.com/rakyll/magicmime
~~~
You may need to install devel packages for `libmagic`. On Debian, Ubuntu and CentOS, get `libmagic-dev` package from your package manager. On Mac OS X get `libmagic` via Homebrew: `brew install libmagic`. If you don't have the required dev packages, compilation will be terminated by an error saying `magic.h` cannot be found.

~~~ go
import "github.com/rakyll/magicmime"
mm, err := magicmime.New()
mimetype, err := mm.TypeByFile("/path/to/file")
~~~

Available functions are:

~~~ go
func (m *Magic) TypeByFile(filePath string) (string, error)
func (m *Magic) TypeByBuffer(buffer []byte) (string, error)
~~~
    
## License
    Copyright 2013 Google Inc. All Rights Reserved.
    
    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at
    
         http://www.apache.org/licenses/LICENSE-2.0
    
    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.