package resource

import (
	_ "embed"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type MyTheme struct{}

//go:embed wqy.ttf
var fontSmiley []byte

var _ fyne.Theme = (*MyTheme)(nil)

func (m MyTheme) Font(s fyne.TextStyle) fyne.Resource {
	//return theme.DefaultTheme().Font(s)
	return &fyne.StaticResource{
		StaticName:    "wqy",
		StaticContent: fontSmiley,
	}
}

func (*MyTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(n, v)
}

func (*MyTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (*MyTheme) Size(n fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(n)
}
