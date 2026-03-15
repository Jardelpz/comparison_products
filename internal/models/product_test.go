package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetProductDefaultFields_IgnoresIDAndSpecsAndCategory(t *testing.T) {
	fields := GetProductDefaultFields()
	require.NotEmpty(t, fields)

	require.NotContains(t, fields, "id")
	require.NotContains(t, fields, "specs")

	require.Contains(t, fields, "name")
	require.Contains(t, fields, "price")
}
