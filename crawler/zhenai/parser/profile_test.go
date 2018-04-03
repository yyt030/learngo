package parser

import (
	"io/ioutil"
	"testing"

	"learngo/crawler/engine"
	"learngo/crawler/model"
)

func TestParseProfile(t *testing.T) {

	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "http://album.zhenai.com/u/106684546", "月娆")
	if len(result.Items) != 1 {
		t.Errorf("items should contain 1 element; but was %v", result.Items)
	}

	actual := result.Items[0]
	expected := engine.Item{
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
			Cat:        "未购车",
		}}
	if actual != expected {
		t.Errorf("expected %v; but was %v", expected, actual)
	}

}
