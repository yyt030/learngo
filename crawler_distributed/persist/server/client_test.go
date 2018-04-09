package main

import (
	"fmt"
	"testing"
	"time"

	"learngo/crawler/engine"
	"learngo/crawler/model"
	"learngo/crawler_distributed/config"
	"learngo/crawler_distributed/rpcsupport"
)

func TestItemSaver(t *testing.T) {
	host := fmt.Sprintf(":%d", config.ItemSaverPort)
	// start ItemSaverServer
	go serveRpc(host, "test1")
	time.Sleep(time.Second)
	// start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		t.Error("error")
	}
	// Call save
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
		}}

	result := ""
	err = client.Call(config.ItemSaverRpc, item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s", result, err)
	}
}
