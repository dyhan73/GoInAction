package main

import (
	"fmt"
	"sync"
	"time"
)

func scoreboardManager(in <-chan func(map[string]int), done <-chan struct{}) {
	scoreboard := map[string]int{}
	for {
		select {
		case <-done:
			return
		case f := <-in:
			f(scoreboard)
		}
	}
}

type ChannelScoreboardManager chan func(map[string]int)

func NewChannelScoreboardManager() (ChannelScoreboardManager, func()) {
	ch := make(ChannelScoreboardManager)
	done := make(chan struct{})
	go scoreboardManager(ch, done)
	return ch, func() {
		close(done)
	}
}

func (csm ChannelScoreboardManager) Update(name string, val int) {
	csm <- func(m map[string]int) {
		m[name] = val
	}
}

func (csm ChannelScoreboardManager) Read(name string) (int, bool) {
	var out int
	var ok bool
	done := make(chan struct{})
	csm <- func(m map[string]int) {
		time.Sleep(10 * time.Millisecond) // 스코어보드에서 처리될 때 지연효과 주려고
		out, ok = m[name]
		close(done)
	}
	<-done // csm 으로 전달한 함수가 종료된 것 확인한 후 리턴하려고
	return out, ok
}

func scoreboardChanVersion() {
	csm, cf := NewChannelScoreboardManager()
	fmt.Println(csm.Read("dyhan"))
	csm.Update("dyhan", 40)
	fmt.Println(csm.Read("dyhan"))
	cf()
}

type MutexScoreboardManager struct {
	l          sync.RWMutex
	scoreboard map[string]int
}

func NewMutexScoreboardManager() *MutexScoreboardManager {
	return &MutexScoreboardManager{
		scoreboard: map[string]int{},
	}
}

func (msm *MutexScoreboardManager) Update(name string, val int) {
	msm.l.Lock()
	defer msm.l.Unlock()
	msm.scoreboard[name] = val
}

func (msm *MutexScoreboardManager) Read(name string) (int, bool) {
	msm.l.RLock()
	defer msm.l.RUnlock()
	val, ok := msm.scoreboard[name]
	return val, ok
}
func main() {
	scoreboardChanVersion()
	msm := NewMutexScoreboardManager()
	fmt.Println(msm.Read("han"))
	msm.Update("han", 20)
	fmt.Println(msm.Read("han"))
}
