package elements

import "github.com/sudak-91/wasm-test/pkg/htmlelement"

type Header struct {
}

func CreateHeader(container *htmlelement.Div) *htmlelement.Div {
	header := container.CreateChildDiv()
	header.AddClass("row")
	imageDiv := header.CreateChildDiv()
	imageDiv.AddClass("col-3")
	label := header.CreateChildDiv()
	label.AddClass("col-9")
	return header
}
