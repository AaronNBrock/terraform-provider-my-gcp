package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
//	"cloud.google.com/go/storage"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		// Schema: map[string]*schema.Schema{
		// 	"keyfile": {
		// 		Type:        schema.TypeString,
		// 		Required:    true,
		// 		Description: "Path to google cloud credential keyfile",
		// 	},

		// 	"project_id": {
		// 		Type:        schema.TypeString,
		// 		Required:    true,
		// 		Description: "The project id you'd like to create resources in.",
		// 	},
		// },

		ResourcesMap: map[string]*schema.Resource{
			"my-gcp_bucket": resourceServer(),
		},
	}
}
