package main

import (
	"fmt"
	"magicmime"
	"os"
)

var mgc = "/usr/share/misc/magic.mgc"

func init() {
	if mgcpath := os.Getenv("MGC"); mgcpath != "" {
		mgc = mgcpath
	}
}

func main() {
	d, err := magicmime.NewWithPath(mgc, magicmime.MAGIC_NONE)
	if err != nil {
		fmt.Println("NewDecoderWithPath", err)
		return
	}
	t, err := d.TypeByFile(os.Args[1])
	fmt.Println(t, err)
}
