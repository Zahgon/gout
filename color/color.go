package color

import (
	"os"
	"strings"

	"github.com/mattn/go-isatty"
)

var (
	// NoColor 关闭颜色开关 本行代码来自github.com/fatih/color, 感谢fatih的付出
	NoColor = os.Getenv("TERM") == "dumb" ||
		(!isatty.IsTerminal(os.Stdout.Fd()) && !isatty.IsCygwinTerminal(os.Stdout.Fd()))
)

type attr int

const (
	// FgBlack 黑色
	FgBlack attr = iota + 30
	// FgRed 红色
	FgRed
	// FgGreen 绿色
	FgGreen
	// FgYellow 黄色
	FgYellow
	// FgBlue 蓝色
	FgBlue
	// FgMagenta 品红
	FgMagenta
	// FgCyan 青色
	FgCyan
	// FgWhite 白色
	FgWhite
)

const (
	// Purple 紫色
	Purple = 35
	// Blue 蓝色
	Blue = 34
)

// Color 着色核心数据结构
type Color struct {
	openColor bool
	attr      attr
}

// New 着色模块构造函数
func New(openColor bool, c ...attr) *Color { _ = "STUB: not implemented"; return nil }

func (c *Color) set(buf *strings.Builder, attr attr) { _ = "STUB: not implemented"; return }

func (c *Color) unset(buf *strings.Builder) { _ = "STUB: not implemented"; return }

func (c *Color) color(a ...interface{}) string { _ = "STUB: not implemented"; return "" }

func (c *Color) colorf(format string, a ...interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

// Sbluef 蓝色函数
func (c *Color) Sbluef(format string, a ...interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

// Sblue 蓝色函数
func (c *Color) Sblue(a ...interface{}) string { _ = "STUB: not implemented"; return "" }

// Spurplef 紫色函数, TODO删除该函数
func (c *Color) Spurplef(format string, a ...interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

// Spurple 紫色函数
func (c *Color) Spurple(a ...interface{}) string { _ = "STUB: not implemented"; return "" }
