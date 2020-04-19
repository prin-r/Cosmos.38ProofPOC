package main

import (
	"fmt"

	tmtypes "github.com/tendermint/tendermint/types"
	"github.com/tendermint/tendermint/version"
)

func BlockHash() {
	tmHeader := tmtypes.Header{
		Version: version.Consensus{Block: 10, App: 0},
		ChainID: "bandchain",
		Height:  10000,
		Time:    ParseTime("2020-04-14T05:16:11.764948446Z"),
		LastBlockID: tmtypes.BlockID{
			Hash: HexToBytes("C6399145BC6AFB251CE0C225EBF9D2243E08AB69068EF676CA4226E9DF2E7D50"),
			PartsHeader: tmtypes.PartSetHeader{
				Total: 1,
				Hash:  HexToBytes("B8A525FAE1925947AA3378FFCA796396FD5B3E3FA2E7701E252B76DF84AD4ECF"),
			},
		},
		LastCommitHash:     HexToBytes("EC2A9FBF731509D8B03B5B9E7DD5098D2774D3179E2EEEC3C2349B3176302732"),
		DataHash:           nil,
		ValidatorsHash:     HexToBytes("3AEB137B43144B229F0CA7AC43033E03FCEE25877A3661E88848E436C3D6DD65"),
		NextValidatorsHash: HexToBytes("3AEB137B43144B229F0CA7AC43033E03FCEE25877A3661E88848E436C3D6DD65"),
		ConsensusHash:      HexToBytes("AD82B220C509602720D74FD75BCE7CFE9B148039958F236D8894E00EB1599E04"),
		AppHash:            HexToBytes("9592EB9B13206F557F123FB98E9B4BC9B4963F9F8A2FA46A67BB421944FB2B08"),
		LastResultsHash:    nil,
		EvidenceHash:       nil,
		ProposerAddress:    HexToBytes("F0C23921727D869745C4F9703CF33996B1D2B715"),
	}

	elems := [][]byte{
		EncodeIfNotNil(tmHeader.Version),
		EncodeIfNotNil(tmHeader.ChainID),
		EncodeIfNotNil(tmHeader.Height),
		EncodeIfNotNil(tmHeader.Time),
		EncodeIfNotNil(tmHeader.LastBlockID),
		EncodeIfNotNil(tmHeader.LastCommitHash),
		EncodeIfNotNil(tmHeader.DataHash),
		EncodeIfNotNil(tmHeader.ValidatorsHash),
		EncodeIfNotNil(tmHeader.NextValidatorsHash),
		EncodeIfNotNil(tmHeader.ConsensusHash),
		EncodeIfNotNil(tmHeader.AppHash),
		EncodeIfNotNil(tmHeader.LastResultsHash),
		EncodeIfNotNil(tmHeader.EvidenceHash),
		EncodeIfNotNil(tmHeader.ProposerAddress),
	}
	s := ""
	for _, e := range elems {
		s += fmt.Sprintf("%x\n", e)
	}

	fmt.Println(s)
	fmt.Println(fmt.Sprintf("%x ¢¢¢¢£££", HashHeader(elems)))
}
