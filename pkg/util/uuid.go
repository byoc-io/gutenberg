package util

import (
	"crypto/rand"
	"encoding/base32"
	"io"
	"strings"
)

var encoding = base32.NewEncoding("abcdefghijklmnopqrstuvwxyz234567")

// NewUUID returns a random string which can be used as an ID for objects.
func NewUUID() string {
	buff := make([]byte, 16) // 128 bit random ID.
	if _, err := io.ReadFull(rand.Reader, buff); err != nil {
		panic(err)
	}
	// Trim padding
	return strings.TrimRight(encoding.EncodeToString(buff), "=")
}
