package api

import (
	"context"
	. "github.com/theplant/htmlgo"
	"strconv"
)

func MovieTableHead() HTMLComponent {
	return ComponentFunc(func(ctx context.Context) (r []byte, err error) {
		return Thead(
			Tr(
				Th("ID").Class("py-2"),
				Th("Title").Class("py-2"),
				Th("Director").Class("py-2"),
				Th("Operation").Class("py-2"),
			),
		).MarshalHTML(ctx)
	})
}

func MovieTableBody(movie Movie) HTMLComponent {
	return ComponentFunc(func(ctx context.Context) (r []byte, err error) {
		return Tr(
			Td().Text(strconv.Itoa(movie.Id)).Class("py-2"),
			Td().Text(movie.Title).Class("py-2"),
			Td().Text(movie.Director).Class("py-2"),
			Td(
				Button("edit").Attr("onclick", "update_model.showModal()").Class("btn btn-xs btn-neutral mr-2"),
				Button("delete").Class("btn btn-xs btn-error"),
			).Class("py-2"),
		).MarshalHTML(ctx)
	})
}
