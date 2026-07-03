{{- define "go-web-app.name" -}}
{{ .Release.Name }}
{{- end }}

{{- define "go-web-app.labels" -}}
app: {{ include "go-web-app.name" . }}
{{- end }}
