package api

import (
	"context"
	. "github.com/theplant/htmlgo"
	"htmxgo/render"
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
				Button("edit").Attr("onclick", "update_model.showModal()", render.HxGet, "http://127.0.0.1:8080/movie?id="+strconv.Itoa(movie.Id), render.HxSwap, "outerHTML", render.HxTarget, "#UpdateDialog").Class("btn btn-xs btn-neutral mr-2"),
				Button("delete").Attr(render.HxDelete, "http://127.0.0.1:8080/movie?id="+strconv.Itoa(movie.Id)).Class("btn btn-xs btn-error"),
			).Class("py-2"),
		).MarshalHTML(ctx)
	})
}

func MovieTable(prevUrl string, page int, afterUrl string) func(movies []Movie) HTMLComponent {
	table := func(movies []Movie) HTMLComponent {
		movieTrs := make([]HTMLComponent, 0)
		for _, movie := range movies {
			movieTrs = append(movieTrs, MovieTableBody(movie))
		}

		return ComponentFunc(func(ctx context.Context) (r []byte, err error) {
			table :=
				Table(
					MovieTableHead(),
					Tbody(movieTrs...),
				).Class("table min-w-full divide-y divide-gray-300")
			pagination :=
				Div(
					Button("«").Attr(render.HxGet, prevUrl, render.HxSwap, "innerHTML", render.HxTarget, "#movieTable").Class("join-item btn btn-sm"),
					Button(strconv.Itoa(page)).Class("join-item btn btn-sm"),
					Button("»").Attr(render.HxGet, afterUrl, render.HxSwap, "innerHTML", render.HxTarget, "#movieTable").Class("join-item btn btn-sm"),
				).Class("join flex justify-center mt-4")
			hs := HTMLComponents{
				table,
				pagination,
			}
			return hs.MarshalHTML(ctx)
		})
	}
	return table
}
