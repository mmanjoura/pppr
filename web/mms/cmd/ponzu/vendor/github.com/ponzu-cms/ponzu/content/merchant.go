package content

import (
	"fmt"

	"github.com/bosssauce/reference"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Merchant struct {
	item.Item

	Name        string   `json:"name"`
	PriceOption string   `json:"price_option"`
	Address     []string `json:"address"`
	Mid         []string `json:"mid"`
	Tid         []string `json:"tid"`
	Description string   `json:"description"`
}

// MarshalEditor writes a buffer of html to edit a Merchant within the CMS
// and implements editor.Editable
func (m *Merchant) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(m,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Merchant field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Name", m, map[string]string{
				"label":       "Name",
				"type":        "text",
				"placeholder": "Enter the Name here",
			}),
		},
		editor.Field{
			View: reference.Select("PriceOption", m, map[string]string{
				"label": "PriceOption",
			},
				"PriceOption",
				`{{ .name }} `,
			),
		},
		editor.Field{
			View: editor.InputRepeater("Address", m, map[string]string{
				"label":       "Address",
				"type":        "text",
				"placeholder": "Enter the Address here",
			}),
		},
		editor.Field{
			View: editor.InputRepeater("Mid", m, map[string]string{
				"label":       "Mid",
				"type":        "text",
				"placeholder": "Enter the Mid here",
			}),
		},
		editor.Field{
			View: editor.InputRepeater("Tid", m, map[string]string{
				"label":       "Tid",
				"type":        "text",
				"placeholder": "Enter the Tid here",
			}),
		},
		editor.Field{
			View: editor.Richtext("Description", m, map[string]string{
				"label":       "Description",
				"placeholder": "Enter the Description here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Merchant editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Merchant"] = func() interface{} { return new(Merchant) }
}

// String defines how a Merchant is printed. Update it using more descriptive
// fields from the Merchant struct type
func (m *Merchant) String() string {
	return fmt.Sprintf("Merchant: %s", m.Name)
}
