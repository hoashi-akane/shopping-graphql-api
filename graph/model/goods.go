package model

type Goods struct{
	Id int `json:"id"`
	GoodsName string `json:"name"`
	Price int `json:"price"`
	Stock int `json:"stock"`
}
