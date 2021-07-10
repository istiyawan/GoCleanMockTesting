package repository

import (
"cleango/entity"
"context"
"google.golang.org/api/iterator"
"log"
"os"

"cloud.google.com/go/firestore"
)

type repo struct{}

func NewFirestoreRepository() PostRepository {
	return &repo{}
}

const (
	projectId      = "goclean-a661b"
	collectionName = "posts"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	var credentialsFilePath = "D:\\wawan\\go_project\\cleango\\gocleanjson.json"
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", credentialsFilePath)

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("failed to create firestore client:%v", err)
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatalf("failed adding a new post :%v", err)
		return nil, err
	}

	return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	var credentialsFilePath = "D:\\wawan\\go_project\\cleango\\gocleanjson.json"
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", credentialsFilePath)

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("failed to create firestore client:%v", err)
		return nil, err
	}

	defer client.Close()
	var posts []entity.Post
	itr := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := itr.Next()
		//if err != nil {
		//	log.Fatalf("failed to iterate the list of posts:%v", err)
		//	return nil, err
		//}
		if err == iterator.Done{
			break
		}

		if err != nil {
			log.Fatalf("failed to iterate the list of posts:%v", err)
			return nil, err
		}

		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}
