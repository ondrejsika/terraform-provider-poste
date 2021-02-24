package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ondrejsika/poste-go"
)

func updateBoxSieve(api poste.PosteAPI, email string, sieve string) error {
	if sieve == "" {
		// Default sieve script for poste.io
		sieve = `require ["fileinto"];
if header :contains "subject" "*****SPAM*****"
{
      fileinto "Junk";
}`
	}
	return api.UpdateBoxSieve(email, sieve)
}

func resourceBoxCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	email := d.Get("email").(string)
	password := d.Get("password").(string)
	sieve := d.Get("sieve").(string)

	d.SetId(email)

	api := PosteApi(m)
	var err error
	err = api.CreateBox(email, password)
	if err != nil {
		return diag.FromErr(err)
	}
	err = updateBoxSieve(api, email, sieve)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceBoxRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	email := d.Get("email").(string)
	d.Set("email", email)
	return diags
}

func resourceBoxUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	email := d.Get("email").(string)
	password := d.Get("password").(string)
	sieve := d.Get("sieve").(string)
	api := PosteApi(m)
	var err error
	err = api.UpdateBoxPassword(email, password)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("password", password)
	err = updateBoxSieve(api, email, sieve)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("sieve", sieve)
	return diags
}

func resourceBoxDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	email := d.Get("email").(string)

	api := PosteApi(m)
	err := api.DeleteBox(email)

	if err != nil {
		return diag.FromErr(err)
	}

	var diags diag.Diagnostics
	return diags
}

func resourceBoxImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	email := d.Id()
	d.Set("email", email)
	api := PosteApi(m)
	var err error
	sieve, err := api.GetBoxSieve(email)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	d.Set("sieve", sieve)
	return []*schema.ResourceData{d}, nil
}

func resourceBox() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBoxCreate,
		ReadContext:   resourceBoxRead,
		UpdateContext: resourceBoxUpdate,
		DeleteContext: resourceBoxDelete,
		Importer: &schema.ResourceImporter{
			State: resourceBoxImport,
		},

		Schema: map[string]*schema.Schema{
			"email": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"sieve": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
