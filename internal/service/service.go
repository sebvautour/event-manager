package service

import (
	"context"

	"github.com/sirupsen/logrus"
)

// Service contains all the variables each service should have
type Service struct {
	Context context.Context
	Log     *logrus.Entry
}
