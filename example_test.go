package magicmime_test

import (
	"log"

	"github.com/rakyll/magicmime"
)

func Example_TypeByFile() {
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
