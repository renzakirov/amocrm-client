package tests

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAccount(t *testing.T) {
	client := CreateAmoClient(t)
	acc, err := client.GetAccount(false)
	require.NoError(t, err)
	require.NotNil(t, acc)
	require.NotEmpty(t, acc.ID)
	require.NotEmpty(t, acc.Name)
	require.NotEmpty(t, acc.Subdomain)
}
