package fr_parser

import (
	mod_konsol "fr/src/konsol"
	"log"
	"strconv"
	"strings"
)

func Parse(syntax string) {
	split_baris := strings.Split(syntax, "\n")
	for counter, syntaxxx := range split_baris {
		indexke1dan2 := syntaxxx[:2]
		if indexke1dan2 == "f_" {
			//parse fungsi
			awal := strings.Replace(syntaxxx, "f_", "", -1)
			get_kurung_awal := strings.Split(awal, "->")
			pkgname := get_kurung_awal[0]
			fnname := get_kurung_awal[1]
			get_valuefunc := strings.Split(awal, "::")
			if len(get_valuefunc) < 3 {
				log.Fatal("maaf, syntax error dibaris " + strconv.Itoa(counter+1))
			}
			vfunc := get_valuefunc

			if pkgname == "konsol" {
				if fnname == "print" {
					mod_konsol.Print(vfunc[1])
				} else if fnname == "println" {
					mod_konsol.Println(vfunc[1])
				}

			}
		}

		//if get_kurung_awal[1] == "konsol" {
		// fmt.Println()
		//}
	}

}
