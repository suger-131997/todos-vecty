package components

import (
	"todos-vecty/model"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

type todoList struct {
	vecty.Core

	Todos   []model.Todo `vecty:"prop"`
	OnClick func(int)    `vecty:"prop"`
}

func (t *todoList) Render() vecty.ComponentOrHTML {
	var items vecty.List

	for _, todo := range t.Todos {
		items = append(items, &todoListItem{
			Text:      todo.Title,
			Completed: todo.Completed,
			OnClick: func(id int, f func(int)) func(*vecty.Event) {
				return func(e *vecty.Event) {
					f(id)
				}
			}(todo.Id, t.OnClick),
		})
	}

	return elem.UnorderedList(
		vecty.Markup(),
		items,
	)
}
