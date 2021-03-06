package main

import (
	"fmt"
	"strconv"
)

func main()  {
	fmt.Println("hello")

	var i string = "200";
	fmt.Println("hanzs " + i);

	j := "300k"
	a := true
	const gaoyany = "hanlx"

	point := &i;
	value := *point;

	fmt.Println("dd", j, gaoyany, point, value);

	hh, error := strconv.Atoi(j);

	if (a) {
		fmt.Println(hh, error, point);
	}

}
