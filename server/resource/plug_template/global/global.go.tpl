package global

{{- if .HasGlobal }}

import "custom-server/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}