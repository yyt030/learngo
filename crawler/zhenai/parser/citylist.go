package parser

import (
	"regexp"

	"learngo/crawler/engine"
	"learngo/crawler_distributed/config"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		url := string(m[1])
		result.Requests = append(result.Requests, engine.Request{
			Url:    url,
			Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
		})
	}
	return result
}
