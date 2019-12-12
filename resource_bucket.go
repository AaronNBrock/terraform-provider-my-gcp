package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/customdiff"
	"fmt"
	"google.golang.org/api/option"
	"cloud.google.com/go/storage"
	"context"
	"strings"
)

const (
	keyfile = "keyfile.json"
	projectID = "contino-9e7fb483b94887bf"
)

func resourceServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,

		Schema: map[string]*schema.Schema{
			"bucket_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},

		CustomizeDiff: customdiff.ForceNewIfChange(
			"bucket_name",
			func (old, new, meta interface{}) bool {
                // "size" can only increase in-place, so we must create a new resource
                // if it is decreased.
                return new.(string) != old.(string)
            },
		),
	}
}

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	fmt.Println("Create run")
	bucketName := d.Get("bucket_name").(string)

	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(keyfile))
	if err != nil {
		return err
	}

	bkt := client.Bucket(bucketName)
	if err := bkt.Create(ctx, projectID, nil); err != nil {
		return err
	}

	d.SetId(bucketName)
	return resourceServerRead(d, m)
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	bucketName := d.Get("bucket_name").(string)

	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(keyfile))
	if err != nil {
		return err
	}

	bkt := client.Bucket(bucketName)

	_, err = bkt.Attrs(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "bucket doesn't exist") {
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	fmt.Println("Update run")
	return fmt.Errorf("todo: this shouldn't ever be called")
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	fmt.Println("Delete run")
	bucketName := d.Get("bucket_name").(string)

	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(keyfile))
	if err != nil {
		return err
	}

	bkt := client.Bucket(bucketName)
	if err := bkt.Delete(ctx); err != nil {
		return err
	}
	return nil
}
