// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethereum

import (
	"math/big"

	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func (l *listener) handleErc20DepositedEvent(
	destId msg.ChainId,
	nonce msg.Nonce,
	amount *big.Int,
	resourceId msg.ResourceId,
	recipient []byte,
) (msg.Message, error) {
	l.log.Info("Handling fungible deposit event", "dest", destId, "nonce", nonce)

	return msg.NewFungibleTransfer(
		l.cfg.id,
		destId,
		nonce,
		amount,
		resourceId,
		recipient,
	), nil
}

func (l *listener) handleErc721DepositedEvent(
	destId msg.ChainId,
	nonce msg.Nonce,
	tokenID *big.Int,
	resourceId msg.ResourceId,
	recipient []byte,
	metadata []byte,
) (msg.Message, error) {
	l.log.Info("Handling nonfungible deposit event", "dest", destId, "nonce", nonce)

	return msg.NewNonFungibleTransfer(
		l.cfg.id,
		destId,
		nonce,
		resourceId,
		tokenID,
		recipient,
		metadata,
	), nil
}

func (l *listener) handleGenericDepositedEvent(destId msg.ChainId, nonce msg.Nonce) (msg.Message, error) {
	l.log.Info("Handling generic deposit event")

	record, err := l.genericHandlerContract.GetDepositRecord(&bind.CallOpts{From: l.conn.Keypair().CommonAddress()}, uint64(nonce), uint8(destId))
	if err != nil {
		l.log.Error("Error Unpacking Generic Deposit Record", "err", err)
		return msg.Message{}, nil
	}

	return msg.NewGenericTransfer(
		l.cfg.id,
		destId,
		nonce,
		record.ResourceID,
		record.MetaData[:],
	), nil
}
