package gologger_test

import (
	"testing"

	logger "github.com/tapvanvn/gologger"
)

func TestLogger(t *testing.T) {
	l := logger.NewLogger("test")
	l.AddBehavior(&logger.PrintBehavior{})
	l.Log(
		logger.Pair("key1", "value1"),
		logger.Pair("key2", "value2"),
	)
}

func TestHubLogger(t *testing.T) {

	logger.AddGlobalBehavior(&logger.PrintBehavior{})

	l := logger.GetGlobalLogger("test")

	l.Log(
		logger.Pair("key1", "value1"),
		logger.Pair("key2", "value2"),
	)

	l2 := logger.GetGlobalLogger("test2")

	l2.Log(
		logger.Pair("key1", "value1"),
		logger.Pair("key2", "value2"),
	)
}
