package kubevirt

import (
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceKubevirtVirtualMachine() *schema.Resource {
	return &schema.Resource{
		Create: resourceKubevirtVirtualMachineCreate,
		Read:   resourceKubevirtVirtualMachineRead,
		Update: resourceKubevirtVirtualMachineUpdate,
		Delete: resourceKubevirtVirtualMachineDelete,
		Exists: resourceKubevirtVirtualMachineExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"metadata": namespacedMetadataSchema("virtualmachine", true),
			"spec": {
				Type:        schema.TypeList,
				Description: "Specification of the desired behavior of the virtual machine.",
				Required:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: virtualMachineSpecFields(),
				},
			},
		},
	}
}

func resourceKubevirtVirtualMachineCreate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceKubevirtVirtualMachineRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceKubevirtVirtualMachineUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceKubevirtVirtualMachineDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceKubevirtVirtualMachineExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	return false, nil
}
