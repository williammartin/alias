package main

import (
	"fmt"
	"github.com/williammartin/converter"
	"github.com/williammartin/temperature"
)

func main() {
	c := converter.New()
	t := &temperature.Temperature{Value: 10, Unit: "Celsius"}
	fmt.Println(t)

	c.Convert(t)
}
