package main

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceBoxCreate(d *schema.ResourceData, m interface{}) error {
	email := d.Get("email").(string)
	password := d.Get("password").(string)

	d.SetId(email)

	api := PosteApi(m)
	err := api.CreateBox(email, password)

	if err != nil {
		return fmt.Errorf("%s", err)
	}

	return nil
}

func resourceBoxRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceBoxUpdate(d *schema.ResourceData, m interface{}) error {
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

func resourceBox() *schema.Resource {
	return &schema.Resource{
		Create: resourceBoxCreate,
		Read:   resourceBoxRead,
		Update: resourceBoxUpdate,
		Delete: resourceBoxDelete,

		Schema: map[string]*schema.Schema{
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
