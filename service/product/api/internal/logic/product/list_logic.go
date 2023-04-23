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

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.ListReq) (resp *types.ListResp, err error) {

	// 调用RPC服务
	reply, err := l.svcCtx.ProductRpc.List(l.ctx, &productclient.ListReq{
		Search:   req.Search,
		Category: req.Category,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		switch s, _ := status.FromError(err); s.Code() {
		case codes.Aborted:
			err = response.ErrResp(1, productcode.List, response.InternalError, s.Message())
		default:
			l.Error(errors.Wrap(err, "product rpc call failed"))
			err = response.ErrResp(2, productcode.List, response.ServiceError, s.Message())
		}
		return
	}

	// 转换数据
	list := make([]types.ListItem, len(reply.List))
	for i, item := range reply.List {
		list[i] = types.ListItem{
			ID:       item.Id,
			Title:    item.Title,
			Category: item.Category,
			Stock:    item.Stock,
			Price:    item.Price,
		}
	}

	resp = &types.ListResp{
		List: list,
		Page: types.Page{
			Page:     reply.Page.Page,
			PageSize: reply.Page.PageSize,
			Total:    reply.Page.Total,
		},
	}

	return
}
