package main

var pkgTemplate = `{{define "examples_md"}}{{range .}}
#### <a name="example_{{.Name}}">Example{{example_suffix .Name}}</a>
{{comment_md .Doc}}
Code:
{{pre .Code}}{{if .Output}}
Output:
{{output .Output}}
{{end}}
{{end}}{{end}}{{with .PDoc}}
{{if $.IsMain}}
> {{ base .ImportPath }}
{{comment_md .Doc}}
{{else}}
# {{ .Name }}
` + "`" + `import "{{.ImportPath}}"` + "`" + `

* [Overview](#pkg-overview)
* [Index](#pkg-index){{if and show_examples $.Examples}}
* [Examples](#pkg-examples){{- end}}{{if $.Dirs}}
* [Subdirectories](#pkg-subdirectories){{- end}}

## <a name="pkg-overview">Overview</a>
{{comment_md .Doc}}
{{template "examples_md" (examples $ "")}}

## <a name="pkg-index">Index</a>{{if .Consts}}
* [Constants](#pkg-constants){{end}}{{if .Vars}}
* [Variables](#pkg-variables){{end}}{{- range .Funcs -}}{{$name_html := html .Name}}
* [{{node_html $ .Decl false | sanitize}}](#{{$name_html}}){{- end}}{{- range .Types}}{{$tname_html := html .Name}}
* [type {{$tname_html}}](#{{$tname_html}}){{- range .Funcs}}{{$name_html := html .Name}}
  * [{{node_html $ .Decl false | sanitize}}](#{{$name_html}}){{- end}}{{- range .Methods}}{{$name_html := html .Name}}
  * [{{node_html $ .Decl false | sanitize}}](#{{$tname_html}}.{{$name_html}}){{- end}}{{- end}}{{- if $.Notes}}{{- range $marker, $item := $.Notes}}
* [{{noteTitle $marker | html}}s](#pkg-note-{{$marker}}){{end}}{{end}}
{{if and show_examples $.Examples}}
#### <a name="pkg-examples">Examples</a>{{- range examples $ "*" }}
* [{{example_name .Name}}](#example_{{.Name}}){{- end}}{{- end}}
{{with .Filenames}}
#### <a name="pkg-files">Package files</a>
{{range .}}[{{.|filename|html}}]({{.|srcLink|html}}) {{end}}
{{end}}

{{with .Consts}}## <a name="pkg-constants">Constants</a>
{{range .}}{{node $ .Decl | pre}}
{{comment_md .Doc}}{{end}}{{end}}
{{with .Vars}}## <a name="pkg-variables">Variables</a>
{{range .}}{{node $ .Decl | pre}}
{{comment_md .Doc}}{{end}}{{end}}

{{range .Funcs}}{{$name_html := html .Name}}## <a name="{{$name_html}}">func</a> [{{$name_html}}]({{posLink_url $ .Decl}})
{{node $ .Decl | pre}}
{{comment_md .Doc}}
{{template "examples_md" (examples $ .Name)}}
{{callgraph_html $ "" .Name}}{{end}}
{{range .Types}}{{$tname := .Name}}{{$tname_html := html .Name}}## <a name="{{$tname_html}}">type</a> [{{$tname_html}}]({{posLink_url $ .Decl}})
{{node $ .Decl | pre}}
{{comment_md .Doc}}{{range .Consts}}
{{node $ .Decl | pre }}
{{comment_md .Doc}}{{end}}{{range .Vars}}
{{node $ .Decl | pre }}
{{comment_md .Doc}}{{end}}

{{template "examples_md" (examples $ $tname)}}
{{implements_html $ $tname}}
{{methodset_html $ $tname}}

{{range .Funcs}}{{$name_html := html .Name}}### <a name="{{$name_html}}">func</a> [{{$name_html}}]({{posLink_url $ .Decl}})
{{node $ .Decl | pre}}
{{comment_md .Doc}}
{{template "examples_md" (examples $ .Name)}}{{end}}
{{callgraph_html $ "" .Name}}

{{range .Methods}}{{$name_html := html .Name}}### <a name="{{$tname_html}}.{{$name_html}}">func</a> ({{md .Recv}}) [{{$name_html}}]({{posLink_url $ .Decl}})
{{node $ .Decl | pre}}
{{comment_md .Doc}}
{{$name := printf "%s_%s" $tname .Name}}{{template "examples_md" (examples $ $name)}}
{{callgraph_html $ .Recv .Name}}
{{end}}{{end}}{{end}}

{{with $.Notes}}
{{range $marker, $content := .}}
## <a name="pkg-note-{{$marker}}">{{noteTitle $marker | html}}s
<ul style="list-style: none; padding: 0;">
{{range .}}
<li><a href="{{posLink_url $ .}}">&#x261e;</a> {{html .Body}}</li>
{{end}}
</ul>
{{end}}
{{end}}
{{end}}
`
