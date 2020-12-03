package core

import (
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	
	Inbound []*InboundHandlerConfig `protobuf:"bytes,1,rep,name=inbound,proto3" json:"inbound,omitempty"`
	
}

type InboundHandlerConfig struct {
	
}