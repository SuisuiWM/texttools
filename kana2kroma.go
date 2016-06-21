// Copyright (C) 2016 Makoto Imaizumi <Suisui@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Japanese Hiragana to Kunrei type romanization convert script
// Usage) kana2kroma filename

// Unsupported:
// 長音 as  ー
// separator with "'" whien n[aiueoy] case
// 2nd convert table

package main

import(
	"fmt"
	"bufio"
	"os"
)

var KunreiMap = map[string]string{
	"きゃ" : "kya",
	"きゅ" : "kyu",
	"きょ" : "kyo",
	"しゃ" : "sya",
	"しゅ" : "syu",
	"しょ" : "syo",
	"しぇ" : "sye",
	"ちゃ" : "tya",
	"ちゅ" : "tyu",
	"ちょ" : "tyo",
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
	"じゃ" : "zya",
	"じゅ" : "zyu",
	"じぇ" : "zye",
	"じょ" : "zyo",
	"ぢゃ" : "zya",
	"ぢゅ" : "zyu",
	"ぢょ" : "zyo",
	"ぴゃ" : "pya",
	"ぴゅ" : "pyu",
	"ぴょ" : "pyo",
	"びゃ" : "bya",
	"びゅ" : "byu",
	"びょ" : "byo",
	"ゔゃ" : "vya",
	"ゔゅ" : "vyu",
	"ゔょ" : "vyo",
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
	"し" : "si",
	"す" : "su",
	"せ" : "se",
	"そ" : "so",
	"ざ" : "za",
	"じ" : "zi",
	"ず" : "zu",
	"ぜ" : "ze",
	"ぞ" : "zo",
	"た" : "ta",
	"ち" : "ti",
	"つ" : "tu",
	"て" : "te",
	"と" : "to",
	"だ" : "da",
	"ぢ" : "di",
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
	"ふ" : "hu",
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

func kunreiConvert(rns []rune, start int, end int) string {
	return KunreiMap[string(rns[start:end])]
}

func Kana2Kunrei(str string) (roma string) {
	lru := []rune(str)
	l := len(lru)
	for pos := 0; pos < l ; pos++ {
		if pos < (l-2) {
			switch(string(lru[pos:pos + 1])) {
			case "っ":
			//	nextr := KunreiMap[string(lru[pos + 1:pos + 2])]
			//	nr := []rune(nextr)
				nr := []rune(kunreiConvert(lru, pos+1, pos+2))
				roma += string(nr[0:1])
				continue
			}
			rns := string(lru[pos + 1:pos + 2])
			switch(rns){
			case "ゃ","ゅ","ょ":
			//	roma += KunreiMap[string(lru[pos:pos+2])]
				roma += kunreiConvert(lru, pos, pos+2)
				pos += 1
				continue
			}
		}
		if pos < (l-1) {
		//	roma += KunreiMap[string(lru[pos:pos+1])]
			roma += kunreiConvert(lru,pos,pos+1)
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
		roma := Kana2Kunrei(line)
		fmt.Printf("%s\n",roma)
	}
}
