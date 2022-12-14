package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"zhycan/internal/config"
	"zhycan/internal/logger"
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

	l, _ := logger.GetManager().GetLogger()
	if l != nil {
		l.Log(logger.NewLogObject(
			logger.INFO, "test", logger.FuncMaintenanceType, time.Now().UTC(), "this is a test", nil))
	}

	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
}
