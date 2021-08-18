package gologger

import (
	"fmt"

	"github.com/tapvanvn/gologger/entity"
)

type IBehavior interface {
	Process(log *entity.Log)
}

type PrintBehavior struct {
	//TODO: support log format
}

func (print *PrintBehavior) Process(log *entity.Log) {
	fmt.Printf("%s %d ", log.Agent, log.Timestamp)
	for _, event := range log.Events {
		if event.Key == "" {
			fmt.Printf("%s ", event.Value)
		} else {
			fmt.Printf("%s=%s ", event.Key, event.Value)
		}
	}
	fmt.Println()
}
