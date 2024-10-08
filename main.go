package main

import (
	"sync"

	"github.com/Apt-Play/node/module"
	rmtpsetup "github.com/Apt-Play/node/rtmp_setup"
)

func main() {
	var wg sync.WaitGroup
	rtmpModule := rmtpsetup.RTMPModule{}
	modules := make([]module.Module, 1)
	modules[0] = &rtmpModule
	wg.Add(len(modules))
	startNode(modules, &wg)
	wg.Wait()
}

func startNode(modules []module.Module, wg *sync.WaitGroup) []error {
	errors := make([]error, len(modules))
	for i := range len(modules) {
		go func() {
			if err := modules[i].Run(wg); err != nil {
				errors[i] = err
			}
		}()
	}
	return errors
}
