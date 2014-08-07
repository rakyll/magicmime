package magicmime_test

import (
	"fmt"

	"github.com/rakyll/magicmime"
)

func main() {
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
