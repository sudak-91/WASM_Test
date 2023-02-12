package elements

import (
	"crypto/sha512"
	"fmt"
	"log"
	"syscall/js"

	"github.com/sudak-91/wasm-test/internal/pkg/updater"
	update_types "github.com/sudak-91/wasm-test/pkg/const"
	"github.com/sudak-91/wasm-test/pkg/htmlelement"
	"github.com/sudak-91/wasm-test/pkg/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Signin struct {
	p      *htmlelement.Paragraph
	Sender chan updater.Update
}

func NewSignIn(c chan updater.Update) Signin {
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
	btn := htmlelement.NewButton(container, "signInButton", "button", "Логин")
	btn.AddClass("btn")
	btn.AddClass("btn-primary")
	var signInFunc js.Func
	signInFunc = js.FuncOf(s.SignInFunc)
	btn.AddClickEventListener(&signInFunc)
	s.p = htmlelement.NewParagaph(container, "testP", "Look here")
	return container

}

func (s Signin) SignInFunc(this js.Value, args []js.Value) any {
	var (
		user   repository.User
		Update updater.Update
	)
	fmt.Println("Click")
	Update.Type = update_types.LoginUpdater

	login := htmlelement.GetDocument().Call("getElementById", "loginInput")
	vLogin := login.Get("value")
	fmt.Println(vLogin.String())
	r := htmlelement.GetInputValue("password")
	fmt.Println(r)
	user.ID = primitive.NewObjectID()
	user.Login = fmt.Sprintf("%x", sha512.Sum512([]byte(vLogin.String())))
	user.Password = fmt.Sprintf("%x", sha512.Sum512([]byte(r)))
	Update.SignIn = &user
	log.Println(Update)
	s.Sender <- Update

	return nil
}
