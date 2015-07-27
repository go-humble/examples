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

	if err = g.AddPartial("errors", `<div class="alert alert-danger" role="alert">
   <ul>
   {{ range $err := . }}
      <li>{{ $err }}</li>
   {{ end }}
   </ul>
</div>
`); err != nil {
		panic(err)
	}

	if err = g.AddPartial("head", `<head>
	<title>Temple Testing</title>	
	<link href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.4/css/bootstrap.min.css" rel="stylesheet">	
</head>`); err != nil {
		panic(err)
	}

	if err = g.AddPartial("people/edit", `<h2>Edit Person</h2>
<div class="panel panel-default">
   <div class="panel-body">
      {{ template "partials/people/form" . }}
   </div>
</div>
`); err != nil {
		panic(err)
	}

	if err = g.AddPartial("people/form", `<form method="post" id="person-form" action="/people">
   <div class="form-group">
      <label for="name">Name</label>
      <input type="text" class="form-control" id="name" name="name" placeholder="Name" {{ if . }}value="{{.Name}}"{{ end }}>
   </div>
   <div class="form-group">
      <label for="age">Age</label>
      <input type="number" class="form-control" id="age" name="age" placeholder="Age" {{ if . }}value="{{.Age}}"{{ end }}>
   </div>
   <input type="submit" class="btn btn-default"></input>
</form>
`); err != nil {
		panic(err)
	}

	if err = g.AddPartial("people/index", `<div class="row">
	<div class="col-lg-12">
		<h2>All People ({{ len . }})</h2>
		<table class="table">
			<tr>
				<th>Id</th>
				<th>Name</th>
				<th>Age</th>
			</tr>
			{{ range . }}
			<tr>
				<td><a href="/people/{{ .Id }}">{{ .Id }}</a></td>
				<td>{{ .Name }}</td>
				<td>{{ .Age }}</td>
			</tr>
			{{ end }}
		</table>
	</div>
</div>
<div class="row">
	<div class="col-lg-12">
		<a href="/people/new" class="btn btn-default">New Person</a>
	</div>
</div>
`); err != nil {
		panic(err)
	}

	if err = g.AddPartial("people/new", `<h2>New Person</h2>
<div class="panel panel-default">
   <div class="panel-body">
      {{ template "partials/people/form" . }}
   </div>
</div>
`); err != nil {
		panic(err)
	}

	if err = g.AddPartial("people/show", `<h2>Person</h2>
<div class="panel panel-default">
	<div class="panel-body">
		<ul>
			<li>
				<strong>Id:</strong> {{ .Id }}
			</li>
			<li>
				<strong>Name:</strong> {{ .Name }}
			</li>
			<li>
				<strong>Age:</strong> {{ .Age }}
			</li>
		</ul>
	</div>
</div>`); err != nil {
		panic(err)
	}

	if err = g.AddLayout("app", `<!doctype html>
<html>
	{{ template "partials/head" }}
	<body>
		<div class="container" id="errors">
			{{ if .Errors }}
				{{ template "partials/errors" .Errors }}
			{{ end }}
		</div>
		<div class="container" id="main">
			{{ template "content" . }}
		</div>
		<script type="text/javascript" src="/js/app.js"></script>
	</body>
</html>
`); err != nil {
		panic(err)
	}

	if err = g.AddTemplate("people/edit", `{{ define "content" }}
	{{ template "partials/people/edit" .Person }}
{{ end }}
{{ template "layouts/app" . }}
`); err != nil {
		panic(err)
	}

	if err = g.AddTemplate("people/index", `{{ define "content" }}
	{{ template "partials/people/index" .People }}
{{ end }}
{{ template "layouts/app" . }}
`); err != nil {
		panic(err)
	}

	if err = g.AddTemplate("people/new", `{{ define "content" }}
	{{ template "partials/people/new" .Person }}
{{ end }}
{{ template "layouts/app" . }}
`); err != nil {
		panic(err)
	}

	if err = g.AddTemplate("people/show", `{{ define "content" }}
	{{ template "partials/people/show" .Person }}
{{ end }}
{{ template "layouts/app" . }}
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
