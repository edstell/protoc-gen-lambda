# protoc-gen-lambda
A protobuf plugin for generating a lambda client and router for rpc-like function invocation.

## Installation
Install `protoc-gen-lambda` binary:
```
go install ./../protoc-gen-lambda
```
This will compile the binary and install it to be available to 'protoc' when 
generating protobuf files.

## Generating '.pb.go' files
Once the plugin has been installed (follow installation instructions above), generate
lambda client and router with the following (this will output generated files to
the same directory as the source proto file):
```
protoc -I=example --go_out=$GOPATH/src --lambda_out=example example/example.proto
```

## Usage
Once you've generated your client and router, use them in your lambda service as
follows.

### Router
Configure the receiving lambda function to route requests like so:
```
// handler should be your concrete implementation.
var handler handlerproto.Handler
// marshaler should marshal your error implementation to json bytes.
var marshaler func(error) (json.RawMessage, error)
router := exampleproto.NewRouter(handler, marshaler)
lambda.Start(router.Handle)
```

### Client
Configure client the the calling lambda function like so:
```
// unmarshaler should unmarshal json bytes to your error implementation.
var unmarshaler func(json.RawMessage) error
example := exampleproto.NewClient(lambda.New(session), "example-arn", unmarshaler)
// Call the 'Do' procedure in the lambda function:
rsp, err := example.Do(ctx, &exampleproto.DoRequest{
	Id: "the-thing",
})
```
