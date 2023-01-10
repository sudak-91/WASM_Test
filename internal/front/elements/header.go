package elements

import "github.com/sudak-91/wasm-test/pkg/htmlelement"

type Header struct {
}

func CreateHeader(container htmlelement.Parent) *htmlelement.Div {
	header := htmlelement.NewDiv(container, "header")
	header.AddClass("row")
	//img := htmlelement.NewImage("./src/logo.png")
	//	img.AddClass("mw-100")
	imageDiv := htmlelement.NewDiv(header, "imageDiv")
	imageDiv.AddClass("col-3")
	img := htmlelement.NewImage(imageDiv, "./src/logo.png", "logoImage")
	img.AddClass("mw-100")
	label := htmlelement.NewDiv(header, "labelDiv")
	label.AddClass("col-9")
	labelText := htmlelement.NewParagaph(label, "labelText", "Я вам че автоматизатор")
	labelText.AddClass("h1")
	labelText.AddClass("mw-100")
	return header
}
