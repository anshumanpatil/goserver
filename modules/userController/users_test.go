package usercontroller

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestLoginResult(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LoginResult(tt.args.c)
		})
	}
}

func TestAllUsers(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AllUsers(tt.args.c)
		})
	}
}
