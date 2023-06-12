package utils_test

import (
	"testing"

	"github.com/heeropunjabi/go-handson/utils"
)

func TestOms(t *testing.T) {
	oms := utils.Oms()
	if oms != "o1" {
		t.Errorf("Expected o1, got %s", oms)
		
	}
}