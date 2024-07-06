package font

import (
	"github.com/twgh/xcgui/objectbase"
	"github.com/twgh/xcgui/res"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// Font 炫彩字体.
type Font struct {
	objectbase.ObjectBase
}

// New 字体_创建, 创建炫彩字体. 当字体句柄与元素关联后, 会自动释放.
//
//	size: 字体大小,单位(pt, 磅).
func New(size int32) *Font {
	p := &Font{}
	p.SetHandle(xc.XFont_Create(size))
	return p
}

// NewEX 字体_创建扩展. 创建炫彩字体.
//
//	pName: 字体名称.
//
//	size: 字体大小, 单位(pt, 磅).
//
//	style: 字体样式, xcc.FontStyle_ .
func NewEX(pName string, size int32, style xcc.FontStyle_) *Font {
	p := &Font{}
	p.SetHandle(xc.XFont_CreateEx(pName, size, style))
	return p
}

// NewLOGFONTW 字体_创建从LOGFONT. 创建炫彩字体.
//
//	pFontInfo: 字体信息.
func NewByLOGFONTW(pFontInfo *xc.LOGFONTW) *Font {
	p := &Font{}
	p.SetHandle(xc.XFont_CreateFromLOGFONTW(pFontInfo))
	return p
}

// NewByHFONT 字体_创建从HFONT. 创建炫彩字体从现有HFONT字体.
//
//	hFont: 字体句柄.
func NewByHFONT(hFont uintptr) *Font {
	p := &Font{}
	p.SetHandle(xc.XFont_CreateFromHFONT(hFont))
	return p
}

// NewByFont 字体_创建从Font. 创建炫彩字体从GDI+字体.
//
//	pFont: GDI+ 字体指针.
func NewByFont(pFont uintptr) *Font {
	p := &Font{}
	p.SetHandle(xc.XFont_CreateFromFont(pFont))
	return p
}

// NewByFile 字体_创建从文件. 创建字体从文件.
//
//	pFontFile: 字体文件名.
//
//	size: 字体大小, 单位(pt, 磅).
//
//	style: 字体样式, xcc.FontStyle_ .
func NewByFile(pFontFile string, size int32, style xcc.FontStyle_) *Font {
	p := &Font{}
	p.SetHandle(xc.XFont_CreateFromFile(pFontFile, size, style))
	return p
}

// NewByZip 字体_创建从ZIP.
//
//	pZipFileName: zip文件名.
//
//	pFileName: 字体文件名.
//
//	pPassword: zip密码.
//
//	fontSize: 字体大小, 单位(pt, 磅).
//
//	style: 字体样式: xcc.FontStyle_ .
func NewByZip(pZipFileName, pFileName, pPassword string, fontSize int32, style xcc.FontStyle_) *Font {
	p := &Font{}
	p.SetHandle(xc.XFont_CreateFromZip(pZipFileName, pFileName, pPassword, fontSize, style))
	return p
}

// NewByZipMem 字体_创建从内存ZIP.
//
//	data: zip数据.
//
//	pFileName: 字体文件名.
//
//	pPassword: zip密码.
//
//	fontSize: 字体大小, 单位(pt, 磅).
//
//	style: 字体样式: xcc.FontStyle_ .
func NewByZipMem(data []byte, pFileName, pPassword string, fontSize int32, style xcc.FontStyle_) *Font {
	p := &Font{}
	p.SetHandle(xc.XFont_CreateFromZipMem(data, pFileName, pPassword, fontSize, style))
	return p
}

// NewByMem 字体_创建从内存. 创建炫彩字体从内存.
//
//	data: 字体文件数据.
//
//	fontSize: 字体大小, 单位(pt, 磅).
//
//	style: 字体样式, xcc.FontStyle_ .
func NewByMem(data []byte, fontSize int32, style xcc.FontStyle_) *Font {
	p := &Font{}
	p.SetHandle(xc.XFont_CreateFromMem(data, fontSize, style))
	return p
}

// NewByRes 字体_创建从资源. 创建字体从资源.
//
//	id: xx.
//
//	pType: xx.
//
//	fontSize: 字体大小, 单位(pt, 磅).
//
//	style: 字体样式, xcc.FontStyle_ .
//
//	hModule: xx.
func NewByRes(id int32, pType string, fontSize int32, style xcc.FontStyle_, hModule uintptr) *Font {
	p := &Font{}
	p.SetHandle(xc.XFont_CreateFromRes(id, pType, fontSize, style, hModule))
	return p
}

// NewByHandle 从句柄创建对象.
//
//	handle: 炫彩字体句柄.
func NewByHandle(handle int) *Font {
	p := &Font{}
	p.SetHandle(handle)
	return p
}

// NewByName 根据资源文件中的name创建对象, 失败返回nil.
//
//	name: name名称.
func NewByName(name string) *Font {
	handle := res.GetFont(name)
	if handle > 0 {
		p := &Font{}
		p.SetHandle(handle)
		return p
	}
	return nil
}

// EnableAutoDestroy 字体_启用自动销毁. 是否自动销毁.
//
//	bEnable: 是否启用.
func (f *Font) EnableAutoDestroy(bEnable bool) *Font {
	xc.XFont_EnableAutoDestroy(f.Handle, bEnable)
	return f
}

// GetFont 字体_取Font. 获取字体. 返回GDI+ Font指针
func (f *Font) GetFont() uintptr {
	return xc.XFont_GetFont(f.Handle)
}

// GetFontInfo 字体_取信息. 获取字体信息.
//
//	pInfo: 接收返回的字体信息.
func (f *Font) GetFontInfo(pInfo *xc.Font_Info_) *Font {
	xc.XFont_GetFontInfo(f.Handle, pInfo)
	return f
}

// GetLOGFONTW 字体_取LOGFONTW. 获取字体LOGFONTW.
//
//	hdc: hdc句柄.
//
//	pOut: 接收返回信息.
func (f *Font) GetLOGFONTW(hdc uintptr, pOut *xc.LOGFONTW) bool {
	return xc.XFont_GetLOGFONTW(f.Handle, hdc, pOut)
}

// Destroy 字体_销毁. 强制销毁炫彩字体, 谨慎使用, 建议使用 Release() 释放.
func (f *Font) Destroy() *Font {
	xc.XFont_Destroy(f.Handle)
	return f
}

// AddRef 字体_增加引用计数.
func (f *Font) AddRef() *Font {
	xc.XFont_AddRef(f.Handle)
	return f
}

// GetRefCount 字体_取引用计数.
func (f *Font) GetRefCount() int32 {
	return xc.XFont_GetRefCount(f.Handle)
}

// Release 字体_释放引用计数. 释放引用计数, 当引用计数为0时自动销毁.
func (f *Font) Release() *Font {
	xc.XFont_Release(f.Handle)
	return f
}

// GetUnderlineEdit 字体_取下划线. 仅供edit字体使用, 因为edit不支持下划线字体, 所以需要单独设置. bUnderline: 返回是否启用下划线. bStrikeout: 返回是否启用删除线.
func (f *Font) GetUnderlineEdit() (bUnderline, bStrikeout bool) {
	return xc.XFont_GetUnderlineEdit(f.Handle)
}

// SetUnderlineEdit 字体_置下划线. 仅供edit字体使用, 因为edit不支持下划线字体, 所以需要单独设置.
//
//	bUnderline: 是否启用下划线.
//
//	bStrikeout: 是否启用删除线.
func (f *Font) SetUnderlineEdit(bUnderline, bStrikeout bool) *Font {
	xc.XFont_SetUnderlineEdit(f.Handle, bUnderline, bStrikeout)
	return f
}
