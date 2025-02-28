/*
 * Copyright (c) 2017-2020 The qitmeer developers
 */

package consensus

import (
	"github.com/Qitmeer/qng/core/types"
	"github.com/libp2p/go-libp2p/core/peer"
)

// Notify interface manage message announce & relay & notification between mempool, websocket, gbt long pull
// and rpc server.
type Notify interface {
	AnnounceNewTransactions(newTxs []*types.TxDesc, filters []peer.ID)
	RelayInventory(data interface{}, filters []peer.ID)
	BroadcastMessage(data interface{})
	TransactionConfirmed(tx *types.Tx)
	AddRebroadcastInventory(newTxs []*types.TxDesc)
}
