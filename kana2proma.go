// Copyright (C) 2016 Makoto Imaizumi <Suisui@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Japanese Hiragana to Hepburn romanization (
// Passport type, Ministry of Foreign Affairs Passport Standard)
// convert script
// Usage) kana2proma filename

// Kana2PassportRoman returns romanization strings and "hint"
// hint shows output may have long bowel omit case.
// such as おおの as ONO,  こうの as KONO.

package main

import(
	"fmt"
	"bufio"
	"os"
)

var PassportMap = map[string]string{
	"きゃ" : "kya",
	"きゅ" : "kyu",
	"きょ" : "kyo",
	"しゃ" : "sha",
	"しゅ" : "shu",
	"しょ" : "sho",
	"しぇ" : "sye",
	"ちゃ" : "cha",
	"ちゅ" : "chu",
	"ちょ" : "cho",
	"にゃ" : "nya",
	"にゅ" : "nyu",
	"にょ" : "nyo",
	"ひゃ" : "hya",
	"ひゅ" : "hyu",
	"ひょ" : "hyo",
	"みゃ" : "mya",
	"みゅ" : "myu",
	"みょ" : "myo",
	"りゃ" : "rya",
	"りゅ" : "ryu",
	"りょ" : "ryo",
	"ぎゃ" : "gya",
	"ぎゅ" : "gyu",
	"ぎょ" : "gyo",
	"じゃ" : "ja",
	"じゅ" : "ju",
	"じょ" : "jo",
	"ぴゃ" : "pya",
	"ぴゅ" : "pyu",
	"ぴょ" : "pyo",
	"びゃ" : "bya",
	"びゅ" : "byu",
	"びょ" : "byo",
	"いぇ" : "ie",
	"うぃ" : "ui",
	"うぇ" : "ue",
	"ゔぁ" : "bua",
	"ゔぃ" : "bui",
	"ゔ" : "bu",
	"ゔぇ" : "bue",
	"ゔぉ" : "buo",
	"くぁ" : "kua",
	"くぃ" : "kui",
	"くぇ" : "kue",
	"くぉ" : "kuo",
	"ぐぁ" : "gua",
	"ぐぃ" : "gui",
	"ぐぇ" : "gue",
	"ぐぉ" : "guo",
	"じぇ" : "jie",
	"ちぇ" : "chie",
	"つぁ" : "tsua",
	"つぃ" : "tsui",
	"つぇ" : "tsue",
	"つぉ" : "tsuo",
	"てぃ" : "tei",
	"でぃ" : "dei",
	"でゅ" : "deyu",
	"どぅ" : "dou",
	"ふぁ" : "fua",
	"ふぃ" : "fui",
	"ふぇ" : "fue",
	"ふぉ" : "fuo",
	"ふょ" : "fuyo",
	"あ" : "a",
	"い" : "i",
	"う" : "u",
	"え" : "e",
	"お" : "o",
	"か" : "ka",
	"き" : "ki",
	"く" : "ku",
	"け" : "ke",
	"こ" : "ko",
	"が" : "ga",
	"ぎ" : "gi",
	"ぐ" : "gu",
	"げ" : "ge",
	"ご" : "go",
	"さ" : "sa",
	"し" : "shi",
	"す" : "su",
	"せ" : "se",
	"そ" : "so",
	"ざ" : "za",
	"じ" : "ji",
	"ず" : "zu",
	"ぜ" : "ze",
	"ぞ" : "zo",
	"た" : "ta",
	"ち" : "chi",
	"つ" : "tsu",
	"て" : "te",
	"と" : "to",
	"だ" : "da",
	"ぢ" : "ji",
	"づ" : "du",
	"で" : "de",
	"ど" : "do",
	"な" : "na",
	"に" : "ni",
	"ぬ" : "nu",
	"ね" : "ne",
	"の" : "no",
	"は" : "ha",
	"ひ" : "hi",
	"ふ" : "fu",
	"へ" : "he",
	"ほ" : "ho",
	"ば" : "ba",
	"び" : "bi",
	"ぶ" : "bu",
	"べ" : "be",
	"ぼ" : "bo",
	"ぱ" : "pa",
	"ぴ" : "pa",
	"ぷ" : "pu",
	"ぺ" : "pe",
	"ぽ" : "po",
	"ま" : "ma",
	"み" : "mi",
	"む" : "mu",
	"め" : "me",
	"も" : "mo",
	"や" : "ya",
	"ゆ" : "yu",
	"よ" : "yo",
	"ら" : "ra",
	"り" : "ri",
	"る" : "ru",
	"れ" : "re",
	"ろ" : "ro",
	"わ" : "wa",
	"ゐ" : "i",
	"ゑ" : "e",
	"を" : "o",
	"ん" : "n",
}

func Kana2PassportRoman(str string) (roma string, hint string) {
	lru := []rune(str)
	l := len(lru)
	
	for pos := 0; pos < l ; pos++ {
		if pos < (l-2) {
			switch(string(lru[pos:pos + 1])) {
			case "っ":
				nextr := PassportMap[string(lru[pos + 1:pos + 2])]
				nr := []rune(nextr)
				if string(nr[0:1]) == "c" {
					roma += "t"
				} else {
					roma += string(nr[0:1])
				}
				continue
			case "ん":
				switch(string(lru[pos + 1:pos + 2])){
				case 	"ば", "び", "ぶ", "べ", "ぼ", 
					"ぱ", "ぴ", "ぷ", "ぺ", "ぽ", 
					"ま", "み", "む", "め", "も":
					roma += "m"
					continue
				}
				roma += "n"
				continue
			}
			rns := string(lru[pos + 1:pos + 2])
			switch(rns){
			case "ゃ","ゅ","ょ","ぁ","ぃ","ぅ","ぇ","ぉ":
				roma += PassportMap[string(lru[pos:pos+2])]
				if pos < (l-3) {
					switch(string(lru[pos+1:pos+3])){
					case "ゅう","ょう","ょお":
						hint += "o"
					}
				}
				pos += 1
				continue
			case "お","う":
				roma += PassportMap[string(lru[pos:pos+1])]
				hint += "o"
				continue
			}
		}
		if pos < (l-1) {
			roma += PassportMap[string(lru[pos:pos+1])]
			
		}
	}
	return
}

func main() {
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
		roma, hint := Kana2PassportRoman(line)
		fmt.Printf("%s\t%s\n",roma, hint)
	}
}
