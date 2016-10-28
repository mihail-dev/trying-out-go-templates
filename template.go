package main

var nginxConfTemplate = `{{.Count}} items are made of {{.Material}}

{{if eq "jenkins" "gerrit" }} 
blablabla 
{{ else if true }}
some other text 
{{end}}

{{ if false }} don't write this {{ end }} 






`
