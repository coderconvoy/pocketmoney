package main

import (
	"encoding/hex"
	"encoding/json"
)

func CookieMarshal(d interface{}) (string, error) {
	rb, err := json.Marshal(d)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(rb), nil
}

func CookieUnmarshal(s string, d interface{}) error {
	b, err := hex.DecodeString(s)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, d)

}
