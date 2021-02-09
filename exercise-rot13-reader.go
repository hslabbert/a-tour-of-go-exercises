package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	for {
		n, err := r.r.Read(b)
		for i := 0; i < n; i++ {
			if (int(b[i]) >= 65 && int(b[i]) <= 77) || (int(b[i]) >= 97 && int(b[i]) <= 109) {
				b[i] = b[i] + byte(13)
			} else if (int(b[i]) >= 78 && int(b[i]) <= 90) || (int(b[i]) >= 110 && int(b[i]) <= 122) {
				b[i] = b[i] - byte(13)
			}
		}
		return n, err
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
