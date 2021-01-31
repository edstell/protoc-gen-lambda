package template

//go:generate go-bindata -modtime 1 -pkg template -ignore=\.go$ -ignore=.DS_Store .
//go:generate go fmt bindata.go
