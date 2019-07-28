package util

import (
"bufio"
"fmt"
"io"
"os"
)

func Read(file string,f func(string,int)) {

	fi, err := os.Open(file)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	lineN:= 1
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		f(string(a),lineN)
		lineN++
	}
}