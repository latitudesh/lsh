package renderer

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/latitudesh/lsh/internal/output/table"
)

type ResponseData interface {
	Validate(formats strfmt.Registry) error
	ContextValidate(ctx context.Context, formats strfmt.Registry) error
	MarshalBinary() ([]byte, error)
	UnmarshalBinary(b []byte) error
	TableRow() table.Row
}

type Renderer interface {
	Render(data []ResponseData)
}

type RenderContext struct {
	renderer Renderer
}

func NewRenderContext(renderer Renderer) *RenderContext {
	return &RenderContext{
		renderer: renderer,
	}
}

func (rc *RenderContext) Render(data []ResponseData) {
	rc.renderer.Render(data)
}
