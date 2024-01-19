package entity

import (
	"testing"

	"github.com/photoprism/photoprism/internal/form"
	"github.com/photoprism/photoprism/pkg/authn"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	m := NewClient()
	assert.Equal(t, "", m.AuthScope)
	assert.Equal(t, m.AuthScope, m.Scope())
	m.SetScope(" metrics WEBdav!")
	assert.Equal(t, "metrics webdav", m.AuthScope)
	assert.Equal(t, m.AuthScope, m.Scope())
}

func TestFindClient(t *testing.T) {
	t.Run("Alice", func(t *testing.T) {
		expected := ClientFixtures.Get("alice")

		m := FindClientByUID("cs5gfen1bgxz7s9i")

		if m == nil {
			t.Fatal("result should not be nil")
		}

		assert.Equal(t, m.UserUID, UserFixtures.Get("alice").UserUID)
		assert.Equal(t, expected.ClientUID, m.UID())
		assert.NotEmpty(t, m.CreatedAt)
		assert.NotEmpty(t, m.UpdatedAt)
	})
	t.Run("Bob", func(t *testing.T) {
		expected := ClientFixtures.Get("bob")

		m := FindClientByUID("cs5gfsvbd7ejzn8m")

		if m == nil {
			t.Fatal("result should not be nil")
		}

		assert.Equal(t, m.UserUID, UserFixtures.Get("bob").UserUID)
		assert.Equal(t, expected.ClientUID, m.UID())
		assert.NotEmpty(t, m.CreatedAt)
		assert.NotEmpty(t, m.UpdatedAt)
	})
	t.Run("Metrics", func(t *testing.T) {
		expected := ClientFixtures.Get("metrics")

		m := FindClientByUID("cs5cpu17n6gj2qo5")

		if m == nil {
			t.Fatal("result should not be nil")
		}

		assert.Empty(t, m.UserUID)
		assert.Equal(t, expected.ClientUID, m.UID())
		assert.NotEmpty(t, m.CreatedAt)
		assert.NotEmpty(t, m.UpdatedAt)
	})
	t.Run("Invalid", func(t *testing.T) {
		m := FindClientByUID("123")
		assert.Nil(t, m)
	})
}

func TestClient_User(t *testing.T) {
	t.Run("Alice", func(t *testing.T) {
		alice := ClientFixtures.Get("alice")

		m := alice.User()

		if m == nil {
			t.Fatal("result should not be nil")
		}

		assert.Equal(t, "alice", m.UserName)
		assert.Equal(t, "uqxetse3cy5eo9z2", m.UserUID)
		assert.Equal(t, "admin", m.UserRole)

	})
	t.Run("Metrics", func(t *testing.T) {
		metrics := ClientFixtures.Get("metrics")

		m := metrics.User()

		if m == nil {
			t.Fatal("result should not be nil")
		}

		assert.Empty(t, m.UserName)
		assert.Empty(t, m.UserUID)
		assert.Empty(t, m.UserRole)

	})
	t.Run("UnknownUser", func(t *testing.T) {
		c := Client{ClientName: "test",
			UserUID: "123",
		}

		m := c.User()

		if m == nil {
			t.Fatal("result should not be nil")
		}

		assert.Empty(t, m.UserName)
		assert.Empty(t, m.UserUID)
		assert.Empty(t, m.UserRole)
	})
	t.Run("Bob", func(t *testing.T) {
		c := Client{ClientName: "bob",
			UserUID: "uqxc08w3d0ej2283",
		}

		m := c.User()

		if m == nil {
			t.Fatal("result should not be nil")
		}

		assert.Equal(t, "bob", m.UserName)
		assert.Equal(t, "uqxc08w3d0ej2283", m.UserUID)
		assert.Equal(t, "admin", m.UserRole)
	})
}

func TestClient_SetUser(t *testing.T) {
	t.Run("john", func(t *testing.T) {
		c := Client{ClientName: "test"}
		u := &User{UserUID: "uqxc08w3d0ej2111", UserName: "john"}

		assert.Empty(t, c.User().UserName)

		m := c.SetUser(u)

		if m == nil {
			t.Fatal("result should not be nil")
		}

		assert.NotEmpty(t, c.User().UserName)
	})
}

func TestClient_Create(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var m = Client{ClientName: "test"}
		if err := m.Create(); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("AlreadyExists", func(t *testing.T) {
		var m = ClientFixtures.Get("alice")
		err := m.Create()
		assert.Error(t, err)
	})
}

func TestClient_Save(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		c := FindClientByUID("cs5cpu17n6gj2aaa")
		assert.Nil(t, c)

		var m = Client{ClientName: "New Client", ClientUID: "cs5cpu17n6gj2aaa"}
		if err := m.Save(); err != nil {
			t.Fatal(err)
		}

		c = FindClientByUID("cs5cpu17n6gj2aaa")

		if c == nil {
			t.Fatal("result should not be nil")
		}
	})
}

func TestClient_Delete(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var m = Client{ClientName: "David", ClientUID: "cs5cpu17n6gj2bbb"}
		if err := m.Save(); err != nil {
			t.Fatal(err)
		}

		err := m.Delete()

		assert.NoError(t, err)
	})
	t.Run("EmptyUID", func(t *testing.T) {
		var m = Client{ClientName: "No UUID"}

		err := m.Delete()

		assert.Error(t, err)
	})
}

func TestClient_Deleted(t *testing.T) {
	assert.False(t, ClientFixtures.Pointer("alice").Deleted())
	assert.True(t, ClientFixtures.Pointer("deleted").Deleted())
}

func TestClient_Updates(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var m = Client{ClientName: "Peter", ClientUID: "cs5cpu17n6gj2ddd"}

		assert.Empty(t, m.AuthScope)

		err := m.Updates(Client{AuthScope: "metrics"})

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "metrics", m.AuthScope)
	})
}

func TestClient_NewSecret(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var m = Client{ClientName: "Anna", ClientUID: "cs5cpu17n6gj2eee"}
		if err := m.Save(); err != nil {
			t.Fatal(err)
		}

		s, err := m.NewSecret()

		if err != nil {
			t.Fatal(err)
		}
		assert.True(t, m.HasSecret(s))
		assert.NotEmpty(t, s)
	})
	t.Run("EmptyUID", func(t *testing.T) {
		var m = Client{ClientName: "No UUID"}

		s, err := m.NewSecret()

		assert.Error(t, err)
		assert.False(t, m.HasSecret(s))
		assert.Empty(t, s)
	})
}

func TestClient_Provider(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		client := NewClient()
		assert.Equal(t, authn.ProviderClientCredentials, client.Provider())
	})
	t.Run("Alice", func(t *testing.T) {
		client := ClientFixtures.Get("alice")
		assert.Equal(t, authn.ProviderClientCredentials, client.Provider())
	})
	t.Run("Bob", func(t *testing.T) {
		client := ClientFixtures.Get("bob")
		assert.Equal(t, authn.ProviderClientCredentials, client.Provider())
	})
}

func TestClient_Method(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		client := NewClient()
		assert.Equal(t, authn.MethodOAuth2, client.Method())
	})
	t.Run("Alice", func(t *testing.T) {
		client := ClientFixtures.Get("alice")
		assert.Equal(t, authn.MethodOAuth2, client.Method())
	})
	t.Run("Bob", func(t *testing.T) {
		client := ClientFixtures.Get("bob")
		assert.Equal(t, authn.MethodOAuth2, client.Method())
	})
}

func TestClient_UpdateLastActive(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var m = Client{ClientName: "Anne", ClientUID: "cs5cpu17n6gj2fff"}
		if err := m.Save(); err != nil {
			t.Fatal(err)
		}

		assert.Empty(t, m.LastActive)

		c := m.UpdateLastActive()

		assert.NotEmpty(t, c.LastActive)
	})
	t.Run("EmptyUID", func(t *testing.T) {
		var m = Client{ClientName: "No UUID"}

		c := m.UpdateLastActive()

		assert.Empty(t, c.LastActive)
	})
}

func TestClient_EnforceAuthTokenLimit(t *testing.T) {
	t.Run("EmptyUID", func(t *testing.T) {
		var m = Client{ClientName: "No UUID"}

		r := m.EnforceAuthTokenLimit()

		assert.Equal(t, r, 0)
	})
	t.Run("NoToken", func(t *testing.T) {
		var m = Client{ClientName: "David", ClientUID: "cs5cpu17n6gj2bbb"}

		r := m.EnforceAuthTokenLimit()

		assert.Equal(t, r, 0)
	})
	t.Run("NegativeTokenLimit", func(t *testing.T) {
		var m = Client{ClientName: "David", ClientUID: "cs5cpu17n6gj2bbb", AuthTokens: -1}

		r := m.EnforceAuthTokenLimit()

		assert.Equal(t, r, 0)
	})
}

func TestClient_HasPassword(t *testing.T) {
	t.Run("Alice", func(t *testing.T) {
		expected := ClientFixtures.Get("alice")

		m := FindClientByUID("cs5gfen1bgxz7s9i")

		if m == nil {
			t.Fatal("result should not be nil")
		}

		assert.Equal(t, expected.ClientUID, m.UID())
		assert.False(t, m.HasSecret("xcCbOrw6I0vcoXzhnOmXhjpVSyFq0l0e"))
		assert.False(t, m.HasSecret("aaCbOrw6I0vcoXzhnOmXhjpVSyFq0l0e"))
		assert.False(t, m.HasSecret(""))
		assert.True(t, m.WrongSecret("xcCbOrw6I0vcoXzhnOmXhjpVSyFq0l0e"))
		assert.True(t, m.WrongSecret("aaCbOrw6I0vcoXzhnOmXhjpVSyFq0l0e"))
		assert.True(t, m.WrongSecret(""))
		assert.NotEmpty(t, m.CreatedAt)
		assert.NotEmpty(t, m.UpdatedAt)
	})
	t.Run("Metrics", func(t *testing.T) {
		expected := ClientFixtures.Get("metrics")

		m := FindClientByUID("cs5cpu17n6gj2qo5")

		if m == nil {
			t.Fatal("result should not be nil")
		}

		assert.Equal(t, expected.ClientUID, m.UID())
		assert.True(t, m.HasSecret("xcCbOrw6I0vcoXzhnOmXhjpVSyFq0l0e"))
		assert.False(t, m.HasSecret("aaCbOrw6I0vcoXzhnOmXhjpVSyFq0l0e"))
		assert.False(t, m.HasSecret(""))
		assert.False(t, m.WrongSecret("xcCbOrw6I0vcoXzhnOmXhjpVSyFq0l0e"))
		assert.True(t, m.WrongSecret("aaCbOrw6I0vcoXzhnOmXhjpVSyFq0l0e"))
		assert.True(t, m.WrongSecret(""))
		assert.NotEmpty(t, m.CreatedAt)
		assert.NotEmpty(t, m.UpdatedAt)
	})
}

func TestClient_Expires(t *testing.T) {
	t.Run("Metrics", func(t *testing.T) {
		m := ClientFixtures.Get("metrics")

		r := m.Expires()

		assert.Equal(t, r.String(), "1h0m0s")
	})
	t.Run("Alice", func(t *testing.T) {
		m := ClientFixtures.Get("alice")

		r := m.Expires()

		assert.Equal(t, r.String(), "24h0m0s")
	})
}

func TestClient_UserInfo(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		assert.Equal(t, "", NewClient().UserInfo())
	})
	t.Run("Alice", func(t *testing.T) {
		assert.Equal(t, "alice", ClientFixtures.Pointer("alice").UserInfo())
	})
	t.Run("Metrics", func(t *testing.T) {
		assert.Equal(t, "", ClientFixtures.Pointer("metrics").UserInfo())
	})
}

func TestClient_AuthInfo(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		assert.Equal(t, "Client Credentials (OAuth2)", NewClient().AuthInfo())
	})
	t.Run("Alice", func(t *testing.T) {
		assert.Equal(t, "Client Credentials (OAuth2)", ClientFixtures.Pointer("alice").AuthInfo())
	})
	t.Run("Metrics", func(t *testing.T) {
		assert.Equal(t, "Client Credentials (OAuth2)", ClientFixtures.Pointer("metrics").AuthInfo())
	})
}

func TestClient_Report(t *testing.T) {
	t.Run("Metrics", func(t *testing.T) {
		m := ClientFixtures.Get("metrics")

		rows, _ := m.Report(true)
		assert.NotEmpty(t, rows)
	})
}

func TestClient_SetFormValues(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var m = Client{ClientName: "Neo", ClientUID: "cs5cpu17n6gj3aab"}

		if err := m.Save(); err != nil {
			t.Fatal(err)
		}

		var values = form.Client{
			ClientName:   "New Name",
			AuthProvider: authn.ProviderClientCredentials.String(),
			AuthMethod:   authn.MethodOAuth2.String(),
			AuthScope:    "test",
			AuthExpires:  4000,
			AuthTokens:   3,
			AuthEnabled:  false,
		}

		c := m.SetFormValues(values)

		assert.Equal(t, "New Name", c.ClientName)
		assert.Equal(t, int64(4000), c.AuthExpires)
		assert.Equal(t, int64(3), c.AuthTokens)
		assert.Equal(t, false, c.AuthEnabled)
	})
	t.Run("Success2", func(t *testing.T) {
		var m = Client{ClientName: "Neo", ClientUID: "cs5cpu17n6gj3aab", AuthTokens: -4}

		if err := m.Save(); err != nil {
			t.Fatal(err)
		}

		var values = form.Client{
			ClientName:   "Annika",
			AuthProvider: authn.ProviderClientCredentials.String(),
			AuthMethod:   authn.MethodOAuth2.String(),
			AuthScope:    "metrics",
			AuthExpires:  -4000,
			AuthTokens:   -5,
			AuthEnabled:  true,
		}

		c := m.SetFormValues(values)

		assert.Equal(t, "Annika", c.ClientName)
		assert.Equal(t, int64(3600), c.AuthExpires)
		assert.Equal(t, int64(-1), c.AuthTokens)
		assert.Equal(t, true, c.AuthEnabled)
	})
	t.Run("Success3", func(t *testing.T) {
		var m = Client{ClientName: "Neo", ClientUID: "cs5cpu17n6gj3aab"}

		if err := m.Save(); err != nil {
			t.Fatal(err)
		}

		var values = form.Client{
			ClientName:   "Friend",
			AuthProvider: authn.ProviderClientCredentials.String(),
			AuthMethod:   authn.MethodOAuth2.String(),
			AuthScope:    "test",
			AuthExpires:  4000000,
			AuthTokens:   3000000000,
			AuthEnabled:  true,
			UserUID:      "uqxqg7i1kperxvu7",
		}

		c := m.SetFormValues(values)

		assert.Equal(t, "Friend", c.ClientName)
		assert.Equal(t, int64(2678400), c.AuthExpires)
		assert.Equal(t, int64(2147483647), c.AuthTokens)
		assert.Equal(t, true, c.AuthEnabled)
	})
}

func TestClient_Validate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		m := Client{
			ClientName:   "test",
			ClientType:   "test",
			AuthProvider: authn.ProviderClientCredentials.String(),
			AuthMethod:   "basic",
			AuthScope:    "all",
		}

		err := m.Validate()

		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("ClientNameEmpty", func(t *testing.T) {
		m := Client{
			ClientName:   "",
			ClientType:   "test",
			AuthProvider: authn.ProviderClientCredentials.String(),
			AuthMethod:   "basic",
			AuthScope:    "all",
		}

		err := m.Validate()

		if err == nil {
			t.Fatal("error expected")
		}
	})
	t.Run("ClientTypeEmpty", func(t *testing.T) {
		m := Client{
			ClientName:   "test",
			ClientType:   "",
			AuthProvider: authn.ProviderClientCredentials.String(),
			AuthMethod:   "basic",
			AuthScope:    "all",
		}

		err := m.Validate()

		if err == nil {
			t.Fatal("error expected")
		}
	})
	t.Run("AuthMethodEmpty", func(t *testing.T) {
		m := Client{
			ClientName:   "test",
			ClientType:   "test",
			AuthProvider: authn.ProviderClientCredentials.String(),
			AuthMethod:   "",
			AuthScope:    "all",
		}

		err := m.Validate()

		if err == nil {
			t.Fatal("error expected")
		}
	})
	t.Run("AuthScopeEmpty", func(t *testing.T) {
		m := Client{
			ClientName:   "test",
			ClientType:   "test",
			AuthProvider: authn.ProviderClientCredentials.String(),
			AuthMethod:   "basic",
			AuthScope:    "",
		}

		err := m.Validate()

		if err == nil {
			t.Fatal("error expected")
		}
	})
}
