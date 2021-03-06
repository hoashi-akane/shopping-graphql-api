package model

type Goods struct {
	ID        int     `json:"id"`
	GoodsName string  `json:"goodsName"`
	Price     int     `json:"price"`
	Stock     int     `json:"stock"`
	BrandId   *int `json:"brandId"`
	GoodsCategoryId *int `json:"goodsCategoryId"`
}

type NewGoods struct {
	GoodsName string `json:"goodsName"`
	Price     int    `json:"price"`
	Stock     int    `json:"stock"`
	BrandId	  *int	`json:"brandId"`
	GoodsCategoryId *int `json:"goodsCategoryId"`
}