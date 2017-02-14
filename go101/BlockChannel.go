package main

import (
	"fmt"
	"time"
	"runtime"
)

func main() {


	fmt.Println(runtime.NumCPU());
	fmt.Println(runtime.NumGoroutine());

	var channel chan string;
	channel = make(chan string);
	go goRoutine1Receive(channel);
	go goRoutine2Send(channel);

	fmt.Println(runtime.NumGoroutine());




	time.Sleep(time.Duration(20 * time.Second));
}

func goRoutine1Receive(channel chan string) {

	for {
		fmt.Println("Before read");
		time.Sleep(time.Second * 5);
		var data = <-channel;
		fmt.Println("Read already " + data);
	}

}

func goRoutine2Send(channel chan string) {

	fmt.Println("Before Insert");
	channel <- "Hello " + time.Now().String();
	fmt.Println("Inserted 1");

	channel <- "Hello " + time.Now().String();
	fmt.Println("Inserted 2");

}
