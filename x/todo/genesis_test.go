package todo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "vbi-cosmos-basic/testutil/keeper"
	"vbi-cosmos-basic/testutil/nullify"
	"vbi-cosmos-basic/x/todo"
	"vbi-cosmos-basic/x/todo/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		PaperList: []types.Paper{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		PaperCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TodoKeeper(t)
	todo.InitGenesis(ctx, *k, genesisState)
	got := todo.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.PaperList, got.PaperList)
	require.Equal(t, genesisState.PaperCount, got.PaperCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
