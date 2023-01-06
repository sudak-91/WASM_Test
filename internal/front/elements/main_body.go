package elements

import "github.com/sudak-91/wasm-test/pkg/htmlelement"

type MainBody struct {
}

func CreateMainBody(container *htmlelement.Div) *htmlelement.Div {
	mainBody := container.CreateChildDiv()
	mainBody.AddClass("row")
	leftSidebar := mainBody.CreateChildDiv()
	leftSidebar.AddClass("col-4")
	mainFrame := mainBody.CreateChildDiv()
	mainFrame.AddClass("col-8")
	return mainBody
}
