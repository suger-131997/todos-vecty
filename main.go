package main

import (
	"todos-vecty/pages"
	"todos-vecty/storeutil"

	"github.com/hexops/vecty"
)

func main() {
	vecty.SetTitle("Todos")

	storeutil.TodosStoreInit()
	storeutil.FilterStoreInit()

	p := &pages.PageView{}
	vecty.RenderBody(p)
}
