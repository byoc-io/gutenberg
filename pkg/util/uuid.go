package util

import (
	"crypto/rand"
	"encoding/base32"
	"github.com/satori/go.uuid"
	"io"
	"regexp"
	"strings"
)

var encoding = base32.NewEncoding("abcdefghijklmnopqrstuvwxyz234567")

// NewUUID returns a random generated UUID V4.
func NewUUID() string {
	return uuid.NewV4().String()
}

// NewFastUUID returns a random string which can be used as an ID for objects.
func NewFastUUID() string {
	buff := make([]byte, 16) // 128 bit random ID.
	if _, err := io.ReadFull(rand.Reader, buff); err != nil {
		panic(err)
	}
	// Trim padding
	return strings.TrimRight(encoding.EncodeToString(buff), "=")
}

// IsValidUUID checks if string is a valid UUID V4
func IsValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}
