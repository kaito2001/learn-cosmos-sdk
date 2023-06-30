package chat_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "vbi-cosmos-basic/testutil/keeper"
	"vbi-cosmos-basic/testutil/nullify"
	"vbi-cosmos-basic/x/chat"
	"vbi-cosmos-basic/x/chat/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		DataList: []types.Data{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		DataCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ChatKeeper(t)
	chat.InitGenesis(ctx, *k, genesisState)
	got := chat.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.DataList, got.DataList)
	require.Equal(t, genesisState.DataCount, got.DataCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
