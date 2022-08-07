package db

import (
	"context"
	"fmt"
	"strconv"

	"github.com/amar-jay/todolist_v2/src/model"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
  "github.com/go-pg/pg/extra/pgdebug"
)


type DB  *pg.DB

func Connection() *pg.DB {
  opt, err := pg.ParseURL("postgres://postgres:Abdel-manan1978@localhost:5432/todolistv2")
if err != nil {
   panic(err)
} else {
  fmt.Println("Connected to DB")
}

db := pg.Connect(opt)
// check if running in Background

db.AddQueryHook(pgdebug.DebugHook{
  Verbose: true,
})

ctx := context.Background()

if err:= db.Ping(ctx); err != nil {
  panic(err)
} else {
  fmt.Println("DB running...")
}

return db
}
// create database tables based on defined schema
func createSchema(db *pg.DB) error {
  model := []interface{}{
    (*model.Todo)(nil),
  }

  for _, m := range model {
    err := db.Model(m).CreateTable(&orm.CreateTableOptions{
      Temp: true,
    })
    if err != nil {
      return err
    } 
  }

  fmt.Println("schema created")
  return nil
}

func Create(db *pg.DB, text string, done bool, id string) *model.Todo {
  err := createSchema(db)
  if err != nil {
    panic(err)
  }

  todo := &model.Todo{
    ID: fmt.Sprint(id),
    Text: text,
    Done: done,
  }

  res, err := db.Model(todo).Insert()
  if res == nil || err != nil {
    fmt.Println("error creating todo to DB")
    panic(err)
  }

  return todo
}



func GetByID(db *pg.DB, id string) *model.Todo {

  if _, err := strconv.Atoi(id); err != nil {
    fmt.Println("id not a number"+ id)
    panic(err)
  } 
 
  // select todo by Primary Key
  todo := &model.Todo{ID: id}
  err := db.Model(todo).WherePK().Select()
  if err != nil {
    panic(err)
  }

  fmt.Println(todo.Text)

  return todo
}

func GetAll(db *pg.DB) []*model.Todo {
  var todos []*model.Todo
  err := db.Model(&todos).Select()
  if err != nil {
    panic(err)
  }

  for _, todo := range todos {
    fmt.Println(todo.Text)
  }

  return todos
}




