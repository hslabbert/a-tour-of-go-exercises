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
		for i := 0; i < 20; i++ {
			if 65 < int(b[i]) && int(b[i]) < 90 {
				b[i] = b[i] + byte(13)
				if int(b[i]) > 90 {
					b[i] = b[i] - byte(26)
				}
			} else if 97 < int(b[i]) && int(b[i]) < 122 {
				b[i] = b[i] + byte(13)
				if int(b[i]) > 122 {
					b[i] = b[i] - byte(26)
				}
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
