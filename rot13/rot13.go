package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)

	for i := 0; i < n; i++ {
		switch {
		case b[i] >= 'A' && b[i] <= 'M':
			fallthrough
		case b[i] >= 'a' && b[i] <= 'm':
			b[i] += 13
		case b[i] > 'M' && b[i] <= 'Z':
			fallthrough
		case b[i] > 'm' && b[i] <= 'z':
			b[i] -= 13
		}
	}
	return n, err
}

var tests = []string{
	"Lbh penpxrq gur pbqr!",
	"WByyl TBbQ FubJ, JunG JUNg?!?!",
	"GUR DHVPX OEBJA SBK WHZCF BIRE GUR YNML QBT!!",
	"gur dhvpx oebja sbk whzcf bire gur ynml qbt!!",
}

func main() {
	for _, test := range tests {
		s := strings.NewReader(test)
		r := rot13Reader{s}
		io.Copy(os.Stdout, &r)
		fmt.Println()
	}
}
