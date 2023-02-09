package service

import (
	"strings"

	"github.com/andersonlira/album/model"
)

var list = []model.File{}

func AddByKey(key string) []model.File {
	if len(key) == 0 {
		return list
	}
	parts := strings.Split(key, "/")

	var parent *model.File
	for idx, name := range parts {
		file := findFolder(name, parent)

		if file == nil {
			isFile := idx == len(parts)-1
			newFolder := model.File{
				Name:   name,
				Files:  []model.File{},
				IsFile: isFile,
			}
			if parent == nil {
				list = append(list, newFolder)
				file = &list[len(list)-1]
			} else {
				parent.Files = append(parent.Files, newFolder)
				file = &parent.Files[len(parent.Files)-1]
			}
		}
		parent = file
	}
	return list
}

func findFolder(name string, parent *model.File) *model.File {
	ref := list
	if parent != nil {
		ref = parent.Files
	}

	for idx, v := range ref {
		if v.Name == name {
			return &ref[idx]
		}
	}
	return nil
}
