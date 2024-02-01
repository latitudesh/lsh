package tablerows

import (
	"github.com/latitudesh/lsh/internal/output/table"
	"github.com/latitudesh/lsh/models"
)

func NewSSHKeyRow(SSHKey *models.SSHKeyData) table.Row {
	attr := SSHKey.Attributes

	return table.Row{
		"id": table.Cell{
			Label: "ID",
			Value: table.String(SSHKey.ID),
		},
		"name": table.Cell{
			Label: "Name",
			Value: table.String(attr.Name),
		},
		"user": table.Cell{
			Label: "User",
			Value: table.String(attr.User.Email),
		},
		"public_key": table.Cell{
			Label:     "Public Key",
			Value:     table.String(attr.PublicKey),
			MaxLength: 25,
		},
		"fingerprint": table.Cell{
			Label: "Fingerprint",
			Value: table.String(attr.Fingerprint),
		},
	}
}
