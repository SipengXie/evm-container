package logic

import (
	"context"
	"strconv"

	"evm-container/api/internal/svc"
	"evm-container/api/internal/types"
	"evm-container/common"
	"evm-container/rpc/types/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetLogic {
	return &ResetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetLogic) Reset(req *types.ResetRequest) (resp *types.ResetResponse, err error) {
	index, err := strconv.ParseInt(req.Index, 10, 32)
	if err != nil {
		return &types.ResetResponse{
			Code: err.Error(),
		}, err
	}
	res, err := l.svcCtx.EvmRpc.Reset(l.ctx, &rpc.ResetRequest{
		TxHash: common.Hex2Bytes(req.TxHash),
		Index:  int32(index),
		TxCtx:  []byte(req.TxCtx),
	})
	return &types.ResetResponse{
		Code: res.Code,
	}, err
}
