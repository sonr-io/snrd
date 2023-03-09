package identity

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonrhq/core/x/identity/keeper"
	"github.com/sonrhq/core/x/identity/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the didDocument
	for _, elem := range genState.DidDocumentList {
		k.SetDidDocument(ctx, elem)
	}
	// Set all the DomainRecord
	for _, elem := range genState.ServiceList {
		k.SetService(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.DidDocumentList = k.GetAllDidDocument(ctx)
	genesis.ServiceList = k.GetAllServices(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
