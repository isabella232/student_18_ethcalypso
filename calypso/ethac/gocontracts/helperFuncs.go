package gocontracts

import (
	"crypto/ecdsa"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func DeployEmptyCalypso(privateKey *ecdsa.PrivateKey, client *ethclient.Client, owners []common.Address) (*common.Address, error) {
	rrHolderAddr, _, _, e := ServiceDeployReadRequestHolder(privateKey, client)
	if e != nil {
		return nil, e
	}
	wrHolderAddr, _, _, e := ServiceDeployWriteRequestHolder(privateKey, client)
	if e != nil {
		return nil, e
	}
	ownersAddress, _, _, e := ServiceDeployOwners(privateKey, client)
	if e != nil {
		return nil, e
	}
	calAddr, _, _, e := ServiceCalypso(privateKey, client, ownersAddress, wrHolderAddr, rrHolderAddr)
	if e != nil {
		return nil, e
	}
	for _, owner := range owners {
		_, e = ServiceUpdatOwners(owner, calAddr, client, privateKey)
		if e != nil {
			return nil, errors.New("Could not update the policy")
		}
	}
	return &calAddr, e
}
