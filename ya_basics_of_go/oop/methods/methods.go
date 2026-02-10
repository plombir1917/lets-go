package methods

import (
	"fmt"
	"time"
)

type Stopwatch struct {
	startTime time.Time
	splits    []time.Duration
}

func (s *Stopwatch) Start() {
	s.startTime = time.Now()
	s.splits = nil
}

func (s *Stopwatch) SaveSplit() {
	passedTime := time.Since(s.startTime)
	s.splits = append(s.splits, passedTime)
}

func (s Stopwatch) GetResults() (retResults []time.Duration) {
	// возвращаем копию слайса, чтобы извне не дать доступ к оригинальному
	for _, result := range s.splits {
		retResults = append(retResults, result)
	}

	return
}

func Execute() {
	sw := Stopwatch{}
	sw.Start()

	time.Sleep(1 * time.Second)
	sw.SaveSplit()

	time.Sleep(500 * time.Millisecond)
	sw.SaveSplit()

	time.Sleep(300 * time.Millisecond)
	sw.SaveSplit()

	fmt.Println(sw.GetResults())
}
