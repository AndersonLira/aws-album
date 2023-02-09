package service_test

import (
	"testing"

	"github.com/andersonlira/album/service"
)

func TestAddByKeyEmpty(t *testing.T) {
	list := service.AddByKey("")
	size := len(list)

	if size != 0 {
		t.Fatalf("List should be empty, but size %d", size)
	}

}
func TestAddByKey(t *testing.T) {
	list := service.AddByKey("abc/def/ghi")
	size := len(list)

	if size != 1 {
		t.Fatalf("It should have 1 item, but was %d", size)
	}

	folder := list[0]

	if folder.Name != "abc" {
		t.Fatalf("Folder name should be 'abc' but %s", folder.Name)
	}

	sub1 := folder.Files[0]

	if sub1.Name != "def" {
		t.Fatalf("Sub folder should be named 'def' but was %s", sub1.Name)
	}

	sub2 := sub1.Files[0]

	if sub2.Name != "ghi" {
		t.Fatalf("Sub folder should be named 'def' but was %s", sub2.Name)
	}

	service.AddByKey("abc/123/ghi")
	list = service.AddByKey("abc/123/ijk")
	size = len(list)

	if size != 1 {
		t.Fatalf("List size should still be 1 but %d", size)
	}

	folder = list[0]
	sub1_2 := folder.Files[1]

	if sub1_2.Name != "123" {
		t.Fatalf("sub 1.2 should be named '123' but %s", sub1_2.Name)
	}

	name2_2 := sub1_2.Files[1].Name

	if name2_2 != "ijk" {
		t.Fatalf("Name 2 2 should be 'ijk' but was %s", name2_2)
	}
}
