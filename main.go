package main

import "github.com/gopherjs/gopherjs/js"

func main() {
	js.Global.Set("WhisperGo", map[string]interface{}{
		"EncryptSymmetric": EncryptSymmetric,
		"Key":              Key,
	})
}
