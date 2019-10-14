package server

import (
	"github.com/liftbridge-io/liftbridge/server/logger"
	"github.com/liftbridge-io/liftbridge/server/proto"
	"google.golang.org/grpc"
)

type Plugin interface {
	Initialize(logger.Logger) error
	Name() string
	RegisterGrpcServer(*grpc.Server) error
	LeadershipAcquired() error
	LeadershipLost() error
	ProcessMessage(proto.Message) bool
}
