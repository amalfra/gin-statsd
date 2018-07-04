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
