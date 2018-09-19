package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Module.Get("exports").Set(
		"EncryptSymmetric", EncryptSymmetric,
	)
	js.Module.Get("exports").Set(
		"Key", Key,
	)
}
