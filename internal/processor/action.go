package processor

import (
	"github.com/sebvautour/event-manager/internal/vm"
	"gocloud.dev/pubsub"
)

func (p *Processor) ActionProcessor(sub *pubsub.Subscription) error {
	p.vm = vm.New(p.s)
	return p.run(p.ProcessActionMsgFunc, sub)
}

func (p *Processor) ProcessActionMsgFunc(msg *pubsub.Message) (err error) {
	// Do work based on the message, for example:

	p.vm.Run(string(msg.Body))

	// Messages must always be acknowledged with Ack.
	msg.Ack()
	return nil
}
