package Controller

import "github.com/kataras/iris"

type IndexController struct{}

func (index *IndexController) Index(ctx iris.Context) {
	ctx.View("index.html")
}
