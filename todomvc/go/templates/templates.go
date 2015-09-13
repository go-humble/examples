package templates

// This package has been automatically generated with temple.
// Do not edit manually!

import (
	"github.com/go-humble/temple/temple"
)

var (
	GetTemplate     func(name string) (*temple.Template, error)
	GetPartial      func(name string) (*temple.Partial, error)
	GetLayout       func(name string) (*temple.Layout, error)
	MustGetTemplate func(name string) *temple.Template
	MustGetPartial  func(name string) *temple.Partial
	MustGetLayout   func(name string) *temple.Layout
)

func init() {
	var err error
	g := temple.NewGroup()

	if err = g.AddPartial("footer", `<!-- This should be 0 items left by default -->
<span class="todo-count">
	<strong>{{ len .Remaining }}</strong>
	item{{ if ne (len .Remaining) 1}}s{{end}} left
</span>
<!-- Remove this if you don't implement routing -->
<ul class="filters">
	<li>
		<a class="selected" href="#/">All</a>
	</li>
	<li>
		<a href="#/active">Active</a>
	</li>
	<li>
		<a href="#/completed">Completed</a>
	</li>
</ul>
<!-- Hidden if no completed items are left ↓ -->
{{ if len .Completed}}
<button class="clear-completed">Clear completed</button>
{{ end }}`); err != nil {
		panic(err)
	}

	if err = g.AddPartial("todo", `<li {{ if .Completed }}class="completed"{{ end }}>
	<div class="view">
		<input class="toggle" type="checkbox" {{ if .Completed }}checked{{ end }}>
		<label>{{ .Title }}</label>
		<button class="destroy"></button>
	</div>
	<input class="edit" value="{{ .Title }}">
</li>`); err != nil {
		panic(err)
	}

	if err = g.AddTemplate("app", `<header class="header">
	<h1>todos</h1>
	<input class="new-todo" placeholder="What needs to be done?" autofocus>
</header>
<!-- This section should be hidden by default and shown when there are todos -->
{{ if gt (len .All) 0 }}
<section class="main">
	<input class="toggle-all" type="checkbox" {{ if eq (len .All) (len .Completed) }}checked{{ end }}>
	<label for="toggle-all">Mark all as complete</label>
	<ul class="todo-list">
	</ul>
</section>
{{ end }}
<!-- This footer should hidden by default and shown when there are todos -->
{{ if gt (len .All) 0 }}
<footer class="footer">
	{{ template "partials/footer" . }}
</footer>
{{ end }}
`); err != nil {
		panic(err)
	}

	GetTemplate = g.GetTemplate
	GetPartial = g.GetPartial
	GetLayout = g.GetLayout
	MustGetTemplate = g.MustGetTemplate
	MustGetPartial = g.MustGetPartial
	MustGetLayout = g.MustGetLayout
}
