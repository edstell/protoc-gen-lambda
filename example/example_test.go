package exampleproto

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	awsreq "github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/lambda"
	router "github.com/edstell/lambda-router"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type LambdaInvokerFunc func(context.Context, *lambda.InvokeInput, ...awsreq.Option) (*lambda.InvokeOutput, error)

func (f LambdaInvokerFunc) InvokeWithContext(ctx context.Context, input *lambda.InvokeInput, opts ...awsreq.Option) (*lambda.InvokeOutput, error) {
	return f(ctx, input, opts...)
}

type HandlerFunc func(context.Context, *DoRequest) (*DoResponse, error)

func (f HandlerFunc) Do(ctx context.Context, req *DoRequest) (*DoResponse, error) {
	return f(ctx, req)
}

func TestDoSuccess(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	id := "test-id"
	li := LambdaInvokerFunc(func(ctx context.Context, input *lambda.InvokeInput, opts ...awsreq.Option) (*lambda.InvokeOutput, error) {
		r := NewRouter(HandlerFunc(func(ctx context.Context, req *DoRequest) (*DoResponse, error) {
			assert.Equal(t, id, req.Id)
			return &DoResponse{Success: true}, nil
		}), nil)
		req := &router.Request{}
		if err := json.Unmarshal(input.Payload, req); err != nil {
			return nil, err
		}
		rsp, err := r.Handle(ctx, *req)
		if err != nil {
			return nil, err
		}
		bytes, err := json.Marshal(rsp)
		if err != nil {
			return nil, err
		}
		return &lambda.InvokeOutput{Payload: bytes}, nil
	})
	client := NewClient(li, "test-arn", nil)
	rsp, err := client.Do(ctx, &DoRequest{
		Id: id,
	})
	require.NoError(t, err)
	assert.True(t, rsp.Success)
}

func TestDoError(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	message := "test error message"
	li := LambdaInvokerFunc(func(ctx context.Context, input *lambda.InvokeInput, opts ...awsreq.Option) (*lambda.InvokeOutput, error) {
		r := NewRouter(HandlerFunc(func(ctx context.Context, req *DoRequest) (*DoResponse, error) {
			return nil, errors.New(message)
		}), func(err error) (json.RawMessage, error) {
			return []byte(fmt.Sprintf(`"%s"`, err.Error())), nil
		})
		req := &router.Request{}
		if err := json.Unmarshal(input.Payload, req); err != nil {
			return nil, err
		}
		rsp, err := r.Handle(ctx, *req)
		if err != nil {
			return nil, err
		}
		bytes, err := json.Marshal(rsp)
		if err != nil {
			return nil, err
		}
		return &lambda.InvokeOutput{Payload: bytes}, nil
	})
	client := NewClient(li, "test-arn", func(m json.RawMessage) error {
		var e interface{}
		if err := json.Unmarshal(m, &e); err != nil {
			return err
		}
		return errors.New(fmt.Sprint(e))
	})
	_, err := client.Do(ctx, &DoRequest{})
	require.Error(t, err)
	assert.Equal(t, message, err.Error())
}
