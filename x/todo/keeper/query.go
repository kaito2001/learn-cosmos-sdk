package keeper

import (
	"vbi-cosmos-basic/x/todo/types"
)

var _ types.QueryServer = Keeper{}
