package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetServerInfo(t *testing.T) {
	expectedAPIVersion := "1"
	expectedAuthor := "DrakeEsdon & HugoKlepsch"
	expectedColor := "#03fcf4"
	expectedHead := "pixel"
	expectedTail := "pixel"
	var actual = GetServerInfo()

	assert.Equal(t, expectedAPIVersion, actual.APIVersion, "APIVersion")
	assert.Equal(t, expectedAuthor, actual.Author, "Author")
	assert.Equal(t, expectedColor, actual.Color, "Color")
	assert.Equal(t, expectedHead, actual.Head, "Head")
	assert.Equal(t, expectedTail, actual.Tail, "Tail")
}