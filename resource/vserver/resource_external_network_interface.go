package vserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vserver"
)

func ResourceExternalNetworkInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceExternalNetworkInterfaceCreate,
		Read:   resourceExternalNetworkInterfaceRead,
		Update: resourceExternalNetworkInterfaceUpdate,
		Delete: resourceExternalNetworkInterfaceDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				idParts := strings.Split(d.Id(), ":")
				if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
					return nil, fmt.Errorf("Unexpected format of ID (%q), expected ProjectID:ExternalNetworkInterfaceID", d.Id())
				}
				projectID := idParts[0]
				externalNetworkInterfaceID := idParts[1]
				d.SetId(externalNetworkInterfaceID)
				d.Set("project_id", projectID)
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
func resourceExternalNetworkInterfaceStateRefreshFunc(cli *client.Client, externalNetworkInterfaceID string, projectID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		resp, httpResponse, _ := cli.VserverClient.NetworkInterfaceElasticRestControllerV2Api.GetNetworkInterfaceElasticUsingGET(context.TODO(), externalNetworkInterfaceID, projectID)
		if httpResponse.StatusCode != http.StatusOK {
			return nil, "", fmt.Errorf("Error describing External network interface: %s", GetResponseBody(httpResponse))
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		externalNetworkInterface := resp.Data
		return externalNetworkInterface, externalNetworkInterface.Status, nil
	}
}
func resourceExternalNetworkInterfaceCreate(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	externalNetworkInterface := vserver.CreateNetworkInterfaceRequest{
		Name: d.Get("name").(string),
	}
	cli := m.(*client.Client)
	resp, httpResponse, err := cli.VserverClient.NetworkInterfaceElasticRestControllerV2Api.CreateNetworkInterfaceElasticUsingPOST(context.TODO(), externalNetworkInterface, projectID)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	stateConf := &resource.StateChangeConf{
		Pending:    externalNetworkInterfaceCreating,
		Target:     externalNetworkInterfaceCreated,
		Refresh:    resourceExternalNetworkInterfaceStateRefreshFunc(cli, resp.Data.Uuid, projectID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for create External network interface (%s) %s", resp.Data.Uuid, err)
	}
	d.SetId(resp.Data.Uuid)
	return resourceExternalNetworkInterfaceRead(d, m)
}

func resourceExternalNetworkInterfaceRead(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	externalNetworkInterfaceID := d.Id()
	cli := m.(*client.Client)
	resp, httpResponse, _ := cli.VserverClient.NetworkInterfaceElasticRestControllerV2Api.GetNetworkInterfaceElasticUsingGET(context.TODO(), externalNetworkInterfaceID, projectID)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	externalNetworkInterface := resp.Data
	d.Set("name", externalNetworkInterface.Name)
	return nil
}

func resourceExternalNetworkInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	externalNetworkInterfaceID := d.Id()
	if d.HasChange("name") {
		renameNetworkInterfaceElasticRequest := vserver.RenameNetworkInterfaceRequest{
			Name: d.Get("name").(string),
		}
		cli := m.(*client.Client)
		resp, httpResponse, _ := cli.VserverClient.NetworkInterfaceElasticRestControllerV2Api.RenameNetworkInterfaceElasticUsingPUT(context.TODO(), externalNetworkInterfaceID, projectID, renameNetworkInterfaceElasticRequest)
		if CheckErrorResponse(httpResponse) {
			responseBody := GetResponseBody(httpResponse)
			errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
			return errorResponse
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
	}
	return resourceExternalNetworkInterfaceRead(d, m)

}

func resourceExternalNetworkInterfaceDelete(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	externalNetworkInterfaceID := d.Id()
	cli := m.(*client.Client)
	httpResponse, _ := cli.VserverClient.NetworkInterfaceElasticRestControllerV2Api.DeleteNetworkInterfaceElasticUsingDELETE(context.TODO(), externalNetworkInterfaceID, projectID)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}
	d.SetId("")
	return nil
}
