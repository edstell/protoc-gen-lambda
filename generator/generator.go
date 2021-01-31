package generator

import (
	google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

// Generator implementations generate a plugin response file from the passed
// file descriptor.
type Generator interface {
	Generate(*google_protobuf.FileDescriptorProto) (*plugin.CodeGeneratorResponse_File, error)
}

// GeneratorFunc type is an adapter to allow the use of ordinary functions as
// Generator implementations. If f is a function with the appropriate signature,
// GeneratorFunc(f) is a Generator that calls f.
type GeneratorFunc func(*google_protobuf.FileDescriptorProto) (*plugin.CodeGeneratorResponse_File, error)

// Generate calls f(desc).
func (f GeneratorFunc) Generate(desc *google_protobuf.FileDescriptorProto) (*plugin.CodeGeneratorResponse_File, error) {
	return f(desc)
}
