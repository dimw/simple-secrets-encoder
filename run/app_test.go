package run

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateApp(t *testing.T) {
	app := CreateApp()

	assert.Equal(t, 3, len(app.Commands))
	assert.Equal(t, "encrypt", app.Commands[0].Name)
	assert.Equal(t, "decrypt", app.Commands[1].Name)
	assert.Equal(t, "generate-keys", app.Commands[2].Name)
}
