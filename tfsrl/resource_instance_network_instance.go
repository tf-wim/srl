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
	"regexp"
	
	
	
	"fmt"
	
	
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	log "github.com/sirupsen/logrus"
)

// resourceNetworkInstanceInstanceString function
func resourceNetworkInstanceInstanceString(d resourceIDStringer) string {
	return resourceIDString(d, "network_instance_instance")
}

// resourceNetworkInstanceInstance function
func resourceNetworkInstanceInstance() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNetworkInstanceInstanceCreate,
		ReadContext:   resourceNetworkInstanceInstanceRead,
		UpdateContext: resourceNetworkInstanceInstanceUpdate,
		DeleteContext: resourceNetworkInstanceInstanceDelete,

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
        "network_instance": {
            Type:     schema.TypeList,
            Optional: true,
            MaxItems: 1,
            Elem: &schema.Resource{
            	Schema: map[string]*schema.Schema{
                    "admin_state": {
                        Type:     schema.TypeString,
                        Optional: true,
                        Default: "enable",
                    },
                    "aft": {
                        Type:     schema.TypeList,
                        Optional: true,
                        MaxItems: 16,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "address_family": {
                                    Type:     schema.TypeString,
                                    Required: true,
                                    ForceNew: true,
                                },
                            },
                        },
                    },
                    "bridge_table": {
                        Type:     schema.TypeList,
                        Optional: true,
                        MaxItems: 1,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "discard_unknown_dest_mac": {
                                    Type:     schema.TypeBool,
                                    Optional: true,
                                    Default: false,
                                },
                                "mac_duplication": {
                                    Type:     schema.TypeList,
                                    Optional: true,
                                    MaxItems: 1,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "action": {
                                                Type:     schema.TypeString,
                                                Optional: true,
                                                Default: "stop-learning",
                                                ValidateFunc: validation.All(
                                                    validation.StringInSlice([]string{
                                                        "blackhole",
                                                        "oper-down",
                                                        "stop-learning",
                                                    }, false),
                                                ),
                                            },
                                            "admin_state": {
                                                Type:     schema.TypeString,
                                                Optional: true,
                                                Default: "enable",
                                            },
                                            "hold_down_time": {
                                                Type:     schema.TypeInt,
                                                Optional: true,
                                                Default: "9",
                                                ValidateFunc: validation.All(
                                                    validation.IntBetween(2, 60),
                                                ),
                                            },
                                            "monitoring_window": {
                                                Type:     schema.TypeInt,
                                                Optional: true,
                                                Default: "3",
                                                ValidateFunc: validation.All(
                                                    validation.IntBetween(1, 15),
                                                ),
                                            },
                                            "num_moves": {
                                                Type:     schema.TypeInt,
                                                Optional: true,
                                                Default: "5",
                                                ValidateFunc: validation.All(
                                                    validation.IntBetween(3, 10),
                                                ),
                                            },
                                        },
                                    },
                                },
                                "mac_learning": {
                                    Type:     schema.TypeList,
                                    Optional: true,
                                    MaxItems: 1,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "admin_state": {
                                                Type:     schema.TypeString,
                                                Optional: true,
                                                Default: "enable",
                                            },
                                            "aging": {
                                                Type:     schema.TypeList,
                                                Optional: true,
                                                MaxItems: 1,
                                                Elem: &schema.Resource{
                                                	Schema: map[string]*schema.Schema{
                                                        "admin_state": {
                                                            Type:     schema.TypeString,
                                                            Optional: true,
                                                            Default: "enable",
                                                        },
                                                        "age_time": {
                                                            Type:     schema.TypeInt,
                                                            Optional: true,
                                                            Default: "300",
                                                            ValidateFunc: validation.All(
                                                                validation.IntBetween(60, 86400),
                                                            ),
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                    },
                                },
                                "mac_limit": {
                                    Type:     schema.TypeList,
                                    Optional: true,
                                    MaxItems: 1,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "maximum_entries": {
                                                Type:     schema.TypeInt,
                                                Optional: true,
                                                Default: "250",
                                                ValidateFunc: validation.All(
                                                    validation.IntBetween(1, 8192),
                                                ),
                                            },
                                            "warning_threshold_pct": {
                                                Type:     schema.TypeInt,
                                                Optional: true,
                                                Default: "95",
                                                ValidateFunc: validation.All(
                                                    validation.IntBetween(6, 100),
                                                ),
                                            },
                                        },
                                    },
                                },
                                "static_mac": {
                                    Type:     schema.TypeList,
                                    Optional: true,
                                    MaxItems: 1,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "mac": {
                                                Type:     schema.TypeList,
                                                Optional: true,
                                                MaxItems: 16,
                                                Elem: &schema.Resource{
                                                	Schema: map[string]*schema.Schema{
                                                        "address": {
                                                            Type:     schema.TypeString,
                                                            Required: true,
                                                            ForceNew: true,
                                                            ValidateFunc: validation.All(
                                                                validation.StringMatch(regexp.MustCompile(`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`), "must match regex"),
                                                            ),
                                                        },
                                                        "destination": {
                                                            Type:     schema.TypeString,
                                                            Optional: true,
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
                        Optional: true,
                        ValidateFunc: validation.All(
                            validation.StringLenBetween(1, 255),
                            validation.StringMatch(regexp.MustCompile("[A-Za-z0-9 !@#$%!^(MISSING)&()|+=`~.,'/_:;?-]*"), "must match regex"),
                        ),
                    },
                    "interface": {
                        Type:     schema.TypeList,
                        Optional: true,
                        MaxItems: 16,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "name": {
                                    Type:     schema.TypeString,
                                    Required: true,
                                    ForceNew: true,
                                    ValidateFunc: validation.All(
                                        validation.StringLenBetween(5, 24),
                                        validation.StringMatch(regexp.MustCompile(`(mgmt0\.0|lo(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])\.(0|[1-9](\d){0,3})|ethernet-([1-9](\d){0,1}(/[abcd])?(/[1-9](\d){0,1})?/[1-9](\d){0,1})\.([0]|[1-9](\d){0,3})|irb(0|1[0-9][0-9]|2([0-4][0-9]|5[0-5])|[1-9][0-9]|[1-9])\.(0|[1-9](\d){0,3})|lag(6[0-4][0-9][0-9][0-9]|65[0-4][0-9][0-9]|655[0-2][0-9]|6553[0-5]|[1-5][0-9][0-9][0-9][0-9]|[1-9][0-9][0-9][0-9]|[1-9][0-9][0-9]|[1-9][0-9]|[1-9])\.(0|[1-9](\d){0,3}))`), "must match regex"),
                                    ),
                                },
                            },
                        },
                    },
                    "ip_forwarding": {
                        Type:     schema.TypeList,
                        Optional: true,
                        MaxItems: 1,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "receive_ipv4_check": {
                                    Type:     schema.TypeBool,
                                    Optional: true,
                                },
                                "receive_ipv6_check": {
                                    Type:     schema.TypeBool,
                                    Optional: true,
                                },
                            },
                        },
                    },
                    "ip_load_balancing": {
                        Type:     schema.TypeList,
                        Optional: true,
                        MaxItems: 1,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "resilient_hash_prefix": {
                                    Type:     schema.TypeList,
                                    Optional: true,
                                    MaxItems: 16,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "hash_buckets_per_path": {
                                                Type:     schema.TypeInt,
                                                Optional: true,
                                                Default: "1",
                                                ValidateFunc: validation.All(
                                                    validation.IntBetween(1, 32),
                                                ),
                                            },
                                            "ip_prefix": {
                                                Type:     schema.TypeString,
                                                Required: true,
                                                ForceNew: true,
                                                ValidateFunc: validation.Any(
                                                    validation.StringMatch(regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))`), "must match regex"),
                                                    validation.StringMatch(regexp.MustCompile(`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`), "must match regex"),
                                                ),
                                            },
                                            "max_paths": {
                                                Type:     schema.TypeInt,
                                                Optional: true,
                                                Default: "1",
                                                ValidateFunc: validation.All(
                                                    validation.IntBetween(1, 64),
                                                ),
                                            },
                                        },
                                    },
                                },
                            },
                        },
                    },
                    "mpls": {
                        Type:     schema.TypeList,
                        Optional: true,
                        MaxItems: 1,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "admin_state": {
                                    Type:     schema.TypeString,
                                    Optional: true,
                                    Default: "disable",
                                },
                                "static_mpls_entry": {
                                    Type:     schema.TypeList,
                                    Optional: true,
                                    MaxItems: 16,
                                    Elem: &schema.Resource{
                                    	Schema: map[string]*schema.Schema{
                                            "collect_stats": {
                                                Type:     schema.TypeBool,
                                                Optional: true,
                                                Default: false,
                                            },
                                            "next_hop_group": {
                                                Type:     schema.TypeString,
                                                Optional: true,
                                            },
                                            "operation": {
                                                Type:     schema.TypeString,
                                                Optional: true,
                                                Default: "swap",
                                                ValidateFunc: validation.All(
                                                    validation.StringInSlice([]string{
                                                        "pop",
                                                        "swap",
                                                    }, false),
                                                ),
                                            },
                                            "preference": {
                                                Type:     schema.TypeInt,
                                                Optional: true,
                                                Default: "5",
                                                ValidateFunc: validation.All(
                                                    validation.IntBetween(0, 255),
                                                ),
                                            },
                                            "top_label": {
                                                Type:     schema.TypeString,
                                                Required: true,
                                                ForceNew: true,
                                            },
                                        },
                                    },
                                },
                                "ttl_propagation": {
                                    Type:     schema.TypeBool,
                                    Optional: true,
                                    Default: false,
                                },
                            },
                        },
                    },
                    "mtu": {
                        Type:     schema.TypeList,
                        Optional: true,
                        MaxItems: 1,
                        Elem: &schema.Resource{
                        	Schema: map[string]*schema.Schema{
                                "path_mtu_discovery": {
                                    Type:     schema.TypeBool,
                                    Optional: true,
                                },
                            },
                        },
                    },
                    "name": {
                        Type:     schema.TypeString,
                        Required: true,
                        ForceNew: true,
                        ValidateFunc: validation.All(
                            validation.StringLenBetween(1, 255),
                            validation.StringMatch(regexp.MustCompile("[A-Za-z0-9 !@#$%!^(MISSING)&()|+=`~.,'/_:;?-]*"), "must match regex"),
                        ),
                    },
                    "router_id": {
                        Type:     schema.TypeString,
                        Optional: true,
                        ValidateFunc: validation.All(
                            validation.StringMatch(regexp.MustCompile(`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`), "must match regex"),
                        ),
                    },
                    "type": {
                        Type:     schema.TypeString,
                        Optional: true,
                        Default: "default",
                    },
                },
            },
        },

        },
    }
}

func resourceNetworkInstanceInstanceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Create: %s", resourceNetworkInstanceInstanceString(d))
	target := meta.(*Target)
	 
	rn := "network_instance"
	rk := "name"
	key, err := getResourceListKey(&rn, &rk, d)
	if err != nil {
		return diag.FromErr(err)
	}
	
	p := "/"
	
	v := ""
	
	
	hid := make([]string, 0)
	//hid := ""
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
	return resourceNetworkInstanceInstanceRead(ctx, d, meta)
}

func resourceNetworkInstanceInstanceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Read: %s", resourceNetworkInstanceInstanceString(d))
	target := meta.(*Target)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	

	 
	//rn := "network_instance"
	rk := "name"
	key:= d.Id()

	
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

func resourceNetworkInstanceInstanceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Infof("Beginning Update: %s", resourceNetworkInstanceInstanceString(d))
	target := meta.(*Target)
	 
	rn := "network_instance"
	rk := "name"
	key, err := getResourceListKey(&rn, &rk, d)
	if err != nil {
		return diag.FromErr(err)
	}
	
	p := "/"
	
	v := ""
	
	
	hid := make([]string, 0)
	//hid := ""
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
	return resourceNetworkInstanceInstanceRead(ctx, d, meta)
}

func resourceNetworkInstanceInstanceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Debugf("Beginning delete: %s", resourceNetworkInstanceInstanceString(d))
	target := meta.(*Target)

	
	 
	//p := fmt.Sprintf("fmt.Sprintf("/network-instance[name=%s]", d.Id())", d.Id())
	p := fmt.Sprintf("/network-instance[name=%s]", d.Id())
	
	
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
