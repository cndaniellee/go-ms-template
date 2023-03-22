package order

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/mr"
	"goms/app/order/rpc/orderclient"
	"goms/app/product/rpc/productclient"
	"goms/common/request"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"goms/app/order/api/internal/svc"
	"goms/app/order/api/internal/types"

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
		l.Logger.Error(errors.Wrap(err, "user id parse failed"))
		err = response.ErrResp(0, ordercode.List, response.InternalError, err.Error())
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
			err = response.ErrResp(1, ordercode.List, response.InternalError, s.Message())
		default:
			l.Logger.Error(errors.Wrap(err, "order rpc call failed"))
			err = response.ErrResp(2, ordercode.List, response.ServiceError, s.Message())
		}
		return
	}

	// 使用MapReduce分别到产品服务获取产品数据
	list, err := mr.MapReduce[*orderclient.ListItem, types.ListItem, []types.ListItem](func(source chan<- *orderclient.ListItem) {
		for _, item := range reply.List {
			source <- item
		}
	}, func(item *orderclient.ListItem, writer mr.Writer[types.ListItem], cancel func(error)) {
		itemIds := make([]int64, len(item.Products))
		products := make([]types.ProductsResp, len(item.Products))
		for i, product := range item.Products {
			itemIds[i] = product.Id
			products[i] = types.ProductsResp{
				ID:     product.Id,
				Amount: product.Amount,
				Price:  product.Price,
			}
		}
		// 调用RPC服务
		byIds, err := l.svcCtx.ProductRpc.ListByIds(l.ctx, &productclient.ListByIdsReq{
			Ids: itemIds,
		})
		if err != nil {
			cancel(err)
		}
		// 将获取到的产品信息会写
		for i, product := range byIds.List {
			products[i].Title = product.Title
		}
		listItem := types.ListItem{
			ID:          item.Id,
			Status:      item.Status,
			Products:    products,
			TotalAmount: item.TotalAmount,
			TotalPrice:  item.TotalPrice,
		}
		writer.Write(listItem)
	}, func(pipe <-chan types.ListItem, writer mr.Writer[[]types.ListItem], cancel func(error)) {
		list := make([]types.ListItem, 0, len(reply.List))
		for _, item := range list {
			list = append(list, item)
		}
		writer.Write(list)
	})
	if err != nil {
		if s, ok := status.FromError(err); ok {
			switch s.Code() {
			case codes.Aborted:
				err = response.ErrResp(3, ordercode.List, response.InternalError, s.Message())
			default:
				l.Logger.Error(errors.Wrap(err, "product rpc call failed"))
				err = response.ErrResp(4, ordercode.List, response.ServiceError, s.Message())
			}
		} else {
			l.Logger.Error(errors.Wrap(err, "map reduce process failed"))
			err = response.ErrResp(5, ordercode.List, response.ServiceError, s.Message())
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
