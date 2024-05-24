package houdinimanager

import (
	"errors"
	"sync"
	"time"
)

type HoudiniManager struct {
	routineCounter int
	maxRoutines    int
	syncGroup      sync.WaitGroup
	needToWait     bool
}

func NewHoudiniManager(setMaxRoutines, needToWait bool, maxRoutines int) (*HoudiniManager, error) {

	if setMaxRoutines && maxRoutines <= 0 {
		return nil, errors.New("if you need to set it, max houdinis must be greater than 0")
	} else if !setMaxRoutines && maxRoutines > 0 {
		return nil, errors.New("if you don't need to set it, max houdinis must be 0")
	}

	manager := &HoudiniManager{
		routineCounter: 0,
		needToWait:     needToWait,
	}

	if setMaxRoutines {
		manager.maxRoutines = maxRoutines
	}

	manager.syncGroup = sync.WaitGroup{}

	return manager, nil
}

func (r *HoudiniManager) Execute(f func()) {
	r.waitAvailableRoutine()
	r.addCount()
	r.addWait()
	go func(r *HoudiniManager) {
		defer r.reduceCount()
		defer r.doneWait()
		f()
	}(r)
}

func (r *HoudiniManager) addWait() {
	if r.needToWait {
		r.syncGroup.Add(1)
	}
}

func (r *HoudiniManager) doneWait() {
	if r.needToWait {
		r.syncGroup.Done()
	}
}

func (r *HoudiniManager) Wait() {
	if r.needToWait {
		r.syncGroup.Wait()
	}
}

func (r *HoudiniManager) PrintHoudiniCounter() {
	println("Houdini counter: ", r.routineCounter)
}

func (r *HoudiniManager) waitAvailableRoutine() {
	if r.maxRoutines > 0 {
		for r.routineCounter == r.maxRoutines {
			time.Sleep(500 * time.Millisecond)
		}
	}
}
func (r *HoudiniManager) addCount() {
	if r.maxRoutines > 0 {
		r.routineCounter++
	}
}

func (r *HoudiniManager) reduceCount() {
	if r.maxRoutines > 0 {
		r.routineCounter--
	}
}
