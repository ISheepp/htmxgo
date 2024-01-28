package api

import "fmt"

func Go(f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("panic: %+v", err)
			}
		}()
		f()
	}()
}
