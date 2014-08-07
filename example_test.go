package magicmime_test

import (
	"fmt"
	"testing"

	"github.com/rakyll/magicmime"
)

// TODO: Remove after Go 1.4.
// Related to https://codereview.appspot.com/107320046
func TestA(t *testing.T) {}

func Example_1() {
	mm, err := magicmime.New(magicmime.MAGIC_MIME_TYPE | magicmime.MAGIC_SYMLINK | magicmime.MAGIC_ERROR)
	if err != nil {
		panic(err)
	}

	filepath := "/bin/ls"

	mimetype, err := mm.TypeByFile(filepath)
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}

	fmt.Printf("%s -> %s\n", filepath, mimetype)
}
