package fr_parser

import (
	"fmt"
	mod_konsol "fr/src/konsol"
	"log"
	"strconv"
	"strings"
)

func Parse(syntax string) {
	data := make(map[string]string)
	split_baris := strings.Split(syntax, "\n")
	for counter, syntaxxx := range split_baris {
		if syntaxxx == "" {

		} else {
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
				sv := get_valuefunc[1]
				getvaridenty := sv[:2]
				if pkgname == "konsol" {
					if fnname == "print" {
						if getvaridenty == "$_" {
							awal := strings.Replace(sv, "$_", "", -1)
							mod_konsol.Print(data[awal])
						} else {
							mod_konsol.Print(vfunc[1])
						}
					} else if fnname == "println" {
						if getvaridenty == "$_" {
							awal := strings.Replace(sv, "$_", "", -1)
							mod_konsol.Println(data[awal])
						} else {
							mod_konsol.Println(vfunc[1])
						}
					} else {
						fmt.Println("fungsi \"" + fnname + "\" tidak ditermukan dibaris " + strconv.Itoa(counter+1))
					}
				} else {
					fmt.Println("package \"" + pkgname + "\" tidak ditermukan dibaris " + strconv.Itoa(counter+1))
				}
			} else if indexke1dan2 == "$_" {
				awal := strings.Replace(syntaxxx, "$_", "", -1)
				getvars := strings.Split(awal, "=")

				data[getvars[0]] = getvars[1]

			} else if indexke1dan2 == "//" {
			}

			//if get_kurung_awal[1] == "konsol" {
			// fmt.Println()
			//}
		}

	}

}
