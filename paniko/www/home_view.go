package www

import (
	"html/template"

	"encoding/json"

	"github.com/CJ-Jackson/ctx"
	"github.com/CJ-Jackson/paniko/paniko/shared"
	"github.com/CJ-Jackson/paniko/paniko/uri"
)

type HomeView interface {
	Index(context ctx.Context, data HomeViewIndexData)
	Json(context ctx.Context, data JsonData)
}

func NewHomeView(context ctx.BackgroundContext) HomeView {
	indexName := "index-08a83f56f55ae2aed3d2b78c58d2645c"
	indexUriReverse := uri.Reverse{
		"IAmAlive": uri.IAmAlive,
		"InDanger": uri.InDanger,
	}
	indexTemplate := template.Must(shared.CloneMasterTemplate(context).Funcs(template.FuncMap{
		"index": func(context ctx.Context) HomeViewIndexData {
			return context.Data(indexName).(HomeViewIndexData)
		},
		"indexUri": func() uri.Reverse { return indexUriReverse },
	}).Parse(`
{{ define "content" }}{{- $index := index . -}}{{- $uri := indexUri -}}
<p>Hello, please report by pressing a button</p>

<code>Expires: <span id="expiry">{{ $index.Expiry }}</span> <span id="alert"></span></code>

<div class="clickReport btn btn-success mt-3 mb-3" data-confirm="false" data-uri='{{ $uri.Print "IAmAlive" }}' style="width: 100%">I am Alive</div>

<div class="clickReport btn btn-danger" data-confirm="true" data-uri='{{ $uri.Print "InDanger" }}' style="width: 100%">In Danger</div>

<p class="mt-3">If the system expires, an email will be sent out to the authority every hour.</p>
{{- end }}

{{ define "js" }}
<script type="application/javascript">
	(function($) {
	    $('.clickReport').on('click', function() {
	        var doConfirm = $(this).data('confirm');
	        var uri = $(this).data('uri');
	        
	        if (!doConfirm || confirm('Are you sure?')) {
				$.ajax(uri, {
				    method: 'PUT'
				}).done(function(data) {
				    data = $.parseJSON(data);
				    $('#expiry').text(data.When);
				    $('#alert').text(data.Alert);
				})
	        }
	    })
	})(jQuery);
</script>
{{ end }}
`))

	return homeView{
		indexName:     indexName,
		indexTemplate: indexTemplate,
	}
}

type homeView struct {
	indexName     string
	indexTemplate *template.Template
}

type HomeViewIndexData struct {
	Expiry string
}

func (v homeView) Index(context ctx.Context, data HomeViewIndexData) {
	context.SetTitle("Paniko")
	context.SetData(v.indexName, data)

	v.indexTemplate.Execute(context.Response(), context)
}

type JsonData struct {
	Alert string
	When  string
}

func (v homeView) Json(context ctx.Context, data JsonData) {
	json.NewEncoder(context.Response()).Encode(data)
}
