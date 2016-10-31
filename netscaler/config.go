/*
Copyright 2016 Citrix Systems, Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package netscaler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

func (c *NitroClient) BindResource(bindToResourceType string, bindToResourceName string, bindingResourceName string, bindingResourceType string, bindingStruct interface{}) error {
	if c.ResourceExists(bindToResourceType, bindToResourceName) == false {
		return fmt.Errorf("BindTo Resource %s of type %s does not exist", bindToResourceType, bindToResourceName)
	}

	if c.ResourceExists(bindingResourceType, bindingResourceName) == false {
		return fmt.Errorf("Binding Resource %s of type %s does not exist", bindingResourceType, bindingResourceName)
	}
	binding_name := bindToResourceType + "_" + bindingResourceType + "_binding"
	nsBinding := make(map[string]interface{})
	nsBinding[binding_name] = bindingStruct

	resourceJson, err := json.Marshal(nsBinding)

	body, err := c.createResource(binding_name, resourceJson)
	if err != nil {
		log.Fatal("Failed to bind resource %s to resource %s, err=%s", bindToResourceName, bindingResourceName, err)
		return err
	}
	_ = body
	return nil
}

func (c *NitroClient) AddResource(resourceType string, name string, resourceStruct interface{}) (string, error) {

	if c.ResourceExists(resourceType, name) == false {

		nsResource := make(map[string]interface{})
		nsResource[resourceType] = resourceStruct

		resourceJson, err := json.Marshal(nsResource)

		log.Println("Resourcejson is " + string(resourceJson))

		body, err := c.createResource(resourceType, resourceJson)
		if err != nil {
			log.Fatal("Failed to create resource of type %s, name=%s, err=%s", resourceType, name, err)
			return "", err
		}
		_ = body
	}

	return name, nil
}

func (c *NitroClient) UpdateResource(resourceType string, name string, resourceStruct interface{}) (string, error) {

	if c.ResourceExists(resourceType, name) == true {
		nsResource := make(map[string]interface{})
		nsResource[resourceType] = resourceStruct
		resourceJson, err := json.Marshal(nsResource)

		log.Println("Resourcejson is " + string(resourceJson))

		body, err := c.updateResource(resourceType, name, resourceJson)
		if err != nil {
			log.Fatal(fmt.Sprintf("Failed to update resource of type %s, name=%s err=%s", resourceType, name, err))
			return "", err
		}
		_ = body
	}

	return name, nil
}

func (c *NitroClient) DeleteResource(resourceType string, resourceName string) error {

	_, err := c.listResource(resourceType, resourceName)
	if err == nil { // resource exists
		log.Printf("Found resource of type %s: %s", resourceType, resourceName)
		_, err = c.deleteResource(resourceType, resourceName)
		if err != nil {
			log.Println(fmt.Sprintf("Failed to delete resourceType %s: %s, err=%s", resourceType, resourceName, err))
			return err
		}
	} else {
		log.Printf("Resource %s already deleted ", resourceName)
	}
	return nil
}

func (c *NitroClient) UnbindResource(boundToResourceType string, boundToResourceName string, boundResourceType string, boundResourceName string, bindingFilterName string) error {

	if c.ResourceExists(boundToResourceType, boundToResourceName) == false {
		log.Println(fmt.Sprintf("Unbind: BoundTo Resource %s of type %s does not exist", boundToResourceType, boundToResourceName))
		return nil
	}

	if c.ResourceExists(boundResourceType, boundResourceName) == false {
		log.Println("Unbind: Bound Resource %s of type %s does not exist", boundResourceType, boundResourceName)
		return nil
	}

	_, err := c.unbindResource(boundToResourceType, boundToResourceName, boundResourceType, boundResourceName, bindingFilterName)
	if err != nil {
		return fmt.Errorf("Failed to unbind  %s:%s from %s:%s, err=%s", boundResourceType, boundResourceName, boundToResourceType, boundToResourceName, err)
	}

	return nil
}

func (c *NitroClient) ResourceExists(resourceType string, resourceName string) bool {
	_, err := c.listResource(resourceType, resourceName)
	if err != nil {
		log.Printf("No %s %s found", resourceType, resourceName)
		return false
	}
	log.Printf("%s %s is already present", resourceType, resourceName)
	return true
}

func (c *NitroClient) FindResource(resourceType string, resourceName string) (map[string]interface{}, error) {
	var data map[string]interface{}
	result, err := c.listResource(resourceType, resourceName)
	if err != nil {
		log.Printf("No %s %s found", resourceType, resourceName)
		return data, errors.New(fmt.Sprintf("No resource %s of type %s found", resourceName, resourceType))
	}
	if err = json.Unmarshal(result, &data); err != nil {
		log.Println("Failed to unmarshal Netscaler Response!")
		return data, errors.New(fmt.Sprintf("Failed to unmarshal Netscaler Response:resource %s of type %s", resourceName, resourceType))
	}
	if data[resourceType] == nil {
		log.Printf("No %s %s found", resourceType, resourceName)
		return data, errors.New(fmt.Sprintf("No resource %s of type %s found", resourceName, resourceType))
	}
	resource := data[resourceType].([]interface{})[0] //only one resource obviously

	return resource.(map[string]interface{}), nil
}

func (c *NitroClient) ResourceBindingExists(resourceName string, resourceType string, boundResourceType string, boundResourceFilterName string, boundResourceFilterValue string) bool {
	result, err := c.listBoundResources(resourceName, resourceType, boundResourceType, boundResourceFilterName, boundResourceFilterValue)
	if err != nil {
		log.Printf("No %s %s to %s %s binding found", resourceType, resourceName, boundResourceType, boundResourceFilterValue)
		return false
	}

	var data map[string]interface{}
	if err := json.Unmarshal(result, &data); err != nil {
		log.Println("Failed to unmarshal Netscaler Response!")
		return false
	}
	if data[fmt.Sprintf("%s_%s_binding", resourceType, boundResourceType)] == nil {
		return false
	}

	log.Printf("%s of type  %s is bound to %s type and name %s", resourceType, resourceName, boundResourceType, boundResourceFilterValue)
	return true
}
