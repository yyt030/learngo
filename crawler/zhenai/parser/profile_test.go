package parser

import (
	"io/ioutil"
	"testing"

	"learngo/crawler/model"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "月娆")
	if len(result.Items) != 1 {
		t.Errorf("items should contain 1 element; but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)
	expected := model.Profile{
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
	}
	if profile != expected {
		t.Errorf("expected %v; but was %v", expected, profile)
	}

}
