package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	acc, err := NewAccount("a", "b", "hftamayo")
	assert.Nil(t, err)

	fmt.Printf("%+v\n", acc)

}
