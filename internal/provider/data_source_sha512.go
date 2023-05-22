package provider

import (
	"context"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
)

var _ datasource.DataSource = (*sha512DataSource)(nil)

func NewSha512DataSource() datasource.DataSource {
	return &sha512DataSource{}
}

type sha512DataSource struct{}

func (d *sha512DataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "sha"
}

func (d *sha512DataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},

			"input": schema.StringAttribute{
				Required: true,
			},
			"sha": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (d *sha512DataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var model modelV0
	diags := req.Config.Get(ctx, &model)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	enc, err := getEncoding("ISO-8859-1")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	var input = model.Input.ValueString()
	encodedBytes, err := convertToBytes(input, enc)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	hash := sha512.New()
	hash.Write(encodedBytes)
	model.Sha = types.StringValue(
		hex.EncodeToString(hash.Sum(nil)),
	)
	model.ID = types.StringValue(input)
	diags = resp.State.Set(ctx, model)
	resp.Diagnostics.Append(diags...)
}

type modelV0 struct {
	ID    types.String `tfsdk:"id"`
	Input types.String `tfsdk:"input"`
	Sha   types.String `tfsdk:"sha"`
}

func getEncoding(encodingName string) (encoding.Encoding, error) {
	switch encodingName {
	case "ISO-8859-1":
		return charmap.ISO8859_1, nil
	default:
		return nil, fmt.Errorf("unsupported encoding: %s", encodingName)
	}
}

func convertToBytes(inputString string, enc encoding.Encoding) ([]byte, error) {
	return enc.NewEncoder().Bytes([]byte(inputString))
}
