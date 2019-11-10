package trylock

import (
	"sync"
	"testing"
	"time"
)

func Test_tryMutexImpl_Lock(t *testing.T) {
	locker := New()
	locker.Lock()
	defer locker.Unlock()
}

func Test_tryMutexImpl_TryLock(t *testing.T) {
	locker := New()
	if locker.TryLock() {
		locker.Unlock()
	}
}

func Test_tryMutexImpl_TryLockTimeout(t *testing.T) {
	locker := New()
	locker.Lock()
	wg := new(sync.WaitGroup)
	wg.Add(2)
	defer wg.Wait()
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 5)
		locker.Unlock()
	}()
	go func() {
		defer wg.Done()
		if locker.TryLock() {
			panic("must failed")
		}
		if !locker.TryLockTimeout(time.Second * 7) {
			panic("must failed")
		}
		t.Log("lock after 7s, success!")
	}()
}

func TestConcurrencyLock(t *testing.T) {
	locker := New()
	wg := new(sync.WaitGroup)
	wg.Add(3)
	defer wg.Wait()
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			locked := locker.TryLockTimeout(time.Second)
			if locked {
				t.Log("g1 lock success")
				time.Sleep(time.Second)
				locker.Unlock()
			} else {
				t.Log("g1 lock failed")
			}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			locked := locker.TryLockTimeout(time.Second)
			if locked {
				t.Log("g2 lock success")
				time.Sleep(time.Second)
				locker.Unlock()
			} else {
				t.Log("g2 lock failed")
			}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			locked := locker.TryLockTimeout(time.Second)
			if locked {
				t.Log("g3 lock success")
				time.Sleep(time.Second)
				locker.Unlock()
			} else {
				t.Log("g3 lock failed")
			}
		}
	}()
}
