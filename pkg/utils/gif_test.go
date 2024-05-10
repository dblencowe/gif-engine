package utils_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"vcs.services.strawberryelk.internal/strawberryelk/gif-engine/pkg/utils"
)

func TestGIFEditor_Join(t *testing.T) {
	editor := &utils.GIFEditor{}

	err := editor.LoadToBuffer("../../test/resources/utils/gif/anakin.gif")
	assert.NoError(t, err)
	err = editor.LoadToBuffer("../../test/resources/utils/gif/explosion.gif")
	assert.NoError(t, err)

	rst, err := editor.Join()
	assert.NoError(t, err)
	assert.Less(t, 0, rst.Len())
}
