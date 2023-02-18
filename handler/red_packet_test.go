package handler

import (
	"ginDemo/model"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestCheckParams(t *testing.T) {
	m := model.SendRpReq{}
	b := checkParams(m)
	t.Log(b)

	assert.Equal(t, b, false)
}

func TestSendRedPacket(t *testing.T) {

}
