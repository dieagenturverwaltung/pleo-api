package api

import "net/http"

type Api struct {
	client *http.Client
}

func New(client *http.Client) *Api {
	return &Api{
		client: client,
	}
}
