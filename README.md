# README

## About

This is the official Wails React-TS template.

You can configure the project by editing `wails.json`. More information about the project settings can be found
here: https://wails.io/docs/reference/project-config

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.


## 项目说明

这是一个使用wails和react-ts创建的桌面应用程序。

## 运行

在项目目录下运行`wails dev`来运行开发服务器。

在项目目录下运行`wails build -clean`来构建可发布的应用程序。
`wails build -clean`
`wails build -clean -devtools -debug -platform windows/arm64`


wails doctor 检查您是否安装了正确的依赖项

## 项目配置

项目配置文件是`wails.json`。你可以通过编辑这个文件来配置你的项目。
## 项目结构

项目结构如下：

```
├── README.md
├── frontend  前端代码
│   ├── index.html
│   ├── index.tsx
│   ├── main.css
│   └── main.tsx
├── wailsjs golang 编译代码
├── golang golang代码
│   ├── go
│   │   └── main.go

```


desktop-notice.exe --title="标题" --content="放假"