// Copyright (c) 2016 arithmetric
// Based on btcd by the btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package zcashjson

import (
	"github.com/btcsuite/btcd/btcjson"
)

// ListAllTransactionsResult models the data from the listalltransactions command.
type ListAllTransactionsResult struct {
	btcjson.TxRawResult
	Amount          float64                       `json:"amount"`
	Fee             float64                       `json:"fee,omitempty"`
}

// ZOperationStatusError models the error data in ZGetOperationStatusResult.
type ZOperationStatusError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ZOperationAmounts models the amounts data in ZOperationParams.
type ZOperationAmounts struct {
	Address string  `json:"address"`
	Amount  float64 `json:"amount"`
	Memo    string  `json:"memo,omitempty"`
}

// ZOperationParams models the params data in ZGetOperationStatusResult.
type ZOperationParams struct {
	From    string              `json:"fromaddress"`
	MinConf int                 `json:"minconf"`
	Amounts []ZOperationAmounts `json:"amounts"`
	Fee     float64             `json:"fee"`
}

// ZGetOperationStatusResult models the data from the z_getoperationresult and
// z_getoperationstatus commands.
type ZGetOperationStatusResult struct {
	Id            string                `json:"id"`
	Status        string                `json:"status"`
	CreationTime  int                   `json:"creation_time"`
	Result        map[string]string     `json:"result"`
	Error         ZOperationStatusError `json:"error"`
	Params        ZOperationParams      `json:"params"`
	ExecutionSecs float64               `json:"execution_secs"`
}

// ZGetTotalBalanceResult models the data from the z_gettotalbalance command.
type ZGetTotalBalanceResult struct {
	Transparent float64 `json:"transparent,string"`
	Private     float64 `json:"private,string"`
	Total       float64 `json:"total,string"`
}

// ZListReceivedByAddressResult models the data from the z_listreceivedbyaddress
// command.
type ZListReceivedByAddressResult struct {
	TxID   string  `json:"txid"`
	Amount float64 `json:"amount"`
	Memo   string  `json:"memo"`
}
