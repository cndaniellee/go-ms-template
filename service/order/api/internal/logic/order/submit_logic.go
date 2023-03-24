package order

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"goms/common/request"
	"goms/common/response/errcode/usercode"
	"goms/service/order/rpc/orderclient"
	"goms/service/order/rpc/pb/order"
	"goms/service/product/rpc/pb/product"
	"goms/service/product/rpc/productclient"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"

	"goms/service/order/api/internal/svc"
	"goms/service/order/api/internal/types"

	"github.com/CNDanielLee/go-utils"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"
	"goms/common/response"
	"goms/common/response/errcode/ordercode"
)

type SubmitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubmitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitLogic {
	return &SubmitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubmitLogic) Submit(req *types.SubmitReq) (err error) {

	// 解析用户ID
	userId, err := request.ParseUserId(l.ctx)
	if err != nil {
		l.Logger.Error(errors.Wrap(err, "user id parse failed"))
		err = response.ErrResp(1, usercode.Register, response.InternalError, err.Error())
		return
	}

	// 构建RPC连接信息
	orderRpc, err := l.svcCtx.Config.OrderRpcConf.BuildTarget()
	if err != nil {
		l.Logger.Error(errors.Wrap(err, "order rpc build failed"))
		err = response.ErrResp(2, ordercode.Submit, response.InternalError, err.Error())
		return
	}
	productRpc, err := l.svcCtx.Config.ProductRpcConf.BuildTarget()
	if err != nil {
		l.Logger.Error(errors.Wrap(err, "product rpc build failed"))
		err = response.ErrResp(3, ordercode.Submit, response.InternalError, err.Error())
		return
	}

	// 构建Saga事务
	reply, err := l.svcCtx.DtmClient.NewGid(l.ctx, &emptypb.Empty{})
	if err != nil {
		l.Logger.Error(errors.Wrap(err, "dtm gid fetch failed"))
		err = response.ErrResp(4, ordercode.Submit, response.InternalError, err.Error())
		return
	}
	saga := dtmgrpc.NewSagaGrpc(l.svcCtx.Config.DtmService, reply.Gid)

	// 依次添加每个产品的扣减事务
	ids := make([]int64, len(req.Products))
	products := make([]*orderclient.Product, len(req.Products))
	for i, item := range req.Products {
		saga = saga.Add(
			fmt.Sprint(productRpc, product.Product_Deduct_FullMethodName),
			fmt.Sprint(productRpc, product.Product_DeductRollback_FullMethodName),
			&productclient.DeductReq{
				Id:     item.ID,
				Amount: item.Amount,
			})
		ids[i] = item.ID
		products[i] = &orderclient.Product{
			Id:     item.ID,
			Amount: item.Amount,
		}
	}

	// 调用RPC服务
	reply2, err := l.svcCtx.ProductRpc.ListByIds(l.ctx, &productclient.ListByIdsReq{
		Ids: ids,
	})
	if err != nil {
		switch s, _ := status.FromError(err); s.Code() {
		case codes.Aborted:
			err = response.ErrResp(5, ordercode.Submit, response.InternalError, s.Message())
		default:
			l.Logger.Error(errors.Wrap(err, "product rpc call failed"))
			err = response.ErrResp(6, ordercode.Submit, response.ServiceError, s.Message())
		}
		return
	}

	// 回写产品价格
	for i, item := range reply2.List {
		products[i].Price = item.Price
	}

	// 添加订单创建事务
	saga = saga.Add(
		fmt.Sprint(orderRpc, order.Order_Create_FullMethodName),
		fmt.Sprint(orderRpc, order.Order_CreateRollback_FullMethodName),
		&orderclient.CreateReq{
			UserId:    userId,
			RefNo:     fmt.Sprint(time.Now().Format("20060102150405"), utils.RandStr(6, utils.Number)),
			Products:  products,
			Consignee: req.Consignee,
			Phone:     req.Phone,
			Address:   req.Address,
		})

	// 提交Saga事务
	if err = saga.Submit(); err != nil {
		l.Logger.Error(errors.Wrap(err, "dtm saga submit failed"))
		err = response.ErrResp(7, ordercode.Submit, response.InternalError, err.Error())
		return
	}

	return
}
