package elements

import (
	"fmt"
	"syscall/js"

	"github.com/sudak-91/wasm-test/internal/types"
	"github.com/sudak-91/wasm-test/pkg/htmlelement"
)

type Signin struct {
	Sender chan types.Update
}

func NewSignIn(c chan types.Update) Signin {
	var s Signin
	s.Sender = c
	return s

}

func (s Signin) CreateSignIn() *htmlelement.Div {
	container := htmlelement.NewDiv()
	container.AddClass("container")
	loginDiv := htmlelement.NewDiv()
	loginDiv.AddClass("input-group")
	loginDiv.AddClass("mb-3")
	passDiv := htmlelement.NewDiv()
	passDiv.AddClass("input-group")
	passDiv.AddClass("mb-3")
	loginSpan := htmlelement.NewSpan("Login")
	loginSpan.AddClass("input-group-text")
	loginInput := htmlelement.NewInput("text")
	loginInput.AddClass("form-control")
	loginInput.Set("id", "signin")
	passSpan := htmlelement.NewSpan("Password")
	passSpan.AddClass("input-group-text")
	passInput := htmlelement.NewInput("password")
	passInput.Set("id", "password")
	passInput.AddClass("form-control")
	btn := htmlelement.NewButton()
	btn.AddClass("btn")
	btn.AddClass("btn-outline-#orange-600")
	btn.Set("innerHTML", "SIGN IN")
	var signInFunc js.Func
	signInFunc = js.FuncOf(s.SignInFunc)
	btn.AddClickEventListener(&signInFunc)
	loginDiv.AddChild(loginSpan)
	loginDiv.AddChild(loginInput)
	passDiv.AddChild(passSpan)
	passDiv.AddChild(passInput)
	container.AddChild(loginDiv)
	container.AddChild(passDiv)
	container.AddChild(btn)
	return container

}

func (s Signin) SignInFunc(this js.Value, args []js.Value) any {
	var (
		signIn types.SignIn
		Update types.Update
	)
	Update.Type = "sign_in"

	login := htmlelement.GetDocument().Call("getElementById", "signin")
	vLogin := login.Get("value")
	fmt.Println(vLogin.String())
	r := htmlelement.GetInputValue("password")
	fmt.Println(r)
	signIn.Login = vLogin.String()
	signIn.Password = r
	Update.Data = signIn
	s.Sender <- Update
	return nil
}
