package main

import (
	"context"
	"fmt"
	"golang-react/utils"
)

// App struct
type OtherApp struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewOtherApp() *OtherApp {
	return &OtherApp{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *OtherApp) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *OtherApp) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *OtherApp) RandomGrandLotto(count int) [][]int {
	randomNum := utils.RandomGrandLotto(count)
	return randomNum
	// return fmt.Sprintf("RandomGrandLotto!", name)
}

func (a *OtherApp) Test(name string) string {
	return fmt.Sprintf("Test!", name)
}
