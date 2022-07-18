package frontend

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type ToDoList struct {
	app.Compo
	toDos           []ToDo
	addItemBox      *AddItemComponent
	updateAvailable bool
}

type ToDo struct {
	ID        uuid.UUID
	Done      bool
	Important bool
	Text      string
}

func (t *ToDoList) OnInit() {
	t.toDos = []ToDo{
		{
			ID:   uuid.New(),
			Text: "ToDo Item Number 1",
		},
		{
			ID:        uuid.New(),
			Important: true,
			Text:      "Clean the kitchen",
		},
		{
			ID:   uuid.New(),
			Text: "Create PR",
		},
		{
			ID:   uuid.New(),
			Text: "Relax!!!",
		},
		{
			ID:   uuid.New(),
			Done: true,
			Text: "Done Item",
		},
	}

	t.addItemBox = &AddItemComponent{
		addItemHandler: t.addItemHandler,
	}

}

func (t *ToDoList) OnAppUpdate(ctx app.Context) {
	t.updateAvailable = ctx.AppUpdateAvailable() // Reports that an app update is available.
}

func (t *ToDoList) Render() app.UI {
	return app.Div().Styles(map[string]string{"width": "300px", "margin": "0 auto"}).Body(
		func() app.UI {
			if t.updateAvailable {
				return app.H1().Style("text-align", "center").Text("Update available, please reload.")
			}
			return app.Div()
		}(),
		app.H1().Text("ToDo List"),

		t.addItemBox,
		app.Div().Styles(map[string]string{}).Body(
			app.Ul().Style("padding", "0px").Body(
				app.Range(t.toDos).Slice(func(i int) app.UI {
					return &ToDoItem{
						item: t.toDos[i],
						editHandler: func(ctx app.Context, e app.Event) {
							t.addItemBox.editItem(t.toDos[i])
						},
						deleteHandler: func(ctx app.Context, e app.Event) {
							t.toDos = append(t.toDos[:i], t.toDos[i+1:]...)
						},
					}
				}),
			),
		),
	)
}

func (t *ToDoList) addItemHandler(ctx app.Context, e app.Event) {
	if t.addItemBox.isUpdate {
		for i := range t.toDos {
			if t.toDos[i].ID == t.addItemBox.toDo.ID {
				t.toDos[i].Important = t.addItemBox.toDo.Important
				t.toDos[i].Text = t.addItemBox.toDo.Text
				break
			}
		}
	} else {
		t.addItemBox.toDo.ID = uuid.New()
		t.toDos = append(
			[]ToDo{t.addItemBox.toDo},
			t.toDos...)
	}

	fmt.Println(t.toDos)

	t.addItemBox.reset()
}
