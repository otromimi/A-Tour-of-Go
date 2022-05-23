package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (inst *rot13Reader) Read(p []byte) (n int, err error) {

	rot13 := make([]byte, len(p))
	n, err = inst.r.Read(rot13)
	
	for i := range p{
	
		if rot13[i] >= 'a' && rot13[i] <= 'z' {
		
			p[i] = 'a' + (rot13[i] - 'a' + 13) % 26 
			
		}else if rot13[i] >= 'A' && rot13[i] <= 'Z' { 
		
			p[i] = 'A' + (rot13[i] - 'A' + 13) % 26 
			
		}else {
		
			p[i] = rot13[i]
			
		}
	}
	
	return
}
		

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
