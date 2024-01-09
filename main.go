package main

import (
	"context"
	"embed"
	"flag"
	"fmt"

	"desktop-notice/backend"
	"desktop-notice/backend/test"
	"desktop-notice/backend/window"

	"github.com/getlantern/systray"
	"github.com/lxn/win"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS
var title string
var DataType string
var DataText string

type Data struct {
	DataType string
	DataText string
}

func init() {
	// 定义命令行参数
	flag.StringVar(&title, "title", "自动化组通知", "帮助信息")
	flag.StringVar(&DataText, "dataText", "默认内容", "通知信息内容")
	flag.StringVar(&DataType, "dataType", "text", "通知信息类型")
}

func main() {
	// 在Wails初始化之前解析命令行参数
	flag.Parse()
	// 打印接收到的参数值
	fmt.Printf("标题参数: %s\n", title)

	// Create an instance of the app structure
	app := NewApp()
	testApp := test.NewTest()
	_window := window.NewWindow()
	_backend := backend.NewBackend()
	width := 336
	height := 200

	data := &window.Data{
		DataType: "exampleType",
		DataText: "exampleText",
	}
	// go systray.Run(trayReady, trayExit)
	// 配置 https://wails.io/zh-Hans/docs/reference/options/
	err := wails.Run(&options.App{
		Title:     title,
		Width:     width,
		Height:    height,
		MinWidth:  width,
		MinHeight: height,
		MaxWidth:  width,
		MaxHeight: height,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		DisableResize:     true,  // 禁止调整窗口大小
		Frameless:         false, // 窗口无边框
		StartHidden:       false, // 窗口是否隐藏
		AlwaysOnTop:       true,  // 窗口固定在最顶层
		HideWindowOnClose: true,  // 窗口关闭时隐藏窗口
		// BackgroundColour:  &options.RGBA{R: 255, G: 0, B: 0, A: 128},
		Windows: &windows.Options{
			WebviewIsTransparent: true,            // Webview 透明
			WindowIsTranslucent:  true,            // 窗口半透
			BackdropType:         windows.Acrylic, // 窗口背景
		},
		OnStartup: func(ctx context.Context) {
			app.SetContext(ctx)
			_window.SetContext(ctx, data)
			_backend.InitServe(ctx)
			// 获取当前窗口句柄
			handle := win.GetForegroundWindow()
			// 获取当前窗口的长样式
			style := win.GetWindowLong(handle, win.GWL_EXSTYLE)
			// 设置窗口样式为工具窗口（Tool window），这将使其不显示在任务栏上
			win.SetWindowLong(handle, win.GWL_EXSTYLE, style|win.WS_EX_TOOLWINDOW)
		}, // 建窗口并即将开始加载前端资源时的回调
		OnShutdown: app.shutdown, //  应用程序即将退出时的回调
		Bind: []interface{}{ // 向前端暴露的一部分结构体实例
			app,
			testApp,
			_window,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

func trayReady() {
	// 在这里初始化系统托盘图标
	systray.SetIcon(getIconData())

	// 添加退出菜单项
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()
}

func trayExit() {
	// 清理工作，如果有的话
}
func getIconData() []byte {
	// 加载图标数据
	// 这里返回你的图标字节数据
	return nil
}
