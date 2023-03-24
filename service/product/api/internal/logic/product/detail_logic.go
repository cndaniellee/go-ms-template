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

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.IdReq) (resp *types.DetailResp, err error) {

	// 调用RPC服务
	reply, err := l.svcCtx.ProductRpc.Detail(l.ctx, &productclient.IdReq{
		Id: req.ID,
	})
	if err != nil {
		switch s, _ := status.FromError(err); s.Code() {
		case codes.NotFound:
			err = response.ErrResp(1, productcode.Detail, response.NoneMatching, s.Message())
		case codes.Aborted:
			err = response.ErrResp(2, productcode.Detail, response.InternalError, s.Message())
		default:
			l.Logger.Error(errors.Wrap(err, "product rpc call failed"))
			err = response.ErrResp(3, productcode.Detail, response.ServiceError, s.Message())
		}
		return
	}

	resp = &types.DetailResp{
		ID:          reply.Id,
		Title:       reply.Title,
		Category:    reply.Category,
		Stock:       reply.Stock,
		Price:       reply.Price,
		Description: reply.Description,
		CreatedAt:   reply.CreatedAt,
	}

	return
}
