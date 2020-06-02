package content

import (
	"fmt"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type PriceOption struct {
	item.Item

	Name           string `json:"name"`
	Scheme         string `json:"scheme"`
	PriceType      string `json:"price_type"`
	DomesticMSCPPT string `json:"domestic_m_s_c_p_p_t"`
	EEAMSCRate     string `json:"e_e_a_m_s_c_rate"`
	EEAMSCPPT      string `json:"e_e_a_m_s_c_p_p_t"`
	MSCRate        string `json:"m_s_c_rate"`
	Description    string `json:"description"`
}

// MarshalEditor writes a buffer of html to edit a PriceOption within the CMS
// and implements editor.Editable
func (p *PriceOption) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(p,
		// Take note that the first argument to these Input-like functions
		// is the string version of each PriceOption field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Name", p, map[string]string{
				"label":       "Name",
				"type":        "text",
				"placeholder": "Enter the Name here",
			}),
		},
		editor.Field{
			View: editor.Select("Scheme", p, map[string]string{
				"label": "Scheme",
			}, map[string]string{
				"1": "Visa",
				"2": "Master Card",
			}),
		},
		editor.Field{
			View: editor.Select("PriceType", p, map[string]string{
				"label": "PriceType",
			}, map[string]string{
				"1": "IC++",
				"2": "Extended Blended",
			}),
		},
		editor.Field{
			View: editor.Input("DomesticMSCPPT", p, map[string]string{
				"label":       "DomesticMSCPPT",
				"type":        "text",
				"placeholder": "Enter the DomesticMSCPPT here",
			}),
		},
		editor.Field{
			View: editor.Input("EEAMSCRate", p, map[string]string{
				"label":       "EEAMSCRate",
				"type":        "text",
				"placeholder": "Enter the EEAMSCRate here",
			}),
		},
		editor.Field{
			View: editor.Input("EEAMSCPPT", p, map[string]string{
				"label":       "EEAMSCPPT",
				"type":        "text",
				"placeholder": "Enter the EEAMSCPPT here",
			}),
		},
		editor.Field{
			View: editor.Input("MSCRate", p, map[string]string{
				"label":       "MSCRate",
				"type":        "text",
				"placeholder": "Enter the MSCRate here",
			}),
		},
		editor.Field{
			View: editor.Richtext("Description", p, map[string]string{
				"label":       "Description",
				"placeholder": "Enter the Description here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render PriceOption editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["PriceOption"] = func() interface{} { return new(PriceOption) }
}

// String defines how a PriceOption is printed. Update it using more descriptive
// fields from the PriceOption struct type
func (p *PriceOption) String() string {
	return fmt.Sprintf("PriceOption: %s", p.Name)
}
