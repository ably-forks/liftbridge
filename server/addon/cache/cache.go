package cache

import (
	context "context"
	"fmt"
	"sync"

	client "github.com/liftbridge-io/liftbridge-grpc/go"
	"github.com/liftbridge-io/liftbridge/server"
	"github.com/liftbridge-io/liftbridge/server/addon/cache/api"
	"github.com/liftbridge-io/liftbridge/server/proto"
	"google.golang.org/grpc"
)

type messageMap map[string][]byte

type safeMessageMap struct {
	sync.Mutex
	m messageMap
}

type CachePlugin struct {
	s        *server.Server
	isLeader bool
	m        safeMessageMap
	config   *Config
}

func New() *CachePlugin {
	return &CachePlugin{
		m: safeMessageMap{
			m: make(map[string][]byte),
		},
	}
}

func (p *CachePlugin) Initialize(s interface{}) error {
	p.s = s.(*server.Server)

	var err error
	p.config, err = NewConfig("cache-addon.conf")
	if err != nil {
		return err
	}

	fmt.Printf("Loaded addon with value %v\n", p.config.Value)

	return nil
}

func (p CachePlugin) Name() string {
	return "Memory Cache"
}

func (p *CachePlugin) RegisterGrpcServer(srv *grpc.Server) error {
	api.RegisterCacheAPIServer(srv, p)
	return nil
}

func (p *CachePlugin) LeadershipAcquired() error {
	p.isLeader = true

	return nil
}

func (p *CachePlugin) LeadershipLost() error {
	p.isLeader = false

	return nil
}

func (p *CachePlugin) ProcessMessage(stream string, subject string, m proto.Message) bool {
	//fmt.Printf("message %v\n", m)

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

func (p *CachePlugin) MessageReceived(stream string, m *client.Message) {
	p.m.Lock()
	defer p.m.Unlock()

	if len(m.Key) > 0 {
		key := stream + m.Subject + string(m.Key)
		p.m.m[key] = m.Value
	}
}

func (p *CachePlugin) PartitionCreated(part int32) {
}

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
