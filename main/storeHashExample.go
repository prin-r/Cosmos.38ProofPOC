package main

import (
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

func StoreHash() {
	var opr2 MultiStoreProofOp
	_ = (amino.NewCodec()).UnmarshalBinaryBare(
		Base64ToBytes("CoMFCi4KA2FjYxInCiUIkE4SIFLcE1ahhs4cq3V5OdCkp37QgcR2DLAHJEghUt8kOtfACjEKBm9yYWNsZRInCiUIkE4SIIwuNVuWpFTGVuBnZyWiXX2CcBQkjH7E/J5WRio+K+VOCjcKDGRpc3RyaWJ1dGlvbhInCiUIkE4SIBbWxbUhOXsbPBHbCmCv1/xo5hiGWtPIg9sLWHvt1ObrChAKB3VwZ3JhZGUSBQoDCJBOCi8KBG1haW4SJwolCJBOEiAjmXfb2YDKqg/X5EeV386YpNoE9Zr01JqJ9noPdoLJuwouCgNnb3YSJwolCJBOEiBgzRN8GWLsrGFjidaANIM/KSFQnC0oWqX1FTmXzpaKdAoyCgdzdGFraW5nEicKJQiQThIg6saY1G4hHFlWdh6pVvA+b2Z595TYExXNNsIlS2GDPfwKMQoGc3VwcGx5EicKJQiQThIgMCYYZhE3gW2M7aEMSGPNy3J+Nt83WGOZ3zf5Btxyu1YKMQoGcGFyYW1zEicKJQiQThIgxt7VueGVn9BcqHC3ZsEOS4QJLJeNAlH2LXrd6aNh6ssKEQoIZXZpZGVuY2USBQoDCJBOCi8KBG1pbnQSJwolCJBOEiCjAPoTDCcCcs1ElKEZ9a0vncO6GhGFErVmiyFNkl0VDwouCgNpYmMSJwolCJBOEiDYV07lM+os9Q9vR4MrLPvm/wXy0piJsZSTW/3b3Wj/TAozCghzbGFzaGluZxInCiUIkE4SIGMT9FctkDNMV3jsmW5EjiZEFSorUx394uVp4oYa8cm3Ci8KBGJhbmsSJwolCJBOEiD6OUYBrOLarjGKi+0ARmHBpOrcyNKlABOoi1XH+Q4DkA=="),
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

	fmt.Println(r)

	mySortedKeys := []string{"acc", "bank", "distribution", "evidence", "gov", "ibc", "main", "mint", "oracle", "params", "slashing", "staking", "supply", "upgrade"}
	mySortedEncoded := [][]byte{}
	s := "\n"
	for _, sk := range mySortedKeys {
		s += fmt.Sprintf("%s. 0x%x\n", sk, m[sk])
		// s += fmt.Sprintf("encodeKV 0x%x\n", EncodeKV(sk, m[sk]))
		mySortedEncoded = append(mySortedEncoded, EncodeKV(sk, m[sk]))
	}

	// fmt.Println(s)

	myHash := HashHeader(mySortedEncoded)

	fmt.Println(fmt.Sprintf("%x\n", myHash))
}
