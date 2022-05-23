package main

import (
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	gen, err := generateTmplUI("../internal/static/ui_tmpl.gohtml")
	if err != nil {
		panic(err)
	}
	code := fmt.Sprintf("// Code generated DO NOT EDIT.\npackage webpb\n\n%s\n", gen)
	if err = os.WriteFile("../gen.go", []byte(code), 0775); err != nil {
		panic(err)
	}
	fmt.Println("File gen.go successfully generated")
}

func generateTmplUI(filepath string) (result string, err error) {
	var data []byte
	if data, err = os.ReadFile(filepath); err != nil {
		return
	}

	hexData := hex.EncodeToString(data)

	for i := 0; i < len(hexData); i++ {
		if i%2 == 0 {
			result += `\x`
		}
		result += hexData[i : i+1]
	}

	result = fmt.Sprintf("var _uiTmpl = []byte(\"%s\")", result)
	return
}
