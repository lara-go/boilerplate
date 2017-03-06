package conf_test

import (
	"testing"

	"github.com/lara-go/boilerplate/app/conf"
	"github.com/stretchr/testify/assert"
)

func TestConfigLoading(t *testing.T) {
	config := conf.ConfigLoader()

	assert.Equal(t, "production", config.Env())
	assert.False(t, config.Debug())
}

func TestGetValue(t *testing.T) {
	config := conf.ConfigLoader()
	value := config.Get("App.Env")

	assert.Equal(t, "production", value)
	assert.Panics(t, func() {
		config.Get("App.Foo")
	})
}

func TestSetValue(t *testing.T) {
	config := conf.ConfigLoader()
	config.Set("App.Env", "test")

	value := config.Get("App.Env")
	assert.Equal(t, "test", value)

	assert.Panics(t, func() {
		config.Set("App.Foo", "foo")
	})
}
