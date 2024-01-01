package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func TailString(str string, digit int) string {
	if digit > len(str) {
		digit = len(str)
	}
	return "..." + str[len(str)-digit:]
}

func ShowErrorMessage(window fyne.Window, err error) {
	showToastWithFade(window, "❌ "+err.Error(), time.Second)
}

func showToastWithFade(win fyne.Window, message string, duration time.Duration) {
	text := canvas.NewText(message, color.White)
	text.Alignment = fyne.TextAlignCenter
	toastContainer := container.NewStack(text)
	toastContainer.Resize(fyne.NewSize(win.Canvas().Size().Width, text.MinSize().Height))
	toastContainer.Move(fyne.NewPos(0, win.Canvas().Size().Height-text.MinSize().Height))
	win.Canvas().Overlays().Add(toastContainer)
	fadeIn(text, 500*time.Millisecond)
	time.AfterFunc(duration, func() {
		fadeOut(text, 500*time.Millisecond, func() {
			win.Canvas().Overlays().Remove(toastContainer)
			win.Canvas().Refresh(toastContainer)
		})
	})
}

func fadeIn(obj *canvas.Text, duration time.Duration) {
	step := 10
	for i := 0; i <= step; i++ {
		alpha := float64(i) / float64(step)
		time.AfterFunc(time.Duration(int64(i)*int64(duration)/int64(step)), func() {
			obj.Color = color.NRGBA{R: 255, G: 255, B: 255, A: uint8(alpha * 255)}
			obj.Refresh()
		})
	}
}

func fadeOut(obj *canvas.Text, duration time.Duration, onCompleted func()) {
	step := 10
	for i := step; i >= 0; i-- {
		alpha := float64(i) / float64(step)
		time.AfterFunc(time.Duration(int64(step-i)*int64(duration)/int64(step)), func() {
			obj.Color = color.NRGBA{R: 255, G: 255, B: 255, A: uint8(alpha * 255)}
			obj.Refresh()

			if alpha == 0 {
				onCompleted()
			}
		})
	}
}
