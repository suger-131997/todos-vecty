package components

import (
	"todos-vecty/model"
	"todos-vecty/storeutil"

	"github.com/hexops/vecty"
)

type filterLink struct {
	vecty.Core

	Type  model.FilterType `vecty:"prop"`
	Label string           `vecty:"prop"`
}

func NewFilterLink(f model.FilterType, l string) *filterLink {
	c := &filterLink{Type: f, Label: l}

	storeutil.FilterDispacher.Subscribe(func() {
		vecty.Rerender(c)
	})

	return c
}

func (f *filterLink) Render() vecty.ComponentOrHTML {
	isActive := storeutil.FilterDispacher.StateOf(storeutil.FilterReducer).(model.FilterType) == f.Type

	return &link{
		Type:     f.Type,
		IsActive: isActive,
		Label:    f.Label,
		OnClick: func(event *vecty.Event) {
			storeutil.FilterDispacher.Dispatch(storeutil.FilterReducer.Chenge.With(f.Type))
		},
	}
}
