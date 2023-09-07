package interest

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {

	// No days
	r, err := NewRate(0, 0)
	require.NotNil(t, err)

	// Values
	rate := 970./349.8 - 1
	days := 297
	r, err = NewRate(rate, days)
	require.Nil(t, err)
	assert.Equal(t, r.Value(), rate)
	assert.Equal(t, r.Days(), days)

}

func TestResample(t *testing.T) {

	// Prepare
	rate := 970./349.8 - 1
	days := 297
	r, err := NewRate(rate, days)
	require.Nil(t, err)

	// Resample
	annual, err := r.Resample(0)
	require.NotNil(t, err)

	annual, err = r.Resample(365)
	assert.Nil(t, err)
	const delta = 0.0000000001
	expected := 2.50242248572357
	assert.InDelta(t, expected, annual.value, delta)
}

func TestNominalYearly(t *testing.T) {

	// Prepare
	rate := 970./349.8 - 1
	days := 297
	r, err := NewRate(rate, days)
	require.Nil(t, err)

	// Resample
	nominal, err := r.NominalYearly()
	require.Nil(t, err)

	const delta = 0.0000000001
	expected := 1.30221200490504
	assert.InDelta(t, expected, nominal, delta)
}
