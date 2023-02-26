package elements

import (
	"github.com/sudak-91/wasm-test/pkg/htmlelement"
	pubupdater "github.com/sudak-91/wasm-test/pkg/updater"
)

type MainBody struct {
}

func CreateMainBody(container htmlelement.Parent, ctx chan pubupdater.Update) *htmlelement.Div {
	mainBody := htmlelement.NewDiv(container, "mainBody")
	mainBody.AddClass("row")
	leftSidebar := htmlelement.NewDiv(mainBody, "leftSideBar")
	leftSidebar.AddClass("col-4")

	singIn := NewSignIn(ctx)
	singIn.CreateSignIn(leftSidebar)

	mainFrame := htmlelement.NewDiv(mainBody, "mainFrame")
	mainFrame.AddClass("col-8")
	registerFrame := htmlelement.NewDiv(mainFrame, "registers")

	regForm := NewRegistrtion(ctx)
	regForm.CreateRegistrationForm(registerFrame)
	return mainBody
}
