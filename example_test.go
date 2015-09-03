package magicmime_test

import (
	"fmt"
	"testing"

	"github.com/rakyll/magicmime"
)

func Example_TypeByFile() {
	mm, err := magicmime.New(magicmime.MAGIC_MIME_TYPE | magicmime.MAGIC_SYMLINK | magicmime.MAGIC_ERROR)
	if err != nil {
		panic(err)
	}
	defer mm.Close()

	mimetype, err := mm.TypeByFile("/path/to/file")
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}

	fmt.Printf("%s -> %s\n", filepath, mimetype)
}
