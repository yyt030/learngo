package parser

import (
	"regexp"

	"learngo/crawler/engine"
)

const cityre = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityre)
	match := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range match {
		userName := string(m[2])
		result.Items = append(result.Items, "User: "+userName)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParseFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, userName)
			},
		})
	}
	return result
}
