package doc

import (
	"errors"
	"fmt"

	"github.com/gomutex/godocx/wml/docxpara"
)

// Return a heading paragraph newly added to the end of the document.
// The heading paragraph will contain text and have its paragraph style determined by level.
// If level is 0, the style is set to Title.
// The style is set to Heading {level}.
// if level is outside the range 0-9, error will be returned
func (rd *RootDoc) AddHeading(text string, level uint) (*docxpara.Paragraph, error) {
	if level < 0 || level > 9 {
		return nil, errors.New("Heading level not supported")
	}

	p := &docxpara.Paragraph{
		Children: []*docxpara.ParagraphChild{},
	}
	p.Property = docxpara.DefaultParaProperty()

	style := "Title"
	if level != 0 {
		style = fmt.Sprintf("Heading%d", level)
	}

	p.Property.Style = docxpara.NewParagraphStyle(style)

	bodyElem := DocumentChild{
		Para: p,
	}
	rd.Document.Body.Children = append(rd.Document.Body.Children, bodyElem)

	p.AddText(text)
	return p, nil
}
