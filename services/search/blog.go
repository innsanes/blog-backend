package search

import (
	"blog-backend/structs/msearch"
	"fmt"

	"github.com/meilisearch/meilisearch-go"
)

func CreateBlogIndex(client meilisearch.ServiceManager) (err error) {
	info, err := client.Index("blogs").FetchInfo()
	if err != nil || info == nil {
		fmt.Println("Blog index not found")
		_, err = client.CreateIndex(&meilisearch.IndexConfig{
			Uid:        "blogs",
			PrimaryKey: "id",
		})
		fmt.Println("Blog index created", info)
		if err != nil {
			fmt.Println("Blog index created error", err)
			return err
		}
		fmt.Println("Blog index created")
		return nil
	}
	return
}

func InsertBlog(client meilisearch.ServiceManager, document *msearch.Blog) (err error) {
	index := client.Index("blogs")
	primaryKey := "id"
	_, err = index.AddDocuments(document, &primaryKey)
	if err != nil {
		return err
	}
	return
}

func UpdateBlog(client meilisearch.ServiceManager, document *msearch.Blog) (err error) {
	index := client.Index("blogs")
	primaryKey := "id"
	_, err = index.UpdateDocuments(document, &primaryKey)
	if err != nil {
		return err
	}
	return
}

func DeleteBlog(client meilisearch.ServiceManager, id string) (err error) {
	index := client.Index("blogs")
	_, err = index.DeleteDocument(id)
	if err != nil {
		return err
	}
	return
}

func SearchBlog(client meilisearch.ServiceManager, query string) (result *meilisearch.SearchResponse, err error) {
	index := client.Index("blogs")
	searchRes, err := index.Search(query, &meilisearch.SearchRequest{
		AttributesToCrop: []string{"name", "content"},
		CropLength:       15,
		CropMarker:       "...",
		HighlightPreTag:  "<em>",
		HighlightPostTag: "</em>",
	})
	if err != nil {
		return result, err
	}
	return searchRes, nil
}
