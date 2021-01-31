// DO NOT EDIT: Client was generated from 'example.proto'
package exampleproto

import (
	"context"
	"encoding/json"

	awsreq "github.com/aws/aws-sdk-go/aws/request"
	invoker "github.com/edstell/lambda-invoker"
	"google.golang.org/protobuf/encoding/protojson"
)

// Client is a generated client for the service defined in 'Service'.
// It exposes methods to call procedures available on the service, and handles
// the packing and unpacking of requests and responses for transport.
type Client interface {
	Do(context.Context, *DoRequest, ...awsreq.Option) (*DoResponse, error)
}

// client is the internal implementation of Client.
type client struct {
	do   *invoker.Invoker
}

// NewClient initializes a Client, configuring it to use the provided 
// unmarshaler for unpacking errors to the error implementation of your choice.
func NewClient(i invoker.LambdaInvoker, arn string, unmarshaler func(json.RawMessage) error) Client {
	return &client{
		do:   invoker.New(i, arn, invoker.AsProcedure("Do", unmarshaler)),
	}
}

func (c *client) Do(ctx context.Context, req *DoRequest, opts ...awsreq.Option) (*DoResponse, error) {
	payload, err := protojson.Marshal(req)
	if err != nil {
		return nil, err
	}
	rsp, err := c.do.Invoke(ctx, payload, opts...)
	if err != nil {
		return nil, err
	}
	out := &DoResponse{}
	if err := protojson.Unmarshal(rsp, out); err != nil {
		return nil, err
	}
	return out, nil
}
