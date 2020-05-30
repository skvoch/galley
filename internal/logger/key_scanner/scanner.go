package key_scanner

import (
	"github.com/robotn/gohook"
	"sync/atomic"
	"time"
)

func New(duration time.Duration) *Scanner {
	scanner := &Scanner{
		countChannel: make(chan int64, 0),
	}
	go scanner.run()
	go scanner.timer(duration)

	return scanner
}

type Scanner struct {
	countChannel chan int64
	count        int64
	Duration     time.Duration
}

func (s *Scanner) GetCountChannel() chan int64 {
	return s.countChannel
}

func (s *Scanner) timer(duration time.Duration) {
	timer := time.NewTicker(duration)

	for {
		<-timer.C
		s.countChannel <- s.count

		atomic.StoreInt64(&s.count, 0)
	}
}

func (s *Scanner) run() {

	EvChan := hook.Start()
	defer hook.End()

	for ev := range EvChan {
		if ev.Kind == hook.KeyDown {
			atomic.AddInt64(&s.count, 1)
		}
	}

	e := hook.Start()
	<-hook.Process(e)
}
