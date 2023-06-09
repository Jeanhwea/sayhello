package main

import (
	"testing"
)

func TestMain01(t *testing.T) {
	t.Log(GetHost())
}

func TestMain02(t *testing.T) {
	t.Log(GetIP())
}
