# timed_mutex

![go report](https://goreportcard.com/badge/github.com/ooopSnake/timed_mutex)
![license](https://img.shields.io/badge/license-MIT-brightgreen.svg)
[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://github.com/ooopSnake/timed_mutex)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat)](https://github.com/ooopSnake/timed_mutex/pulls)
[![Ask Me Anything !](https://img.shields.io/badge/Ask%20me-anything-1abc9c.svg)](https://github.com/ooopSnake/timed_mutex/issues)


## Getting Started

### 📌 Lock()

acquire a lock,block until lock obtained

### 📌 UnLock()

release lock

### 📌 TryLock() (locked bool) 

try acquire a lock, `true` if lock obtained otherwise return `false ` immediately.


### 📌 TryLockTimeout(t time.Duration) (locked bool)

try acquire a lock  in specified duration, `true` if lock obtained.
return `false` if the mutex still unavailable until specified time point has been reached.

## Example


init lock 😏

```go
locker := timed_mutex.New()
```

unlock 😐

```go
locker.UnLock()
```

direct lock 😐

```go
locker.Lock()
```

trylock 😎
```go
locked :=  locker.TryLock()
```

trylock with timeout 😎

```go
locked := locker.TryLockTimeout(time.Second*5)
```

## License

Released under the [MIT License](https://github.com/ooopSnake/timed_mutex/blob/master/LICENSE)