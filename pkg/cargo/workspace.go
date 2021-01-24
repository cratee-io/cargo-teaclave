package cargo

import (
	"bytes"
	"fmt"
	"text/template"
)

const workspaceTmpl = `
[workspace]
members = [{{range .Members}}"{{.}}", {{end}}]
{{range $src, $crate := .Patches}}
[patch."{{$src}}"]
{{range $name,$config := $crate}}{{$name}} = { git = "{{$config.Git}}", rev = "{{$config.Rev}}" }
{{end}}{{end}}
`

type Patch struct {
	Old    string
	NewGit string
	Crates map[string]string // crate name => revision
}

type miniCrate struct {
	Git string
	Rev string
}

type miniWorkspace struct {
	Members []string
	Patches map[string]map[string]miniCrate
}

func NewWorkspaceManifest(members []string, patches ...Patch) ([]byte, error) {
	patch := make(map[string]map[string]miniCrate)
	for _, v := range patches {
		p := make(map[string]miniCrate)
		for k, vv := range v.Crates {
			p[k] = miniCrate{Git: v.NewGit, Rev: vv}
		}

		patch[v.Old] = p
	}

	workspace := &miniWorkspace{Members: members, Patches: patch}

	out := new(bytes.Buffer)
	if err := template.Must(template.New("testing-workspace").Parse(workspaceTmpl)).
		Execute(out, workspace); err != nil {
		return nil, fmt.Errorf("fill manifest: %w", err)
	}

	return out.Bytes(), nil
}
