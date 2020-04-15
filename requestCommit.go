package main

import (
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/levigross/grequests"
)

type Version struct {
	Block string `json:"block"`
	App   string `json:"app"`
}
type Parts struct {
	Total string `json:"total"`
	Hash  string `json:"hash"`
}
type LastBlockID struct {
	Hash  string `json:"hash"`
	Parts Parts  `json:"parts"`
}
type Header struct {
	Version            Version     `json:"version"`
	ChainID            string      `json:"chain_id"`
	Height             string      `json:"height"`
	Time               time.Time   `json:"time"`
	LastBlockID        LastBlockID `json:"last_block_id"`
	LastCommitHash     string      `json:"last_commit_hash"`
	DataHash           string      `json:"data_hash"`
	ValidatorsHash     string      `json:"validators_hash"`
	NextValidatorsHash string      `json:"next_validators_hash"`
	ConsensusHash      string      `json:"consensus_hash"`
	AppHash            string      `json:"app_hash"`
	LastResultsHash    string      `json:"last_results_hash"`
	EvidenceHash       string      `json:"evidence_hash"`
	ProposerAddress    string      `json:"proposer_address"`
}
type BlockID struct {
	Hash  string `json:"hash"`
	Parts Parts  `json:"parts"`
}
type Signatures struct {
	BlockIDFlag      int       `json:"block_id_flag"`
	ValidatorAddress string    `json:"validator_address"`
	Timestamp        time.Time `json:"timestamp"`
	Signature        string    `json:"signature"`
}
type Commit struct {
	Height     string       `json:"height"`
	Round      string       `json:"round"`
	BlockID    BlockID      `json:"block_id"`
	Signatures []Signatures `json:"signatures"`
}
type SignedHeader struct {
	Header Header `json:"header"`
	Commit Commit `json:"commit"`
}
type Result struct {
	SignedHeader SignedHeader `json:"signed_header"`
	Canonical    bool         `json:"canonical"`
}

type RPCResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  Result `json:"result"`
}

func main() {
	resp, err := grequests.Get("http://d3n-debug.bandprotocol.com:26657/commit?height=10000", &grequests.RequestOptions{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	})

	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	if resp.Ok != true {
		fmt.Println("resp not ok", resp.Error)
		return
	}

	rpcR := RPCResponse{}
	err = resp.JSON(&rpcR)
	if err != nil {
		fmt.Println("resp.JSON(&rpcR) error: ", err)
		return
	}

	spew.Dump(rpcR)
	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
	spew.Dump(err)
}
