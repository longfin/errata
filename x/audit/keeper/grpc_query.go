package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mattverse/errata/x/audit/types"
)

type queryServer struct {
	keeper Keeper
}

// NewQueryServerImpl returns an implementation of the bond QueryServer interface
// for the provided Keeper.
func NewQueryServerImpl(keeper Keeper) types.QueryServer {
	return &queryServer{keeper: keeper}
}

func (q queryServer) Protocol(ctx context.Context, req *types.QueryProtocolRequest) (*types.QueryProtocolResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	protocolID := req.Id
	protocol, err := q.keeper.GetProtocolByID(sdkCtx, protocolID)
	if err != nil {
		return nil, err
	}

	return &types.QueryProtocolResponse{
		Title:       protocol.Title,
		Description: protocol.Description,
		SourceCode:  protocol.SourceCode,
		ProjectHome: protocol.ProjectHome,
		Category:    protocol.Category,
		AttackPool:  &protocol.AttackPool,
		DefensePool: &protocol.DefensePool,
	}, nil
}
