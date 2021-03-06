// Code generated by go-bindata.
// sources:
// tmpl/monitor.tmpl
// tmpl/screenboard.tmpl
// tmpl/timeboard.tmpl
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

var _tmplMonitorTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x54\x4d\x6f\xdb\x30\x0c\xbd\xfb\x57\x10\x3e\x6f\xf9\x05\xed\x61\x68\x33\x34\x87\x25\x58\x11\xa0\x87\x61\x10\x04\x8b\xb1\x85\x2a\x52\x2a\xcb\x31\x02\x8d\xff\x7d\xd0\x87\x1d\x3b\x75\x9a\x13\xf9\xf8\xde\x13\xc5\x50\xb6\xd8\x9a\xce\x56\x08\xa5\xe0\x8e\x0b\x53\xb3\xa3\xd1\xd2\x19\x5b\x42\xe9\x3d\xac\x36\x02\x88\x4a\xf0\x05\x80\xe6\x47\x84\xf9\xef\x31\x91\xb6\xa1\x42\x54\x16\x00\xee\x72\xba\x43\xda\x87\x4a\x22\x79\x0f\xf2\x00\xab\x3d\xaf\x5b\x20\x0a\xaa\x10\xdd\xaa\xfe\x78\x6f\xb9\xae\x31\x11\x89\x4a\xef\x57\x44\xe5\x37\xef\x51\x0b\xa2\xbf\xd1\xe8\x3b\xa0\x16\xc9\xe4\x88\x6d\xcb\x6b\x9c\x9b\x3c\x3c\xac\x77\xfb\x22\x9c\xff\x2b\x97\x89\x8a\x00\x01\x60\x5b\x71\xc5\x9d\x34\x9a\x0d\xd2\xdc\xea\xee\x14\xd0\x76\xb5\x1e\x19\x57\x71\x59\x14\x00\x1f\x1d\xda\x0b\x3c\x86\x8b\xac\x7e\xc7\xf8\x5f\xb4\x3b\xe1\x53\xc3\x2d\xaf\x1c\xda\x78\xb3\xdc\x62\x2f\x5d\x33\xba\xa6\x66\xf3\x08\xb6\xc6\xc9\xc3\x65\x6b\x9e\xb9\xe3\x40\xa4\x63\xca\xb4\x61\xe1\xdf\x48\x37\x88\x03\x9e\xd3\x66\xd7\xce\x4e\xaf\x98\xc4\x1b\xed\xd0\x9e\xb9\x02\x22\x9b\x21\x26\x07\x2c\xb9\x2d\x50\x61\xc9\x32\x9d\xfa\xa3\x13\xd2\x5d\x7b\xe3\x31\x9d\xb6\x35\x10\x96\x2c\xf6\xf2\x88\xa6\x73\x2f\x40\xe4\x52\xc8\x9a\x2c\x9e\x94\x96\x94\x1b\x5d\xa9\x4e\x60\xde\x11\x99\x32\x16\x17\x25\xe9\xa7\x84\xbb\x33\xf9\xe8\xa4\xc5\x9f\x9d\x52\x6f\x52\x0b\xd3\xc7\xa1\x44\x8c\x1d\x3a\xa5\x58\x9f\xd0\x61\x2c\x9f\xd9\x8b\x63\xc1\xfe\xc5\xb4\xee\x19\x15\xbf\x84\xb9\x60\xcf\x1a\xd3\x3a\x26\x22\x90\x27\x33\xe7\x2c\xd9\xac\xcf\x5c\x75\x71\xb9\x06\x16\x8e\xc8\xcc\xeb\x33\x71\x62\x97\x37\x2c\xce\xba\xb1\xd8\x36\x46\x89\x71\xc5\x86\xcd\xbb\xad\xb8\x6b\x1e\x1e\xf6\xd8\xd2\xee\x1d\x88\xcc\x7b\x3e\x37\x66\xb3\xce\x47\xe2\x1b\xb7\x5a\xea\x1a\x88\xfa\x1c\x25\xc9\x15\xff\x52\xf7\x8a\x95\x39\x87\x57\x33\xea\x99\x1d\xa0\x99\xd1\x84\xb8\x6c\xf8\x64\xa5\x93\x55\x5c\xe1\x6a\x08\x93\xc3\xa4\xf2\xb5\x74\x72\xc6\x60\x71\xdb\xcd\x02\x75\xe6\x49\xb7\x9f\xa2\x7b\x49\x41\xc5\xff\x00\x00\x00\xff\xff\x85\x80\x7d\xfd\x71\x05\x00\x00")

func tmplMonitorTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplMonitorTmpl,
		"tmpl/monitor.tmpl",
	)
}

func tmplMonitorTmpl() (*asset, error) {
	bytes, err := tmplMonitorTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/monitor.tmpl", size: 1393, mode: os.FileMode(420), modTime: time.Unix(1530282546, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplScreenboardTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x5a\xdb\x6e\xe3\x38\x12\x7d\xcf\x57\x10\x79\x5e\xf7\x7c\x41\x3f\xa4\x93\xee\x49\x00\x67\x27\x1b\x67\xd2\x7b\xc1\x42\x60\xa4\xb2\x4c\x44\x12\xdd\x14\xe5\xc4\x31\xfc\xef\x0b\xde\x59\x12\x75\x09\x16\x93\x17\x77\xd5\x29\x1e\x92\xc5\x62\x91\x25\xb6\x80\x96\x77\x22\x07\x72\x59\x50\x49\x0b\x5e\x66\x6d\x2e\x00\x9a\x17\x4e\x45\x71\x49\x2e\x4f\x27\xf2\xe5\xae\x20\xe7\xf3\x25\x39\x5d\x10\x22\x99\xac\x80\x7c\x35\xfa\x27\x2d\xac\xce\xe7\xcb\x0b\x42\x4e\x27\xc2\xb6\xe4\xcb\x23\xd0\xe2\x8f\xa6\x3a\x2a\xf5\x05\x21\x02\x68\x91\x71\x25\x7f\x55\x16\x01\xd6\xe8\xe9\xb4\x22\xd0\x14\xd6\xd6\x32\x6c\x76\x54\x80\xd3\xb5\x46\x30\x8d\x2d\x82\x9a\xc6\x2d\x9f\xa0\xde\x57\x54\xc2\x33\x15\x8c\xbe\x54\xd0\x7a\xe2\x15\x79\x63\x72\x37\x6d\x21\x68\x53\x02\xf9\x62\x35\xd2\x5a\x66\x07\x6b\xaa\xa7\x4f\x48\x43\x6b\x50\xbf\xd6\x05\x7f\x57\xa2\x71\x00\x21\x7b\x01\x5b\xf6\xee\xb1\x07\x23\x3a\xb4\x80\x2d\xed\x2a\xe9\xd0\x1b\x2b\x2a\xcf\x5e\x10\x92\x72\xc7\x98\x84\xa6\xf4\x93\x15\x25\xc8\xf1\x89\xbc\x69\xdc\x8e\x5e\x1e\xf7\x61\xf5\xd4\xbf\xdd\xe0\xde\x9d\xf6\x9f\x5e\x75\x74\xaa\x7f\x79\x95\xf3\xb3\x5b\x77\xc3\x39\x0c\x89\x60\x1f\x2d\x12\x6e\xfe\x04\xef\x12\x53\x64\x52\xa9\x62\x1e\x6d\xb3\x80\xeb\xaa\x62\x65\xd3\x23\xa3\x5a\x17\xb3\x19\xab\x05\x74\x1b\xf6\xd1\x9b\x5d\xd6\x2a\x55\x4c\xa6\x6d\xa6\xb9\x6e\x81\x95\xbb\x30\xc7\x9d\x11\x2d\x89\x05\xa7\x19\x7e\xb2\x42\xee\x3c\xc1\x9b\x96\x6c\x7b\x03\x8d\x36\x5f\xd9\xfd\x10\xfb\x38\xf6\x2e\x76\x6c\x14\x66\xbe\xef\x6b\x5e\x71\xe1\x95\xb9\x96\x6c\x6b\x03\xcd\x39\xb2\x86\x88\xd2\xef\xbf\x58\x2d\x95\x60\xe2\x92\x90\x8a\x1d\x20\x6b\xf7\xd4\xaf\xd9\x9a\x1d\x60\xa3\x64\xd7\xd1\x79\x64\xb4\x89\xc1\x3f\xb1\x0a\x6e\x60\x4b\x12\xfd\x1b\x20\x0c\xa1\x82\xac\x80\xad\x1f\x86\x25\x78\x66\x1f\xde\x86\x90\x03\xfb\x70\xa3\x52\x80\x1b\xd0\x60\xee\xc1\xf3\xd7\x5d\x2b\x79\xfd\x67\xc3\x64\xc0\x72\xad\xcb\x3a\xa5\x74\x9e\x8c\xcd\x66\x49\xaf\x3a\xc9\xdb\x9c\xea\xfd\x65\x21\xea\x55\x96\x31\xb6\x99\x25\x54\x61\xe0\x76\x85\x85\x54\x94\xf4\xf6\x4e\x64\x34\xc1\x68\x93\xff\xaf\x0e\x5a\x9f\x8c\xb0\xef\x47\x40\x9c\xad\xd4\x9f\x30\x86\x7e\x4d\x3c\xfd\x3f\x3a\x10\xc7\xc8\x90\x90\x5f\x6e\x94\x06\x0a\x23\x4c\x8c\x31\x44\x87\x4a\x7c\x31\xcb\x68\x56\x9c\xe6\xd1\x7d\x0e\xc8\x7e\x29\x6d\x16\x53\x06\xbb\x14\xef\x6a\x48\x7c\x0f\x52\xb0\x1c\x21\xb5\x51\x59\x46\x6b\xb0\x90\x4e\xad\xe0\x0f\x56\x49\x10\x78\xd6\x6a\xa9\xb7\x46\x1f\xad\xb5\xb5\x5c\xc8\xbd\x66\x35\x93\x08\xa8\xb4\xc6\x6f\xe2\x1a\xc5\xf6\x24\xd7\x55\x59\x0a\x28\xa9\xe4\x78\x9c\x34\xa8\x5d\x90\x07\xcd\x42\xea\x6b\x5e\xef\xa9\x80\x27\x8e\xc0\xdc\x68\x33\xc9\x43\x6a\x73\x76\x4b\x89\x77\x2a\x7c\x07\x41\x90\x6b\x35\x8a\x82\xc8\x72\x21\xf7\x1f\xa2\x00\xf1\x0d\x07\x3c\x57\xba\xec\xc5\x9f\xcd\xce\xe6\x33\x94\x37\x4c\x24\x38\x0b\x26\x10\xa9\xb2\x5a\xc8\xfa\xfd\x5d\x0a\x7a\xcd\x2b\x84\x81\x52\x66\x39\xaf\x1c\xab\xb7\x5a\xc8\x7a\xd7\xe4\x02\x68\x0b\xbf\x73\x8e\x71\x66\x81\xac\x54\x88\x65\x47\xd6\xcb\xf6\xef\x35\x6f\x0a\x26\x19\x6f\x68\xf5\x83\x8b\x9a\xa2\xd4\x14\x67\xae\x84\x61\xcf\x6e\x98\xc4\x54\x74\xf9\x56\xd9\x56\x37\x8b\x32\xda\xc8\x39\xeb\x5a\x8e\x9e\xb6\xe3\x5e\xf3\x8c\x0f\xb4\x02\x29\xa1\x07\xee\xad\xd6\xdd\x4d\xad\x98\xe6\x4d\xd1\x9a\xad\x31\xd8\x9d\x6e\x17\xc5\xfb\x33\x32\x5d\xce\x7f\xd7\x1c\x40\xc8\x1e\x37\x33\x4a\xbf\xc6\x5a\xfa\x84\x2f\x9e\x69\xd5\xf5\x3d\x71\xd0\x3a\x77\x9e\x6b\xe1\x13\x8c\x77\x35\x2d\xe1\xcf\xc7\x75\x7f\xa0\x4a\x9d\x75\xc2\x47\xbb\xb7\x9b\xf5\xc0\x39\xb5\x19\xc8\x4c\xfc\x8e\xec\x99\x8d\x3c\x56\x30\x12\xc5\x43\xac\xd5\x9a\x44\x50\xfe\x45\x21\x34\x48\x93\x33\x47\xef\x34\x1b\xbe\x1d\x9b\xbf\x89\x3b\xf2\xb8\xff\xc9\x6f\xbf\x69\x4f\xcc\x24\xa5\xc4\x32\xd8\x6b\x14\x2d\xcd\xa1\x89\xd2\x82\xa4\xa5\x3d\x5d\x5b\xf2\x95\xfc\xe7\x74\xb2\x39\x22\x58\x9f\xcf\x97\xa7\xd3\x97\xf3\xf9\xf2\x6f\xa7\x13\x34\xc5\xf9\xfc\xdf\xf1\xde\xd4\x20\xed\xbd\x68\xea\x56\x37\xaf\xd0\xf9\xfa\x00\xcd\xc8\x25\x2d\x09\x0d\xb3\x1b\x28\xb3\x28\x72\x46\xef\x61\x6a\xdc\xda\xf8\x62\xcc\xb5\xd3\x9a\xde\xc8\xef\xa9\x78\x55\x0e\x4d\x0d\x3d\x8d\x0d\xc7\x5e\x6b\x3b\x74\xbb\x5c\xfd\x1f\xf7\xc2\x7e\x90\xe8\x5b\x11\x7d\x01\x7c\x1a\x56\x5a\xe3\x6e\x45\x5a\x58\x4a\x36\xcc\x61\x33\x19\x6c\x48\x36\xb6\x0a\x8b\x63\x47\x11\xb8\x5a\xe9\x82\xa4\xea\xae\x55\xb2\x0e\x83\x77\x89\xab\x69\x75\xe7\x44\xc5\xb4\xb3\x18\xad\x46\x8d\x13\x74\x8d\xe4\x94\x71\x09\x85\x8a\xa7\xe1\x30\x4c\xeb\x07\x01\x39\x6b\x19\x0f\x1f\x09\xf6\x5e\x13\x3e\xd6\x58\xc5\x38\x59\xaf\x68\x42\x93\x9a\xa9\x99\xd2\x75\xea\x0f\xde\x60\xff\x6c\x79\x83\xfd\xe3\x2d\x66\x46\x75\x55\x81\x90\x77\x37\x5e\x4d\x95\x9c\x31\x7f\x3d\x72\xf8\xf4\x78\x54\xf5\xf8\x08\x5b\x01\x6d\x48\xac\xaa\xc4\xcc\x84\x55\x46\x55\xa6\xb3\x9b\x66\x5c\x43\x19\x6b\x2b\x23\xba\x7d\x60\xa4\x25\x0c\xc8\x4b\x86\x05\xf9\x29\xb2\x9a\x89\x24\x5c\x44\xea\x6a\x2d\x9d\xbc\xc6\x82\x29\x3e\xfc\xa3\x03\x3f\x3e\xeb\xc7\x9a\x5e\xef\x20\x7f\x0d\x1f\x55\xb4\xe4\xcb\x03\x25\xcc\x11\xfc\x2e\x78\xb7\x67\x4d\xe9\x81\xd2\x29\x2c\x8d\x37\x58\xc4\x84\x69\x10\xc7\x2c\xc1\x13\xcb\x5f\x1f\x78\x1b\x7d\x4a\xc9\x5f\xb3\x3d\x6f\xc3\x67\x32\x83\x2f\xe1\xf9\x5e\x94\x80\x89\x40\x69\x22\x26\x6d\x31\x47\x75\xfb\x74\x1f\x96\x66\x27\x6b\xbf\x36\x1a\x58\x32\x10\x34\x88\xb8\xff\xd9\xd6\xdf\xca\x1c\xdd\xe4\x5f\x4a\x74\x8b\x77\xf0\x1c\x8d\x3e\x7f\x51\xb0\xeb\xbc\x8d\x62\x3d\xd8\xcc\xb1\x6d\xd8\x47\x1c\x2c\xad\x11\x2d\x8d\x05\xe7\x38\xee\xa9\x28\x59\x48\x76\xb5\x11\xdd\x87\x08\x23\xcd\xce\xaa\x39\x84\xf9\x34\x07\x3f\x91\xe6\x30\x3f\x05\x10\x07\x96\x83\xfd\x09\x53\x31\x72\x66\x7f\xfd\x9c\xb0\xf5\x42\x72\xfd\x11\xbf\xcf\xac\x3f\xf4\x63\x5a\xf4\xb1\x7f\xc2\xe7\xf0\x0c\x02\x9d\x36\x6a\xf5\xb2\x83\x55\x06\xf7\x7b\xbb\x39\xce\x35\x3d\xf2\x4e\xf6\x59\x2b\xad\xed\xf3\x62\xdb\xd9\xd5\xed\x5a\xb9\xd9\xf1\xb7\x5b\x16\xdd\xfa\xea\xae\x95\x59\xbb\xe3\x6f\xd9\x4e\xa9\xdd\x5a\xc7\xa6\x4b\x79\xbf\x0b\xc1\x45\x8a\x19\x0c\xd0\xe3\xb6\xe6\x4b\xd9\xd7\x54\x42\x93\x1f\x13\xf4\x95\x45\x7a\xfc\xae\xc1\xd2\x0e\xbe\x09\xa0\xaf\x05\x7f\x6b\x12\x5d\xbc\x78\xac\xd7\x49\x68\xb4\xb4\x9b\x1b\xd6\x4a\xc1\x5e\x3a\x19\x2f\x6f\xe8\xa9\x88\xe1\x5e\x67\xa8\xe9\xd2\xfe\x1e\xed\xa3\xdf\x9a\xb5\x32\xd1\x9f\x7b\x13\xcc\x2a\x85\xf7\x3a\x44\x6d\x67\x6e\x25\x37\xac\xdd\x57\xf4\x68\x3e\x9a\x78\xb0\x30\x5a\xf7\x51\xc4\xbd\x87\x21\xdb\x05\x8f\x13\x0f\x02\xb6\x20\xa0\x89\xb2\x82\xce\xb1\xd9\xde\x03\x2d\xfa\x88\x12\x35\x98\xa6\xbf\x65\x05\xfc\x1b\x04\xbf\xe6\x5d\x5c\x0c\xed\x58\x01\xd9\x07\x08\x9e\xe5\x06\x70\xa7\x0b\x36\x9f\xe6\xbe\xa7\x0d\x2d\x61\x23\xa9\xec\x5a\xe5\x4e\xfc\x96\x56\x6b\x34\x6b\x35\x6c\x16\x03\x3d\xaf\xa5\x5b\x2f\xef\x71\xf8\xf4\x86\x7b\x1c\x3e\xc4\xa5\x5b\x7f\xb2\x47\x74\xa0\xa5\x7a\x8c\x8f\xb7\x74\xeb\x4f\xf6\x88\x6f\xe8\xa9\x2e\xd1\x85\x7d\xa4\xfd\xcc\x55\xf2\x81\x0a\x5a\xb7\x48\x6d\x4a\x51\x0b\xb8\x5a\xc3\x48\xa7\xa8\xb4\xd2\x07\x05\x17\xd1\xeb\x50\xcb\xc3\x77\x2e\x8b\xf4\x5f\x5b\x56\x83\x42\xde\xae\x85\xd5\x8e\xbe\xf0\x4d\x51\xe8\xa0\x8d\x1e\xa9\xb4\xe8\x77\x8d\xc1\xe6\x59\x36\x92\xa2\xc9\x68\xd1\xcd\xc6\x62\x63\x2c\x63\x15\xe4\xc8\x15\x9a\x57\x5d\xdd\x78\xdf\xe6\x56\x0c\xfb\xdc\xa2\x93\x0b\xb7\xe6\x65\x0b\x61\x0f\x54\x46\x74\x07\xa8\x91\x92\xef\x9b\xaa\x0e\x36\x0f\xea\xa9\xff\xbd\x60\xcd\xce\x17\xff\x0b\x00\x00\xff\xff\xd8\x2a\x63\x57\x55\x21\x00\x00")

func tmplScreenboardTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplScreenboardTmpl,
		"tmpl/screenboard.tmpl",
	)
}

func tmplScreenboardTmpl() (*asset, error) {
	bytes, err := tmplScreenboardTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/screenboard.tmpl", size: 8533, mode: os.FileMode(420), modTime: time.Unix(1546534573, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplTimeboardTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x55\xcd\x6e\xdb\x30\x0c\xbe\xe7\x29\x08\xa3\x87\x0d\x68\xf2\x00\x03\x72\x28\x10\xac\x18\xb0\x9f\x2e\x2d\xb2\xc3\x30\x04\x4a\x44\x3b\xc2\x14\xdb\x95\xe5\xb4\xa9\xa0\x77\x1f\x44\xfd\x39\xae\x57\x0c\x58\x4e\xe2\xc7\x8f\xe2\x47\x9a\x8c\x14\x76\x4d\xaf\xf6\x08\x05\x67\x9a\xf1\xa6\xda\x6a\x71\xc4\x5d\xc3\x14\x2f\xa0\x30\x06\x16\x9f\x38\x58\x5b\x80\x99\x01\x68\xa1\x25\x82\xff\x2d\xbd\xf7\x81\x20\x6b\x8b\x19\x00\xc7\x6e\xaf\x44\xab\x45\x53\x47\xf7\x6a\x00\x79\x92\x42\xc6\xb7\x4d\x2d\xcf\x74\x87\xe3\xac\x91\xf1\x6f\x0e\xb0\x76\x06\x60\xcc\x93\xd0\x07\x58\xdc\x2a\xd6\x1e\x3a\x98\x07\x50\xb1\xba\x42\x58\x78\x4e\xe5\x7c\xa4\x28\x6a\x9a\x50\x93\xaf\x5a\x61\x29\x6a\x41\x1a\xfc\x75\x00\x27\xf1\x92\x6b\xd8\x88\x17\xe7\x88\x41\xa2\x84\xc5\x4d\xaf\x9b\x6e\xcf\x24\x3a\x07\x4b\x46\x08\xc8\x5e\x6b\x0b\x63\xe6\x80\x35\x0f\x17\xfb\xf0\x3b\x85\x7b\xd1\x85\x84\x6d\x32\x42\x78\xf6\xfe\x25\xfc\x56\x35\x7d\x4b\xb5\x57\xee\x04\x4b\xf8\x69\xcc\x55\xe5\xd1\x0f\xcb\x48\xb0\x36\xf6\xe5\x4a\xd4\x1c\x9f\xaf\xe1\x0a\x25\x1e\x47\x0c\x51\x06\xb7\xb5\xd7\xc6\x50\xaa\xc2\x18\x62\xd2\x89\x90\x5f\x53\x32\xee\xf7\x4d\x8b\x24\xa3\x73\xa7\x20\xa3\xf3\xa8\x4b\xe2\x09\x6f\xc9\xc8\x8c\xff\x90\xa1\xcf\xfe\x3b\x10\xd6\x91\xe5\xbf\x3d\x40\xcb\x24\x6a\x8d\x17\x23\x49\xfc\xc5\x5d\xf0\xc4\xef\x9a\xb8\xdb\x52\x8a\x76\x92\xfb\xd1\x39\x22\xdf\x8e\xa5\xcc\xc1\x89\xf9\xc2\xd4\x6f\x54\xd4\x14\x07\xf9\x01\xbb\x00\xc3\xa4\x86\xb8\x23\xb9\x92\x5e\x7d\x6e\x31\xef\x8e\x33\xb2\xbe\x13\x93\x7d\x1a\xb1\x0d\x19\xd9\x69\x0c\xa5\xff\xcc\x76\x28\x5d\x1e\x49\x87\x40\xf6\xe8\xab\x61\xca\x25\xbc\x3e\xa4\xa2\x7c\x05\x6b\x7c\xec\xb1\xd3\x93\x25\x28\xef\x4b\x35\x3c\x0e\x7a\xfd\xbd\x47\x75\xce\xcb\x93\x74\x52\x69\x73\x6b\xa9\xde\x51\xb9\x41\x03\x84\xeb\x53\xcc\x4d\x55\x29\xac\x98\x6e\x94\x57\xe1\xc0\x1a\xa1\x28\xe0\xdd\x0a\xd7\x58\xde\x6b\x25\xea\x6a\xc8\x7b\x4f\xdb\x99\xc3\xe2\x7a\x66\x64\xd8\x93\xcb\xda\x53\xda\x34\x5e\xb9\x1b\x97\x13\x37\x9e\xb9\xb4\xe4\x61\xc2\xdc\x8a\x87\x63\x5c\xf0\x3c\x7b\xe3\xa4\x31\xf8\x87\xe0\xfa\xe0\x42\x9f\xe8\x10\x02\x3d\xfa\x46\xd8\x3f\xf4\x35\xc5\xe4\x3a\x27\x1b\x90\x47\xe0\x72\x20\xa6\x2d\x1b\x36\x74\x88\xe5\x86\x3d\xe0\xb1\x95\x4c\xe3\x86\x29\xc1\x76\x12\xd3\x7e\x0c\xfe\xb6\xdd\xfb\x11\x68\xdb\x53\xe0\x85\x9e\xd6\xec\x88\x83\xa1\xfa\xea\xcc\x38\x52\xad\xc2\x52\x3c\xc3\xe0\xbf\xd3\x99\xd1\xcb\xb1\x64\xbd\xd4\xf9\xb9\xf1\xa6\x7b\xb0\x26\x45\xdb\xd9\x9f\x00\x00\x00\xff\xff\x64\x22\x23\x98\xf1\x06\x00\x00")

func tmplTimeboardTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplTimeboardTmpl,
		"tmpl/timeboard.tmpl",
	)
}

func tmplTimeboardTmpl() (*asset, error) {
	bytes, err := tmplTimeboardTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/timeboard.tmpl", size: 1777, mode: os.FileMode(420), modTime: time.Unix(1546514320, 0)}
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
	"tmpl/monitor.tmpl": tmplMonitorTmpl,
	"tmpl/screenboard.tmpl": tmplScreenboardTmpl,
	"tmpl/timeboard.tmpl": tmplTimeboardTmpl,
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
	"tmpl": &bintree{nil, map[string]*bintree{
		"monitor.tmpl": &bintree{tmplMonitorTmpl, map[string]*bintree{}},
		"screenboard.tmpl": &bintree{tmplScreenboardTmpl, map[string]*bintree{}},
		"timeboard.tmpl": &bintree{tmplTimeboardTmpl, map[string]*bintree{}},
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

