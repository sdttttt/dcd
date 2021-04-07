package huck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const TestConfigYamlFileName = "huck.test.yml"

func TestConfigYamlParse(t *testing.T) {
	conf := FromConfigFile(TestConfigYamlFileName)
	assert.Equal(t, "/test_1", conf.Counter[0].Path, "test yaml counter one item Test.")
	assert.Equal(t, "伞兵lbw准备就绪!", conf.Counter[0].Name, "test yaml counter one item Test.")
}
