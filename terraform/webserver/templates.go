// Code generated by go-bindata.
// sources:
// templates/db_mysql.tf
// templates/devbox.tf
// templates/load_balancer.tf
// templates/network_security_group.tf
// templates/output.tf
// templates/provider.tf
// templates/resource_group.tf
// templates/storage_account.tf
// templates/subnet.tf
// templates/vars.tf
// templates/vmss.tf
// templates/vnet.tf
// templates/web_server.tf
// DO NOT EDIT!

package webserver

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

var _templatesDb_mysqlTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x53\xcd\x8e\x9b\x30\x10\xbe\xf3\x14\x23\xab\xa7\x1e\xa2\xfe\x45\x3d\xe5\xd6\x6a\xaf\xfb\x06\xd6\x80\x27\xc8\x5d\x33\xa6\x33\x36\x55\x1a\xf1\xee\x15\x06\x42\x92\x45\xe5\xf8\xcd\xf7\x27\x3c\x33\xa0\x78\xac\x03\x81\x21\x6e\x3d\x93\x1d\x48\xd4\x47\x36\x70\xad\x00\x1c\x9d\x31\x87\x04\x27\x30\xc7\xc3\x77\x53\x8d\x55\xb5\x09\x9a\x28\xa4\x06\xae\x0f\xa0\xa6\x28\xd8\xd2\x33\x5c\x63\xf3\x96\x7b\x2b\x94\x88\x93\x8f\x6c\x1d\x5e\xf4\x39\xe3\xeb\xf1\x31\xc0\x61\xc2\x1a\x95\x6c\x56\x12\xc6\xee\x9d\xeb\x8d\xd0\xa3\xea\x9f\x28\x6e\x26\x08\x69\xcc\xd2\x10\x18\xfc\x9b\x85\xa4\xb3\xdd\x45\x7f\x07\xab\x24\x03\x89\x01\xe3\xea\x19\x99\x0b\x4c\xce\xb0\xfb\x9d\xc0\x7c\xb8\x0e\x28\x07\xe2\xc1\x4e\xb4\xd1\x54\x00\xab\xbd\x6d\x25\xe6\xde\xde\xc9\x0b\x7f\xcd\x7c\xa4\x1d\x7e\x79\x76\x31\x3f\xa3\x37\xd3\x10\x1b\x9c\x7e\xcc\x7f\x4a\xac\x94\xd1\x54\x15\x80\xbe\xe5\xd2\x7e\xe9\x7f\x02\xf3\xf2\x6a\x5f\x88\x8f\x76\x66\x97\xd7\x29\xd6\x00\x0d\xf6\xd8\xf8\x74\xd9\xbc\xee\xa7\xc9\x93\x14\x3d\x31\x09\x86\xd7\x2c\x7d\x54\x9a\x67\x67\xec\x7c\xb8\x2c\xd3\xe3\x84\x8d\x25\x7c\x7e\x65\xdb\x4b\x3c\xfb\x40\x4b\x91\x15\xed\xea\x2d\x68\xc1\xe0\x23\x7c\xfe\xf4\xe5\xdb\x92\xb8\xbb\x0d\x9b\x66\x77\xbc\x48\x5b\x8a\x56\xc8\x65\x76\xc8\xc9\xce\xcc\x49\xf9\xc3\xeb\xb4\x13\xee\x56\x11\x5d\xe7\xd9\x6b\x12\x4c\x51\x6c\x88\xad\xe7\x2d\xe1\xdd\x66\x15\xf7\x1d\xc9\x6d\xb3\x76\xb4\xeb\xa8\x68\x97\xb3\xb9\x5f\x99\xfb\x73\x2a\x1c\xd5\x60\x89\xcf\x51\x1a\xea\x88\xcb\x55\xfd\xe4\xa5\xf5\x58\xfd\x0b\x00\x00\xff\xff\x21\x2a\x15\x7f\x8a\x03\x00\x00")

func templatesDb_mysqlTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesDb_mysqlTf,
		"templates/db_mysql.tf",
	)
}

func templatesDb_mysqlTf() (*asset, error) {
	bytes, err := templatesDb_mysqlTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/db_mysql.tf", size: 906, mode: os.FileMode(480), modTime: time.Unix(1541042574, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesDevboxTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x55\x4b\x6f\xe3\x36\x10\xbe\xeb\x57\x4c\xd5\x5e\x63\x34\x69\xd0\x9b\x4f\x29\x0a\x14\x48\xb1\x8b\x78\x1f\x87\xc5\x82\x18\x8b\x23\x79\xd6\x12\x29\x0c\x49\xe5\x05\xff\xf7\x05\x45\x49\x6b\xc7\x56\xb2\x0f\xc4\x07\x43\x10\xbf\x79\x7d\xfc\xbe\x91\x90\xb3\x41\x0a\x82\x1c\x1f\x82\x90\x34\xaa\x0d\xeb\x9a\x0b\xc5\x6d\x0e\xf9\x17\x36\xda\x06\xa5\xa9\x5b\xdb\xbb\xfd\x93\xc7\x0c\xc0\x60\x43\x30\xf7\x5b\xce\x07\x67\x00\xb5\x2d\xd0\xb3\x35\x73\xa1\x7f\x3c\x8e\xdd\x8c\xfd\xa9\x4a\x6c\x68\x17\x43\xce\x27\x6f\xc7\x74\xbb\x98\xfb\xf0\x4c\x1d\x76\xf9\xe3\xb9\x63\x7c\x9f\x77\x1a\x40\xa1\xd6\x42\xce\x29\xac\xa7\x39\x96\x90\x3b\x8f\x9e\x8b\x3c\xdb\x65\xd9\x31\xa9\x86\xfc\xad\x95\xad\x62\xe3\x49\x4a\x2c\xe8\x88\x5c\xc3\xc5\x3c\xad\x7d\xdb\x1d\xca\x82\x4c\xd7\x4f\xb4\x3b\x4b\x61\x67\x31\x6c\x86\xd0\x57\xe0\xf1\xa7\xe9\xcb\x00\xb8\x55\x85\x35\x25\x57\x41\x52\xab\x71\xd8\x17\x54\xf4\xfc\xe0\x49\x4b\x00\x2e\xac\x0d\x79\xc5\xfa\x65\x39\x25\xe8\xc2\xa1\x19\x1f\x59\xef\x52\x96\x56\xb8\x43\x4f\xf3\x17\xac\xef\x0d\x36\x89\xee\x53\x6a\xd8\x2f\x7f\x50\x73\x82\x2e\x66\x1c\x31\xf6\xb0\x3b\xad\x9d\x8e\xc5\x07\xac\x55\x83\xc5\x86\xcd\x91\x72\x9e\x33\xe3\x09\xfa\x52\xec\xc0\xe2\xbc\x17\x5f\xc5\x84\xbf\xe0\xbe\x23\x03\x29\xd6\x0e\x96\xf0\x69\x2f\xe3\x11\x66\x71\x64\xb1\x9e\xea\xcf\x19\x40\xd7\x28\xc7\x0f\x4f\x29\x5b\x42\xbe\xf2\x68\x34\x8a\x56\xff\xac\x2e\x54\x77\x91\x67\x51\xbb\xbf\xc3\x95\x6d\x1a\x32\x1e\xfc\x86\x1d\xd4\x6c\x08\xbc\x85\x2d\x51\x0b\x7e\x43\xf0\x66\x05\x9a\xdd\x16\xb0\xf4\x24\xa0\xa9\x26\xcf\xa6\xea\x8f\x3e\xfc\x9f\x41\x7a\x43\xca\x3a\x15\x61\xca\x1a\xe5\x49\x1a\x36\xa3\xb6\xbc\x04\xfa\x9e\x3a\x1a\x3d\xf6\x95\xdc\x4b\xa5\x22\xb2\x2f\xe6\xe6\xab\x39\x6f\x05\x2b\x52\xdc\xc4\x7f\xa1\x92\x84\x4c\x41\x83\x31\x7b\x79\xba\x0d\x49\x64\xe5\x0a\x8d\x35\x5c\x60\x9d\xd4\x6f\xcb\x92\x64\x64\xec\xfd\x3a\x18\x1f\x56\x24\x1d\xc9\x60\xc8\x6d\xf8\x46\xe8\xf9\xdf\x8b\x3f\x2f\xcf\xae\xdf\xad\xd2\x59\x47\xe2\x92\xda\x96\x90\xd7\xe8\xc9\xf9\xa4\xfd\xbd\x86\x06\x9e\x66\x37\xc4\xf4\x71\x49\x17\x1b\xb1\x29\x79\xd1\x5b\xa4\x3a\x84\xde\x10\xea\x8f\xc2\x9e\x06\x8c\x50\xb4\xb9\x6d\x27\xd1\x2f\x21\xff\x57\x6c\xf3\x5f\xa4\x21\x61\x1a\x34\x58\x91\x4e\xb7\xe5\xef\x5b\x3a\x50\xc6\xf5\xcd\x6a\x6a\xd9\x3a\xd5\x8a\x2d\xb9\x1e\x69\x2b\x6c\xd3\x06\x4f\x32\x68\x3e\xee\x8d\xc9\x6a\x00\xa8\x1b\x36\x2a\x38\x92\x71\x9f\x46\x02\xfa\xb7\xfb\x80\x16\x9d\xbb\xb5\xa2\x23\xe0\xed\xf0\x7c\x7e\xf1\xd7\xe5\x6f\x27\xea\xaa\x9a\x4d\xb8\x1b\x76\xeb\xd0\x84\x66\x87\xeb\x9a\xa6\x3c\x0a\x83\xdf\x90\xf1\x3c\x6d\xb3\x12\x6b\x47\x63\x32\x8f\x95\x1b\x22\xc9\x74\x2c\xd6\xf4\x1a\x4c\x1f\xb5\x8a\x4d\x35\x6e\xa7\xaf\x01\x00\x00\xff\xff\xb4\xf4\xad\xbe\x2e\x08\x00\x00")

func templatesDevboxTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesDevboxTf,
		"templates/devbox.tf",
	)
}

func templatesDevboxTf() (*asset, error) {
	bytes, err := templatesDevboxTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/devbox.tf", size: 2094, mode: os.FileMode(480), modTime: time.Unix(1541042574, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesLoad_balancerTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x94\x3f\x8f\xd4\x30\x10\xc5\xfb\x7c\x0a\xcb\xa2\xcd\x42\x49\xb3\x05\x27\x21\x81\x44\xb1\x42\xf4\x96\x63\xfb\xf6\x4c\x7c\x1e\xe3\x3f\x5b\xb0\xca\x77\x47\x76\xe2\xdc\x65\xe3\x39\x09\x0e\x0a\xb6\xdc\x79\x79\x9e\xf9\x3d\x7b\xbc\x0a\x90\xbc\x50\x84\xf2\x9f\xc9\x2b\xff\xc8\x5c\x1a\x8c\x16\x4c\x3b\x4a\xe8\x77\x6d\x25\x24\x66\x80\x4b\x36\x70\xc3\xad\x50\xbe\x54\xae\x1d\x21\x96\x3f\x2a\x82\xfd\x8e\x84\xbe\xb9\x5e\xb8\x3f\x28\x7b\x61\x59\x39\xf5\xd9\xa5\xaf\x2e\xbd\x76\xb4\x23\xc4\x80\xe0\x51\x83\xc5\x3d\x6a\x5b\xb5\x51\x76\xf6\x90\xdc\x61\xe9\xec\xe6\xdf\x6a\x37\x65\xef\x6d\x8d\x6d\xdb\xfd\x7d\xef\x32\x44\xf6\x5d\x01\x31\x2e\xa5\x57\x21\x30\x6e\xd6\x39\x8e\x84\x86\xc8\xa3\x16\x59\x19\xc6\x84\x02\x7a\x22\xb4\xa5\x1b\xc6\x34\xd1\x6e\xea\xba\x7d\x32\x66\x40\x22\xc1\xf3\x68\xc5\x30\x60\xe0\xff\x01\xef\x57\x60\x6e\xc1\x7b\x91\x19\x21\xf7\x1e\x6c\x54\x56\xe6\x6c\x04\xd8\x7b\x7d\x4e\x7e\x9e\x32\xf3\x41\x6e\xec\x91\xd0\x53\x49\xf4\xf3\xe9\xc3\x9c\x27\x2d\xe2\x7d\xcc\x5a\x6e\xe7\x59\x15\x07\xe4\x9d\x1c\xb4\x2c\x7d\x4d\x58\x9e\x6c\xe0\x62\xcc\x0d\xd7\x23\x1c\x80\xc1\xde\x5d\x5b\x7b\xfd\xfb\xe4\xf3\xb1\x4f\x53\xc8\xfd\xdd\x30\x43\x73\xe2\x3a\x2e\x72\x11\xef\xb8\x18\x3f\x5a\xb9\x40\x3e\xe5\xee\x51\x2e\xce\xc3\xa0\x30\x10\x4b\xf1\xff\x99\xfc\x21\x46\xd7\x07\xe5\x2f\xda\x9e\xfb\xb9\xfb\xbc\x47\x3c\x44\x10\x60\x6e\xc4\x9f\x62\x74\xf3\x73\xfa\x91\x54\x88\xcc\xf1\xf8\xf0\xac\xfc\xb6\x7c\x0a\x3e\xee\xcf\x79\xff\x0e\xe5\xe9\x93\x29\x38\xa1\x81\x73\xae\x61\x34\x5f\xb5\x1f\x5e\x04\x8b\x39\xff\x01\xe3\xad\xd5\x97\xbb\xaf\x79\x24\x04\xf1\x56\xfb\x4d\xb8\xcd\xe6\x68\x91\x5d\xd8\x12\x52\x5f\x60\x5b\x54\x55\xe8\x16\x5a\xaf\x67\x63\xdf\xb4\x1e\xf7\x33\x50\x37\x80\x9a\xbb\xa0\xbd\x85\x9a\xca\x05\x65\xb9\x8a\x8d\x38\x90\x43\x8b\xbc\x7d\xca\x5c\x2a\xb6\xd3\xaf\x00\x00\x00\xff\xff\xc6\xaa\xb2\x46\x53\x08\x00\x00")

func templatesLoad_balancerTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesLoad_balancerTf,
		"templates/load_balancer.tf",
	)
}

func templatesLoad_balancerTf() (*asset, error) {
	bytes, err := templatesLoad_balancerTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/load_balancer.tf", size: 2131, mode: os.FileMode(480), modTime: time.Unix(1541042574, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesNetwork_security_groupTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x93\x3f\x4f\xc3\x30\x10\xc5\x77\x7f\x8a\x93\xc5\x84\xd4\x28\x6c\x2c\x1d\xd8\x60\x67\xb7\x8c\x7d\x44\x86\xd4\x17\x9d\xed\x16\xa8\xf2\xdd\x91\xf3\x07\xda\x34\x89\xaa\x66\xcc\xfd\xf2\xee\xdd\xbb\x0b\x63\xa0\xc4\x06\x41\xea\x9f\xc4\xc8\x3b\xe5\x31\x1e\x88\x3f\x55\x40\x93\xd8\xc5\x6f\x55\x31\xa5\x46\x82\xfc\x70\xde\x52\x5a\xac\x1f\x05\x80\xd7\x3b\x84\xc9\xb3\x05\x79\x77\xdc\x6b\x2e\xd0\xef\x55\x06\xda\x8d\x0f\x95\x14\x00\x35\x19\x1d\x1d\xf9\x0b\x7a\xb4\x32\x9a\xeb\x5b\x14\x83\x81\xc9\xdb\x51\xa5\xcd\x92\xe7\xb5\xae\xdd\x0d\x92\x9d\x4b\x29\x5a\x21\xae\x88\x87\x53\x8d\x2b\xe9\xf4\xe5\xa5\x70\x4e\xc6\x7e\xaa\x6b\x3a\x6c\x9e\x63\x6c\xf2\x20\x0d\x3b\xca\xdf\xcf\xc3\x0f\x65\x29\x00\xac\x63\x34\xd3\x00\xff\x05\x5f\xfc\x1b\x25\x6f\xb3\x9a\x36\x06\x43\x58\x6f\xdd\x77\xa5\x48\x86\xea\x05\xee\xd5\x74\xde\x86\xac\x1a\xe2\xa8\x58\xfb\x0a\xcf\xa9\xfb\xcc\x58\x0c\xd1\xf9\x6e\x31\x17\xe0\x16\xe4\x63\x79\x22\xa4\xad\x65\x0c\x41\x35\x8c\xef\xee\x6b\x45\x68\x0a\x8e\xcc\xdc\xd6\x6f\x3f\xa8\x61\xfb\x00\xf3\x97\x3e\x73\x54\xf3\x60\xb1\xfa\xc3\xfc\x1d\xd9\x6f\x00\x00\x00\xff\xff\x26\xe1\x0c\xa2\x82\x03\x00\x00")

func templatesNetwork_security_groupTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesNetwork_security_groupTf,
		"templates/network_security_group.tf",
	)
}

func templatesNetwork_security_groupTf() (*asset, error) {
	bytes, err := templatesNetwork_security_groupTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/network_security_group.tf", size: 898, mode: os.FileMode(480), modTime: time.Unix(1541042574, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesOutputTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\xcd\x41\xaa\xc2\x40\x0c\xc6\xf1\xfd\x9c\x22\x0c\x6f\xdd\x1b\xbc\xb3\x84\xa0\x41\x46\x3b\x49\x89\x49\x04\x4b\xef\x2e\x0a\x15\x5b\xc1\xed\xc7\x3f\xbf\x68\xf8\x14\x0e\x95\x25\x9b\xa9\x74\x16\xaf\x30\x17\x80\xa4\x31\x18\xfe\xa1\xfe\xcd\x49\x36\xd0\x3d\x8c\xf1\xa3\x5a\x6a\x59\x4a\x59\xcf\x8d\xaf\x1a\x76\x60\x3c\x99\xc6\x84\x42\x9d\xbf\x98\x17\x61\x1d\xb7\xed\x70\x6e\x72\xd4\xd8\xaf\x4f\x61\xfb\x22\x85\xfd\x37\x9c\xcd\x3c\x68\x44\x61\xbf\xa9\x5d\x56\x79\x3f\xbf\xe9\x47\x00\x00\x00\xff\xff\x3e\xf4\xe9\x4b\xfc\x00\x00\x00")

func templatesOutputTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesOutputTf,
		"templates/output.tf",
	)
}

func templatesOutputTf() (*asset, error) {
	bytes, err := templatesOutputTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/output.tf", size: 252, mode: os.FileMode(480), modTime: time.Unix(1541042574, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesProviderTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8e\xc1\x0a\xc3\x20\x0c\x40\xef\x7e\x45\x90\x1d\x36\x18\xfb\x83\x7d\x4b\x71\x36\x87\xb0\x36\x4a\x8c\x1e\x56\xfc\xf7\x51\x0b\xb5\x14\x36\x8f\xbe\xf7\x92\x44\x09\x85\x46\x14\xb0\xee\x93\x05\x65\xb6\xb0\x18\x00\x3f\x11\xb2\x0e\x34\xc2\xf6\x9e\x60\x2f\x4b\x71\xf2\x68\xd6\xb0\xe3\x6a\xbb\x9c\xd0\x0b\xea\x4f\x79\xc3\x2d\x48\xf9\x95\xbc\x50\x54\x0a\xbc\xee\x38\x05\x27\xdc\x12\x45\x76\x7f\x0e\xda\x71\x93\x91\x0b\x49\xe0\x19\x59\xbb\x3c\x85\xf0\xce\xf1\xba\x36\x07\x9e\xee\xd0\xa7\x1c\xfe\x6f\xd5\x9a\x6a\xbe\x01\x00\x00\xff\xff\xcd\xed\xb3\xce\x1e\x01\x00\x00")

func templatesProviderTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesProviderTf,
		"templates/provider.tf",
	)
}

func templatesProviderTf() (*asset, error) {
	bytes, err := templatesProviderTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/provider.tf", size: 286, mode: os.FileMode(480), modTime: time.Unix(1541042574, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesResource_groupTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x4a\x2d\xce\x2f\x2d\x4a\x4e\x55\x50\x4a\xac\x2a\x2d\x4a\x2d\xca\x8d\x87\x89\xc4\xa7\x17\xe5\x97\x16\x28\x29\x28\x65\x65\xe6\xa5\xe4\x97\x62\x88\x57\x73\x29\x28\xe4\x25\xe6\xa6\x2a\x80\x80\xad\x82\x92\x4a\x75\x59\x62\x91\x5e\x6a\x5e\x59\x3c\x48\xb4\x56\x17\xa2\x4d\x89\x4b\x41\x21\x27\x3f\x39\xb1\x24\x33\x3f\x0f\xa1\x0a\x26\x52\xab\xc4\x55\xcb\x05\x08\x00\x00\xff\xff\x0a\xde\x17\x64\x83\x00\x00\x00")

func templatesResource_groupTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesResource_groupTf,
		"templates/resource_group.tf",
	)
}

func templatesResource_groupTf() (*asset, error) {
	bytes, err := templatesResource_groupTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/resource_group.tf", size: 131, mode: os.FileMode(480), modTime: time.Unix(1541042574, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesStorage_accountTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\xd0\x41\x4e\xeb\x30\x10\x06\xe0\x7d\x4e\x31\xb2\xde\xe2\x21\x95\xde\xa0\x37\x60\x45\x0f\x30\x9a\x38\x93\x62\x70\xc7\xd6\x64\x5c\x29\x54\xbd\x3b\x72\x8a\x45\xa0\x45\x78\x69\x7f\xfe\xfd\x8f\x95\xa7\x54\xd4\x33\x38\x7a\x2f\xca\x7a\xc4\xc9\x92\xd2\x81\x91\xbc\x4f\x45\xcc\x81\x1b\x78\xa4\x12\xed\xf6\xe4\xdc\x01\x08\x1d\x19\xee\xae\x1d\xb8\x7f\x67\xe5\x1c\xc9\xf3\xff\x13\xe9\x96\xe5\x84\x95\x6f\xc0\x3d\xba\x0d\x38\xf7\x70\x71\x1d\x40\xab\x80\x07\x4d\x25\xe3\x2a\x70\x49\x68\xbd\xbe\xb3\xed\x6b\x90\x21\x95\x9f\xbb\xf5\xf2\x12\x1a\x93\x27\x0b\x49\x7e\xa9\x55\xeb\x34\xb2\xf8\xcf\x99\xf0\x2d\xc8\x70\xe3\xf7\xd7\xc1\xd7\xce\x02\xeb\x1d\x47\x32\x90\x0e\x6b\x58\x3f\x20\x5c\x1f\x42\x9b\x33\x2f\xf0\xe9\x79\x5f\x0d\x0b\xf5\x91\xb1\x8f\xa9\x47\x16\xaf\x73\x6e\x95\x77\x60\x5a\xf8\x8b\x8c\x21\xf2\x1f\xe4\xc5\x2c\x4f\x68\x4a\xe3\x18\x3c\x26\x89\x73\x23\x97\xee\x23\x00\x00\xff\xff\x78\x75\x10\x85\xe6\x01\x00\x00")

func templatesStorage_accountTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesStorage_accountTf,
		"templates/storage_account.tf",
	)
}

func templatesStorage_accountTf() (*asset, error) {
	bytes, err := templatesStorage_accountTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/storage_account.tf", size: 486, mode: os.FileMode(480), modTime: time.Unix(1541042574, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSubnetTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x91\x4d\x4e\xc4\x30\x0c\x85\xf7\x39\x85\x15\xcd\x02\xa4\xa1\x9a\x25\x9b\x39\x09\x42\x51\x48\xcc\xc8\xc0\x24\x95\x9d\xf0\x57\xf5\xee\xa8\x3f\x69\xd5\x68\xba\xa0\xbb\xda\xcf\x9f\xe3\xf7\x18\x25\x66\x76\x08\xda\xfe\x66\x46\xbe\x1a\xc9\x2f\x01\x93\x06\x2d\x36\x2c\x3f\x9d\x02\x08\xf6\x8a\x50\x7f\xe7\x51\xf7\x30\xeb\x14\x40\x01\x9a\x0b\xc7\xdc\x9a\x69\xe8\x0c\xfa\xd0\x95\x05\x5b\x45\xf3\x46\xc1\xc7\x5c\x57\x87\xb9\x7e\xe0\x7d\x12\xa7\x6c\x3f\x4c\xc0\xf4\x15\xf9\x7d\x02\x6e\x78\x95\xa2\x00\xeb\xf2\x42\xb4\xde\x33\x8a\x98\x96\xf1\x95\xbe\xd7\x3b\x0e\x9d\x23\xcf\xd3\x25\x77\xff\x84\x17\xa6\xb4\xd6\xe1\xd3\xe9\xf9\x08\x8f\x47\x38\xdd\xf7\x5a\xf5\x4a\xed\x79\xbc\xdc\x24\xe8\x32\x53\xfa\x99\x3d\xb3\x22\xd1\x91\x4d\x14\x83\x06\x3d\x6f\x0c\x72\xd9\x36\x86\x48\x66\x0e\xf9\x5b\xb9\xac\x0e\x4d\xb2\x66\x0d\xb4\x21\x3f\x5a\xb1\xf3\x00\xf2\xdb\xf9\xdb\xb2\xe2\xc5\x4e\x77\x5c\xd1\xab\xbf\x00\x00\x00\xff\xff\x6e\xe7\xaf\xff\x61\x02\x00\x00")

func templatesSubnetTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesSubnetTf,
		"templates/subnet.tf",
	)
}

func templatesSubnetTf() (*asset, error) {
	bytes, err := templatesSubnetTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/subnet.tf", size: 609, mode: os.FileMode(480), modTime: time.Unix(1541042574, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesVarsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x91\x41\x4e\xc3\x30\x10\x45\xf7\x39\xc5\xc8\x07\x00\x5a\x24\x76\x59\xa0\x22\xf5\x00\x15\x2b\x84\xac\x89\x3d\x14\x53\xc7\x8e\x66\x6c\x57\xb4\xea\xdd\x51\xd2\x42\x93\x16\x48\x56\xe3\xff\xf4\xc7\xff\xbb\x20\x3b\x6c\x3c\x81\xa2\x50\x74\xc0\x96\x14\xec\x0f\x55\x35\x39\x77\x1c\x43\x4b\x21\x89\x82\x7d\x05\x90\x3e\x3b\x82\x1a\x54\x8b\x9d\xaa\x00\x2c\xbd\x61\xf6\x09\xea\x41\x04\x78\xdc\x65\xa6\x85\x8f\xd9\xc2\xe9\xab\x41\x75\xb9\xf1\xce\xa8\x33\xf0\xbc\x5a\xc6\x42\x3c\xf8\xf6\x40\x96\xf5\xcf\x3c\xc2\x96\xc4\x2d\x86\x93\x5b\x0d\x6a\x3d\xcc\x23\x60\xf1\xee\x02\x7e\x6f\xab\x41\x99\x7e\xee\xf5\x43\x75\x18\xa5\xc0\x9e\xd5\xa3\x2c\x97\x31\x8f\x40\xa2\x80\x21\x69\x67\x7f\x97\x25\x37\x62\xd8\x75\xc9\xc5\xf0\x27\x64\xbc\xa3\x7f\x3c\x4e\xb2\x90\x61\xba\xba\x85\x8f\x06\x7b\xf3\x63\xd1\xe7\x66\x15\xa1\xa4\x2c\x73\x55\x4d\xf0\x0f\x17\x6c\xcc\xba\x38\x4e\x19\xbd\x0e\x94\xb6\x91\x37\x1a\xad\x65\x12\xd1\xd2\xa1\xa1\xd1\x9b\x1d\x1b\xf2\x4e\xd2\xf4\xdd\x5e\xd4\xec\xee\x66\xf8\x6f\x67\x0f\xea\x75\xba\x63\x4b\x8d\x16\xe2\x42\xac\x4b\xab\x4d\xcc\xd7\xdd\x4d\x11\x71\x3b\xba\xba\xff\x2a\x61\xb0\xc8\x56\x3f\xcd\x45\x97\xfb\x8b\x1c\x3e\xa2\xd5\x0d\x7a\x0c\x86\x58\xcb\x26\x0f\x1b\xbe\x02\x00\x00\xff\xff\x08\x71\x7e\x0f\x9c\x02\x00\x00")

func templatesVarsTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesVarsTf,
		"templates/vars.tf",
	)
}

func templatesVarsTf() (*asset, error) {
	bytes, err := templatesVarsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/vars.tf", size: 668, mode: os.FileMode(480), modTime: time.Unix(1541042574, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesVmssTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\xc1\x8a\xe4\x36\x10\xbd\xfb\x2b\x84\xc9\x75\x1a\xf7\x64\x37\x90\x83\x0f\x21\x24\xb0\xb0\x81\x61\x67\x43\x0e\x21\x88\xb2\x54\xee\x51\x46\x96\x44\x49\xf2\x6c\xef\xd2\xff\xbe\x48\xb6\xa6\xdd\x6e\x37\x73\x58\x9f\xba\x55\x4f\xaf\x9e\xf4\xca\x55\x26\xf4\x36\x92\x40\x56\xc3\xd7\x48\x48\x03\xd7\x1d\x37\x10\xb8\xb3\x56\xd7\xac\xfe\x5f\x19\x69\x23\x1f\x07\xef\x17\xcb\xdf\x2a\xc6\xca\x4e\x7e\x20\x1b\x1d\x37\x30\x20\x5b\x3c\x2d\xab\x7f\xfa\x56\x38\x2f\xb1\xbb\x99\x74\xb5\x9a\x18\x4e\x75\xc5\xd8\x9a\x6a\xf5\xb4\x45\xd5\x9d\xf7\x4f\x09\xaf\x2d\xc8\x0e\x34\x18\x81\xc4\x95\xdc\xc0\x9f\x95\xe8\xae\x64\x4f\xbb\x78\xd9\xb6\x53\x32\xa7\x76\x64\x83\x15\x56\xdf\x4e\xfd\x59\xb8\x04\xec\xc9\x9a\x80\x46\x72\x67\x29\x70\x1f\x80\xc2\x25\xf0\x7d\xd3\x34\xcd\x15\x10\x8d\x5c\x33\xbe\x6f\xf6\xfb\x5f\x2b\xc6\x3a\x10\xcf\x05\xb7\x99\xfa\xfe\x7e\x49\xa7\x1c\x17\xd6\xf4\xea\x10\x09\x82\xb2\x66\x32\xa0\x65\xf5\x43\xec\xb4\x12\x1f\x1e\x7e\x93\x92\xd0\xfb\xba\x3a\x55\xd5\xb5\xcb\xa3\xa2\x10\x41\xf3\x01\xc4\x93\x32\xc8\xbd\x00\x8d\xdc\x63\xb8\xf4\x7c\xb2\x7a\xcb\x90\x96\xd5\xc3\x31\xa0\x0f\x79\xa7\xc7\x70\xb7\x9f\xbc\x10\x59\xce\x8f\x95\x42\x61\xc9\x9e\x6c\x15\xda\x0f\x54\x57\x74\x07\x02\x89\xdc\x59\xad\xc4\x91\x0f\x56\x66\xba\xbf\xc0\x44\xd0\x75\x55\x31\xe6\x9f\x63\x3e\xf6\xe2\xe0\x39\xdf\x08\xb4\x7b\xc1\x8e\x7b\xa4\x11\x89\x8f\x03\xf7\xea\xeb\x44\xca\x58\x50\x48\x05\xfa\x18\xc0\x48\x20\x39\x45\x04\x38\x10\x2a\x1c\x6f\x91\x08\x1b\x4d\xc8\x2c\xa7\x9c\x3d\x58\x82\x03\x72\x47\xb6\x57\x1a\xb9\x1a\xd2\x3f\xc2\x1e\x09\x8d\xc0\x59\x99\x4b\x2e\xfb\x27\xa4\xc4\xfa\x3b\x18\x6b\x94\x48\xf2\x53\xcc\xf6\xfd\x59\xcb\xdf\x5d\x34\x21\x3e\xe6\x6c\x53\x38\x1d\xef\xd5\x98\xfd\x2f\xbb\xe6\xdd\xdd\xc7\xcf\x8f\x53\x6c\x44\xf2\x93\x7b\x2d\xab\x35\x24\x7f\x6f\x0a\xb3\x9e\x4b\xe5\x9f\xd7\x57\xb5\xb4\xbd\x5c\x40\xaa\xb1\xc3\x65\xe8\x13\x82\xfc\x87\x54\xc0\x19\x43\x08\x01\xb9\x75\xaf\xc5\xd3\xb2\xfa\x4f\xb2\xc3\x87\x74\xfc\x09\x33\x80\x81\x03\xca\x9c\x95\x87\xa3\xc3\xe5\x5d\xf3\x8f\x9f\x1e\x6f\x4a\x95\x10\x60\x29\x56\xc7\x55\x81\x36\x5b\x42\xdf\x54\xd9\xb2\xfa\x8f\xc1\x85\xe3\x14\xcc\xb2\x52\x45\xf0\x43\x97\x77\xef\x9b\xa2\xc7\xfa\x22\x65\x16\x20\xec\xe0\x62\x40\xca\xd5\xcc\x1d\x61\xaf\xbe\x9c\xeb\x03\xcd\x98\x03\xa7\xbb\x71\x98\xb8\x41\x0e\xca\xf0\xe8\x91\x16\xf7\x9c\x5f\xc1\x1c\x59\x82\x1c\x78\xff\x62\x49\x9e\x41\x0f\x69\x25\x2d\xed\xef\x7f\x7e\x57\x5f\x6b\xe2\x5a\x99\xf8\x65\x6e\x27\xb3\x40\xa9\x3c\x74\x1a\x5f\xd9\x38\xc4\xf0\x84\x26\xa8\xf9\xf5\x6e\x59\x0f\xda\x63\x21\x33\x18\x5e\x2c\x3d\xaf\x4e\x59\xb4\xbe\x76\xec\xd4\x52\x66\xe8\x8c\x9c\x94\x3b\x52\x03\x50\x7a\x43\x02\x45\xac\xf2\xda\xba\xc1\xcd\x9c\x6f\x8e\x87\x85\x7b\x57\x49\x95\xbb\xa0\xac\x67\xc6\x92\xfe\x4d\xc6\xac\x6e\xfa\xe3\x63\x67\x30\x6c\x8c\x9b\x2b\x15\xe7\x46\x35\xed\xd9\x79\x30\xe5\xe7\x3c\x76\xd2\x73\x31\x8f\x78\x99\x06\x30\x75\xf1\x3c\x77\xb9\x92\x9e\xb5\xec\xdf\x8b\x71\xb6\x89\xdc\x9c\x71\xdb\xc8\xa4\xe0\xbf\x4d\x09\xca\x74\x36\x1a\x99\xc7\x3e\x45\x8d\x3e\xe7\xcf\x67\x5a\x49\x28\x1f\x06\xbb\xad\xaf\x85\x45\x86\x53\x2e\x97\xd3\xf7\x00\x00\x00\xff\xff\xa5\x28\xe7\x7b\x74\x08\x00\x00")

func templatesVmssTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesVmssTf,
		"templates/vmss.tf",
	)
}

func templatesVmssTf() (*asset, error) {
	bytes, err := templatesVmssTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/vmss.tf", size: 2164, mode: os.FileMode(480), modTime: time.Unix(1541042574, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesVnetTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8e\x41\xaa\xc3\x30\x0c\x44\xf7\x3e\x85\x10\x59\xe7\xff\x6c\xba\xcb\x49\x4a\x31\x22\x11\x25\x6d\x62\x05\xd9\x4e\xa1\xc6\x77\x2f\x09\x75\xa1\xae\xb4\xd3\xcc\xe8\x8d\xb2\x97\xa8\x03\x03\xd2\x33\x2a\xeb\x62\xb7\x49\x43\xa4\xd9\x3a\x0e\x0f\xd1\x3b\x02\xde\x26\x37\x4a\xfc\x15\x92\x01\x70\xb4\x30\x54\xd3\x97\x08\x1a\x80\x02\xb0\x57\x95\xb8\xda\xc3\xdf\x03\x36\xa9\xf0\xbe\x0d\xed\x1b\x56\x5d\xf7\x58\xde\xdf\xd1\x38\x2a\x7b\x6f\xfd\x4a\x03\x7f\x78\x67\xec\xfe\xdb\x63\xff\xba\x13\x5e\x0c\xc0\x2c\x03\x85\x49\x5c\xd5\xab\x49\x1b\x69\x5b\xc4\x8c\x26\x9b\x57\x00\x00\x00\xff\xff\x60\x01\xba\xdc\x02\x01\x00\x00")

func templatesVnetTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesVnetTf,
		"templates/vnet.tf",
	)
}

func templatesVnetTf() (*asset, error) {
	bytes, err := templatesVnetTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/vnet.tf", size: 258, mode: os.FileMode(480), modTime: time.Unix(1541042574, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesWeb_serverTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x55\x4d\x8f\xdb\x36\x10\xbd\xeb\x57\x4c\xd5\x9c\x0a\xac\xd1\xb4\x41\x6f\x3e\x05\x28\x50\x20\x45\x8b\xb8\x1f\x87\x20\x20\xc6\xe2\x58\x9e\x5a\x1c\x0a\x43\x52\x9b\x64\xe1\xff\x5e\x90\x92\x1c\x7b\x2d\x77\xb7\x01\xe2\x83\x21\x89\x6f\x3e\xf8\x38\xef\xb1\x52\x0a\x3e\x69\x43\x50\xe3\xa7\xa4\xa4\xce\x08\xc5\x7b\xaf\x07\xc3\x12\x49\x77\xd8\x50\x0d\xf5\x3f\x2c\xd6\x27\x23\xdc\xd4\xf0\x50\x01\x08\x3a\x82\x47\xbf\x35\xd4\x2f\x1e\x06\xd4\x15\xc9\x60\x32\xe0\x78\x27\xdc\xdc\xbd\x78\x68\x7c\x92\xb8\x62\xb1\xf4\xe1\x58\x57\x00\xe5\xfd\x56\xf0\x3d\x6d\x4d\x20\x1d\x48\xcd\xe0\x4c\x81\x96\xa0\xce\x37\x18\xd9\xcb\x55\xd0\xdc\xf6\xbc\x11\xd3\xaa\x4f\xfd\x6a\xea\xf8\xd1\xd7\x39\x4b\x49\x79\xb9\x56\x5a\xfe\x82\x94\x65\xa7\x75\x55\x01\x70\x6f\x1a\x2f\x3b\x6e\x93\x8e\xad\x66\xa6\x96\xb9\x7a\x8a\x35\xee\x17\x78\x03\x08\x69\x2b\x14\x0d\xdb\xdb\xa9\xe6\xde\x47\xe8\x2a\xa0\xcc\x8f\x6c\xa7\x2c\xbd\xf2\x80\x91\x0c\xf7\x06\xad\x55\x0a\xc1\x60\x77\xe2\x77\x0d\xb5\xfd\x28\xe8\xb8\xc9\xe8\x63\x75\xac\x16\x66\x64\x60\x8d\x09\x3b\xe3\xb0\xd9\xb3\x9c\x4d\xc8\xe0\x6e\x0f\xc8\xe2\x66\xc7\xb8\xbb\xc1\x2d\x6c\x78\xe9\xcc\xbf\xca\xa9\x7f\x49\xd2\xe9\xdc\x01\xae\xf4\x62\xd8\x06\x58\xc3\xbb\xb3\x8c\x57\x98\xd5\x67\x45\xad\xbe\x5b\xb1\x7d\x77\xb6\xf7\xf7\xc7\xfa\x7d\x05\x30\x38\x13\xf8\xd3\x63\x16\x6f\x08\x25\x23\x6f\x8a\xeb\x49\x79\xe5\xe9\xfd\x16\x5e\x7b\xe7\x48\x22\xc4\x3d\x07\xe8\x58\x08\xa2\x87\x03\x51\x0f\x71\x4f\xf0\xdb\x06\x2c\x87\x03\xe0\x2e\x92\x82\xa5\x8e\x22\x4b\x5b\x96\xfe\xfa\xb5\x82\xf1\x0b\x19\x1f\x4c\x86\x19\x2f\x26\x92\x3a\x96\x79\xaa\xa2\x26\x7a\x4e\x1d\x8b\x11\x4b\xa5\xf0\x54\xa9\x8c\x2c\xc5\xc2\xed\x6a\x21\x7a\xc5\x96\x0c\xbb\xfc\xaf\xb4\x23\x25\x69\x68\x92\x66\x9f\xb6\x1d\x87\x3d\x69\x26\xe8\x35\x8a\x17\x6e\xb0\x1b\x55\xe2\x77\x3b\xd2\x99\xbc\x3f\xb7\x49\x62\xda\x14\xe2\x26\x29\x1e\xd2\x67\x6e\x5f\xfe\xb4\xfa\xfe\xd5\xdd\x9b\x3f\x36\xe3\xda\x40\x1a\xc6\xb1\x5d\x43\xdd\x61\xa4\x10\x47\x2d\x9d\x35\x34\xf1\x74\xd3\x23\xd6\xb3\xa4\x06\x97\x71\x4b\x6e\xd0\x14\xf1\xb5\x97\x41\x6f\x09\xed\xdf\xca\x91\x26\x8c\x52\x96\xba\xef\x4f\x3a\x5a\x43\xfd\xb3\x7a\xf7\x4b\x26\x64\xc4\x38\x14\x6c\xc9\x8e\xe7\x16\x3f\xf6\xc5\x05\x37\x11\xc5\xa2\x5a\xf3\xe6\xed\xe6\xd4\xbc\x0f\xa6\x57\xbf\xe3\x6e\x26\xb0\xf1\xae\x4f\x91\x74\x92\xd1\x1a\xea\xbd\x0f\x31\xbf\x2c\x35\x8c\xd6\xb1\x98\x14\x48\x67\xaf\xcd\xd4\x94\xaf\xe7\x80\x1e\x43\xb8\xf7\x6a\x33\xe0\xf7\xe9\xf9\xe5\x0f\x3f\xbe\xfa\x66\xa1\x0f\xd3\xb1\xa4\x0f\x93\xef\x4e\x4d\x59\x0e\xb8\xed\xe8\x94\xc7\x60\x8a\x7b\x92\xc8\x27\x87\xdb\x61\x17\x68\x4e\x16\xb1\x0d\x53\x24\xc9\xc0\xea\xa5\x4c\xe7\x1a\xea\x10\xb1\x65\x69\xff\xc3\x07\xaf\xb5\xbf\xc5\xe6\x40\x62\x4f\xce\xda\x7b\xdf\x19\x0c\xc1\x37\x5c\xaa\x9f\x5b\x65\x56\xff\x29\xe0\x1a\xf8\x70\x53\xce\xcf\xb8\x2f\x97\x5c\xe9\xca\x3d\xff\xaf\x2d\x2d\x5c\x72\x17\x06\xfa\xbc\x7b\x6c\x91\x22\xb6\x97\xad\x75\xdb\x45\x2a\xe7\x06\x3b\x8f\xd6\x6c\xb1\x43\x69\x48\x97\x91\xe5\xbe\x3b\x56\xff\x06\x00\x00\xff\xff\xc7\xed\x3e\xd1\xe2\x08\x00\x00")

func templatesWeb_serverTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesWeb_serverTf,
		"templates/web_server.tf",
	)
}

func templatesWeb_serverTf() (*asset, error) {
	bytes, err := templatesWeb_serverTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/web_server.tf", size: 2274, mode: os.FileMode(480), modTime: time.Unix(1541043837, 0)}
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
	"templates/db_mysql.tf": templatesDb_mysqlTf,
	"templates/devbox.tf": templatesDevboxTf,
	"templates/load_balancer.tf": templatesLoad_balancerTf,
	"templates/network_security_group.tf": templatesNetwork_security_groupTf,
	"templates/output.tf": templatesOutputTf,
	"templates/provider.tf": templatesProviderTf,
	"templates/resource_group.tf": templatesResource_groupTf,
	"templates/storage_account.tf": templatesStorage_accountTf,
	"templates/subnet.tf": templatesSubnetTf,
	"templates/vars.tf": templatesVarsTf,
	"templates/vmss.tf": templatesVmssTf,
	"templates/vnet.tf": templatesVnetTf,
	"templates/web_server.tf": templatesWeb_serverTf,
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
	"templates": &bintree{nil, map[string]*bintree{
		"db_mysql.tf": &bintree{templatesDb_mysqlTf, map[string]*bintree{}},
		"devbox.tf": &bintree{templatesDevboxTf, map[string]*bintree{}},
		"load_balancer.tf": &bintree{templatesLoad_balancerTf, map[string]*bintree{}},
		"network_security_group.tf": &bintree{templatesNetwork_security_groupTf, map[string]*bintree{}},
		"output.tf": &bintree{templatesOutputTf, map[string]*bintree{}},
		"provider.tf": &bintree{templatesProviderTf, map[string]*bintree{}},
		"resource_group.tf": &bintree{templatesResource_groupTf, map[string]*bintree{}},
		"storage_account.tf": &bintree{templatesStorage_accountTf, map[string]*bintree{}},
		"subnet.tf": &bintree{templatesSubnetTf, map[string]*bintree{}},
		"vars.tf": &bintree{templatesVarsTf, map[string]*bintree{}},
		"vmss.tf": &bintree{templatesVmssTf, map[string]*bintree{}},
		"vnet.tf": &bintree{templatesVnetTf, map[string]*bintree{}},
		"web_server.tf": &bintree{templatesWeb_serverTf, map[string]*bintree{}},
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

