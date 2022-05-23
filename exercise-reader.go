package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (my MyReader) Read( str []byte) (int, error) {
	for i := range str {
		str[i] = 'A'
	}
	return len(str), nil
}

func main() {
	reader.Validate(MyReader{})
}
