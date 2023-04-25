package order

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/mr"
	"goms/common/request"
	"goms/service/order/rpc/orderclient"
	"goms/service/product/rpc/productclient"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"goms/service/order/api/internal/svc"
	"goms/service/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"goms/common/response"
	"goms/common/response/errcode/ordercode"
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

	// 解析用户ID
	userId, err := request.ParseUserId(l.ctx)
	if err != nil {
		l.Error(errors.Wrap(err, "user id parse failed"))
		err = response.ErrResp(1, ordercode.List, response.InternalError, err.Error())
		return
	}

	// 调用RPC服务
	reply, err := l.svcCtx.OrderRpc.List(l.ctx, &orderclient.ListReq{
		UserId:   userId,
		Status:   req.Status,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		switch s, _ := status.FromError(err); s.Code() {
		case codes.Aborted:
			err = response.ErrResp(2, ordercode.List, response.InternalError, s.Message())
		default:
			l.Error(errors.Wrap(err, "order rpc call failed"))
			err = response.ErrResp(3, ordercode.List, response.ServiceError, s.Message())
		}
		return
	}

	// 转换数据
	list := make([]types.ListItem, len(reply.List))
	for i, item := range reply.List {
		products := make([]types.ProductsResp, len(item.Products))
		for j, product := range item.Products {
			products[j] = types.ProductsResp{
				ID:     product.Id,
				Price:  product.Price,
				Amount: product.Amount,
			}
		}
		list[i] = types.ListItem{
			ID:          item.Id,
			Status:      item.Status,
			Products:    products,
			TotalAmount: item.TotalAmount,
			TotalPrice:  item.TotalPrice,
		}
	}

	// 使用MapReduce分别到产品服务获取产品数据
	err = mr.MapReduceVoid[types.ListItem, any](func(source chan<- types.ListItem) {
		for _, item := range list {
			source <- item
		}
	}, func(item types.ListItem, writer mr.Writer[any], cancel func(error)) {
		productIds := make([]int64, len(item.Products))
		for i, product := range item.Products {
			productIds[i] = product.ID
		}
		// 调用RPC服务
		byIds, err := l.svcCtx.ProductRpc.ListByIds(l.ctx, &productclient.ListByIdsReq{
			Ids: productIds,
		})
		if err != nil {
			cancel(err)
		}
		// 将获取到的产品信息回写
		for i, product := range byIds.List {
			item.Products[i].Title = product.Title
		}
	}, func(pipe <-chan any, cancel func(error)) {})
	if err != nil {
		if s, ok := status.FromError(err); ok {
			switch s.Code() {
			case codes.Aborted:
				err = response.ErrResp(4, ordercode.List, response.InternalError, s.Message())
			default:
				l.Error(errors.Wrap(err, "product rpc call failed"))
				err = response.ErrResp(5, ordercode.List, response.ServiceError, s.Message())
			}
		} else {
			l.Error(errors.Wrap(err, "map reduce process failed"))
			err = response.ErrResp(6, ordercode.List, response.InternalError, s.Message())
		}
		return
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
