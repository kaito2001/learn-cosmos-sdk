package todo

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"vbi-cosmos-basic/x/todo/keeper"
	"vbi-cosmos-basic/x/todo/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the paper
	for _, elem := range genState.PaperList {
		k.SetPaper(ctx, elem)
	}

	// Set paper count
	k.SetPaperCount(ctx, genState.PaperCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PaperList = k.GetAllPaper(ctx)
	genesis.PaperCount = k.GetPaperCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
