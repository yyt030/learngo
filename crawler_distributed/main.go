package main

import (
	"flag"
	"log"
	"net/rpc"
	"strings"

	"learngo/crawler/engine"
	"learngo/crawler/scheduler"
	"learngo/crawler/zhenai/parser"
	"learngo/crawler_distributed/config"
	itemsaver "learngo/crawler_distributed/persist/client"
	"learngo/crawler_distributed/rpcsupport"
	worker "learngo/crawler_distributed/worker/client"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")
	workHosts     = flag.String("worker_hosts", "", "worker hosts(comma separated")
)

func main() {
	flag.Parse()
	itemChan, err := itemsaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workHosts, ","))
	processor := worker.CreateProcessor(pool)
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:      &scheduler.QueuedScheduler{},
		WorkCount:      100,
		ItemChan:       itemChan,
		RequestProcess: processor,
	}

	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
		} else {
			log.Printf("error connection to %s: %v", h, err)
		}
	}
	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
