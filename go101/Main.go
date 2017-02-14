package main

import "fmt"
import (
	"time"
	"math/rand"
)

func main(){

	/*var channel chan string;
	channel = make(chan string);
	go push("Hello", channel);
	go take1(channel);
	go take2(channel);*/
	/*for i := 0; i < 10 ; i++  {
		var a = <-channel;
		fmt.Println(a);
	}*/


	/*var p Person;
	p.name = "Binh"
	p.age = 27
	p.email = "hb@gmail.com"
	fmt.Println(p.toString());*/


	/*var chan1 = push1("Binh");
	var chan2 = push1("Mia");
	//var results = merge(chan1, chan2);
	var results2 = mergeUsingSelect(chan1, chan2);
	for i := 0; i < 15 ; i++  {
		var a = <-results2;
		fmt.Println(a);
	}*/


	var chan1 = make(chan string);
	var chan2 = make(chan string);
	var chan3 = make(chan string);



	producer(chan1, "chan 1");
	producer(chan2, "chan 2");
	producer(chan3, "chan 3");

	time.Local.String();


	//consumer(chan1);

	var result = select2(chan1, chan2, chan3);

	for {
		var a = <-result;
		fmt.Println(a);
	}

	time.Sleep(time.Duration(15 * time.Second));
}


func producer(queue chan string, name string) {
	go func() {
		for
		{
			queue <- fmt.Sprintf("%s: %s", name,time.Now());
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
}

func consumer(queue chan string) {
	go func() {
		for
		{
			var item = <-queue;
			fmt.Println(item);
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
}

func select2(chan1 chan string, chan2 chan string, chan3 chan string) <- chan string {

	var chanS = make(chan string);

	go func() {
		for
		{
			select
			{
				case v1 := <-chan1 : chanS <- v1;
				case v2 := <-chan2 : chanS <- v2;
				case v3 := <-chan3 : chanS <- v3;
			}
		}
	}()

	return chanS;
}

func select1()  {
	var c1, c2, c3 chan int

	select {
	case v1 := <-c1:
		fmt.Printf("received %v from c1\n", v1)
	case v2 := <-c2:
		fmt.Printf("received %v from c2\n", v2)
	case c3 <- 23:
		fmt.Printf("sent %v to c3\n", 23)
	default:
		fmt.Printf("no one was ready to communicate\n")
	}
}

func mergeUsingSelect(chan1 chan string, chan2 chan string) <- chan string {

	var chan3 = make(chan string);

	go func(){
		for
		{
			select
			{
				case s := <-chan1 : chan3 <- s;
				case s := <-chan2 : chan3 <- s;
 			}

		}
	}()

	return chan3;

}


func merge(chan1 chan string, chan2 chan string) <-chan string {

	var chan3 = make(chan string);


	go func() {
		for
		{
			var a = <- chan1;
			chan3 <- a;
		}
	}();

	go func() {
		for
		{
			var a = <- chan2;
			chan3 <- a;
		}
	}()

	return chan3;
}

func push1(data string) chan string {
	var chan1 = make(chan string);
	go func() {

		for i :=0 ; ; i++   {
			chan1 <- fmt.Sprintf("%s %d", data, i);
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
		}

	}();
	return chan1;
}

func push(data string, channel chan string){
	for i :=0 ; ; i++  {
		channel <- fmt.Sprintf("%s %d", data, i);
		time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
	}
}

func take1(channel chan string){
	for i := 0; i < 10 ; i++  {
		var a = <-channel;
		fmt.Printf("%s %s i: %d \n","take 1 a: ", a, i);
	}
}

func take2 (channel chan string){
	for i := 0; i < 10 ; i++  {
		var a = <-channel;
		fmt.Printf("%s %s i: %d \n","take 2 a: ", a, i);
	}
}

func printSomething(something string, howManyTime int){
	for i := 0; i < howManyTime ; i++  {
		fmt.Printf("%s %d \n", something, i);
		time.Sleep(time.Second);
	}
}


type Person struct {
	name string;
	age int;
	email string;
}


func (p Person) toString() string {
	return  fmt.Sprintf("Person [Name: %s, Email: %s, Age: %d]",p.name, p.email, p.age);
}
