package times

import (
	"fmt"
	"os"
	"time"
)

func Times() {
	var time1 = time.Now()
	fmt.Printf("time1 %v\n", time1)
	// time1 2015-09-01 17:59:31.73600891 +0700 WIB

	var time2 = time.Date(2011, 12, 24, 10, 20, 0, 0, time.UTC)
	fmt.Printf("time2 %v\n", time2)
}

func ParsingTime() {
	var layoutFormat, value string
	var date time.Time

	layoutFormat = "2006-01-02 15:04:05"
	value = "2015-09-02 08:04:00"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t->", date.String())

	layoutFormat = "02/01/2006 MST"
	value = "02/09/2015 WIB"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t\t->", date.String())
}

func PredenfinedTime() {
	var date, _ = time.Parse(time.RFC822, "02 Sep 15 08:00 WIB")
	fmt.Println(date.String())
}

func TimeToString() {
	var date, _ = time.Parse(time.RFC822, "02 Sep 15 08:00 WIB")

	var dateS1 = date.Format("Monday 02, January 2006 15:04 MST")
	fmt.Println("dateS1", dateS1)
	// Wednesday 02, September 2015 08:00 WIB

	var dateS2 = date.Format(time.RFC3339)
	fmt.Println("dateS2", dateS2)
}

func HandleErrorTime() {
	var date, err = time.Parse("06 Jan 15", "02 Sep 15 08:00 WIB")

	if err != nil {
		fmt.Println("error", err.Error())
		return
	}

	fmt.Println(date)
}

func TimesSlepp() {
	fmt.Println("start")
	time.Sleep(time.Second * 4)
	fmt.Println("after 4 seconds")
}

func ScedhulerSleep() {
	for true {
		fmt.Println("Hallo")
		time.Sleep(time.Second * 1)
	}
}

func NewTimers() {
	var timer = time.NewTimer(4 * time.Second)

	fmt.Println("start")
	<-timer.C
	fmt.Println("finish")
}

func AftersTime() {
	var ch = make(chan bool)

	time.AfterFunc(4*time.Second, func() {
		fmt.Println("expired")
		ch <- true
	})

	fmt.Println("start")
	<-time.After(4 * time.Second)
	fmt.Println("expired")
	fmt.Println("finish")
}

func SchedulerTicker() {
	done := make(chan bool)

	ticker := time.NewTicker(time.Second)

	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()

	for {
		select {
		case <-done:
			ticker.Stop()
			return
		case t := <-ticker.C:
			fmt.Println("Hello !!", t)
		}
	}
}

func timer(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		ch <- true
	})
}

func watcher(timeout int, ch <-chan bool) {
	<-ch
	fmt.Println("\ntimer out! no answer more than", timeout, "seconds")
	os.Exit(0)
}

func TimerGoruotine() {
	var timeout = 5
	var ch = make(chan bool)

	go timer(timeout, ch)
	go watcher(timeout, ch)

	var input string
	fmt.Print("what is 725/25 ? ")
	fmt.Scan(&input)

	if input == "29" {
		fmt.Println("the anser is right!!")
	} else {
		fmt.Println("th answer wrong!")
	}
}

func Duration() {
	start := time.Now()

	time.Sleep(5 * time.Second)

	duration := time.Since(start)

	fmt.Println("time elapsed in seconds:", duration.Seconds())
	fmt.Println("time elapsed in minutes:", duration.Minutes())
	fmt.Println("time elapsed in hours:", duration.Hours())
}

func KalkulasiSelisih() {
	t1 := time.Now()

	time.Sleep(5 * time.Second)
	t2 := time.Now()

	duration := t1.Sub(t2)

	fmt.Println("time elapsed in seconds:", duration.Seconds())
	fmt.Println("time elapsed in minutes:", duration.Minutes())
	fmt.Println("time elapsed in hours:", duration.Hours())
}
