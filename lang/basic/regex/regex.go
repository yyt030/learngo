package main

import (
	"fmt"
	"regexp"
)

const text = `My email is test@gmail.com
my email is ok@baidu.com
email3 is kkk@qq.com
`

func main() {
	re := regexp.MustCompile("([a-zA-Z0-9]+)@[a-zA-Z]+.com")
	s := re.FindString(text)
	fmt.Println(s)
}
