package cloudconnexa

import (
	"context"

	"github.com/openvpn/cloudconnexa-go-client/v2/cloudconnexa"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRoute() *schema.Resource {
	return &schema.Resource{
		Description:   "Use `cloudconnexa_route` to create a route on an CloudConnexa network.",
		CreateContext: resourceRouteCreate,
		UpdateContext: resourceRouteUpdate,
		ReadContext:   resourceRouteRead,
		DeleteContext: resourceRouteDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"IP_V4", "IP_V6"}, false),
				Description:  "The type of route. Valid values are `IP_V4`, `IP_V6`, and `DOMAIN`.",
			},
			"subnet": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The target value of the default route.",
			},
			"network_item_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the network on which to create the route.",
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Managed by Terraform",
			},
		},
	}
}

func resourceRouteCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*cloudconnexa.Client)
	var diags diag.Diagnostics
	networkItemId := d.Get("network_item_id").(string)
	routeType := d.Get("type").(string)
	routeSubnet := d.Get("subnet").(string)
	descriptionValue := d.Get("description").(string)
	r := cloudconnexa.Route{
		Type:        routeType,
		Subnet:      routeSubnet,
		Description: descriptionValue,
	}
	route, err := c.Routes.Create(networkItemId, r)
	if err != nil {
		return append(diags, diag.FromErr(err)...)
	}
	d.SetId(route.ID)
	if routeType == "IP_V4" || routeType == "IP_V6" {
		d.Set("subnet", route.Subnet)
	}
	return diags
}

func resourceRouteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*cloudconnexa.Client)
	var diags diag.Diagnostics
	routeId := d.Id()
	r, err := c.Routes.Get(routeId)
	if err != nil {
		return append(diags, diag.FromErr(err)...)
	}
	if r == nil {
		d.SetId("")
	} else {
		d.Set("type", r.Type)
		if r.Type == "IP_V4" || r.Type == "IP_V6" {
			d.Set("subnet", r.Subnet)
		}
		d.Set("description", r.Description)
		d.Set("network_item_id", r.NetworkItemID)
	}
	return diags
}

func resourceRouteUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*cloudconnexa.Client)
	var diags diag.Diagnostics
	if !d.HasChanges("description", "subnet") {
		return diags
	}

	_, description := d.GetChange("description")
	_, subnet := d.GetChange("subnet")
	r := cloudconnexa.Route{
		ID:          d.Id(),
		Description: description.(string),
		Subnet:      subnet.(string),
	}

	err := c.Routes.Update(r)
	if err != nil {
		return append(diags, diag.FromErr(err)...)
	}
	return diags
}

func resourceRouteDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*cloudconnexa.Client)
	var diags diag.Diagnostics
	routeId := d.Id()
	err := c.Routes.Delete(routeId)
	if err != nil {
		return append(diags, diag.FromErr(err)...)
	}
	return diags
}
