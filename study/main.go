/**
 *
	@goroutine:Go 允许使用 go 语句开启一个新的运行期线程， 即 goroutine，以一个不同的、
新创建的 goroutine 来执行一个函数。 同一个程序中的所有 goroutine 共享同一个地址空间。

	chan:通道可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。操作符 <- 用于指定通道的方向，发送或接收。如果未指定方向，则为双向通道。
	重点：
		通道在使用前必须先创建;
		默认情况下，通道是不带缓冲区的。发送端发送数据，同时必须有接收端接收相应数据
		带缓冲区的通道允许发送端的数据发送和接收端的数据获取处于异步状态，就是说发送端发送的数据可以放在缓冲区里面，可以等待接收端去获取数据，而不是立刻需要接收端去获取数据。
不过由于缓冲区的大小是有限的，所以还是必须有接收端来接收数据的，否则缓冲区一满，数据发送端就无法再发送数据了。
		带有缓冲区，发送方会先把数据发送到缓冲区。在接收方接受数据之前会阻塞。
		没有缓冲区，在发送数据时会死锁
	一个进程在启动的时候，会创建一个主线程，这个主线程结束的时候，程序进程也就终止了，所以一个进程至少有一个线程，这也是我们在main函数里，使用goroutine的时候，
要让主线程等待的原因，因为主线程结束了，程序就终止了，那么就有可能会看不到goroutine的输出。
	并行指的是在不同的物理处理器上同时执行不同的代码片段，并行可以同时做很多事情，而并发是同时管理很多事情，因为操作系统和硬件的总资源比较少，
所以并发的效果要比并行好的多，使用较少的资源做更多的事情，也是Go语言提倡的。
	Go默认是给每个可用的物理处理器都分配一个逻辑处理器，因为我的电脑是4核的，所以上面的例子默认创建了4个逻辑处理器
	runtime.GOMAXPROCS(1):强制只使用一个逻辑处理器
 **/

/**
 * @Param 协程
	1:defer函数会在普通函数返回之后执行
	2:sync同步锁:互斥锁sync.Mutex和读写锁sync.RWMutex.
	3：当有一个goroutine获取了互斥锁后，任何goroutine都不可以获取互斥锁，只能等待这个goroutine将互斥锁释放。互斥锁是串行的
	4：range关键字在使用channel的时候，会自动等待channel的动作一直到channel关闭。通俗点将就是可以channel可以自动开关。

 * @return
 **/
package main

import (
	"fmt"
	"gomodtest/study/channel"
)

/**
 * @Description //运行study目录中的demo
 * @Param
 * @return
 **/
func main() {
	channel.ChanDemo2()
	//mutex.MutexDemo3()
	//singleThread()
}

func printOne(cs chan int) {
	fmt.Println(1)
	cs <- 1
}

func printTwo(cs chan int) {
	<-cs
	fmt.Println(2)
	defer close(cs)
}
