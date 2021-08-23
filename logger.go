package gologger

import (
	"sync"
	"time"

	"github.com/tapvanvn/gologger/entity"
)

func NewLogger(agent string) *Logger {

	return &Logger{
		Agent:     agent,
		behaviors: map[string]IBehavior{},
	}
}

type Logger struct {
	Agent     string
	behaviors map[string]IBehavior
	mux       sync.Mutex
}

func (logger *Logger) AddBehavior(behave IBehavior) {
	logger.mux.Lock()
	defer logger.mux.Unlock()
	logger.behaviors[behave.GetName()] = behave
}
func (logger *Logger) RemoveBehavior(behave IBehavior) {
	logger.mux.Lock()
	defer logger.mux.Unlock()
	delete(logger.behaviors, behave.GetName())
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
