package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "vbi-cosmos-basic/testutil/keeper"
	"vbi-cosmos-basic/testutil/nullify"
	"vbi-cosmos-basic/x/todo/keeper"
	"vbi-cosmos-basic/x/todo/types"
)

func createNPaper(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Paper {
	items := make([]types.Paper, n)
	for i := range items {
		items[i].Id = keeper.AppendPaper(ctx, items[i])
	}
	return items
}

func TestPaperGet(t *testing.T) {
	keeper, ctx := keepertest.TodoKeeper(t)
	items := createNPaper(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetPaper(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestPaperRemove(t *testing.T) {
	keeper, ctx := keepertest.TodoKeeper(t)
	items := createNPaper(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePaper(ctx, item.Id)
		_, found := keeper.GetPaper(ctx, item.Id)
		require.False(t, found)
	}
}

func TestPaperGetAll(t *testing.T) {
	keeper, ctx := keepertest.TodoKeeper(t)
	items := createNPaper(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPaper(ctx)),
	)
}

func TestPaperCount(t *testing.T) {
	keeper, ctx := keepertest.TodoKeeper(t)
	items := createNPaper(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetPaperCount(ctx))
}
