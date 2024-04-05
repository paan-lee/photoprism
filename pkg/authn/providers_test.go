package authn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProviderType_String(t *testing.T) {
	assert.Equal(t, "default", ProviderUndefined.String())
	assert.Equal(t, "default", ProviderDefault.String())
	assert.Equal(t, "none", ProviderNone.String())
	assert.Equal(t, "local", ProviderLocal.String())
	assert.Equal(t, "ldap", ProviderLDAP.String())
	assert.Equal(t, "link", ProviderLink.String())
	assert.Equal(t, "access_token", ProviderAccessToken.String())
	assert.Equal(t, "client_credentials", ProviderClientCredentials.String())
}

func TestProviderType_Is(t *testing.T) {
	assert.False(t, ProviderLocal.Is(ProviderLDAP))
	assert.True(t, ProviderLDAP.Is(ProviderLDAP))
	assert.False(t, ProviderClient.Is(ProviderLDAP))
	assert.False(t, ProviderClientCredentials.Is(ProviderLDAP))
	assert.False(t, ProviderApplication.Is(ProviderLDAP))
	assert.False(t, ProviderAccessToken.Is(ProviderLDAP))
	assert.False(t, ProviderNone.Is(ProviderLDAP))
	assert.False(t, ProviderDefault.Is(ProviderLDAP))
	assert.False(t, ProviderUndefined.Is(ProviderLDAP))
}

func TestProviderType_IsNot(t *testing.T) {
	assert.False(t, ProviderLocal.IsNot(ProviderLocal))
	assert.True(t, ProviderLDAP.IsNot(ProviderLocal))
	assert.False(t, ProviderClient.IsNot(ProviderClient))
	assert.False(t, ProviderClientCredentials.IsNot(ProviderClientCredentials))
	assert.False(t, ProviderApplication.IsNot(ProviderApplication))
	assert.False(t, ProviderAccessToken.IsNot(ProviderAccessToken))
	assert.False(t, ProviderNone.IsNot(ProviderNone))
	assert.False(t, ProviderDefault.IsNot(ProviderDefault))
	assert.False(t, ProviderUndefined.IsNot(ProviderUndefined))
}

func TestProviderType_IsUndefined(t *testing.T) {
	assert.True(t, ProviderUndefined.IsUndefined())
	assert.False(t, ProviderLocal.IsUndefined())
}

func TestProviderType_IsRemote(t *testing.T) {
	assert.False(t, ProviderLocal.IsRemote())
	assert.True(t, ProviderLDAP.IsRemote())
	assert.False(t, ProviderClient.IsRemote())
	assert.False(t, ProviderClientCredentials.IsRemote())
	assert.False(t, ProviderApplication.IsRemote())
	assert.False(t, ProviderAccessToken.IsRemote())
	assert.False(t, ProviderNone.IsRemote())
	assert.False(t, ProviderDefault.IsRemote())
	assert.False(t, ProviderUndefined.IsRemote())
}

func TestProviderType_IsLocal(t *testing.T) {
	assert.True(t, ProviderLocal.IsLocal())
	assert.False(t, ProviderLDAP.IsLocal())
	assert.False(t, ProviderClient.IsLocal())
	assert.False(t, ProviderClientCredentials.IsLocal())
	assert.False(t, ProviderApplication.IsLocal())
	assert.False(t, ProviderAccessToken.IsLocal())
	assert.False(t, ProviderNone.IsLocal())
	assert.False(t, ProviderDefault.IsLocal())
	assert.False(t, ProviderUndefined.IsLocal())
}

func TestProviderType_SupportsPasscode(t *testing.T) {
	assert.True(t, ProviderLocal.Supports2FA())
	assert.True(t, ProviderLDAP.Supports2FA())
	assert.False(t, ProviderClient.Supports2FA())
	assert.False(t, ProviderClientCredentials.Supports2FA())
	assert.False(t, ProviderApplication.Supports2FA())
	assert.False(t, ProviderAccessToken.Supports2FA())
	assert.False(t, ProviderNone.Supports2FA())
	assert.True(t, ProviderDefault.Supports2FA())
	assert.False(t, ProviderUndefined.Supports2FA())
}

func TestProviderType_IsDefault(t *testing.T) {
	assert.False(t, ProviderLocal.IsDefault())
	assert.False(t, ProviderLDAP.IsDefault())
	assert.False(t, ProviderNone.IsDefault())
	assert.True(t, ProviderDefault.IsDefault())
	assert.True(t, ProviderUndefined.IsDefault())
}

func TestProviderType_IsClient(t *testing.T) {
	assert.False(t, ProviderLocal.IsClient())
	assert.False(t, ProviderLDAP.IsClient())
	assert.False(t, ProviderNone.IsClient())
	assert.False(t, ProviderDefault.IsClient())
	assert.True(t, ProviderClient.IsClient())
	assert.True(t, ProviderClientCredentials.IsClient())
}

func TestProviderType_Equal(t *testing.T) {
	assert.True(t, ProviderClient.Equal("Client"))
	assert.False(t, ProviderLocal.Equal("Client"))
}

func TestProviderType_NotEqual(t *testing.T) {
	assert.False(t, ProviderClient.NotEqual("Client"))
	assert.True(t, ProviderLocal.NotEqual("Client"))
}

func TestProviderType_Pretty(t *testing.T) {
	assert.Equal(t, "Local", ProviderLocal.Pretty())
	assert.Equal(t, "LDAP/AD", ProviderLDAP.Pretty())
	assert.Equal(t, "None", ProviderNone.Pretty())
	assert.Equal(t, "Default", ProviderDefault.Pretty())
	assert.Equal(t, "Default", ProviderUndefined.Pretty())
	assert.Equal(t, "Client", ProviderClient.Pretty())
	assert.Equal(t, "Access Token", ProviderAccessToken.Pretty())
	assert.Equal(t, "Client Credentials", ProviderClientCredentials.Pretty())
}

func TestProvider(t *testing.T) {
	assert.Equal(t, ProviderLocal, Provider("pass"))
	assert.Equal(t, ProviderLDAP, Provider("ad"))
	assert.Equal(t, ProviderDefault, Provider(""))
	assert.Equal(t, ProviderLink, Provider("url"))
	assert.Equal(t, ProviderDefault, Provider("default"))
	assert.Equal(t, ProviderClientCredentials, Provider("oauth2"))
}

func TestProviderType_IsApplication(t *testing.T) {
	assert.True(t, ProviderApplication.IsApplication())
	assert.False(t, ProviderLocal.IsApplication())
}
