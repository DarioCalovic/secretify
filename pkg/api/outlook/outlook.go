package outlook

import (
	"bytes"
	_ "embed"
	"text/template"
)

//go:embed assets/manifest.xml
var data []byte

type Data struct {
	AppID string
	UIURL string
}

// Manifest returns manifest xml for office add-in integration
func (o *Outlook) Manifest() []byte {
	tmpl, err := template.New("manifest").Parse(string(data))
	if err != nil {
		panic(err)
	}
	sweaters := Data{
		AppID: o.cfgSvc.Outlook().AppID,
		UIURL: o.cfgSvc.Outlook().UIURL,
	}

	var result bytes.Buffer
	err = tmpl.Execute(&result, sweaters)
	if err != nil {
		panic(err)
	}

	return result.Bytes()
}
