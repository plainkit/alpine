# Alpine.js for Plain

Alpine.js attribute functions for the Plain component library.

## Installation

```bash
go get github.com/plainkit/alpine
```

## Usage

### Basic Example

```go
package main

import (
    "github.com/plainkit/html"
    "github.com/plainkit/alpine"
)

func main() {
    // Create a simple Alpine.js component
    component := html.Div(
        alpine.XData("{ open: false }"),

        html.Button(
            alpine.AtClick("open = !open"),
            html.Text("Toggle"),
        ),

        html.Div(
            alpine.XShow("open"),
            html.Text("Hello Alpine.js!"),
        ),
    )
}
```

### Core Directives

#### Data & Initialization
- `XData(data)` - Declares a new Alpine component and its data
- `XInit(code)` - Runs code when a component initializes

#### Display & Rendering
- `XShow(expression)` - Shows/hides an element based on a condition
- `XIf(expression)` - Conditionally renders elements (use with `<template>`)
- `XFor(expression)` - Creates DOM elements by iterating (use with `<template>`)
- `XHtml(expression)` - Sets the inner HTML of an element
- `XText(expression)` - Sets the text content of an element
- `XCloak()` - Hides elements until Alpine is initialized

#### Event Handling
- `XOn(event, handler)` - Listens for browser events
- `AtClick(handler)` - Shorthand for @click
- `AtSubmit(handler)` - Shorthand for @submit
- `AtChange(handler)` - Shorthand for @change
- `AtInput(handler)` - Shorthand for @input

#### Binding
- `XBind(attribute, expression)` - Dynamically sets HTML attributes
- `ColonClass(expression)` - Shorthand for :class
- `ColonStyle(expression)` - Shorthand for :style
- `XModel(expression)` - Two-way data binding for form inputs

#### Advanced
- `XTransition()` - Adds transitions to elements
- `XEffect(expression)` - Re-evaluates when dependencies change
- `XRef(name)` - References DOM elements
- `XTeleport(selector)` - Moves elements to another location
- `XIgnore()` - Tells Alpine to ignore a block of HTML
- `XId(expression)` - Generates unique IDs

### Event Modifiers

```go
// Click outside/away
alpine.AtClickAway("open = false")

// Prevent default
alpine.AtClickPrevent("handleClick()")

// Stop propagation
alpine.AtClickStop("handleClick()")

// Keyboard events
alpine.AtKeydownEscape("open = false")
alpine.AtKeydownEnter("submit()")
```

### Model Modifiers

```go
// Lazy update (on change instead of input)
alpine.XModelLazy("value")

// Number conversion
alpine.XModelNumber("age")

// Debounce
alpine.XModelDebounce("searchQuery", "500ms")
```

### Transitions

```go
html.Div(
    alpine.XShow("open"),
    alpine.XTransitionEnter("transition ease-out duration-300"),
    alpine.XTransitionEnterStart("transform opacity-0 scale-90"),
    alpine.XTransitionEnterEnd("transform opacity-100 scale-100"),
    alpine.XTransitionLeave("transition ease-in duration-300"),
    alpine.XTransitionLeaveStart("transform opacity-100 scale-100"),
    alpine.XTransitionLeaveEnd("transform opacity-0 scale-90"),
    html.Text("Animated content"),
)
```

### Serving Alpine.js

The package includes the Alpine.js library which can be served directly:

```go
http.HandleFunc("/js/alpine.min.js", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/javascript")
    w.Header().Set("Cache-Control", "public, max-age=31536000")
    w.Write(alpine.JavaScript())
})
```

### Complete Example

```go
package main

import (
    "net/http"
    "github.com/plainkit/html"
    "github.com/plainkit/alpine"
)

func TodoList() html.Node {
    return html.Div(
        alpine.XData(`{
            todos: [],
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

        html.Form(
            alpine.AtSubmitPrevent("addTodo()"),
            html.Input(
                html.Type("text"),
                alpine.XModel("newTodo"),
                html.Placeholder("Add a todo..."),
            ),
            html.Button(
                html.Type("submit"),
                html.Text("Add"),
            ),
        ),

        html.Template(
            alpine.XFor("todo in todos"),
            html.Li(
                html.Label(
                    html.Input(
                        html.Type("checkbox"),
                        alpine.XModel("todo.done"),
                    ),
                    html.Span(
                        alpine.XText("todo.text"),
                        alpine.ColonClass("{ 'line-through': todo.done }"),
                    ),
                ),
                html.Button(
                    alpine.AtClick("removeTodo(todo.id)"),
                    html.Text("Delete"),
                ),
            ),
        ),
    )
}
```

## License

MIT