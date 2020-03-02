package server

import (
	"fmt"
	"time"
)

// stream is a message stream consisting of one or more partitions. Each
// partition maps to a NATS subject and is the unit of replication.
type stream struct {
	name              string
	subject           string
	partitions        map[int32]*partition
	AutoCloseDuration time.Duration
	AutoCloseTimer    *time.Timer
}

// Close the stream by closing each of its partitions.
func (p *stream) Close() error {
	for _, partition := range p.partitions {
		if err := partition.Close(); err != nil {
			return err
		}
	}
	return nil
}

func (p *stream) Pause() error {
	for _, partition := range p.partitions {
		if err := partition.Pause(); err != nil {
			return err
		}
	}
	return nil
}

func (p *stream) SetupAutoclose() {
	p.AutoCloseTimer = time.NewTimer(p.AutoCloseDuration)
	go func() { // TODO: use proper goroutine management
		<-p.AutoCloseTimer.C
		p.Pause()
		fmt.Printf("Closing Stream")
	}()
}
