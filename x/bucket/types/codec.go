package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgDefineBucket{}, "bucket/CreateWhereIs", nil)
	cdc.RegisterConcrete(&MsgGenerateBucket{}, "bucket/GenerateBucket", nil)
	cdc.RegisterConcrete(&MsgDeactivateBucket{}, "bucket/DeactivateBucket", nil)
	cdc.RegisterConcrete(&MsgBurnBucket{}, "bucket/BurnBucket", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDefineBucket{},
		&MsgGenerateBucket{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeactivateBucket{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBurnBucket{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
