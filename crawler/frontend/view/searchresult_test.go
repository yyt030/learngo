package view

import (
	"os"
	"testing"

	"learngo/crawler/engine"
	m "learngo/crawler/frontend/model"
	"learngo/crawler/model"
)

func TestSearchResultView_Render(t *testing.T) {
	//temp := template.Must(template.ParseFiles("searchresult.html"))

	temp := CreateSearchResultView("searchresult.html")

	data := m.SearchResult{}
	data.Hits = 123
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/106684546",
		Type: "zhenai",
		Id:   "106684546",
		Payload: model.Profile{
			Name:       "月娆",
			Gender:     "女",
			Age:        28,
			Height:     165,
			Weight:     57,
			Income:     "3001-5000元",
			Marriage:   "未婚",
			Education:  "中专",
			Occupation: "健身教练",
			Hokou:      "四川遂宁",
			Xinzuo:     "双鱼座",
			House:      "单位宿舍",
			Car:        "未购车",
		},
	}

	for i := 0; i < 10; i++ {
		data.Items = append(data.Items, item)
	}

	err := temp.Render(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
