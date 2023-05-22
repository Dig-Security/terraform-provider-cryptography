package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func New() provider.Provider {
	return &shaProvider{}
}

var _ provider.Provider = (*shaProvider)(nil)

type shaProvider struct{}

func (p *shaProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "sha-calculator"
}

func (p *shaProvider) Schema(context.Context, provider.SchemaRequest, *provider.SchemaResponse) {
}

func (p *shaProvider) Configure(context.Context, provider.ConfigureRequest, *provider.ConfigureResponse) {
}

func (p *shaProvider) Resources(context.Context) []func() resource.Resource {
	return nil
}

func (p *shaProvider) DataSources(context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewSha512DataSource,
	}
}
