package pages

import (
	"todos-vecty/components"
	"todos-vecty/model"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

type footer struct {
	vecty.Core
}

func (f *footer) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Style("display", "flex"),
		),
		elem.Span(
			vecty.Markup(),
			vecty.Text("Show:"),
		),
		components.NewFilterLink(model.All, "All"),
		components.NewFilterLink(model.Active, "Active"),
		components.NewFilterLink(model.Completed, "Completed"),
	)
}
