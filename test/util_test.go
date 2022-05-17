package test

import (
	"mygomodule/tool"
	"testing"
)

func TestAdd(t *testing.T) {
	if tool.Add(1, 3) != 4 {
		t.Fatal("TestAdd fail")
	}
}
