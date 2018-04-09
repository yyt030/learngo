package main

import (
	"fmt"

	"learngo/crawler_distributed/config"
	"learngo/crawler_distributed/rpcsupport"
	"learngo/crawler_distributed/worker"
)

func main() {
	rpcsupport.ServeRpc(fmt.Sprintf(":%d", config.WorkPort0),
		worker.CrawlService{})
}
