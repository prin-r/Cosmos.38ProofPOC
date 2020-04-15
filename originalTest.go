package main

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"time"

	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/crypto/tmhash"
	tmtypes "github.com/tendermint/tendermint/types"
	"github.com/tendermint/tendermint/version"
)

func parseTime(s string) time.Time {
	t, _ := time.Parse(time.RFC3339, s)
	return t
}

func hexToBytes(s string) []byte {
	x, _ := hex.DecodeString(s)
	return x
}

func isTypedNil(o interface{}) bool {
	rv := reflect.ValueOf(o)
	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Slice:
		return rv.IsNil()
	default:
		return false
	}
}

func myEncode(item interface{}) []byte {
	if item == nil || isTypedNil(item) {
		return nil
	}
	return (amino.NewCodec()).MustMarshalBinaryBare(item)
}

func leafHash(item []byte) []byte {
	// leaf prefix is 0
	return tmhash.Sum(append([]byte{0}, item...))
}

func branchHash(left, right []byte) []byte {
	// branch prefix is 1
	return tmhash.Sum(append([]byte{1}, append(left, right...)...))
}

func hashHeader(items [][]byte) []byte {
	if len(items) != 14 {
		panic("hashHeader error, len of items must be 14")
	}
	return branchHash(
		branchHash(
			branchHash(
				branchHash(leafHash(items[0]), leafHash(items[1])),
				branchHash(leafHash(items[2]), leafHash(items[3])),
			),
			branchHash(
				branchHash(leafHash(items[4]), leafHash(items[5])),
				branchHash(leafHash(items[6]), leafHash(items[7])),
			),
		),
		branchHash(
			branchHash(
				branchHash(leafHash(items[8]), leafHash(items[9])),
				branchHash(leafHash(items[10]), leafHash(items[11])),
			),
			branchHash(leafHash(items[12]), leafHash(items[13])),
		),
	)
}

func MyTest1() {
	tmHeader := tmtypes.Header{
		Version: version.Consensus{Block: 10, App: 0},
		ChainID: "bandchain",
		Height:  10000,
		Time:    parseTime("2020-04-14T05:16:11.764948446Z"),
		LastBlockID: tmtypes.BlockID{
			Hash: hexToBytes("C6399145BC6AFB251CE0C225EBF9D2243E08AB69068EF676CA4226E9DF2E7D50"),
			PartsHeader: tmtypes.PartSetHeader{
				Total: 1,
				Hash:  hexToBytes("B8A525FAE1925947AA3378FFCA796396FD5B3E3FA2E7701E252B76DF84AD4ECF"),
			},
		},
		LastCommitHash:     hexToBytes("EC2A9FBF731509D8B03B5B9E7DD5098D2774D3179E2EEEC3C2349B3176302732"),
		DataHash:           nil,
		ValidatorsHash:     hexToBytes("3AEB137B43144B229F0CA7AC43033E03FCEE25877A3661E88848E436C3D6DD65"),
		NextValidatorsHash: hexToBytes("3AEB137B43144B229F0CA7AC43033E03FCEE25877A3661E88848E436C3D6DD65"),
		ConsensusHash:      hexToBytes("AD82B220C509602720D74FD75BCE7CFE9B148039958F236D8894E00EB1599E04"),
		AppHash:            hexToBytes("9592EB9B13206F557F123FB98E9B4BC9B4963F9F8A2FA46A67BB421944FB2B08"),
		LastResultsHash:    nil,
		EvidenceHash:       nil,
		ProposerAddress:    hexToBytes("F0C23921727D869745C4F9703CF33996B1D2B715"),
	}

	elems := [][]byte{
		myEncode(tmHeader.Version),
		myEncode(tmHeader.ChainID),
		myEncode(tmHeader.Height),
		myEncode(tmHeader.Time),
		myEncode(tmHeader.LastBlockID),
		myEncode(tmHeader.LastCommitHash),
		myEncode(tmHeader.DataHash),
		myEncode(tmHeader.ValidatorsHash),
		myEncode(tmHeader.NextValidatorsHash),
		myEncode(tmHeader.ConsensusHash),
		myEncode(tmHeader.AppHash),
		myEncode(tmHeader.LastResultsHash),
		myEncode(tmHeader.EvidenceHash),
		myEncode(tmHeader.ProposerAddress),
	}
	s := ""
	for _, e := range elems {
		s += fmt.Sprintf("%x\n", e)
	}

	fmt.Println(s)
	fmt.Println(fmt.Sprintf("%x ¢¢¢¢£££", hashHeader(elems)))
}

func main() {
	MyTest1()
}
