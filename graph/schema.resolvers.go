package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hoashi-akane/shopping-graphql/graph/generated"
	"github.com/hoashi-akane/shopping-graphql/graph/model"
)

func (r *goodsResolver) ID(ctx context.Context, obj *model.Goods) (string, error) {
	panic(fmt.Errorf("not implemented"))
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
		GoodsName: input.GoodsName,
		Price:     input.Price,
		Stock:     input.Stock,
	}
	return good, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	//r.DB.Exec("SELECT * FROM ")
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Goodes(ctx context.Context) ([]*model.Goods, error) {
	rows, err := r.DB.Query("SELECT goods_name, price, stock FROM goods;")
	if err != nil {
		return nil, err
	}
	var results []*model.Goods
	for rows.Next() {
		var goods *model.Goods
		// 参照渡しをすることでコピー処理を発生させない。
		// 構造体の配列に入れるための一時領域のため書き換えても問題ない。
		goods = &model.Goods{}
		err := rows.Scan(&goods.GoodsName, &goods.Price, &goods.Stock)
		if err != nil {
			panic(fmt.Errorf("DBエラー"))
		}
		results = append(results, goods)
	}
	return results, nil
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *goodsResolver) Price(ctx context.Context, obj *model.Goods) (int, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *goodsResolver) GoodsName(ctx context.Context, obj *model.Goods) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *goodsResolver) Stock(ctx context.Context, obj *model.Goods) (int, error) {
	panic(fmt.Errorf("not implemented"))
}
