package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Println(
		// true
		strings.Contains("TEST", "ES"),

		// 2
		strings.Count("TEST", "T"),

		// true
		strings.HasPrefix("test", "te"),

		// true
		strings.HasSuffix("test", "st"),

		// 1
		strings.Index("test", "e"),

		// 1
		strings.Index("teste", "e"),

		// 1
		strings.IndexAny("teste", "e"),

		// a-b
		strings.Join([]string{"a", "b"}, "-"),

		// aaaaa
		strings.Repeat("a", 5),

		// bbba
		strings.Replace("aaaa", "a", "b", 3),

		// []string{a,b,c,d,e,f}
		strings.Split("a b c d e f", " "),

		// test
		strings.ToLower("TEST"),

		// TEST
		strings.ToUpper("test"),

		// [116 101 115 116]
		[]byte("test"),

		// test
		string([]byte{'t', 'e', 's', 't'}),
	)

	var buf bytes.Buffer
	buf.Write([]byte("test"))
	fmt.Println(buf)

	file, err := os.Open("test.txt")
	if err != nil {
		// handle the error
		return
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}

	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)

	// shorter way of reading files
	txt, error := ioutil.ReadFile("test.txt")
	if error != nil {
		return
	}
	fmt.Println(string(txt))

	// create new file
	file3, err3 := os.Create("test2.txt")
	if err3 != nil {
		return
	}
	defer file3.Close()
	file3.WriteString("HELP!")

}
