package controllers

import (
	"github.com/revel/revel"
	"fmt"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (this App) add(name string) revel.Result {
	fmt.Println("name=",name)
	return this.RenderText("hi,"+name)

}

















