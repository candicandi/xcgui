package app

import (
	"github.com/twgh/xcgui/bkmanager"
)

// NewBkManager 背景_创建, 创建背景管理器.
func NewBkManager() *bkmanager.BkManager {
	return bkmanager.New()
}

// NewBkManagerByHandle 从句柄创建背景管理器对象.
func NewBkManagerByHandle(handle int) *bkmanager.BkManager {
	return bkmanager.NewByHandle(handle)
}

// NewBkManagerByName 从name创建背景管理器对象, 失败返回nil.
func NewBkManagerByName(name string) *bkmanager.BkManager {
	return bkmanager.NewByName(name)
}
