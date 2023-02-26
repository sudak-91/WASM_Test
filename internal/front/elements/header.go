package elements

import "github.com/sudak-91/wasm-test/pkg/htmlelement"

type Header struct {
}

func CreateHeader(container htmlelement.Parent) *htmlelement.Div {
	header := htmlelement.NewDiv(container, "header")
	header.AddClass("row")
	header.AddClass("align-items-center")
	imageDiv := htmlelement.NewDiv(header, "imageDiv")
	imageDiv.AddClass("col-sm-1")
	imageDiv.AddClass("align-center")
	img := htmlelement.NewImage(imageDiv, "./src/logo.png", "logoImage")
	img.AddClass("mw-100")
	label := htmlelement.NewDiv(header, "labelDiv")
	label.AddClass("col-5")
	label.AddClass("align-items-center")
	labelText := htmlelement.NewParagaph(label, "labelText", "Я вам че автоматизатор")
	labelText.AddClass("h1")
	labelText.AddClass("mw-100")
	signInDiv := htmlelement.NewDiv(header, "signInDiv")
	signInDiv.AddClass("col-1")
	signInDiv.AddClass("align-items-center")
	/*signInCollapseButton := htmlelement.NewButton(signInDiv, "signInCollapseBtn", "button", "")
	signInCollapseButton.AddClass("btn")
	signInCollapseButton.AddClass("btn-primary")
	signInCollapseDiv := htmlelement.NewDiv(header, "collapseDiv")
	signInCollapseDiv.AddClass("collapse")
	signInCollapseDiv.AddClass("collapsexample")
	innerSignInDinv := htmlelement.NewDiv(signInCollapseDiv, "innerSignInDiv")
	innerSignInDinv.AddClass("card")
	innerSignInDinv.AddClass("card-body")*/
	//signInCollapseButton.GetJs().Set("data-bs-toggle", "collapse")
	return header
}
