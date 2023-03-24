package product

import (
	"context"
	"github.com/pkg/errors"
	"goms/service/product/rpc/productclient"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"goms/service/product/api/internal/svc"
	"goms/service/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"goms/common/response"
	"goms/common/response/errcode/productcode"
)

type RemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLogic {
	return &RemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveLogic) Remove(req *types.IdReq) (err error) {

	// 调用RPC服务
	_, err = l.svcCtx.ProductRpc.Remove(l.ctx, &productclient.IdReq{
		Id: req.ID,
	})
	if err != nil {
		switch s, _ := status.FromError(err); s.Code() {
		case codes.Aborted:
			err = response.ErrResp(1, productcode.Remove, response.InternalError, s.Message())
		default:
			l.Logger.Error(errors.Wrap(err, "product rpc call failed"))
			err = response.ErrResp(2, productcode.Remove, response.ServiceError, s.Message())
		}
		return
	}

	return
}
