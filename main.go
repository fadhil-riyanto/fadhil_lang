
package main

import (
	"io/ioutil"
	"log"
	"os"
	parser_engine "fr/fr_parser"
)

func main() {
	fadhil_bukaFile := os.Args
	if len(fadhil_bukaFile) < 2 {
		log.Fatal("maaf,gunakan ./fr {namafile}.fr")
	} else {
		files, err := ioutil.ReadFile(string(fadhil_bukaFile[1]))
		if err != nil {
			log.Fatal(err)
		} else {
			parser_engine.Parse(string(files))
		}
	}
}
