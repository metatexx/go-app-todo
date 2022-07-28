package frontend

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type ToDoList struct {
	app.Compo
	todos           []*todo
	updateAvailable bool
	inputTodo       string
	filterMode      int
}

type todo struct {
	Done bool
	Text string
}

func (c *ToDoList) OnInit() {
	c.todos = []*todo{
		{Text: "todo Item Number 1"},
		{Text: "Clean the kitchen"},
		{Text: "Create PR"},
		{Text: "Relax!!!"},
		{Text: "Done Item", Done: true},
	}
}

func (c *ToDoList) OnAppUpdate(ctx app.Context) {
	c.updateAvailable = ctx.AppUpdateAvailable() // Reports that an app update is available.
}

func (c *ToDoList) Render() app.UI {
	if app.IsServer {
		// this gets called on the server before the page is delivered
		return app.Div().Text("app is loading")
	}
	return app.Section().
		Class("todoapp").Body(
		func() app.UI {
			if c.updateAvailable {
				return app.H1().
					Style("text-align", "center").
					Text("Update available, please reload.")
			}
			return app.Div()
		}(),
		app.H1().
			Text("todos"),
		app.Div().
			Styles(map[string]string{"display": "grid", "grid-template-columns": "auto auto"}).
			Body(
				app.Div().Body(
					app.Button().
						Text("✓").
						OnClick(c.toggleAllDone),
				),
				app.Form().
					OnSubmit(c.onSubmit).Body(
					app.Div().Body(
						app.Input().
							Type("text").
							Value(c.inputTodo).
							Class("new-todo").
							OnInput(c.ValueTo(&c.inputTodo)),
					),
				),
			),
		app.Div().Body(
			app.Ul().Body(
				app.Range(c.todos).Slice(func(i int) app.UI {
					if (c.filterMode == 1 && c.todos[i].Done) ||
						(c.filterMode == 2 && !c.todos[i].Done) {
						return app.Span()
					}
					return app.Li().
						Body(
							app.Label().Body(
								app.Input().
									Type("checkbox").
									Checked(c.todos[i].Done).
									OnChange(func(ctx app.Context, e app.Event) {
										c.todos[i].Done = !c.todos[i].Done
									}),
								app.Span().Styles(map[string]string{
									"text-decoration": func() string {
										if c.todos[i].Done {
											return "line-through"
										}
										return "none"
									}(),
								}).Body(
									app.Text(c.todos[i].Text),
								),
							),
							app.Div().
								Class("delete").
								Text("✕").
								OnClick(func(ctx app.Context, e app.Event) {
									c.todos = append(c.todos[:i], c.todos[i+1:]...)
								}))
				}),
				app.Li().Body(
					app.Div().
						Class("grid").Body(
						app.Div().
							Text(fmt.Sprintf("%d items left", c.countUncompleted())),
						app.Div().
							Class("selection").Body(
							c.generateFilterButton("All", 0, c.filterMode == 0),
							c.generateFilterButton("Active", 1, c.filterMode == 1),
							c.generateFilterButton("Completed", 2, c.filterMode == 2),
						),
						func() app.UI {
							if !c.hasCompletedTodo() {
								return app.Span()
							}
							return app.Div().
								Styles(map[string]string{
									"text-align": "right",
								}).Body(
								app.Button().
									Text("Clear completed").
									OnClick(c.clearCompleted))
						}(),
					),
				),
			),
		),
	)
}

func (c *ToDoList) onSubmit(_ app.Context, e app.Event) {
	e.PreventDefault()
	if c.inputTodo == "" {
		return
	}
	c.todos = append(c.todos, &todo{Text: c.inputTodo})
	c.inputTodo = ""
}

func (c *ToDoList) countUncompleted() int {
	sum := 0
	for _, v := range c.todos {
		if !v.Done {
			sum++
		}
	}
	return sum
}

func (c *ToDoList) hasCompletedTodo() bool {
	for _, v := range c.todos {
		if v.Done {
			return true
		}
	}
	return false
}

func (c *ToDoList) clearCompleted(_ app.Context, _ app.Event) {
	var tmp []*todo

	for _, v := range c.todos {
		if !v.Done {
			tmp = append(tmp, v)
		}
	}
	c.todos = tmp
}

func (c *ToDoList) switchSelection(mode int) func(ctx app.Context, e app.Event) {
	return func(ctx app.Context, e app.Event) {
		c.filterMode = mode
	}
}

func (c *ToDoList) toggleAllDone(_ app.Context, e app.Event) {
	e.PreventDefault()
	newDone := false
	if !c.hasCompletedTodo() {
		newDone = true
	}

	for _, v := range c.todos {
		v.Done = newDone
	}

}

func (c *ToDoList) generateFilterButton(text string, mode int, active bool) app.UI {
	if active {
		return app.Button().
			Text(text).
			OnClick(c.switchSelection(mode)).
			Class("active")
	}
	return app.Button().
		Text(text).
		OnClick(c.switchSelection(mode))
}
