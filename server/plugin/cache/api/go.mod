module api

go 1.13

replace github.com/liftbridge-io/liftbridge/server/plugin/cache/api => ./

require (
	github.com/golang/protobuf v1.3.2
	google.golang.org/grpc v1.24.0
)
