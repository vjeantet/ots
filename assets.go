// Code generated by go-bindata.
// sources:
// frontend/application.coffee
// frontend/application.js
// frontend/index.html
// DO NOT EDIT!

package main

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

var _frontendApplicationCoffee = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x55\x6d\x4f\xdb\x30\x10\xfe\xee\x5f\x71\x6b\x91\x9c\x48\x25\x14\xc6\xf6\xa1\x52\x26\x6d\xec\x55\x02\x34\xd1\xa1\x7d\x98\xd0\xe4\xc4\x97\xc6\x90\xda\x95\xed\x50\x26\xe8\x7f\x9f\x1c\xe7\xb5\x14\x04\xfd\xd2\xf8\x9e\x7b\x71\xee\xb9\x7b\x62\x30\x2d\x35\xfe\x64\xc6\xac\x95\xe6\x10\x83\x2c\x8b\x82\x90\x54\x23\xb3\x38\xc7\x54\xa3\x85\x18\x82\x10\xf6\x3f\x10\x00\xd3\x18\xf6\x02\x3a\xce\x94\x5e\x9e\xf4\xfc\x68\x18\x65\x42\xf2\x80\x5a\xbc\xb3\x4c\x23\xa3\x61\x74\xcb\x8a\x20\x24\x3e\x72\x58\xe7\x8c\xd9\x3c\xd2\x4c\x72\xb5\x0c\xc2\xc8\xaa\xb9\xd5\x42\x2e\x82\xb7\xef\xc3\xc8\x94\x89\xf1\xa7\xa3\xb0\x5f\xf4\x9b\x48\x12\xd4\xc2\xe4\x1f\xbf\xcc\x23\x94\x69\xe0\x91\xc9\x56\xf2\xaa\xde\x5e\xc4\xae\xd9\x1d\x50\xb6\x12\x07\xfe\x65\xe8\x84\x00\x00\x70\x66\xd9\xac\x7a\x6a\x32\xcf\xea\xff\x16\xfe\xf5\x6f\x85\x33\x18\x5d\x1b\x25\x47\x95\xd1\x58\x66\x4b\x73\xa2\x38\x36\x91\x47\xd3\xc3\x26\xcc\xb7\x80\xd7\xc0\xf1\x74\x3a\x03\xa3\x96\x68\x73\x21\x17\xbf\xb5\x92\x8b\x1a\x79\xf7\x24\x72\x3c\x3d\x9e\xb5\x2d\xf6\xbf\x31\x9c\xa9\xf4\x06\x32\xa5\x41\x48\x8b\x3a\x63\x29\x82\x45\x63\x45\x1b\x05\x3b\xeb\x77\xf6\xbf\x82\xcf\x80\x66\x4a\x25\x4c\x53\xd7\x92\x8c\x15\x06\x09\x71\x6f\x78\xae\xec\x57\x55\x4a\xde\xa3\xd6\x31\x2a\x95\xcd\x9c\x99\x86\x91\xc9\xd5\xda\x31\x97\x33\x93\x9f\x2a\xd6\xf7\x74\x26\x88\x61\x2d\x24\x57\xeb\xa8\x50\x29\xb3\x42\xc9\xc8\x99\x09\x80\xc8\x2a\x87\xa8\x40\xb9\xb0\x39\xc4\x31\x4c\xab\xab\x69\xb4\xa5\x96\xa4\x8b\xe7\x98\x2a\x8e\x97\x17\x3f\x4e\xd4\x72\xa5\x24\x4a\x1b\x38\xa4\x62\x6f\xc5\xb4\x35\x10\xfb\x4c\x66\x55\x08\x0b\xf4\x81\xfa\xec\x15\xd6\x4b\x7f\x54\xa5\xaf\x93\x56\xe0\x9f\xe9\x15\xa9\xfb\x30\x9c\x39\x8f\x1e\x5e\xb9\x12\x82\xb7\xf9\xdb\x71\x3b\x0c\xbb\xc9\x19\xb9\xc9\x59\xa0\x3d\x18\xdf\x0b\xbe\x19\x4d\x5e\x33\x1d\x15\xa1\xfd\x3e\xb7\x53\xe3\x46\x20\x57\xeb\xcf\xcc\x32\x42\x84\x14\xf6\x93\x90\xdc\x6c\xd1\xb0\x63\xb1\x12\x21\x39\x50\x53\x26\x4b\x61\xe9\x04\xfa\xeb\xd9\x70\x87\xeb\x79\xbd\x0e\x91\x64\xb7\x09\xd3\xfb\x89\x5b\xaf\x36\x38\x2d\x44\x7a\x43\x27\xd0\x3a\x56\x81\x9e\xc5\xc6\xc7\xf5\x23\xcd\x99\x5c\x20\x9d\x40\x43\x3d\x21\x6d\x48\xef\xa2\x1d\xef\x1a\x33\x88\x87\x67\xcf\x59\x40\xc7\x34\xf4\x64\xd4\xb3\x37\x98\x58\x97\xcc\x35\x69\xa0\x2d\xdf\xeb\xe1\x60\x96\x45\xed\x1c\x7b\xe2\xb7\xe8\x7c\x53\x6b\x55\x37\xf1\x75\xec\x68\x7c\xdf\x9d\x37\x0f\xd5\xa9\x17\xb8\x71\xac\x95\xba\xf0\x9e\xcf\x5c\x7b\x33\x1e\x24\x1a\x91\xba\xd3\x2b\x26\xb1\x38\x6f\x5a\x42\xc3\x28\x17\x1c\x83\xb0\x8f\x7a\xe8\xf2\xe2\xb4\xdb\xa4\x27\x50\xaf\x98\x42\xae\x4a\xeb\xe5\xd2\x5d\xed\x65\xce\x99\x4a\x4b\xf3\xd2\xcc\x06\x0b\x4c\xed\x13\x5a\xec\x35\xbf\x19\xcc\x9d\xbc\x40\x0c\x7d\x52\x5e\xc8\xc8\xb6\x68\x73\x7c\x5e\xb4\x5f\xdf\xdd\x01\xda\xd3\xaf\x81\xbd\x55\xdd\xb5\x53\xdd\x9d\x39\x2f\x90\xf1\xb6\xe4\x63\xca\x06\xf0\xae\xaf\x5c\xf3\x0d\x21\x43\x85\xdf\xda\xec\x47\x17\x69\x64\x96\xec\x79\xa7\x56\x12\xaa\xf2\xcd\x06\x06\x21\xf9\x1f\x00\x00\xff\xff\xd4\x91\x95\x3e\xa8\x07\x00\x00")

func frontendApplicationCoffeeBytes() ([]byte, error) {
	return bindataRead(
		_frontendApplicationCoffee,
		"frontend/application.coffee",
	)
}

func frontendApplicationCoffee() (*asset, error) {
	bytes, err := frontendApplicationCoffeeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "frontend/application.coffee", size: 1960, mode: os.FileMode(436), modTime: time.Unix(1501852583, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _frontendApplicationJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\x5f\x6f\xdb\x36\x10\x7f\xf7\xa7\xb8\x39\x01\x44\xa1\x1a\x63\x67\xd9\x1e\x1c\xe8\x61\xf3\xb6\x76\x40\x57\x0c\xf1\x8a\x3d\x0c\xc5\x40\x89\x94\xc5\x4e\x26\x0d\x92\x8a\x5b\xb4\xfe\xee\x03\xff\x48\xa2\x64\xbb\x59\xb6\xbc\x28\xc7\xfb\xf1\xee\x78\xf7\xbb\x3b\xdf\xdc\xc0\x4b\x26\x98\x22\x86\x51\x28\x3e\xc2\x5a\x56\x15\x63\x9b\x52\xf1\xbd\x81\x25\x5e\xde\xe2\xbb\x19\xaa\x5a\x51\x1a\x2e\x05\x4a\xe1\xd3\x0c\xe0\x91\x28\x28\x15\x23\x86\x6d\x58\xa9\x98\xc9\x80\x12\x43\xde\x48\xf3\xb3\x6c\x05\xcd\xa0\x26\xba\x7e\x2d\x09\xcd\x80\x0b\x6e\x7e\xe0\x82\xea\x0c\x04\x3b\x74\x68\xed\xbe\x6b\x67\x81\x3a\xb1\x55\xec\x37\xa2\xf5\x41\x2a\x2b\xd7\xf2\xf0\x23\x31\x24\x03\x2d\x77\xcc\xd4\x5c\x6c\xff\x50\x52\x6c\xef\x67\x33\x98\x80\x21\x07\xd1\x36\x8d\xd3\xc4\x11\x41\x0e\x93\x98\x7d\xd4\xde\xf3\xbd\x93\x75\x87\xbc\x46\xc9\x55\x25\xd5\x6e\x1d\x19\x48\x52\x5c\x71\x41\x51\x62\xd8\x07\x43\x14\x23\x49\x8a\x1f\x49\x83\xd2\xfe\xee\x38\x8a\x5f\x89\xa9\xb1\x22\x82\xca\x1d\x4a\xb1\x91\x1b\xa3\xb8\xd8\xa2\x6f\xbe\x4b\xb1\x6e\x0b\xed\xa5\xdb\x74\xe2\xf9\x25\x2f\x0a\xa6\xb8\xae\xbf\xff\x69\x83\x99\x28\x91\x1e\x32\x14\xd9\x0f\xd7\xae\x31\x79\x4f\x3e\xa0\x84\xec\xf9\x8d\x7f\x6c\x92\x85\xb7\x81\xab\xc0\xaa\x97\x3a\x1f\xab\xf0\x0d\xc7\xc7\x2c\x42\xff\xfe\x71\xcf\x56\x30\x7f\xaf\xa5\x98\x77\xe7\xda\x10\xd3\xea\xb5\xa4\x2c\xb6\x75\xbb\x58\xae\x26\x45\xeb\x75\x77\x8b\xc5\x6a\x52\xa7\x41\xf9\xed\x97\x94\x77\x8b\xbb\xd5\x69\x95\xfc\x9f\x62\xa6\x55\x62\xec\x13\xc5\x80\xee\x81\x7f\x71\xba\x82\xa4\x92\xb2\x20\x2a\x89\xf4\xc7\x90\x33\xf7\xff\x2c\xfe\x76\x9a\xe0\xa2\x22\x8d\x66\xf6\xe4\xe8\x48\x14\x13\xf9\x1c\x89\xc2\x2d\x4b\x19\x21\x4d\x65\x61\x49\x8a\x2d\x63\x3d\x35\xbc\x95\xae\x01\x2e\xd1\xd0\xea\x33\xe0\x34\x83\x3d\x51\x46\xfb\x80\xec\x21\xe4\x70\xe0\x82\xca\x03\x6e\x64\x49\xec\x3d\x6c\x8f\x3d\x80\x57\x80\xac\x84\x1b\x26\xb6\xa6\x86\x3c\xcf\x61\x31\xe4\xcd\x87\xe6\xa1\xc7\xd8\x22\x65\xa5\xa4\xec\xed\xc3\x2f\x6b\xb9\xdb\x4b\xc1\x84\x71\x66\x42\x1e\x5c\x04\x90\x3b\x30\xd6\xfb\x86\x1b\x94\x7c\x4e\xd2\xc1\xa5\x03\xc4\x3e\x6f\x07\x9f\xc1\x83\x83\xfc\xb9\x78\xd7\xe5\xfc\xa4\x3d\x3c\x60\xf9\x2e\x8e\x8e\xd3\xde\x6b\xdf\x21\xcb\x71\x71\x02\xe3\xe7\x96\xf1\x5b\x66\x6e\xe6\xf0\xc2\x65\xed\xd3\x7f\xa2\xb1\x23\xdc\x68\x52\x45\x0c\xb7\x44\x0d\x73\xe7\x0c\x5d\x7c\x59\xfb\x69\x76\xae\xae\x17\xa6\x48\xe1\xa6\x88\x6e\x8b\x1d\x37\x49\x36\x1a\x52\x5d\x5b\x5b\x2e\x0d\xc3\x11\x0b\xf2\x58\x10\xf5\x75\x61\xa7\x49\x6f\xa0\x6c\x78\xf9\x77\x12\x4d\xd1\x49\xa2\x90\xa7\x4d\x07\xb7\x69\x2d\x6b\x22\xb6\x76\x44\x74\x6c\x8c\x5e\xd2\x9b\x39\xf7\x92\x81\x7a\x8a\x55\x90\x8f\xe5\x8e\x22\x57\x49\xda\x17\xfc\x7c\x2b\x8d\x9a\x37\xf6\x63\x4b\x70\x3a\x94\x5f\xb9\x9e\x68\x55\x13\xcf\xc8\x57\x81\xc1\xc4\x10\xdc\x37\xfc\x40\xcd\x09\xcf\xbe\xca\xfd\x3a\x18\xf8\x39\xb2\x12\x09\x2f\x60\xfe\xd9\x92\x69\x6c\x20\x66\x67\xab\x1a\xc8\x01\x7d\xe1\xed\xa9\xb5\x72\x15\xac\x04\xc3\x43\x45\xf7\x44\xb0\xe6\x4d\x97\xe6\x24\xc5\x35\xa7\x0c\xa5\x13\x80\xd7\xbe\x7d\x78\x3d\x1a\x22\x17\x00\x7e\x21\x71\xb1\x6f\x4d\xd8\x46\xad\x6a\xfe\xfd\x85\x4a\x96\xad\x7e\x86\x03\xcd\x1a\x56\x1a\x34\x26\xda\x85\xf5\xdb\xd7\x3c\xb4\xd0\xd3\xe5\x9e\x6c\xc2\xa8\xc2\xcf\x2e\xef\x74\x91\x52\xf6\xc4\x22\x3d\xfe\x8f\x2a\x4d\x01\xd1\x16\x98\xaa\xfa\xc5\x77\xb0\x8b\xef\x92\xf1\x07\x46\x68\xef\x3e\xe6\x40\xb4\x6a\x4e\x81\x67\x7f\x9c\xe8\x61\x32\x84\x72\x8c\x56\xef\x13\xfb\xec\x24\xdc\xe9\x56\xbb\x9e\xfe\x0e\x8c\x26\xe2\x24\xe8\x6e\xe4\x84\xeb\xe9\xfd\x6c\x76\x4c\x71\x49\x9a\x06\x99\x9a\xeb\xf4\x7e\xf6\x4f\x00\x00\x00\xff\xff\xf1\x16\xe6\x08\x78\x0a\x00\x00")

func frontendApplicationJsBytes() ([]byte, error) {
	return bindataRead(
		_frontendApplicationJs,
		"frontend/application.js",
	)
}

func frontendApplicationJs() (*asset, error) {
	bytes, err := frontendApplicationJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "frontend/application.js", size: 2680, mode: os.FileMode(436), modTime: time.Unix(1501852616, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _frontendIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x59\xeb\x52\xdc\xb8\x12\xfe\x9d\x3c\x45\x63\xaa\x4e\x76\xb3\xd8\x9e\x4b\x80\x40\x3c\x53\xc5\x6d\x09\x2c\x30\xe1\x12\x08\x6c\xed\x0f\xd9\x6e\xdb\x1a\x64\xc9\x48\xf2\x98\xe1\xd4\x79\xf7\x53\xf2\x65\x98\x5b\x36\x6c\x92\x85\x62\xc6\x92\x5b\xdd\xfa\xba\x3f\x75\x4b\xc2\x5b\xd9\x1f\xec\x5d\xdd\x7e\x3a\x80\x44\xa7\xac\xff\xda\x33\x5f\xc0\x08\x8f\x7b\x16\x72\xab\xff\x1a\xc0\x4b\x90\x84\xe6\x01\xc0\x4b\x51\x13\x08\x12\x22\x15\xea\x9e\x95\xeb\xc8\x7e\x6f\x4d\xbf\x4a\xb4\xce\x6c\x7c\xc8\xe9\xa8\x67\x7d\xb1\x3f\xef\xd8\x7b\x22\xcd\x88\xa6\x3e\x43\x0b\x02\xc1\x35\x72\xdd\xb3\x8e\x0e\x7a\x18\xc6\x38\x33\x92\x93\x14\x7b\xd6\x88\x62\x91\x09\xa9\xa7\x84\x0b\x1a\xea\xa4\x17\xe2\x88\x06\x68\x97\x8d\x35\xa0\x9c\x6a\x4a\x98\xad\x02\xc2\xb0\xd7\x6e\x14\xad\xd8\x36\x5c\x25\x08\xc4\x17\x23\x84\x2e\x94\x8a\x35\x89\x15\xbc\x4d\x73\xa5\xdf\x42\x20\x52\x84\x88\x4a\xa5\x81\x72\xd0\x09\x82\xc1\xf6\x01\x08\x1f\x83\xd0\x09\xca\xb2\xdd\xd8\x06\x33\xa8\x1a\xf3\x96\x44\x1a\xe5\x5b\x33\x44\x61\xa5\xd2\xb6\x6b\xab\x9a\x6a\x86\xfd\xc1\xd5\x25\xd8\x30\xe0\x08\x57\x34\x45\xb8\xc4\x40\xa2\x56\x9e\x5b\xbd\x7d\xfd\x3c\xc1\x5d\x21\xb4\xd2\x92\x64\xa5\x86\x57\xaf\x3c\x46\xf9\x3d\x48\x64\x3d\x4b\xe9\x31\x43\x95\x20\x6a\x0b\x12\x89\x51\xcf\x32\xfe\x54\xdb\xae\x1b\x84\x7c\xa8\x9c\x80\x89\x3c\x8c\x18\x91\xe8\x04\x22\x75\xc9\x90\x3c\xba\x8c\xfa\xca\xf5\x8d\xce\x82\xe8\x20\x71\xbb\x4e\xd7\xd9\x74\x23\x46\x34\x1b\x57\xfd\xc6\x96\x93\x52\xee\x04\x4a\x59\xe5\x3c\xaa\x1f\xca\x35\xc6\x92\xea\x71\xcf\x52\x09\xe9\xac\x6f\xd8\xb2\x7d\x43\x87\x37\xee\xe5\xd9\x69\x3c\x28\xee\xd7\x4f\xf6\xef\x37\xcf\x2f\x3e\xca\x0d\x71\x7d\x7b\xeb\xdf\xde\x9c\x16\x6e\xfb\x7e\xf0\x25\xba\x3d\x8e\xe2\x9e\x05\x81\x14\x4a\x09\x49\x63\xca\x7b\x16\xe1\x82\x8f\x53\x91\x2b\x0b\xdc\xda\x2f\x3f\x01\x57\x24\xb8\xb6\x49\x81\x4a\xa4\xe8\xbe\x73\x36\x9d\x96\x1b\xa8\xd9\xee\x97\x21\xc3\x3b\x29\x8f\x83\x62\x3f\x70\xbb\xf9\x7e\xa2\x42\xbd\xd1\x56\x27\x1d\x31\xd8\xbd\xed\x6e\x74\x1e\x4e\xbb\x4c\xf0\x76\x3c\x3e\x78\xbc\x3f\x69\xfd\x3d\xb2\xe7\x38\x7e\xbc\x3a\x3d\x59\x07\x95\xd0\x14\x08\x0f\xe1\x02\x55\x26\x78\xe8\x0c\x15\x44\x42\xc2\xd1\xc1\x7b\x50\x79\x66\xb8\x0c\x22\xaa\x85\x91\x61\x8a\x5c\xab\x72\x40\x8a\x21\x25\xf0\x90\xa3\xa4\x38\xc5\x26\xa3\xfa\x66\xe7\xe2\xec\xe8\xec\x70\x7b\x5a\x69\x28\x50\xf1\x37\x1a\x0a\x21\xef\x81\x46\x30\x16\x39\x98\xd5\x52\xb2\x38\x23\x31\xc2\x88\x12\x88\x28\xc3\x6d\xd7\x9d\x51\xf7\x27\x8d\x80\x69\x38\x3a\x80\xad\xbf\xfa\xb5\x97\x3c\x15\x48\x9a\x69\x50\x32\x78\x71\x30\x4c\x66\x58\x57\x09\x1d\xb9\x5d\x67\xd3\xe9\x3e\xb7\xcb\x10\x0c\x67\x22\xb0\x3c\x0a\xdd\xe3\xb1\x7b\xe8\x5f\x9e\xc8\xb8\x25\xb6\xc6\xeb\x77\xeb\xbc\x9d\x17\xad\x87\xc7\xbb\x83\xbd\x8f\x9b\x7b\x1b\x83\xf3\xec\x7a\x37\x3e\xfb\xfd\x96\xb4\xfe\x86\x5f\x7d\xcf\xad\x26\xff\x23\x58\xe4\xc4\xb1\x6e\xdb\x79\xe7\x74\x26\x1d\x2f\x85\x12\x6f\xd0\x9d\x68\x74\x97\xfd\xc6\xf7\xcf\x3b\x57\xe1\xd5\x85\x7b\x7d\xfd\xc7\x71\xd4\xf5\x0f\xa5\x78\x97\xfb\xeb\xd1\xe8\xe6\xf2\xe6\xfa\x82\x76\xce\x0e\x5e\x0c\xc5\x5b\xf9\x13\x79\x48\xa3\xbf\x4c\xf4\xaa\x9e\x72\xe5\x34\x40\x57\x0d\xdf\x75\x42\x79\x5c\x48\xc1\xe3\x35\x58\xe5\x42\x47\x22\xe7\xe1\x1a\xac\x66\x84\x23\xbb\x40\x12\x56\x99\xa7\xe9\xa9\x5a\x9f\x2f\x4e\xe0\xbf\x13\x4c\x21\x55\x19\x23\xe3\x6d\xe0\x82\xe3\x87\xba\xfb\x7f\xf5\xb7\x13\x09\xa1\x51\x4e\x89\x97\xab\x4d\xd1\x27\xdc\x86\x96\xb3\x85\xe9\x87\xc9\x1b\x8d\x8f\xda\x26\x8c\xc6\x7c\x1b\x02\xe4\x1a\xe5\xf3\xbb\x40\x30\x21\xb7\x61\xb5\x13\x99\xdf\x59\x2b\x9e\x3b\xc1\xe5\xb9\x4d\x75\xf1\x7c\x11\x8e\xcb\x84\xc8\xc9\x08\x02\x46\x94\xea\x59\x9c\x8c\x7c\x22\xa1\xfa\xb2\x43\x8c\x48\xce\xb4\x65\xa4\x5e\x79\x21\x9d\x88\x99\x84\x4d\x28\x47\x69\x47\x2c\xa7\x61\x25\xf0\xaa\x4a\xb7\xd2\x2c\x36\xf3\xa7\x45\x1c\x33\x84\x18\x35\xc4\x52\xe4\x19\x86\xe5\x4a\xf5\x51\x1b\xbc\xa9\xf0\x29\xc3\xc6\x37\x75\x6e\x9e\xb5\x52\xcf\xc2\x4c\x18\x65\x6d\xe3\x95\xe7\xe7\x5a\x0b\x0e\x7a\x9c\x61\xcf\xaa\x1a\xd6\xdc\x88\xda\x72\x20\x18\x23\x99\xc2\xd0\x82\x90\x68\x52\x77\x9b\xd9\x57\xfd\x4d\x37\x91\xb1\xa9\xae\xab\xbe\xb2\xf1\x91\xa4\x19\x43\xbb\x56\xd4\x48\xda\x6d\x0b\x88\xa4\xc4\xc6\xc7\x8c\xf0\x10\xc3\x9e\x15\x11\xa6\xb0\x99\xd4\x2b\x4f\x65\x84\x37\xb3\x50\xd2\x16\x9c\x8d\xad\xfe\x55\x35\x0f\x4e\x46\x34\x26\x9a\x0a\xee\xb9\x46\x6e\xe9\x20\x1a\x08\x6e\xfb\x44\x96\x24\xfd\x37\x84\x3c\xb7\x72\x56\xd3\x24\x73\x4e\xf3\x4d\xe0\x9a\x7a\xb1\x6a\x7d\xb5\xc0\x92\x3a\x50\x6e\x48\x47\xfd\xd7\xcf\x81\xdf\x13\x8c\x61\xa0\xcb\xfc\x68\x18\x65\xca\x91\x5a\x33\x21\x4f\xd5\x5a\x49\x88\xaa\xe2\x37\xc5\xde\x70\xa1\x8c\x08\xe5\xf1\xd2\xf0\x37\xce\x87\xb9\x60\x58\x40\xc3\x9e\xf5\xb7\xc1\x6a\x30\xe6\x6c\x0a\x64\xa3\x67\xea\x51\xd2\x38\xd1\xcf\x41\x64\xb4\xef\x91\x89\x07\x4a\x33\x1c\x8b\x0a\xb9\xd5\xf7\x68\xa3\x2c\x22\x10\x11\x3b\x63\x55\x4a\xa1\x7d\x38\xc3\xa2\x76\x90\xf1\x8f\xe7\x32\x3a\x71\x7a\xce\xa6\xfd\x55\x7a\xca\x75\xe6\x66\xdc\xc0\x9f\x91\x99\x5b\x63\xcd\xee\xc5\xe5\x64\xd4\x87\x3a\x5d\x2d\x5b\x92\xd6\x24\x4f\x4f\xbd\x95\xa2\x98\xf4\xcf\x8f\x63\x76\x1a\xda\xef\xa1\x7e\xc8\x72\x95\xd8\x9d\x29\xe1\x59\x71\xc2\x50\x6a\x28\x3f\xed\x90\xf0\x18\xa5\x05\x52\x98\x55\x55\xf6\xd5\x4e\xab\xd3\xe4\x8c\x16\x80\x79\x07\x3e\xe4\xa8\xcc\xaa\xb0\x03\x2a\x03\xb3\x61\x2d\x97\x58\x42\xc3\x10\x79\xcf\xd2\x32\xc7\xca\xbd\x33\x4a\xae\x12\xaa\x80\x2a\xe0\xa2\xa2\x9a\x2a\xdd\x5e\x96\x67\x22\x11\x98\x10\xf7\x86\x50\x91\x90\xff\x49\x90\x31\x9a\x7d\x98\x46\x52\x91\xf6\x07\xa0\xcd\x96\x84\x7f\x03\xe0\x65\x63\x01\x0a\xb3\x4c\x4a\x3b\x0e\x1c\xbd\x49\x61\x84\x72\x0c\x4a\x48\x39\x36\x3b\xee\xdc\xe0\xa7\xea\xdb\x28\x67\x9a\xcd\xb2\xfd\x0e\x8a\xb4\x0d\x2b\xbe\xe2\xbb\xb2\xf6\x41\xf9\x69\x67\x92\xa6\x44\x8e\x2b\x7f\x95\x5d\x67\xcf\x8b\x68\xd6\x5f\xf3\x1a\xca\x5c\x4f\x17\xfc\x6a\xce\x44\xdd\x59\xc1\x72\x9b\x6f\xf5\xf7\x24\x12\x8d\x40\x80\x63\x51\x53\xc1\x73\x93\xee\x9c\x99\xf9\xa8\x2f\xb3\x6c\xaa\xe1\xa2\x59\x93\xbd\x4a\x1c\xe6\xa1\x32\xb6\x14\xc9\xbc\x4e\x23\x6e\x97\x55\x6f\x89\xa0\xd9\xab\x13\x1f\x99\x21\x69\xcf\x52\xb5\xbe\x4a\x6f\x59\x92\xb6\x3d\xb7\x14\x58\x3a\xd4\x6c\x02\x88\x44\x32\x63\xca\xac\x7e\x29\x98\x21\x6c\xa1\x7a\xd6\x7a\x4d\xd6\x26\x73\xb9\xcd\xa0\x25\xb3\x5e\x74\x4e\xd9\x4d\x79\x96\xeb\xc6\x86\xaf\x39\xf8\x9a\xdb\x2a\x0f\x02\x54\xca\xaa\xab\xaf\xca\xfd\x94\x6a\x0b\x46\x84\xe5\xd8\xb3\xea\x60\x3c\xaf\xca\x95\x45\x87\xba\x66\xba\xdf\x08\xcf\x2c\x47\x97\x86\xab\xa6\xda\x64\x3e\x13\xaa\x4d\xf6\x5f\x3f\x9d\x6a\x75\x7c\x82\x12\x64\xb8\xf2\x53\x69\x96\x2d\x06\xe0\x56\xe4\xb2\x49\x6e\x05\x51\x8d\xdd\xb2\x90\x2a\x2d\x24\x86\x90\x2b\x93\x25\x4c\x12\x80\xcf\x17\x27\xdb\x0b\xae\x5e\xd0\xfa\x62\x86\xd6\xd1\xaf\xa2\x6c\xb8\x63\x7d\x85\x6d\x48\x42\xb3\xd5\x59\x0c\xf3\x12\x52\x2d\x43\xf9\x89\x21\x51\x08\x12\x53\x4c\x7d\x94\x55\x56\x17\x10\x0b\xf3\xd9\x20\x33\xc9\x5d\x2a\x64\x11\x10\x05\x3a\x21\xe6\x50\x96\xb3\x10\x42\x54\x5a\x8a\xf1\x14\xe1\x1c\x38\xce\x95\x86\x8c\x28\x05\xb4\xd4\x64\xf2\xb5\xe0\x08\xc8\x14\xae\x7c\xcb\x43\x3f\xc0\xc4\xc5\xa4\xf7\x7c\x36\xf8\xe9\x54\xbc\xa8\xe4\x4b\xbf\x34\xc8\x1d\xe7\xa7\x32\xf2\xc5\x4c\x79\x61\x3a\x9a\x30\xe5\xeb\xb9\x68\xd1\xd9\xf0\x15\xd6\x78\x26\xee\x3c\xee\xef\x68\xb3\x93\xa4\x82\x6f\x9b\x13\x4e\xd9\x65\xd6\xcd\x1b\x89\x60\x6c\x81\x42\x9c\xac\x10\xc1\x03\x74\x60\x47\x81\x12\x82\x1b\x22\x99\x1d\x83\x44\x26\x48\xf8\x7c\xa4\x9f\xda\x50\x14\x94\x31\xf0\x11\x62\xc3\x1e\x25\x20\x25\x63\xdf\x1c\x29\xb2\xb1\x61\x16\x17\xc5\x92\xc2\x0b\xdf\x45\xaa\xaa\x0d\x93\x6d\x5f\x5d\x6c\x27\x77\x0a\xf3\x12\x52\x14\x30\x39\xb1\xfe\xe3\x1a\x0e\xd5\x79\x73\x26\x94\x9f\x44\x81\x26\xa1\xf8\x63\x98\xec\x7f\x9b\x83\x7d\x4c\x75\x92\xfb\xe5\x71\xfe\x24\x7f\xa2\x11\x4a\x57\x68\x65\xf5\x9b\xc6\xe0\xea\xb2\x3c\x13\x2c\x85\x3b\xd5\x98\x46\x5d\x02\x19\x9e\xe7\x66\x3f\xf3\x0b\x47\x93\xc2\x89\x1c\x97\x27\x82\xc9\x6d\xdd\x1b\x05\xc7\x64\x44\x2e\xab\xbb\x86\x8c\xe5\x31\xe5\xea\xd7\xe7\x8b\x96\xef\xb8\x84\x18\x3e\x18\x8b\x6e\xdb\x69\x77\x9c\x77\x75\x6b\xe9\x05\xc4\xe2\xf5\xc3\x9d\x50\x07\xfe\xc5\x89\x7f\x76\xfe\x74\x92\xf1\x3f\x8e\xee\x0f\x42\xf9\x69\xb4\xc9\x06\xe3\xad\xbd\xce\x66\xf2\xf1\xfc\xb7\x2f\xd9\x7b\xf2\xee\xf4\x71\xe7\xfc\x1f\x5c\x3f\xd8\x36\x1c\xf1\x80\xe5\x21\x02\x61\x0c\x02\x91\x66\x94\x61\xd8\xa0\x85\x5f\x7c\x64\xa2\xf8\x75\x0d\x84\x04\x5a\x0b\x52\x1e\xd2\x11\x0d\x73\xc2\xca\xcb\x27\x65\xa8\xcc\x11\x43\x0c\x7f\xc8\x37\xba\xa0\xe6\x58\x6e\x4f\x6e\x30\xeb\x8b\xcd\xa1\x9a\xbb\xd4\xfc\xb6\xaf\x3e\xaf\xdf\x1d\xe0\x1f\xd1\xe1\xd9\x60\x48\x5a\xad\xcd\xd3\xd3\xfd\xee\xed\xee\x51\x6b\xa7\x3b\xb8\xbc\x1b\x9c\xfb\x78\xb8\xf1\xd4\x89\x3a\xb7\xad\x24\x7f\xff\x72\x5f\x7d\x07\xa4\x98\xfa\x3e\x4a\xaa\x12\x9b\xa0\x72\xdb\x4e\xcb\x69\xcd\xf6\xbd\x10\xcf\xc7\xad\x5d\xf7\xcb\xc5\x3e\x1d\x9e\xd1\xdf\xbb\x24\xbe\x1e\x6a\xfa\xd4\x3e\x3a\xb9\xa6\xdd\x40\x1d\x84\x87\x4f\x5b\xb7\xbb\xf1\xfa\xd5\xf0\x46\xe4\xa7\x2f\xc2\xb3\x08\x88\x64\x19\xa3\x41\x79\x43\x60\xe6\x33\x83\xdd\x73\xab\x7b\x1a\xcf\xad\xfe\x61\xf0\xfa\xff\x01\x00\x00\xff\xff\x7d\xb0\xa0\xdc\x42\x18\x00\x00")

func frontendIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_frontendIndexHtml,
		"frontend/index.html",
	)
}

func frontendIndexHtml() (*asset, error) {
	bytes, err := frontendIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "frontend/index.html", size: 6210, mode: os.FileMode(436), modTime: time.Unix(1501852563, 0)}
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
	"frontend/application.coffee": frontendApplicationCoffee,
	"frontend/application.js": frontendApplicationJs,
	"frontend/index.html": frontendIndexHtml,
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
	"frontend": &bintree{nil, map[string]*bintree{
		"application.coffee": &bintree{frontendApplicationCoffee, map[string]*bintree{}},
		"application.js": &bintree{frontendApplicationJs, map[string]*bintree{}},
		"index.html": &bintree{frontendIndexHtml, map[string]*bintree{}},
	}},
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

