package ticker

import (
	"fmt"
	"time"
)

func Schedule() chan bool {
	ticker := time.NewTicker(1 * time.Second)

	stopChan := make(chan bool)
	go func(ticker *time.Ticker) {
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				fmt.Println("Ticker 2...")
			case stop := <-stopChan:
				if stop {
					fmt.Println("Ticker 2 stop")
					return
				}
			}
		}
	}(ticker)
	return stopChan
}

func Stop(stopChan chan bool) {
	time.Sleep(10 * time.Second)
	stopChan <- true
}
