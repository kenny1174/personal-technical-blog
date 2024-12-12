package main

import (
	"technicalBlog/model"
	"technicalBlog/routes"
)

func main() {
	//引用数据库
	model.InitDb()

	routes.InitRouter()
}
