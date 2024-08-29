package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		fmt.Println("Sending request", i, time.Now().Format("15:04:05.000"))
		wg.Add(1)
		go func() {
			response, err := http.Get("http://localhost:8000/statistics")
			if err != nil {
				fmt.Println("Error", err)
				return
			}
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					return
				}
			}(response.Body)

			body, err := io.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Response received", i, time.Now().Format("15:04:05.000"), len(body),
				'\n', string(body))
			wg.Done()
		}()
	}
	wg.Wait()
}
