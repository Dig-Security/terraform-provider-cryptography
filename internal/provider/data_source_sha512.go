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
	"golang.org/x/text/encoding/ianaindex"
)

var _ datasource.DataSource = (*sha512DataSource)(nil)

func NewSha512DataSource() datasource.DataSource {
	return &sha512DataSource{}
}

type sha512DataSource struct{}

func (d *sha512DataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "sha512"
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
			"encoding": schema.StringAttribute{
				Required: true,
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

	enc, err := ianaindex.IANA.Encoding(model.Encoding.ValueString())
	if err != nil || enc == nil {
		resp.Diagnostics.AddError(
			"Not a supported IANA encoding",
			fmt.Sprintf("%s is not a supported IANA encoding name or alias in this Terraform version", model.Encoding.ValueString()),
		)
		return
	}
	var input = model.Input.ValueString()
	encodedBytes, err := convertToBytes(input, enc)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error convert input to bytes",
			fmt.Sprintf("Error: %v\n", err),
		)
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
	ID       types.String `tfsdk:"id"`
	Input    types.String `tfsdk:"input"`
	Sha      types.String `tfsdk:"sha"`
	Encoding types.String `tfsdk:"encoding"`
}

func convertToBytes(inputString string, enc encoding.Encoding) ([]byte, error) {
	return enc.NewEncoder().Bytes([]byte(inputString))
}
