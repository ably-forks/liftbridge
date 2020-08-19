module github.com/ably-forks/liftbridge

go 1.14

replace github.com/liftbridge-io/liftbridge-api => github.com/ably-forks/liftbridge-api v1.9.0

replace github.com/liftbridge-io/liftbridge => github.com/ably-forks/liftbridge v1.9.0

replace github.com/liftbridge-io/go-liftbridge => github.com/ably-forks/go-liftbridge v1.9.0

require (
	github.com/Workiva/go-datastructures v1.0.52
	github.com/dustin/go-humanize v1.0.0
	github.com/golang/protobuf v1.4.2
	github.com/hako/durafmt v0.0.0-20200605151348-3a43fc422dd9
	github.com/hashicorp/raft v1.1.2
	github.com/liftbridge-io/go-liftbridge v1.0.1-0.20200707183953-f9c0b883e534
	github.com/liftbridge-io/liftbridge v0.0.0-00010101000000-000000000000
	github.com/liftbridge-io/liftbridge-api v1.1.0
	github.com/liftbridge-io/nats-on-a-log v0.0.0-20200303015016-68120bc11e03
	github.com/liftbridge-io/raft-boltdb v0.0.0-20200414234651-aaf6e08d8f73
	github.com/natefinch/atomic v0.0.0-20200526193002-18c0533a5b09
	github.com/nats-io/nats-server/v2 v2.1.4
	github.com/nats-io/nats.go v1.10.0
	github.com/nats-io/nuid v1.0.1
	github.com/nsip/gommap v0.0.0-20181229045655-f7881c3a959f
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/viper v1.7.0
	github.com/stretchr/testify v1.4.0
	github.com/urfave/cli v1.22.4
	google.golang.org/grpc v1.30.0
)
