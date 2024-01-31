package tablerows

import (
	"github.com/latitudesh/lsh/internal/output/table"
	"github.com/latitudesh/lsh/models"
)

type SSHKeyRow struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	User        string `json:"user,omitempty"`
	PublicKey   string `json:"public_key,omitempty"`
	Fingerprint string `json:"fingerprint,omitempty"`
}

func NewSSHKeyRow(SSHKey *models.SSHKeyData) *SSHKeyRow {
	attr := SSHKey.Attributes

	return &SSHKeyRow{
		ID:          table.String(SSHKey.ID),
		Name:        table.String(attr.Name),
		User:        table.String(attr.User.Email),
		PublicKey:   table.String(attr.PublicKey),
		Fingerprint: table.String(attr.Fingerprint),
	}
}

func CreateSSHKeyRows(SSHKeys []*models.SSHKeyData) []interface{} {
	var rows []SSHKeyRow
	var rowsInterface []interface{}

	for _, SSHKey := range SSHKeys {
		rows = append(rows, *NewSSHKeyRow(SSHKey))
	}

	for _, row := range rows {
		rowsInterface = append(rowsInterface, row)
	}

	return rowsInterface
}
