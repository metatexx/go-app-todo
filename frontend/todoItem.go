package frontend

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type ToDoItem struct {
	app.Compo
	item          ToDo
	editHandler   func(ctx app.Context, e app.Event)
	deleteHandler func(ctx app.Context, e app.Event)
}

func (t *ToDoItem) Render() app.UI {
	return app.Li().Styles(map[string]string{
		"clear":           "both",
		"list-style-type": "none",
		"text-decoration": func() string {
			if t.item.Done {
				return "line-through"
			}
			return "none"
		}(),
		"font-weight": func() string {
			if t.item.Important {
				return "bold"
			}
			return "normal"
		}(),
	}).Body(
		app.Label().Body(
			app.Input().Type("checkbox").
				Checked(t.item.Done).
				OnChange(func(ctx app.Context, e app.Event) {
					t.item.Done = !t.item.Done
				}),
			app.Text(t.item.Text),
		),
		app.Div().Style("float", "right").Body(
			func() app.UI {
				if t.item.Done {
					return app.Span()
				}

				return app.Span().Styles(map[string]string{
					"cursor":       "pointer",
					"margin-right": "10px",
				}).
					Text("✎").OnClick(t.editHandler)
			}(),
			app.Span().Styles(map[string]string{
				"color":        "red",
				"cursor":       "pointer",
				"margin-right": "10px",
			}).Text("✕").OnClick(t.deleteHandler),
			func() app.UI {
				if t.item.Important {
					return app.Span().Styles(map[string]string{
						"color":  "orange",
						"cursor": "pointer",
					}).
						Text("★").OnClick(func(ctx app.Context, e app.Event) {
						t.item.Important = false
					})
				}
				return app.Span().Styles(map[string]string{
					"cursor": "pointer",
				}).
					Text("☆").OnClick(func(ctx app.Context, e app.Event) {
					t.item.Important = true
				})
			}(),
		),
	)
}
