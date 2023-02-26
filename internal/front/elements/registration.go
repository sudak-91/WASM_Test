package elements

import (
	"crypto/sha512"
	"fmt"
	"log"
	"syscall/js"
	"time"

	update_types "github.com/sudak-91/wasm-test/pkg/const"
	"github.com/sudak-91/wasm-test/pkg/htmlelement"
	"github.com/sudak-91/wasm-test/pkg/repository"
	pubupdater "github.com/sudak-91/wasm-test/pkg/updater"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Registration struct {
	Sender chan pubupdater.Update
}

func NewRegistrtion(c chan pubupdater.Update) Registration {
	var r Registration
	r.Sender = c
	return r
}

func (r Registration) CreateRegistrationForm(parent htmlelement.Parent) *htmlelement.Div {
	container := htmlelement.NewDiv(parent, "regContainer")
	container.AddClass("container")
	loginDiv := htmlelement.NewDiv(container, "loginDiv")
	loginDiv.AddClass("input-group")
	loginDiv.AddClass("mb-3")
	loginSpan := htmlelement.NewSpan(loginDiv, "loginSpan", "Login")
	loginSpan.AddClass("input-group-text")
	loginInput := htmlelement.NewInput(loginDiv, "loginInput", "text")
	loginInput.AddClass("form-control")

	passDiv := htmlelement.NewDiv(container, "passDiv")
	passDiv.AddClass("input-group")
	passDiv.AddClass("mb-3")
	passSpan := htmlelement.NewSpan(passDiv, "passSpan", "Password")
	passSpan.AddClass("input-group-text")
	passInput := htmlelement.NewInput(passDiv, "passInput", "password")
	passInput.AddClass("form-control")

	emailDiv := htmlelement.NewDiv(container, "emailDiv")
	emailDiv.AddClass("input-group")
	emailDiv.AddClass("mb-3")
	emailSpan := htmlelement.NewSpan(emailDiv, "emailSpan", "E-Mail")
	emailSpan.AddClass("input-group-text")
	emailInput := htmlelement.NewInput(emailDiv, "emailInput", "email")
	emailInput.AddClass("form-control")

	return container

}

func (r Registration) RegFunc(this js.Value, args []js.Value) any {
	var (
		user   repository.User
		Update pubupdater.Update
	)
	Update.Type = update_types.Registration
	login := htmlelement.GetDocument().Call("getElementById", "loginInput")
	vLogin := login.Get("value")
	vPassword := htmlelement.GetInputValue("passInput")
	vEmail := htmlelement.GetInputValue("emailInput")
	if len(vLogin.String()) == 0 {
		fmt.Println("EmptyLogin")
		return nil
	}
	user.ID = primitive.NewObjectID()
	user.Login = fmt.Sprintf("%x", sha512.Sum512([]byte(vLogin.String())))
	user.Password = fmt.Sprintf("%x", sha512.Sum512([]byte(vPassword)))
	user.Email = vEmail
	user.RegistrationDate = primitive.NewDateTimeFromTime(time.Now())
	user.IsTemporary = true
	user.Role = 0
	Update.Data = user
	log.Println(Update)
	r.Sender <- Update
	return nil
}
