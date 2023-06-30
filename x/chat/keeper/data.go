package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"vbi-cosmos-basic/x/chat/types"
)

// GetDataCount get the total number of data
func (k Keeper) GetDataCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.DataCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetDataCount set the total number of data
func (k Keeper) SetDataCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.DataCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendData appends a data in the store with a new id and update the count
func (k Keeper) AppendData(
	ctx sdk.Context,
	data types.Data,
) uint64 {
	// Create the data
	count := k.GetDataCount(ctx)

	// Set the ID of the appended value
	data.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DataKey))
	appendedValue := k.cdc.MustMarshal(&data)
	store.Set(GetDataIDBytes(data.Id), appendedValue)

	// Update data count
	k.SetDataCount(ctx, count+1)

	return count
}

// SetData set a specific data in the store
func (k Keeper) SetData(ctx sdk.Context, data types.Data) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DataKey))
	b := k.cdc.MustMarshal(&data)
	store.Set(GetDataIDBytes(data.Id), b)
}

// GetData returns a data from its id
func (k Keeper) GetData(ctx sdk.Context, id uint64) (val types.Data, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DataKey))
	b := store.Get(GetDataIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveData removes a data from the store
func (k Keeper) RemoveData(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DataKey))
	store.Delete(GetDataIDBytes(id))
}

// GetAllData returns all data
func (k Keeper) GetAllData(ctx sdk.Context) (list []types.Data) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DataKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Data
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetDataIDBytes returns the byte representation of the ID
func GetDataIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetDataIDFromBytes returns ID in uint64 format from a byte array
func GetDataIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
