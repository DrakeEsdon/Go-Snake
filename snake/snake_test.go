package snake

import (
	"github.com/DrakeEsdon/Go-Snake/datatypes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDirection(t *testing.T) {
	allDirs := datatypes.AllDirections

	expectedDirs := []datatypes.Direction{datatypes.DirectionRight, datatypes.DirectionDown, datatypes.DirectionLeft}
	actualDirs := removeDirection(allDirs, datatypes.DirectionUp)
	assert.Equal(t, expectedDirs, actualDirs)

	expectedDirs = []datatypes.Direction{datatypes.DirectionUp, datatypes.DirectionDown, datatypes.DirectionLeft}
	actualDirs = removeDirection(allDirs, datatypes.DirectionRight)
	assert.Equal(t, expectedDirs, actualDirs)

	expectedDirs = []datatypes.Direction{datatypes.DirectionUp, datatypes.DirectionRight, datatypes.DirectionLeft}
	actualDirs = removeDirection(allDirs, datatypes.DirectionDown)
	assert.Equal(t, expectedDirs, actualDirs)

	expectedDirs = []datatypes.Direction{datatypes.DirectionUp, datatypes.DirectionRight, datatypes.DirectionDown}
	actualDirs = removeDirection(allDirs, datatypes.DirectionLeft)
	assert.Equal(t, expectedDirs, actualDirs)
}
