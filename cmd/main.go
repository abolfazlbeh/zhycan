package main

import (
	"fmt"
	"time"
	"zhycan/internal/config"
)

func main() {
	fmt.Println("Zhycan Project")

	path := "/Users/abolfazl.beh/Projects/zhycan/"
	initialMode := "dev"
	prefix := "ZHYCAN"

	err := config.CreateManager(path, initialMode, prefix)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		flag := config.GetManager().IsInitialized()
		if flag {
			break
		}
		time.Sleep(1 * time.Second)
	}
	fmt.Println("All modules is initialized")
}
