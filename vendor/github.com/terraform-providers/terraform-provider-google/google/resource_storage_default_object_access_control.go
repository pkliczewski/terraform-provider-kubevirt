// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceStorageDefaultObjectAccessControl() *schema.Resource {
	return &schema.Resource{
		Create: resourceStorageDefaultObjectAccessControlCreate,
		Read:   resourceStorageDefaultObjectAccessControlRead,
		Update: resourceStorageDefaultObjectAccessControlUpdate,
		Delete: resourceStorageDefaultObjectAccessControlDelete,

		Importer: &schema.ResourceImporter{
			State: resourceStorageDefaultObjectAccessControlImport,
		},

		Schema: map[string]*schema.Schema{
			"bucket": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"entity": {
				Type:     schema.TypeString,
				Required: true,
			},
			"role": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"OWNER", "READER"}, false),
			},
			"object": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"email": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"generation": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"project_team": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"project_number": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"team": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"editors", "owners", "viewers", ""}, false),
						},
					},
				},
			},
		},
	}
}

func resourceStorageDefaultObjectAccessControlCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	bucketProp, err := expandStorageDefaultObjectAccessControlBucket(d.Get("bucket"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("bucket"); !isEmptyValue(reflect.ValueOf(bucketProp)) && (ok || !reflect.DeepEqual(v, bucketProp)) {
		obj["bucket"] = bucketProp
	}
	entityProp, err := expandStorageDefaultObjectAccessControlEntity(d.Get("entity"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("entity"); !isEmptyValue(reflect.ValueOf(entityProp)) && (ok || !reflect.DeepEqual(v, entityProp)) {
		obj["entity"] = entityProp
	}
	objectProp, err := expandStorageDefaultObjectAccessControlObject(d.Get("object"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("object"); !isEmptyValue(reflect.ValueOf(objectProp)) && (ok || !reflect.DeepEqual(v, objectProp)) {
		obj["object"] = objectProp
	}
	roleProp, err := expandStorageDefaultObjectAccessControlRole(d.Get("role"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("role"); !isEmptyValue(reflect.ValueOf(roleProp)) && (ok || !reflect.DeepEqual(v, roleProp)) {
		obj["role"] = roleProp
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/storage/v1/b/{{bucket}}/defaultObjectAcl")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new DefaultObjectAccessControl: %#v", obj)
	res, err := sendRequest(config, "POST", url, obj)
	if err != nil {
		return fmt.Errorf("Error creating DefaultObjectAccessControl: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{bucket}}/{{entity}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating DefaultObjectAccessControl %q: %#v", d.Id(), res)

	return resourceStorageDefaultObjectAccessControlRead(d, meta)
}

func resourceStorageDefaultObjectAccessControlRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://www.googleapis.com/storage/v1/b/{{bucket}}/defaultObjectAcl/{{entity}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("StorageDefaultObjectAccessControl %q", d.Id()))
	}

	if err := d.Set("domain", flattenStorageDefaultObjectAccessControlDomain(res["domain"])); err != nil {
		return fmt.Errorf("Error reading DefaultObjectAccessControl: %s", err)
	}
	if err := d.Set("email", flattenStorageDefaultObjectAccessControlEmail(res["email"])); err != nil {
		return fmt.Errorf("Error reading DefaultObjectAccessControl: %s", err)
	}
	if err := d.Set("entity", flattenStorageDefaultObjectAccessControlEntity(res["entity"])); err != nil {
		return fmt.Errorf("Error reading DefaultObjectAccessControl: %s", err)
	}
	if err := d.Set("entity_id", flattenStorageDefaultObjectAccessControlEntityId(res["entityId"])); err != nil {
		return fmt.Errorf("Error reading DefaultObjectAccessControl: %s", err)
	}
	if err := d.Set("generation", flattenStorageDefaultObjectAccessControlGeneration(res["generation"])); err != nil {
		return fmt.Errorf("Error reading DefaultObjectAccessControl: %s", err)
	}
	if err := d.Set("object", flattenStorageDefaultObjectAccessControlObject(res["object"])); err != nil {
		return fmt.Errorf("Error reading DefaultObjectAccessControl: %s", err)
	}
	if err := d.Set("project_team", flattenStorageDefaultObjectAccessControlProjectTeam(res["projectTeam"])); err != nil {
		return fmt.Errorf("Error reading DefaultObjectAccessControl: %s", err)
	}
	if err := d.Set("role", flattenStorageDefaultObjectAccessControlRole(res["role"])); err != nil {
		return fmt.Errorf("Error reading DefaultObjectAccessControl: %s", err)
	}

	return nil
}

func resourceStorageDefaultObjectAccessControlUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	bucketProp, err := expandStorageDefaultObjectAccessControlBucket(d.Get("bucket"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("bucket"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, bucketProp)) {
		obj["bucket"] = bucketProp
	}
	entityProp, err := expandStorageDefaultObjectAccessControlEntity(d.Get("entity"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("entity"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, entityProp)) {
		obj["entity"] = entityProp
	}
	objectProp, err := expandStorageDefaultObjectAccessControlObject(d.Get("object"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("object"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, objectProp)) {
		obj["object"] = objectProp
	}
	roleProp, err := expandStorageDefaultObjectAccessControlRole(d.Get("role"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("role"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, roleProp)) {
		obj["role"] = roleProp
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/storage/v1/b/{{bucket}}/defaultObjectAcl/{{entity}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating DefaultObjectAccessControl %q: %#v", d.Id(), obj)
	_, err = sendRequest(config, "PUT", url, obj)

	if err != nil {
		return fmt.Errorf("Error updating DefaultObjectAccessControl %q: %s", d.Id(), err)
	}

	return resourceStorageDefaultObjectAccessControlRead(d, meta)
}

func resourceStorageDefaultObjectAccessControlDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://www.googleapis.com/storage/v1/b/{{bucket}}/defaultObjectAcl/{{entity}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting DefaultObjectAccessControl %q", d.Id())
	res, err := sendRequest(config, "DELETE", url, obj)
	if err != nil {
		return handleNotFoundError(err, d, "DefaultObjectAccessControl")
	}

	log.Printf("[DEBUG] Finished deleting DefaultObjectAccessControl %q: %#v", d.Id(), res)
	return nil
}

func resourceStorageDefaultObjectAccessControlImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	parseImportId([]string{"(?P<bucket>[^/]+)/(?P<entity>[^/]+)"}, d, config)

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{bucket}}/{{entity}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenStorageDefaultObjectAccessControlDomain(v interface{}) interface{} {
	return v
}

func flattenStorageDefaultObjectAccessControlEmail(v interface{}) interface{} {
	return v
}

func flattenStorageDefaultObjectAccessControlEntity(v interface{}) interface{} {
	return v
}

func flattenStorageDefaultObjectAccessControlEntityId(v interface{}) interface{} {
	return v
}

func flattenStorageDefaultObjectAccessControlGeneration(v interface{}) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		} // let terraform core handle it if we can't convert the string to an int.
	}
	return v
}

func flattenStorageDefaultObjectAccessControlObject(v interface{}) interface{} {
	return v
}

func flattenStorageDefaultObjectAccessControlProjectTeam(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	transformed := make(map[string]interface{})
	transformed["project_number"] =
		flattenStorageDefaultObjectAccessControlProjectTeamProjectNumber(original["projectNumber"])
	transformed["team"] =
		flattenStorageDefaultObjectAccessControlProjectTeamTeam(original["team"])
	return []interface{}{transformed}
}
func flattenStorageDefaultObjectAccessControlProjectTeamProjectNumber(v interface{}) interface{} {
	return v
}

func flattenStorageDefaultObjectAccessControlProjectTeamTeam(v interface{}) interface{} {
	return v
}

func flattenStorageDefaultObjectAccessControlRole(v interface{}) interface{} {
	return v
}

func expandStorageDefaultObjectAccessControlBucket(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandStorageDefaultObjectAccessControlEntity(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandStorageDefaultObjectAccessControlObject(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandStorageDefaultObjectAccessControlRole(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}
