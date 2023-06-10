package main

import (
	"github.com/wolfogre/go-pprof-practice/demo"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetOutput(os.Stdout)

	runtime.GOMAXPROCS(1)
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)

	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	// 原本的demo
	//for {
	//	for _, v := range animal.AllAnimals {
	//		v.Live()
	//	}
	//	time.Sleep(time.Second)
	//}
	// 自己造的demo
	demo.TestGoroutine()
}
