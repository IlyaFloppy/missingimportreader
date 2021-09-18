package pb

import "google.golang.org/protobuf/runtime/protoimpl"

type SomeEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content float32 `protobuf:"fixed32,1,opt,name=content,proto3" json:"content,omitempty"`
}
