/*
@Time : 2020/10/19 10:25 上午
@Author : mac
@File : mutexstudy.go
@Software: GoLand
*/
package mutex

import (
	"fmt"
	"sync"
)

func MutexPrint1() {
	mutex := sync.Mutex{}
	mutex.Lock()
	fmt.Println(1)
	defer mutex.Unlock()
}

func MutexPrint2() {
	mutex := sync.Mutex{}
	mutex.Lock()
	fmt.Println(2)
	defer mutex.Unlock()
}

func MutexPrint3() {
	mutex := sync.Mutex{}
	mutex.Lock()
	fmt.Println(3)
	defer mutex.Unlock()
}

/**
 * @Description //同步执行
 * @Param
 * @return
 **/
func MutexDemo2() {
	wg := sync.WaitGroup{}
	//wg.Add(10)来增加wg的内部计数器为10
	wg.Add(10)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			//wg.Done()减少(1)wg的内部计数器
			wg.Done()
		}
	}()
	// 告诉程序等待wg中的计数器清零才能继续执行
	wg.Wait()
}

/**
 * @Description
	1、可以随便读，多个goroutine同时读。
	2、写的时候，不能读也不能写。
	主要有下面四个API构成，读锁RLock,RUnlock,写锁Lock,Unlock。
 * @Param
 * @return
 **/
func MutexDemo3() {
	wg := sync.RWMutex{}
	wg.Lock()
	fmt.Println("读写锁一般用于读多写少的情况")
	wg.Unlock()
}

/**
 * @Description 互斥锁是串行执行的
 * @Param
 * @return
 **/
func MutextDemo1() {
	MutexPrint1()
	MutexPrint2()
	MutexPrint3()
}
