// Code generated by: `make actors-gen`. DO NOT EDIT.

package market

import (
	"bytes"
	"fmt"

	"github.com/filecoin-project/go-state-types/abi"
	actorstypes "github.com/filecoin-project/go-state-types/actors"
	"github.com/filecoin-project/go-state-types/manifest"
	lotusactors "github.com/filecoin-project/lotus/chain/actors"
	"github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/chain/actors/adt"

	verifregtypes "github.com/filecoin-project/go-state-types/builtin/v9/verifreg"

	market11 "github.com/filecoin-project/go-state-types/builtin/v11/market"
	adt11 "github.com/filecoin-project/go-state-types/builtin/v11/util/adt"
	markettypes "github.com/filecoin-project/go-state-types/builtin/v9/market"
)

var _ State = (*state11)(nil)

func load11(store adt.Store, root cid.Cid) (State, error) {
	out := state11{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func make11(store adt.Store) (State, error) {
	out := state11{store: store}

	s, err := market11.ConstructState(store)
	if err != nil {
		return nil, err
	}

	out.State = *s

	return &out, nil
}

type state11 struct {
	market11.State
	store adt.Store
}

func (s *state11) StatesChanged(otherState State) (bool, error) {
	otherState11, ok := otherState.(*state11)
	if !ok {
		// there's no way to compare different versions of the state, so let's
		// just say that means the state of balances has changed
		return true, nil
	}
	return !s.State.States.Equals(otherState11.State.States), nil
}

func (s *state11) States() (DealStates, error) {
	stateArray, err := adt11.AsArray(s.store, s.State.States, market11.StatesAmtBitwidth)
	if err != nil {
		return nil, err
	}
	return &dealStates11{stateArray}, nil
}

func (s *state11) ProposalsChanged(otherState State) (bool, error) {
	otherState11, ok := otherState.(*state11)
	if !ok {
		// there's no way to compare different versions of the state, so let's
		// just say that means the state of balances has changed
		return true, nil
	}
	return !s.State.Proposals.Equals(otherState11.State.Proposals), nil
}

func (s *state11) Proposals() (DealProposals, error) {
	proposalArray, err := adt11.AsArray(s.store, s.State.Proposals, market11.ProposalsAmtBitwidth)
	if err != nil {
		return nil, err
	}
	return &dealProposals11{proposalArray}, nil
}

type dealStates11 struct {
	adt.Array
}

func (s *dealStates11) Get(dealID abi.DealID) (*DealState, bool, error) {
	var deal11 market11.DealState
	found, err := s.Array.Get(uint64(dealID), &deal11)
	if err != nil {
		return nil, false, err
	}
	if !found {
		return nil, false, nil
	}
	deal := fromV11DealState(deal11)
	return &deal, true, nil
}

func (s *dealStates11) ForEach(cb func(dealID abi.DealID, ds DealState) error) error {
	var ds11 market11.DealState
	return s.Array.ForEach(&ds11, func(idx int64) error {
		return cb(abi.DealID(idx), fromV11DealState(ds11))
	})
}

func (s *dealStates11) decode(val *cbg.Deferred) (*DealState, error) {
	var ds11 market11.DealState
	if err := ds11.UnmarshalCBOR(bytes.NewReader(val.Raw)); err != nil {
		return nil, err
	}
	ds := fromV11DealState(ds11)
	return &ds, nil
}

func (s *dealStates11) array() adt.Array {
	return s.Array
}

func fromV11DealState(v11 market11.DealState) DealState {
	ret := DealState{
		SectorStartEpoch: v11.SectorStartEpoch,
		LastUpdatedEpoch: v11.LastUpdatedEpoch,
		SlashEpoch:       v11.SlashEpoch,
		VerifiedClaim:    0,
	}

	ret.VerifiedClaim = verifregtypes.AllocationId(v11.VerifiedClaim)

	return ret
}

type dealProposals11 struct {
	adt.Array
}

func (s *dealProposals11) Get(dealID abi.DealID) (*DealProposal, bool, error) {
	var proposal11 market11.DealProposal
	found, err := s.Array.Get(uint64(dealID), &proposal11)
	if err != nil {
		return nil, false, err
	}
	if !found {
		return nil, false, nil
	}

	proposal, err := fromV11DealProposal(proposal11)
	if err != nil {
		return nil, true, xerrors.Errorf("decoding proposal: %w", err)
	}

	return &proposal, true, nil
}

func (s *dealProposals11) ForEach(cb func(dealID abi.DealID, dp DealProposal) error) error {
	var dp11 market11.DealProposal
	return s.Array.ForEach(&dp11, func(idx int64) error {
		dp, err := fromV11DealProposal(dp11)
		if err != nil {
			return xerrors.Errorf("decoding proposal: %w", err)
		}

		return cb(abi.DealID(idx), dp)
	})
}

func (s *dealProposals11) decode(val *cbg.Deferred) (*DealProposal, error) {
	var dp11 market11.DealProposal
	if err := dp11.UnmarshalCBOR(bytes.NewReader(val.Raw)); err != nil {
		return nil, err
	}

	dp, err := fromV11DealProposal(dp11)
	if err != nil {
		return nil, err
	}

	return &dp, nil
}

func (s *dealProposals11) array() adt.Array {
	return s.Array
}

func fromV11DealProposal(v11 market11.DealProposal) (DealProposal, error) {

	label, err := fromV11Label(v11.Label)

	if err != nil {
		return DealProposal{}, xerrors.Errorf("error setting deal label: %w", err)
	}

	return DealProposal{
		PieceCID:     v11.PieceCID,
		PieceSize:    v11.PieceSize,
		VerifiedDeal: v11.VerifiedDeal,
		Client:       v11.Client,
		Provider:     v11.Provider,

		Label: label,

		StartEpoch:           v11.StartEpoch,
		EndEpoch:             v11.EndEpoch,
		StoragePricePerEpoch: v11.StoragePricePerEpoch,

		ProviderCollateral: v11.ProviderCollateral,
		ClientCollateral:   v11.ClientCollateral,
	}, nil
}

func (s *state11) DealProposalsAmtBitwidth() int {
	return market11.ProposalsAmtBitwidth
}

func (s *state11) DealStatesAmtBitwidth() int {
	return market11.StatesAmtBitwidth
}

func (s *state11) ActorKey() string {
	return manifest.MarketKey
}

func (s *state11) ActorVersion() actorstypes.Version {
	return actorstypes.Version11
}

func (s *state11) Code() cid.Cid {
	code, ok := lotusactors.GetActorCodeID(s.ActorVersion(), s.ActorKey())
	if !ok {
		panic(fmt.Errorf("didn't find actor %v code id for actor version %d", s.ActorKey(), s.ActorVersion()))
	}

	return code
}

func fromV11Label(v11 market11.DealLabel) (DealLabel, error) {
	if v11.IsString() {
		str, err := v11.ToString()
		if err != nil {
			return markettypes.EmptyDealLabel, xerrors.Errorf("failed to convert string label to string: %w", err)
		}
		return markettypes.NewLabelFromString(str)
	}

	bs, err := v11.ToBytes()
	if err != nil {
		return markettypes.EmptyDealLabel, xerrors.Errorf("failed to convert bytes label to bytes: %w", err)
	}
	return markettypes.NewLabelFromBytes(bs)
}
