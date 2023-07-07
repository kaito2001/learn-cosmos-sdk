package todo

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"vbi-cosmos-basic/testutil/sample"
	todosimulation "vbi-cosmos-basic/x/todo/simulation"
	"vbi-cosmos-basic/x/todo/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = todosimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreatePaper = "op_weight_msg_paper"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreatePaper int = 100

	opWeightMsgUpdatePaper = "op_weight_msg_paper"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdatePaper int = 100

	opWeightMsgDeletePaper = "op_weight_msg_paper"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeletePaper int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	todoGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PaperList: []types.Paper{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		PaperCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&todoGenesis)
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

	var weightMsgCreatePaper int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreatePaper, &weightMsgCreatePaper, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePaper = defaultWeightMsgCreatePaper
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePaper,
		todosimulation.SimulateMsgCreatePaper(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdatePaper int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdatePaper, &weightMsgUpdatePaper, nil,
		func(_ *rand.Rand) {
			weightMsgUpdatePaper = defaultWeightMsgUpdatePaper
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdatePaper,
		todosimulation.SimulateMsgUpdatePaper(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeletePaper int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeletePaper, &weightMsgDeletePaper, nil,
		func(_ *rand.Rand) {
			weightMsgDeletePaper = defaultWeightMsgDeletePaper
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeletePaper,
		todosimulation.SimulateMsgDeletePaper(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
