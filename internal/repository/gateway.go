//
// Copyright (c) 2018 Tencent
//
// SPDX-License-Identifier: Apache-2.0
//

package repository

import (
	"github.com/edgexfoundry/edgex-ui-go/internal/domain"
	"github.com/edgexfoundry/edgex-ui-go/internal/repository/mm"
	"github.com/edgexfoundry/edgex-ui-go/internal/repository/mongo"
)

type GatewayRepos interface {
	Select(id string) (domain.Gateway, error)
	SelectAll() ([]domain.Gateway, error)
	Exists(id string) (bool, error)
	Insert(gateway *domain.Gateway) (string, error)
	Update(gateway domain.Gateway) error
	Delete(id string) error
}

func GetGatewayRepos() GatewayRepos {
	if mongo.DS.S == nil {
		return GatewayRepos(&mm.GatewayRepository{})
	} else {
		return GatewayRepos(&mongo.GatewayMongoRepository{})
	}
}
