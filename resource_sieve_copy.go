package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSieveCopyCreateReadUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
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
	return diags
}

func resourceSieveCopyDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceSieveCopyDiff(ctx context.Context, d *schema.ResourceDiff, m interface{}) error {
	if d.HasChange("emails") {
		d.SetNewComputed("sieve")
	}
	return nil
}

// func resourceSieveCopyDiff(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics
// 	if d.HasChange("emails") {
// 		d.SetNewComputed("sieve")
// 	}
// 	return diags
// }

func resourceSieveCopy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSieveCopyCreateReadUpdate,
		ReadContext:   resourceSieveCopyCreateReadUpdate,
		UpdateContext: resourceSieveCopyCreateReadUpdate,
		DeleteContext: resourceSieveCopyDelete,
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
