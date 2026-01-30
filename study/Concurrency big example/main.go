package main

import (
	"concurrency/miner"
	"concurrency/postman"
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var coal atomic.Int64
	var mails []string
	mtx := sync.Mutex{}

	minerContext, minerCancel := context.WithCancel(context.Background())
	postmanContext, postmanCancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("--------->Рабочий день шахтёров окончен")
		minerCancel()
	}()

	go func() {
		time.Sleep(6 * time.Second)
		fmt.Println("--------->Рабочий день почтальонов окончен")
		postmanCancel()
	}()

	coalTransferPoint := miner.MinerPool(minerContext, 100)
	mailTransferPoint := postman.PostmanPool(postmanContext, 100)
	initTime := time.Now()

	wg := sync.WaitGroup{}

	wg.Go(func() {

		for v := range coalTransferPoint {
			coal.Add(int64(v))
		}
	})

	wg.Go(func() {

		for v := range mailTransferPoint {
			mtx.Lock()
			mails = append(mails, v)
			mtx.Unlock()
		}
	})

	wg.Wait()

	fmt.Println("Суммарно добытый уголь:", coal.Load())
	mtx.Lock()
	fmt.Println("Суммарно отправленные письма:", len(mails))
	mtx.Unlock()
	fmt.Println("Времени затрачено:", time.Since(initTime))
}
