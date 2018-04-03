package persist

import (
	"context"
	"encoding/json"
	"testing"

	"learngo/crawler/model"

	"gopkg.in/olivere/elastic.v5"
)

func TestSaver(t *testing.T) {
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
	id, err := saver(expected)

	// TODO: try to start up elasticsearch search here using
	// docker go client
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		t.Errorf("connect error: %v", err)
	}
	resp, err := client.Get().Index("dating_profile").Type("zhenai").Id(id).Do(context.Background())
	if err != nil {
		t.Errorf("get failed: %v", err)
	}
	var actual model.Profile
	err = json.Unmarshal([]byte(*resp.Source), &actual)
	if err != nil {
		t.Errorf("json unmarshal error: %v", err)
	}
	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}
