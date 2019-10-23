package plugin

import (
	client "github.com/liftbridge-io/liftbridge-grpc/go"
	"github.com/liftbridge-io/liftbridge/server/proto"
	"google.golang.org/grpc"
)

// Plugin represents a Liftbridge plugin
type Plugin interface {
	// Initialize is called when the plugin should perform any setup procedures
	// This function takes a pointer that can be type asserted into a
	// server.Server.
	// This is done to prevent cyclic import issues.
	Initialize(interface{}) error
	// Name returns this plugin's name.
	Name() string
	// RegisterGrpcServer should be used by a plugin to register any Grpc
	// service to the Liftbridge server.
	RegisterGrpcServer(*grpc.Server) error
	// LeadershipAcquired is called every time the Liftbridge server gains
	// leadership.
	LeadershipAcquired() error
	// LeadershipLost is called every time the Liftbridge server looses
	// leadership.
	LeadershipLost() error
	// ProcessMessage is called every time a message arrives.
	// This function should return true is the message should be processed
	// or false otherwise
	ProcessMessage(stream string, subject string, msg proto.Message) bool
	// MessageReceived is called every time a message has been received
	// or restored from the log. Note that this function is currently
	// called multiple times per message.
	MessageReceived(stream string, msg *client.Message)
}
