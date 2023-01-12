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

func ResourceNetworkInterfaceElastic() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkInterfaceElasticCreate,
		Read:   resourceNetworkInterfaceElasticRead,
		Update: resourceNetworkInterfaceElasticUpdate,
		Delete: resourceNetworkInterfaceElasticDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				idParts := strings.Split(d.Id(), ":")
				if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
					return nil, fmt.Errorf("Unexpected format of ID (%q), expected ProjectID:NetworkInterfaceElasticID", d.Id())
				}
				projectID := idParts[0]
				networkInterfaceElasticID := idParts[1]
				d.SetId(networkInterfaceElasticID)
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
func resourceNetworkInterfaceElasticStateRefreshFunc(cli *client.Client, networkInterfaceElasticID string, projectID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		resp, httpResponse, _ := cli.VserverClient.NetworkInterfaceElasticRestControllerV2Api.GetNetworkInterfaceElasticUsingGET(context.TODO(), networkInterfaceElasticID, projectID)
		if httpResponse.StatusCode != http.StatusOK {
			return nil, "", fmt.Errorf("Error describing network interface elastic: %s", GetResponseBody(httpResponse))
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		networkInterfaceElastic := resp.Data
		return networkInterfaceElastic, networkInterfaceElastic.Status, nil
	}
}
func resourceNetworkInterfaceElasticCreate(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	networkInterfaceElastic := vserver.CreateNetworkInterfaceRequest{
		Name: d.Get("name").(string),
	}
	cli := m.(*client.Client)
	resp, httpResponse, err := cli.VserverClient.NetworkInterfaceElasticRestControllerV2Api.CreateNetworkInterfaceElasticUsingPOST(context.TODO(), networkInterfaceElastic, projectID)
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
		Pending:    networkInterfaceElasticCreating,
		Target:     networkInterfaceElasticCreated,
		Refresh:    resourceNetworkInterfaceElasticStateRefreshFunc(cli, resp.Data.Uuid, projectID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for create network interface elastic (%s) %s", resp.Data.Uuid, err)
	}
	d.SetId(resp.Data.Uuid)
	return resourceNetworkInterfaceElasticRead(d, m)
}

func resourceNetworkInterfaceElasticRead(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	networkInterfaceElasticID := d.Id()
	cli := m.(*client.Client)
	resp, httpResponse, _ := cli.VserverClient.NetworkInterfaceElasticRestControllerV2Api.GetNetworkInterfaceElasticUsingGET(context.TODO(), networkInterfaceElasticID, projectID)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	networkInterfaceElastic := resp.Data
	d.Set("name", networkInterfaceElastic.Name)
	return nil
}

func resourceNetworkInterfaceElasticUpdate(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	networkInterfaceElasticID := d.Id()
	if d.HasChange("name") {
		renameNetworkInterfaceElasticRequest := vserver.RenameNetworkInterfaceRequest{
			Name: d.Get("name").(string),
		}
		cli := m.(*client.Client)
		resp, httpResponse, _ := cli.VserverClient.NetworkInterfaceElasticRestControllerV2Api.RenameNetworkInterfaceElasticUsingPUT(context.TODO(), networkInterfaceElasticID, projectID, renameNetworkInterfaceElasticRequest)
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
	return resourceNetworkInterfaceElasticRead(d, m)

}

func resourceNetworkInterfaceElasticDelete(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	networkInterfaceElasticID := d.Id()
	cli := m.(*client.Client)
	httpResponse, _ := cli.VserverClient.NetworkInterfaceElasticRestControllerV2Api.DeleteNetworkInterfaceElasticUsingDELETE(context.TODO(), networkInterfaceElasticID, projectID)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}
	d.SetId("")
	return nil
}
