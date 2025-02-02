// Code generated by: `make actors-gen`. DO NOT EDIT.
package reward

import (
	"fmt"
	"github.com/filecoin-project/go-state-types/abi"
	reward0 "github.com/filecoin-project/specs-actors/actors/builtin/reward"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-state-types/cbor"

	builtin0 "github.com/filecoin-project/specs-actors/actors/builtin"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"

	builtin3 "github.com/filecoin-project/specs-actors/v3/actors/builtin"

	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"

	builtin5 "github.com/filecoin-project/specs-actors/v5/actors/builtin"

	builtin6 "github.com/filecoin-project/specs-actors/v6/actors/builtin"

	builtin7 "github.com/filecoin-project/specs-actors/v7/actors/builtin"

	builtin11 "github.com/filecoin-project/go-state-types/builtin"

	"github.com/filecoin-project/lotus/chain/types"

	"github.com/filecoin-project/lily/chain/actors/adt"
	"github.com/filecoin-project/lily/chain/actors/builtin"
	lotusactors "github.com/filecoin-project/lotus/chain/actors"

	actorstypes "github.com/filecoin-project/go-state-types/actors"
	"github.com/filecoin-project/go-state-types/manifest"
)

var (
	Address = builtin11.RewardActorAddr
	Methods = builtin11.MethodsReward
)

func Load(store adt.Store, act *types.Actor) (State, error) {
	if name, av, ok := lotusactors.GetActorMetaByCode(act.Code); ok {
		if name != manifest.RewardKey {
			return nil, fmt.Errorf("actor code is not reward: %s", name)
		}

		switch actorstypes.Version(av) {

		case actorstypes.Version8:
			return load8(store, act.Head)

		case actorstypes.Version9:
			return load9(store, act.Head)

		case actorstypes.Version10:
			return load10(store, act.Head)

		case actorstypes.Version11:
			return load11(store, act.Head)

		}
	}

	switch act.Code {

	case builtin0.RewardActorCodeID:
		return load0(store, act.Head)

	case builtin2.RewardActorCodeID:
		return load2(store, act.Head)

	case builtin3.RewardActorCodeID:
		return load3(store, act.Head)

	case builtin4.RewardActorCodeID:
		return load4(store, act.Head)

	case builtin5.RewardActorCodeID:
		return load5(store, act.Head)

	case builtin6.RewardActorCodeID:
		return load6(store, act.Head)

	case builtin7.RewardActorCodeID:
		return load7(store, act.Head)

	}

	return nil, fmt.Errorf("unknown actor code %s", act.Code)
}

type State interface {
	cbor.Marshaler

	Code() cid.Cid
	ActorKey() string
	ActorVersion() actorstypes.Version

	ThisEpochBaselinePower() (abi.StoragePower, error)
	ThisEpochReward() (abi.StoragePower, error)
	ThisEpochRewardSmoothed() (builtin.FilterEstimate, error)

	EffectiveBaselinePower() (abi.StoragePower, error)
	EffectiveNetworkTime() (abi.ChainEpoch, error)

	TotalStoragePowerReward() (abi.TokenAmount, error)

	CumsumBaseline() (abi.StoragePower, error)
	CumsumRealized() (abi.StoragePower, error)

	InitialPledgeForPower(abi.StoragePower, abi.TokenAmount, *builtin.FilterEstimate, abi.TokenAmount) (abi.TokenAmount, error)
	PreCommitDepositForPower(builtin.FilterEstimate, abi.StoragePower) (abi.TokenAmount, error)
}

type AwardBlockRewardParams = reward0.AwardBlockRewardParams

func AllCodes() []cid.Cid {
	return []cid.Cid{
		(&state0{}).Code(),
		(&state2{}).Code(),
		(&state3{}).Code(),
		(&state4{}).Code(),
		(&state5{}).Code(),
		(&state6{}).Code(),
		(&state7{}).Code(),
		(&state8{}).Code(),
		(&state9{}).Code(),
		(&state10{}).Code(),
		(&state11{}).Code(),
	}
}

func VersionCodes() map[actorstypes.Version]cid.Cid {
	return map[actorstypes.Version]cid.Cid{
		actorstypes.Version0:  (&state0{}).Code(),
		actorstypes.Version2:  (&state2{}).Code(),
		actorstypes.Version3:  (&state3{}).Code(),
		actorstypes.Version4:  (&state4{}).Code(),
		actorstypes.Version5:  (&state5{}).Code(),
		actorstypes.Version6:  (&state6{}).Code(),
		actorstypes.Version7:  (&state7{}).Code(),
		actorstypes.Version8:  (&state8{}).Code(),
		actorstypes.Version9:  (&state9{}).Code(),
		actorstypes.Version10: (&state10{}).Code(),
		actorstypes.Version11: (&state11{}).Code(),
	}
}
