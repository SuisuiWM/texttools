// Copyright (C) 2016 Makoto Imaizumi <Suisui@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Japanese Half width kana to kana converter
// Usage) han2zen filename
  

package main

import(
	"fmt"
	"bufio"
	"os"
	"golang.org/x/text/unicode/norm"
)

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
		fmt.Printf("%s",string(norm.NFKC.Bytes([]byte(line))))
	}
}

