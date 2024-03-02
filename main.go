package main

import (
	"context"
	"embed"
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"desktop-notice/backend/test"
	"desktop-notice/backend/utils"
	"desktop-notice/backend/window"

	// "desktop-notice/backend/wails/v2"

	"github.com/getlantern/systray"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS
var DataTitle string
var DataType string
var DataHtml string
var DataText string // text link img
var DataIcon string
var TextAlign string
var Width int
var Height int
var AutoCloseWindowTimer int // 自动关闭窗口的时间 -1 表示不关闭，单位 毫秒

// const defaultHtml = `<div>我是div<button id="button" onclick="window.DesktopNotice.request(http://localhost:3000,{ data: { name: 1 } })">按钮</button>
// <div onClick="window.DesktopNotice.request(http://localhost:3000/postTest,{method: 'post',data: {name: '工具人' }} )">post请求</div>
// <div onClick="window.DesktopNotice.myGlobalFunction()">点击div</div>
// </div>`

// const defaultHtml = `E:\ZHIXIN\automation_auto\DesktopNotice\test\index.html`
const defaultHtml = ``

func init() {
	envVar := os.Getenv("ENV_VAR")
	// 定义命令行参数
	flag.StringVar(&DataTitle, "title", "自动化组通知", "标题")
	if envVar == "dev" {
		flag.StringVar(&DataHtml, "html", defaultHtml, "通知信息内容")
		flag.StringVar(&DataText, "text", "[{\"type\":\"text\",\"text\":\"最新通知，链接如下：最新通知，链接如下：最新通知，链接如下：\"},{\"type\":\"row\",\"text\":\"\"},{\"type\":\"link\",\"text\":\"www.baidu.com\"},{\"type\":\"row\",\"text\":\"\"},{\"type\":\"img\",\"text\":\"https://cdn.wwads.cn/creatives/m88Dv8ffgDW2NO9TVOfe2Ee3QYRtwORH2acMe3Id.png\",\"style\":{\"width\":\"120px\",\"height\":\"120px\"}}]", "通知信息内容")
	} else {
		flag.StringVar(&DataHtml, "html", "", "通知信息内容")
		flag.StringVar(&DataText, "text", "", "通知信息内容")
	}
	flag.StringVar(&DataType, "type", "text", "通知信息类型")
	flag.StringVar(&DataIcon, "icon", "default", "图标")
	flag.StringVar(&TextAlign, "textAlign", "start", "内容剧中")
	flag.IntVar(&Width, "width", 336, "宽度")
	flag.IntVar(&Height, "height", 200, "高度")
	flag.IntVar(&AutoCloseWindowTimer, "autoCloseWindowTimer", 5000, "自动关闭时间")

}

func main() {

	// 在Wails初始化之前解析命令行参数
	flag.Parse()

	// Create an instance of the app structure
	_app := NewApp()
	_testApp := test.NewTest()
	_window := window.NewWindow()

	width := Width
	height := Height

	// 如果图片路径不是http开头的转base64
	if !strings.HasPrefix(DataIcon, "http") && (DataIcon != "default") {
		res, err := GetBase64FromImagePath(DataIcon)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Base64编码:", res)
		DataIcon = "data:image/png;base64," + res
	}

	// 如果html是以html结尾的转为html字符串
	if strings.HasSuffix(DataHtml, ".html") {
		res, err := utils.HtmlString(DataHtml)
		if err != nil {
			fmt.Println("Error:", err)
		}
		DataHtml = res
	}

	data := &window.Data{
		DataTitle:            DataTitle,
		DataType:             DataType,
		DataText:             DataText,
		DataHtml:             DataHtml,
		DataIcon:             DataIcon,
		TextAlign:            TextAlign,
		Width:                width,
		Height:               height,
		AutoCloseWindowTimer: AutoCloseWindowTimer,
	}
	// go systray.Run(trayReady, trayExit)
	// 配置 https://wails.io/zh-Hans/docs/reference/options/
	err := wails.Run(&options.App{
		Title:     DataTitle,
		Width:     width,
		Height:    height,
		MinWidth:  width,
		MinHeight: height,
		MaxWidth:  width,
		MaxHeight: height,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		DisableResize:     false, // 禁止调整窗口大小
		Frameless:         true,  // 窗口无边框
		StartHidden:       false, // 窗口是否隐藏
		AlwaysOnTop:       true,  // 窗口固定在最顶层
		HideWindowOnClose: true,  // 窗口关闭时隐藏窗口
		// BackgroundColour:  "transparent", // &options.RGBA{R: 255, G: 255, B: 255, A: 100},

		Windows: &windows.Options{
			WebviewIsTransparent: true,            // Webview 透明
			WindowIsTranslucent:  true,            // 窗口半透
			BackdropType:         windows.Acrylic, // 窗口背景
			// DisableFramelessWindowDecorations: false,
			// DisableWindowIcon:                 true, // 根据需要禁用窗口图标
		},
		OnStartup: func(ctx context.Context) {
			_app.SetContext(ctx)
			_window.SetContext(ctx, data)
			// 获取当前窗口句柄
			// handle := win.GetForegroundWindow()
			// // 获取当前窗口的长样式
			// style := win.GetWindowLong(handle, win.GWL_EXSTYLE)
			// // 设置窗口样式为工具窗口（Tool window），这将使其不显示在任务栏上
			// win.SetWindowLong(handle, win.GWL_EXSTYLE, style|win.WS_EX_TOOLWINDOW)
		}, // 建窗口并即将开始加载前端资源时的回调
		OnShutdown: _app.shutdown, //  应用程序即将退出时的回调
		Bind: []interface{}{ // 向前端暴露的一部分结构体实例
			_app,
			_testApp,
			_window,
			data,
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

func GetBase64FromImagePath(imagePath string) (string, error) {
	// 读取图标文件
	imageFile, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return "", err
	}

	// 将图标文件转换为base64编码
	base64Encoded := base64.StdEncoding.EncodeToString(imageFile)
	return base64Encoded, nil
}
