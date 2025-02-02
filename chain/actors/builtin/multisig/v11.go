// Code generated by: `make actors-gen`. DO NOT EDIT.
package multisig

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors"
	"github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"

	"github.com/filecoin-project/lily/chain/actors/adt"

	actorstypes "github.com/filecoin-project/go-state-types/actors"
	"github.com/filecoin-project/go-state-types/manifest"

	"crypto/sha256"

	builtin11 "github.com/filecoin-project/go-state-types/builtin"
	msig11 "github.com/filecoin-project/go-state-types/builtin/v11/multisig"
	adt11 "github.com/filecoin-project/go-state-types/builtin/v11/util/adt"
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

type state11 struct {
	msig11.State
	store adt.Store
}

func (s *state11) LockedBalance(currEpoch abi.ChainEpoch) (abi.TokenAmount, error) {
	return s.State.AmountLocked(currEpoch - s.State.StartEpoch), nil
}

func (s *state11) StartEpoch() (abi.ChainEpoch, error) {
	return s.State.StartEpoch, nil
}

func (s *state11) UnlockDuration() (abi.ChainEpoch, error) {
	return s.State.UnlockDuration, nil
}

func (s *state11) InitialBalance() (abi.TokenAmount, error) {
	return s.State.InitialBalance, nil
}

func (s *state11) Threshold() (uint64, error) {
	return s.State.NumApprovalsThreshold, nil
}

func (s *state11) Signers() ([]address.Address, error) {
	return s.State.Signers, nil
}

func (s *state11) ForEachPendingTxn(cb func(id int64, txn Transaction) error) error {
	arr, err := adt11.AsMap(s.store, s.State.PendingTxns, builtin11.DefaultHamtBitwidth)
	if err != nil {
		return err
	}
	var out msig11.Transaction
	return arr.ForEach(&out, func(key string) error {
		txid, n := binary.Varint([]byte(key))
		if n <= 0 {
			return fmt.Errorf("invalid pending transaction key: %v", key)
		}
		return cb(txid, (Transaction)(out)) //nolint:unconvert
	})
}

func (s *state11) PendingTxnChanged(other State) (bool, error) {
	other11, ok := other.(*state11)
	if !ok {
		// treat an upgrade as a change, always
		return true, nil
	}
	return !s.State.PendingTxns.Equals(other11.PendingTxns), nil
}

func (s *state11) PendingTransactionsMap() (adt.Map, error) {
	return adt11.AsMap(s.store, s.PendingTxns, builtin11.DefaultHamtBitwidth)
}

func (s *state11) PendingTransactionsMapBitWidth() int {

	return builtin11.DefaultHamtBitwidth

}

func (s *state11) PendingTransactionsMapHashFunction() func(input []byte) []byte {

	return func(input []byte) []byte {
		res := sha256.Sum256(input)
		return res[:]
	}

}

func (s *state11) decodeTransaction(val *cbg.Deferred) (Transaction, error) {
	var tx msig11.Transaction
	if err := tx.UnmarshalCBOR(bytes.NewReader(val.Raw)); err != nil {
		return Transaction{}, err
	}
	return Transaction(tx), nil
}

func (s *state11) ActorKey() string {
	return manifest.MultisigKey
}

func (s *state11) ActorVersion() actorstypes.Version {
	return actorstypes.Version11
}

func (s *state11) Code() cid.Cid {
	code, ok := actors.GetActorCodeID(s.ActorVersion(), s.ActorKey())
	if !ok {
		panic(fmt.Errorf("didn't find actor %v code id for actor version %d", s.ActorKey(), s.ActorVersion()))
	}

	return code
}
