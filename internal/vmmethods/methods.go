package vmmethods

import (
	"github.com/sebvautour/event-manager/internal/service"
	"github.com/sirupsen/logrus"
)

type Methods struct {
	Message interface{}
	s       *service.Service
	log     *logrus.Entry
	drop    bool
}

func New(svc *service.Service, log *logrus.Entry, Message interface{}) *Methods {
	return &Methods{
		Message: Message,
		s:       svc,
		log:     log,
	}
}

func (m *Methods) LogInfo(msg interface{}) {
	m.log.Info(msg)
}

func (m *Methods) LogWarn(msg interface{}) {
	m.log.Warn(msg)
}
func (m *Methods) LogErr(msg interface{}) {
	m.log.Error(msg)
}

func (m *Methods) Drop() {
	m.drop = true
}

func (m *Methods) IsDropped() bool {
	return m.drop
}
