package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"syscall/js"
	"time"

	"github.com/sudak-91/wasm-test/internal/front/elements"
	"github.com/sudak-91/wasm-test/pkg/htmlelement"
	pubupdater "github.com/sudak-91/wasm-test/pkg/updater"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

var data = make(chan pubupdater.Update)

func main() {
	fmt.Println("Start Main")
	var (
		body = htmlelement.GetBody()
		//onClick   js.Func
		//withParam js.Func

		//contDiv         = htmlelement.NewDiv()
		//header          = contDiv.CreateChildDiv()
		//centr           = contDiv.CreateChildDiv()
		//leftColumn      = header.CreateChildDiv()
		//rightColumn     = header.CreateChildDiv()
		//imageDiv        = centr.CreateChildDiv()
		//TestButton      = htmlelement.NewButton()
		//ButtonWithParam = htmlelement.NewButton()
		u = url.URL{Scheme: "ws",
			Host: "0.0.0.0:8000",
			Path: "/ws"}
	)
	Render(body)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	fmt.Println("Start Connect to", u.String())
	c, responce, err := websocket.Dial(ctx, u.String(), nil)
	if err != nil {
		log.Printf("we have an error:%s", err.Error())
	}
	fmt.Println(responce.StatusCode)
	defer c.Close(websocket.StatusInternalError, "the sky is falling")
	go func() {
		for {
			_, data, err := c.Read(ctx)
			if err != nil {
				cancel()
				return
			}
			fmt.Println(data)
		}
	}()
	go Writer(c, ctx)

	/*body.AddChild(contDiv)
	contDiv.AddClass("container")
	header.AddClass("row")
	centr.AddClass("row")
	imageDiv.AddClass("col")
	leftColumn.AddClass("col")
	leftColumn.GetJs().Set("innerHTML", "Look at here")
	rightColumn.AddClass("col")
	rightColumn.AddChild(TestButton)
	rightColumn.AddChild(ButtonWithParam)
	TestButton.AddClass("btn")
	TestButton.AddClass("btn-outline-primary")
	TestButton.AddType("button")
	TestButton.Set("innerHTML", "Click Me")
	ButtonWithParam.AddClass("btn")
	ButtonWithParam.AddClass("btn-outline-primary")
	ButtonWithParam.Set("innerHTML", "Param")
	ButtonWithParam.Set("value", "testing")
	onClick = js.FuncOf(TestButtonClick)
	withParam = js.FuncOf(TestButtonWithParam)
	TestButton.AddClickEventListener(&onClick)
	js.Global().Set("wParam", withParam)
	ButtonWithParam.GetJs().Call("setAttribute", "onclick", "wParam(1,2)")

	rawParam := js.Global().Get("location").Get("href")
	param := rawParam.String()
	route := strings.Split(param, "/")
	fmt.Println(route)
	leftColumn.GetJs().Set("id", "main")*/
	fmt.Println("Hello WASM")
	b := make(chan bool)
	<-b
	fmt.Println("Close")
}
func Render(body *htmlelement.Body) {
	fmt.Println("Start Render")
	container := htmlelement.NewDiv(body, "container")
	container.AddClass("container")
	fmt.Println("Render header")
	elements.CreateHeader(container)
	fmt.Println("Start Main Body")
	elements.CreateMainBody(container, data)
	fmt.Println("Stop Render")
	body.Render()

}
func Writer(c *websocket.Conn, ctx context.Context) {
	for {
		k := <-data
		err := wsjson.Write(ctx, c, k)
		if err != nil {
			log.Println(err.Error())
		}

	}
}

func TestButtonClick(this js.Value, args []js.Value) any {
	fmt.Println("click")
	jsGlobal := js.Global().Get("document")
	div := jsGlobal.Call("getElementById", "main")
	div.Set("innerHTML", "Yes.sir")

	return nil
}
func TestButtonWithParam(this js.Value, args []js.Value) any {
	//a := this.Get("value")
	//fmt.Println(a)
	fmt.Println(args[0].String())
	data <- pubupdater.Update{Type: "test"}
	return nil
}

func Newtest() {
	fmt.Println("Yeppy")
}
