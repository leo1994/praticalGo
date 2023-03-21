package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	sig, err := sha1Sum("file.gz")
	if err != nil {
		log.Fatalf("%#v", err)
	}

	fmt.Println(sig)
}
func sha1Sum(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var r io.Reader = file

	if strings.HasSuffix(fileName, ".gz") {
		gz, err := gzip.NewReader(file)
		if err != nil {
			return "", err
		}
		defer gz.Close()
		r = gz
	}

	// io.Copy(os.Stdout, r)

	w := sha1.New()

	if _, err := io.Copy(w, r); err != nil {
		return "", err
	}
	sig := w.Sum(nil)

	return fmt.Sprintf("%x", sig), nil
}
