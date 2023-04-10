package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "vbi-cosmos-basic/testutil/keeper"
	"vbi-cosmos-basic/x/vbicosmosbasic/keeper"
	"vbi-cosmos-basic/x/vbicosmosbasic/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.VbicosmosbasicKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
