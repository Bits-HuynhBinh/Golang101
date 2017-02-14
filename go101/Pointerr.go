package main;

import (
	"fmt"
)

func main()  {

	testPointer2();


}


type Vertex struct {
	X int
	Y int
}

func testPointer2()  {

	var p *Vertex;

	var vertext = Vertex{1,5};

	p = &vertext;


	fmt.Println((*p).X);
	fmt.Println((*p).Y);

	fmt.Println(p.X);
	fmt.Println(p.Y);



	fmt.Println(p);
	fmt.Println(vertext);
	fmt.Println(*p);

}

func testPointer1()  {

	var p *int;

	var v int = 13;

	p = &v;

	fmt.Println(p);
	fmt.Println(*p);
	fmt.Println(&v);
	fmt.Println(v);
}
