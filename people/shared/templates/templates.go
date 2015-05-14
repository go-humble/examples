package templates

// This package has been automatically generated with temple.
// Do not edit manually!

import (
	"github.com/go-humble/temple"
)

func init() {
	var err error

	if err = temple.AddPartial("head", `<head>
	<title>Temple Testing</title>	
	<link href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.4/css/bootstrap.min.css" rel="stylesheet">	
</head>`); err != nil {
		panic(err)
	}

	if err = temple.AddPartial("people", `<div class="row">
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
`); err != nil {
		panic(err)
	}

	if err = temple.AddPartial("person", `<h2>Person</h2>
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

	if err = temple.AddLayout("app", `<!doctype html>
<html>
	{{ template "partials/head" }}
	<body>
		<div class="container" id="main">
			{{ template "content" . }}
		</div>		
		<script type="text/javascript" src="/js/app.js"></script>
	</body>
</html>`); err != nil {
		panic(err)
	}

	if err = temple.AddTemplate("people/index", `{{ define "content" }}
	{{ template "partials/people" . }}
{{ end }}
{{ template "layouts/app" . }}`); err != nil {
		panic(err)
	}

	if err = temple.AddTemplate("people/show", `{{ define "content" }}
	{{ template "partials/person" . }}
{{ end }}
{{ template "layouts/app" . }}`); err != nil {
		panic(err)
	}

}
