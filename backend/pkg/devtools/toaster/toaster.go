package toaster

import (
	"fmt"

	"steamdeckhomebrew.decktweaks/pkg/devtools/icon"
	"steamdeckhomebrew.decktweaks/pkg/devtools/steamdeck"
)

const (
	ICON_COLOR = "white"
)

type ToastData struct {
	Title    string `json:"title"`
	Body     string `json:"body"`
	Icon     string `json:"icon"`
	Critical bool   `json:"critical"`
}

func NewToast(title string, body string) *ToastData {
	return &ToastData{
		Title:    title,
		Body:     body,
		Icon:     `""`,
		Critical: false,
	}
}

// Sets a given SVG Image as the toast icon, by creating a React Element, which
// the toaster expects.
func (t *ToastData) SetIcon(svgFilepath string) error {
	svg_react_elt_js, err := icon.CreateSvgReactElement(svgFilepath, ICON_COLOR)
	if err != nil {
		return fmt.Errorf("failed to create an svg react element: %v", err)
	}
	t.Icon = svg_react_elt_js
	return nil
}

// Invokes the toast data.
func (t *ToastData) Toast() error {
	if err := steamdeck.InjectJs(fmt.Sprintf(
		`window.DeckyPluginLoader.toaster.toast({
         title: "%s",
         body: "%s",
         icon: %s,
         critical: %v,
      })`,
		t.Title,
		t.Body,
		t.Icon,
		t.Critical,
	)); err != nil {
		return fmt.Errorf("toast failed: %v", err)
	}

	return nil
}
