package app

import (
	"github.com/twgh/xcgui/drawx"
)

// NewDraw 绘制_创建, 创建图形绘制模块实例, 返回句柄.
//
// hWindow: 窗口句柄.
func NewDraw(hWindow int) *drawx.Draw {
	return drawx.New(hWindow)
}

// NewDrawGDI 绘制_创建GDI, 创建图形绘制模块实例, 返回图形绘制模块实例句柄.
//
// hWindow: 窗口句柄.
//
// hdc: hdc句柄.
func NewDrawGDI(hWindow int, hdc uintptr) *drawx.Draw {
	return drawx.NewGDI(hWindow, hdc)
}

// NewDrawByHandle 从图形绘制模块实例句柄创建对象.
func NewDrawByHandle(handle int) *drawx.Draw {
	return drawx.NewByHandle(handle)
}
