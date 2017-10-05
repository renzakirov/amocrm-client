package tests

import (
	"testing"

	"github.com/esporo-org/amocrm-client/amo"
	"github.com/stretchr/testify/require"
)

// CreateAmoClient creates Amo Client and tests authorization
func CreateAmoClient(t *testing.T) *amo.Client {
	envs, err := newEnvParams()
	require.NoError(t, err)
	client, err := amo.NewClient(envs.AccountWebAddress, nil)
	require.NoError(t, err)
	auth, err := client.Authorize(envs.UserLogin, envs.UserHash)
	require.NoError(t, err)
	require.NotNil(t, auth)
	require.True(t, auth.Response.Auth)
	return client
}
