package es

import (
	"context"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"goms/common/es"
	"goms/service/product/rpc/model"
	"goms/service/product/rpc/model/enum"
	"reflect"
	"strconv"
	"time"
)

type ProductES struct {
	index  string
	client *elastic.Client
}

func NewProductES(url string) *ProductES {
	index := "product"
	client, err := elastic.NewClient(elastic.SetURL(url))
	logx.Must(err)
	ctx := context.Background()
	// 检查索引
	exist, err := client.IndexExists(index).Do(ctx)
	logx.Must(err)
	if !exist {
		mapping := &es.IndexMapping{}
		mapping.Settings.NumberOfShards = 3
		mapping.Settings.NumberOfReplicas = 0
		properties := es.Properties{
			"id":          {"type": "long"},
			"title":       {"type": "text", "analyzer": es.Analyzer},
			"category":    {"type": "integer"},
			"last_update": {"type": "date", "format": "epoch_millis"},
		}
		mapping.Mappings.Properties = properties
		// 创建索引
		create, err := client.CreateIndex(index).BodyJson(mapping).Do(ctx)
		logx.Must(err)
		if !create.Acknowledged {
			logx.Must(errors.New("create product index not acknowledged"))
		}
	}
	return &ProductES{index: index, client: client}
}

func (es *ProductES) Upsert(ctx context.Context, product *model.Product) error {
	body := map[string]any{
		"id":          product.ID,
		"title":       product.Title,
		"category":    product.Category,
		"last_update": time.Now().UnixMilli(),
	}
	_, err := es.client.Index().Index(es.index).Id(strconv.FormatInt(product.ID, 10)).BodyJson(body).Do(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (es *ProductES) Delete(ctx context.Context, id int64) error {
	_, err := es.client.Delete().Index(es.index).Id(strconv.FormatInt(id, 10)).Do(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (es *ProductES) Search(ctx context.Context, search string, category enum.ProductCategory, page, pageSize int) ([]model.Product, int64, error) {
	// 构建查询
	query := elastic.NewBoolQuery()
	if search != "" {
		query.Filter(elastic.NewMatchQuery("title", search))
	}
	if category != 0 {
		query.Filter(elastic.NewTermQuery("category", category))
	}
	// 查找结果
	result, err := es.client.Search().
		Index(es.index).Query(query).
		Sort("last_update", false).
		From((page - 1) * pageSize).Size(pageSize).
		Pretty(true).Do(ctx)
	if err != nil {
		return nil, 0, err
	}
	products := make([]model.Product, result.TotalHits())
	for i, item := range result.Each(reflect.TypeOf(model.Product{})) {
		products[i] = item.(model.Product)
	}
	// 统计数量
	count, err := es.client.Count(es.index).Query(query).Do(ctx)
	if err != nil {
		return nil, 0, err
	}
	return products, count, nil
}
