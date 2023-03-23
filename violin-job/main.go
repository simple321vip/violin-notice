package main

import "violin-home.cn/violin-job/config"
import "violin-home.cn/violin-job/queue"

func main() {

	// config
	conf := config.InitConfig()

	// setting mongo
	config.ConnectToDB(conf)

	// create job queue
	go func() {
		q := queue.NewQueue()
		q.Push()
	}()

	go func() {
		q := queue.NewQueue()
		q.Pop()
	}()

	// listen and execute

}
