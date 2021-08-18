package gologger

import (
	"time"

	"github.com/tapvanvn/gologger/entity"
)

func NewLogger(agent string) *Logger {
	return &Logger{
		Agent:     agent,
		behaviors: make([]IBehavior, 0),
	}
}

type Logger struct {
	Agent     string
	behaviors []IBehavior
}

func (logger *Logger) AddBehavior(behave IBehavior) {
	logger.behaviors = append(logger.behaviors, behave)
}
func (logger *Logger) Log(pairs ...*entity.LogEvent) {
	log := &entity.Log{
		Agent:     logger.Agent,
		Timestamp: time.Now().Unix(),
		Events:    make([]*entity.LogEvent, 0),
	}
	log.Events = append(log.Events, pairs...)
	for _, behave := range logger.behaviors {
		behave.Process(log)
	}
}

func Pair(key string, value string) *entity.LogEvent {
	return &entity.LogEvent{
		Key:   key,
		Value: value,
	}
}
