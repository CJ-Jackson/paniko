package www

import (
	"html/template"

	"github.com/CJ-Jackson/ctx"
	"github.com/CJ-Jackson/paniko/paniko/shared"
	"github.com/CJ-Jackson/paniko/paniko/uri"
)

const indexName = "index-08a83f56f55ae2aed3d2b78c58d2645c"

func buildIndexTemplate(context ctx.BackgroundContext) *template.Template {
	indexUriReverse := uri.Reverse{
		"IAmAlive": uri.IAmAlive,
		"InDanger": uri.InDanger,
		"LogOut":   uri.Logout,
	}

	funcMaps := template.FuncMap{
		"index":    func(context ctx.Context) HomeViewIndexData { return context.Data(indexName).(HomeViewIndexData) },
		"indexUri": func() uri.Reverse { return indexUriReverse },
	}

	return template.Must(shared.CloneMasterTemplate(context).Funcs(funcMaps).Parse(indexTemplate))
}

type HomeViewIndexData struct {
	Expiry string
}

const indexTemplate = `{{ define "content" }}{{- $index := index . -}}{{- $uri := indexUri -}}
<p>Hello, please report by pressing a button</p>

<pre>Expires: <span id="expiry">{{ $index.Expiry }}</span> <span id="alert"></span></pre>

<div class="clickReport btn btn-success mt-3 mb-3" data-confirm="false" data-uri='{{ $uri.Print "IAmAlive" }}' style="width: 100%">I am Alive</div>

<div class="clickReport btn btn-danger" data-confirm="true" data-uri='{{ $uri.Print "InDanger" }}' style="width: 100%">In Danger</div>

<p class="mt-3">If the system expires, an email will be sent out to the authority every hour.</p>

<a href='{{ $uri.Print "LogOut" }}' class="btn btn-secondary mt-3" style="width: 100%">Logout</a>
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
				});
	        }
	    });
	})(jQuery);
</script>
{{ end }}`
