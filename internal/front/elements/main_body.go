package elements

import (
	"github.com/sudak-91/wasm-test/internal/types"
	"github.com/sudak-91/wasm-test/pkg/htmlelement"
)

type MainBody struct {
}

func CreateMainBody(container htmlelement.Parent, ctx chan types.Update) *htmlelement.Div {
	mainBody := htmlelement.NewDiv(container, "mainBody")
	mainBody.AddClass("row")
	leftSidebar := htmlelement.NewDiv(mainBody, "leftSideBar")
	leftSidebar.AddClass("col-4")

	singIn := NewSignIn(ctx)
	singIn.CreateSignIn(leftSidebar)

	mainFrame := htmlelement.NewDiv(mainBody, "mainFrame")
	mainFrame.AddClass("col-8")
	return mainBody
}
