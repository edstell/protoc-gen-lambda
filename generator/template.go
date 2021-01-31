package generator

import (
	"bytes"
	"fmt"
	"path"
	"text/template"

	google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

// FromTemplate returns a Generator that generates an output file from the
// template passed, naming the output with the suffix.
func FromTemplate(suffix string, t *template.Template) Generator {
	return GeneratorFunc(func(desc *google_protobuf.FileDescriptorProto) (*plugin.CodeGeneratorResponse_File, error) {
		filepath := *desc.Name
		name := filepath[0:len(filepath)-len(path.Ext(filepath))] + fmt.Sprintf(".%s.pb.go", suffix)
		var b bytes.Buffer
		if err := t.Execute(&b, desc); err != nil {
			return nil, err
		}
		content := b.String()
		return &plugin.CodeGeneratorResponse_File{
			Name:    &name,
			Content: &content,
		}, nil
	})
}
