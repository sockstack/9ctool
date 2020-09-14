package response

import (
	"github.com/sockstack/9ctool/response"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResponse(t *testing.T) {
	response := response.NewResponse(200, "message", nil)
	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "message", response.Message)
	assert.Equal(t, nil, response.Data)
}

func TestSUCCESS(t *testing.T) {
	assert.Equal(t, 20000, response.SUCCESS.Code)
	assert.Equal(t, "success", response.SUCCESS.Message)
	assert.Equal(t, nil, response.SUCCESS.Data)
}

func TestFAIL(t *testing.T) {
	assert.Equal(t, 40000, response.FAIL.Code)
	assert.Equal(t, "fail", response.FAIL.Message)
	assert.Equal(t, nil, response.FAIL.Data)
}

func TestERROR(t *testing.T) {
	assert.Equal(t, 50000, response.ERROR.Code)
	assert.Equal(t, "error", response.ERROR.Message)
	assert.Equal(t, nil, response.ERROR.Data)
}
