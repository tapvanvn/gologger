package entity

type LogEvent struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

type Log struct {
	Timestamp int64  `json:"Timestamp"`
	Agent     string `json:"Agent"`
	Events    []*LogEvent
}
