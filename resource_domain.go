package main

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceDomainCreate(d *schema.ResourceData, m interface{}) error {
	domain := d.Get("name").(string)

	d.SetId(domain)

	api := PosteApi(m)
	err := api.CreateDomain(domain)

	if err != nil {
		return fmt.Errorf("%s", err)
	}

	return nil
}

func resourceDomainRead(d *schema.ResourceData, m interface{}) error {
	domain := d.Id()
	d.Set("name", domain)
	return nil
}

func resourceDomainUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceDomainDelete(d *schema.ResourceData, m interface{}) error {
	domain := d.Get("name").(string)

	api := PosteApi(m)
	err := api.DeleteDomain(domain)

	if err != nil {
		return fmt.Errorf("%s", err)
	}

	return nil
}

func resourceDomainImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	domain := d.Id()
	d.Set("name", domain)
	return []*schema.ResourceData{d}, nil
}

func resourceDoamin() *schema.Resource {
	return &schema.Resource{
		Create: resourceDomainCreate,
		Read:   resourceDomainRead,
		Update: resourceDomainUpdate,
		Delete: resourceDomainDelete,
		Importer: &schema.ResourceImporter{
			State: resourceDomainImport,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
