// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package charms_test

import (
	"github.com/golang/mock/gomock"
	charm "github.com/juju/charm/v8"
	gc "gopkg.in/check.v1"

	basemocks "github.com/juju/juju/api/base/mocks"
	apicommoncharms "github.com/juju/juju/api/common/charms"
	"github.com/juju/juju/apiserver/params"
	coretesting "github.com/juju/juju/testing"
)

type charmsMockSuite struct {
	coretesting.BaseSuite
	charmsCommonClient *apicommoncharms.CharmsClient
}

var _ = gc.Suite(&charmsMockSuite{})

func (s *charmsMockSuite) TestCharmInfo(c *gc.C) {
	ctrl := gomock.NewController(c)
	defer ctrl.Finish()

	mockFacadeCaller := basemocks.NewMockFacadeCaller(ctrl)

	url := "local:quantal/dummy-1"
	args := params.CharmURL{URL: url}
	info := new(params.Charm)

	params := params.Charm{
		Revision: 1,
		URL:      url,
		Config: map[string]params.CharmOption{
			"config": {
				Type:        "type",
				Description: "config-type option",
			},
		},
		LXDProfile: &params.CharmLXDProfile{
			Description: "LXDProfile",
			Devices: map[string]map[string]string{
				"tun": {
					"path": "/dev/net/tun",
					"type": "unix-char",
				},
			},
		},
	}

	mockFacadeCaller.EXPECT().FacadeCall("CharmInfo", args, info).SetArg(2, params).Return(nil)

	client := apicommoncharms.NewCharmsClient(mockFacadeCaller)
	got, err := client.CharmInfo(url)
	c.Assert(err, gc.IsNil)

	want := &apicommoncharms.CharmInfo{
		Revision: 1,
		URL:      url,
		Config: &charm.Config{
			Options: map[string]charm.Option{
				"config": {
					Type:        "type",
					Description: "config-type option",
				},
			},
		},
		LXDProfile: &charm.LXDProfile{
			Description: "LXDProfile",
			Config:      map[string]string{},
			Devices: map[string]map[string]string{
				"tun": {
					"path": "/dev/net/tun",
					"type": "unix-char",
				},
			},
		},
	}
	c.Assert(got, gc.DeepEquals, want)
}