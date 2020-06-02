package content

import (
	"fmt"

	"github.com/bosssauce/reference"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Acquirer struct {
	item.Item

	Aquirerid   string   `json:"aquirerid"`
	Name        string   `json:"name"`
	Merchant    []string `json:"merchant"`
	Description string   `json:"description"`
}

// MarshalEditor writes a buffer of html to edit a Acquirer within the CMS
// and implements editor.Editable
func (a *Acquirer) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(a,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Acquirer field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Aquirerid", a, map[string]string{
				"label":       "Aquirerid",
				"type":        "text",
				"placeholder": "Enter the Aquirerid here",
			}),
		},
		editor.Field{
			View: editor.Input("Name", a, map[string]string{
				"label":       "Name",
				"type":        "text",
				"placeholder": "Enter the Name here",
			}),
		},
		editor.Field{
			View: reference.SelectRepeater("Merchant", a, map[string]string{
				"label": "Merchant",
			},
				"Merchant",
				`{{ .name }} `,
			),
		},
		editor.Field{
			View: editor.Richtext("Description", a, map[string]string{
				"label":       "Description",
				"placeholder": "Enter the Description here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Acquirer editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Acquirer"] = func() interface{} { return new(Acquirer) }
}

// String defines how a Acquirer is printed. Update it using more descriptive
// fields from the Acquirer struct type
func (a *Acquirer) String() string {
	return fmt.Sprintf("Acquirer: %s", a.Name)
}
