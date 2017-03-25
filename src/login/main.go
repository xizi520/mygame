// login project main.go
package main

import (
	"fmt"
	"net"
	"net/http"
	"sync"
)

func main() {
	fmt.Println("Hello World!")

	http.HandleFunc()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			fmt.Printf("start server failed, err: %s", err.Error())
			wg.Done()
		}
	}()

	go func() {
		err := http.ListenAndServeTLS(":8000", "", "", nil)
		if err != nil {
			fmt.Printf("start TLS server failed, err: %s", err.Error())
			wg.Done()
		}
	}()

	wg.Wait()
}
