package bitcoin

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/ordishs/gocore"
)

var logger = gocore.Log("go-bitcoin")

// A Bitcoind represents a Bitcoind client
type Bitcoind struct {
	client *rpcClient
}

// New return a new bitcoind
func New(host string, port int, user, passwd string, useSSL bool) (*Bitcoind, error) {
	rpcClient, err := newClient(host, port, user, passwd, useSSL)
	if err != nil {
		return nil, err
	}
	return &Bitcoind{rpcClient}, nil
}

// GetConnectionCount returns the number of connections to other nodes.
func (b *Bitcoind) GetConnectionCount() (count uint64, err error) {
	r, err := b.client.call("getconnectioncount", nil)
	if err != nil {
		return 0, err
	}
	count, err = strconv.ParseUint(string(r.Result), 10, 64)
	return
}

// GetBlockchainInfo returns the number of connections to other nodes.
func (b *Bitcoind) GetBlockchainInfo() (info BlockchainInfo, err error) {
	r, err := b.client.call("getblockchaininfo", nil)
	if err != nil {
		return
	}

	if r.Err != nil {
		rr := r.Err.(map[string]interface{})
		err = fmt.Errorf("ERROR %s: %s", rr["code"], rr["message"])
		return
	}

	err = json.Unmarshal(r.Result, &info)
	return
}

// GetNetworkInfo returns the number of connections to other nodes.
func (b *Bitcoind) GetNetworkInfo() (info NetworkInfo, err error) {
	r, err := b.client.call("getnetworkinfo", nil)
	if err != nil {
		return
	}

	if r.Err != nil {
		rr := r.Err.(map[string]interface{})
		err = fmt.Errorf("ERROR %s: %s", rr["code"], rr["message"])
		return
	}

	err = json.Unmarshal(r.Result, &info)
	return
}

// GetNetTotals returns the number of connections to other nodes.
func (b *Bitcoind) GetNetTotals() (totals NetTotals, err error) {
	r, err := b.client.call("getnettotals", nil)
	if err != nil {
		return
	}

	if r.Err != nil {
		rr := r.Err.(map[string]interface{})
		err = fmt.Errorf("ERROR %s: %s", rr["code"], rr["message"])
		return
	}

	err = json.Unmarshal(r.Result, &totals)
	return
}

// GetMiningInfo comment
func (b *Bitcoind) GetMiningInfo() (info MiningInfo, err error) {
	r, err := b.client.call("getmininginfo", nil)
	if err != nil {
		return
	}

	if r.Err != nil {
		rr := r.Err.(map[string]interface{})
		err = fmt.Errorf("ERROR %s: %s", rr["code"], rr["message"])
		return
	}

	err = json.Unmarshal(r.Result, &info)
	return
}

// Uptime returns the number of connections to other nodes.
func (b *Bitcoind) Uptime() (uptime uint64, err error) {
	r, err := b.client.call("uptime", nil)
	if err != nil {
		return 0, err
	}
	uptime, err = strconv.ParseUint(string(r.Result), 10, 64)
	return
}

// GetPeerInfo returns the number of connections to other nodes.
func (b *Bitcoind) GetPeerInfo() (info PeerInfo, err error) {
	r, err := b.client.call("getpeerinfo", nil)
	if err != nil {
		return
	}

	if r.Err != nil {
		rr := r.Err.(map[string]interface{})
		err = fmt.Errorf("ERROR %s: %s", rr["code"], rr["message"])
		return
	}

	err = json.Unmarshal(r.Result, &info)
	return
}

// GetRawMempool returns the number of connections to other nodes.
func (b *Bitcoind) GetRawMempool() (raw RawMemPool, err error) {
	p := []interface{}{false}
	r, err := b.client.call("getrawmempool", p)
	if err != nil {
		return
	}

	if r.Err != nil {
		rr := r.Err.(map[string]interface{})
		err = fmt.Errorf("ERROR %s: %s", rr["code"], rr["message"])
		return
	}

	err = json.Unmarshal(r.Result, &raw)
	return
}

// GetChainTxStats returns the number of connections to other nodes.
func (b *Bitcoind) GetChainTxStats(blockcount int) (stats ChainTXStats, err error) {
	p := []interface{}{blockcount}
	r, err := b.client.call("getchaintxstats", p)
	if err != nil {
		return
	}

	if r.Err != nil {
		rr := r.Err.(map[string]interface{})
		err = fmt.Errorf("ERROR %s: %s", rr["code"], rr["message"])
		return
	}

	err = json.Unmarshal(r.Result, &stats)
	return
}

// ValidateAddress returns the number of connections to other nodes.
func (b *Bitcoind) ValidateAddress(address string) (addr Address, err error) {
	p := []interface{}{address}
	r, err := b.client.call("validateaddress", p)
	if err != nil {
		return
	}

	if r.Err != nil {
		rr := r.Err.(map[string]interface{})
		err = fmt.Errorf("ERROR %s: %s", rr["code"], rr["message"])
		return
	}

	err = json.Unmarshal(r.Result, &addr)
	return
}

// GetHelp returns the number of connections to other nodes.
func (b *Bitcoind) GetHelp() (j []byte, err error) {
	r, err := b.client.call("help", nil)
	if err != nil {
		return
	}
	j, err = json.Marshal(r.Result)

	return
}

// GetBestBlockHash comment
func (b *Bitcoind) GetBestBlockHash() (hash string, err error) {
	r, err := b.client.call("getbestblockhash", nil)
	if err != nil {
		return "", err
	}
	json.Unmarshal(r.Result, &hash)
	return
}

// GetBlockHash comment
func (b *Bitcoind) GetBlockHash(blockHeight int) (blockHash string, err error) {
	p := []interface{}{blockHeight}
	r, err := b.client.call("getblockhash", p)
	if err != nil {
		return "", err
	}
	json.Unmarshal(r.Result, &blockHash)
	return
}

// SendRawTransaction comment
func (b *Bitcoind) SendRawTransaction(hex string) (txid string, err error) {
	r, err := b.client.call("sendrawtransaction", []interface{}{hex})
	if err != nil {
		return "", err
	}
	json.Unmarshal(r.Result, &txid)
	return
}

// GetBlock returns information about the block with the given hash.
func (b *Bitcoind) GetBlock(blockHash string) (block Block, err error) {
	r, err := b.client.call("getblock", []string{blockHash})

	if err != nil {
		return
	}

	if r.Err != nil {
		rr := r.Err.(map[string]interface{})
		err = fmt.Errorf("ERROR %s: %s", rr["code"], rr["message"])
		return
	}

	err = json.Unmarshal(r.Result, &block)
	return
}

// GetRawTransaction returns raw transaction representation for given transaction id.
func (b *Bitcoind) GetRawTransaction(txID string) (rawTx RawTransaction, err error) {
	r, err := b.client.call("getrawtransaction", []interface{}{txID, 1})
	if err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &rawTx)
	return
}

// GetRawTransactionHex returns raw transaction representation for given transaction id.
func (b *Bitcoind) GetRawTransactionHex(txID string) (rawTx string, err error) {
	r, err := b.client.call("getrawtransaction", []interface{}{txID, 0})
	if err != nil {
		return
	}

	err = json.Unmarshal(r.Result, &rawTx)
	return
}

// GetBlockTemplate comment
func (b *Bitcoind) GetBlockTemplate() (template *BlockTemplate, err error) {
	params := gbtParams{
		Mode:         "",
		Capabilities: []string{},
		Rules:        []string{"segwit"},
	}

	r, err := b.client.call("getblocktemplate", []gbtParams{params})
	if err != nil {
		return nil, err
	}

	json.Unmarshal(r.Result, &template)
	return
}

// GetMiningCandidate comment
func (b *Bitcoind) GetMiningCandidate() (template *MiningCandidate, err error) {

	r, err := b.client.call("getminingcandidate", nil)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(r.Result, &template)
	return
}

// SubmitBlock comment
func (b *Bitcoind) SubmitBlock(hexData string) (result string, err error) {
	r, err := b.client.call("submitblock", []interface{}{hexData})
	if err != nil || r.Err != nil || string(r.Result) != "null" {
		msg := fmt.Sprintf("******* BLOCK SUBMIT FAILED with error: %+v and result: %s\n", err, string(r.Result))
		logger.Warnf("%s", msg)
		return "", errors.New(msg)
	}

	logger.Infof("******* BLOCK SUBMITTED SUCCESS: %s\n", string(r.Result))
	return string(r.Result), nil
}

// SubmitMiningSolution comment
func (b *Bitcoind) SubmitMiningSolution(miningCandidateID string, nonce uint32, coinbase string, time uint32, version uint32) (result string, err error) {
	params := submitMiningSolutionParams{
		ID:       miningCandidateID,
		Nonce:    nonce,
		Coinbase: coinbase,
		Time:     time,
		Version:  version,
	}

	logger.Infof("Sending submitminingsolution to bitcoin: %+v", params)

	r, err := b.client.call("submitminingsolution", []interface{}{params})
	if (err != nil && err.Error() != "") || r.Err != nil || (string(r.Result) != "null" && string(r.Result) != "true") {
		msg := fmt.Sprintf("******* BLOCK SUBMIT FAILED with error: %+v and result: %s\n", err, string(r.Result))
		logger.Warnf("%s", msg)
		return "", errors.New(msg)
	}

	logger.Infof("******* BLOCK SUBMITTED SUCCESS.")
	return string(r.Result), nil
}