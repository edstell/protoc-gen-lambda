package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/edstell/protoc-gen-lambda/generator"
	"github.com/edstell/protoc-gen-lambda/template"
	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	req := &plugin.CodeGeneratorRequest{}
	if err := proto.Unmarshal(data, req); err != nil {
		panic(fmt.Sprint("failed to unmarshal the generation request:", err))
	}
	funcs := map[string]interface{}{
		"toLower": strings.ToLower,
		"rhs": func(s string) string {
			ss := strings.Split(s, ".")
			return ss[len(ss)-1]
		},
	}
	names := []string{"client", "router"}
	files := make([]*plugin.CodeGeneratorResponse_File, 0, len(req.ProtoFile)*len(names))
	for _, protoFile := range req.ProtoFile {
		if len(protoFile.Service) == 0 {
			continue
		}
		if !contains(req.FileToGenerate, *protoFile.Name) {
			continue
		}
		for _, name := range names {
			template, err := template.FromAsset(name, funcs)
			if err != nil {
				panic(err)
			}
			file, err := generator.FromTemplate(name, template).Generate(protoFile)
			if err != nil {
				panic(fmt.Sprint("failed to generate file", err))
			}
			files = append(files, file)
		}
	}
	data, err = proto.Marshal(&plugin.CodeGeneratorResponse{
		File: files,
	})
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(data)
}

func contains(strs []string, str string) bool {
	for _, s := range strs {
		if s == str {
			return true
		}
	}
	return false
}
