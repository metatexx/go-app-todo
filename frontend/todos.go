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
		{
			Text: "todo Item Number 1",
		},
		{
			Text: "Clean the kitchen",
		},
		{
			Text: "Create PR",
		},
		{
			Text: "Relax!!!",
		},
		{
			Done: true,
			Text: "Done Item",
		},
	}
}

func (c *ToDoList) OnAppUpdate(ctx app.Context) {
	c.updateAvailable = ctx.AppUpdateAvailable() // Reports that an app update is available.
}

func (c *ToDoList) Render() app.UI {
	return app.Section().Class("todoapp").Body(
		func() app.UI {
			if c.updateAvailable {
				return app.H1().Style("text-align", "center").Text("Update available, please reload.")
			}
			return app.Div()
		}(),
		app.H1().Text("todos"),
		app.Form().Body(
			app.Div().Styles(map[string]string{"display": "grid", "grid-template-columns": "auto auto"}).Body(
				app.Div().Body(
					app.Button().Text("✓"),
				).OnClick(c.handleToggleDone),
				app.Div().Body(
					app.Input().Class("new-todo").Type("text").
						Value(c.inputTodo).
						OnInput(c.ValueTo(&c.inputTodo)),
				),
			),
		).OnSubmit(c.onSubmit),
		app.Div().Body(
			app.Ul().Body(
				app.Range(c.todos).Slice(func(i int) app.UI {
					if (c.filterMode == 1 && c.todos[i].Done) || (c.filterMode == 2 && !c.todos[i].Done) {
						return app.Span()
					}
					return app.Li().Styles(map[string]string{
						"text-decoration": func() string {
							if c.todos[i].Done {
								return "line-through"
							}
							return "none"
						}(),
					}).Body(
						app.Label().Body(
							app.Input().Type("checkbox").
								Checked(c.todos[i].Done).
								OnChange(c.toggleDone(c.todos[i])),
							app.Text(c.todos[i].Text),
						),
						app.Div().Class("destroy").Text("✕").OnClick(c.deleteHandler(c.todos[i])))
				}),
				app.Li().Body(
					app.Div().Class("grid").Body(
						app.Div().Text(c.generateLeftItemsOutput()),
						app.Div().Class("selection").Body(
							func() app.UI {
								if c.filterMode == 0 {
									return app.Button().Text("All").OnClick(c.switchSelection(0)).Class("active")
								}
								return app.Button().Text("All").OnClick(c.switchSelection(0))
							}(),
							func() app.UI {
								if c.filterMode == 1 {
									return app.Button().Text("Active").OnClick(c.switchSelection(1)).Class("active")
								}
								return app.Button().Text("Active").OnClick(c.switchSelection(1))
							}(),
							func() app.UI {
								if c.filterMode == 2 {
									return app.Button().Text("Completed").OnClick(c.switchSelection(2)).Class("active")
								}
								return app.Button().Text("Completed").OnClick(c.switchSelection(2))
							}(),
						),
						func() app.UI {
							if c.hasCompletedTodo() {
								return app.Div().Styles(map[string]string{
									"text-align": "right",
								}).Body(app.Button().Text("Clear completed")).OnClick(c.handleClearCompleted)
							}
							return app.Div()
						}(),
					),
				),
			),
		),
	)
}

func (c *ToDoList) deleteHandler(item *todo) func(ctx app.Context, e app.Event) {
	return func(ctx app.Context, e app.Event) {
		for i, v := range c.todos {
			if v == item {
				c.todos = append(c.todos[:i], c.todos[i+1:]...)
				return
			}
		}
	}
}

func (c *ToDoList) toggleDone(item *todo) func(ctx app.Context, e app.Event) {
	return func(ctx app.Context, e app.Event) {
		item.Done = !item.Done
	}
}

func (c *ToDoList) onSubmit(ctx app.Context, e app.Event) {
	e.PreventDefault()

	c.todos = append(c.todos, &todo{Text: c.inputTodo})
	c.inputTodo = ""
}

func (c *ToDoList) generateLeftItemsOutput() string {
	sum := 0
	for _, v := range c.todos {
		if !v.Done {
			sum++
		}
	}
	return fmt.Sprintf("%v items left", sum)
}

func (c *ToDoList) hasCompletedTodo() bool {
	for _, v := range c.todos {
		if v.Done {
			return true
		}
	}
	return false
}

func (c *ToDoList) hasUncompletedTodo() bool {
	for _, v := range c.todos {
		if !v.Done {
			return true
		}
	}
	return false
}

func (c *ToDoList) handleClearCompleted(ctx app.Context, e app.Event) {
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

func (c *ToDoList) handleToggleDone(ctx app.Context, e app.Event) {
	e.PreventDefault()
	newDone := false
	if c.hasUncompletedTodo() {
		newDone = true
	}

	for _, v := range c.todos {
		v.Done = newDone
	}

}
