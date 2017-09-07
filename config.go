package main

import (
	"os"
	"path"

	"github.com/coderconvoy/lazyf"
	"github.com/pkg/errors"
)

func confLocs() []string {
	home := os.Getenv("HOME")
	return []string{
		"test_data/.pmconf",
		".pmconf",
		path.Join(home, ".config/pocketmoney/init"),
	}
}

func getConfig(s string) (lazyf.LZ, error) {
	locs := confLocs()
	if s != "" {
		locs = []string{s}
	}

	carr, err := lazyf.GetConfig(locs...)

	if err != nil {
		return lazyf.LZ{}, err
	}

	if len(carr) < 1 {
		return lazyf.LZ{}, errors.Errorf("No Entry in Config")
	}

	return carr[0], nil

}
