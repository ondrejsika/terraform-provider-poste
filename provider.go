package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Config struct {
	Origin   string
	Username string
	Password string
}

func Provider() *schema.Provider {
	p := &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"poste_domain":     resourceDoamin(),
			"poste_box":        resourceBox(),
			"poste_sieve_copy": resourceSieveCopy(),
		},
		Schema: map[string]*schema.Schema{
			"origin": {
				Type:     schema.TypeString,
				Required: true,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
	p.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {

		config := Config{
			Origin:   d.Get("origin").(string),
			Username: d.Get("username").(string),
			Password: d.Get("password").(string),
		}
		return &config, nil
	}
	return p
}
