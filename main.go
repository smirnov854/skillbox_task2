package main

import "fmt"

type SetName interface {
	SetName(name int)
}

type SetPath interface {
	SetPath(path string)
}


type Category struct {
	id int
	name int
	path string
}

func (cat *Category) SetName(name int){
		cat.name = name
}

func (cat *Category) SetPath(path string){
	cat.path = path
}

func (cat *Category) getName() int{
	return cat.name
}


type SetCategoryId interface {
	SetCategoryId(id int)
}


type Product struct {
	id int
	categoryId int
	name string
}

func (prod *Product) SetCategoryId(categoryId int){
	prod.categoryId = categoryId
}


func main(){
	firstCategory:= Category{5,11,"tetst"}
	secondCategory:= Category{55,33,"testcat"}
	firstProduct := Product{id: 1,categoryId: 5,name:"first"}
	secondProduct:= Product{id:2, categoryId: 11, name:"second"}

	fmt.Println(firstCategory.name)
	fmt.Println(secondCategory.path)

	firstCategory.SetName(123)
	secondCategory.SetPath("test_path")

	fmt.Println(firstCategory.name)
	fmt.Println(secondCategory.path)


	fmt.Println(firstProduct)
	fmt.Println(secondProduct)

	firstProduct.SetCategoryId(333)
	secondProduct.SetCategoryId(654)

	fmt.Println(firstProduct)
	fmt.Println(secondProduct)

}
