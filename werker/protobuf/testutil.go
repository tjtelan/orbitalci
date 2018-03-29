package protobuf

import (
	"google.golang.org/grpc"
	"io"
	"golang.org/x/net/context"
)

//type BuildClient interface {
//	BuildInfo(ctx context.Context, in *Request, opts ...grpc.CallOption) (Build_BuildInfoClient, error)
//}

func NewFakeBuildClient(logLines []string) *fakeBuildClient {
	return &fakeBuildClient{
		outputLines: logLines,
	}
}

type fakeBuildClient struct {
	index int
	outputLines []string
	grpc.ClientStream
}

func (c *fakeBuildClient) CloseSend() error {
	return nil
}

func (c *fakeBuildClient) Recv() (*Response, error) {
	if c.index + 1 > len(c.outputLines) {
		return nil, io.EOF
	}
	resp := &Response{OutputLine: c.outputLines[c.index]}
	c.index++
	return resp, nil
}

func (c *fakeBuildClient) Context() context.Context {
	return context.TODO()
}

func (c *fakeBuildClient) SendMsg(m interface{}) error {
	return nil
}

func (c *fakeBuildClient) RecvMsg(m interface{}) error {
	if c.index + 1 > len(c.outputLines) {
		return io.EOF
	}
	original, ok := m.(Response)
	if ok {
		original.OutputLine = c.outputLines[c.index]
	}
	c.index++
	return nil
}
