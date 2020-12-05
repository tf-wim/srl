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
	
	
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	log "github.com/sirupsen/logrus"
)

// dataNetworkInstanceInstanceString function
func dataNetworkInstanceInstanceString(d resourceIDStringer) string {
	return resourceIDString(d, "network_instance_instance")
}

// dataNetworkInstanceInstance function
func dataNetworkInstanceInstance() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataNetworkInstanceInstanceRead,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Read:   schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
        "network_instance": {
            Type:     schema.TypeList,
            Required: true,
            Elem: &schema.Resource{
            	Schema: map[string]*schema.Schema{
                    "admin_state": {
                        Type:     schema.TypeString,
                        Computed: true,
                    },
                    "aft": {
                        Type:     schema.TypeList,
                        Computed: true,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "address_family": {
                                    Type:     schema.TypeString,
                                    Required: true,
                                },
                            },
                        },
                    },
                    "bridge_table": {
                        Type:     schema.TypeList,
                        Computed: true,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "discard_unknown_dest_mac": {
                                    Type:     schema.TypeBool,
                                    Computed: true,
                                },
                                "mac_duplication": {
                                    Type:     schema.TypeList,
                                    Computed: true,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "action": {
                                                Type:     schema.TypeString,
                                                Computed: true,
                                            },
                                            "admin_state": {
                                                Type:     schema.TypeString,
                                                Computed: true,
                                            },
                                            "hold_down_time": {
                                                Type:     schema.TypeInt,
                                                Computed: true,
                                            },
                                            "monitoring_window": {
                                                Type:     schema.TypeInt,
                                                Computed: true,
                                            },
                                            "num_moves": {
                                                Type:     schema.TypeInt,
                                                Computed: true,
                                            },
                                        },
                                    },
                                },
                                "mac_learning": {
                                    Type:     schema.TypeList,
                                    Computed: true,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "admin_state": {
                                                Type:     schema.TypeString,
                                                Computed: true,
                                            },
                                            "aging": {
                                                Type:     schema.TypeList,
                                                Computed: true,
                                                Elem: &schema.Resource{
                                                	Schema: map[string]*schema.Schema{
                                                        "admin_state": {
                                                            Type:     schema.TypeString,
                                                            Computed: true,
                                                        },
                                                        "age_time": {
                                                            Type:     schema.TypeInt,
                                                            Computed: true,
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                    },
                                },
                                "mac_limit": {
                                    Type:     schema.TypeList,
                                    Computed: true,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "maximum_entries": {
                                                Type:     schema.TypeInt,
                                                Computed: true,
                                            },
                                            "warning_threshold_pct": {
                                                Type:     schema.TypeInt,
                                                Computed: true,
                                            },
                                        },
                                    },
                                },
                                "static_mac": {
                                    Type:     schema.TypeList,
                                    Computed: true,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "mac": {
                                                Type:     schema.TypeList,
                                                Computed: true,
                                                Elem: &schema.Resource{
                                                	Schema: map[string]*schema.Schema{
                                                        "address": {
                                                            Type:     schema.TypeString,
                                                            Required: true,
                                                        },
                                                        "destination": {
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
                    "description": {
                        Type:     schema.TypeString,
                        Computed: true,
                    },
                    "interface": {
                        Type:     schema.TypeList,
                        Computed: true,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "name": {
                                    Type:     schema.TypeString,
                                    Required: true,
                                },
                            },
                        },
                    },
                    "ip_forwarding": {
                        Type:     schema.TypeList,
                        Computed: true,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "receive_ipv4_check": {
                                    Type:     schema.TypeBool,
                                    Computed: true,
                                },
                                "receive_ipv6_check": {
                                    Type:     schema.TypeBool,
                                    Computed: true,
                                },
                            },
                        },
                    },
                    "ip_load_balancing": {
                        Type:     schema.TypeList,
                        Computed: true,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "resilient_hash_prefix": {
                                    Type:     schema.TypeList,
                                    Computed: true,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "hash_buckets_per_path": {
                                                Type:     schema.TypeInt,
                                                Computed: true,
                                            },
                                            "ip_prefix": {
                                                Type:     schema.TypeString,
                                                Required: true,
                                            },
                                            "max_paths": {
                                                Type:     schema.TypeInt,
                                                Computed: true,
                                            },
                                        },
                                    },
                                },
                            },
                        },
                    },
                    "mpls": {
                        Type:     schema.TypeList,
                        Computed: true,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "admin_state": {
                                    Type:     schema.TypeString,
                                    Computed: true,
                                },
                                "static_mpls_entry": {
                                    Type:     schema.TypeList,
                                    Computed: true,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "collect_stats": {
                                                Type:     schema.TypeBool,
                                                Computed: true,
                                            },
                                            "next_hop_group": {
                                                Type:     schema.TypeString,
                                                Computed: true,
                                            },
                                            "operation": {
                                                Type:     schema.TypeString,
                                                Computed: true,
                                            },
                                            "preference": {
                                                Type:     schema.TypeInt,
                                                Computed: true,
                                            },
                                            "top_label": {
                                                Type:     schema.TypeString,
                                                Required: true,
                                            },
                                        },
                                    },
                                },
                                "ttl_propagation": {
                                    Type:     schema.TypeBool,
                                    Computed: true,
                                },
                            },
                        },
                    },
                    "mtu": {
                        Type:     schema.TypeList,
                        Computed: true,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "path_mtu_discovery": {
                                    Type:     schema.TypeBool,
                                    Computed: true,
                                },
                            },
                        },
                    },
                    "name": {
                        Type:     schema.TypeString,
                        Required: true,
                    },
                    "router_id": {
                        Type:     schema.TypeString,
                        Computed: true,
                    },
                    "type": {
                        Type:     schema.TypeString,
                        Computed: true,
                    },
                },
            },
        },

        },
    }
}

func dataNetworkInstanceInstanceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Read: %s", dataNetworkInstanceInstanceString(d))
	target := meta.(*Target)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	

	 
	rn := "network_instance"
	rk := "name"
	key, err := getResourceListKey(&rn, &rk, d)

	
	//p := fmt.Sprintf("fmt.Sprintf("/network-instance[name=%s]", key)", key)
	p := fmt.Sprintf("/network-instance[name=%s]", key)
	
	

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
			switch x := upd.Values["network-instance"].(type) {
			case map[string]interface{}:
				for k, v := range x {
					log.Debugf("BEFORE KEY: %s, VALUE: %v", k, v)
					sk := strings.Split(k, ":")[len(strings.Split(k, ":"))-1]

					switch sk {
					
					case "protocols":
						delete(x, k)
					
					case "protocols/bgp":
						delete(x, k)
					
					case "protocols/isis":
						delete(x, k)
					
					case "protocols/ospf":
						delete(x, k)
					
					case "static_routes":
						delete(x, k)
					
					case "aggregate_routes":
						delete(x, k)
					
					case "next_hop_groups":
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
				 
				// add key to the get resp data since it is not returned in the gnmi data
				x[rk] = key
				// append the get resp to data
				
				data = append(data, x)
			}
			log.Debugf("get response: index: %d, data: %v", i, data)
			if err := d.Set("network_instance", data); err != nil {
				return diag.FromErr(err)
			}
			// always run
			 
			d.SetId(key)
			
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