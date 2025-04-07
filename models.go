package main

import "fmt"

type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"is_done"`
}
