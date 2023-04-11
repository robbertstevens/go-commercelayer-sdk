/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the BundleDataRelationshipsSkuList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BundleDataRelationshipsSkuList{}

// BundleDataRelationshipsSkuList struct for BundleDataRelationshipsSkuList
type BundleDataRelationshipsSkuList struct {
	Data *POSTBundlesRequestDataRelationshipsSkuListData `json:"data,omitempty"`
}

// NewBundleDataRelationshipsSkuList instantiates a new BundleDataRelationshipsSkuList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleDataRelationshipsSkuList() *BundleDataRelationshipsSkuList {
	this := BundleDataRelationshipsSkuList{}
	return &this
}

// NewBundleDataRelationshipsSkuListWithDefaults instantiates a new BundleDataRelationshipsSkuList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleDataRelationshipsSkuListWithDefaults() *BundleDataRelationshipsSkuList {
	this := BundleDataRelationshipsSkuList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BundleDataRelationshipsSkuList) GetData() POSTBundlesRequestDataRelationshipsSkuListData {
	if o == nil || IsNil(o.Data) {
		var ret POSTBundlesRequestDataRelationshipsSkuListData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleDataRelationshipsSkuList) GetDataOk() (*POSTBundlesRequestDataRelationshipsSkuListData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BundleDataRelationshipsSkuList) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTBundlesRequestDataRelationshipsSkuListData and assigns it to the Data field.
func (o *BundleDataRelationshipsSkuList) SetData(v POSTBundlesRequestDataRelationshipsSkuListData) {
	o.Data = &v
}

func (o BundleDataRelationshipsSkuList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BundleDataRelationshipsSkuList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableBundleDataRelationshipsSkuList struct {
	value *BundleDataRelationshipsSkuList
	isSet bool
}

func (v NullableBundleDataRelationshipsSkuList) Get() *BundleDataRelationshipsSkuList {
	return v.value
}

func (v *NullableBundleDataRelationshipsSkuList) Set(val *BundleDataRelationshipsSkuList) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleDataRelationshipsSkuList) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleDataRelationshipsSkuList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleDataRelationshipsSkuList(val *BundleDataRelationshipsSkuList) *NullableBundleDataRelationshipsSkuList {
	return &NullableBundleDataRelationshipsSkuList{value: val, isSet: true}
}

func (v NullableBundleDataRelationshipsSkuList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleDataRelationshipsSkuList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
