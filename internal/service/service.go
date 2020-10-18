package service

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/sebvautour/event-manager/pkg/client"
)

// Service contains all the variables each service should have
type Service struct {
	Context context.Context
	API     *client.Client
	Log     *logrus.Entry
}
