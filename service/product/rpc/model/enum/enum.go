package enum

type ProductCategory int32

const (
	ProductCategoryTool ProductCategory = iota + 1
	ProductCategoryFood
)
