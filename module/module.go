package module

import "sync"

type Module interface {
	Run(wg *sync.WaitGroup) error
}
