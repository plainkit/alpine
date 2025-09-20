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
				Meta(Charset("utf-8")),
				Meta(Name("viewport"), Content("width=device-width, initial-scale=1.0")),
				HeadTitle(Text("Alpine.js with Plain Demo")),
				Script(ScriptSrc("/js/alpine.min.js"), Defer()),
				HeadStyle(Text(`
					body {
						font-family: system-ui, -apple-system, sans-serif;
						max-width: 800px;
						margin: 2rem auto;
						padding: 0 1rem;
						line-height: 1.6;
						color: #1f2937;
						background: #fafafa;
					}

					h1 {
						color: #111827;
						border-bottom: 2px solid #d1d5db;
						padding-bottom: 0.5rem;
					}

					h2 {
						color: #374151;
						padding: 0;
						margin: 0;
					}

					.container {
						margin: 1rem;
						padding: 1rem;
						background: #f3f4f6;
						border-radius: 12px;
						border: 1px solid #e5e7eb;
					}

					button {
						padding: 0.5rem 1rem;
						margin: 0.5rem;
						cursor: pointer;
						background: #4b5563;
						color: white;
						border: none;
						border-radius: 6px;
						transition: all 0.2s;
						font-weight: 500;
					}

					button:hover {
						background: #374151;
						transform: translateY(-1px);
						box-shadow: 0 2px 4px rgba(0,0,0,0.1);
					}

					button:active {
						transform: translateY(0);
					}

					button.active {
						background: #1f2937;
						font-weight: bold;
						box-shadow: inset 0 2px 4px rgba(0,0,0,0.2);
					}

					.hidden { display: none; }

					.card {
						border: 1px solid #d1d5db;
						padding: 1.25rem;
						margin: 1rem 0;
						border-radius: 10px;
						background: white;
						box-shadow: 0 1px 3px rgba(0,0,0,0.08);
					}

					/* Todo form styles */
					form {
						display: flex;
						gap: 0.5rem;
						margin-bottom: 1rem;
					}

					input[type="text"] {
						flex: 1;
						padding: 0.75rem;
						border: 2px solid #e5e7eb;
						border-radius: 6px;
						font-size: 1rem;
						background: white;
						transition: all 0.2s;
					}

					input[type="text"]:focus {
						outline: none;
						border-color: #9ca3af;
						background: #fafafa;
						box-shadow: 0 0 0 3px rgba(107, 114, 128, 0.1);
					}

					input[type="text"]::placeholder {
						color: #9ca3af;
					}

					form button {
						background: #6b7280;
						padding: 0.75rem 1.5rem;
						margin: 0;
						font-weight: 600;
					}

					form button:hover {
						background: #4b5563;
					}

					input[type="checkbox"] {
						margin-right: 0.75rem;
						width: 18px;
						height: 18px;
						cursor: pointer;
						accent-color: #6b7280;
					}

					.todo-item {
						display: flex;
						align-items: center;
						padding: 1rem;
						margin: 0.5rem 0;
						background: white;
						border-radius: 8px;
						border: 1px solid #e5e7eb;
						transition: all 0.3s;
					}

					.todo-item:hover {
						box-shadow: 0 3px 6px rgba(0,0,0,0.08);
						border-color: #d1d5db;
					}

					.todo-item.done {
						background: #f9fafb;
						opacity: 0.6;
					}

					.todo-item span {
						flex: 1;
						margin: 0 1rem;
						font-size: 1.05rem;
						color: #374151;
					}

					.todo-item.done span {
						text-decoration: line-through;
						color: #9ca3af;
					}

					.todo-item button {
						background: #6b7280;
						padding: 0.4rem 0.9rem;
						font-size: 0.875rem;
						margin: 0;
					}

					.todo-item button:hover {
						background: #4b5563;
					}

					/* Counter styles */
					.counter-display {
						display: inline-block;
						margin: 0 1.5rem;
						font-weight: bold;
						color: #1f2937;
						min-width: 4rem;
						text-align: center;
						background: #f3f4f6;
						border-radius: 8px;
						padding: 0.5rem 1rem;
						border: 2px solid #e5e7eb;
					}

					/* Dropdown styles */
					.dropdown-container {
						position: relative;
						display: inline-block;
					}

					.dropdown-menu {
						position: absolute;
						top: 100%;
						left: 0;
						margin-top: 0.5rem;
						background: white;
						border: 1px solid #d1d5db;
						border-radius: 10px;
						box-shadow: 0 4px 12px rgba(0,0,0,0.08);
						min-width: 160px;
						z-index: 100;
					}

					.dropdown-menu ul {
						list-style: none;
						padding: 0.5rem 0;
						margin: 0;
					}

					.dropdown-menu li {
						padding: 0;
						margin: 0;
					}

					.dropdown-menu a {
						display: block;
						padding: 0.75rem 1.25rem;
						color: #374151;
						text-decoration: none;
						transition: background 0.15s;
					}

					.dropdown-menu a:hover {
						background: #f3f4f6;
						color: #1f2937;
					}

					.dropdown-menu li:first-child a {
						border-radius: 10px 10px 0 0;
					}

					.dropdown-menu li:last-child a {
						border-radius: 0 0 10px 10px;
					}

					/* Tab styles */
					.tab-buttons {
						display: flex;
						gap: 0.5rem;
						margin-bottom: 1.25rem;
						background: #e5e7eb;
						padding: 0.25rem;
						border-radius: 8px;
					}

					.tab-buttons button {
						background: transparent;
						color: #6b7280;
						margin: 0;
						flex: 1;
					}

					.tab-buttons button:hover {
						background: rgba(255,255,255,0.5);
						color: #374151;
					}

					.tab-buttons button.active {
						background: white;
						color: #1f2937;
						box-shadow: 0 1px 3px rgba(0,0,0,0.1);
					}

					[x-cloak] { display: none !important; }
				`)),
			),
			Body(
				H1(Text("Alpine.js with Plain Demo")),

				// Simple Toggle Example
				Div(Class("container"),
					H2(Text("Simple Toggle")),
					Div(
						alpine.XData("{ open: false }"),
						Button(
							alpine.AtClick("open = !open"),
							Text("Toggle Content"),
						),
						Div(
							alpine.XShow("open"),
							Class("card"),
							P(Text("This content is toggled with Alpine.js!")),
						),
					),
				),

				// Counter Example
				Div(Class("container"),
					H2(Text("Counter")),
					Div(
						alpine.XData("{ count: 0 }"),
						Button(
							alpine.AtClick("count++"),
							Text("+"),
						),
						Span(
							Class("counter-display"),
							alpine.XText("count"),
						),
						Button(
							alpine.AtClick("count--"),
							Text("-"),
						),
					),
				),

				// Todo List Example
				Div(Class("container"),
					H2(Text("Todo List")),
					Div(
						alpine.XData(`{
							todos: [
								{ id: 1, text: 'Learn Alpine.js', done: false },
								{ id: 2, text: 'Build with Plain', done: false }
							],
							newTodo: '',
							addTodo() {
								if (this.newTodo.trim()) {
									this.todos.push({
										id: Date.now(),
										text: this.newTodo,
										done: false
									});
									this.newTodo = '';
								}
							},
							removeTodo(id) {
								this.todos = this.todos.filter(t => t.id !== id);
							}
						}`),

						Form(
							alpine.AtSubmitPrevent("addTodo()"),
							Input(
								InputType("text"),
								alpine.XModel("newTodo"),
								Placeholder("Add a todo..."),
							),
							Button(
								ButtonType("submit"),
								Text("Add Todo"),
							),
						),

						Div(
							Template(
								alpine.XFor("todo in todos"),
								alpine.Colon("key", "todo.id"),
								Div(
									Class("todo-item"),
									alpine.ColonClass("{ 'done': todo.done }"),
									Input(
										InputType("checkbox"),
										alpine.XModel("todo.done"),
									),
									Span(
										alpine.XText("todo.text"),
									),
									Button(
										alpine.AtClick("removeTodo(todo.id)"),
										Text("Delete"),
									),
								),
							),
						),
					),
				),

				// Tabs Example
				Div(Class("container"),
					H2(Text("Tabs")),
					Div(
						alpine.XData("{ activeTab: 'tab1' }"),
						Div(Class("tab-buttons"),
							Button(
								alpine.AtClick("activeTab = 'tab1'"),
								alpine.ColonClass("{ 'active': activeTab === 'tab1' }"),
								Text("Tab 1"),
							),
							Button(
								alpine.AtClick("activeTab = 'tab2'"),
								alpine.ColonClass("{ 'active': activeTab === 'tab2' }"),
								Text("Tab 2"),
							),
							Button(
								alpine.AtClick("activeTab = 'tab3'"),
								alpine.ColonClass("{ 'active': activeTab === 'tab3' }"),
								Text("Tab 3"),
							),
						),
						Div(Class("card"),
							Div(
								alpine.XShow("activeTab === 'tab1'"),
								P(Text("Content for Tab 1")),
							),
							Div(
								alpine.XShow("activeTab === 'tab2'"),
								P(Text("Content for Tab 2")),
							),
							Div(
								alpine.XShow("activeTab === 'tab3'"),
								P(Text("Content for Tab 3")),
							),
						),
					),
				),

				// Dropdown Example
				Div(Class("container"),
					H2(Text("Dropdown")),
					Div(Class("dropdown-container"),
						alpine.XData("{ open: false }"),
						Button(
							alpine.AtClick("open = !open"),
							Text("Options ▼"),
						),
						Div(
							alpine.XShow("open"),
							alpine.AtClickAway("open = false"),
							Class("dropdown-menu"),
							Ul(
								Li(A(Href("#"), Text("Option 1"))),
								Li(A(Href("#"), Text("Option 2"))),
								Li(A(Href("#"), Text("Option 3"))),
							),
						),
					),
				),
			),
		)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = fmt.Fprint(w, "<!DOCTYPE html>\n")
		_, _ = fmt.Fprint(w, Render(page))
	})

	fmt.Println("Server starting on http://localhost:8080")
	_ = http.ListenAndServe(":8080", nil)
}
