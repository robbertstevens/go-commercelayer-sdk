/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETSkuLists200ResponseDataInnerRelationships struct for GETSkuLists200ResponseDataInnerRelationships
type GETSkuLists200ResponseDataInnerRelationships struct {
	Customer     *GETCouponRecipients200ResponseDataInnerRelationshipsCustomer `json:"customer,omitempty"`
	Skus         *GETBundles200ResponseDataInnerRelationshipsSkus              `json:"skus,omitempty"`
	SkuListItems *GETSkuLists200ResponseDataInnerRelationshipsSkuListItems     `json:"sku_list_items,omitempty"`
	Bundles      *GETSkuLists200ResponseDataInnerRelationshipsBundles          `json:"bundles,omitempty"`
}

// NewGETSkuLists200ResponseDataInnerRelationships instantiates a new GETSkuLists200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETSkuLists200ResponseDataInnerRelationships() *GETSkuLists200ResponseDataInnerRelationships {
	this := GETSkuLists200ResponseDataInnerRelationships{}
	return &this
}

// NewGETSkuLists200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETSkuLists200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETSkuLists200ResponseDataInnerRelationshipsWithDefaults() *GETSkuLists200ResponseDataInnerRelationships {
	this := GETSkuLists200ResponseDataInnerRelationships{}
	return &this
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *GETSkuLists200ResponseDataInnerRelationships) GetCustomer() GETCouponRecipients200ResponseDataInnerRelationshipsCustomer {
	if o == nil || o.Customer == nil {
		var ret GETCouponRecipients200ResponseDataInnerRelationshipsCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkuLists200ResponseDataInnerRelationships) GetCustomerOk() (*GETCouponRecipients200ResponseDataInnerRelationshipsCustomer, bool) {
	if o == nil || o.Customer == nil {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *GETSkuLists200ResponseDataInnerRelationships) HasCustomer() bool {
	if o != nil && o.Customer != nil {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given GETCouponRecipients200ResponseDataInnerRelationshipsCustomer and assigns it to the Customer field.
func (o *GETSkuLists200ResponseDataInnerRelationships) SetCustomer(v GETCouponRecipients200ResponseDataInnerRelationshipsCustomer) {
	o.Customer = &v
}

// GetSkus returns the Skus field value if set, zero value otherwise.
func (o *GETSkuLists200ResponseDataInnerRelationships) GetSkus() GETBundles200ResponseDataInnerRelationshipsSkus {
	if o == nil || o.Skus == nil {
		var ret GETBundles200ResponseDataInnerRelationshipsSkus
		return ret
	}
	return *o.Skus
}

// GetSkusOk returns a tuple with the Skus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkuLists200ResponseDataInnerRelationships) GetSkusOk() (*GETBundles200ResponseDataInnerRelationshipsSkus, bool) {
	if o == nil || o.Skus == nil {
		return nil, false
	}
	return o.Skus, true
}

// HasSkus returns a boolean if a field has been set.
func (o *GETSkuLists200ResponseDataInnerRelationships) HasSkus() bool {
	if o != nil && o.Skus != nil {
		return true
	}

	return false
}

// SetSkus gets a reference to the given GETBundles200ResponseDataInnerRelationshipsSkus and assigns it to the Skus field.
func (o *GETSkuLists200ResponseDataInnerRelationships) SetSkus(v GETBundles200ResponseDataInnerRelationshipsSkus) {
	o.Skus = &v
}

// GetSkuListItems returns the SkuListItems field value if set, zero value otherwise.
func (o *GETSkuLists200ResponseDataInnerRelationships) GetSkuListItems() GETSkuLists200ResponseDataInnerRelationshipsSkuListItems {
	if o == nil || o.SkuListItems == nil {
		var ret GETSkuLists200ResponseDataInnerRelationshipsSkuListItems
		return ret
	}
	return *o.SkuListItems
}

// GetSkuListItemsOk returns a tuple with the SkuListItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkuLists200ResponseDataInnerRelationships) GetSkuListItemsOk() (*GETSkuLists200ResponseDataInnerRelationshipsSkuListItems, bool) {
	if o == nil || o.SkuListItems == nil {
		return nil, false
	}
	return o.SkuListItems, true
}

// HasSkuListItems returns a boolean if a field has been set.
func (o *GETSkuLists200ResponseDataInnerRelationships) HasSkuListItems() bool {
	if o != nil && o.SkuListItems != nil {
		return true
	}

	return false
}

// SetSkuListItems gets a reference to the given GETSkuLists200ResponseDataInnerRelationshipsSkuListItems and assigns it to the SkuListItems field.
func (o *GETSkuLists200ResponseDataInnerRelationships) SetSkuListItems(v GETSkuLists200ResponseDataInnerRelationshipsSkuListItems) {
	o.SkuListItems = &v
}

// GetBundles returns the Bundles field value if set, zero value otherwise.
func (o *GETSkuLists200ResponseDataInnerRelationships) GetBundles() GETSkuLists200ResponseDataInnerRelationshipsBundles {
	if o == nil || o.Bundles == nil {
		var ret GETSkuLists200ResponseDataInnerRelationshipsBundles
		return ret
	}
	return *o.Bundles
}

// GetBundlesOk returns a tuple with the Bundles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETSkuLists200ResponseDataInnerRelationships) GetBundlesOk() (*GETSkuLists200ResponseDataInnerRelationshipsBundles, bool) {
	if o == nil || o.Bundles == nil {
		return nil, false
	}
	return o.Bundles, true
}

// HasBundles returns a boolean if a field has been set.
func (o *GETSkuLists200ResponseDataInnerRelationships) HasBundles() bool {
	if o != nil && o.Bundles != nil {
		return true
	}

	return false
}

// SetBundles gets a reference to the given GETSkuLists200ResponseDataInnerRelationshipsBundles and assigns it to the Bundles field.
func (o *GETSkuLists200ResponseDataInnerRelationships) SetBundles(v GETSkuLists200ResponseDataInnerRelationshipsBundles) {
	o.Bundles = &v
}

func (o GETSkuLists200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Customer != nil {
		toSerialize["customer"] = o.Customer
	}
	if o.Skus != nil {
		toSerialize["skus"] = o.Skus
	}
	if o.SkuListItems != nil {
		toSerialize["sku_list_items"] = o.SkuListItems
	}
	if o.Bundles != nil {
		toSerialize["bundles"] = o.Bundles
	}
	return json.Marshal(toSerialize)
}

type NullableGETSkuLists200ResponseDataInnerRelationships struct {
	value *GETSkuLists200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETSkuLists200ResponseDataInnerRelationships) Get() *GETSkuLists200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETSkuLists200ResponseDataInnerRelationships) Set(val *GETSkuLists200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETSkuLists200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETSkuLists200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETSkuLists200ResponseDataInnerRelationships(val *GETSkuLists200ResponseDataInnerRelationships) *NullableGETSkuLists200ResponseDataInnerRelationships {
	return &NullableGETSkuLists200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETSkuLists200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETSkuLists200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
