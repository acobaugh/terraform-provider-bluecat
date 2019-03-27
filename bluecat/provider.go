package bluecat

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("BLUECAT_USERNAME", nil),
				Description: "The user name for Bluecat API operations.",
			},

			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("BLUECAT_PASSWORD", nil),
				Description: "The user password for Bluecat API operations.",
			},

			"base_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("BLUECAT_BASE_URL", nil),
				Description: "The base URL for Bluecat API operations.",
			},
			"allow_unverified_ssl": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("BLUECAT_ALLOW_UNVERIFIED_SSL", false),
				Description: "If set, will permit unverifiable SSL certificates.",
			},
			"configuration_name": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("BLUECAT_CONFIGURATION_NAME", nil),
				Deprecated:  "",
			},
			"dns_view": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("BLUECAT_DNS_VIEW", false),
				Description: "",
			},
		},

		ResourcesMap: map[string]*schema.Resource{},

		DataSourcesMap: map[string]*schema.Resource{},

		ConfigureFunc: providerConfigure,
	}
}
