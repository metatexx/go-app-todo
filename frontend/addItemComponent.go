package frontend

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type AddItemComponent struct {
	app.Compo
	toDo           ToDo
	show           bool
	addItemHandler func(ctx app.Context, e app.Event)
	isUpdate       bool
	addBtnText     string
}

func (a *AddItemComponent) Render() app.UI {
	if a.addBtnText == "" {
		a.addBtnText = "Add"
	}

	if !a.show {
		return app.Div().Body(
			app.Div().Body(
				app.H3().Text("Add new Item").Styles(map[string]string{
					"margin": "4px 0px",
					"float":  "left",
				}),
				app.Button().Style("float", "right").Text("Open").OnClick(a.openInputBox),
			),
			app.Br(),
		)
	}

	return app.Div().Body(
		app.Div().Body(
			app.H3().Text("Add new Item").Styles(map[string]string{
				"margin": "4px 0px",
				"float":  "left",
			}),
			app.Button().Style("float", "right").Text("Close").OnClick(a.openInputBox),
		),
		app.Br(),
		app.Div().Styles(map[string]string{
			"border":     "1px solid",
			"padding":    "4px",
			"margin-top": "10px",
		}).
			Body(
				app.Input().Type("text").
					Styles(map[string]string{
						"width":         "98%",
						"margin-bottom": "5px",
					}).
					Value(a.toDo.Text).
					Placeholder("What`s next?").
					AutoFocus(true).
					OnChange(a.ValueTo(&a.toDo.Text)),
				app.Label().Style("margin-bottom", "5px").Body(
					app.Input().Type("checkbox").Checked(a.toDo.Important).OnChange(func(ctx app.Context, e app.Event) {
						a.toDo.Important = !a.toDo.Important
					}),
					app.Text("Mark as important"),
				),
				app.Div().Style("margin-top", "5px").Body(
					app.Button().Styles(map[string]string{
						"background-color": "red",
						"padding":          "4px",
					}).Text("Cancel").OnClick(a.closeBoxAction),
					app.Button().Styles(map[string]string{
						"background-color": "orange",
						"padding":          "4px",
						"margin-left":      "5px",
					}).Text("Clear").OnClick(a.clearBox),
					app.Button().Styles(map[string]string{
						"background-color": "green",
						"padding":          "4px",
						"margin-left":      "5px",
					}).Text(a.addBtnText).OnClick(a.addItemHandler),
				),
			),
	)
}

func (a *AddItemComponent) closeBoxAction(ctx app.Context, e app.Event) {
	a.reset()
	a.show = false
}

func (a *AddItemComponent) openInputBox(ctx app.Context, e app.Event) {
	if !a.show {
		a.show = true
	} else {
		a.reset()
	}
}

func (a *AddItemComponent) clearBox(ctx app.Context, e app.Event) {
	a.toDo = ToDo{}
	a.isUpdate = false
	a.addBtnText = "Add"
}

func (a *AddItemComponent) reset() {
	a.toDo = ToDo{}

	a.show = false
	a.isUpdate = false
	a.addBtnText = "Add"
}

func (a *AddItemComponent) editItem(item ToDo) {
	a.show = true
	a.isUpdate = true
	a.addBtnText = "Update"

	a.toDo = item

	a.Update()
}
