package order

import (
	"context"
	"github.com/pkg/errors"
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

	// 解析用户ID
	userId, err := request.ParseUserId(l.ctx)
	if err != nil {
		l.Logger.Error(errors.Wrap(err, "user id parse failed"))
		err = response.ErrResp(0, ordercode.List, response.InternalError, err.Error())
		return
	}

	// 调用RPC服务
	reply, err := l.svcCtx.OrderRpc.Detail(l.ctx, &orderclient.IdReq{
		Id:     req.ID,
		UserId: userId,
	})
	if err != nil {
		switch s, _ := status.FromError(err); s.Code() {
		case codes.NotFound:
			err = response.ErrResp(0, ordercode.Detail, response.NoneMatching, s.Message())
		case codes.Aborted:
			err = response.ErrResp(1, ordercode.Detail, response.InternalError, s.Message())
		default:
			l.Logger.Error(errors.Wrap(err, "order rpc call failed"))
			err = response.ErrResp(2, ordercode.Detail, response.ServiceError, s.Message())
		}
		return
	}

	ids := make([]int64, len(reply.Products))
	products := make([]types.ProductsResp, len(reply.Products))
	for i, item := range reply.Products {
		ids[i] = item.Id
		products[i] = types.ProductsResp{
			ID:     item.Id,
			Amount: item.Amount,
			Price:  item.Price,
		}
	}

	// 调用RPC服务
	reply2, err := l.svcCtx.ProductRpc.ListByIds(l.ctx, &productclient.ListByIdsReq{
		Ids: ids,
	})
	if err != nil {
		switch s, _ := status.FromError(err); s.Code() {
		case codes.Aborted:
			err = response.ErrResp(3, ordercode.Detail, response.InternalError, s.Message())
		default:
			l.Logger.Error(errors.Wrap(err, "product rpc call failed"))
			err = response.ErrResp(4, ordercode.Detail, response.ServiceError, s.Message())
		}
		return
	}

	// 回填数据
	for i, item := range reply2.List {
		products[i].Title = item.Title
	}

	resp = &types.DetailResp{
		ID:          reply.Id,
		Status:      reply.Status,
		Products:    products,
		Consignee:   reply.Consignee,
		Phone:       reply.Phone,
		Address:     reply.Address,
		TotalAmount: reply.TotalAmount,
		TotalPrice:  reply.TotalPrice,
		CreatedAt:   reply.CreatedAt,
	}

	return
}
