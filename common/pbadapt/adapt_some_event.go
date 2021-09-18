package pbadapt

import (
	"mir/common/lib"
	"mir/common/pb"
)

func MarshalSomeEvent(e lib.SomeEvent) *pb.SomeEvent {
	return &pb.SomeEvent{
		Content: 0,
	}
}

func UnmarshalSomeEvent(e *pb.SomeEvent) lib.SomeEvent {
	if e == nil {
		return lib.SomeEvent{}
	}

	return lib.SomeEvent{
		Content: float64(e.Content),
	}
}
