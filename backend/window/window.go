package window

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Data struct {
	DataType             string
	DataText             string
	DataTitle            string
	DataIcon             string
	TextAlign            string
	DataHtml             string
	Width                int
	Height               int
	AutoCloseWindowTimer int
	ShowClose            int
}

type Window struct {
	ctx  context.Context
	Data *Data
}

// NewApp creates a new App application struct
func NewWindow() *Window {
	// Window.ctx = app.Context()
	return &Window{}
}

func (a *Window) SetContext(ctx context.Context, data *Data) {
	a.ctx = ctx
	a.Data = data
}

// 显示窗口
func (w *Window) ShowWindow() {
	// runtime.Show(w.ctx)
	fmt.Println("设置显示窗口", w.ctx)
	runtime.WindowShow(w.ctx)
}

// 隐藏窗口
func (w *Window) HideWindow() {
	runtime.WindowHide(w.ctx)
}

// 设置窗口位置
func (w *Window) GetWindow() *Window {
	return w
}
