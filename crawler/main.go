package main

import (
	"bufio"
	"fmt"
	"io"
	"regexp"

	"learngo/crawler/engine"
	"learngo/crawler/zhenai/parser"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList})
}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func printCityList(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		fmt.Printf("City:%s, URL:%s\n", m[2], m[1])
	}
	fmt.Printf("Matches found: %d\n", len(matches))
}
