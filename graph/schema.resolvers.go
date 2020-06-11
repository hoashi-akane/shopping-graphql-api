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
	ins, err := r.DB.Prepare("INSERT INTO goods(goods_name, price, stock) VALUES(?,?,?);")
	if err != nil {
		log.Fatal(err)
	}
	ins.Exec(good.GoodsName, good.Price, good.Stock)
	return good, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	//r.DB.Exec("SELECT * FROM ")
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Goodes(ctx context.Context) ([]*model.Goods, error) {
	rows, err := r.DB.Query("SELECT id, goods_name, price, stock FROM goods;")
	if err != nil {
		return nil, err
	}
	var results []*model.Goods
	for rows.Next() {
		var goods *model.Goods
		// 参照渡しをすることでコピー処理を発生させない。
		// 構造体の配列に入れるための一時領域のため書き換えても問題ない。
		goods = &model.Goods{}
		err := rows.Scan(&goods.Id, &goods.GoodsName, &goods.Price, &goods.Stock)
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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
