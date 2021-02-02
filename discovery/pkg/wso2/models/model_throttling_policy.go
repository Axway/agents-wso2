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

// ThrottlingPolicy struct for ThrottlingPolicy
type ThrottlingPolicy struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	PolicyLevel *string `json:"policyLevel,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	// Custom attributes added to the policy policy 
	Attributes *map[string]string `json:"attributes,omitempty"`
	// Maximum number of requests which can be sent within a provided unit time 
	RequestCount int64 `json:"requestCount"`
	UnitTime int64 `json:"unitTime"`
	TimeUnit *string `json:"timeUnit,omitempty"`
	// This attribute declares whether this policy is available under commercial or free 
	TierPlan string `json:"tierPlan"`
	// By making this attribute to false, you are capabale of sending requests even if the request count exceeded within a unit time 
	StopOnQuotaReach bool `json:"stopOnQuotaReach"`
	// Properties of a tier plan which are related to monetization
	MonetizationProperties *map[string]string `json:"monetizationProperties,omitempty"`
}

// NewThrottlingPolicy instantiates a new ThrottlingPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThrottlingPolicy(name string, requestCount int64, unitTime int64, tierPlan string, stopOnQuotaReach bool, ) *ThrottlingPolicy {
	this := ThrottlingPolicy{}
	this.Name = name
	this.RequestCount = requestCount
	this.UnitTime = unitTime
	this.TierPlan = tierPlan
	this.StopOnQuotaReach = stopOnQuotaReach
	return &this
}

// NewThrottlingPolicyWithDefaults instantiates a new ThrottlingPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThrottlingPolicyWithDefaults() *ThrottlingPolicy {
	this := ThrottlingPolicy{}
	return &this
}

// GetName returns the Name field value
func (o *ThrottlingPolicy) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ThrottlingPolicy) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ThrottlingPolicy) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ThrottlingPolicy) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThrottlingPolicy) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ThrottlingPolicy) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ThrottlingPolicy) SetDescription(v string) {
	o.Description = &v
}

// GetPolicyLevel returns the PolicyLevel field value if set, zero value otherwise.
func (o *ThrottlingPolicy) GetPolicyLevel() string {
	if o == nil || o.PolicyLevel == nil {
		var ret string
		return ret
	}
	return *o.PolicyLevel
}

// GetPolicyLevelOk returns a tuple with the PolicyLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThrottlingPolicy) GetPolicyLevelOk() (*string, bool) {
	if o == nil || o.PolicyLevel == nil {
		return nil, false
	}
	return o.PolicyLevel, true
}

// HasPolicyLevel returns a boolean if a field has been set.
func (o *ThrottlingPolicy) HasPolicyLevel() bool {
	if o != nil && o.PolicyLevel != nil {
		return true
	}

	return false
}

// SetPolicyLevel gets a reference to the given string and assigns it to the PolicyLevel field.
func (o *ThrottlingPolicy) SetPolicyLevel(v string) {
	o.PolicyLevel = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ThrottlingPolicy) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThrottlingPolicy) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ThrottlingPolicy) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ThrottlingPolicy) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ThrottlingPolicy) GetAttributes() map[string]string {
	if o == nil || o.Attributes == nil {
		var ret map[string]string
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThrottlingPolicy) GetAttributesOk() (*map[string]string, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ThrottlingPolicy) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]string and assigns it to the Attributes field.
func (o *ThrottlingPolicy) SetAttributes(v map[string]string) {
	o.Attributes = &v
}

// GetRequestCount returns the RequestCount field value
func (o *ThrottlingPolicy) GetRequestCount() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.RequestCount
}

// GetRequestCountOk returns a tuple with the RequestCount field value
// and a boolean to check if the value has been set.
func (o *ThrottlingPolicy) GetRequestCountOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestCount, true
}

// SetRequestCount sets field value
func (o *ThrottlingPolicy) SetRequestCount(v int64) {
	o.RequestCount = v
}

// GetUnitTime returns the UnitTime field value
func (o *ThrottlingPolicy) GetUnitTime() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.UnitTime
}

// GetUnitTimeOk returns a tuple with the UnitTime field value
// and a boolean to check if the value has been set.
func (o *ThrottlingPolicy) GetUnitTimeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UnitTime, true
}

// SetUnitTime sets field value
func (o *ThrottlingPolicy) SetUnitTime(v int64) {
	o.UnitTime = v
}

// GetTimeUnit returns the TimeUnit field value if set, zero value otherwise.
func (o *ThrottlingPolicy) GetTimeUnit() string {
	if o == nil || o.TimeUnit == nil {
		var ret string
		return ret
	}
	return *o.TimeUnit
}

// GetTimeUnitOk returns a tuple with the TimeUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThrottlingPolicy) GetTimeUnitOk() (*string, bool) {
	if o == nil || o.TimeUnit == nil {
		return nil, false
	}
	return o.TimeUnit, true
}

// HasTimeUnit returns a boolean if a field has been set.
func (o *ThrottlingPolicy) HasTimeUnit() bool {
	if o != nil && o.TimeUnit != nil {
		return true
	}

	return false
}

// SetTimeUnit gets a reference to the given string and assigns it to the TimeUnit field.
func (o *ThrottlingPolicy) SetTimeUnit(v string) {
	o.TimeUnit = &v
}

// GetTierPlan returns the TierPlan field value
func (o *ThrottlingPolicy) GetTierPlan() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.TierPlan
}

// GetTierPlanOk returns a tuple with the TierPlan field value
// and a boolean to check if the value has been set.
func (o *ThrottlingPolicy) GetTierPlanOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TierPlan, true
}

// SetTierPlan sets field value
func (o *ThrottlingPolicy) SetTierPlan(v string) {
	o.TierPlan = v
}

// GetStopOnQuotaReach returns the StopOnQuotaReach field value
func (o *ThrottlingPolicy) GetStopOnQuotaReach() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.StopOnQuotaReach
}

// GetStopOnQuotaReachOk returns a tuple with the StopOnQuotaReach field value
// and a boolean to check if the value has been set.
func (o *ThrottlingPolicy) GetStopOnQuotaReachOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StopOnQuotaReach, true
}

// SetStopOnQuotaReach sets field value
func (o *ThrottlingPolicy) SetStopOnQuotaReach(v bool) {
	o.StopOnQuotaReach = v
}

// GetMonetizationProperties returns the MonetizationProperties field value if set, zero value otherwise.
func (o *ThrottlingPolicy) GetMonetizationProperties() map[string]string {
	if o == nil || o.MonetizationProperties == nil {
		var ret map[string]string
		return ret
	}
	return *o.MonetizationProperties
}

// GetMonetizationPropertiesOk returns a tuple with the MonetizationProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThrottlingPolicy) GetMonetizationPropertiesOk() (*map[string]string, bool) {
	if o == nil || o.MonetizationProperties == nil {
		return nil, false
	}
	return o.MonetizationProperties, true
}

// HasMonetizationProperties returns a boolean if a field has been set.
func (o *ThrottlingPolicy) HasMonetizationProperties() bool {
	if o != nil && o.MonetizationProperties != nil {
		return true
	}

	return false
}

// SetMonetizationProperties gets a reference to the given map[string]string and assigns it to the MonetizationProperties field.
func (o *ThrottlingPolicy) SetMonetizationProperties(v map[string]string) {
	o.MonetizationProperties = &v
}

func (o ThrottlingPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.PolicyLevel != nil {
		toSerialize["policyLevel"] = o.PolicyLevel
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["requestCount"] = o.RequestCount
	}
	if true {
		toSerialize["unitTime"] = o.UnitTime
	}
	if o.TimeUnit != nil {
		toSerialize["timeUnit"] = o.TimeUnit
	}
	if true {
		toSerialize["tierPlan"] = o.TierPlan
	}
	if true {
		toSerialize["stopOnQuotaReach"] = o.StopOnQuotaReach
	}
	if o.MonetizationProperties != nil {
		toSerialize["monetizationProperties"] = o.MonetizationProperties
	}
	return json.Marshal(toSerialize)
}

type NullableThrottlingPolicy struct {
	value *ThrottlingPolicy
	isSet bool
}

func (v NullableThrottlingPolicy) Get() *ThrottlingPolicy {
	return v.value
}

func (v *NullableThrottlingPolicy) Set(val *ThrottlingPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableThrottlingPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableThrottlingPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThrottlingPolicy(val *ThrottlingPolicy) *NullableThrottlingPolicy {
	return &NullableThrottlingPolicy{value: val, isSet: true}
}

func (v NullableThrottlingPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThrottlingPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

