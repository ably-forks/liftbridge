package addon

import (
	client "github.com/liftbridge-io/liftbridge-grpc/go"
	"github.com/liftbridge-io/liftbridge/server/proto"
	"google.golang.org/grpc"
)

type Addon interface {
	Initialize(interface{}) error
	Name() string
	RegisterGrpcServer(*grpc.Server) error
	LeadershipAcquired() error
	LeadershipLost() error
	ProcessMessage(stream string, subject string, msg proto.Message) bool
	MessageReceived(stream string, msg *client.Message)
	PartitionCreated(int32)
}
