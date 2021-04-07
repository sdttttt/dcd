package huck

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHuckCore(t *testing.T) {
	go Run(TestConfigYamlFileName)
	time.Sleep(time.Second)

	resp, err := http.Get("http://127.0.0.1:10086/test_1")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode, "test_1 should is Ok.")

	counter := kernel.stat["test_1"].(*Counter)
	assert.NotNil(t, counter)
	assert.Greater(t, counter.count, uint64(0), "test_1 counts should greater 0.")
}
