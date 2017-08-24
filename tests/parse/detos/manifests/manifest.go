//+build ignore

// Package manifests is auto-generated and should not be modified by hand.
// This package contains a virtual file system for generate resources which are not accessed
// through a remote endpoint (i.e those resources generated from the manifests that are local in the
// filesystem and are not marked as remote in access).
package manifests

import (
	"path/filepath"

	"github.com/spf13/afero"
)

// AppFS defines the global handler for which all assets generated from manifests
// files which are not remote resources are provided as binary embedded assets.
var AppFS = afero.NewMemMapFs()

// addFile adds a giving file name into the file system.
func addFile(path string, content []byte) {
	dir, _ := filepath.Split(path)
	if dir != "" {
		AppFS.MkdirAll(dir, 0755)
	}

	afero.WriteFile(AppFS, path, content, 0644)
}

func init() {

	addFile("sandles.js", []byte("Ly8gUGFja2FnZSBkZXRvcyBkZWZpbmVzIGEgc2VyaWVzIG9mIGNvbXBvbmVudCBmb3IgdGhlIGRldG9zIHByb2plY3QKLy8KLy9zaGVsbDpjb21wb25lbnQ6Z2xvYmFsCi8vCi8vIFJlc291cmNlIHsKLy8gICAgIE5hbWU6IGRldG9zLmZvbnQuanMKLy8gICAgIFBhdGg6IGh0dHBzOi8vZm9udHMuZ29vZ2xlYXBpcy5jb20vY3NzP2ZhbWlseT1MYXRvfE9wZW4rU2Fuc3xSb2JvdG8KLy8gICAgIEhvb2s6IGVtYmVkZGVkLWNzcwovLwkJIExvY2FsaXplOiB0cnVlCi8vIH0KLy8KcGFja2FnZSBkZXRvcwoKaW1wb3J0ICgKCSJnaXRodWIuY29tL2d1LWlvL2d1L3Rlc3RzL3BhcnNlL2RldHJvIgopCgovLyBIb2xkQ29tcG9uZW50IGRlZmluZXMgYSBjb21wb25lbnQgdG8gcHJvdmlkZSBhIGhvbGQgdHlwZS4KLy8KLy9zaGVsbDpjb21wb25lbnQKLyoKUmVzb3VyY2UgewogICAgTmFtZTogZmFudG9tLmpzCiAgICBQYXRoOiBodHRwOi8vY2RsLmdvZy5jb20vcmVzL3N0YXRpY3MvZmFudG9tLmpzCiAgICBIb29rOiBqYXZhc2NyaXB0Cn0KClJlc291cmNlIHsKCQlWZXJzaW9uOiAxLjQuMQoJCVBrZzogRGVudHVzVnV6CiAgICBOYW1lOiBodWxfaHViLmpzCiAgICBDb250ZW50OiBgYGAKICAgICAgICA8c2NyaXB0IHNyYz0iIj4KCQkJCQl2YXIgbW8gPSBmdW5jKCkgUmVzb3VyY2UgewoJCQkJCQlyZXR1cm4gYGBgYmxvY2sgUmVzb3VyY2UgewoJCQkJCQkJbmFtZTogJ2JvY2snLAoJCQkJCQl9YGBgCgkJCQkJfQoJCQkJPC9zY3JpcHQ+CiAgICBgYGAKICAgIEhvb2s6IGVtYmVkLmpzCiAgICBIb29rUmVwbzogZ2l0aHViLmNvbS9yZXNrL2NvbXBvLXdlYi9ob29rcwogICAgU2l6ZTogNDQwMDMwCn0KKi8KdHlwZSBIb2xkQ29tcG9uZW50IHN0cnVjdCB7CglkZXRyby5EZXRyb0NvbXBvbmVudAoJR3VsTmFtZSBzdHJpbmcgICAgICAgICBganNvbjoiZ3VsX25hbWUiYAoJQnVsICAgICBTb2xpZENvbXBvbmVudCBganNvbjoiYnVsImAKfQoKLy8gU29saWRDb21wb25lbnQgZGVmaW5lcyBhIGNvbXBvbmVudCB0byBwcm92aWRlIGEgc29saWQgYm94IHR5cGUuCi8vCi8vIHNoZWxsOmNvbXBvbmVudAovLyBSZXNvdXJjZSB7Ci8vICAgICBOYW1lOiBzaGVsbC1zdS5qcwovLyAgICAgUGF0aDogaHR0cDovL2NkbC5nb2cuY29tL3Jlcy9zdGF0aWNzL3NoZWxsLXN1LmpzCi8vICAgICBIb29rOiBqYXZhc2NyaXB0Ci8vIH0KLy8KLy8gUmVzb3VyY2UgewovLyAgICAgTmFtZTogY3JvY2suanMKLy8gICAgIENvbnRlbnQ6IGBgYAovLyAgICAgICAgIDxzY3JpcHQgc3JjPSIiPgovLwkJCQkJewovLwkJCQkJCW5hbWU6ICJib2IiLAovLwkJCQkJCWFnZTogMTIsCi8vCQkJCQl9Ci8vCQkJCQk8L3NjcmlwdD4KLy8gICAgIGBgYAovLyAgICAgSG9vazogZW1iZWQuanMKLy8gICAgIEhvb2tSZXBvOiBnaXRodWIuY29tL3Jlc2svY29tcG8td2ViL2hvb2tzCi8vICAgICBTaXplOiA0NDAwMzAKLy8gfQovLwp0eXBlIFNvbGlkQ29tcG9uZW50IHN0cnVjdCB7CglOYW1lIHN0cmluZyBganNvbjoibmFtZSJgCn0K"))

	addFile("detos.font.js", []byte("QGZvbnQtZmFjZSB7CiAgZm9udC1mYW1pbHk6ICdMYXRvJzsKICBmb250LXN0eWxlOiBub3JtYWw7CiAgZm9udC13ZWlnaHQ6IDQwMDsKICBzcmM6IGxvY2FsKCdMYXRvIFJlZ3VsYXInKSwgbG9jYWwoJ0xhdG8tUmVndWxhcicpLCB1cmwoaHR0cHM6Ly9mb250cy5nc3RhdGljLmNvbS9zL2xhdG8vdjExL3YwU2RjR0ZBbDJhZXpNOVZxX2FGVFEudHRmKSBmb3JtYXQoJ3RydWV0eXBlJyk7Cn0KQGZvbnQtZmFjZSB7CiAgZm9udC1mYW1pbHk6ICdPcGVuIFNhbnMnOwogIGZvbnQtc3R5bGU6IG5vcm1hbDsKICBmb250LXdlaWdodDogNDAwOwogIHNyYzogbG9jYWwoJ09wZW4gU2FucycpLCBsb2NhbCgnT3BlblNhbnMnKSwgdXJsKGh0dHBzOi8vZm9udHMuZ3N0YXRpYy5jb20vcy9vcGVuc2Fucy92MTMvY0paS2VPdUJybjRrRVJ4cXRhVUgzYUNXY3luZl9jRHhYd0NMeGlpeEcxYy50dGYpIGZvcm1hdCgndHJ1ZXR5cGUnKTsKfQpAZm9udC1mYWNlIHsKICBmb250LWZhbWlseTogJ1JvYm90byc7CiAgZm9udC1zdHlsZTogbm9ybWFsOwogIGZvbnQtd2VpZ2h0OiA0MDA7CiAgc3JjOiBsb2NhbCgnUm9ib3RvJyksIGxvY2FsKCdSb2JvdG8tUmVndWxhcicpLCB1cmwoaHR0cHM6Ly9mb250cy5nc3RhdGljLmNvbS9zL3JvYm90by92MTUvek43R0JGd2ZNUDR1QTZBUjBIQ29MUS50dGYpIGZvcm1hdCgndHJ1ZXR5cGUnKTsKfQo="))

	addFile("hul_hub.js", []byte("<script src=\"\">\n\t\t\t\t\tvar mo = func() Resource {\n\t\t\t\t\t\treturn ```block Resource {\n\t\t\t\t\t\t\tname: 'bock',\n\t\t\t\t\t\t}```\n\t\t\t\t\t}\n\t\t\t\t</script>"))

	addFile("crock.js", []byte("<script src=\"\">\n\t\t\t\t\t{\n\t\t\t\t\t\tname: \"bob\",\n\t\t\t\t\t\tage: 12,\n\t\t\t\t\t}\n\t\t\t\t\t</script>"))

	addFile("manifest.json", []byte("[\n\t{\n\t\t\"global_scope\": true,\n\t\t\"id\": \"JUpX7LXPV7OLVPT\",\n\t\t\"name\": \"detos.detos\",\n\t\t\"depends\": null,\n\t\t\"manifests\": [\n\t\t\t{\n\t\t\t\t\"size\": 2028,\n\t\t\t\t\"remote\": false,\n\t\t\t\t\"localize\": false,\n\t\t\t\t\"b64_encode\": true,\n\t\t\t\t\"content_b64\": true,\n\t\t\t\t\"appmanifest_id\": \"JUpX7LXPV7OLVPT\",\n\t\t\t\t\"name\": \"sandles.js\",\n\t\t\t\t\"path\": \"./detos.go\",\n\t\t\t\t\"content\": \"Ly8gUGFja2FnZSBkZXRvcyBkZWZpbmVzIGEgc2VyaWVzIG9mIGNvbXBvbmVudCBmb3IgdGhlIGRldG9zIHByb2plY3QKLy8KLy9zaGVsbDpjb21wb25lbnQ6Z2xvYmFsCi8vCi8vIFJlc291cmNlIHsKLy8gICAgIE5hbWU6IGRldG9zLmZvbnQuanMKLy8gICAgIFBhdGg6IGh0dHBzOi8vZm9udHMuZ29vZ2xlYXBpcy5jb20vY3NzP2ZhbWlseT1MYXRvfE9wZW4rU2Fuc3xSb2JvdG8KLy8gICAgIEhvb2s6IGVtYmVkZGVkLWNzcwovLwkJIExvY2FsaXplOiB0cnVlCi8vIH0KLy8KcGFja2FnZSBkZXRvcwoKaW1wb3J0ICgKCSJnaXRodWIuY29tL2d1LWlvL2d1L3Rlc3RzL3BhcnNlL2RldHJvIgopCgovLyBIb2xkQ29tcG9uZW50IGRlZmluZXMgYSBjb21wb25lbnQgdG8gcHJvdmlkZSBhIGhvbGQgdHlwZS4KLy8KLy9zaGVsbDpjb21wb25lbnQKLyoKUmVzb3VyY2UgewogICAgTmFtZTogZmFudG9tLmpzCiAgICBQYXRoOiBodHRwOi8vY2RsLmdvZy5jb20vcmVzL3N0YXRpY3MvZmFudG9tLmpzCiAgICBIb29rOiBqYXZhc2NyaXB0Cn0KClJlc291cmNlIHsKCQlWZXJzaW9uOiAxLjQuMQoJCVBrZzogRGVudHVzVnV6CiAgICBOYW1lOiBodWxfaHViLmpzCiAgICBDb250ZW50OiBgYGAKICAgICAgICA8c2NyaXB0IHNyYz0iIj4KCQkJCQl2YXIgbW8gPSBmdW5jKCkgUmVzb3VyY2UgewoJCQkJCQlyZXR1cm4gYGBgYmxvY2sgUmVzb3VyY2UgewoJCQkJCQkJbmFtZTogJ2JvY2snLAoJCQkJCQl9YGBgCgkJCQkJfQoJCQkJPC9zY3JpcHQ+CiAgICBgYGAKICAgIEhvb2s6IGVtYmVkLmpzCiAgICBIb29rUmVwbzogZ2l0aHViLmNvbS9yZXNrL2NvbXBvLXdlYi9ob29rcwogICAgU2l6ZTogNDQwMDMwCn0KKi8KdHlwZSBIb2xkQ29tcG9uZW50IHN0cnVjdCB7CglkZXRyby5EZXRyb0NvbXBvbmVudAoJR3VsTmFtZSBzdHJpbmcgICAgICAgICBganNvbjoiZ3VsX25hbWUiYAoJQnVsICAgICBTb2xpZENvbXBvbmVudCBganNvbjoiYnVsImAKfQoKLy8gU29saWRDb21wb25lbnQgZGVmaW5lcyBhIGNvbXBvbmVudCB0byBwcm92aWRlIGEgc29saWQgYm94IHR5cGUuCi8vCi8vIHNoZWxsOmNvbXBvbmVudAovLyBSZXNvdXJjZSB7Ci8vICAgICBOYW1lOiBzaGVsbC1zdS5qcwovLyAgICAgUGF0aDogaHR0cDovL2NkbC5nb2cuY29tL3Jlcy9zdGF0aWNzL3NoZWxsLXN1LmpzCi8vICAgICBIb29rOiBqYXZhc2NyaXB0Ci8vIH0KLy8KLy8gUmVzb3VyY2UgewovLyAgICAgTmFtZTogY3JvY2suanMKLy8gICAgIENvbnRlbnQ6IGBgYAovLyAgICAgICAgIDxzY3JpcHQgc3JjPSIiPgovLwkJCQkJewovLwkJCQkJCW5hbWU6ICJib2IiLAovLwkJCQkJCWFnZTogMTIsCi8vCQkJCQl9Ci8vCQkJCQk8L3NjcmlwdD4KLy8gICAgIGBgYAovLyAgICAgSG9vazogZW1iZWQuanMKLy8gICAgIEhvb2tSZXBvOiBnaXRodWIuY29tL3Jlc2svY29tcG8td2ViL2hvb2tzCi8vICAgICBTaXplOiA0NDAwMzAKLy8gfQovLwp0eXBlIFNvbGlkQ29tcG9uZW50IHN0cnVjdCB7CglOYW1lIHN0cmluZyBganNvbjoibmFtZSJgCn0K\",\n\t\t\t\t\"cache\": \"\",\n\t\t\t\t\"meta\": {},\n\t\t\t\t\"hook_name\": \"embed-js\"\n\t\t\t}\n\t\t],\n\t\t\"relation\": {\n\t\t\t\"name\": \"detos\",\n\t\t\t\"package\": \"github.com/gu-io/gu/tests/parse/detos\"\n\t\t}\n\t},\n\t{\n\t\t\"global_scope\": true,\n\t\t\"id\": \"wxMKds88RZVqHup\",\n\t\t\"name\": \"detos.detos\",\n\t\t\"depends\": null,\n\t\t\"manifests\": [\n\t\t\t{\n\t\t\t\t\"size\": 920,\n\t\t\t\t\"remote\": true,\n\t\t\t\t\"localize\": true,\n\t\t\t\t\"b64_encode\": true,\n\t\t\t\t\"content_b64\": true,\n\t\t\t\t\"appmanifest_id\": \"wxMKds88RZVqHup\",\n\t\t\t\t\"name\": \"detos.font.js\",\n\t\t\t\t\"path\": \"https://fonts.googleapis.com/css?family=Lato|Open+Sans|Roboto\",\n\t\t\t\t\"content\": \"QGZvbnQtZmFjZSB7CiAgZm9udC1mYW1pbHk6ICdMYXRvJzsKICBmb250LXN0eWxlOiBub3JtYWw7CiAgZm9udC13ZWlnaHQ6IDQwMDsKICBzcmM6IGxvY2FsKCdMYXRvIFJlZ3VsYXInKSwgbG9jYWwoJ0xhdG8tUmVndWxhcicpLCB1cmwoaHR0cHM6Ly9mb250cy5nc3RhdGljLmNvbS9zL2xhdG8vdjExL3YwU2RjR0ZBbDJhZXpNOVZxX2FGVFEudHRmKSBmb3JtYXQoJ3RydWV0eXBlJyk7Cn0KQGZvbnQtZmFjZSB7CiAgZm9udC1mYW1pbHk6ICdPcGVuIFNhbnMnOwogIGZvbnQtc3R5bGU6IG5vcm1hbDsKICBmb250LXdlaWdodDogNDAwOwogIHNyYzogbG9jYWwoJ09wZW4gU2FucycpLCBsb2NhbCgnT3BlblNhbnMnKSwgdXJsKGh0dHBzOi8vZm9udHMuZ3N0YXRpYy5jb20vcy9vcGVuc2Fucy92MTMvY0paS2VPdUJybjRrRVJ4cXRhVUgzYUNXY3luZl9jRHhYd0NMeGlpeEcxYy50dGYpIGZvcm1hdCgndHJ1ZXR5cGUnKTsKfQpAZm9udC1mYWNlIHsKICBmb250LWZhbWlseTogJ1JvYm90byc7CiAgZm9udC1zdHlsZTogbm9ybWFsOwogIGZvbnQtd2VpZ2h0OiA0MDA7CiAgc3JjOiBsb2NhbCgnUm9ib3RvJyksIGxvY2FsKCdSb2JvdG8tUmVndWxhcicpLCB1cmwoaHR0cHM6Ly9mb250cy5nc3RhdGljLmNvbS9zL3JvYm90by92MTUvek43R0JGd2ZNUDR1QTZBUjBIQ29MUS50dGYpIGZvcm1hdCgndHJ1ZXR5cGUnKTsKfQo=\",\n\t\t\t\t\"cache\": \"\",\n\t\t\t\t\"meta\": {},\n\t\t\t\t\"hook_name\": \"embedded-css\"\n\t\t\t}\n\t\t],\n\t\t\"relation\": {\n\t\t\t\"name\": \"detos\",\n\t\t\t\"package\": \"github.com/gu-io/gu/tests/parse/detos\"\n\t\t}\n\t},\n\t{\n\t\t\"global_scope\": false,\n\t\t\"id\": \"CeWvadUX8pu62Ro\",\n\t\t\"name\": \"detos.HoldComponent\",\n\t\t\"depends\": null,\n\t\t\"manifests\": [\n\t\t\t{\n\t\t\t\t\"size\": 0,\n\t\t\t\t\"remote\": true,\n\t\t\t\t\"localize\": false,\n\t\t\t\t\"b64_encode\": true,\n\t\t\t\t\"content_b64\": false,\n\t\t\t\t\"appmanifest_id\": \"CeWvadUX8pu62Ro\",\n\t\t\t\t\"name\": \"fantom.js\",\n\t\t\t\t\"path\": \"http://cdl.gog.com/res/statics/fantom.js\",\n\t\t\t\t\"content\": \"\",\n\t\t\t\t\"cache\": \"\",\n\t\t\t\t\"meta\": {},\n\t\t\t\t\"hook_name\": \"javascript\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"size\": 440030,\n\t\t\t\t\"remote\": false,\n\t\t\t\t\"localize\": false,\n\t\t\t\t\"b64_encode\": true,\n\t\t\t\t\"content_b64\": false,\n\t\t\t\t\"appmanifest_id\": \"CeWvadUX8pu62Ro\",\n\t\t\t\t\"name\": \"hul_hub.js\",\n\t\t\t\t\"path\": \"\",\n\t\t\t\t\"content\": \"\\u003cscript src=\\\"\\\"\\u003e\\n\\t\\t\\t\\t\\tvar mo = func() Resource {\\n\\t\\t\\t\\t\\t\\treturn ```block Resource {\\n\\t\\t\\t\\t\\t\\t\\tname: 'bock',\\n\\t\\t\\t\\t\\t\\t}```\\n\\t\\t\\t\\t\\t}\\n\\t\\t\\t\\t\\u003c/script\\u003e\",\n\t\t\t\t\"cache\": \"\",\n\t\t\t\t\"meta\": {\n\t\t\t\t\t\"HookRepo\": \"github.com/resk/compo-web/hooks\",\n\t\t\t\t\t\"Pkg\": \"DentusVuz\",\n\t\t\t\t\t\"Version\": \"1.4.1\"\n\t\t\t\t},\n\t\t\t\t\"hook_name\": \"embed.js\"\n\t\t\t}\n\t\t],\n\t\t\"relation\": {\n\t\t\t\"name\": \"HoldComponent\",\n\t\t\t\"package\": \"github.com/gu-io/gu/tests/parse/detos\",\n\t\t\t\"composites\": [\n\t\t\t\t\"DetroComponent\"\n\t\t\t],\n\t\t\t\"fieldtypes\": [\n\t\t\t\t\"SolidComponent\"\n\t\t\t]\n\t\t}\n\t},\n\t{\n\t\t\"global_scope\": false,\n\t\t\"id\": \"76LH18PnvIuuN2H\",\n\t\t\"name\": \"detos.SolidComponent\",\n\t\t\"depends\": null,\n\t\t\"manifests\": [\n\t\t\t{\n\t\t\t\t\"size\": 0,\n\t\t\t\t\"remote\": true,\n\t\t\t\t\"localize\": false,\n\t\t\t\t\"b64_encode\": true,\n\t\t\t\t\"content_b64\": false,\n\t\t\t\t\"appmanifest_id\": \"76LH18PnvIuuN2H\",\n\t\t\t\t\"name\": \"shell-su.js\",\n\t\t\t\t\"path\": \"http://cdl.gog.com/res/statics/shell-su.js\",\n\t\t\t\t\"content\": \"\",\n\t\t\t\t\"cache\": \"\",\n\t\t\t\t\"meta\": {},\n\t\t\t\t\"hook_name\": \"javascript\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"size\": 440030,\n\t\t\t\t\"remote\": false,\n\t\t\t\t\"localize\": false,\n\t\t\t\t\"b64_encode\": true,\n\t\t\t\t\"content_b64\": false,\n\t\t\t\t\"appmanifest_id\": \"76LH18PnvIuuN2H\",\n\t\t\t\t\"name\": \"crock.js\",\n\t\t\t\t\"path\": \"\",\n\t\t\t\t\"content\": \"\\u003cscript src=\\\"\\\"\\u003e\\n\\t\\t\\t\\t\\t{\\n\\t\\t\\t\\t\\t\\tname: \\\"bob\\\",\\n\\t\\t\\t\\t\\t\\tage: 12,\\n\\t\\t\\t\\t\\t}\\n\\t\\t\\t\\t\\t\\u003c/script\\u003e\",\n\t\t\t\t\"cache\": \"\",\n\t\t\t\t\"meta\": {\n\t\t\t\t\t\"HookRepo\": \"github.com/resk/compo-web/hooks\"\n\t\t\t\t},\n\t\t\t\t\"hook_name\": \"embed.js\"\n\t\t\t}\n\t\t],\n\t\t\"relation\": {\n\t\t\t\"name\": \"SolidComponent\",\n\t\t\t\"package\": \"github.com/gu-io/gu/tests/parse/detos\"\n\t\t}\n\t}\n]"))

}
