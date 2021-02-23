package processor

import (
	"errors"

	"github.com/sebvautour/event-manager/internal/service"
	"github.com/segmentio/kafka-go"
)

type processorFunc func(msg kafka.Message) (retry bool, err error)

type Processor struct {
	s *service.Service
}

func New(svc *service.Service) (p *Processor) {
	return &Processor{
		s: svc,
	}

}

func (p *Processor) run(fc processorFunc, r *kafka.Reader) error {
	defer func() {
		if r := recover(); r != nil {
			p.s.Log.Error(r)
		}
	}()
	defer p.s.Log.Info("processor stopping")

	p.s.Log.Info("starting processor message loop")
	for {
		msg, err := r.FetchMessage(p.s.Context)
		p.s.Log.Info("processor fetch messsage " + string(msg.Key))
		if err != nil {
			p.s.Log.Error(err)
			return err
		}

		if err := p.retryLoop(msg, fc); err != nil {
			p.s.Log.Error("Err processing msg: " + err.Error())
		}

		if err := r.CommitMessages(p.s.Context, msg); err != nil {
			p.s.Log.Error("Err committing msg: " + err.Error())
		}
		p.s.Log.Info("processor messsage commit" + string(msg.Key))

		select {
		case <-p.s.Context.Done():
			return errors.New("Context is done")
		default:
		}
	}
}

func (p *Processor) retryLoop(msg kafka.Message, fc processorFunc) (err error) {
	for {

		retry, err := fc(msg)

		if err == nil {
			return nil
		}
		if retry == false {
			return err
		}
		p.s.Log.Error("Retriable error processing msg: " + err.Error())

	}
}
