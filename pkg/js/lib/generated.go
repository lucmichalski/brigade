// Code generated by go-bindata.
// sources:
// js/event.js
// js/job.js
// js/mock8s.js
// js/run.js
// js/run_mock.js
// js/runner.js
// js/waitgroup.js
// DO NOT EDIT!

package lib

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

var _jsEventJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8f\x41\x4a\xc6\x30\x10\x85\xf7\x39\xc5\xdb\x55\x41\xd2\x13\x74\xe1\x42\x70\xed\x0d\xa6\xc9\xd4\x04\x62\x52\x33\x53\x6b\x91\xde\x5d\x42\x94\x9f\xf2\xef\x86\x79\x8f\xef\x9b\x19\x47\xbc\x7c\x71\xd6\x57\xca\x3e\x71\x85\x67\x71\x35\xce\x2c\xd0\xc0\x48\x51\x14\x65\x01\xb7\x4a\x5b\x91\xe2\xd9\x45\x8f\x28\xa0\x9d\x2a\xa3\x2c\xd6\x2c\x5b\x76\x1a\x4b\xbe\x90\x1e\x1e\xf1\x63\x80\xce\xaf\x47\x47\x20\xfc\x69\xde\x59\xbb\x61\xa5\x4a\x1f\x18\x3c\x29\x0d\x4f\xd8\x43\x74\xa1\xc1\x5b\x34\x17\x7f\x34\x79\x9b\x2b\x7f\x6e\x2c\x6a\x0d\xa0\x21\x8a\x5d\x37\x09\x98\xf0\x6f\x6e\xb2\xf3\x96\xa5\xf4\xd6\xfb\x77\x95\xd3\x18\xfe\x5e\x4b\x55\xb1\x97\xb7\xa7\xcb\xed\xe6\x37\x00\x00\xff\xff\x38\xd9\xc5\x1d\x16\x01\x00\x00")

func jsEventJsBytes() ([]byte, error) {
	return bindataRead(
		_jsEventJs,
		"js/event.js",
	)
}

func jsEventJs() (*asset, error) {
	bytes, err := jsEventJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/event.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jsJobJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x54\xc1\x8e\xdb\x36\x10\xbd\xeb\x2b\x5e\x7d\xb2\x5b\x47\xda\x20\x6d\x0f\x59\xf8\xb0\xe8\x6e\x80\xa4\x40\x10\x24\xdb\xf4\x18\x50\xd4\xc8\x62\x96\x22\x55\x72\x68\xaf\x11\xec\xbf\x17\x43\xda\xb2\xb3\x4d\x2f\x02\x38\xe2\xbc\x99\x79\xef\x0d\x9b\x06\xf7\x03\xa1\xa3\x5e\x25\xcb\x30\xa3\xda\x12\x4c\x44\x64\xaf\x1f\x90\xda\xe4\x38\xe1\xe5\xef\xf5\xd5\xaf\xf8\x05\xa3\x7a\x20\x28\xd7\x61\x6b\xb8\xae\x76\x2a\x40\x69\xd3\xbd\xcd\x39\x1b\x2c\xe4\xf0\xa2\xa4\xbc\xb6\x8a\x29\xf2\xa2\xaa\x9a\x06\x1f\x82\x67\xcf\x87\x89\xd0\xfb\x80\x77\xbe\xad\xab\x3e\x39\xcd\xc6\x3b\x39\x2d\x9d\x1a\x69\x0d\x56\xf1\x21\xae\xf0\xad\x02\x04\x79\x3c\x60\x03\x1e\x4c\xac\x2a\xc0\xf4\x58\xfe\x44\x8f\x93\x0f\x1c\xeb\x2f\xb4\x23\xc7\xe5\x26\xc0\x43\xf0\x7b\x2c\x72\x0c\xce\x33\x7a\x9f\x5c\xb7\xa8\x80\x27\xc9\x6c\x1a\xbc\x57\x23\x61\x6f\xac\x45\x4b\xda\x8f\x04\x1e\x08\x53\xa0\xde\x3c\xe6\x86\xf2\xd1\x77\x8d\xf6\xae\x37\xdb\x51\x4d\x90\x86\x62\x5d\x21\xd7\xaf\xe5\x84\x4d\x0e\x5e\x17\xc4\x7b\x69\x55\x58\x92\x54\x6b\x22\xc3\xf7\xa5\x7f\xb0\x47\x48\xae\x16\x52\x0f\x50\x81\x40\x8f\xa4\x13\x53\x07\xe3\x10\xe9\x9f\x44\x4e\x13\x8c\x8b\xa6\x23\xf8\xbe\xe0\x29\xc4\x81\xac\xc5\xb2\x69\x8d\x6b\xe2\xb0\x9a\x6b\x17\xd0\x4d\x01\xbf\x3e\x0e\x74\x03\xed\xad\xa5\x42\xa0\xef\x73\x67\xcd\x4e\xd9\x44\x98\x94\x09\x51\x62\xe4\x76\x26\x78\x37\x0a\x2b\x3b\x15\x8c\x6a\xed\xc5\x48\xe4\x76\xd8\xe0\xdb\xd3\x09\x51\x2c\x50\xa4\x17\x75\x95\x83\x9f\x04\x5c\x59\xb0\xda\xce\x59\xe6\x28\xf4\x2c\xfa\xb9\x21\x61\xcd\xf7\xb8\x7b\xff\xf9\xcb\xe7\x9b\x8f\x85\xc0\x8c\xf5\x89\x74\x20\x3e\x32\x8a\xfb\xc1\xc4\x22\xc5\xe4\xa7\x24\x1e\xc9\x14\x5e\x34\x5b\x00\x4f\x1d\x63\x6f\x78\xc8\x57\xca\x78\x59\x5b\x61\x52\x42\x31\x43\xd7\xa7\x09\x4e\xc8\x7e\x47\x21\x08\xbb\x0a\xa3\x62\x3d\x18\xb7\x95\x02\xd9\x54\x7d\xf0\xe3\xa9\xa2\xf4\x3c\x8f\x56\xb0\x62\x26\xe5\x38\xd4\xe4\xbb\xec\x1c\x59\x06\x62\xb4\x07\x11\x76\xb9\x5a\xe7\xb1\xb4\x77\xac\x8c\x2b\x0e\xc8\x06\x11\x07\x14\x23\x41\x07\x52\x4c\xdd\x0c\x7e\x44\x3a\xe2\x86\x24\x46\x70\x5d\xcc\x3f\xf1\xd5\xb7\x62\x9a\x3f\x53\x4b\xc1\x11\x5f\x88\x24\x17\x37\x38\x6d\xca\xf2\x6c\x78\x13\xeb\x56\xe9\x87\x6d\x10\x36\x96\xcf\xd6\xe2\x7c\x67\xaf\x0c\x2f\x57\x55\x0e\x04\xe2\x14\x5c\xd9\x27\xa0\xe8\xfe\x0c\xe8\x7f\x6b\x9d\x88\xd8\x64\x02\x24\xb4\xc6\xb3\xa2\xd7\x33\x68\xd3\x40\xea\xfe\xe5\xd8\xd8\x5b\xef\x32\x7d\x03\x85\xb2\xfa\x52\x6c\xaf\x42\x17\xa1\xfd\x38\x29\x36\xad\xb1\x86\x0f\x6b\xb4\x89\xd1\x79\x8a\xb2\xc0\x22\xd8\x51\xd4\xdb\xbb\x0f\x1f\xef\xfe\xb8\xb9\xbf\xbb\x7d\x8d\xbf\xcb\x02\x23\xd0\xe8\x77\xd4\xa1\x4b\x41\x94\xbd\xb1\xd3\xa0\xaa\x8b\x89\xcf\x95\x9f\xcd\xf3\x23\x1e\x2e\x3a\xce\x9f\x88\x24\xe9\x50\x59\xc7\x41\x02\x8b\x4f\x49\x6b\xa2\x8e\xf2\x9b\xd2\x34\xe7\x8c\xe5\x0a\x5a\x39\xe9\x49\x2b\x6b\xa9\x83\x77\x38\xd3\x29\xe7\xf6\x2b\x69\x8e\xf5\x7f\xf3\x4c\x84\x4a\xec\x47\xc5\x46\x72\x0f\x27\x84\xf6\x30\x4b\x7f\x99\xf5\xb6\xcf\xfd\x94\xe6\x23\x16\x6f\x94\xb1\xd4\x2d\xd6\xc5\x41\xf9\x05\x94\x6d\x03\x3d\x6a\xca\x8b\x5b\x7f\x9f\x97\x5c\xcc\xfc\x8f\x3e\xc8\xc2\x29\x87\x97\xbf\x61\x34\x2e\x31\x45\x2c\x5f\x5d\x5d\xe1\x67\xbc\x7a\x11\x49\xfb\xbc\x5d\x4c\x61\xa7\x6c\x5c\xad\x67\x68\xb0\x19\xc9\x27\xfe\xbe\xc2\xcc\xf9\x8f\xac\x23\xf1\x37\x3e\xc8\xfb\x3e\x1e\x56\xc5\x1f\x4f\x55\x75\x32\xce\x3b\xdf\x62\x23\xaf\x7f\xf5\x6f\x00\x00\x00\xff\xff\xdd\x8d\x70\x61\x82\x06\x00\x00")

func jsJobJsBytes() ([]byte, error) {
	return bindataRead(
		_jsJobJs,
		"js/job.js",
	)
}

func jsJobJs() (*asset, error) {
	bytes, err := jsJobJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/job.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jsMock8sJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\x4f\x8b\xdb\x30\x14\xc4\xef\xfe\x14\x83\xe9\x21\x01\xe3\xd0\xab\xc1\xa7\x42\x6f\x2d\xa5\x81\x5e\x4a\x29\x8a\x34\x4e\xb4\xb1\x25\xa3\x3f\xbb\x0b\x8b\xbf\xfb\x22\xc7\x56\x92\xbd\xe4\xb2\x17\xe5\x45\x9a\xf7\xe3\xcd\xf8\xed\x76\xd8\x07\xeb\x88\xc1\xca\x33\x46\xab\xa0\xd8\x69\xa3\x83\xb6\xc6\xe3\x44\x47\x08\x8f\xd2\x88\x81\x25\x1a\x7c\x71\xf4\x36\x3a\xc9\x22\xe9\x7f\x59\xe5\xd1\xe2\xad\x98\x8a\x22\x83\xa4\x35\x9d\x3e\x0e\x62\x7c\xd8\xfe\x6d\x56\xfe\x48\xca\x05\x72\x8e\x07\x3a\xc3\xc0\xcb\x0d\xf0\xa2\xc3\xe9\xe7\xbe\x41\x17\x8d\x4c\x23\x6d\x8c\xdf\xce\x0f\x80\x63\x88\xce\xe0\x02\x72\x2c\x80\x29\x21\xd6\xff\x0b\x40\x5a\xc7\x3f\x5f\x9b\xa5\x67\xb4\x6a\x2d\x81\x23\xc3\x2d\x58\x0c\xdc\xe6\xb7\x8c\xff\x5f\x77\xda\xa8\xcd\xea\xb6\xba\x36\xb0\xbf\xd3\xe7\x0e\xf6\xac\x07\x06\xa1\x44\x10\x75\xa2\xa2\x6d\x91\x7e\xb3\x74\xda\x2e\xe5\x54\x2d\x85\x74\x14\x81\x37\xd3\x28\x76\xb7\xf0\x14\x6e\x94\x92\x54\x54\x15\x7e\x47\x63\xb4\x39\x56\x10\x46\xe1\xbb\xd0\x3d\x15\x84\x23\xbc\x1d\x88\x67\xd1\x6b\x95\xce\x48\x5f\x67\x80\x62\x57\xfb\x20\x42\x9c\x83\xc5\x78\x12\x9e\x0d\xca\x0c\x2d\x31\x65\xed\xea\xf5\x6f\x6a\xca\x4e\x7a\x71\x60\xef\xeb\x27\x7b\x48\x5e\xfe\xa1\x4d\xcc\x8f\x69\x5d\xaf\x2e\xb8\x74\xce\x26\xf9\x1a\x68\x7c\xda\xa9\x35\xff\xbc\x26\xd7\x0f\xf2\x28\x85\xfb\x9d\xf9\xac\xf1\x8a\xa9\x78\x0f\x00\x00\xff\xff\x1e\x5a\xf4\x09\x05\x03\x00\x00")

func jsMock8sJsBytes() ([]byte, error) {
	return bindataRead(
		_jsMock8sJs,
		"js/mock8s.js",
	)
}

func jsMock8sJs() (*asset, error) {
	bytes, err := jsMock8sJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/mock8s.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jsRunJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x57\xdd\x6f\xdb\x38\x12\x7f\xf7\x5f\x31\x20\x8a\x5d\x19\x75\xe8\x2e\xf6\x65\xe1\x9c\x1f\xb2\x69\xda\xba\x6d\x3e\x2e\xc9\x16\x77\x28\x82\x80\x96\xc6\x16\x63\x8a\x14\x48\xca\xa9\xaf\xc8\xff\x7e\x18\x92\x96\xe4\x24\x4e\x16\xfb\xd2\xa8\x9c\x99\xdf\x7c\x7f\x78\x3c\x86\xeb\x12\x41\x6a\x8f\x56\x0b\xe5\x60\x61\x2c\xd8\x46\x6b\xa9\x97\xe0\x85\x5b\x39\x0e\xd7\xa5\x74\x50\x35\xce\xc3\x1c\x41\x19\x51\x60\x01\x73\x5c\x18\x8b\x20\xf4\x06\xcc\x02\x7c\x89\x83\xf1\x18\xcc\xfc\x0e\x73\xef\xc0\x97\xc2\x43\xe3\x90\x80\xb2\x21\x1f\x8c\xc7\x44\x3d\x52\x0a\xbe\x34\x73\xb4\x1a\x3d\x3a\x38\xba\x98\x41\x2e\x94\x72\xe0\x4a\xd3\xa8\x22\x82\xe7\x42\xc9\xff\x61\x01\x25\x5a\xe4\x70\xee\x4b\xb4\x50\x99\xa2\x51\xd8\xf2\x69\xe3\x09\x8e\x64\xe1\xd7\x55\x0b\xf8\x2b\x14\xd2\x62\xee\xd5\x86\x0f\x88\x7e\x2f\xa4\x07\x2f\x56\xe8\x40\xc0\x9d\x99\x1f\x28\xb9\xc2\x64\x22\x08\x5d\x04\x06\x07\x8d\xf6\x52\x81\xf4\xe0\x9a\x3c\x47\x2c\x1c\x18\x0b\x0b\x21\x95\xe3\x83\x45\xa3\x73\x2f\x8d\x0e\xac\x1f\x8c\xfd\x6c\xe6\xd9\x9d\x99\x0f\xe1\xe7\x00\x60\x2d\x2c\x54\x1b\x98\x12\xf6\x60\x00\x21\x70\x19\x3d\x4a\x98\xc2\xbb\x43\x90\xf0\x2f\xf8\xfd\x1d\x7d\xbc\x7d\x1b\x25\x00\x72\xa3\x9d\x51\xc8\x95\x59\x66\x2c\x2f\x31\x5f\x51\x98\x9d\x17\xbe\x71\x14\x47\x06\x6f\xa1\xda\xf0\xda\x14\x67\xa2\xc2\x61\x90\x21\xc8\x15\x4c\xa1\xf3\x94\xdf\x4b\x5f\x9e\x5d\x65\xac\xc0\x85\x68\x94\x67\x1d\x63\xb5\xa9\x4d\x41\xcc\x3c\x37\x16\xbf\xfd\x46\x50\x7c\x89\x3e\xeb\xa3\x3e\x31\xe5\xf3\xd5\xf9\x19\x77\xde\x4a\xbd\x94\x8b\x4d\x16\x40\x86\xc3\xa7\x16\x5f\x98\xe2\x91\x89\xf0\x16\x18\x48\x07\x52\x07\x2f\x30\x91\x49\x6b\xf4\x8a\xd7\xa5\x70\x5b\x9d\x72\x01\xd9\x53\x22\x4c\xa7\xc0\x3e\x08\xa9\xb0\x60\xdb\x40\x01\xf8\xd2\x9a\x7b\xd8\xa7\x72\x11\xd8\xc1\x1b\x2a\x31\xfa\x93\x9b\xaa\x56\x48\xc9\x62\x87\x01\xe1\xe1\x35\x8d\x57\x31\xdf\x3b\x4a\x2d\xfa\xc6\x6a\xf0\xb6\xc1\x1e\xc8\x78\x0c\x57\x0a\xb1\x0e\x29\x16\x50\xe0\x42\x6a\x2c\x40\x54\xa6\xd1\x3e\x94\xbf\xac\x90\x07\x56\x47\x7c\xd9\xef\x14\x3b\x92\x4d\x4e\x10\xbd\x00\xd3\xf8\x50\x48\x94\x72\x42\xaa\x9f\x77\x2d\xfa\xc4\x06\x0f\xa1\x8a\xc9\x3d\xdb\xe8\x54\xc3\xd1\x02\x0d\xb8\x46\xed\x21\xc3\x21\x87\x99\x87\xc2\xa0\xa3\xae\x88\x25\x4f\x2c\xbe\xc4\xc0\xde\x05\x06\xa9\x0b\x63\x27\x4b\x02\x53\xe6\xfe\x40\xe1\x1a\x15\xd4\x56\x56\xd2\xcb\x35\xf6\xea\x9d\xfa\xf6\xce\xcc\x47\x80\x31\x36\xe3\x31\xbc\xf9\x7c\xfe\xe7\xc1\x9b\xeb\xd9\xe9\xc9\xc1\x9b\x8f\xb3\xeb\xab\x4f\x47\xa9\x07\x56\x7f\xb8\x60\x7b\x68\x04\xae\x93\x1b\x07\xe4\xd9\x7b\xe1\x91\x6b\x73\x9f\x0d\xdb\x27\xe4\xb9\xa9\x2a\xe9\xb9\x6b\xe6\xb1\xe2\xb2\x77\x23\xf8\x63\x78\x98\xd0\xf2\x2a\x81\x25\xd8\xf4\x4c\x13\x09\xed\xb3\xa4\xbc\x82\x29\x68\xbc\x3f\x3e\xcd\xa2\xf0\x70\x47\x26\x12\x2f\xc3\xf7\x85\x29\xb2\x0e\x69\x14\x0c\x96\x95\x58\xc6\xfa\x8c\x14\x5e\xa1\x17\x85\xf0\x82\x2b\x31\x47\xe5\xf8\x9d\x99\xeb\x5d\xff\xf6\xf3\xce\x51\x19\xbd\x74\xde\xc0\x14\x90\x5b\xac\x4d\x10\xa0\x2f\x25\x72\xcc\xd8\x98\x8d\x28\x10\xc3\xfd\x10\x31\x3c\x41\x3e\x7e\x0e\x62\xfc\x8f\x8a\x02\x50\xaf\xc9\x33\xc7\x93\x87\xa8\xd7\xdf\x84\x75\x30\x85\xef\x37\x89\xed\x96\xa3\xc8\x4b\x4a\x1e\x47\xbd\x1e\xc1\x36\xa5\xd9\x5a\xa8\x11\xac\x70\x33\x02\x15\x73\xfa\x2a\xe7\xb6\x2d\x92\x16\x5e\x37\xae\xcc\x7e\x92\x3f\x93\x08\xb4\x16\xaa\xc1\x09\xfd\x79\x08\xf9\x8b\xff\x26\x5b\x1d\xe6\x16\xbd\x03\xe1\x76\xcc\xee\x69\x4d\x1c\x7b\x34\xb7\x9d\x67\x2a\x6c\xc1\xee\x11\xf0\x47\xad\x64\x2e\xbd\xda\xc0\x5c\x99\x7c\xc5\xdb\x3e\xbf\xe5\xb9\xd1\x5e\x48\xed\xb2\xef\x2c\x4a\xb0\x9b\x60\xe5\xf0\x71\x87\xa7\xe6\xde\xea\xf8\x88\xbe\x55\xb1\xb0\xa6\x0a\xfd\xb3\x94\x6b\xd4\xf4\x9c\x6a\xed\x71\x20\x12\x60\x17\x8f\xf4\x10\xa2\xf2\xc1\x9a\x6a\xd2\x2a\x85\x84\xfe\x05\x37\x97\xb8\x98\x40\x0a\x22\xf2\xda\x1a\x5a\x47\xb3\x22\xb8\x1d\x43\x99\x64\xe2\xdf\x2e\xb0\x83\xe8\x64\xaa\x2a\xe7\xca\x2f\xfb\x32\xb4\x63\x18\x3b\x3a\x9e\xbd\xbf\xbd\x3c\xb9\x38\xbf\xfd\x72\xf2\x5f\xb6\x63\xe4\x04\x76\xd0\x92\xc2\x41\x8a\x4c\xca\xa3\x37\x75\x1a\x14\x6d\x1a\xe9\x62\x70\x18\x2f\x02\xb3\x46\x6b\x65\x11\x2f\x01\xe1\x3d\x56\xb5\xa7\xb1\xe3\xd0\x87\x28\x06\x45\x2e\xc2\xd1\xb3\xa9\xd0\x97\x34\x01\x51\xb9\x30\x33\x77\x6d\xdf\x5a\x7d\xfc\xf5\xfc\xec\xe4\xf6\xaf\xcb\xaf\x6c\xf4\xc8\xd6\x5c\x19\x8d\x7f\x5d\x7e\x8d\x96\x3e\x2f\x7d\x75\xf5\xe9\x59\x59\xe7\xca\x57\x24\x3f\xce\xae\x9f\x95\x5c\x4a\xff\x8a\xe4\xa7\x93\xa3\xf7\xb7\xc7\xe7\xa7\xa7\xb3\xeb\xdb\xd9\xfb\x3e\x40\xea\xe8\x87\x5e\xd3\xbb\x1a\xf3\x6d\xb5\xa2\x75\xdf\xdf\xdd\x50\x0f\x52\xd3\x47\xec\x5e\xfc\x73\xa3\x17\x72\x09\x95\xa8\x61\x6d\x54\xd3\x9f\x3e\x01\x25\x3e\x86\x19\x10\x12\xb8\x35\x28\xce\xc2\x51\x92\x3f\x15\x75\x5b\x76\x69\xc4\x3e\xb4\xcb\x6d\xd4\x0a\x31\x59\x58\x27\xd8\x28\xd5\xeb\x04\x7e\xa6\xaf\xb3\x40\x4e\xcd\x10\x45\x6f\x0e\x5f\xf4\x27\xda\x75\x4a\x2b\x72\xbf\x71\x61\x83\x5e\x08\x5f\x4e\x80\x8d\x4b\x63\x56\x63\x1a\x85\xec\x61\xf4\x82\x69\x4f\x65\x9c\x2b\xd9\x08\x2c\x8a\xe2\x5c\xab\xcd\x24\xac\xef\x64\x61\x8c\xe3\x67\x23\x75\xa8\xc6\x70\xd0\x52\x1d\x56\x62\x85\x20\x68\x33\xd0\x82\xac\x84\x2e\x26\x69\xa6\xd2\x26\xa9\x8a\x34\xed\xe3\x01\x7c\x67\xa4\xce\x18\xfc\xf2\x0b\xb0\xb0\x28\xf2\x8a\x93\x99\xdf\x59\x25\xa4\xe6\xae\x64\x37\x69\x03\x55\xc5\x60\xf0\x37\x0f\xb6\xc1\xa3\xd3\xea\xd8\xa2\x08\x97\x41\xcc\x17\xa5\x9b\xd6\x65\x5e\x75\x0b\x42\xa7\xdd\xf6\xc2\xe5\x96\x57\xe1\x6c\x5b\x71\xfc\xe1\x51\x3b\x69\x34\x2d\x94\x04\xc8\x73\xd2\x81\xc4\xb4\x57\xfb\xf6\x26\xd9\x9f\xd7\xbf\x61\x45\x14\x4e\x96\xf4\xee\xd0\xa4\x3f\x91\x1f\xdb\x90\x7e\x74\x70\xce\x63\x78\xd2\x2d\xd6\xed\xeb\x43\x3a\x89\xda\x13\x65\x67\xa9\xd7\xa6\xd0\xa1\xa0\x44\x2e\x8b\x59\xd8\xe8\x61\x34\x26\x8c\x38\x12\xd9\x4a\xea\x82\x4d\xc2\x4d\x99\x06\x21\x13\xb5\xfc\x86\x96\x02\x45\x84\xf5\x6f\xdb\xf7\x6d\xd0\x59\x37\xc4\x19\x69\x60\x13\xd8\xea\xea\x3f\xbb\x5a\xe4\x44\x6b\x13\xdc\x52\xe3\x56\x67\xfd\x5d\xc0\x4a\xb4\xd2\x8b\x65\x10\xf8\x77\x63\x56\x2b\xd1\xf2\x93\x6a\xa1\xc5\x12\x8b\x3f\x37\x44\x26\x7f\xd8\xee\x46\x48\x16\x52\x6a\xfa\xd6\x59\x74\x5e\x58\x7f\x61\x94\xcc\x83\xe8\x19\xae\xd1\x76\x86\x74\x59\x64\x93\xd4\x8d\xa1\x23\xdb\xaf\xce\xc3\xa0\x95\x6e\xd0\x51\x9f\x18\x0e\x25\x36\xe9\x42\xbc\x43\x4d\x5d\xb4\x83\x1d\x08\xa1\x45\xa9\x49\x7a\xcf\x37\x7d\xd1\xf1\x18\x3e\xcc\xfe\x73\x7a\x32\x81\xe3\x52\xe8\x25\x52\x7b\xb2\xd9\xe2\xcc\xf8\x0b\x8b\x0e\xb5\x67\x4f\xac\xb8\x68\x94\xea\x1c\x3d\x52\xf7\x62\xe3\x3a\xae\xed\x06\xbd\x69\x4f\xf8\x87\x50\x3a\x3b\xb5\x73\x7c\x9a\x85\x4a\x7e\xa1\x4c\x8e\xb7\x93\xf3\x1f\x14\x4b\x17\xcc\x7e\xad\xbc\x56\x2d\xcf\xd7\xcb\xeb\x35\x93\x6a\x62\x5b\x1a\x4f\x6d\xd9\x0e\xaa\x09\x30\xcc\x4b\x03\x25\x2a\x65\x68\xa0\x85\xff\x2d\x8d\x29\xe6\x1b\x64\x1d\x48\x0c\x18\xfe\xa8\x8d\xf5\x8e\xd3\x6f\x90\x29\x35\x62\xfb\xd2\xfd\x1e\x86\x69\xef\xc7\xf1\xe0\xff\x01\x00\x00\xff\xff\x4b\x73\x3a\xfc\x4e\x10\x00\x00")

func jsRunJsBytes() ([]byte, error) {
	return bindataRead(
		_jsRunJs,
		"js/run.js",
	)
}

func jsRunJs() (*asset, error) {
	bytes, err := jsRunJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/run.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jsRun_mockJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xce\xcf\x4a\x03\x41\x0c\xc7\xf1\xfb\x3c\xc5\x8f\x9e\x76\x69\x4d\x3d\x0a\xd2\x9b\x08\x7a\xf6\x05\x66\x87\xb8\x9d\xad\x9b\x2c\x49\x86\x0a\xe2\xbb\xcb\x1e\xec\x9f\x6b\xf2\x81\xdf\x37\xed\xf7\xf8\x38\x56\xc7\x92\xcb\x29\x8f\x8c\x59\xcb\xc9\x11\x47\x86\x35\xa1\xc9\xf1\xd9\xa4\x44\x55\x71\xc2\x5b\xa0\x64\x81\xf1\xf2\x95\xcb\x15\xa8\x21\xd8\xa3\xca\x48\x29\xfd\x73\x9c\x73\x8d\x57\xb5\x77\x1d\xba\x49\x87\x1e\x3f\x09\x30\x8e\x66\x82\xb0\xc6\xe9\xf7\xc6\x5a\x93\x15\xed\xc0\x77\x6e\xd2\x81\x24\xcf\x8c\x2d\x36\x0f\x1b\x6c\xf1\x92\x83\x49\xf4\xdc\xf5\x97\x13\x53\xd1\x79\xae\x41\xde\x06\x0f\xab\x32\x76\x8f\x3b\x3c\xf5\xcf\xeb\x00\x7f\x2f\x6a\xe1\x74\x6d\xc1\xe1\x26\xec\xf2\xb7\x26\x38\xac\x15\xe9\x2f\x00\x00\xff\xff\x05\x8c\x6a\x5e\x10\x01\x00\x00")

func jsRun_mockJsBytes() ([]byte, error) {
	return bindataRead(
		_jsRun_mockJs,
		"js/run_mock.js",
	)
}

func jsRun_mockJs() (*asset, error) {
	bytes, err := jsRun_mockJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/run_mock.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jsRunnerJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\xc1\x8a\x14\x31\x10\x86\xef\x79\x8a\xdf\x3e\x75\xe3\xd2\xfb\x04\x73\x10\x5d\x51\xf0\xe8\x4d\x44\x42\x52\xdd\x1d\x08\x55\x6d\x55\x7a\x7b\x17\xd9\x77\x97\x74\xdc\x71\x66\xf0\xe0\x31\xf9\xc9\x97\xbf\x3e\xea\xfe\x1e\x5f\x97\x64\x48\x86\xb2\x10\x74\x63\x26\xc5\xae\x7e\x5d\x13\xcf\xb0\xa0\x69\x2d\xa3\x0b\xc2\x26\x99\xc6\x2c\x73\xdf\x7d\x11\x1f\x6b\xf8\xee\xfd\xe7\x0f\x08\xa2\xd4\x0d\xce\x4d\x1b\x87\x92\x84\x31\x25\xa5\x87\x47\xe2\xd2\xd3\x80\x5f\x0e\x78\xf4\x0a\xaa\x17\x9f\x3c\xc7\x4c\x8a\x13\x98\x76\x3c\x5c\x5c\xf5\x83\x73\xc0\x6b\x97\x59\xc8\x90\xf8\x28\x64\x41\x56\xaa\x87\x7d\x49\x61\xc1\xd4\x0f\xb5\x2a\x3d\x51\xd8\x0a\x45\x98\xa0\x2c\xbe\x60\x27\x04\xcf\xf0\x21\x90\x59\x43\xd5\xd7\xa2\x69\x4e\xec\x73\xfb\x1f\x7b\x2a\x8b\x6c\x05\x4a\x3f\xb7\xa4\x75\x84\x5e\xf4\x0e\x4a\x3e\xe7\xe7\x3b\xf8\x9c\x65\x4f\x3c\x0f\xc7\xdb\xcd\x48\x51\xa4\xc1\x56\x6f\x4d\x50\x03\x79\x95\x8d\xe3\xe8\x00\x7a\x5a\x45\x8b\x8d\x3f\x5a\x70\x02\xd5\x49\xd2\x84\xfe\x8d\xd2\x9c\xac\x90\x1e\x83\x5a\x73\x01\x5c\x99\x64\xf9\x03\x5c\x9a\x07\x43\xa4\x29\x31\xc5\xee\x10\x02\x28\x95\x4d\xd9\x01\x2f\xf5\x7c\x4d\xec\x2f\xa5\x0e\xee\x06\x7d\x84\x86\x2c\x3e\x52\x1c\xf1\xb1\xcd\xdb\xe1\x2d\x68\x2c\xcf\x2b\x0d\xe7\x9e\x97\x98\x6f\x2d\xfc\xfe\x5f\x6d\xcf\x7d\x28\x62\x12\xbd\x85\xdf\xb6\xaf\x7b\x30\x55\x43\xff\xf8\xaf\xe6\x53\x4f\x83\x7b\x71\xee\x55\xe9\x79\x91\x70\xfa\xbb\x54\xee\x77\x00\x00\x00\xff\xff\x0b\x8b\xea\xff\xb1\x02\x00\x00")

func jsRunnerJsBytes() ([]byte, error) {
	return bindataRead(
		_jsRunnerJs,
		"js/runner.js",
	)
}

func jsRunnerJs() (*asset, error) {
	bytes, err := jsRunnerJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/runner.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jsWaitgroupJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x92\xb1\x8e\x9c\x30\x10\x86\x7b\x3f\xc5\x5f\x82\x14\x99\x27\xd8\x32\x8a\x92\x22\x4d\x22\xa5\x88\x52\x18\x18\x62\x73\xc6\x83\xec\xe1\xd8\xd5\x89\x77\x3f\xd9\xbb\x2c\x48\x57\x5e\x61\xad\xd6\x78\xbe\xff\x9b\xb1\x9b\x06\x7f\x8c\x93\x6f\x91\x97\x19\xab\x71\x92\x30\x70\xc4\xb4\x78\x71\xb3\x27\x8c\xdc\x26\x08\x63\x70\xc1\x25\xab\xf1\x5d\xb0\x3a\xef\x21\x36\xf2\x0a\x13\x40\x31\x72\x54\x4d\x03\x93\x90\x98\x43\xfe\x35\xb9\x0c\x91\x66\x8e\x92\x9e\x87\xb4\x1a\x96\xd0\x89\xe3\x70\x44\x56\x35\xde\x14\x20\xd6\x25\x5d\xa2\x2e\xf8\xfb\x4f\x29\x20\x03\xfb\x3e\xaf\x8c\x0b\xb4\x16\xa4\x30\xc4\x52\xf1\xfc\x9f\xcb\xf7\xd2\x7c\xf4\x82\x1d\x5f\x8d\xdc\xde\xb9\x27\xb2\x9e\x97\x64\xcb\x17\x05\x6c\x8f\x88\xb8\x84\xbc\x12\xe8\x95\xe2\xad\x44\xb8\x50\x22\x0a\xfe\x0b\x4c\xe8\xf3\xdf\x70\x1a\x8d\x58\x9a\xb2\x48\xc7\xd3\xec\x49\x48\xef\x12\x19\x76\x92\xf8\x68\x30\x70\xfc\x6a\x3a\x5b\x3d\xc7\x50\x8d\xfb\x21\x60\xd4\xad\xe9\x5e\x72\x6c\xe8\xab\xba\x6c\x6e\xf5\x01\xc8\xf9\xd5\x59\x3d\x6f\x3c\xac\x96\x20\xce\xdf\x6f\xca\x44\x3a\xc4\xf0\x93\x85\x20\xd6\x48\x61\xa0\x67\x4a\x08\x2c\xa5\xed\xdc\x64\xb1\xc2\x6f\x4b\xb7\x3b\x73\x5a\x92\xa0\x25\x24\x31\x51\xa8\x07\x5d\x85\x62\x30\xde\xdf\x34\xaa\x5f\x44\xc7\xbd\x95\x66\x39\xe2\x07\xb7\x27\xef\x5a\x9d\x6c\x3f\x35\x8b\x67\xbb\x8f\x29\x6c\x6a\x53\x8a\xae\xe5\x41\xe9\xe3\xc1\x5e\x0e\x23\xf5\x1e\x00\x00\xff\xff\x60\x6f\x0e\xe7\xca\x02\x00\x00")

func jsWaitgroupJsBytes() ([]byte, error) {
	return bindataRead(
		_jsWaitgroupJs,
		"js/waitgroup.js",
	)
}

func jsWaitgroupJs() (*asset, error) {
	bytes, err := jsWaitgroupJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/waitgroup.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"js/event.js": jsEventJs,
	"js/job.js": jsJobJs,
	"js/mock8s.js": jsMock8sJs,
	"js/run.js": jsRunJs,
	"js/run_mock.js": jsRun_mockJs,
	"js/runner.js": jsRunnerJs,
	"js/waitgroup.js": jsWaitgroupJs,
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
	"js": &bintree{nil, map[string]*bintree{
		"event.js": &bintree{jsEventJs, map[string]*bintree{}},
		"job.js": &bintree{jsJobJs, map[string]*bintree{}},
		"mock8s.js": &bintree{jsMock8sJs, map[string]*bintree{}},
		"run.js": &bintree{jsRunJs, map[string]*bintree{}},
		"run_mock.js": &bintree{jsRun_mockJs, map[string]*bintree{}},
		"runner.js": &bintree{jsRunnerJs, map[string]*bintree{}},
		"waitgroup.js": &bintree{jsWaitgroupJs, map[string]*bintree{}},
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

