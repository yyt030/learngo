package main

import (
	"fmt"
	"testing"
	"time"

	"learngo/crawler_distributed/config"
	"learngo/crawler_distributed/rpcsupport"
	"learngo/crawler_distributed/worker"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		t.Error(err)
	}
	req := worker.Request{
		Url: "http://album.zhenai.com/u/106684546",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "月娆",
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
