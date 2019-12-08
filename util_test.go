package gowizz

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenTimeRanges(t *testing.T) {

	ranges := GenTimeRanges(time.Date(2009, 11, 17, 1, 34, 58, 651387237, time.UTC), 30*Day, 4)

	// then
	require.Len(t, ranges, 4)
	assert.Equal(t, time.Date(2009, 11, 17, 0, 0, 0, 0, time.UTC), ranges[0].From)
	assert.Equal(t, time.Date(2009, 12, 17, 0, 0, 0, 0, time.UTC), ranges[0].To)
	assert.Equal(t, time.Date(2009, 12, 18, 0, 0, 0, 0, time.UTC), ranges[1].From)
	assert.Equal(t, time.Date(2010, 01, 17, 0, 0, 0, 0, time.UTC), ranges[1].To)
	assert.Equal(t, time.Date(2010, 01, 18, 0, 0, 0, 0, time.UTC), ranges[2].From)
	assert.Equal(t, time.Date(2010, 02, 17, 0, 0, 0, 0, time.UTC), ranges[2].To)
	assert.Equal(t, time.Date(2010, 02, 18, 0, 0, 0, 0, time.UTC), ranges[3].From)
	assert.Equal(t, time.Date(2010, 03, 20, 0, 0, 0, 0, time.UTC), ranges[3].To)
}
