package components

import (
	"todos-vecty/model"
	"todos-vecty/storeutil"

	"github.com/hexops/vecty"
)

type visibleTodoList struct {
	vecty.Core
}

func NewVisibleTodoList() *visibleTodoList {
	c := &visibleTodoList{}

	storeutil.TodosDispacher.Subscribe(func() {
		vecty.Rerender(c)
	})

	storeutil.FilterDispacher.Subscribe(func() {
		vecty.Rerender(c)
	})

	return c
}

func (a *visibleTodoList) Render() vecty.ComponentOrHTML {
	todosState, _ := storeutil.TodosDispacher.StateOf(storeutil.TodosReducer).([]model.Todo)
	filterState, _ := storeutil.FilterDispacher.StateOf(storeutil.FilterReducer).(model.FilterType)

	todos := make([]model.Todo, 0)

	if filterState == model.All {
		todos = append(todos, todosState...)
	} else if filterState == model.Active {

		for _, todo := range todosState {
			if !todo.Completed {
				todos = append(todos, todo)
			}
		}
	} else if filterState == model.Completed {
		for _, todo := range todosState {
			if todo.Completed {
				todos = append(todos, todo)
			}
		}
	}

	return &todoList{
		Todos: todos,
		OnClick: func(id int) {
			storeutil.TodosDispacher.Dispatch(storeutil.TodosReducer.Complete.With(storeutil.NewCompleteTodoAction(id)))
		},
	}
}
