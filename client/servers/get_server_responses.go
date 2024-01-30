package servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/spf13/viper"

	apierrors "github.com/latitudesh/lsh/internal/api/errors"
	"github.com/latitudesh/lsh/internal/output"
	"github.com/latitudesh/lsh/internal/output/table"
	"github.com/latitudesh/lsh/models"
)

// GetServerReader is a Reader for the GetServer structure.
type GetServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := apierrors.NewUnauthorized()
		if err := result.ReadResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := apierrors.NewNotFound()
		if err := result.ReadResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /servers/{server_id}] get-server", response, response.Code())
	}
}

// NewGetServerOK creates a GetServerOK with default headers values
func NewGetServerOK() *GetServerOK {
	return &GetServerOK{}
}

/*
GetServerOK describes a response with status code 200, with default header values.

Success
*/
type GetServerOK struct {
	Payload *models.Server
}

// IsSuccess returns true when this get server o k response has a 2xx status code
func (o *GetServerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get server o k response has a 3xx status code
func (o *GetServerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get server o k response has a 4xx status code
func (o *GetServerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get server o k response has a 5xx status code
func (o *GetServerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get server o k response a status code equal to that given
func (o *GetServerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get server o k response
func (o *GetServerOK) Code() int {
	return 200
}

func (o *GetServerOK) Error() string {
	return fmt.Sprintf("[GET /servers/{server_id}][%d] getServerOK  %+v", 200, o.Payload)
}

func (o *GetServerOK) String() string {
	return fmt.Sprintf("[GET /servers/{server_id}][%d] getServerOK  %+v", 200, o.Payload)
}

func (o *GetServerOK) GetPayload() *models.Server {
	return o.Payload
}

type GetServerTableRow struct {
	ID              string `json:"id,omitempty"`
	Hostname        string `json:"hostname,omitempty"`
	PrimaryIPV4     string `json:"primary_ipv4,omitempty"`
	Status          string `json:"status,omitempty"`
	IPMIStatus      string `json:"ipmi_status,omitempty"`
	Location        string `json:"location,omitempty"`
	Project         string `json:"project,omitempty"`
	Team            string `json:"team,omitempty"`
	Plan            string `json:"plan,omitempty"`
	OperatingSystem string `json:"operating_system,omitempty"`
	CPU             string `json:"cpu,omitempty"`
	Disk            string `json:"disk,omitempty"`
	RAM             string `json:"ram,omitempty"`
}

func (o *GetServerOK) Render() {
	formatAsJSON := viper.GetBool("json")

	if formatAsJSON {
		o.RenderJSON()
		return
	}

	formatOutputFlag := viper.GetString("output")

	switch formatOutputFlag {
	case "json":
		o.RenderJSON()
	case "table":
		o.RenderTable()
	default:
		fmt.Println("Unsupported output format")
	}
}

func (o *GetServerOK) RenderJSON() {
	if !swag.IsZero(o) && !swag.IsZero(o.Payload) {
		JSONString, err := json.Marshal(o.Payload)
		if err != nil {
			fmt.Println("Could not decode the result as JSON.")
		}

		output.RenderJSON(JSONString)
	}
}

func (o *GetServerOK) RenderTable() {
	resource := o.Payload.Data

	var rows []GetServerTableRow

	attributes := resource.Attributes

	row := GetServerTableRow{
		ID:              table.RenderString(resource.ID),
		Hostname:        table.RenderString(attributes.Hostname),
		PrimaryIPV4:     table.RenderString(*attributes.PrimaryIPV4),
		Status:          table.RenderString(attributes.Status),
		IPMIStatus:      table.RenderString(attributes.IpmiStatus),
		Location:        table.RenderString(attributes.Region.Site.Slug),
		Project:         table.RenderString(attributes.Project.Name),
		Team:            table.RenderString(attributes.Team.Name),
		Plan:            table.RenderString(attributes.Plan.Name),
		OperatingSystem: table.RenderString(attributes.OperatingSystem.Slug),
		CPU:             table.RenderString(attributes.Specs.CPU),
		Disk:            table.RenderString(attributes.Specs.Disk),
		RAM:             table.RenderString(attributes.Specs.RAM),
	}
	rows = append(rows, row)

	headers := table.ExtractHeaders(rows[0])

	var values [][]string

	for _, row := range rows {
		var tr []string

		for _, key := range headers {
			value, err := table.GetFieldValue(row, key)
			if err != nil {
				fmt.Printf("Error accessing field %s: %v\n", key, err)
				continue
			}

			tr = append(tr, fmt.Sprintf("%v", value))
		}

		values = append(values, tr)
	}

	table.Render(table.Table{Headers: headers, Rows: values})
}

func (o *GetServerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Server)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
