package model

type File struct {
	Name   string `json:"name"`
	Files  []File `json:"files"`
	IsFile bool   `json:"isFile"`
}
