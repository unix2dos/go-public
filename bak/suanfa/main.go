package main

import (
	"fmt"

	"google.golang.org/genproto/googleapis/datastore/admin/v1"
)

type stu struct {
	Name string
}

func aa(v interface{}) {
	a := v.(type)
	fmt.Print(a)
	switch a:= v.(type) {
	case *stu,stu:
		a.Name
	}
}
