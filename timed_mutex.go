package trylock

import (
	"sync"
	"time"
)

type TryLocker interface {
	sync.Locker
	TryLock() (locked bool)
	TryLockTimeout(t time.Duration) (locked bool)
}

type tryMutexImpl struct {
	c chan uint8
}

func (this *tryMutexImpl) Lock() {
	this.c <- 0
}

func (this *tryMutexImpl) Unlock() {
	<-this.c
}

func (this *tryMutexImpl) TryLock() (locked bool) {
	select {
	case this.c <- 0:
		locked = true
	default:
		locked = false
	}
	return
}

func (this *tryMutexImpl) TryLockTimeout(t time.Duration) (locked bool) {
	timer := time.NewTimer(t)
	defer timer.Stop()
	select {
	case <-timer.C:
		locked = false
	case this.c <- 0:
		locked = true
	}
	return
}

func New() TryLocker {
	return &tryMutexImpl{
		c: make(chan uint8, 1),
	}
}
