package elements

import (
	"github.com/sudak-91/wasm-test/internal/types"
	"github.com/sudak-91/wasm-test/pkg/htmlelement"
)

type MainBody struct {
}

func CreateMainBody(container *htmlelement.Div, ctx chan types.Update) *htmlelement.Div {
	mainBody := container.CreateChildDiv()
	mainBody.AddClass("row")
	leftSidebar := mainBody.CreateChildDiv()
	leftSidebar.AddClass("col-4")

	singIn := NewSignIn(ctx)
	s := singIn.CreateSignIn()

	leftSidebar.AddChild(s)
	mainFrame := mainBody.CreateChildDiv()
	mainFrame.AddClass("col-8")
	return mainBody
}
