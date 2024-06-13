package houdinimanager

import (
	"errors"
	"sync"
)

type HoudiniManager struct {
	channel        chan int
	waitGroup      sync.WaitGroup
	needToWait     bool
	setMaxRoutines bool
}

func NewHoudiniManager(setMaxRoutines, needToWait bool, maxRoutines int) (*HoudiniManager, error) {

	if setMaxRoutines && maxRoutines <= 0 {
		return nil, errors.New("if you need to set it, max houdinis must be greater than 0")
	} else if !setMaxRoutines && maxRoutines > 0 {
		return nil, errors.New("if you don't need to set it, max houdinis must be 0")
	}

	manager := &HoudiniManager{
		needToWait: needToWait,
	}

	if setMaxRoutines {
		manager.channel = make(chan int, maxRoutines)
		manager.setMaxRoutines = true
	} else {
		manager.channel = make(chan int)
		manager.setMaxRoutines = false
	}

	if needToWait {
		manager.waitGroup = sync.WaitGroup{}
	}

	return manager, nil
}

func (r *HoudiniManager) Execute(f func()) {
	if r.setMaxRoutines {
		r.channel <- 1
	}
	if r.needToWait {
		r.waitGroup.Add(1)
	}
	go func(r *HoudiniManager) {
		f()
		if r.setMaxRoutines {
			<-r.channel
		}
		if r.needToWait {
			r.waitGroup.Done()
		}
	}(r)
}

func (r *HoudiniManager) Wait() {
	if r.needToWait {
		r.waitGroup.Wait()
	}
}

func (r *HoudiniManager) Close() {
	close(r.channel)
}
