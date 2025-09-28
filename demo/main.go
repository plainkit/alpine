package main

import (
	"fmt"
	"net/http"

	"github.com/plainkit/alpine"
	. "github.com/plainkit/html"
)

func main() {
	// Serve Alpine.js library
	http.HandleFunc("/js/alpine.min.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/javascript")
		w.Header().Set("Cache-Control", "public, max-age=31536000")
		_, _ = w.Write(alpine.JavaScript())
	})

	// Main page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		page := Html(
			Head(
				Meta(ACharset("utf-8")),
				Meta(AName("viewport"), AContent("width=device-width, initial-scale=1.0")),
				Title(Text("Todo App")),
				Script(ASrc("/js/alpine.min.js"), ADefer()),
				Style(Text(`
					@import url('https://fonts.googleapis.com/css2?family=Inter:ital,wght@0,100;0,200;0,300;0,400;0,500;0,600;0,700;0,800;0,900;1,100;1,200;1,300;1,400;1,500;1,600;1,700;1,800;1,900&display=swap');

					* {
						margin: 0;
						padding: 0;
						box-sizing: border-box;
						font-family: Inter, sans-serif;
					}

					:root {
						--background: #ffffff;
						--foreground: #0a0a0a;
						--card: #ffffff;
						--primary: #0a0a0a;
						--primary-foreground: #fafafa;
						--secondary: #f5f5f5;
						--secondary-foreground: #0a0a0a;
						--muted-foreground: #737373;
						--border: #e5e5e5;
						--radius: 0.5rem;
					}

					body {
						background: var(--background);
						color: var(--foreground);
						line-height: 1.5;
						min-height: 100vh;
						display: flex;
						align-items: center;
						justify-content: center;
						padding: 1rem;
					}

					.container {
						max-width: 500px;
						width: 100%;
						background: var(--card);
						border: 1px solid var(--border);
						border-radius: var(--radius);
						padding: 2rem;
						box-shadow: 0 1px 3px 0 rgb(0 0 0 / 0.1), 0 1px 2px -1px rgb(0 0 0 / 0.1);
					}

					h1 {
						font-family: Inter, sans-serif;
						font-size: 2rem;
						font-weight: 800;
						margin-bottom: 0.5rem;
						text-align: center;
						letter-spacing: -0.025em;
					}

					.subtitle {
						color: var(--muted-foreground);
						text-align: center;
						margin-bottom: 2rem;
						font-size: 0.875rem;
					}

					.add-todo {
						display: flex;
						gap: 0.5rem;
						margin-bottom: 1.5rem;
					}

					input {
						flex: 1;
						padding: 0.5rem 0.75rem;
						border: 1px solid var(--border);
						border-radius: calc(var(--radius) - 2px);
						font-size: 0.875rem;
						outline: none;
						transition: border-color 0.2s;
					}

					input:focus {
						border-color: var(--primary);
						box-shadow: 0 0 0 2px rgb(0 0 0 / 0.1);
					}

					button {
						padding: 0.5rem 1rem;
						background: var(--primary);
						color: var(--primary-foreground);
						border: none;
						border-radius: calc(var(--radius) - 2px);
						font-size: 0.875rem;
						font-weight: 600;
						cursor: pointer;
						transition: all 0.2s ease;
						white-space: nowrap;
						box-shadow: 0 1px 2px 0 rgb(0 0 0 / 0.05);
					}

					button:hover {
						background: rgb(0 0 0 / 0.8);
						box-shadow: 0 1px 3px 0 rgb(0 0 0 / 0.1), 0 1px 2px -1px rgb(0 0 0 / 0.1);
						transform: translateY(-1px);
					}

					button.secondary {
						background: var(--secondary);
						color: var(--secondary-foreground);
						border: 1px solid var(--border);
						padding: 0.25rem 0.5rem;
						font-size: 0.75rem;
					}

					button.secondary:hover {
						background: rgb(0 0 0 / 0.05);
					}

					.todo-list {
						list-style: none;
					}

					.todo-item {
						display: flex;
						align-items: center;
						gap: 0.75rem;
						padding: 0.75rem 0;
						border-bottom: 1px solid var(--border);
					}

					.todo-item:last-child {
						border-bottom: none;
					}

					.todo-checkbox {
						width: 1.25rem;
						height: 1.25rem;
						min-width: 1.25rem;
						max-width: 1.25rem;
						border: 2px solid var(--border);
						border-radius: 0.375rem;
						cursor: pointer;
						appearance: none;
						position: relative;
						background: var(--background);
						transition: all 0.2s ease;
						flex-shrink: 0;
						flex-grow: 0;
						display: inline-block;
					}

					.todo-checkbox:hover {
						border-color: var(--primary);
						box-shadow: 0 0 0 2px rgb(0 0 0 / 0.1);
					}

					.todo-checkbox:checked {
						background: var(--primary);
						border-color: var(--primary);
						box-shadow: 0 0 0 2px rgb(0 0 0 / 0.1);
					}

					.todo-checkbox:checked::after {
						content: '✓';
						position: absolute;
						top: 50%;
						left: 50%;
						transform: translate(-50%, -50%);
						color: white;
						font-size: 0.875rem;
						font-weight: 700;
						line-height: 1;
					}

					.todo-checkbox:focus {
						outline: 2px solid var(--primary);
						outline-offset: 2px;
					}

					.todo-text {
						flex: 1;
						font-size: 0.875rem;
					}

					.todo-text.completed {
						text-decoration: line-through;
						color: var(--muted-foreground);
					}

					.empty-state {
						text-align: center;
						padding: 3rem 1rem;
						color: var(--muted-foreground);
						font-size: 0.875rem;
					}

					.stats {
						margin-top: 1.5rem;
						padding-top: 1rem;
						border-top: 1px solid var(--border);
						display: flex;
						justify-content: space-between;
						align-items: center;
						font-size: 0.75rem;
						color: var(--muted-foreground);
					}
				`)),
			),
			Body(
				Div(AClass("container"),
					alpine.XData(`{
							newTodo: '',
							todos: [
								{id: 1, text: 'Learn Alpine.js', completed: false},
								{id: 2, text: 'Build awesome UI', completed: false},
								{id: 3, text: 'Deploy to production', completed: true}
							],
							nextId: 4,
							addTodo() {
								if (this.newTodo.trim()) {
									this.todos.push({
										id: this.nextId++,
										text: this.newTodo.trim(),
										completed: false
									});
									this.newTodo = '';
								}
							},
							deleteTodo(id) {
								this.todos = this.todos.filter(todo => todo.id !== id);
							},
							toggleTodo(id) {
								const todo = this.todos.find(t => t.id === id);
								if (todo) todo.completed = !todo.completed;
							},
							get completedCount() {
								return this.todos.filter(t => t.completed).length;
							},
							get totalCount() {
								return this.todos.length;
							}
						}`),

					H1(Text("Todo App")),
					P(AClass("subtitle"), Text("Stay organized and productive")),

					Div(
						AClass("add-todo"),
						Input(
							AType("text"),
							APlaceholder("Add a new todo..."),
							alpine.XModel("newTodo"),
							alpine.AtKeydownEnter("addTodo()"),
						),
						Button(
							alpine.XOnClick("addTodo()"),
							Text("Add"),
						),
					),

					Div(
						alpine.XShow("todos.length === 0"),
						AClass("empty-state"),
						Text("No todos yet. Add one above to get started!"),
					),

					Ul(
						AClass("todo-list"),
						alpine.XShow("todos.length > 0"),
						Template(
							alpine.XFor("todo in todos"),
							Li(
								AClass("todo-item"),
								Input(
									AType("checkbox"),
									AClass("todo-checkbox"),
									alpine.XModel("todo.completed"),
								),
								Span(
									AClass("todo-text"),
									alpine.XBindClass("{'completed': todo.completed}"),
									alpine.XText("todo.text"),
								),
								Button(
									AClass("secondary"),
									alpine.XOnClick("deleteTodo(todo.id)"),
									Text("Delete"),
								),
							),
						),
					),

					Div(
						AClass("stats"),
						alpine.XShow("todos.length > 0"),
						Span(alpine.XText("`${completedCount} of ${totalCount} completed`")),
						Button(
							AClass("secondary"),
							alpine.XOnClick("todos = todos.filter(t => !t.completed)"),
							alpine.XShow("completedCount > 0"),
							Text("Clear completed"),
						),
					),
				),
			),
		)

		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(Render(page)))
	})

	fmt.Println("🚀 Alpine + Plain Demo Server starting on :8080")
	fmt.Println("📦 Using embedded Alpine JavaScript (no CDN required)")
	fmt.Println("🔗 Open http://localhost:8080 to view the demo")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
