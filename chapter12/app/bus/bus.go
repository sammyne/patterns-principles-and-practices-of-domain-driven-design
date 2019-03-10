package bus

import (
	"encoding/json"

	nats "github.com/nats-io/go-nats"
)

// PublishTo encode a message as JSON and then send the resultant bytes into the bus
func PublishTo(bus *nats.Conn, subject string, msg interface{}) error {
	data, err := json.Marshal(&msg)
	if nil != err {
		return err
	}

	return bus.Publish(subject, data)
}

/*
func New() (*nats.Conn, func(string, interface{}) error, error) {
	broker, err = nats.Connect(nats.DefaultURL)
	if nil != err {
		return nil, nil, err
	}

	return broker, func(sub)
}
*/
