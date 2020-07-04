package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hoashi-akane/shopping-graphql/graph/generated"
	"github.com/hoashi-akane/shopping-graphql/graph/model"
)

func (r *goodsResolver) Brand(ctx context.Context, obj *model.Goods) (*model.Brands, error) {
	// オブジェクト内のBrandIdがセットされているか識別
	if obj.BrandId == nil {
		return nil, nil
	}
	stmt, err := r.DB.Prepare("SELECT id, brand_name FROM brands WHERE id = ?")
	if err != nil {
		return nil, err
	}
	brands := &model.Brands{}
	err = stmt.QueryRow(obj.BrandId).Scan(&brands.ID, &brands.BrandName)
	return brands, err
}

func (r *goodsResolver) Category(ctx context.Context, obj *model.Goods) (*model.GoodsCategorys, error) {
	// オブジェクト内のBrandIdがセットされているか識別
	if obj.GoodsCategoryId == nil {
		return nil, nil
	}
	stmt, err := r.DB.Prepare("SELECT id, category_name FROM goods_category WHERE id = ?")
	if err != nil {
		return nil, err
	}
	category := &model.GoodsCategorys{}
	err = stmt.QueryRow(obj.BrandId).Scan(&category.ID, &category.CategoryName)
	return category, err
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: input.UserID, // fix this line
	}
	return todo, nil
}

func (r *mutationResolver) CreateGoods(ctx context.Context, input model.NewGoods) (*model.Goods, error) {
	good := &model.Goods{
		GoodsName:       input.GoodsName,
		Price:           input.Price,
		Stock:           input.Stock,
		BrandId:         input.BrandId,
		GoodsCategoryId: input.GoodsCategoryId,
	}
	ins, err := r.DB.Prepare("INSERT INTO goods(goods_name, price, stock, brand_id, category_id) VALUES(?,?,?,?,?);")
	if err != nil {
		log.Fatal(err)
	}
	ins.Exec(good.GoodsName, good.Price, good.Stock, good.BrandId, good.GoodsCategoryId)
	return good, nil
}

func (r *mutationResolver) DeleteGoods(ctx context.Context, input int) (bool, error) {
	result := false
	ins, err := r.DB.Prepare("DELETE FROM goods WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	_, err = ins.Exec(input)
	if err == nil {
		result = true
	}
	return result, nil
}

func (r *mutationResolver) CreateBrand(ctx context.Context, input model.NewBrand) (*model.Brands, error) {
	brand := &model.Brands{
		BrandName: input.BrandName,
	}
	stmt, err := r.DB.Prepare("INSERT INTO brands(brand_name) VALUES(?)")
	if err != nil {
		fmt.Errorf("DBエラー")
	}
	stmt.Exec(brand.BrandName)
	return brand, nil
}

func (r *mutationResolver) CreateCategory(ctx context.Context, input model.NewCategory) (*model.GoodsCategorys, error) {
	category := &model.GoodsCategorys{
		CategoryName: input.CategoryName,
	}
	stmt, err := r.DB.Prepare("INSERT INTO goods_category(category_name) VALUES(?)")
	if err != nil {
		fmt.Errorf("DBエラー")
	}
	stmt.Exec(category.CategoryName)
	return category, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	//r.DB.Exec("SELECT * FROM ")
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Goodes(ctx context.Context) ([]*model.Goods, error) {
	rows, err := r.DB.Query("SELECT id, goods_name, price, stock, brand_id, category_id FROM goods;")
	if err != nil {
		return nil, err
	}
	var results []*model.Goods
	for rows.Next() {
		var goods *model.Goods
		// 参照渡しをすることでコピー処理を発生させない。
		// 構造体の配列に入れるための一時領域のため書き換えても問題ない。
		goods = &model.Goods{}

		err := rows.Scan(&goods.ID, &goods.GoodsName, &goods.Price, &goods.Stock, &goods.BrandId, &goods.GoodsCategoryId)
		if err != nil {
			panic(fmt.Errorf("DBエラー"))
		}
		results = append(results, goods)
	}
	return results, nil
}

func (r *queryResolver) FindGood(ctx context.Context, id int) (*model.Goods, error) {
	stmt, err := r.DB.Prepare("SELECT id, goods_name, price, stock, brand_id, category_id FROM goods WHERE id = ?")
	if err != nil {
		return nil, err
	}
	goods := &model.Goods{}
	err = stmt.QueryRow(id).Scan(&goods.ID, &goods.GoodsName, &goods.Price, &goods.Stock, &goods.BrandId, &goods.GoodsCategoryId)
	if err != nil {
		panic(fmt.Errorf("DBエラー"))
	}
	return goods, nil
}

func (r *queryResolver) FindBrand(ctx context.Context, id int) (*model.Brands, error) {
	stmt, err := r.DB.Prepare("SELECT id, brand_name FROM brands WHERE id = ?")
	if err != nil {
		return nil, err
	}
	brands := &model.Brands{}
	err = stmt.QueryRow(id).Scan(&brands.ID, &brands.BrandName)
	if err != nil {
		panic(fmt.Errorf("DBエラー"))
	}
	return brands, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}

// Goods returns generated.GoodsResolver implementation.
func (r *Resolver) Goods() generated.GoodsResolver { return &goodsResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type goodsResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
