package model

type Goods struct{
	GoodsName string `json:"name"`
	Price int `json:"price"`
	Stock int `json:"stock"`
}
