// Copyright (c) 2016 arithmetric
// Based on btcrpcclient by the btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package zcashrpcclient

import (
	"encoding/json"

	"github.com/bitbandi/zcashrpcclient/zcashjson"
	"github.com/btcsuite/btcutil"
)

// FutureListAllTransactionsResult is a future promise to deliver the result
// of a ListAllTransactionsAsync, ListAllTransactionsMinConfAsync, or
// ListAllTransactionsIncludeEmptyAsync RPC invocation (or an applicable
// error).
type FutureListAllTransactionsResult chan *response

// Receive waits for the response promised by the future and returns a list of
// balances by account.
func (r FutureListAllTransactionsResult) Receive() ([]zcashjson.ListAllTransactionsResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal as an array of listreceivedbyaccount result objects.
	var received []zcashjson.ListAllTransactionsResult
	err = json.Unmarshal(res, &received)
	if err != nil {
		return nil, err
	}

	return received, nil
}

// ListAllTransactionsAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
// See ListAllTransactions for the blocking version and more details.
func (c *Client) ListAllTransactionsAsync() FutureListAllTransactionsResult {
	cmd := zcashjson.NewListAllTransactionsCmd(nil, nil, nil)
	return c.sendCmd(cmd)
}

// ListAllTransactions lists balances by account using the default number
// of minimum confirmations and including accounts that haven't received any
// payments.
//
// See ListAllTransactionsMinConf to override the minimum number of
// confirmations and ListAllTransactionsIncludeEmpty to filter accounts that
// haven't received any payments from the results.
func (c *Client) ListAllTransactions() ([]zcashjson.ListAllTransactionsResult, error) {
	return c.ListAllTransactionsAsync().Receive()
}

// ListAllTransactionsMinConfAsync returns an instance of a type that can be
// used to get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
// See ListAllTransactionsMinConf for the blocking version and more details.
func (c *Client) ListAllTransactionsMinConfAsync(minConfirms int) FutureListAllTransactionsResult {
	cmd := zcashjson.NewListAllTransactionsCmd(&minConfirms, nil, nil)
	return c.sendCmd(cmd)
}

// ListAllTransactionsMinConf lists balances by account using the specified
// number of minimum confirmations not including accounts that haven't received
// any payments.
//
// See ListAllTransactions to use the default minimum number of confirmations
// and ListAllTransactionsIncludeEmpty to also include accounts that haven't
// received any payments in the results.
func (c *Client) ListAllTransactionsMinConf(minConfirms int) ([]zcashjson.ListAllTransactionsResult, error) {
	return c.ListAllTransactionsMinConfAsync(minConfirms).Receive()
}

// ListAllTransactionsIncludeEmptyAsync returns an instance of a type that can
// be used to get the result of the RPC at some future time by invoking the
// Receive function on the returned instance.
//
// See ListAllTransactionsIncludeEmpty for the blocking version and more details.
func (c *Client) ListAllTransactionsIncludeEmptyAsync(minConfirms int, includeEmpty bool) FutureListAllTransactionsResult {
	cmd := zcashjson.NewListAllTransactionsCmd(&minConfirms, &includeEmpty,
		nil)
	return c.sendCmd(cmd)
}

// ListAllTransactionsIncludeEmpty lists balances by account using the
// specified number of minimum confirmations and including accounts that
// haven't received any payments depending on specified flag.
//
// See ListAllTransactions and ListAllTransactionsMinConf to use defaults.
func (c *Client) ListAllTransactionsIncludeEmpty(minConfirms int, includeEmpty bool) ([]zcashjson.ListAllTransactionsResult, error) {
	return c.ListAllTransactionsIncludeEmptyAsync(minConfirms,
		includeEmpty).Receive()
}


// ***************************
// Operation Listing Functions
// ***************************

// FutureZGetOperationResultResult is a future promise to deliver the result
// of a ZGetOperationResultAsync RPC invocation (or an applicable error).
type FutureZGetOperationResultResult chan *response

// Receive waits for the response promised by the future and returns detailed
// information about a wallet transaction.
func (r FutureZGetOperationResultResult) Receive() ([]zcashjson.ZGetOperationStatusResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a z_getoperationresult result object
	var results []zcashjson.ZGetOperationStatusResult
	err = json.Unmarshal(res, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

// ZGetOperationResultAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
// See ZGetOperationResult for the blocking version and more details.
func (c *Client) ZGetOperationResultAsync() FutureZGetOperationResultResult {
	cmd := zcashjson.NewZGetOperationResultCmd(nil)
	return c.sendCmd(cmd)
}

// ZGetOperationResult returns detailed information about a wallet transaction.
func (c *Client) ZGetOperationResult() ([]zcashjson.ZGetOperationStatusResult, error) {
	return c.ZGetOperationResultAsync().Receive()
}

// FutureZGetOperationStatusResult is a future promise to deliver the result
// of a ZGetOperationStatusAsync RPC invocation (or an applicable error).
type FutureZGetOperationStatusResult chan *response

// Receive waits for the response promised by the future and returns detailed
// information about a wallet transaction.
func (r FutureZGetOperationStatusResult) Receive() ([]zcashjson.ZGetOperationStatusResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a z_getoperationstatus result object
	var results []zcashjson.ZGetOperationStatusResult
	err = json.Unmarshal(res, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

// ZGetOperationStatusAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
// See ZGetOperationStatus for the blocking version and more details.
func (c *Client) ZGetOperationStatusAsync() FutureZGetOperationStatusResult {
	cmd := zcashjson.NewZGetOperationStatusCmd(nil)
	return c.sendCmd(cmd)
}

// ZGetOperationStatus returns detailed information about a wallet transaction.
func (c *Client) ZGetOperationStatus() ([]zcashjson.ZGetOperationStatusResult, error) {
	return c.ZGetOperationStatusAsync().Receive()
}

// FutureZListOperationIdsResult is a future promise to deliver the result of a
// ZListOperationIdsAsync RPC invocation (or an applicable error).
type FutureZListOperationIdsResult chan *response

// Receive waits for the response promised by the future and returns a list of
// the most recent transactions.
func (r FutureZListOperationIdsResult) Receive() ([]string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as an array of result strings.
	var operationIds []string
	err = json.Unmarshal(res, &operationIds)
	if err != nil {
		return nil, err
	}

	return operationIds, nil
}

// ZListOperationIdsAsync returns an instance of a type that can be used to get
// the result of the RPC at some future time by invoking the Receive function on
// the returned instance.
//
// See ZListOperationIds for the blocking version and more details.
func (c *Client) ZListOperationIdsAsync() FutureZListOperationIdsResult {
	cmd := zcashjson.NewZListOperationIdsCmd(nil)
	return c.sendCmd(cmd)
}

// ZListOperationIds returns a list of the most recent transactions.
func (c *Client) ZListOperationIds() ([]string, error) {
	return c.ZListOperationIdsAsync().Receive()
}

// **************************
// Transaction Send Functions
// **************************

// FutureZSendManyResult is a future promise to deliver the result of a
// ZSendManyAsync RPC invocation (or an applicable error).
type FutureZSendManyResult chan *response

// Receive waits for the response promised by the future and returns the
// operation ID.
func (r FutureZSendManyResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	// Unmarshal result as an array of result strings.
	var operationID string
	err = json.Unmarshal(res, &operationID)
	if err != nil {
		return "", err
	}

	return operationID, nil
}

// ZSendManyAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See ZSendMany for the blocking version and more details.
func (c *Client) ZSendManyAsync(fromAccount string, amounts []zcashjson.ZSendManyEntry) FutureZSendManyResult {
	cmd := zcashjson.NewZSendManyCmd(fromAccount, amounts, nil)
	return c.sendCmd(cmd)
}

// ZSendMany sends multiple amounts to multiple addresses using the provided
// account as a source of funds in a single transaction.  Only funds with the
// default number of minimum confirmations will be used.
func (c *Client) ZSendMany(fromAccount string, amounts []zcashjson.ZSendManyEntry) (string, error) {
	return c.ZSendManyAsync(fromAccount, amounts).Receive()
}

// *************************
// Address/Account Functions
// *************************

// FutureZGetNewAddressResult is a future promise to deliver the result of a
// ZGetNewAddressAsync RPC invocation (or an applicable error).
type FutureZGetNewAddressResult chan *response

// Receive waits for the response promised by the future and returns a new
// address.
func (r FutureZGetNewAddressResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	// Unmarshal result as a string.
	var zaddr string
	err = json.Unmarshal(res, &zaddr)
	if err != nil {
		return "", err
	}

	return zaddr, nil
}

// ZGetNewAddressAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See ZGetNewAddress for the blocking version and more details.
func (c *Client) ZGetNewAddressAsync() FutureZGetNewAddressResult {
	cmd := zcashjson.NewZGetNewAddressCmd()
	return c.sendCmd(cmd)
}

// ZGetNewAddress returns a new address.
func (c *Client) ZGetNewAddress() (string, error) {
	return c.ZGetNewAddressAsync().Receive()
}

// ************************
// Amount/Balance Functions
// ************************

// FutureZListAddressesResult is a future promise to deliver the result of a
// ZListAddressesAsync RPC invocation (or an applicable error).
type FutureZListAddressesResult chan *response

// Receive waits for the response promised by the future and returns returns a
// map of account names and their associated balances.
func (r FutureZListAddressesResult) Receive() ([]string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a json object.
	var zaddrs []string
	err = json.Unmarshal(res, &zaddrs)
	if err != nil {
		return nil, err
	}

	return zaddrs, nil
}

// ZListAddressesAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See ZListAddresses for the blocking version and more details.
func (c *Client) ZListAddressesAsync() FutureZListAddressesResult {
	cmd := zcashjson.NewZListAddressesCmd(nil)
	return c.sendCmd(cmd)
}

// ZListAddresses returns a map of account names and their associated balances
// using the default number of minimum confirmations.
func (c *Client) ZListAddresses() ([]string, error) {
	return c.ZListAddressesAsync().Receive()
}

// FutureZGetBalanceResult is a future promise to deliver the result of a
// ZGetBalanceAsync or ZGetBalanceMinConfAsync RPC invocation (or an applicable
// error).
type FutureZGetBalanceResult chan *response

// Receive waits for the response promised by the future and returns the
// available balance from the server for the specified account.
func (r FutureZGetBalanceResult) Receive() (btcutil.Amount, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return 0, err
	}

	// Unmarshal result as a floating point number.
	var balance float64
	err = json.Unmarshal(res, &balance)
	if err != nil {
		return 0, err
	}

	amount, err := btcutil.NewAmount(balance)
	if err != nil {
		return 0, err
	}

	return amount, nil
}

// ZGetBalanceAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See ZGetBalance for the blocking version and more details.
func (c *Client) ZGetBalanceAsync(address string) FutureZGetBalanceResult {
	cmd := zcashjson.NewZGetBalanceCmd(&address, nil)
	return c.sendCmd(cmd)
}

// ZGetBalance returns the available balance from the server for the specified
// account using the default number of minimum confirmations.  The account may
// be "*" for all accounts.
func (c *Client) ZGetBalance(address string) (btcutil.Amount, error) {
	return c.ZGetBalanceAsync(address).Receive()
}

// ZGetBalanceMinConfAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See ZGetBalanceMinConf for the blocking version and more details.
func (c *Client) ZGetBalanceMinConfAsync(address string, minConf int) FutureZGetBalanceResult {
	cmd := zcashjson.NewZGetBalanceCmd(&address, &minConf)
	return c.sendCmd(cmd)
}

// ZGetBalanceMinConf returns the available balance from the server for the specified
// account using the number of confirmations.  The account may
// be "*" for all accounts.
func (c *Client) ZGetBalanceMinConf(address string, minConf int) (btcutil.Amount, error) {
	return c.ZGetBalanceMinConfAsync(address, minConf).Receive()
}

// FutureZGetTotalBalanceResult is a future promise to deliver the result of a
// ZGetTotalBalanceAsync or ZGetTotalBalanceMinConfAsync RPC invocation (or an applicable error).
type FutureZGetTotalBalanceResult chan *response

// Receive waits for the response promised by the future and returns the
// available balance from the server for the specified account.
func (r FutureZGetTotalBalanceResult) Receive() (*zcashjson.ZGetTotalBalanceResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a floating point number.
	var totals zcashjson.ZGetTotalBalanceResult
	err = json.Unmarshal(res, &totals)
	if err != nil {
		return nil, err
	}

	return &totals, nil
}

// ZGetTotalBalanceAsync returns an instance of a type that can be used to get
// the result of the RPC at some future time by invoking the Receive function on
// the returned instance.
//
// See ZGetTotalBalance for the blocking version and more details.
func (c *Client) ZGetTotalBalanceAsync() FutureZGetTotalBalanceResult {
	cmd := zcashjson.NewZGetTotalBalanceCmd(nil)
	return c.sendCmd(cmd)
}

// ZGetTotalBalance returns the transparent, private, and total balances for
// all addresses in the wallet.
func (c *Client) ZGetTotalBalance() (*zcashjson.ZGetTotalBalanceResult, error) {
	return c.ZGetTotalBalanceAsync().Receive()
}

// ZGetTotalBalanceMinConfAsync returns an instance of a type that can be used to get
// the result of the RPC at some future time by invoking the Receive function on
// the returned instance.
//
// See ZGetTotalBalanceMinConf for the blocking version and more details.
func (c *Client) ZGetTotalBalanceMinConfAsync(minConf int) FutureZGetTotalBalanceResult {
	cmd := zcashjson.NewZGetTotalBalanceCmd(&minConf)
	return c.sendCmd(cmd)
}

// ZGetTotalBalanceMinConf returns the transparent, private, and total balances for
// all addresses in the wallet.
func (c *Client) ZGetTotalBalanceMinConf(minConf int) (*zcashjson.ZGetTotalBalanceResult, error) {
	return c.ZGetTotalBalanceMinConfAsync(minConf).Receive()
}

// FutureZListReceivedByAddressResult is a future promise to deliver the result
// of a ZListReceivedByAddressAsync RPC invocation (or an applicable error).
type FutureZListReceivedByAddressResult chan *response

// Receive waits for the response promised by the future and returns a list of
// balances by address.
func (r FutureZListReceivedByAddressResult) Receive() ([]zcashjson.ZListReceivedByAddressResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal as an array of z_listreceivedbyaddress result objects.
	var received []zcashjson.ZListReceivedByAddressResult
	err = json.Unmarshal(res, &received)
	if err != nil {
		return nil, err
	}

	return received, nil
}

// ZListReceivedByAddressAsync returns an instance of a type that can be used to
// get the result of the RPC at some future time by invoking the Receive
// function on the returned instance.
//
// See ZListReceivedByAddress for the blocking version and more details.
func (c *Client) ZListReceivedByAddressAsync(address string) FutureZListReceivedByAddressResult {
	cmd := zcashjson.NewZListReceivedByAddressCmd(address, nil)
	return c.sendCmd(cmd)
}

// ZListReceivedByAddress lists balances by address using the default number
// of minimum confirmations not including addresses that haven't received any
// payments or watching only addresses.
func (c *Client) ZListReceivedByAddress(address string) ([]zcashjson.ZListReceivedByAddressResult, error) {
	return c.ZListReceivedByAddressAsync(address).Receive()
}

// ***********************
// Export/Import Functions
// ***********************

// FutureZExportKeyResult is a future promise to deliver the result of a
// ZExportKeyAsync RPC invocation (or an applicable error).
type FutureZExportKeyResult chan *response

// Receive waits for the response promised by the future and returns the private
// key corresponding to the passed address
func (r FutureZExportKeyResult) Receive() (string, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return "", err
	}

	// Unmarshal result as a string.
	var key string
	err = json.Unmarshal(res, &key)
	if err != nil {
		return "", err
	}

	return key, nil
}

// ZExportKeyAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See ZExportKey for the blocking version and more details.
func (c *Client) ZExportKeyAsync(address string) FutureZExportKeyResult {
	cmd := zcashjson.NewZExportKeyCmd(address)
	return c.sendCmd(cmd)
}

// ZExportKey gets the private key corresponding to the passed address.
func (c *Client) ZExportKey(address string) (string, error) {
	return c.ZExportKeyAsync(address).Receive()
}

// FutureZExportWalletResult is a future promise to deliver the result of a
// ZExportWalletAsync RPC invocation (or an applicable error).
type FutureZExportWalletResult chan *response

// Receive waits for the response promised by the future and returns the private
// key corresponding to the passed address
func (r FutureZExportWalletResult) Receive() error {
	_, err := receiveFuture(r)
	if err != nil {
		return err
	}

	return nil
}

// ZExportWalletAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See ZExportWallet for the blocking version and more details.
func (c *Client) ZExportWalletAsync(filename string) FutureZExportWalletResult {
	cmd := zcashjson.NewZExportWalletCmd(filename)
	return c.sendCmd(cmd)
}

// ZExportWallet gets the private key corresponding to the passed address.
func (c *Client) ZExportWallet(filename string) error {
	return c.ZExportWalletAsync(filename).Receive()
}

// FutureZImportKeyResult is a future promise to deliver the result of an
// ZImportKeyAsync RPC invocation (or an applicable error).
type FutureZImportKeyResult chan *response

// Receive waits for the response promised by the future and returns the result
// of importing the passed private key.
func (r FutureZImportKeyResult) Receive() error {
	_, err := receiveFuture(r)
	if err != nil {
		return err
	}

	return nil
}

// ZImportKeyAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See ZImportKey for the blocking version and more details.
func (c *Client) ZImportKeyAsync(key string, rescan *bool) FutureZImportKeyResult {
	cmd := zcashjson.NewZImportKeyCmd(key, rescan)
	return c.sendCmd(cmd)
}

// ZImportKey imports the passed private key.
func (c *Client) ZImportKey(key string, rescan *bool) error {
	return c.ZImportKeyAsync(key, rescan).Receive()
}

// FutureZImportWalletResult is a future promise to deliver the result of a
// ZImportWalletAsync RPC invocation (or an applicable error).
type FutureZImportWalletResult chan *response

// Receive waits for the response promised by the future and returns the private
// key corresponding to the passed address
func (r FutureZImportWalletResult) Receive() error {
	_, err := receiveFuture(r)
	if err != nil {
		return err
	}

	return nil
}

// ZImportWalletAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See ZImportWallet for the blocking version and more details.
func (c *Client) ZImportWalletAsync(filename string) FutureZImportWalletResult {
	cmd := zcashjson.NewZImportWalletCmd(filename)
	return c.sendCmd(cmd)
}

// ZImportWallet gets the private key corresponding to the passed address.
func (c *Client) ZImportWallet(filename string) error {
	return c.ZImportWalletAsync(filename).Receive()
}
