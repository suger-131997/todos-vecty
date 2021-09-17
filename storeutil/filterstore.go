package storeutil

import (
	"todos-vecty/model"

	"github.com/dannypsnl/redux/v2/rematch"
	"github.com/dannypsnl/redux/v2/store"
)

var FilterDispacher *store.Store

var FilterReducer *filterStore

func FilterStoreInit() {
	FilterReducer = &filterStore{State: model.All}
	FilterDispacher = store.New(FilterReducer)
}

type filterStore struct {
	rematch.Reducer
	State model.FilterType

	Chenge *rematch.Action `action:"ChengeFilter"`
}

func (t *filterStore) ChengeFilter(s, a model.FilterType) model.FilterType {
	return a
}
