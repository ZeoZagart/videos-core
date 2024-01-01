package server

import (
	"context"
	"fmt"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func Search() []*youtube.SearchResult {
	context := context.Background()
	client := client(context).Search.List([]string{})
	caller := client.Context(context).MaxResults(20).RelevanceLanguage("en").Type("video")
	result, err := caller.Do()
	if err != nil {
		fmt.Println("Error finding videos: ", err)
		return nil
	}
	for _, item := range result.Items {
		fmt.Printf("Video: %+v\n", *item)
	}
	return result.Items
}

func Categories() []*youtube.VideoCategory {
	context := context.Background()
	client := client(context).VideoCategories.List([]string{})
	result, err := client.Context(context).Do()
	if err != nil {
		fmt.Println("Error finding categories: ", err)
		return nil
	}
	for _, item := range result.Items {
		fmt.Printf("videocategory: %+v\n", *item)
	}
	return result.Items
}

type YoutubeVideo struct {
	Name string
}

func client(ctxt context.Context) *youtube.Service {
	youtubeService, err := youtube.NewService(ctxt, option.WithAPIKey("AIzaSyCqq6W8RzisT1lNTkiBsnNTRpsv8E0WNi8"))
	if err != nil {
		panic(err)
	}
	return youtubeService
}
