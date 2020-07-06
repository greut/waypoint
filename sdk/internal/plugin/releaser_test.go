package plugin

import (
	"testing"

	"github.com/hashicorp/waypoint/sdk/component"
	"github.com/hashicorp/waypoint/sdk/component/mocks"
)

func TestReleaseManagerDynamicFunc_validateAuth(t *testing.T) {
	testDynamicFunc(t, "releasemanager", &mockReleaseManagerAuthenticator{}, func(v, f interface{}) {
		v.(*mockReleaseManagerAuthenticator).Authenticator.On("ValidateAuthFunc").Return(f)
	}, func(raw interface{}) interface{} {
		return raw.(component.Authenticator).ValidateAuthFunc()
	})
}
func TestReleaseManagerDynamicFunc_auth(t *testing.T) {
	testDynamicFunc(t, "releasemanager", &mockReleaseManagerAuthenticator{}, func(v, f interface{}) {
		v.(*mockReleaseManagerAuthenticator).Authenticator.On("AuthFunc").Return(f)
	}, func(raw interface{}) interface{} {
		return raw.(component.Authenticator).AuthFunc()
	})
}

func TestReleaseManagerConfig(t *testing.T) {
	mockV := &mockReleaseManagerConfigurable{}
	testConfigurable(t, "releasemanager", mockV, &mockV.Configurable)
}

type mockReleaseManagerAuthenticator struct {
	mocks.ReleaseManager
	mocks.Authenticator
}

type mockReleaseManagerConfigurable struct {
	mocks.ReleaseManager
	mocks.Configurable
}