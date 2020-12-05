/*
Package tfsrl is a generated package which contains definitions
of structs which represent a Terraform resource. 

This package was generated by ygocodegen
using the following YANG input files:
	- yang/20_06_2/srl/
Imported modules were sourced from:
	- yang/20_06_2/ietf/
*/
package tfsrl

import (
	"context"
	"strings"

	
	"fmt"
	
	"strconv"
	
	
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	log "github.com/sirupsen/logrus"
)

// resourceInterfacesSubinterfaceIpv6DhcpClientString function
func resourceInterfacesSubinterfaceIpv6DhcpClientString(d resourceIDStringer) string {
	return resourceIDString(d, "interfaces_subinterface_ipv6_dhcp_client")
}

// resourceInterfacesSubinterfaceIpv6DhcpClient function
func resourceInterfacesSubinterfaceIpv6DhcpClient() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInterfacesSubinterfaceIpv6DhcpClientCreate,
		ReadContext:   resourceInterfacesSubinterfaceIpv6DhcpClientRead,
		UpdateContext: resourceInterfacesSubinterfaceIpv6DhcpClientUpdate,
		DeleteContext: resourceInterfacesSubinterfaceIpv6DhcpClientDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
        "interface_id": {
            Type:     schema.TypeString,
            Required: true,
            ForceNew: true,
        },
        "subinterface_id": {
            Type:     schema.TypeString,
            Required: true,
            ForceNew: true,
        },
        "ipv6_id": {
            Type:     schema.TypeString,
            Required: true,
            ForceNew: true,
        },
        "dhcp_client": {
            Type:     schema.TypeList,
            Optional: true,
            MaxItems: 1,
            Elem: &schema.Resource{
            	Schema: map[string]*schema.Schema{
                    "trace_options": {
                        Type:     schema.TypeList,
                        Optional: true,
                        MaxItems: 1,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "trace": {
                                    Type:     schema.TypeString,
                                    Optional: true,
                                    ValidateFunc: validation.All(
                                        validation.StringInSlice([]string{
                                            "messages",
                                        }, false),
                                    ),
                                },
                            },
                        },
                    },
                },
            },
        },

        },
    }
}

func resourceInterfacesSubinterfaceIpv6DhcpClientCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Create: %s", resourceInterfacesSubinterfaceIpv6DhcpClientString(d))
	target := meta.(*Target)
	
	key := "dhcp-client"
	
	
	hkey0 := d.Get("interface_id").(string)
    
	hkey1 := d.Get("subinterface_id").(string)
    
	//hkey := d.Get("[interface_id subinterface_id]").(string)
	//p := fmt.Sprintf("fmt.Sprintf("/interface[name=%s]/subinterface[index=%s]/ipv6/dhcp-client",hkey0,hkey1)", hkey)
	p := fmt.Sprintf("/interface[name=%s]/subinterface[index=%s]/ipv6/dhcp-client",hkey0,hkey1)
	
	v := "dhcp_client"
	
	
	hid := make([]string, 0)
	
    hid = append(hid, "interface_id")
	
    hid = append(hid, "subinterface_id")
	
    hid = append(hid, "ipv6_id")
	
	//hid = append(hid, "ipv6_id")
	//hid := "ipv6_id"
	req, err := target.CreateSetRequest(&p, &v, &hid, d)
	
	if err != nil {
		return diag.FromErr(err)
	}

	response, err := target.Set(ctx, req)
	if err != nil {
		return diag.FromErr(err)
	}

	replaceInKeys(d.Get(v), "-", "_")
	log.Debugf("Set response: %v", response)

	d.SetId(key)
	return resourceInterfacesSubinterfaceIpv6DhcpClientRead(ctx, d, meta)
}

func resourceInterfacesSubinterfaceIpv6DhcpClientRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Read: %s", resourceInterfacesSubinterfaceIpv6DhcpClientString(d))
	target := meta.(*Target)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	
	
	hkey0 := d.Get("interface_id").(string)
    
	hkey1 := d.Get("subinterface_id").(string)
    
	//hkey := d.Get("[interface_id subinterface_id]").(string)
	

	
	
	//p := fmt.Sprintf("fmt.Sprintf("/interface[name=%s]/subinterface[index=%s]/ipv6/dhcp-client",hkey0,hkey1)", hkey)
	p := fmt.Sprintf("/interface[name=%s]/subinterface[index=%s]/ipv6/dhcp-client",hkey0,hkey1)
	
	

	req, err := target.CreateGetRequest(&p, "CONFIG", d)
	if err != nil {
		return diag.FromErr(err)
	}
	log.Infof("Get Request: %v", req)
	response, err := target.Get(ctx, req)
	if err != nil {
		return diag.FromErr(err)
	}
	log.Debugf("Get Gnmi read response: %v", response)

	u, err := target.HandleGetRespone(response)
	if err != nil {
		return diag.FromErr(err)
	}
	for i, upd := range u {
		// we expect a single response in the get since we target the explicit resource
		log.Debugf("get response: index: %d, update: %v", i, upd)
		if i <= 0 {
			data := make([]map[string]interface{}, 0)
			switch x := upd.Values["dhcp-client"].(type) {
			case map[string]interface{}:
				for k, v := range x {
					log.Debugf("BEFORE KEY: %s, VALUE: %v", k, v)
					sk := strings.Split(k, ":")[len(strings.Split(k, ":"))-1]

					switch sk {
					
					default:
						if k != sk {
							delete(x, k)
							x[sk] = v
						}
					}
                }
                for k, v := range x {
                    log.Debugf("AFTER KEY: %s, VALUE: %v", k, v)
				}
				
				data = append(data, x)
			}
			log.Debugf("get response: index: %d, data: %v", i, data)
			if err := d.Set("dhcp_client", data); err != nil {
				return diag.FromErr(err)
			}
			// always run
			
			d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
			
			return diags
		} else {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unexpected multiple response",
				Detail:   "We only expect a single response from the read/get response",
			})
			return diags
		}
	}
	// when the response is empty no data exists in the system
	log.Debugf("get response: empty set id to nill")
	d.SetId("")
	return diags
}

func resourceInterfacesSubinterfaceIpv6DhcpClientUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Update: %s", resourceInterfacesSubinterfaceIpv6DhcpClientString(d))
	target := meta.(*Target)
	
	key := "dhcp-client"
	
	
	hkey0 := d.Get("interface_id").(string)
    
	hkey1 := d.Get("subinterface_id").(string)
    
	//hkey := d.Get("[interface_id subinterface_id]").(string)
	//p := fmt.Sprintf("fmt.Sprintf("/interface[name=%s]/subinterface[index=%s]/ipv6/dhcp-client",hkey0,hkey1)", hkey)
	p := fmt.Sprintf("/interface[name=%s]/subinterface[index=%s]/ipv6/dhcp-client",hkey0,hkey1)
	
	v := "dhcp_client"
	
	
	hid := make([]string, 0)
	
    hid = append(hid, "interface_id")
	
    hid = append(hid, "subinterface_id")
	
    hid = append(hid, "ipv6_id")
	
	//hid = append(hid, "ipv6_id")
	//hid := "ipv6_id"
	req, err := target.CreateSetRequest(&p, &v, &hid, d)
	
	if err != nil {
		return diag.FromErr(err)
	}

	response, err := target.Set(ctx, req)
	if err != nil {
		return diag.FromErr(err)
	}

	replaceInKeys(d.Get(v), "-", "_")
	log.Debugf("Set response: %v", response)

	d.SetId(key)
	return resourceInterfacesSubinterfaceIpv6DhcpClientRead(ctx, d, meta)
}

func resourceInterfacesSubinterfaceIpv6DhcpClientDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Debugf("Beginning delete: %s", resourceInterfacesSubinterfaceIpv6DhcpClientString(d))
	target := meta.(*Target)

	
	
	hkey0 := d.Get("interface_id").(string)
    
	hkey1 := d.Get("subinterface_id").(string)
    
	//hkey := d.Get("[interface_id subinterface_id]").(string)
	
	//p := fmt.Sprintf("fmt.Sprintf("/interface[name=%s]/subinterface[index=%s]/ipv6/dhcp-client",hkey0,hkey1)", hkey)
	p := fmt.Sprintf("/interface[name=%s]/subinterface[index=%s]/ipv6/dhcp-client",hkey0,hkey1)
	
	
	req, err := target.CreateDeleteRequest(&p, d)
	if err != nil {
		return diag.FromErr(err)
	}

	response, err := target.Set(ctx, req)
	if err != nil {
		return diag.FromErr(err)
	}

	log.Debugf("Gnmi Delete Response: %v", response)

	d.SetId("")
	return nil
}