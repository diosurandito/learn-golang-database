package repository

import (
	"context"
	"fmt"
	golang_database "golang-database"
	"golang-database/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(golang_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "dio@test.com",
		Comment: "Test Repository",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestCommentFindById(t *testing.T) {
	commentRepository := NewCommentRepository(golang_database.GetConnection())
	ctx := context.Background()
	id := int32(35)

	comment, err := commentRepository.FindById(ctx, id)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestCommentFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(golang_database.GetConnection())
	ctx := context.Background()

	comments, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}
	for _, comment := range comments {
		fmt.Println(comment)

	}
}
