package main

import (
	"bytes"
	"fmt"

	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

type ProofLeafNode struct {
	Key       []byte `json:"key"`
	ValueHash []byte `json:"value"`
	Version   int64  `json:"version"`
}

type ProofInnerNode struct {
	Height  int8   `json:"height"`
	Size    int64  `json:"size"`
	Version int64  `json:"version"`
	Left    []byte `json:"left"`
	Right   []byte `json:"right"`
}

type RangeProof struct {
	LeftPath   []ProofInnerNode   `json:"left_path"`
	InnerNodes [][]ProofInnerNode `json:"inner_nodes"`
	Leaves     []ProofLeafNode    `json:"leaves"`
}

type ValueOp struct {
	key   []byte
	Proof *RangeProof `json:"proof"`
}

func encodeHeightSizeVersion(height int8, size, version int64) []byte {
	buf := new(bytes.Buffer)
	amino.EncodeInt8(buf, height)
	amino.EncodeVarint(buf, size)
	amino.EncodeVarint(buf, version)
	return buf.Bytes()
}

func IAVLHash() {
	var opr1 ValueOp
	_ = (amino.NewCodec()).UnmarshalBinaryLengthPrefixed(
		Base64ToBytes("twIKtAIKKQgOEDwYkE4qIAqVHc4yEL3Lr0lZ0cQPJ+nMt/TaT+H5Cd/pZOcdWmWZCikIDBAdGJBOKiAOo7imPYU0LXqSghAX34+evA8tHUnu6SI3mWqUAS5rcAopCAgQDBiQTiIgLcT4yW/qM8aqxuhaw+Zv55iofahwEGUGcAmeyzas1AIKKQgGEAgY8ikqIIQMVU13Mxhj2pnxGywAfwoRaRX7qIeWcSAqI/ot7OwsCikIBBAEGIoiKiCe4B7tdcsF0tG2W+f+nMvQdzjOsQkIvqNxbNHXfpL69gopCAIQAhiNDSog/tzkcYxC546anyOplbnELEARIvc8NNk990HEd3X4dZ0aMAoJAQAAAAAAAAABEiBF01aMSserjw6ZditfkKaR1q0SHsDI0mptE2Nzxud0RBioAQ=="),
		&opr1,
	)
	iavlop := opr1.Proof

	leaveBytes := append(
		encodeHeightSizeVersion(0, 1, iavlop.Leaves[0].Version),
		EncodeKV(string(iavlop.Leaves[0].Key), iavlop.Leaves[0].ValueHash)...,
	)
	iavlHash := tmhash.Sum(leaveBytes)

	for i := len(iavlop.LeftPath) - 1; i >= 0; i-- {
		var bs []byte
		if len(iavlop.LeftPath[i].Left) == 0 {
			bs = append([]byte{32}, append(iavlHash, append([]byte{32}, iavlop.LeftPath[i].Right...)...)...)
		} else {
			bs = append([]byte{32}, append(iavlop.LeftPath[i].Left, append([]byte{32}, iavlHash...)...)...)
		}

		iavlHash = tmhash.Sum(append(
			encodeHeightSizeVersion(
				iavlop.LeftPath[i].Height,
				iavlop.LeftPath[i].Size,
				iavlop.LeftPath[i].Version,
			),
			bs...,
		))
	}

	fmt.Println(fmt.Sprintf("%x \n", iavlHash))
}
