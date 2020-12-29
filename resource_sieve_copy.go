package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSieveCopyCreateReadUpdate(d *schema.ResourceData, m interface{}) error {
	id := ""
	sieve := `# poste.io copy filter
require ["copy"];
if true
{
`
	for _, v := range d.Get("emails").([]interface{}) {
		email := v.(string)
		if id == "" {
			id = email
		} else {
			id = id + "," + email
		}
		sieve = sieve + "redirect :copy \"" + email + "\";\n"
	}
	sieve = sieve + "}"
	d.SetId(id)
	d.Set("sieve", sieve)
	return nil
}

func resourceSieveCopyDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceSieveCopyDiff(d *schema.ResourceDiff, m interface{}) error {
	if d.HasChange("emails") {
		d.SetNewComputed("sieve")
	}
	return nil
}

func resourceSieveCopy() *schema.Resource {
	return &schema.Resource{
		Create:        resourceSieveCopyCreateReadUpdate,
		Read:          resourceSieveCopyCreateReadUpdate,
		Update:        resourceSieveCopyCreateReadUpdate,
		Delete:        resourceSieveCopyDelete,
		CustomizeDiff: resourceSieveCopyDiff,

		Schema: map[string]*schema.Schema{
			"emails": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
			"sieve": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
