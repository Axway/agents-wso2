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

// SecurityAuditAttribute struct for SecurityAuditAttribute
type SecurityAuditAttribute struct {
	IsGlobal *bool `json:"isGlobal,omitempty"`
	OverrideGlobal *bool `json:"overrideGlobal,omitempty"`
	ApiToken *string `json:"apiToken,omitempty"`
	CollectionId *string `json:"collectionId,omitempty"`
	BaseUrl *string `json:"baseUrl,omitempty"`
}

// NewSecurityAuditAttribute instantiates a new SecurityAuditAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAuditAttribute() *SecurityAuditAttribute {
	this := SecurityAuditAttribute{}
	return &this
}

// NewSecurityAuditAttributeWithDefaults instantiates a new SecurityAuditAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAuditAttributeWithDefaults() *SecurityAuditAttribute {
	this := SecurityAuditAttribute{}
	return &this
}

// GetIsGlobal returns the IsGlobal field value if set, zero value otherwise.
func (o *SecurityAuditAttribute) GetIsGlobal() bool {
	if o == nil || o.IsGlobal == nil {
		var ret bool
		return ret
	}
	return *o.IsGlobal
}

// GetIsGlobalOk returns a tuple with the IsGlobal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAuditAttribute) GetIsGlobalOk() (*bool, bool) {
	if o == nil || o.IsGlobal == nil {
		return nil, false
	}
	return o.IsGlobal, true
}

// HasIsGlobal returns a boolean if a field has been set.
func (o *SecurityAuditAttribute) HasIsGlobal() bool {
	if o != nil && o.IsGlobal != nil {
		return true
	}

	return false
}

// SetIsGlobal gets a reference to the given bool and assigns it to the IsGlobal field.
func (o *SecurityAuditAttribute) SetIsGlobal(v bool) {
	o.IsGlobal = &v
}

// GetOverrideGlobal returns the OverrideGlobal field value if set, zero value otherwise.
func (o *SecurityAuditAttribute) GetOverrideGlobal() bool {
	if o == nil || o.OverrideGlobal == nil {
		var ret bool
		return ret
	}
	return *o.OverrideGlobal
}

// GetOverrideGlobalOk returns a tuple with the OverrideGlobal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAuditAttribute) GetOverrideGlobalOk() (*bool, bool) {
	if o == nil || o.OverrideGlobal == nil {
		return nil, false
	}
	return o.OverrideGlobal, true
}

// HasOverrideGlobal returns a boolean if a field has been set.
func (o *SecurityAuditAttribute) HasOverrideGlobal() bool {
	if o != nil && o.OverrideGlobal != nil {
		return true
	}

	return false
}

// SetOverrideGlobal gets a reference to the given bool and assigns it to the OverrideGlobal field.
func (o *SecurityAuditAttribute) SetOverrideGlobal(v bool) {
	o.OverrideGlobal = &v
}

// GetApiToken returns the ApiToken field value if set, zero value otherwise.
func (o *SecurityAuditAttribute) GetApiToken() string {
	if o == nil || o.ApiToken == nil {
		var ret string
		return ret
	}
	return *o.ApiToken
}

// GetApiTokenOk returns a tuple with the ApiToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAuditAttribute) GetApiTokenOk() (*string, bool) {
	if o == nil || o.ApiToken == nil {
		return nil, false
	}
	return o.ApiToken, true
}

// HasApiToken returns a boolean if a field has been set.
func (o *SecurityAuditAttribute) HasApiToken() bool {
	if o != nil && o.ApiToken != nil {
		return true
	}

	return false
}

// SetApiToken gets a reference to the given string and assigns it to the ApiToken field.
func (o *SecurityAuditAttribute) SetApiToken(v string) {
	o.ApiToken = &v
}

// GetCollectionId returns the CollectionId field value if set, zero value otherwise.
func (o *SecurityAuditAttribute) GetCollectionId() string {
	if o == nil || o.CollectionId == nil {
		var ret string
		return ret
	}
	return *o.CollectionId
}

// GetCollectionIdOk returns a tuple with the CollectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAuditAttribute) GetCollectionIdOk() (*string, bool) {
	if o == nil || o.CollectionId == nil {
		return nil, false
	}
	return o.CollectionId, true
}

// HasCollectionId returns a boolean if a field has been set.
func (o *SecurityAuditAttribute) HasCollectionId() bool {
	if o != nil && o.CollectionId != nil {
		return true
	}

	return false
}

// SetCollectionId gets a reference to the given string and assigns it to the CollectionId field.
func (o *SecurityAuditAttribute) SetCollectionId(v string) {
	o.CollectionId = &v
}

// GetBaseUrl returns the BaseUrl field value if set, zero value otherwise.
func (o *SecurityAuditAttribute) GetBaseUrl() string {
	if o == nil || o.BaseUrl == nil {
		var ret string
		return ret
	}
	return *o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAuditAttribute) GetBaseUrlOk() (*string, bool) {
	if o == nil || o.BaseUrl == nil {
		return nil, false
	}
	return o.BaseUrl, true
}

// HasBaseUrl returns a boolean if a field has been set.
func (o *SecurityAuditAttribute) HasBaseUrl() bool {
	if o != nil && o.BaseUrl != nil {
		return true
	}

	return false
}

// SetBaseUrl gets a reference to the given string and assigns it to the BaseUrl field.
func (o *SecurityAuditAttribute) SetBaseUrl(v string) {
	o.BaseUrl = &v
}

func (o SecurityAuditAttribute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsGlobal != nil {
		toSerialize["isGlobal"] = o.IsGlobal
	}
	if o.OverrideGlobal != nil {
		toSerialize["overrideGlobal"] = o.OverrideGlobal
	}
	if o.ApiToken != nil {
		toSerialize["apiToken"] = o.ApiToken
	}
	if o.CollectionId != nil {
		toSerialize["collectionId"] = o.CollectionId
	}
	if o.BaseUrl != nil {
		toSerialize["baseUrl"] = o.BaseUrl
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityAuditAttribute struct {
	value *SecurityAuditAttribute
	isSet bool
}

func (v NullableSecurityAuditAttribute) Get() *SecurityAuditAttribute {
	return v.value
}

func (v *NullableSecurityAuditAttribute) Set(val *SecurityAuditAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAuditAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAuditAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAuditAttribute(val *SecurityAuditAttribute) *NullableSecurityAuditAttribute {
	return &NullableSecurityAuditAttribute{value: val, isSet: true}
}

func (v NullableSecurityAuditAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAuditAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


