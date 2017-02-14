package main

import (
	"time"
	"fmt"
	"math/rand"
)

func main() {

	rand.Seed(time.Now().UnixNano());

	var start = time.Now();

	//var results = google("Go lang");

	//var results = google1("Go lang");

	var results = google4("Go lang");

	var elapsed = time.Since(start);

	fmt.Println(results);
	fmt.Println(elapsed);


	//google3("Go lang");

}

type Result string;

type Search func(query string) Result;

func fakeSearch(kind string) Search {

	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

var Web = fakeSearch("WEB");

var Image = fakeSearch("IMAGE");

var Video = fakeSearch("VIDEO");

func google(query string) []Result {

	var results []Result;

	results = append(results, Web(query));
	results = append(results, Image(query));
	results = append(results, Video(query));

	return results;
}

func google1(query string) []Result {

	var chan1 = make(chan Result);

	var results []Result;

	go func() {
		chan1 <- Web(query)
	}();
	go func() {
		chan1 <- Image(query)
	}();
	go func() {
		chan1 <- Video(query)
	}();

	for i := 0; i < 3; i++ {
		var result = <-chan1;
		results = append(results, result);
	}

	return results;
}

func google2(query string) []Result {

	var chan1 = make(chan Result);

	var results []Result;

	go func() {
		chan1 <- Web(query)
	}();
	go func() {
		chan1 <- Image(query)
	}();
	go func() {
		chan1 <- Video(query)
	}();

	var timeout = time.After(80 * time.Millisecond);

	for i := 0 ; i < 3 ; i++ {
		select
		{
			case result := <-chan1:
				results = append(results, result);

			case <-timeout:
				results = append(results, "TIME OUT");
		}
	}

	return results;
}

func firstResult(query string, replicas ...Search ) Result  {

	var chan1 = make(chan Result);

	var searchReplica = func(i int)  { chan1 <- replicas[i](query)};

	for i := range replicas {

		go searchReplica(i);
		
	}

	return <- chan1;
}

func google3(query string)  {

	rand.Seed(time.Now().UnixNano());

	var start = time.Now();

	var results = firstResult(query, fakeSearch("WEB 1"), fakeSearch("WEB 2"));

	var elapsed = time.Since(start);

	fmt.Println(results);
	fmt.Println(elapsed);
}

func google4(query string) []Result  {

	var results []Result;

	var Web1 = fakeSearch("Web server 1");
	var Web2 = fakeSearch("Web server 2");

	var Image1 = fakeSearch("Image server 1");
	var Image2 = fakeSearch("Image server 2");

	var Video1 = fakeSearch("Video server 1");
	var Video2 = fakeSearch("Video server 2");


	var chan1 = make(chan Result);

	go func() { chan1 <- firstResult(query, Web1, Web2) }();
	go func() { chan1 <- firstResult(query, Image1, Image2) }();
	go func() { chan1 <- firstResult(query, Video1, Video2) }();


	var timeout = time.After(time.Millisecond * 80);

	for i := 0; i < 3; i++ {

		select
		{
			case r := <- chan1 : results = append(results, r);

			case <- timeout : results = append(results, "Time out");
		}
	}

	return results;
}



