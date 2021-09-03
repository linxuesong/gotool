package concurrency

import "reflect"

// 任务安排
// 使用chan实现 Or-Done的 多个任务中一个完成就可以触发
// params channels 需要检测任务中完成后会写入的chan
// return 返回一个检测所有任务的chan，  如果一个任务完成了，那么这个chan会被关闭    读取chan会返回默认零值    如果没有任务完成 chan开启但没有值  读取会阻塞
func OrDone(channels ...chan interface{}) <-chan interface{} {
	ordone := make(chan interface{})

	go func() {
		defer close(ordone)
		// 反射创建指定数量的case
		var causes []reflect.SelectCase
		for channel := range channels {
			causes = append(causes, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(channel),
			})
		}
		// 当某一个cause触发后走到defer 关闭ordone 这样外面  <-ordone会退出阻塞
		reflect.Select(causes)
	}()
	return ordone
}
