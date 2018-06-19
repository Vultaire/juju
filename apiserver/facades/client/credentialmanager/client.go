// Copyright 2018 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package credentialmanager

import (
	"github.com/juju/loggo"

	"github.com/juju/juju/apiserver/common"
	"github.com/juju/juju/apiserver/common/credentialcommon"
	"github.com/juju/juju/apiserver/facade"
	"github.com/juju/juju/apiserver/params"
)

var logger = loggo.GetLogger("juju.api.credentialmanager")

// CredentialManager defines the methods on credentialmanager API endpoint.
type CredentialManager interface {
	InvalidateModelCredential(reason string) (params.ErrorResult, error)
}

type CredentialManagerAPI struct {
	*credentialcommon.CredentialManagerAPI

	resources facade.Resources
}

var _ CredentialManager = (*CredentialManagerAPI)(nil)

// NewCredentialManagerAPI creates a new CredentialManager API endpoint on server-side.
func NewCredentialManagerAPI(ctx facade.Context) (*CredentialManagerAPI, error) {
	return internalNewCredentialManagerAPI(NewStateShim(ctx.State()), ctx.Resources(), ctx.Auth())
}

func internalNewCredentialManagerAPI(backend credentialcommon.StateBackend, resources facade.Resources, authorizer facade.Authorizer) (*CredentialManagerAPI, error) {
	hostAuthTag := authorizer.GetAuthTag()
	if hostAuthTag == nil {
		return nil, common.ErrPerm
	}

	return &CredentialManagerAPI{
		resources:            resources,
		CredentialManagerAPI: credentialcommon.NewCredentialManagerAPI(backend),
	}, nil
}
