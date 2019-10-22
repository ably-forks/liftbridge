package cache

import (
	context "context"
	"fmt"
	"sync"

	client "github.com/liftbridge-io/liftbridge-grpc/go"
	"github.com/liftbridge-io/liftbridge/server"
	"github.com/liftbridge-io/liftbridge/server/plugin/cache/api"
	"github.com/liftbridge-io/liftbridge/server/proto"
	"google.golang.org/grpc"
)

// messageMap represents a collection of message data indexed by a string
type messageMap map[string][]byte

// safeMessageMap represents a messageMap that can be used from multiple thread
// Lock and Unlock have to be called before each map access
type safeMessageMap struct {
	sync.Mutex
	m messageMap
}

// CachePlugin implements a plugin that allows caching messages using their key
type CachePlugin struct {
	s        *server.Server
	isLeader bool
	m        safeMessageMap
	config   *Config
}

// New returns a new cache plugin
func New() *CachePlugin {
	return &CachePlugin{
		m: safeMessageMap{
			m: make(map[string][]byte),
		},
	}
}

// Initialize load this plugin's configuration file
func (p *CachePlugin) Initialize(s interface{}) error {
	p.s = s.(*server.Server)

	var err error
	p.config, err = NewConfig("cache-plugin.conf")
	if err != nil {
		return err
	}

	fmt.Printf("Loaded plugin with value %v\n", p.config.Value)

	return nil
}

// Name returns this plugin's name
func (p *CachePlugin) Name() string {
	return "Memory Cache"
}

// RegisterGrpcServer registers this plugin's gRPC API server
func (p *CachePlugin) RegisterGrpcServer(srv *grpc.Server) error {
	api.RegisterCacheAPIServer(srv, p)
	return nil
}

// LeadershipAcquired
func (p *CachePlugin) LeadershipAcquired() error {
	p.isLeader = true

	return nil
}

// LeadershipLost
func (p *CachePlugin) LeadershipLost() error {
	p.isLeader = false

	return nil
}

// ProcessMessage returns false if the provided message is cached, true otherwise
func (p *CachePlugin) ProcessMessage(stream string, subject string, m proto.Message) bool {
	p.m.Lock()
	defer p.m.Unlock()

	if len(m.Key) > 0 {
		key := stream + subject + string(m.Key)
		_, ok := p.m.m[key]
		// We already have this message, so ignore it
		if ok {
			return false
		}

		// Put the message in our cache
		p.m.m[key] = m.Value
	}

	return true
}

// MessageReceived adds a message in the cache
func (p *CachePlugin) MessageReceived(stream string, m *client.Message) {
	p.m.Lock()
	defer p.m.Unlock()

	if len(m.Key) > 0 {
		key := stream + m.Subject + string(m.Key)
		p.m.m[key] = m.Value
	}
}

// Get is called when a client accesses this plugin's API to query it for a
// message
func (p *CachePlugin) Get(ctx context.Context, r *api.GetRequest) (*api.GetResponse, error) {
	if !p.isLeader {
		//TODO: redirect to leader
		return nil, nil
	}

	p.m.Lock()
	defer p.m.Unlock()

	value, ok := p.m.m[r.Stream+r.Subject+r.Key]
	if !ok {
		return &api.GetResponse{Value: nil}, nil
	}

	return &api.GetResponse{Value: value}, nil
}
