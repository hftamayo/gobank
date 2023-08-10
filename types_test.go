package main

import (
	"testing"
	"github.com/zeebo/assert"

func TestNewAccount(t *testing.T) {
	acc, err := NewAccount("a", "b", "hftamayo")
	assert.Nil(t, err)

}
