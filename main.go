package main

import (
	"github.com/bean-du/crawler/engine"
	"github.com/bean-du/crawler/persist"
	"github.com/bean-du/crawler/scheduler"
	"github.com/bean-du/crawler/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedSchedule{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}

	e.Run(engine.Request{
		Url:        "http://city.zhenai.com",
		ParserFunc: parser.ParseCityList,
	})

}
