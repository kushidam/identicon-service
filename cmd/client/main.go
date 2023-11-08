package main

import (
	"flag"
	"fmt"
)


var (
	userText    = flag.String("text", "nullrocks", "use decides the resulting figure")
)

func main(){
	flag.Parse()
	fmt.Println(userText)
	
}