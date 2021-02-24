package main

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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

func resourceBoxCreate(d *schema.ResourceData, m interface{}) error {
	email := d.Get("email").(string)
	password := d.Get("password").(string)
	sieve := d.Get("sieve").(string)

	d.SetId(email)

	api := PosteApi(m)
	var err error
	err = api.CreateBox(email, password)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	err = updateBoxSieve(api, email, sieve)
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	return nil
}

func resourceBoxRead(d *schema.ResourceData, m interface{}) error {
	email := d.Get("email").(string)
	d.Set("email", email)
	return nil
}

func resourceBoxUpdate(d *schema.ResourceData, m interface{}) error {
	email := d.Get("email").(string)
	password := d.Get("password").(string)
	sieve := d.Get("sieve").(string)
	api := PosteApi(m)
	var err error
	err = api.UpdateBoxPassword(email, password)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	d.Set("password", password)
	err = updateBoxSieve(api, email, sieve)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	d.Set("sieve", sieve)
	return nil
}

func resourceBoxDelete(d *schema.ResourceData, m interface{}) error {
	email := d.Get("email").(string)

	api := PosteApi(m)
	err := api.DeleteBox(email)

	if err != nil {
		return fmt.Errorf("%s", err)
	}

	return nil
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
		Create: resourceBoxCreate,
		Read:   resourceBoxRead,
		Update: resourceBoxUpdate,
		Delete: resourceBoxDelete,
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
