package parser

import (
	"regexp"

	"learngo/crawler/engine"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte) engine.ParseResult {
	match := profileRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range match {
		userName := string(m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParseFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, userName)
			},
		})
	}

	submatchs := cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range submatchs {
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			ParseFunc: ParseCity,
		})

	}

	return result
}
