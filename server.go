package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hoashi-akane/shopping-graphql/graph"
	"github.com/hoashi-akane/shopping-graphql/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	// DBアクセス
	db, err := sql.Open("mysql", DATABASESERVER)
	if err != nil{
		log.Print("DB Error")
	}else{
		log.Print("実行完了")
	}
	defer db.Close()

	// DBを渡す。
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	// 確認用url
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// アクセス先
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

