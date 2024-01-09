package backend

import (
	"context"
	"fmt"
	"net/http"

	controllers "desktop-notice/backend/controllers"
)

// App struct
type Backend struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewBackend() *Backend {
	return &Backend{}
}

// 启动一个服务
func (b *Backend) InitServe(ctx context.Context) {

	// 将处理函数与路径关联
	http.HandleFunc("/", controllers.Test)
	// http.HandleFunc("/receive", func(w http.ResponseWriter, r *http.Request) { controllers.Receive(ctx) })
	// http.HandleFunc("/receive", func(w http.ResponseWriter, r *http.Request) { win.ShowWindow() })
	// 启动服务器并监听本地端口
	fmt.Println("Server is running at http://localhost:36000")
	http.ListenAndServe(":36000", nil)
}
