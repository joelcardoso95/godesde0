package users

import (
	"fmt"
	"projects/godesde0/model"
	"time"
)

func CriaUsuario() {
	u := new(model.User)
	u.AddUser(10, "Pablo", time.Now(), true)
	fmt.Println(u)
}
