// Copyright (C) 2016 Makoto Imaizumi <Suisui@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Japanese Katakana to Hiragaana converter
// Usage) kana2hira filename
// Reerence http://www.serendip.ws/archives/6307

package main

import(
	"os"
	"bufio"
	"strings"
	"unicode"
	"fmt"
)


var kanaConv = unicode.SpecialCase{
    // ひらがなをカタカナに変換
    unicode.CaseRange{
        0x3041, // Lo: ぁ
        0x3093, // Hi: ん
        [unicode.MaxCase]rune{
            0x30a1 - 0x3041, // UpperCase でカタカナに変換
            0,               // LowerCase では変換しない
            0x30a1 - 0x3041, // TitleCase でカタカナに変換
        },
    },
    // カタカナをひらがなに変換
    unicode.CaseRange{
        0x30a1, // Lo: ァ
        0x30f3, // Hi: ン
        [unicode.MaxCase]rune{
            0,               // UpperCase では変換しない
            0x3041 - 0x30a1, // LowerCase でひらがなに変換
            0,               // TitleCase では変換しない
        },
    },
}

func main () {
        var f *os.File
        var err error

        if len(os.Args) < 2 {
                f = os.Stdin
        } else {
                f, err = os.Open(os.Args[1])
                if err != nil {
                        panic(err)
                }
                defer f.Close()
        }
        rd := bufio.NewReaderSize(f, 8192)
        for line := ""; err == nil ; line, err = rd.ReadString('\n') {
                if err != nil {
                        panic(err)
                }
		fmt.Printf("%s",strings.ToLowerSpecial(kanaConv, line))
	}
}



