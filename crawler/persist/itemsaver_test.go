package persist

import (
	"context"
	"encoding/json"
	"testing"

	"learngo/crawler/engine"
	"learngo/crawler/model"

	"gopkg.in/olivere/elastic.v5"
)

func TestItemSaver(t *testing.T) {
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

	// TODO: try to start up elasticsearch search here using
	// docker go client
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		t.Errorf("connect error: %v", err)
	}

	//Save expected item
	const index = "dating_test"
	err = saver(client, index, expected)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	// fetch saved item
	resp, err := client.Get().Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())
	if err != nil {
		t.Errorf("get failed: %v", err)
	}

	var actual engine.Item
	err = json.Unmarshal([]byte(*resp.Source), &actual)
	if err != nil {
		t.Errorf("json unmarshal error: %v", err)
	}

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	// Verify result
	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}
