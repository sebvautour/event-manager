package processor

import (
	"github.com/segmentio/kafka-go"
)

func (p *Processor) ActionProcessor(reader *kafka.Reader) error {
	return p.run(p.ProcessActionMsgFunc, reader)
}

func (p *Processor) ProcessActionMsgFunc(msg kafka.Message) (retry bool, err error) {
	// Do work based on the message:

	return false, nil
}
