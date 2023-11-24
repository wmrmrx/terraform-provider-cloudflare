package origin_ca_certificate

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *CloudflareOriginCACertificateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Use this data source to retrieve an existing origin ca certificate.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The user's unique identifier.",
				Computed:    true,
			},
			"certificate": schema.StringAttribute{
				Computed:    true,
				Description: "The Origin CA certificate.",
			},
			"csr": schema.StringAttribute{
				Computed:    true,
				Description: "The Certificate Signing Request. Must be newline-encoded.",
			},
			"expires_on": schema.StringAttribute{
				Computed:    true,
				Description: "The datetime when the certificate will expire.",
			},
			"hostnames": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "A list of hostnames or wildcard names bound to the certificate.",
			},
			"request_type": schema.StringAttribute{
				Computed:    true,
				Description: fmt.Sprintf("The signature type desired on the certificate. %v", []string{"origin-rsa", "origin-ecc", "keyless-certificate"}),
			},
			"requested_validity": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: fmt.Sprintf("The number of days for which the certificate should be valid. %v", []int{7, 30, 90, 365, 730, 1095, 5475}),
			},
			"revoked_at": schema.StringAttribute{
				Computed:    true,
				Description: "The datetime when the certificate was revoked.",
			},
		},
	}
}
