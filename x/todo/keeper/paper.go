package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"vbi-cosmos-basic/x/todo/types"
)

// GetPaperCount get the total number of paper
func (k Keeper) GetPaperCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.PaperCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetPaperCount set the total number of paper
func (k Keeper) SetPaperCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.PaperCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendPaper appends a paper in the store with a new id and update the count
func (k Keeper) AppendPaper(
	ctx sdk.Context,
	paper types.Paper,
) uint64 {
	// Create the paper
	count := k.GetPaperCount(ctx)

	// Set the ID of the appended value
	paper.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PaperKey))
	appendedValue := k.cdc.MustMarshal(&paper)
	store.Set(GetPaperIDBytes(paper.Id), appendedValue)

	// Update paper count
	k.SetPaperCount(ctx, count+1)

	return count
}

// SetPaper set a specific paper in the store
func (k Keeper) SetPaper(ctx sdk.Context, paper types.Paper) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PaperKey))
	b := k.cdc.MustMarshal(&paper)
	store.Set(GetPaperIDBytes(paper.Id), b)
}

// GetPaper returns a paper from its id
func (k Keeper) GetPaper(ctx sdk.Context, id uint64) (val types.Paper, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PaperKey))
	b := store.Get(GetPaperIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePaper removes a paper from the store
func (k Keeper) RemovePaper(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PaperKey))
	store.Delete(GetPaperIDBytes(id))
}

// GetAllPaper returns all paper
func (k Keeper) GetAllPaper(ctx sdk.Context) (list []types.Paper) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PaperKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Paper
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetPaperIDBytes returns the byte representation of the ID
func GetPaperIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetPaperIDFromBytes returns ID in uint64 format from a byte array
func GetPaperIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
