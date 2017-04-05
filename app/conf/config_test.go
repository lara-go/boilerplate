package conf_test

import (
	"testing"

	"github.com/lara-go/boilerplate/app/conf"
	"github.com/stretchr/testify/assert"
)

func TestConfigLoading(t *testing.T) {
	config := conf.NewConfig()

	assert.Equal(t, "production", config.Env())
	assert.False(t, config.Debug())
}
