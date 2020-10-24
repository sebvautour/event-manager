package processor

import (
	"errors"

	"github.com/sebvautour/event-manager/internal/service"

	"gocloud.dev/pubsub"
)

type processorFunc func(msg *pubsub.Message) (err error)

type Processor struct {
	s *service.Service
}

func New(svc *service.Service) (p *Processor) {
	return &Processor{
		s: svc,
	}

}

func (p *Processor) run(fc processorFunc, sub *pubsub.Subscription) error {
	defer func() {
		if r := recover(); r != nil {
			p.s.Log.Error(r)
		}
	}()
	defer p.s.Log.Info("processor stopping")
	defer sub.Shutdown(p.s.Context)

	p.s.Log.Info("starting processor message loop")
	for {
		msg, err := sub.Receive(p.s.Context)
		p.s.Log.Info("received msg")
		if err != nil {
			p.s.Log.Error(err)
			return err
		}

		err = fc(msg)
		if err != nil {
			p.s.Log.Error("Err processing msg: " + err.Error())
		}

		select {
		case <-p.s.Context.Done():
			return errors.New("Context is done")
		default:
		}
	}
}
