module cache

go 1.13

replace github.com/liftbridge-io/liftbridge/server/addon => ../

replace github.com/liftbridge-io/liftbridge/server/addon/cache => ../cache

replace github.com/liftbridge-io/liftbridge/server/addon/cache/api => ./api

require (
	github.com/golang/mock v1.3.1
	github.com/liftbridge-io/liftbridge v0.0.0-20191009142300-f92ab643e701
	github.com/liftbridge-io/liftbridge-grpc v0.0.0-20190829220806-66e3ee4b7943
	github.com/liftbridge-io/liftbridge/server/addon/cache/api v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.24.0
)
