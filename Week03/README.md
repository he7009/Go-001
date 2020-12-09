# GO-并发、并行

## 并发、并行
**并发（concurrency）不是并行（parallelism）**。并行是让不同的代码片段同时在不同的物理处理器上执行。并行的关键是同时做很多事情，而并发是指同时管理很多事情，这些事情可能只做了一半就被暂停去做别的事情了。

并发增加了效率，但是并发读写同一个资源可能导致随机性的Bug。对于这种问题可以通过携程间的通信（chan）或者加锁的形式来解决。

尽量通过通信，而不是加锁的方式

## sync、**sync/atomic**
sync 包提供一些加锁、安全的原子性操作相关功能：

1. sync.WaitGroup 用来等待 goroutines 集合完成
1. sync.RWMutex、sync.RWMutex 对资源加锁
1. sync.Once 只执行一次
1. atomic.AddInt32() ... 等原子性操作



### sync.WaitGrou
```go
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("GoGoGo")
		}()
	}
	wg.Wait()
}
```
### sync.RWMutex、sync.RWMutex
```go
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var cnt int32
var lc sync.Mutex
func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				lc.Lock()
				cnt++
				lc.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Printf("cnt 10 * 1000 = %v",cnt)
}
```
### sync.Once
```go
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
func main() {
	once := sync.Once{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(func() {
				fmt.Println("只执行一次")
			})
		}()
	}
	wg.Wait()
}

```
### atomic package
atomic 针对6中数据类型提供了5种原子操作

- 数据类型：*int32、*int64、*uint32、*uint64、*uintptr、*unsafe.Pointer
- 操作：Swap、CompareAndSwap、Add、Load、Store
- Swap：存储新值，返回旧值
- CompareAndSwap：比较值：
   - 如果相等，返回true并且替换为新的值；
   - 如果不相等，返回false并且不做替换
- Add：增加
- Load：取值
- Store：存储
```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var cnt int32
func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt32(&cnt,1)
			}
		}()
	}
	wg.Wait()
	fmt.Printf("cnt 10 * 1000 = %v",cnt)
}

```