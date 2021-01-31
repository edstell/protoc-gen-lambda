// Code generated by go-bindata.
// sources:
// client.gotmpl
// router.gotmpl
// DO NOT EDIT!

package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _clientGotmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\xcd\x6e\xdc\x36\x10\x3e\x4b\x4f\x31\xdd\x43\x22\x19\x32\x75\x77\x91\x43\x91\xf4\x60\x20\xb6\x83\xd6\x7d\x00\x9a\x1a\x69\x59\x53\x33\x32\x7f\xbc\x76\x15\xbd\x7b\x41\x52\x2b\xdb\xf5\x02\x4d\x0e\x82\xc4\x21\x67\xe6\xfb\x19\xb1\x6d\xe1\xcb\x0d\x5c\xdf\xdc\xc2\xef\x5f\x2e\x6f\x2f\xe0\xb3\xd1\x48\x1e\x0e\xd2\xc1\x80\x84\x56\x7a\x84\xde\xf2\x08\x1f\xe7\x59\x5c\xcb\x11\x97\xe5\x63\x39\x49\x75\x2f\x07\x84\x79\x16\xdf\xf2\xe7\xb2\x94\xa5\x1e\x27\xb6\x1e\xaa\xb2\xd8\x29\x26\x8f\x4f\x7e\x57\x16\x3b\x24\xc5\x9d\xa6\xa1\xfd\xdb\x31\xed\xca\xb2\x90\x07\x67\xf1\x01\x76\x83\xf6\xfb\x70\x27\x14\x8f\xad\x3c\xb8\xf8\x9c\xbb\xee\xfe\x7c\xe0\xb4\xb4\xf8\x10\xd0\xc5\x0a\x9a\x1e\xf9\x1e\xed\x9b\x04\xec\x9c\x47\x63\x5a\x23\xc7\xbb\x4e\x9e\xaf\x47\x62\xbb\x81\x79\x30\x28\x06\x36\x92\x06\xc1\x76\x68\x27\xcb\x9e\xef\x42\xdf\x6e\x48\x52\x24\xc3\xa9\xcb\x79\x3e\x07\x2b\x69\x40\x10\x7f\xa2\x7d\xd4\x2a\x71\x69\xdb\xa3\x12\xda\x81\xdc\xa4\xe8\x40\xe5\x68\xcf\x16\xfc\x1e\xc1\xe5\x14\xe8\xb0\xd7\x84\x1d\x68\x7a\x2d\x94\x88\x75\x2e\x3d\xe0\xd3\xc4\x0e\x1d\x8c\xe8\xf7\xdc\x39\xf0\x0c\x4a\x1a\x03\x93\x65\x85\x5d\xb0\xe8\x40\x3e\x4a\x6d\xe4\x9d\x41\x60\x7a\x5d\xb9\x01\x49\x1d\xec\x25\x75\x06\x5d\x2c\x17\xf7\xa2\xfe\x9a\x86\xb4\x15\xe8\xb8\xe2\x1e\x56\xd5\x5c\xda\xb1\xe8\x26\xa6\xd8\x37\xa1\xb5\x92\x5c\x34\x48\x94\xfe\x79\xc2\x8d\x1e\x79\xb4\xbd\x54\x08\x73\x59\xbc\xd2\xe2\x2a\x41\x5d\x96\x18\x5c\xe9\x54\xab\xab\xe2\x73\x7e\x37\x70\x36\xcf\xe2\x92\xa6\xe0\x6f\x9f\x27\xfc\x6e\xf7\x6e\x59\x1a\x10\x42\x64\x8b\xc5\xcd\xe4\x35\x53\x0d\x55\x3c\x77\x13\xfc\x7f\x0e\xa2\xb5\x6c\xeb\xdc\x15\x29\xf6\xca\xc2\xab\x4d\xf8\xc8\x35\x01\x24\x69\x40\x8f\x93\xc1\x11\xc9\xcb\x58\x35\xb2\xcd\x14\x56\x3e\x6b\x96\xf3\x36\x28\xff\x7f\x64\xbe\x7b\xfe\xca\x07\xb4\xcb\x02\x00\x67\xeb\xf8\x88\xcb\xfc\x7e\x8f\xe8\x1a\x0f\x9b\x5c\xda\x6b\x69\xf4\x3f\xd1\xb2\x15\x40\x03\x8a\xa9\xd7\x43\xb0\xd1\x05\xed\xa3\xbd\xc1\x61\x76\xca\xf2\xa3\xee\xb0\x83\x58\x26\xd0\x28\xad\xdb\x4b\x83\x36\x59\xf2\x62\x5d\x92\x22\xcd\x45\x4c\x4a\xab\x13\x7c\x9f\x39\x58\x50\x7b\xd6\x0a\x45\xd9\x07\x52\x2f\xc0\x2a\x0d\x47\x16\x5f\xd3\x3f\xb1\x72\x69\x40\x5a\x8a\xa2\x68\x1a\x9a\xb7\x00\x02\xa9\x2a\xfe\x04\xe2\x0f\x79\xb8\x42\xe7\xe4\x80\xf5\xea\xc9\x71\x38\xe6\xb2\xb0\xe8\x83\x25\xf8\x90\xe5\x9d\xcb\xe2\xb4\xae\xef\x85\xbd\x00\xd8\x20\x5d\xe3\xa1\xd2\x09\x4a\xb3\xc5\x7e\x73\xdf\x8e\xd3\x5f\xed\xb6\x19\xdb\xbd\x01\x59\xd7\xcd\xda\x30\xbb\x51\x44\x43\x4e\xf5\xcf\x6a\x54\x0a\xce\x32\xce\x1a\x5e\x4d\xad\x7f\x82\x77\x93\x1b\x6f\xa0\x93\xd3\xcb\x93\x77\x3f\x3b\xc2\x51\xa7\x49\x3e\x1b\x96\x5d\x0a\xc1\xc5\x27\xd8\xee\x18\x71\x95\xd9\x54\x16\x1f\xea\xb2\xd0\x7d\x3a\xf1\xcb\x27\x20\x6d\x62\xe2\x51\x61\xd2\x26\x25\x47\x92\x85\x75\xd3\x56\x49\x89\x77\xda\xae\x93\x1a\xa9\x35\xb0\x75\x8e\xd0\x85\x10\x3f\xdc\x84\x83\x8f\xf5\x3f\x9c\x60\x36\x2f\x5b\x91\x37\x5c\xfe\x3a\x7a\x53\x25\x84\x1c\x7c\xfd\xeb\x0f\xf1\xc9\x31\x0e\xbe\x89\x1b\xab\x8b\xd9\xd4\x97\xaf\x7f\x03\x00\x00\xff\xff\x95\x4f\xad\xbb\x8c\x06\x00\x00")

func clientGotmplBytes() ([]byte, error) {
	return bindataRead(
		_clientGotmpl,
		"client.gotmpl",
	)
}

func clientGotmpl() (*asset, error) {
	bytes, err := clientGotmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "client.gotmpl", size: 1676, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _routerGotmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x54\xc1\x6e\xdc\x36\x10\x3d\x8b\x5f\x31\xdd\x43\x2c\x05\x0a\xf7\xee\xc0\xa7\x26\x45\x0d\xd4\x76\x91\xb8\xed\xa1\xe8\x81\x92\x46\x12\x1b\x2d\x29\x0f\xa9\x5d\x6f\x15\xfd\x7b\x31\x22\xa5\xdd\xad\x1d\x14\x05\x7a\x30\x4c\xcc\x0e\xdf\x7b\xf3\xe6\x89\xdb\x2d\x7c\x78\x80\xfb\x87\x47\xf8\xf8\xe1\xf6\xf1\x1a\x3e\xd9\xc1\x23\xc1\x41\x39\x68\xd0\x20\x29\x8f\x15\xd4\x64\x77\x70\x35\x8e\xf2\x5e\xed\x70\x9a\xae\x44\xaf\xca\x2f\xaa\x41\x18\x47\xf9\x73\x38\x4e\x93\x10\x7a\xd7\x5b\xf2\x90\x8a\x64\x53\x5a\xe3\xf1\xd9\x6f\x44\xb2\x41\x53\xda\x4a\x9b\x66\xfb\xa7\xb3\x66\x23\x44\x42\x81\x62\xd3\x68\xdf\x0e\x85\x2c\xed\x6e\x8b\x95\xf3\xd8\x75\xdb\x4e\xed\x8a\x4a\xbd\x0b\x1d\x7c\xb9\xb1\xb6\xe9\x50\x36\xb6\x53\xa6\x91\x96\x9a\x6d\x4f\xd6\xdb\x62\xa8\xb7\x2b\xee\x5c\x09\xe0\x99\x18\x47\x52\xa6\x41\x90\x9f\x91\xf6\xba\x64\x5d\xdb\x2d\xdc\xd6\x70\xb4\x03\x01\xe1\xd3\x80\xce\x83\x3f\xf6\xe8\x40\xef\xfa\x0e\x77\x68\x3c\xf8\x16\x61\xaf\x3a\x5d\x29\x6f\x09\xb4\xf1\x48\xb5\x2a\x91\xeb\x47\x38\xe8\xae\x83\x02\x81\x81\x62\x13\x56\x50\x60\x6d\x09\xa1\x40\x6d\x1a\x68\x95\xa9\xb0\x02\x5b\xd7\xe0\xed\x8c\xc6\x95\x0e\x09\x6a\x4b\xd0\x93\x2d\xd1\x39\x6d\x1a\x09\x82\xa9\x5f\xe5\x1a\x45\xf2\x6b\x44\x4f\x33\x40\x22\x4b\x62\x12\x4c\xfa\x63\xc4\xd2\x6e\x86\xc6\x67\xb6\x19\xab\xb3\xbb\x47\x3b\x80\x6b\xed\xd0\x55\xe7\x43\xd9\xa8\x62\x19\xdb\x81\x36\x8c\x37\x5b\xe1\x82\x3f\x32\x08\x5a\x29\xce\xe5\x8c\xe3\x3b\x88\x6e\xde\xa1\x6f\x6d\x35\x4d\x5c\x8c\x29\x48\xe3\x8e\xe5\xf7\xe1\x7f\x0e\x6f\xc7\x51\xde\x9a\x7e\xf0\x8f\xc7\x1e\xbf\x52\xeb\xa6\x29\x83\x94\xab\x0f\x83\xbf\x28\xe7\x61\xbe\x2c\x70\xa0\x61\xe4\x30\xeb\x92\x3f\x52\x7d\x98\xf6\xea\x22\x13\x57\xf9\xf9\xac\xbd\x72\x2e\xde\x90\x61\x02\x1e\x9a\x61\xc2\x25\xf9\xd9\x2b\xf2\x70\x68\xd1\x80\x36\xda\x6b\xd5\xe9\xbf\x78\x5f\xb3\x03\xa1\x07\xea\xc1\x94\x5e\x5b\x13\x9d\x88\x02\x9c\xa7\xa1\xf4\xec\xc2\xdb\xc0\x2c\xc3\x0f\x51\xe6\x3d\x1e\x62\xe3\x8a\x8b\x0e\x14\xc4\x6c\x1f\x5a\x5d\xb6\x21\x38\x73\xe5\xb4\x81\xcb\x78\x30\x54\x4f\x76\xaf\x2b\xac\x24\xdc\xfa\x70\x65\x70\x73\xf2\x60\xa7\xc8\xb5\x8a\xd7\xb2\xf4\xf0\x75\xfe\xf6\x82\x7d\x6e\x4e\x97\x27\x65\x1c\x27\x42\x0a\x1e\xe5\x24\x2d\x5d\x32\x18\x97\x9b\x9f\x01\x72\x67\x1a\x76\x00\xe9\xef\x7f\x14\x47\x8f\xcb\x4e\x32\x78\x1b\x47\x1b\xd7\x6f\xf5\xfa\x26\x4e\x26\xef\xf1\x90\xc6\xe3\x5d\x40\xfb\x38\x4b\xf9\x4d\xfb\x36\x5d\xf1\xb3\xec\x1b\xe9\x39\xf7\x32\xdd\xac\x59\xda\xe4\x10\xcf\x5f\xbd\xfd\xc9\x1e\x90\xa6\x69\x91\x2f\xd7\xae\xec\x22\x2f\x09\xa1\x1f\xc8\xc0\x9b\xa0\x76\x14\x49\x12\x4e\xd7\x51\x6b\x2e\x12\x4e\xd5\x6b\x3a\x82\x53\xdf\xa4\x0c\xf6\xfc\x0f\xf9\xce\x16\xdb\x96\xef\x6b\x5c\x65\x5f\xfe\xf0\x03\x13\x06\x56\xff\x0c\x2f\x98\x09\x9f\x80\x5f\x38\xf9\x49\x1d\xee\xd0\x39\xd5\x60\x06\xe9\x3f\x2a\x0b\x29\x93\x24\x85\xad\x8e\xbc\xb6\x37\x2f\x35\x8f\x93\x48\x12\x5d\x73\x37\x77\xac\xcf\xa7\xfc\xc5\xc4\x05\xa6\x84\x4f\x39\x30\x44\xf6\x7e\x6e\xfb\xee\x06\x8c\xee\x66\xe0\x45\xbf\xd1\xdd\x4c\x28\x92\x84\xf1\xf6\x8a\xa0\x38\x3d\x1f\xe3\x04\x37\x33\x40\xa0\xda\xe7\x60\xbf\x30\x59\x21\xd3\xf5\xed\xcb\xde\x73\x71\xc6\x3c\xa9\xd9\xcb\xd3\x23\xf8\x92\xfb\x15\xf2\x99\x9d\xff\xc8\xf5\xf9\x82\x12\xf7\xc8\x66\xc6\x39\x4e\x23\xff\xfb\x2c\xfc\x35\xb8\xfc\xa5\x3f\x77\x8b\x3b\xae\xff\x4f\x78\xb1\x1a\x61\x8d\xee\x44\x32\x65\x31\x99\x21\xcb\xa7\xd3\xdf\x01\x00\x00\xff\xff\xf5\x8c\xc9\xec\x90\x07\x00\x00")

func routerGotmplBytes() ([]byte, error) {
	return bindataRead(
		_routerGotmpl,
		"router.gotmpl",
	)
}

func routerGotmpl() (*asset, error) {
	bytes, err := routerGotmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "router.gotmpl", size: 1936, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"client.gotmpl": clientGotmpl,
	"router.gotmpl": routerGotmpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"client.gotmpl": &bintree{clientGotmpl, map[string]*bintree{}},
	"router.gotmpl": &bintree{routerGotmpl, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
