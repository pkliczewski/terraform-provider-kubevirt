package kubevirt

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func virtualMachineSpecFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"affinity": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "If affinity is specifies, obey all the affinity rules",
			Elem: &schema.Resource{
				Schema: affinityFields(),
			},
		},
		"domain": {},
		"hostname": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Specifies the hostname of the vmi If not specified, the hostname will be set to the name of the vmi, if dhcp or cloud-init is configured properly. +optional",
		},
		"liveness_probe": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			ForceNew:    true,
			Description: "Periodic probe of VirtualMachineInstance liveness. VirtualmachineInstances will be stopped if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes +optional",
			Elem:        probeSchema(),
		},
		"networks": {},
		"node_selector": {
			Type:        schema.TypeMap,
			Optional:    true,
			Description: "NodeSelector is a selector which must be true for the vmi to fit on a node. Selector which must match a node’s labels for the vmi to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/ +optional",
		},
		"readiness_probe": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			ForceNew:    true,
			Description: "Periodic probe of VirtualMachineInstance service readiness. VirtualmachineInstances will be removed from service endpoints if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes +optional",
			Elem:        probeSchema(),
		},
		"subdomain": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "If specified, the fully qualified vmi hostname will be \"<hostname>.<subdomain>.<pod namespace>.svc.<cluster domain>\". If not specified, the vmi will not have a domainname at all. The DNS entry will resolve to the vmi, no matter if the vmi itself can pick up a hostname. +optional",
		},
		"termination_grace_period_seconds": {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      30,
			ValidateFunc: validateTerminationGracePeriodSeconds,
			Description:  "Grace period observed after signalling a VirtualMachineInstance to stop after which the VirtualMachineInstance is force terminated.",
		},
		"tolerations": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "If toleration is specified, obey all the toleration rules.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"effect": {
						Type:        schema.TypeSet,
						Optional:    true,
						Description: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
						Elem: &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{
								"NoSchedule",
								"PreferNoSchedule",
								"NoExecute",
							}, false),
						},
						Set: schema.HashString,
					},
					"key": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
					},
					"operator": {
						Type:        schema.TypeSet,
						Optional:    true,
						Default:     "Equal",
						Description: "Operator represents a key’s relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
						Elem: &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{
								"Exists",
								"Equal",
							}, false),
						},
						Set: schema.HashString,
					},
					"toleration_seconds": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
					},
					"value": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
					},
				},
			},
		},
		"volumes": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "List of volumes that can be mounted by disks belonging to the vmi.",
			Elem:        volumesSchema,
		},
	}
}

func volumesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_init_no_cloud": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of volumes that can be mounted by disks belonging to the vmi.",
			},
			"config_map": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of volumes that can be mounted by disks belonging to the vmi.",
			},
			"container_disk": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of volumes that can be mounted by disks belonging to the vmi.",
			},
			"data_volume": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of volumes that can be mounted by disks belonging to the vmi.",
			},
			"empty_disk": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of volumes that can be mounted by disks belonging to the vmi.",
			},
			"ephemeral": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of volumes that can be mounted by disks belonging to the vmi.",
			},
			"host_disk": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of volumes that can be mounted by disks belonging to the vmi.",
			},
			"name": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of volumes that can be mounted by disks belonging to the vmi.",
			},
			"persistent_volume_claim": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of volumes that can be mounted by disks belonging to the vmi.",
			},
			"secret": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of volumes that can be mounted by disks belonging to the vmi.",
			},
			"service_account": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of volumes that can be mounted by disks belonging to the vmi.",
			},
		},
	}
}

func probeSchema() *schema.Resource {
	h := handlerFields()
	h["failure_threshold"] = &schema.Schema{
		Type:         schema.TypeInt,
		Optional:     true,
		Description:  "Minimum consecutive failures for the probe to be considered failed after having succeeded.",
		Default:      3,
		ValidateFunc: validatePositiveInteger,
	}
	h["initial_delay_seconds"] = &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Description: "Number of seconds after the container has started before liveness probes are initiated. More info: http://kubernetes.io/docs/user-guide/pod-states#container-probes",
	}
	h["period_seconds"] = &schema.Schema{
		Type:         schema.TypeInt,
		Optional:     true,
		Default:      10,
		ValidateFunc: validatePositiveInteger,
		Description:  "How often (in seconds) to perform the probe",
	}
	h["success_threshold"] = &schema.Schema{
		Type:         schema.TypeInt,
		Optional:     true,
		Default:      1,
		ValidateFunc: validatePositiveInteger,
		Description:  "Minimum consecutive successes for the probe to be considered successful after having failed.",
	}

	h["timeout_seconds"] = &schema.Schema{
		Type:         schema.TypeInt,
		Optional:     true,
		Default:      1,
		ValidateFunc: validatePositiveInteger,
		Description:  "Number of seconds after which the probe times out. More info: http://kubernetes.io/docs/user-guide/pod-states#container-probes",
	}

	return &schema.Resource{
		Schema: h,
	}
}

func handlerFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"http_get": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Specifies the http request to perform.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"host": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: `Host name to connect to, defaults to the pod IP. You probably want to set "Host" in httpHeaders instead.`,
					},
					"path": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: `Path to access on the HTTP server.`,
					},
					"scheme": {
						Type:        schema.TypeString,
						Optional:    true,
						Default:     "HTTP",
						Description: `Scheme to use for connecting to the host.`,
					},
					"http_header": {
						Type:        schema.TypeList,
						Optional:    true,
						Description: `Scheme to use for connecting to the host.`,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"name": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "The header field name",
								},
								"value": {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "The header field value",
								},
							},
						},
					},
				},
			},
		},
		"tcp_socket": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"host": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Optional: Host name to connect to, defaults to the pod IP.",
					},
				},
			},
		},
	}
}

func affinityFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"node_affinity": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Node affinity is a group of node affinity scheduling rules.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"preferred_during_scheduling_ignored_during_execution": {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding \"weight\" to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"preference": {
									Type:        schema.TypeList,
									Required:    true,
									Description: "A node selector term, associated with the corresponding weight.",
									Elem: &schema.Resource{
										Schema: nodeSelectorTermFields(),
									},
								},
								"weight": {
									Type:         schema.TypeInt,
									Required:     true,
									ValidateFunc: validatePositiveInteger,
									Description:  "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
								},
							},
						},
					},
					"required_during_scheduling_ignored_during_execution": {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"node_selector_terms": {
									Type:        schema.TypeList,
									Required:    true,
									Description: "Required. A list of node selector terms. The terms are ORed.",
									Elem: &schema.Resource{
										Schema: nodeSelectorTermFields(),
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

func nodeSelectorTermFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"match_expressions": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "A list of node selector requirements by node’s labels.",
			Elem: &schema.Resource{
				Schema: nodeSelectorRequirementFields(),
			},
		},
		"match_fields": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "A list of node selector requirements by node’s fields.",
			Elem: &schema.Resource{
				Schema: nodeSelectorRequirementFields(),
			},
		},
	}
}

func nodeSelectorRequirementFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"key": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The label key that the selector applies to.",
		},
		"operator": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Represents a key’s relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
		},
		"values": {
			Type:        schema.TypeSet,
			Description: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Set:         schema.HashString,
		},
	}
}
