package icon

import (
	"encoding/xml"
	"fmt"
	"os"
)

type SvgXML struct {
	Xmlns   xml.Attr `xml:"xmlns,attr"`
	ViewBox xml.Attr `xml:"viewBox,attr"`
	Path    PathXML  `xml:"path"`
}

type PathXML struct {
	Data xml.Attr `xml:"d,attr"`
}

func CreateSvgReactElement(filepath string, color string) (string, error) {
	b, err := os.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("failed to read svg image '%s': %v", filepath, err)
	}

	svgElt := SvgXML{}
	if err := xml.Unmarshal(b, &svgElt); err != nil {
		return "", fmt.Errorf("failed to parse icon svg: %v", err)
	}

	// Construct an SVG React Element.
	path_elt_js := fmt.Sprintf(`
		SP_REACT.createElement(
			'path',
			{
				d: "%s"
			},
			null
		)`,
		svgElt.Path.Data.Value,
	)

	svg_elt_js := fmt.Sprintf(`
		SP_REACT.createElement(
			'svg',
			{
				xmlns: "%s",
				viewBox: "%s",
				style: { fill: "%s" }
			},
			%s
		)`,
		svgElt.Xmlns.Value,
		svgElt.ViewBox.Value,
		color,
		path_elt_js,
	)

	return svg_elt_js, nil
}
