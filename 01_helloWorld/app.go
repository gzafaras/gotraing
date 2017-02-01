package main

import "fmt"



func main() {
	fmt.Println(4545)

	for i := 0;i<= 100; i++{
		fmt.Printf("%d,%b,%#x\n",i,i,i)

	}

	var a int
	var b float64
	var c string
	var d bool

	f,g,h := 47,0.0,"gzaf"

	fmt.Printf("%v\n",a)
	fmt.Printf("%v\n",b)
	fmt.Printf("%v\n",c)
	fmt.Printf("%v\n",d)

	fmt.Printf("%t\n",f)
	fmt.Printf("%t\n",g)
	fmt.Printf("%t\n",h)

}
