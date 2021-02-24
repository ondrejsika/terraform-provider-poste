package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDomainCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	domain := d.Get("name").(string)

	d.SetId(domain)

	api := PosteApi(m)
	err := api.CreateDomain(domain)

	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceDomainRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	domain := d.Id()
	d.Set("name", domain)
	return diags
}

func resourceDomainUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceDomainDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	domain := d.Get("name").(string)

	api := PosteApi(m)
	err := api.DeleteDomain(domain)

	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceDomainImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	domain := d.Id()
	d.Set("name", domain)
	return []*schema.ResourceData{d}, nil
}

func resourceDoamin() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDomainCreate,
		ReadContext:   resourceDomainRead,
		UpdateContext: resourceDomainUpdate,
		DeleteContext: resourceDomainDelete,
		Importer: &schema.ResourceImporter{
			State: resourceDomainImport,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
