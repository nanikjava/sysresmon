package main

import (
	"fmt"
	"github.com/nmrshll/sysresmon"
)

func main() {
	sampler := sysresmon.NewSampler().StartSampling().Aggregate()
	for {
		select {
		case aggregate := <-sampler.AggregateChan:
			fmt.Printf("CPU usage is %f%%\n", aggregate.CPUUsage)
		}
	}
}
