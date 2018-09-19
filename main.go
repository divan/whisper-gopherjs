package main

import "github.com/gopherjs/gopherjs/js"

func main() {
	js.Global.Set("WhisperGo", map[string]interface{}{
		"EncryptSymmetric": EncryptSymmetric,
		"Key":              Key,
		"KeyEncrypt":       KeyEncrypt,
	})

	js.Module.Get("exports").Set(
		"EncryptSymmetric", EncryptSymmetric,
	)
	js.Module.Get("exports").Set(
		"Key", Key,
	)
	js.Module.Get("exports").Set(
		"KeyEncrypt", KeyEncrypt,
	)
}
