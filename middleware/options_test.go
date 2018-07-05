package middleware

import (
	"strconv"
	"strings"
	"testing"
)

func TestDefaultAddress(t *testing.T) {
	oi := &Options{}
	addr := oi.getAddress()
	if !strings.Contains(addr, defaultHost) || !strings.Contains(addr, strconv.Itoa(defaultPort)) {
		t.Error("Incorrect address generated")
	}
}

func TestDefaultRequestKey(t *testing.T) {
	oi := &Options{}
	reqKey := oi.getRequestKey()
	if reqKey != defaultRequestKey {
		t.Error("Incorrect requestKey generated")
	}
}

func TestNonDefaultAddress(t *testing.T) {
	host := "abcd"
	port := 123
	oi := &Options{Host: host, Port: port}
	addr := oi.getAddress()
	if !strings.Contains(addr, host) || !strings.Contains(addr, strconv.Itoa(port)) {
		t.Error("Incorrect address generated")
	}
}

func TestNonDefaultRequestKey(t *testing.T) {
	key := "abcd"
	oi := &Options{RequestKey: key}
	reqKey := oi.getRequestKey()
	if reqKey != key {
		t.Error("Incorrect requestKey generated")
	}
}
