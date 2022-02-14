package repository

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/ardhihdra/go-clean-arch/entity"
)

const (
	projectId      = "go-clean-arch-923f2"
	collectionName = "posts"
)

type repo struct{}

func NewFirestoreRepository() PostRepository {
	return &repo{}
}

func createCtx() (context.Context, *firestore.Client, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Firestore client creation failed: %v", err)
		return nil, client, err
	}
	return ctx, client, nil
}

func (r repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx, client, err := createCtx()

	defer client.Close()
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatalf("Failed adding new post: %v", err)
		return nil, err
	}

	return post, err
}

func (r repo) FindAll() ([]entity.Post, error) {
	ctx, client, err := createCtx()
	var posts []entity.Post

	defer client.Close()
	iterator := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := iterator.Next()
		if err != nil {
			break
		}
		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}
	return posts, err
}

func (r repo) FindByID(id int64) (entity.Post, error) {
	ctx, client, err := createCtx()
	var posts entity.Post

	defer client.Close()
	states := client.Collection(collectionName)
	query := states.Where("ID", "==", id)
	iterator := query.Documents(ctx)
	defer iterator.Stop()
	for {
		doc, err := iterator.Next()
		if err != nil {
			break
		}
		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = post
	}
	return posts, err
}
