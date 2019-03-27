package bluecat

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/acobaugh/terraform-provider-bluecat/client"
)

type Config struct {
	Username     string
	Password     string
	BaseURL      string
	InsecureFlag bool
}

// NewConfig returns a new Config from a supplied ResourceData.
func NewConfig(d *schema.ResourceData) (*Config, error) {
	c := &Config{
		Username:     d.Get("username").(string),
		Password:     d.Get("password").(string),
		InsecureFlag: d.Get("allow_unverified_ssl").(bool),
		BaseURL:      d.Get("base_url").(string),
	}

	return c, nil
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	c, err := NewConfig(d)
	if err != nil {
		return nil, err
	}
	return client.NewClient(&client.Config{
		Username: c.Username,
		Password: c.Password,
		Insecure: c.InsecureFlag,
		BaseURL: c.BaseURL,
	})
}
