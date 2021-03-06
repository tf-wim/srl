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

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	log "github.com/sirupsen/logrus"
)

// dataInterfacesSubinterfaceIpv6String function
func dataInterfacesSubinterfaceIpv6String(d resourceIDStringer) string {
	return resourceIDString(d, "interfaces_subinterface_ipv6")
}

// dataInterfacesSubinterfaceIpv6 function
func dataInterfacesSubinterfaceIpv6() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataInterfacesSubinterfaceIpv6Read,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Read:   schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
        "interface_id": {
            Type:     schema.TypeString,
            Required: true,
        },
        "subinterface_id": {
            Type:     schema.TypeString,
            Required: true,
        },
        "ipv6": {
            Type:     schema.TypeList,
            Computed: true,
            Elem: &schema.Resource{
            	Schema: map[string]*schema.Schema{
                    "neighbor_discovery": {
                        Type:     schema.TypeList,
                        Computed: true,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "duplicate_address_detection": {
                                    Type:     schema.TypeBool,
                                    Computed: true,
                                },
                                "neighbor": {
                                    Type:     schema.TypeList,
                                    Computed: true,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "ipv6_address": {
                                                Type:     schema.TypeString,
                                                Required: true,
                                            },
                                            "link_layer_address": {
                                                Type:     schema.TypeString,
                                                Computed: true,
                                            },
                                        },
                                    },
                                },
                                "reachable_time": {
                                    Type:     schema.TypeInt,
                                    Computed: true,
                                },
                                "stale_time": {
                                    Type:     schema.TypeInt,
                                    Computed: true,
                                },
                            },
                        },
                    },
                    "router_advertisement": {
                        Type:     schema.TypeList,
                        Computed: true,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "router_role": {
                                    Type:     schema.TypeList,
                                    Computed: true,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "admin_state": {
                                                Type:     schema.TypeString,
                                                Computed: true,
                                            },
                                            "current_hop_limit": {
                                                Type:     schema.TypeInt,
                                                Computed: true,
                                            },
                                            "ip_mtu": {
                                                Type:     schema.TypeInt,
                                                Computed: true,
                                            },
                                            "managed_configuration_flag": {
                                                Type:     schema.TypeBool,
                                                Computed: true,
                                            },
                                            "max_advertisement_interval": {
                                                Type:     schema.TypeInt,
                                                Computed: true,
                                            },
                                            "min_advertisement_interval": {
                                                Type:     schema.TypeInt,
                                                Computed: true,
                                            },
                                            "other_configuration_flag": {
                                                Type:     schema.TypeBool,
                                                Computed: true,
                                            },
                                            "prefix": {
                                                Type:     schema.TypeList,
                                                Computed: true,
                                                Elem: &schema.Resource{
                                                	Schema: map[string]*schema.Schema{
                                                        "autonomous_flag": {
                                                            Type:     schema.TypeBool,
                                                            Computed: true,
                                                        },
                                                        "ipv6_prefix": {
                                                            Type:     schema.TypeString,
                                                            Required: true,
                                                        },
                                                        "on_link_flag": {
                                                            Type:     schema.TypeBool,
                                                            Computed: true,
                                                        },
                                                        "preferred_lifetime": {
                                                            Type:     schema.TypeString,
                                                            Computed: true,
                                                        },
                                                        "valid_lifetime": {
                                                            Type:     schema.TypeString,
                                                            Computed: true,
                                                        },
                                                    },
                                                },
                                            },
                                            "reachable_time": {
                                                Type:     schema.TypeInt,
                                                Computed: true,
                                            },
                                            "retransmit_time": {
                                                Type:     schema.TypeInt,
                                                Computed: true,
                                            },
                                            "router_lifetime": {
                                                Type:     schema.TypeInt,
                                                Computed: true,
                                            },
                                        },
                                    },
                                },
                                "trace_options": {
                                    Type:     schema.TypeList,
                                    Computed: true,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "trace": {
                                                Type:     schema.TypeString,
                                                Computed: true,
                                            },
                                        },
                                    },
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

func dataInterfacesSubinterfaceIpv6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Read: %s", dataInterfacesSubinterfaceIpv6String(d))
	target := meta.(*Target)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	
	
	hkey0 := d.Get("interface_id").(string)
    
	hkey1 := d.Get("subinterface_id").(string)
    
	//hkey := d.Get("[interface_id subinterface_id]").(string)
	

	
	
	//p := fmt.Sprintf("fmt.Sprintf("/interface[name=%s]/subinterface[index=%s]/ipv6",hkey0,hkey1)", hkey)
	p := fmt.Sprintf("/interface[name=%s]/subinterface[index=%s]/ipv6",hkey0,hkey1)
	
	

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
			switch x := upd.Values["ipv6"].(type) {
			case map[string]interface{}:
				for k, v := range x {
					log.Debugf("BEFORE KEY: %s, VALUE: %v", k, v)
					sk := strings.Split(k, ":")[len(strings.Split(k, ":"))-1]

					switch sk {
					
					case "neighbor_discovery":
						delete(x, k)
					
					case "dhcp_client":
						delete(x, k)
					
					case "dhcp_relay":
						delete(x, k)
					
					case "vrrp":
						delete(x, k)
					
					case "address":
						delete(x, k)
					
					case "router_advertisement":
						delete(x, k)
					
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
			if err := d.Set("ipv6", data); err != nil {
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
