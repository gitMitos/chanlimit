package main

import (
	"fmt"
	"math/rand"
	"time"

	uuid "github.com/nu7hatch/gouuid"
)

func main() {

	limiter := make(chan int, 3)
	info := make(chan uuid.UUID)

	go func(){
		for {
			u4, _ := uuid.NewV4()
			info <- *u4
		}
	}()

	for i := range info {
		limiter <- 1
		go func(i uuid.UUID){
			process(i)
			<-limiter
		}(i)
		
	}

}

func process(i uuid.UUID) {
	time.Sleep(time.Second * time.Duration(rand.Intn(10)) )
	fmt.Println("Processing: ", &i)
}