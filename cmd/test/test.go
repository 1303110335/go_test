package main

import (
   "fmt"
   "time"
)

func main () {
   var i=3

   go func(a int) {
      fmt.Println(a)
      fmt.Println("1")
   }(i)
   fmt.Println("2")
   //因为程序会优先执行主线程，主线程执行完成后，程序会立即退出，没有多余的时间去执行子线程。如果在程序的最后让主线程休眠1秒钟，那程序就会有足够的时间去执行子线程。
   time.Sleep(1 * time.Second)
}