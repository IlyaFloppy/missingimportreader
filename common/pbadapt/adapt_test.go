package pbadapt

import (
	"fmt"
	"mir/common/lib"
	"testing"
)

func TestAdaptersSomeEvent(t *testing.T) {
	for x := -3; x < 4; x++ {
		t.Run(fmt.Sprintf("content=%d", x), func(t *testing.T) {
			v := lib.SomeEvent{
				Content: float64(x),
			}

			r := UnmarshalSomeEvent(MarshalSomeEvent(v))

			if v != r {
				t.Errorf("%v != %v", v, r)
			}
		})
	}
}
