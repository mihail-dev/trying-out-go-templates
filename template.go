package main

var nginxConfTemplate = `{{.Count}} items are made of {{.Material}}

{{if eq "gerrit" "gerrit" }} blablabla {{end}}

{{ if false }} don't write this {{ end }} 






`
