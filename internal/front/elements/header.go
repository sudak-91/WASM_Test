package elements

import "github.com/sudak-91/wasm-test/pkg/htmlelement"

type Header struct {
}

func CreateHeader(container *htmlelement.Div) *htmlelement.Div {
	header := container.CreateChildDiv()
	img := htmlelement.NewImage("./src/logo.png")
	img.AddClass("mw-100")
	header.AddClass("row")
	imageDiv := header.CreateChildDiv()
	imageDiv.AddClass("col-3")
	imageDiv.AddChild(img)
	label := header.CreateChildDiv()
	label.AddClass("col-9")
	labelText := htmlelement.NewParagaph("Я вам че автоматизатор")
	labelText.AddClass("h1")
	labelText.AddClass("mw-100")
	label.AddChild(labelText)
	return header
}
