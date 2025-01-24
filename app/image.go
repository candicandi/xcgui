package app

import (
	"github.com/twgh/xcgui/imagex"
)

// 图片_加载从图片源.
//
// hImageSrc: 图片源句柄.
func NewImageBySrc(hImageSrc int) *imagex.Image {
	return imagex.NewBySrc(hImageSrc)
}

// 图片_加载从文件.
//
// pFileName: 图片文件.
func NewImageByFile(pFileName string) *imagex.Image {
	return imagex.NewByFile(pFileName)
}

// 图片_加载从文件自适应, 加载图片从文件, 自适应图片.
//
// pFileName: 图片文件.
//
// leftSize: 坐标.
//
// topSize: 坐标.
//
// rightSize: 坐标.
//
// bottomSize: 坐标.
func NewImageByFileAdaptive(pFileName string, leftSize, topSize, rightSize, bottomSize int32) *imagex.Image {
	return imagex.NewByFileAdaptive(pFileName, leftSize, topSize, rightSize, bottomSize)
}

// 图片_加载从文件指定区域, 加载图片, 指定区位置及大小.
//
// pFileName: 图片文件.
//
// x: 坐标.
//
// y: 坐标.
//
// cx: 宽度.
//
// cy: 高度.
func NewImageByFileRect(pFileName string, x, y, cx, cy int32) *imagex.Image {
	return imagex.NewByFileRect(pFileName, x, y, cx, cy)
}

// 图片_加载从资源自适应, 加载图片从资源, 自适应图片.
//
// id: 资源ID.
//
// pType: 资源类型.
//
// leftSize: 坐标.
//
// topSize: 坐标.
//
// rightSize: 坐标.
//
// bottomSize: 坐标.
//
// hModule:	从指定模块加载, 例如:DLL, EXE; 如果为空, 从当前EXE加载.
func NewImageByResAdaptive(id int32, pType string, leftSize, topSize, rightSize, bottomSize int32, hModule uintptr) *imagex.Image {
	return imagex.NewByResAdaptive(id, pType, leftSize, topSize, rightSize, bottomSize, hModule)
}

// 图片_加载从资源.
//
// id: 资源ID.
//
// pType: 资源类型.
//
// bStretch: 是否拉伸图片.
//
// hModule:	从指定模块加载, 例如:DLL, EXE; 如果为空, 从当前EXE加载.
func NewImageByRes(id int32, pType string, bStretch bool, hModule uintptr) *imagex.Image {
	return imagex.NewByRes(id, pType, bStretch, hModule)
}

// 图片_加载从ZIP, 加载图片从ZIP压缩包.
//
// pZipFileName: ZIP压缩包文件名.
//
// pFileName: 图片文件名.
//
// pPassword: ZIP压缩包密码.
func NewImageByZip(pZipFileName string, pFileName string, pPassword string) *imagex.Image {
	return imagex.NewByZip(pZipFileName, pFileName, pPassword)
}

// 图片_加载从资源ZIP.
//
// id: RC资源ID.
//
// pFileName: 图片文件名.
//
// pPassword: ZIP压缩包密码.
//
// hModule: 模块句柄, 可填0.
func NewImageByZipRes(id int32, pFileName string, pPassword string, hModule uintptr) *imagex.Image {
	return imagex.NewByZipRes(id, pFileName, pPassword, hModule)
}

// 图片_加载从ZIP自适应, 加载图片从ZIP压缩包, 自适应图片.
//
// pZipFileName: ZIP压缩包文件名.
//
// pFileName: 图片文件名.
//
// pPassword: ZIP压缩包密码.
//
// x1: 坐标.
//
// x2: 坐标.
//
// y1: 坐标.
//
// y2: 坐标.
func NewImageByZipAdaptive(pZipFileName string, pFileName string, pPassword string, x1, x2, y1, y2 int32) *imagex.Image {
	return imagex.NewByZipAdaptive(pZipFileName, pFileName, pPassword, x1, x2, y1, y2)
}

// 图片_加载从ZIP指定区域, 加载ZIP图片, 指定区位置及大小.
//
// pZipFileName: ZIP文件.
//
// pFileName: 图片名称.
//
// pPassword: 密码.
//
// x: 坐标.
//
// y: 坐标.
//
// cx: 宽度.
//
// cy: 高度.
func NewImageByZipRect(pZipFileName string, pFileName string, pPassword string, x, y, cx, cy int32) *imagex.Image {
	return imagex.NewByZipRect(pZipFileName, pFileName, pPassword, x, y, cx, cy)
}

// 图片_加载从内存ZIP.
//
// data: 图片数据.
//
// pFileName: 图片名称.
//
// pPassword: zip压缩包密码.
func NewImageByZipMem(data []byte, pFileName string, pPassword string) *imagex.Image {
	return imagex.NewByZipMem(data, pFileName, pPassword)
}

// 图片_加载从内存, 加载流图片.
//
// pBuffer: 图片数据.
func NewImageByMem(pBuffer []byte) *imagex.Image {
	return imagex.NewByMem(pBuffer)
}

// 图片_加载从内存指定区域, 加载流图片, 指定区位置及大小.
//
// pBuffer: 图片数据.
//
// x: 坐标.
//
// y: 坐标.
//
// cx: 宽度.
//
// cy: 高度.
func NewImageByMemRect(pBuffer []byte, x, y, cx, cy int32) *imagex.Image {
	return imagex.NewByMemRect(pBuffer, x, y, cx, cy)
}

// 图片_加载从内存自适应, 加载流图片压缩包, 自适应图片(九宫格).
//
// pBuffer: 图片数据.
//
// leftSize: 坐标.
//
// topSize: 坐标.
//
// rightSize: 坐标.
//
// bottomSize: 坐标.
func NewImageByMemAdaptive(pBuffer []byte, leftSize, topSize, rightSize, bottomSize int32) *imagex.Image {
	return imagex.NewByMemAdaptive(pBuffer, leftSize, topSize, rightSize, bottomSize)
}

// 图片_加载从Image, 加载图片从GDI+的Image对象.
//
// pImage: GDI图片对象指针Image*.
func NewImageByImage(pImage uintptr) *imagex.Image {
	return imagex.NewByImage(pImage)
}

// 图片_加载文件图标, 加载文件图标, 从一个EXE文件或DLL文件或图标文件; 例如:*.exe文件的图标.
//
// pFileName: 文件名.
func NewImageByExtractIcon(pFileName string) *imagex.Image {
	return imagex.NewByExtractIcon(pFileName)
}

// 图片_加载从HICON, 创建一个炫彩图片句柄, 从一个现有的图标句柄HICON.
//
// hIcon: 图标句柄.
func NewImageByHICON(hIcon uintptr) *imagex.Image {
	return imagex.NewByHICON(hIcon)
}

// 图片_加载从HBITMAP, 创建一个炫彩图片句柄, 从一个现有的位图句柄HBITMAP.
//
// hBitmap: 位图句柄.
func NewImageByHBITMAP(hBitmap uintptr) *imagex.Image {
	return imagex.NewByHBITMAP(hBitmap)
}

// 图片_加载从SVG.
//
// hSvg: SVG句柄.
func NewImageBySvg(hSvg int) *imagex.Image {
	return imagex.NewBySvg(hSvg)
}

// 图片_加载从SVG文件.
//
// pFileName: 文件名.
func NewImageBySvgFile(pFileName string) *imagex.Image {
	return imagex.NewBySvgFile(pFileName)
}

// 图片_加载从SVG字符串.
//
// pString: 字符串.
func NewImageBySvgString(pString string) *imagex.Image {
	return imagex.NewBySvgString(pString)
}

// 图片_加载从SVG字符串W.
//
// pString: 字符串.
func NewImageBySvgStringW(pString string) *imagex.Image {
	return imagex.NewBySvgStringW(pString)
}

// 图片_加载从SVG字符串UTF8, 更推荐使用 imagex.NewBySvgStringW.
//
// pString: 字符串.
func NewImageBySvgStringUtf8(pString string) *imagex.Image {
	return imagex.NewBySvgStringUtf8(pString)
}

// 从句柄创建对象.
func NewImageByHandle(handle int) *imagex.Image {
	return imagex.NewByHandle(handle)
}

// 根据资源文件中的name创建对象, 失败返回nil.
//
// pName: 资源名称.
func NewImageByName(name string) *imagex.Image {
	return imagex.NewByName(name)
}

// 从指定的资源文件中, 根据name创建对象, 失败返回nil.
//
// pFileName: 资源文件名.
//
// pName: 资源名称.
func NewImageByNameEx(fileName, name string) *imagex.Image {
	return imagex.NewByNameEx(fileName, name)
}
