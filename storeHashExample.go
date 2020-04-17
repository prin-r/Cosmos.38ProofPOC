package main

import (
	"encoding/base64"
	"fmt"

	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

type CommitID struct {
	Version int64
	Hash    []byte
}

type storeCore struct {
	CommitID CommitID
}

type storeInfo struct {
	Name string
	Core storeCore
}

type MultiStoreProof struct {
	StoreInfos []storeInfo
}

type MultiStoreProofOp struct {
	key   []byte
	Proof *MultiStoreProof `json:"proof"`
}

func base64ToBytes(s string) []byte {
	b64, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err.Error())
	}
	return b64
}

func encodeKV(key string, value []byte) []byte {
	lenKey := uint8(len([]byte(key)))
	lenValue := uint8(len(value))
	return append(append([]byte{lenKey}, []byte(key)...), append([]byte{lenValue}, value...)...)
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

func storeHash() {
	var opr2 MultiStoreProofOp
	_ = (amino.NewCodec()).UnmarshalBinaryBare(
		base64ToBytes("CoMFCi4KA2FjYxInCiUIkE4SIFLcE1ahhs4cq3V5OdCkp37QgcR2DLAHJEghUt8kOtfACjEKBm9yYWNsZRInCiUIkE4SIIwuNVuWpFTGVuBnZyWiXX2CcBQkjH7E/J5WRio+K+VOCjcKDGRpc3RyaWJ1dGlvbhInCiUIkE4SIBbWxbUhOXsbPBHbCmCv1/xo5hiGWtPIg9sLWHvt1ObrChAKB3VwZ3JhZGUSBQoDCJBOCi8KBG1haW4SJwolCJBOEiAjmXfb2YDKqg/X5EeV386YpNoE9Zr01JqJ9noPdoLJuwouCgNnb3YSJwolCJBOEiBgzRN8GWLsrGFjidaANIM/KSFQnC0oWqX1FTmXzpaKdAoyCgdzdGFraW5nEicKJQiQThIg6saY1G4hHFlWdh6pVvA+b2Z595TYExXNNsIlS2GDPfwKMQoGc3VwcGx5EicKJQiQThIgMCYYZhE3gW2M7aEMSGPNy3J+Nt83WGOZ3zf5Btxyu1YKMQoGcGFyYW1zEicKJQiQThIgxt7VueGVn9BcqHC3ZsEOS4QJLJeNAlH2LXrd6aNh6ssKEQoIZXZpZGVuY2USBQoDCJBOCi8KBG1pbnQSJwolCJBOEiCjAPoTDCcCcs1ElKEZ9a0vncO6GhGFErVmiyFNkl0VDwouCgNpYmMSJwolCJBOEiDYV07lM+os9Q9vR4MrLPvm/wXy0piJsZSTW/3b3Wj/TAozCghzbGFzaGluZxInCiUIkE4SIGMT9FctkDNMV3jsmW5EjiZEFSorUx394uVp4oYa8cm3Ci8KBGJhbmsSJwolCJBOEiD6OUYBrOLarjGKi+0ARmHBpOrcyNKlABOoi1XH+Q4DkA=="),
		&opr2,
	)

	msp := opr2.Proof
	m := make(map[string][]byte, len(msp.StoreInfos))

	r := ""
	for _, si := range msp.StoreInfos {
		bz := si.Core.CommitID.Hash
		m[si.Name] = tmhash.Sum(tmhash.Sum(bz))
		r += fmt.Sprintf("%s -> %x\n", si.Name, si.Core.CommitID.Hash)
	}

	// fmt.Println(r)

	mySortedKeys := []string{"acc", "bank", "distribution", "evidence", "gov", "ibc", "main", "mint", "oracle", "params", "slashing", "staking", "supply", "upgrade"}
	mySortedEncoded := [][]byte{}
	s := "\n"
	for i, sk := range mySortedKeys {
		s += fmt.Sprintf("%d. 0x%x\n", i+1, m[sk])
		s += fmt.Sprintf("encodeKV 0x%x\n", encodeKV(sk, m[sk]))
		mySortedEncoded = append(mySortedEncoded, encodeKV(sk, m[sk]))
	}

	// fmt.Println(s)

	myHash := hashHeader(mySortedEncoded)

	fmt.Println(fmt.Sprintf("%x", myHash))
}

func main() {
	storeHash()
}
