package chat

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"vbi-cosmos-basic/testutil/sample"
	chatsimulation "vbi-cosmos-basic/x/chat/simulation"
	"vbi-cosmos-basic/x/chat/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = chatsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateData = "op_weight_msg_data"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateData int = 100

	opWeightMsgUpdateData = "op_weight_msg_data"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateData int = 100

	opWeightMsgDeleteData = "op_weight_msg_data"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteData int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	chatGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		DataList: []types.Data{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		DataCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&chatGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateData int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateData, &weightMsgCreateData, nil,
		func(_ *rand.Rand) {
			weightMsgCreateData = defaultWeightMsgCreateData
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateData,
		chatsimulation.SimulateMsgCreateData(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateData int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateData, &weightMsgUpdateData, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateData = defaultWeightMsgUpdateData
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateData,
		chatsimulation.SimulateMsgUpdateData(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteData int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteData, &weightMsgDeleteData, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteData = defaultWeightMsgDeleteData
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteData,
		chatsimulation.SimulateMsgDeleteData(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
