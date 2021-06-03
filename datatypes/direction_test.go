package datatypes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDirectionToStr(t *testing.T) {
	assert.Equal(t, "up", DirectionToStr(DirectionUp))
	assert.Equal(t, "right", DirectionToStr(DirectionRight))
	assert.Equal(t, "down", DirectionToStr(DirectionDown))
	assert.Equal(t, "left", DirectionToStr(DirectionLeft))
}