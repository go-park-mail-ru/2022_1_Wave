package test

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/tools/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBaseTypeConverterToString(t *testing.T) {
	var someStr interface{}
	const expected = "hello, world"
	someStr = expected
	result, err := utils.ToString(someStr)

	require.NoError(t, err)
	require.Equal(t, result, expected)
}

func TestBaseTypeConverterToBool(t *testing.T) {
	var boolean interface{}
	const expected = true
	boolean = expected
	result, err := utils.ToBool(boolean)

	require.NoError(t, err)
	require.Equal(t, result, expected)
}

func TestBaseTypeConverterToUint64(t *testing.T) {
	var uints interface{}
	const value = 555
	const expected = uint64(value)
	uints = float64(expected)
	result, err := utils.ToUint64(uints)

	require.NoError(t, err)
	require.Equal(t, result, expected)
}

func TestBaseTypeConverterToInt(t *testing.T) {
	var ints interface{}
	const value = -543
	const expected = int(value)
	ints = float64(value)
	result, err := utils.ToInt(ints)

	require.NoError(t, err)
	require.Equal(t, result, expected)
}

func TestBaseTypeConverterToInt64(t *testing.T) {
	var ints64 interface{}
	const value = 909
	const expected = int64(value)
	ints64 = float64(value)
	result, err := utils.ToInt64(ints64)

	require.NoError(t, err)
	require.Equal(t, result, expected)
}
