package product

import (
	"context"
	"github.com/pkg/errors"
	"goms/app/product/rpc/productclient"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"goms/app/product/api/internal/svc"
	"goms/app/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"goms/common/response"
	"goms/common/response/errcode/productcode"
)

type EditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditLogic {
	return &EditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditLogic) Edit(req *types.EditReq) (resp *types.IdResp, err error) {

	// 调用RPC服务
	reply, err := l.svcCtx.ProductRpc.Edit(l.ctx, &productclient.EditReq{
		Id:          req.ID,
		Title:       req.Title,
		Category:    req.Category,
		Stock:       req.Stock,
		Price:       req.Price,
		Description: req.Description,
	})
	if err != nil {
		switch s, _ := status.FromError(err); s.Code() {
		case codes.Aborted:
			err = response.ErrResp(0, productcode.Edit, response.InternalError, s.Message())
		default:
			l.Logger.Error(errors.Wrap(err, "product rpc call failed"))
			err = response.ErrResp(1, productcode.Edit, response.ServiceError, s.Message())
		}
		return
	}

	resp = &types.IdResp{ID: reply.Id}

	return
}
