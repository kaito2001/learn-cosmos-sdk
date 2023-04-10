package vbicosmosbasic_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "vbi-cosmos-basic/testutil/keeper"
	"vbi-cosmos-basic/testutil/nullify"
	"vbi-cosmos-basic/x/vbicosmosbasic"
	"vbi-cosmos-basic/x/vbicosmosbasic/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VbicosmosbasicKeeper(t)
	vbicosmosbasic.InitGenesis(ctx, *k, genesisState)
	got := vbicosmosbasic.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
