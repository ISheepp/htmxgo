package render

import (
	"context"
	"fmt"
	. "github.com/theplant/htmlgo"
	"os"
)

const (
	HxGet     = "hx-get"
	HxPost    = "hx-post"
	HxSwap    = "hx-swap"
	HxTarget  = "hx-target"
	HxTrigger = "hx-trigger"
	HxParams  = "hx-params"
)

func Generate() {
	file, err := os.Create("./render/index.html")
	if err != nil {
		panic(err)
	}
	Fprint(file, render(), nil)
}

func render() HTMLComponent {
	return HTML(
		Head(
			Meta().Charset("utf-8"),
			Meta().Name("viewport").Content("width=device-width, initial-scale=1.0"),
			Title("Movies"),
		),
		Link("https://cdn.jsdelivr.net/npm/daisyui@4.6.0/dist/full.min.css").Rel("stylesheet").Type("text/css"),
		RawHTML("<script src=\"https://cdn.tailwindcss.com\"></script>\n"),
		RawHTML("<script src=\"https://unpkg.com/htmx.org@latest\" crossorigin=\"anonymous\"></script>\n"),
		Body(
			Div(
				Div(
					Div(
						H1("Best Movies").Class("text-2xl font-bold"),
						Button("Add").Attr("onclick", "add_model.showModal()").Id("addDataBtn").Class("btn btn-success btn-sm rounded-md"),
					).Class("flex justify-between items-center mb-4"),

					Div().Attr(HxGet, "http://127.0.0.1:8080/movies", HxSwap, "innerHTML", HxTarget, "#movieTable", HxTrigger, "load").Id("movieTable"),
				).Class("bg-white p-8 rounded-md shadow-md"),
			).Class("bg-gray-200 h-screen flex items-center justify-center"),

			movieDialog("Add", "add_model"),
			movieDialog("Update", "update_model"),
		),
	)
}

func movieDialog(action, id string) HTMLComponent {
	return ComponentFunc(func(ctx context.Context) (r []byte, err error) {
		title := fmt.Sprintf("%s Movie!", action)
		return Dialog(
			Div(
				H3(title).Class("font-bold text-lg"),
				Div(
					Input("title").Type("text").Placeholder("Title").Class("input input-bordered w-full max-w-xs"),
					Input("director").Type("text").Placeholder("Director").Class("input input-bordered w-full max-w-xs"),
				).Class("mt-4"),
				Div(
					Button(action).Class("btn btn-success"),
					Form(
						Button("Close").Class("btn"),
					).Method("dialog"),
				).Class("modal-action"),
			).Class("modal-box"),
		).Id(id).Class("modal").MarshalHTML(ctx)
	})
}
