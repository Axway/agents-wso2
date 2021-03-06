/*
 * WSO2 API Manager - Publisher API
 *
 * This document specifies a **RESTful API** for WSO2 **API Manager** - **Publisher**.  # Authentication Our REST APIs are protected using OAuth2 and access control is achieved through scopes. Before you start invoking the the API you need to obtain an access token with the required scopes. This guide will walk you through the steps that you will need to follow to obtain an access token. First you need to obtain the consumer key/secret key pair by calling the dynamic client registration (DCR) endpoint. You can add your preferred grant types in the payload. A Sample payload is shown below. ```   {   \"callbackUrl\":\"www.google.lk\",   \"clientName\":\"rest_api_publisher\",   \"owner\":\"admin\",   \"grantType\":\"client_credentials password refresh_token\",   \"saasApp\":true   } ``` Create a file (payload.json) with the above sample payload, and use the cURL shown bellow to invoke the DCR endpoint. Authorization header of this should contain the base64 encoded admin username and password. **Format of the request** ```   curl -X POST -H \"Authorization: Basic Base64(admin_username:admin_password)\" -H \"Content-Type: application/json\"   \\ -d @payload.json https://<host>:<servlet_port>/client-registration/v0.17/register ``` **Sample request** ```   curl -X POST -H \"Authorization: Basic YWRtaW46YWRtaW4=\" -H \"Content-Type: application/json\"   \\ -d @payload.json https://localhost:9443/client-registration/v0.17/register ``` Following is a sample response after invoking the above curl. ``` { \"clientId\": \"fOCi4vNJ59PpHucC2CAYfYuADdMa\", \"clientName\": \"rest_api_publisher\", \"callBackURL\": \"www.google.lk\", \"clientSecret\": \"a4FwHlq0iCIKVs2MPIIDnepZnYMa\", \"isSaasApplication\": true, \"appOwner\": \"admin\", \"jsonString\": \"{\\\"grant_types\\\":\\\"client_credentials password refresh_token\\\",\\\"redirect_uris\\\":\\\"www.google.lk\\\",\\\"client_name\\\":\\\"rest_api123\\\"}\", \"jsonAppAttribute\": \"{}\", \"tokenType\": null } ``` Next you must use the above client id and secret to obtain the access token. We will be using the password grant type for this, you can use any grant type you desire. You also need to add the proper **scope** when getting the access token. All possible scopes for publisher REST API can be viewed in **OAuth2 Security** section of this document and scope for each resource is given in **authorization** section of resource documentation. Following is the format of the request if you are using the password grant type. ``` curl -k -d \"grant_type=password&username=<admin_username>&password=<admin_passowrd&scope=<scopes seperated by space>\" \\ -H \"Authorization: Basic base64(cliet_id:client_secret)\" \\ https://<host>:<gateway_port>/token ``` **Sample request** ``` curl https://localhost:8243/token -k \\ -H \"Authorization: Basic Zk9DaTR2Tko1OVBwSHVjQzJDQVlmWXVBRGRNYTphNEZ3SGxxMGlDSUtWczJNUElJRG5lcFpuWU1h\" \\ -d \"grant_type=password&username=admin&password=admin&scope=apim:api_view apim:api_create\" ``` Shown below is a sample response to the above request. ``` { \"access_token\": \"e79bda48-3406-3178-acce-f6e4dbdcbb12\", \"refresh_token\": \"a757795d-e69f-38b8-bd85-9aded677a97c\", \"scope\": \"apim:api_create apim:api_view\", \"token_type\": \"Bearer\", \"expires_in\": 3600 } ``` Now you have a valid access token, which you can use to invoke an API. Navigate through the API descriptions to find the required API, obtain an access token as described above and invoke the API with the authentication header. If you use a different authentication mechanism, this process may change.  # Try out in Postman If you want to try-out the embedded postman collection with \"Run in Postman\" option, please follow the guidelines listed below. * All of the OAuth2 secured endpoints have been configured with an Authorization Bearer header with a parameterized access token. Before invoking any REST API resource make sure you run the `Register DCR Application` and `Generate Access Token` requests to fetch an access token with all required scopes. * Make sure you have an API Manager instance up and running. * Update the `basepath` parameter to match the hostname and port of the APIM instance.  [![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/a09044034b5c3c1b01a9) 
 *
 * API version: v1.1
 * Contact: architecture@wso2.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
)

// LifecycleState struct for LifecycleState
type LifecycleState struct {
	State *string `json:"state,omitempty"`
	CheckItems *[]LifecycleStateCheckItems `json:"checkItems,omitempty"`
	AvailableTransitions *[]LifecycleStateAvailableTransitions `json:"availableTransitions,omitempty"`
}

// NewLifecycleState instantiates a new LifecycleState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLifecycleState() *LifecycleState {
	this := LifecycleState{}
	return &this
}

// NewLifecycleStateWithDefaults instantiates a new LifecycleState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLifecycleStateWithDefaults() *LifecycleState {
	this := LifecycleState{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *LifecycleState) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifecycleState) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *LifecycleState) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *LifecycleState) SetState(v string) {
	o.State = &v
}

// GetCheckItems returns the CheckItems field value if set, zero value otherwise.
func (o *LifecycleState) GetCheckItems() []LifecycleStateCheckItems {
	if o == nil || o.CheckItems == nil {
		var ret []LifecycleStateCheckItems
		return ret
	}
	return *o.CheckItems
}

// GetCheckItemsOk returns a tuple with the CheckItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifecycleState) GetCheckItemsOk() (*[]LifecycleStateCheckItems, bool) {
	if o == nil || o.CheckItems == nil {
		return nil, false
	}
	return o.CheckItems, true
}

// HasCheckItems returns a boolean if a field has been set.
func (o *LifecycleState) HasCheckItems() bool {
	if o != nil && o.CheckItems != nil {
		return true
	}

	return false
}

// SetCheckItems gets a reference to the given []LifecycleStateCheckItems and assigns it to the CheckItems field.
func (o *LifecycleState) SetCheckItems(v []LifecycleStateCheckItems) {
	o.CheckItems = &v
}

// GetAvailableTransitions returns the AvailableTransitions field value if set, zero value otherwise.
func (o *LifecycleState) GetAvailableTransitions() []LifecycleStateAvailableTransitions {
	if o == nil || o.AvailableTransitions == nil {
		var ret []LifecycleStateAvailableTransitions
		return ret
	}
	return *o.AvailableTransitions
}

// GetAvailableTransitionsOk returns a tuple with the AvailableTransitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifecycleState) GetAvailableTransitionsOk() (*[]LifecycleStateAvailableTransitions, bool) {
	if o == nil || o.AvailableTransitions == nil {
		return nil, false
	}
	return o.AvailableTransitions, true
}

// HasAvailableTransitions returns a boolean if a field has been set.
func (o *LifecycleState) HasAvailableTransitions() bool {
	if o != nil && o.AvailableTransitions != nil {
		return true
	}

	return false
}

// SetAvailableTransitions gets a reference to the given []LifecycleStateAvailableTransitions and assigns it to the AvailableTransitions field.
func (o *LifecycleState) SetAvailableTransitions(v []LifecycleStateAvailableTransitions) {
	o.AvailableTransitions = &v
}

func (o LifecycleState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.CheckItems != nil {
		toSerialize["checkItems"] = o.CheckItems
	}
	if o.AvailableTransitions != nil {
		toSerialize["availableTransitions"] = o.AvailableTransitions
	}
	return json.Marshal(toSerialize)
}

type NullableLifecycleState struct {
	value *LifecycleState
	isSet bool
}

func (v NullableLifecycleState) Get() *LifecycleState {
	return v.value
}

func (v *NullableLifecycleState) Set(val *LifecycleState) {
	v.value = val
	v.isSet = true
}

func (v NullableLifecycleState) IsSet() bool {
	return v.isSet
}

func (v *NullableLifecycleState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLifecycleState(val *LifecycleState) *NullableLifecycleState {
	return &NullableLifecycleState{value: val, isSet: true}
}

func (v NullableLifecycleState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLifecycleState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


