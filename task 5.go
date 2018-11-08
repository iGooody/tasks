package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)

type request struct {
	time    int
	timeOut bool
}

func getRequest(r request, wg sync.WaitGroup, client http.Client, url string) {
	defer wg.Done()
	start := time.Now()
	_, err := client.Get(url)
	end := time.Now()
	connTime := end.Sub(start) / time.Millisecond
	connTimeOut := false
	if e, ok := err.(net.Error); ok && e.Timeout() {
		connTimeOut = true
	}
	r.time = int(connTime)
	r.timeOut = connTimeOut
}

func printMetriki(req []request, fullTime int) {
	var (
		minTime      int = -1
		maxTime      int
		countTimeOut int
		allTime      int
	)
	for _, request := range req {
		if !request.timeOut {
			allTime += request.time
			if request.time > maxTime {
				maxTime = request.time
			}
			if request.time < minTime || minTime == -1 {
				minTime = request.time
			}
		} else {
			countTimeOut++
		}

	}
	var avgTime int
	succes := len(req) - countTimeOut
	if succes != 0 {
		avgTime = allTime / succes
	} else {
		minTime = 0
	}

	fmt.Println("Общее время %d"+"Среднее время %d"+"Макс время %d"+"Мин время %d"+"Кол-во таймаутов %d", fullTime, avgTime, maxTime, minTime, countTimeOut)
}

func main() {
	var (
		url     string
		count   int
		timeOut int
	)
	flag.StringVar(&url, "url", "", "Ссылка")
	flag.IntVar(&count, "count", 0, "Количество повторов")
	flag.IntVar(&timeOut, "timeout", 0, "Таймаут")
	flag.Parse()

	if url == "" || count == 0 || timeOut == 0 {
		fmt.Println("Неверные выходные данные")
		os.Exit(1)
	}
	requests := make([]request, count)
	client := http.Client{
		Timeout: time.Duration(timeOut) * time.Millisecond,
	}

	var wg sync.WaitGroup
	wg.Add(count)

	startGlobal := time.Now()
	for i := 0; i < count; i++ {
		go getRequest(requests[i], wg, client, url)
	}
	wg.Wait()
	endGlobal := time.Now()
	fullTime := endGlobal.Sub(startGlobal) / time.Millisecond

	printMetriki(requests, int(fullTime))
}