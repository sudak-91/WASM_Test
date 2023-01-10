package elements

import (
	"fmt"
	"syscall/js"

	"github.com/sudak-91/wasm-test/internal/types"
	"github.com/sudak-91/wasm-test/pkg/htmlelement"
)

type Signin struct {
	p      *htmlelement.Paragraph
	Sender chan types.Update
}

func NewSignIn(c chan types.Update) Signin {
	var s Signin
	s.Sender = c
	return s

}

func (s Signin) CreateSignIn(parent htmlelement.Parent) *htmlelement.Div {
	container := htmlelement.NewDiv(parent, "signInContainer")
	container.AddClass("container")
	loginDiv := htmlelement.NewDiv(container, "loginDiv")
	loginDiv.AddClass("input-group")
	loginDiv.AddClass("mb-3")
	passDiv := htmlelement.NewDiv(container, "passDiv")
	passDiv.AddClass("input-group")
	passDiv.AddClass("mb-3")
	loginSpan := htmlelement.NewSpan(loginDiv, "loginSpan", "Логин")
	loginSpan.AddClass("input-group-text")
	loginInput := htmlelement.NewInput(loginDiv, "loginInput", "text")
	loginInput.AddClass("form-control")
	passSpan := htmlelement.NewSpan(passDiv, "passwordSpan", "Пароль")
	passSpan.AddClass("input-group-text")
	passInput := htmlelement.NewInput(passDiv, "password", "password")
	passInput.AddClass("form-control")
	btn := htmlelement.NewButton(container, "signInButton", "button")
	btn.AddClass("btn")
	btn.AddClass("btn-outline-#orange-600")
	var signInFunc js.Func
	signInFunc = js.FuncOf(s.SignInFunc)
	btn.AddClickEventListener(&signInFunc)
	s.p = htmlelement.NewParagaph(container, "testP", "Look here")
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
	s.p.ChangeText("Formaling")
	s.Sender <- Update

	return nil
}
