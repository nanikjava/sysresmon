package main

import (
	"fmt"
	"github.com/nmrshll/sysresmon"
)

func main() {
	sampler := sysresmon.NewSampler().StartSampling()
	var s sysresmon.MemSample
	for {
		select {
		case sampleSet := <-sampler.SampleSetChan:
			s = sampleSet.MemSample

			fmt.Printf("MemTotal = %v KB, MemFree = %v KB, MemUsed = %v KB\n", s.MemTotal,
				s.MemFree, s.MemUsed)
		}
	}
}
