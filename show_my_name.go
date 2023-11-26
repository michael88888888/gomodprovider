package gomodprovider

import (
	"fmt"

	"github.com/michael88888888/gomodproviderx"
)

func ShowMyName() {
	fmt.Printf("Hi, I’m %v, meeting you so nice.", gomodproviderx.GetMyName())
}