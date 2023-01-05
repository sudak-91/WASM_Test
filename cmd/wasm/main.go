package main

import (
	"fmt"
	"strings"
	"syscall/js"

	"github.com/sudak-91/wasm-test/pkg/htmlelement"
)

func main() {
	var (
		onClick     js.Func
		document    htmlelement.Document
		contDiv     = document.CreateDiv()
		header      = document.CreateDiv()
		leftColomn  = document.CreateDiv()
		rightColomn = document.CreateDiv()
		centr       = document.CreateDiv()
		imageDiv    = document.CreateDiv()
		logoImage   = document.CreateImg()
	)
	go func() {
		for i := 1; i < 100; i++ {
			fmt.Println(i)
		}
	}()

	contDiv.AddClass("container")
	header.AddClass("row")
	centr.AddClass("row")
	imageDiv.AddClass("col")
	logoImage.Set("src", "./src/logo.png")
	leftColomn.AddClass("col")
	leftColomn.GetJs().Set("innerHTML", "Look at here")
	rightColomn.AddClass("col")
	//contDiv.AddChild(centr.GetJs())
	rawParam := js.Global().Get("location").Get("href")
	param := rawParam.String()
	route := strings.Split(param, "/")
	fmt.Println(route)
	onClick = js.FuncOf(TestButtonClick)
	body := document.GetJS().Call("querySelector", "body")
	button := document.GetJS().Call("createElement", "button")
	button.Get("classList").Call("add", "btn")
	button.Get("classList").Call("add", "btn-outline-primary")
	button.Set("id", "testbtn")
	button.Set("type", "button")
	button.Set("innerHTML", "ClickMe")
	button.Call("addEventListener", "click", onClick)
	rightColomn.AddChild(button)
	leftColomn.GetJs().Set("id", "main")
	header.AddChild(leftColomn.GetJs())
	header.AddChild(rightColomn.GetJs())
	contDiv.AddChild(header.GetJs())
	imageDiv.AddChild(logoImage)
	centr.AddChild(imageDiv.GetJs())
	contDiv.AddChild(centr.GetJs())
	rightColomn.Render()
	leftColomn.Render()
	header.Render()
	imageDiv.Render()
	centr.Render()
	contDiv.Render()
	body.Call("appendChild", contDiv.GetJs())
	fmt.Println("Hello WASM")
	b := make(chan bool)
	<-b
}

func TestButtonClick(this js.Value, args []js.Value) any {
	fmt.Println("click")
	jsGlobal := js.Global().Get("document")
	div := jsGlobal.Call("getElementById", "main")
	div.Set("innerHTML", "Yes.sir")

	return nil
}
