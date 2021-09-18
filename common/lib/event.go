package lib

import "fmt"

type IEvent interface {
	Apply(h *ObsHolder)
}

type SomeEvent struct {
	Content float64
}

func (e *SomeEvent) Apply(h *ObsHolder) {
	fmt.Println(h, e.Content)
}
