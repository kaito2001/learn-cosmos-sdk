package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "vbi-cosmos-basic/testutil/keeper"
	"vbi-cosmos-basic/testutil/nullify"
	"vbi-cosmos-basic/x/chat/keeper"
	"vbi-cosmos-basic/x/chat/types"
)

func createNData(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Data {
	items := make([]types.Data, n)
	for i := range items {
		items[i].Id = keeper.AppendData(ctx, items[i])
	}
	return items
}

func TestDataGet(t *testing.T) {
	keeper, ctx := keepertest.ChatKeeper(t)
	items := createNData(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetData(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestDataRemove(t *testing.T) {
	keeper, ctx := keepertest.ChatKeeper(t)
	items := createNData(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveData(ctx, item.Id)
		_, found := keeper.GetData(ctx, item.Id)
		require.False(t, found)
	}
}

func TestDataGetAll(t *testing.T) {
	keeper, ctx := keepertest.ChatKeeper(t)
	items := createNData(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllData(ctx)),
	)
}

func TestDataCount(t *testing.T) {
	keeper, ctx := keepertest.ChatKeeper(t)
	items := createNData(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetDataCount(ctx))
}
