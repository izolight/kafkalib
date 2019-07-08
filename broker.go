package kafkalib

type Broker struct {
	ID   string `json:"id"`
	Addr string `json:"addr"`
	Rack string `json:"rack"`
}

