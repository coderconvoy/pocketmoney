// Code generated by go-bindata.
// sources:
// assets/s/main.css
// assets/templates/.viewac.html.swp
// assets/templates/familypage.html
// assets/templates/forms.html
// assets/templates/index.html
// assets/templates/layout.html
// assets/templates/newaccount.html
// assets/templates/newuser.html
// assets/templates/transactions.html
// assets/templates/userhome.html
// assets/templates/viewac.html
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

var _assetsSMainCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x8b\xb1\xae\xc2\x30\x0c\x45\xf7\x7c\x85\xa5\xb7\x3e\xb1\xa0\x0c\x4d\xbf\x26\xc6\x75\x14\xd1\xd8\x91\x49\x24\x10\xe2\xdf\x29\xa1\x2c\x70\xc7\x7b\xce\x41\xa5\xdb\xdd\xc1\x36\x8c\xa7\x73\x32\xed\x42\xe1\x6f\x9a\x98\x11\x67\xf7\x70\xee\x50\x16\xe9\xf0\xab\x78\xcf\x1c\xe3\x50\x58\xad\x7c\x0c\x35\x5a\x2c\xf8\x7a\x85\x8b\xae\x99\x00\xd7\xad\x99\x07\x2b\xd1\x52\x96\x17\x1b\x55\x96\xda\x1b\xfc\x03\xf6\xd6\x54\xf6\xfe\xcb\x69\xb4\xff\x35\x12\x65\x49\xe1\xf8\x06\xcf\x00\x00\x00\xff\xff\x54\x41\xbd\xbe\xb6\x00\x00\x00")

func assetsSMainCssBytes() ([]byte, error) {
	return bindataRead(
		_assetsSMainCss,
		"assets/s/main.css",
	)
}

func assetsSMainCss() (*asset, error) {
	bytes, err := assetsSMainCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/s/main.css", size: 182, mode: os.FileMode(420), modTime: time.Unix(1489057121, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesViewacHtmlSwp = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x9a\xcf\x6b\xd4\x40\x14\xc7\xa7\x05\xf1\x62\xab\xd4\x83\x82\x22\xd3\x75\x05\x0b\x35\x69\xab\xa0\xd4\x18\x28\x96\x4a\xb5\xfe\xa4\x15\x4f\xc2\x64\x32\xbb\x09\x4d\x32\x31\x33\xd9\x52\x43\xea\x7f\xd0\x8b\x57\x8f\xfe\x13\x1e\xed\xc1\x3f\xc2\x93\x07\x8f\x5e\xbd\x8a\x6f\xd2\x74\xb3\xdb\x56\xf6\x22\x8a\xf8\x3e\xf0\xdd\x99\x79\x6f\xe6\xbd\x37\x33\xb9\x84\x8d\x37\xf7\x7c\xf5\x21\xbd\x69\xdd\x20\xc0\x19\x42\xa6\x7e\xec\xbd\x38\x7f\x62\x76\xfc\xd3\x24\x21\x31\xd3\x3a\x10\x5b\x64\x14\xf7\x44\xa6\xb7\x17\x46\x4e\x23\x3b\x75\x40\x5b\xf1\x2c\x4c\xb5\xb2\xbb\x32\x62\x49\xd7\x56\x19\xb7\xbb\xa1\x0e\x72\xcf\xe2\x32\xb6\xb9\xf4\x45\xc6\x65\xd2\x93\xdb\x76\x2a\xf9\xa6\xd0\xb1\x4c\xc4\xb6\xcd\x94\x12\xb0\x48\x8b\x38\x8d\x98\x16\xca\xee\x85\x62\x8b\x71\x2b\xd0\x71\x34\x3a\x39\x82\xfc\xcf\xe4\xba\x73\xed\xd6\x04\xb9\xbe\x30\x3f\x67\x86\x97\x5b\xd3\xf4\xec\xd4\xc6\xdf\xae\x0a\x41\x10\x04\x41\x10\x04\x41\x90\x3f\x88\x4e\xc7\xc8\x1b\x68\xc7\xeb\xf1\xa5\xba\x1d\x3b\xd4\x22\x08\x82\x20\x08\x82\x20\x08\x82\x20\x08\xf2\xef\xc2\x7c\x42\xd4\x29\x42\x4e\x4e\x90\xea\xff\xff\x83\xf7\xff\x6f\xa7\x09\xf9\x0c\xfa\x08\x7a\x0f\x7a\x07\x7a\x0b\x7a\x0d\xea\x80\x9e\x82\x16\x41\xd3\xa0\x73\xa0\x2f\x93\x84\xec\x81\x3e\x80\x76\x41\x6b\xa0\x07\xa0\xef\x10\xf7\x2b\x68\x17\xb4\x03\x7a\x09\x72\x40\xb7\x41\xf3\xa0\x2b\xa0\x8b\xa0\x0b\x13\x75\x0d\x08\x82\x20\x08\x82\x20\x08\x82\x20\xbf\x8b\xa2\x30\x1f\x10\xef\xff\x46\x75\xc3\x74\xbf\x23\x68\xcb\x63\x4a\xb4\xca\x12\xde\x53\x6c\xcd\xbc\x48\xb8\x84\x02\x45\x21\x12\x1f\xac\xa6\x0f\x8e\xcc\x25\x55\xb7\x1a\x6a\xdf\x2d\x0a\xdf\x2c\x6e\xf7\xac\x65\x68\xcb\x12\xa6\xf8\xee\xa1\x19\xe0\x7c\x92\x67\xa9\x54\x47\xfc\x83\xc1\x9b\xf9\xd5\x67\xcf\x26\xe4\x52\x2c\xf3\x44\xd7\x8b\x8c\xaf\xdf\xa9\x82\xae\x64\x32\xde\x50\x22\x2b\x4b\xba\x48\xfb\x96\xa5\xbb\xc7\x64\x89\x4c\xee\xa1\x34\x03\xa1\x7e\x9d\xaf\x0a\xba\x2c\x94\x1e\x4e\x63\x2c\x4d\x9a\x81\x3c\x61\x87\xb2\xc4\xa7\x57\xc5\x2b\xda\xd6\x32\xb5\x56\x62\x11\xd3\x81\x4a\x67\xf6\x5d\x8c\xd3\x7e\xb1\x33\xf4\x70\x61\xde\x50\x49\x1d\x6d\x32\x3b\xb6\xe7\x36\x9b\x72\xcc\x35\xec\x67\xcc\x58\xd2\x85\xe3\xdf\xa4\xb3\xed\x1e\x5d\xbc\x43\xdb\x6b\xa1\xd2\x75\x44\xa8\xd5\x8c\x8c\xd9\x5a\x61\xb1\xb5\xc4\x79\x1e\xe7\xe6\xb2\xd7\x61\x99\x62\x5c\x87\x32\x51\xb4\xae\x92\xf1\xa1\x5b\x6e\x4a\x0a\xdc\xfa\xf6\xc0\x11\x0c\x3b\x1e\xeb\x40\x64\xc7\x98\x37\xd6\x8f\x1a\x57\x1f\x1d\xb5\x3d\xcb\x93\x24\x4c\xba\x8d\xa3\xda\x99\x53\x3f\x7e\x4e\xb0\x00\x47\x51\x95\x77\x70\xf6\xa6\x48\xc7\x06\xfb\xc1\x06\xe1\x28\xcd\xf6\xee\x4b\x8f\xb6\x18\x6f\x35\x3b\x87\xe3\xaf\x3c\x60\x19\x7c\xc8\x03\xc1\x7c\x33\xeb\x67\x00\x00\x00\xff\xff\x4f\x90\x57\x9a\x00\x30\x00\x00")

func assetsTemplatesViewacHtmlSwpBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesViewacHtmlSwp,
		"assets/templates/.viewac.html.swp",
	)
}

func assetsTemplatesViewacHtmlSwp() (*asset, error) {
	bytes, err := assetsTemplatesViewacHtmlSwpBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/.viewac.html.swp", size: 12288, mode: os.FileMode(420), modTime: time.Unix(1489174047, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesFamilypageHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x93\xc1\x6e\x9c\x30\x10\x86\xef\x79\x8a\xd1\x9c\x12\xa9\x81\x86\x28\x97\x08\x90\xaa\x48\x91\x7a\x48\xb5\x52\xd5\x07\x30\x78\x36\x58\xb5\x0d\xb5\x8d\x76\x57\x16\xef\x5e\x83\x61\x59\xb8\x64\x2f\x8b\xcd\xff\xff\x33\xf3\x19\x7b\xef\x48\x75\x92\x39\x02\x6c\x88\x71\x1c\x86\x3b\x08\xbf\xbc\x79\x2a\xdf\x99\x12\xf2\x02\x07\xf6\x49\xaf\xe0\x7d\x12\xd6\x49\xdc\xfb\xc5\x14\x0d\x43\x9e\x06\x51\x54\xf7\x32\x3e\x78\x6f\x98\xfe\x24\x98\xb4\x1f\xa4\x2a\x32\x76\x4e\x9c\x74\x52\x94\xd7\x45\xd4\x27\x7f\x2c\x19\x3d\xe5\xc1\xe3\x63\xd8\xa8\x7e\x93\xa4\xda\x41\x72\x60\x86\xb4\x03\x8c\xff\x08\xf8\xd6\x08\x39\x36\x98\x57\x66\x4d\xc9\xd3\x25\xd3\x7b\xd2\x7c\x69\x3f\x5d\x3a\x1a\xc5\x57\x83\xf7\xe2\x08\xf7\xc2\x76\x31\x3a\x79\x57\xa4\xa6\x5e\x1f\x16\xdf\xb1\x35\x0a\x58\xed\x44\xab\x0b\x64\x9c\xab\x69\x06\x04\x45\xae\x69\x79\x81\x5d\x6b\x1d\x96\x70\xb7\x0e\x10\x12\xc3\xa4\x9b\x29\x9b\x67\xa8\x25\xb3\xb6\x40\x45\xd6\x06\x7c\x58\x86\x41\x83\x08\x26\x68\xcf\xe5\x8d\x7d\xed\x39\x5a\xb3\xf2\x07\xe7\x30\xa3\x8f\x04\x83\x25\x5b\x2d\x23\x7b\x78\x85\x5c\xe8\xae\x77\xe0\x2e\x1d\x15\xe8\xe8\x1c\x00\x8d\x14\x0b\xec\x67\x9e\xb8\x5a\x7e\x5a\x88\x10\x77\xb6\xba\xa1\xfa\x6f\xd5\x9e\x17\x6b\xc4\x82\xe5\x06\xf0\x21\x0c\x72\x6a\x0d\xdf\xd7\xec\xe6\x7d\x84\x8e\x39\x17\x6a\x16\x98\xf8\x97\x6f\xd9\xf7\x01\xc1\xd0\xbf\x5e\x18\xe2\x4b\xee\x89\x3f\xed\x52\xef\x3f\x84\x16\xaa\x57\xf0\x02\x92\x46\xbb\x7d\xd8\xbc\x7f\x6b\xf5\x51\x84\xa3\xf8\xba\xfa\xb5\x44\xb6\x2b\xb1\xd1\xdb\xbe\x52\xc2\xcd\x4c\xf2\x74\x3c\xe6\xed\x57\x73\x7b\x11\x2a\x66\x69\xbc\x08\xff\x03\x00\x00\xff\xff\x3b\xe2\xc7\x9e\x1d\x03\x00\x00")

func assetsTemplatesFamilypageHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesFamilypageHtml,
		"assets/templates/familypage.html",
	)
}

func assetsTemplatesFamilypageHtml() (*asset, error) {
	bytes, err := assetsTemplatesFamilypageHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/familypage.html", size: 797, mode: os.FileMode(420), modTime: time.Unix(1488902512, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesFormsHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xdc\x55\xcb\x6e\xdb\x3a\x10\xdd\xfb\x2b\x08\x22\x8b\x7b\x81\x40\x71\x72\x2f\x50\xc0\xb0\x04\x18\x4e\xbd\xcb\x03\x6d\xba\x36\x18\x73\x1c\x09\x11\x1f\xa5\x28\xbb\x82\xa0\xaf\xe9\x9f\xf4\xcb\x3a\x94\x28\x59\x72\x1c\xdb\x05\xb2\xaa\x16\x09\x1f\x67\x0e\x67\xce\x3c\x5c\x96\x1c\xd6\x89\x04\x42\xd7\x46\x2c\x19\xe7\x4b\xb6\x5a\xa9\x5c\x5a\x5a\x55\x23\x82\xdf\x74\xad\x8c\x20\x6c\x65\x13\x25\x43\x8a\x80\xf6\x9e\x08\xb0\xb1\xe2\x21\xd5\x2a\xb3\x34\x6a\xc0\xf1\x7f\xd1\x8c\x73\x32\x6b\x30\xd3\x2b\xdc\x37\x17\x89\xd4\xb9\x25\xb6\xd0\x10\x52\x0b\x3f\xd0\x5c\x32\x81\xeb\x3c\x03\xe3\x56\x94\x18\x60\x5c\xc9\xb4\x20\x1b\x96\xe6\x78\x53\x96\xc1\x42\x80\xa8\x2a\x4a\xe2\x84\x73\x90\x27\x98\xbc\x5f\x0d\x99\x66\xd6\x22\x71\x48\x83\xf2\xff\xcb\x9b\x71\xe5\xe8\xbf\xe7\x89\x01\x7e\x80\x25\xcb\x9f\x45\x82\x3c\x8e\x2d\xa4\xe8\x3f\x6d\x7d\x70\x6b\x6f\x70\xe5\x74\x88\x46\x65\x09\x92\xa3\x34\xa3\x9e\x70\x42\x2f\x45\xe1\x9f\xcf\x3a\xdd\x32\x48\x61\x65\xbd\x73\x6b\xa3\x84\x67\x72\x5f\x59\x1a\x26\x5f\x80\x5c\xbc\x92\x4b\x72\xb1\x21\x93\x90\xfc\x83\x04\x90\x65\xb3\x79\x46\x82\x05\x13\x81\x97\xd0\xed\x50\x85\xfa\xec\x5f\x4f\x5d\xd3\x2b\xed\x32\xb2\x13\xeb\x62\x13\x7c\xf3\x5a\x56\xd5\xa4\xde\xdf\xd7\x6b\x1a\x1d\xb9\x9c\x38\x5f\x84\x92\x50\xa0\x1b\xc1\x3c\x37\x06\xa4\xad\xaa\xe9\x55\x43\xdf\xf7\xb8\x89\xbb\xd1\xa2\x89\x6d\xa7\xc6\x50\x0c\x96\xa6\xc7\xd5\xb0\xea\xa0\x16\x97\x8d\x12\x83\xf0\x3f\x3e\xe4\x3f\x8d\x6d\x90\x6a\xec\x11\xcd\x8a\xc3\xbd\xe1\x2e\x0e\x35\x45\xdb\x18\x77\xec\x15\x08\x23\x8f\xac\x10\xf0\xe1\xbd\x41\x48\xcd\xb5\xc0\x3a\x9b\x60\x40\x16\x84\x4e\x99\x7d\x5b\x9d\x24\xf0\xbe\x3f\xa9\xb7\xb8\x7e\xe2\x1c\x70\xfa\x6c\x1a\x17\x67\xc2\x1d\x62\xb1\xfc\xfa\x39\xf0\x56\xe6\xe2\x19\x0c\x25\x99\x05\x1d\xd2\x71\x30\xbe\x46\x05\x12\xd4\x62\xdc\x75\x65\x6d\x49\xa3\x8e\xea\x31\x37\xa8\x0d\x4c\x8e\xc4\xad\x1b\x48\xcf\xe8\x44\xcb\xa2\xa4\xe4\x5e\x6d\xbb\xb6\x6d\xf7\xa7\x5b\x17\xf3\x99\x59\x26\x79\x22\x5f\xde\x1d\x78\x1d\xe0\x8c\xe4\x7e\xf5\x58\xf2\x60\x38\x98\x5d\x8e\xff\xc2\xbc\xd4\x56\x18\xaf\xb1\xe4\x16\x3d\x9d\x0c\xed\x38\x1e\xb5\x76\x99\x03\xd1\x5d\xed\xba\xbb\x2f\x8b\xb9\xeb\xd3\xee\xf9\xa7\x18\x24\xf9\xbc\x01\x53\xec\x11\xb5\xa1\xd4\xee\x5f\xb7\x94\x1c\x52\xd7\x6d\x9e\xf2\x53\x9b\xe9\xc1\x94\xa9\x31\x4b\xc7\xd2\xcf\xd5\x70\x88\x70\x56\x64\x34\xba\xc5\xbf\x6f\xc7\xc2\x1e\x14\xa7\xa4\x8d\x11\x7c\x57\xff\x1f\xc2\xbb\xa9\xd1\xc5\x73\x4e\xd9\xce\xb1\xad\x9d\x4a\xfe\x01\xbf\x3d\xab\x68\x57\xb1\x66\x59\x3d\x5d\x87\xe5\xea\xcf\xdf\xfb\x6d\x9e\xc7\xf5\xa8\x7d\x44\xcc\x56\x19\xbe\xab\xcf\x87\x94\x77\xa7\x64\x2f\x03\xda\x9f\xb7\xda\xab\x94\xeb\x2d\xef\xe5\xee\x1e\xb6\x44\x9f\x69\x8c\x96\xd7\x3d\xd3\xb9\x92\xeb\x04\xfd\x3f\x6d\x75\xd3\xb3\x3a\xa6\x69\x1d\xe1\x4e\xd3\x66\x1b\x8d\xf6\xf5\xfc\x1d\x00\x00\xff\xff\x96\xd7\x4c\x57\xf8\x08\x00\x00")

func assetsTemplatesFormsHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesFormsHtml,
		"assets/templates/forms.html",
	)
}

func assetsTemplatesFormsHtml() (*asset, error) {
	bytes, err := assetsTemplatesFormsHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/forms.html", size: 2296, mode: os.FileMode(420), modTime: time.Unix(1489153098, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x54\xcd\xce\x9b\x30\x10\xbc\xf3\x14\x2b\x9f\x5a\xa9\xcd\x0f\x52\x2e\x15\x20\x55\x51\x7b\x0a\x55\x2e\x7d\x00\x03\x4b\xb1\x8a\x6d\x6a\x9b\xa6\x11\xe2\xdd\x6b\xc7\x90\x90\x84\x34\xcd\x97\x13\xbb\x9e\x19\xaf\x66\x36\xee\x3a\x83\xbc\xa9\xa9\x41\x20\x15\xd2\x82\xf4\x7d\x00\xc3\x2f\xaa\xd6\xc9\x5e\xe6\x3f\xd1\x40\x2a\x05\x1e\xa3\xa5\x6d\x04\xe7\xe3\xae\x63\x25\x2c\xa6\xf8\x2c\xe9\xba\x45\x8a\xba\xef\xa3\x65\x96\x4c\x80\x28\x0a\x8b\x0b\x2e\xdc\xa8\x94\x8a\x03\xcd\x0d\x93\x22\x26\xb5\xfc\xc1\x04\x01\x8e\xa6\x92\x45\x4c\x1a\xa9\x0d\xb9\xd0\xfd\x28\x61\xb2\x73\x28\x3b\x43\x78\x7d\xf4\x95\x72\x56\x1f\x3f\x41\xc4\x44\xd3\x1a\x30\xc7\x06\x63\x62\xf0\x8f\x21\x20\x28\xb7\xdf\xe5\x09\x40\xe0\x37\xad\x5b\x5b\xda\x11\x3d\xa5\xef\x49\x12\x65\xea\x5a\x2d\x45\x9e\xa1\x82\x6f\x96\x09\xff\xd0\x6c\x35\x2a\xf7\x35\x55\xfd\x3e\xf4\x66\x75\xf7\x54\xeb\x83\x54\xc5\xad\x68\x33\xf4\x47\xe1\xe6\x50\xcc\xb0\xaf\x28\xba\xcd\x38\xb3\x93\xb8\x79\x62\xb2\xf3\xde\x0d\x63\xf8\xea\x42\x8e\x96\xce\xe8\xe4\xa1\xf3\x02\x0f\xa3\x3d\x4f\xdc\xdf\x2a\x74\x4b\xe2\xad\x83\xcf\x79\x2e\x5b\x61\x1e\xa5\xf1\xd4\x3f\x7f\xe9\xc9\xc1\x99\x0c\x28\x13\xe0\xdc\xfc\xaf\x04\xee\xf9\x5f\x38\x65\xf5\x2d\x17\x5d\x73\x24\xfb\xe2\x2d\x29\x35\xd4\x18\x7b\x71\x4c\x16\xdd\xe6\x43\xb8\xea\x09\x28\xfc\xd5\x32\x85\x05\x18\x66\x6a\x8b\xdd\x7c\x0c\x57\x90\x57\x54\x59\x8f\x51\xe9\x49\xb0\xeb\x99\x1b\xdf\xa5\x4c\x30\xde\x72\xd8\x40\x8d\x4e\x5a\xbf\xbf\xc3\x6c\xa5\x28\x99\x0d\xed\x85\x1d\x0a\x5f\x5b\x22\x1f\xee\x79\x8b\x86\x72\xf2\x5f\x1f\xf7\x68\xfa\x5c\x64\x54\xa3\x7d\x2e\x46\x48\xf0\x37\x00\x00\xff\xff\xdf\x16\x24\x30\x4b\x04\x00\x00")

func assetsTemplatesIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesIndexHtml,
		"assets/templates/index.html",
	)
}

func assetsTemplatesIndexHtml() (*asset, error) {
	bytes, err := assetsTemplatesIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/index.html", size: 1099, mode: os.FileMode(420), modTime: time.Unix(1489069089, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesLayoutHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x51\x3d\x73\xb3\x30\x0c\xde\xf9\x15\x7e\xb5\xbf\xf5\xda\xc1\x78\xe9\xc7\xd4\x5e\x18\xb2\x74\x54\x40\x04\x5f\x8c\x9d\x43\xa2\x57\x8e\xcb\x7f\xaf\xc1\x6d\x43\x87\xc6\x8b\xbe\x9e\x47\x96\xf4\xcc\x73\x43\xad\x0b\xa4\xa0\x23\x6c\xe0\x72\x29\xcc\xbf\xc7\xdd\xc3\xfe\xad\x7a\x52\x9d\xf4\xde\x16\x26\x1b\x95\x9e\x59\x30\xd9\x5d\xc3\x9e\x04\x55\xdd\xe1\xc0\x24\x25\x8c\xd2\xfe\xbf\x87\x4d\xd9\xbb\x70\x52\x03\xf9\x12\x58\x26\x4f\xdc\x11\x09\x28\x99\xce\x54\x82\xd0\x87\xe8\x9a\x19\x54\x37\x50\x5b\x82\x66\xdd\xa3\x0b\x77\x4b\x6a\xd3\x42\x9c\x78\xb2\x55\xac\x4f\x24\xea\x35\x06\x9a\x8c\xce\xb9\x3c\x90\xbe\x4e\x64\x0e\xb1\x99\x36\xd4\xc6\xbd\xab\xda\x23\x73\x09\x3d\x85\x71\xd3\x75\x2d\xe3\xf7\xc7\x2d\xf6\xce\x4f\x60\x9f\x57\x6b\x34\xfe\x05\x3c\xd3\xc0\x31\xa0\x07\x5b\x7d\x79\x37\xc0\x32\x60\x60\xac\xc5\xc5\x90\x16\xda\x6f\xa2\x1b\x24\x1f\x8f\x71\x14\xb0\x2f\xf1\xa8\x76\xa3\xfc\x42\x1a\x9d\x16\xb2\xc5\x3c\x53\x68\x92\x4a\xe9\x5d\xa5\x3b\x20\xd3\x22\x5d\xc6\xe5\x3b\xa4\xd3\xac\xba\xfd\x10\x3e\x03\x00\x00\xff\xff\x7a\x8a\x17\xb6\xea\x01\x00\x00")

func assetsTemplatesLayoutHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesLayoutHtml,
		"assets/templates/layout.html",
	)
}

func assetsTemplatesLayoutHtml() (*asset, error) {
	bytes, err := assetsTemplatesLayoutHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/layout.html", size: 490, mode: os.FileMode(420), modTime: time.Unix(1488553008, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesNewaccountHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb2\x51\x74\xf1\x77\x0e\x89\x0c\x70\x55\xc8\x28\xc9\xcd\xb1\xe3\xe2\xb2\x81\xd0\x0a\x40\x60\x93\x91\x9a\x98\x02\x61\x82\xb9\x25\x99\x25\x39\xa9\x76\x1e\xa9\x39\x39\xf9\x36\xfa\x10\x0e\x42\x32\x37\xb5\x24\x51\x21\x39\x23\xb1\xa8\x38\xb5\xc4\x56\xa9\xb4\x24\x4d\xd7\x42\x09\x6a\x8c\x3e\xc2\x1c\x9b\xa4\xfc\x94\x4a\x24\x5d\x19\x86\x76\x7e\xa9\xe5\x0a\xa1\xc5\xa9\x45\x40\x65\x86\x08\x99\xea\xea\x82\xa2\xcc\xbc\x92\x34\x05\x25\xd5\x62\x25\x05\xbd\xda\x5a\xa8\x51\x10\xfd\x40\xb5\x10\xd7\x02\x02\x00\x00\xff\xff\xd6\x4d\xae\x17\xbf\x00\x00\x00")

func assetsTemplatesNewaccountHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesNewaccountHtml,
		"assets/templates/newaccount.html",
	)
}

func assetsTemplatesNewaccountHtml() (*asset, error) {
	bytes, err := assetsTemplatesNewaccountHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/newaccount.html", size: 191, mode: os.FileMode(420), modTime: time.Unix(1488373318, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesNewuserHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb2\x51\x74\xf1\x77\x0e\x89\x0c\x70\x55\xc8\x28\xc9\xcd\xb1\xe3\xe2\xb2\x81\xd0\x0a\x40\x60\x93\x91\x9a\x98\x02\x61\x82\xb9\x25\x99\x25\x39\xa9\x76\x1e\xa9\x39\x39\xf9\x36\xfa\x10\x0e\x42\x32\x37\xb5\x24\x51\x21\x39\x23\xb1\xa8\x38\xb5\xc4\x56\xa9\xb4\x24\x4d\xd7\x42\x09\x6a\x8c\x3e\xc2\x1c\x9b\xa4\xfc\x94\x4a\x24\x5d\x19\x86\x76\x7e\xa9\xe5\x0a\xa1\xc5\xa9\x45\x40\x65\x86\x08\x99\xea\x6a\xbd\x50\xbf\xc4\xdc\xd4\xda\x5a\x9b\xa4\x22\x24\x0d\x89\x0a\x19\x45\xa9\x69\xb6\x4a\xfa\x40\x6d\x8e\xc9\xc9\xf9\xa5\x79\x25\x4a\x60\x23\xa0\x1c\x1b\xfd\x44\x98\xad\x10\xab\x80\xc6\x42\x3c\x06\x08\x00\x00\xff\xff\xa6\xb7\x94\xc4\xea\x00\x00\x00")

func assetsTemplatesNewuserHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesNewuserHtml,
		"assets/templates/newuser.html",
	)
}

func assetsTemplatesNewuserHtml() (*asset, error) {
	bytes, err := assetsTemplatesNewuserHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/newuser.html", size: 234, mode: os.FileMode(420), modTime: time.Unix(1488373372, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesTransactionsHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x90\xc1\x0a\xc2\x30\x10\x44\xef\xf9\x8a\x50\x7a\x94\x16\x3c\x96\x18\x28\x4a\xcf\x1e\xf4\x03\x52\xbb\x58\xb1\x49\x4a\xb2\x2d\x48\xc8\xbf\xbb\x51\xd4\x0a\xcd\x21\xb0\xbb\x8f\x61\x66\x42\x40\xd0\xe3\xa0\x10\x78\xd6\x83\xea\xb2\x18\x99\xe8\xb7\xf2\xe4\x94\xf1\xea\x82\x37\x6b\xbc\x28\x69\xc1\x04\xaa\x76\x00\xc9\x42\xa0\xd3\x15\x78\x7e\xdf\xe4\x33\xaf\x76\xbc\x68\x94\x2e\x96\x3c\x4f\x1a\xe8\x24\xe3\xf4\x04\x76\x32\x84\x7c\x2e\x1a\x67\xf5\xd9\x83\x8b\xb1\xfa\xce\xf5\x3e\x46\x51\x12\xf1\x8f\x1e\xc0\xe3\x12\x4d\xf3\x1a\xaa\xad\x81\x07\x27\xa0\xd6\x76\x32\xb8\xaa\x75\x9c\xdc\x68\x3d\x7c\x6e\xf4\x93\x31\x0a\x01\xa6\x7b\xd9\x7c\x87\x4a\x9b\x5f\x11\xad\xf2\x90\x8a\x78\x06\x00\x00\xff\xff\x56\x12\x6d\xf6\x1d\x01\x00\x00")

func assetsTemplatesTransactionsHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesTransactionsHtml,
		"assets/templates/transactions.html",
	)
}

func assetsTemplatesTransactionsHtml() (*asset, error) {
	bytes, err := assetsTemplatesTransactionsHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/transactions.html", size: 285, mode: os.FileMode(420), modTime: time.Unix(1489056589, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesUserhomeHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x91\x4f\x6b\xc3\x30\x0c\xc5\xef\xf9\x14\x22\x94\x1e\xca\x88\x61\xbd\x75\x8e\xc7\x18\xec\xb6\x31\x18\x3b\x17\x35\x51\x9b\xb0\xd8\xcd\xec\xa4\xa5\x18\x7f\xf7\xc9\xf9\xb3\x8d\x6e\xb9\x04\xeb\xf7\xa4\xa7\x67\x7b\xdf\x91\x6e\x1b\xec\x08\xd2\x8a\xb0\x4c\x43\x48\x64\x75\xab\xde\x1d\x59\x78\xc5\x03\xc1\x06\xbc\xcf\x9e\x34\xe9\x10\xa4\x60\xc2\x78\xad\x1e\x8a\xe2\xd8\x9b\xce\x71\x65\xad\x12\xef\xeb\x3d\x64\xcf\xe4\x86\xe6\x35\x14\x0d\x3a\x97\xa7\x9a\x9c\xe3\x09\xa9\xe2\x01\x03\x9c\xd5\x64\x4a\x56\x7a\x2f\x56\x6f\xd5\xf1\x0c\xf3\xb0\x95\xe0\x6a\x22\xfb\x26\x6a\x16\xd1\x12\x36\x39\x4c\xde\x5c\xb2\x68\x78\x9f\xc5\xc7\x0d\x2c\x4e\x23\x41\x9d\xcd\xcd\xac\x00\xfe\x86\x55\xe8\x93\x15\x59\x8c\x60\x50\x73\xc7\x34\x21\x72\xd9\xd4\x4a\x22\x54\x96\xf6\x79\x2a\x4e\x35\x9d\xef\xfb\x28\xca\xd9\xf1\xa7\x25\x84\x25\x16\x63\xe9\x65\x38\xc6\x0c\xdf\x07\x29\x50\x2d\xcd\xce\xb5\x77\xde\xeb\xa3\xa1\x4b\x74\x7b\xec\xad\x25\xd3\x45\xca\x16\xd3\x2e\x73\xd0\xf1\x2f\x45\x8c\x96\xfc\xbe\xf2\xbd\xd5\x5b\x2c\xcb\x2d\x8e\x29\x52\xc8\xe2\x15\x5c\x2b\x5a\xbc\x8c\xe4\x1a\xb8\x0e\x4d\x59\x9b\x03\x53\xf8\x07\x17\x55\xcb\x0f\xf1\xb7\x75\x87\x8e\xe2\x43\x7f\x05\x00\x00\xff\xff\x22\xff\xd7\x4d\xfd\x01\x00\x00")

func assetsTemplatesUserhomeHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesUserhomeHtml,
		"assets/templates/userhome.html",
	)
}

func assetsTemplatesUserhomeHtml() (*asset, error) {
	bytes, err := assetsTemplatesUserhomeHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/userhome.html", size: 509, mode: os.FileMode(420), modTime: time.Unix(1489066634, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesViewacHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x52\x5d\x6b\xc2\x30\x14\x7d\xf7\x57\x5c\xa4\x0f\x13\x46\x85\x3d\x4a\x57\x28\x0e\x61\x63\x6c\x63\xe8\x0f\x48\xdb\xab\x2d\x33\x89\x4b\x52\x61\x84\xfe\xf7\xdd\x34\xe9\x6a\xad\xcc\x17\x6f\xcf\x39\xb9\xe7\x7e\x59\x6b\x90\x9f\x8e\xcc\x20\xcc\x2b\x64\xe5\xbc\x6d\x67\x40\x3f\x6b\x23\x23\x4f\xb0\x7a\x84\x78\x40\x58\xd1\x01\x2f\x32\x87\x39\x2b\x9c\x34\xa9\x1e\x52\x6b\xe3\x0d\x47\xde\xb6\xb0\xf2\xa2\xb6\x4d\x96\x84\xcf\x12\xc3\xf2\x23\xa6\xdd\xeb\xc4\x28\x1f\xf8\x8f\x2a\xfd\x6c\x84\xa8\xc5\x21\x59\x52\x3c\x22\x9e\xdf\xa6\xd8\xfb\x6e\x7b\x03\x34\x15\xaa\x29\xfc\xd1\xa8\x93\xd4\x38\x10\x14\x05\x6f\xaa\xee\xb5\xd6\xa6\x6b\x62\xc3\x78\x9c\x15\x45\xc3\x1b\xd7\xfc\x56\x31\xa1\x59\x61\x6a\x29\x34\x74\xed\x40\xd7\x48\x78\x46\xec\x01\x21\xfa\x82\xfb\xe8\xec\x5e\x77\x69\x02\x7b\xd5\x59\x99\x26\x39\x8d\x84\x4b\x81\x3f\x10\x9d\xe3\x6c\x6f\x50\xb9\x89\xe4\x29\x15\x52\x0e\x52\x6b\xeb\x3d\x30\x51\xc2\x1d\x7e\x83\x1b\x76\xb0\x3d\xc7\x1b\x25\xf9\x4e\xa3\x5a\x78\x8a\xa6\x1e\xc0\x6c\xbd\x00\x72\x1d\xbb\xb9\xa4\x2e\xb8\xf4\xe4\xb2\x11\xc6\x99\xfe\x71\x84\x3e\xa1\x36\x2e\x6d\xbf\x28\x8f\x64\xeb\xa0\xbb\x28\x0c\x8f\x1a\x43\x77\xbd\xcd\x7f\xd9\xc7\x36\x7d\xf5\x83\x8d\x2f\xfd\x86\x8d\x28\x27\x2e\x24\x0f\x0b\xbc\xd6\x7b\xbe\x74\x97\xea\x4a\xa7\xff\x5e\x31\xd9\xb2\x4f\x4c\x88\xbf\xbf\xd9\xe5\x91\xe7\x4c\xa3\xbb\xdc\xdf\x00\x00\x00\xff\xff\x21\x62\xce\x52\xf9\x02\x00\x00")

func assetsTemplatesViewacHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesViewacHtml,
		"assets/templates/viewac.html",
	)
}

func assetsTemplatesViewacHtml() (*asset, error) {
	bytes, err := assetsTemplatesViewacHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/viewac.html", size: 761, mode: os.FileMode(420), modTime: time.Unix(1489174047, 0)}
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
	"assets/s/main.css": assetsSMainCss,
	"assets/templates/.viewac.html.swp": assetsTemplatesViewacHtmlSwp,
	"assets/templates/familypage.html": assetsTemplatesFamilypageHtml,
	"assets/templates/forms.html": assetsTemplatesFormsHtml,
	"assets/templates/index.html": assetsTemplatesIndexHtml,
	"assets/templates/layout.html": assetsTemplatesLayoutHtml,
	"assets/templates/newaccount.html": assetsTemplatesNewaccountHtml,
	"assets/templates/newuser.html": assetsTemplatesNewuserHtml,
	"assets/templates/transactions.html": assetsTemplatesTransactionsHtml,
	"assets/templates/userhome.html": assetsTemplatesUserhomeHtml,
	"assets/templates/viewac.html": assetsTemplatesViewacHtml,
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
	"assets": &bintree{nil, map[string]*bintree{
		"s": &bintree{nil, map[string]*bintree{
			"main.css": &bintree{assetsSMainCss, map[string]*bintree{}},
		}},
		"templates": &bintree{nil, map[string]*bintree{
			".viewac.html.swp": &bintree{assetsTemplatesViewacHtmlSwp, map[string]*bintree{}},
			"familypage.html": &bintree{assetsTemplatesFamilypageHtml, map[string]*bintree{}},
			"forms.html": &bintree{assetsTemplatesFormsHtml, map[string]*bintree{}},
			"index.html": &bintree{assetsTemplatesIndexHtml, map[string]*bintree{}},
			"layout.html": &bintree{assetsTemplatesLayoutHtml, map[string]*bintree{}},
			"newaccount.html": &bintree{assetsTemplatesNewaccountHtml, map[string]*bintree{}},
			"newuser.html": &bintree{assetsTemplatesNewuserHtml, map[string]*bintree{}},
			"transactions.html": &bintree{assetsTemplatesTransactionsHtml, map[string]*bintree{}},
			"userhome.html": &bintree{assetsTemplatesUserhomeHtml, map[string]*bintree{}},
			"viewac.html": &bintree{assetsTemplatesViewacHtml, map[string]*bintree{}},
		}},
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

