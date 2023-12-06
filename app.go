package main

import (
	"context"
	"fmt"
	"golang-react/utils"

	"github.com/wailsapp/wails/v2/pkg/menu"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) RandomGrandLotto(count int) [][]int {
	randomNum := utils.RandomGrandLotto(count)
	return randomNum
	// return fmt.Sprintf("RandomGrandLotto!", name)
}

func (a *App) Test(name string) string {
	return fmt.Sprintf("Test!", name)
}

func (a *App) GetNewApp() *menu.Menu {
	// 在 NewMenu 函数中一次性初始化 Menu 实例的 Items 切片
	return &menu.Menu{
		Items: []*menu.MenuItem{
			{
				Label: "Item 1",
				Click: func(cd *menu.CallbackData) {
					fmt.Println("Item 1 clicked!")
				},
				// 设置其他字段的值
			},
			{
				Label: "Item 2",
				Click: func(cd *menu.CallbackData) {
					fmt.Println("Item 2 clicked!")
				},
				// 设置其他字段的值
			},
			// 添加更多 MenuItem 到 Items 切片中
		},
	}
}
