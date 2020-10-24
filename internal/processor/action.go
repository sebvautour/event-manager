package processor

import (
	"gocloud.dev/pubsub"
)

func (p *Processor) ActionProcessor(sub *pubsub.Subscription) error {
	return p.run(p.ProcessActionMsgFunc, sub)
}

func (p *Processor) ProcessActionMsgFunc(msg *pubsub.Message) (err error) {
	// Do work based on the message:

	// Messages must always be acknowledged with Ack.
	msg.Ack()
	return nil
}
