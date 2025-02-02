/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the POSTSkuListPromotionRules201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTSkuListPromotionRules201ResponseDataRelationships{}

// POSTSkuListPromotionRules201ResponseDataRelationships struct for POSTSkuListPromotionRules201ResponseDataRelationships
type POSTSkuListPromotionRules201ResponseDataRelationships struct {
	Promotion *POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion `json:"promotion,omitempty"`
	Versions  *POSTAddresses201ResponseDataRelationshipsVersions                  `json:"versions,omitempty"`
	SkuList   *POSTBundles201ResponseDataRelationshipsSkuList                     `json:"sku_list,omitempty"`
	Skus      *POSTBundles201ResponseDataRelationshipsSkus                        `json:"skus,omitempty"`
}

// NewPOSTSkuListPromotionRules201ResponseDataRelationships instantiates a new POSTSkuListPromotionRules201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTSkuListPromotionRules201ResponseDataRelationships() *POSTSkuListPromotionRules201ResponseDataRelationships {
	this := POSTSkuListPromotionRules201ResponseDataRelationships{}
	return &this
}

// NewPOSTSkuListPromotionRules201ResponseDataRelationshipsWithDefaults instantiates a new POSTSkuListPromotionRules201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTSkuListPromotionRules201ResponseDataRelationshipsWithDefaults() *POSTSkuListPromotionRules201ResponseDataRelationships {
	this := POSTSkuListPromotionRules201ResponseDataRelationships{}
	return &this
}

// GetPromotion returns the Promotion field value if set, zero value otherwise.
func (o *POSTSkuListPromotionRules201ResponseDataRelationships) GetPromotion() POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion {
	if o == nil || IsNil(o.Promotion) {
		var ret POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion
		return ret
	}
	return *o.Promotion
}

// GetPromotionOk returns a tuple with the Promotion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTSkuListPromotionRules201ResponseDataRelationships) GetPromotionOk() (*POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion, bool) {
	if o == nil || IsNil(o.Promotion) {
		return nil, false
	}
	return o.Promotion, true
}

// HasPromotion returns a boolean if a field has been set.
func (o *POSTSkuListPromotionRules201ResponseDataRelationships) HasPromotion() bool {
	if o != nil && !IsNil(o.Promotion) {
		return true
	}

	return false
}

// SetPromotion gets a reference to the given POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion and assigns it to the Promotion field.
func (o *POSTSkuListPromotionRules201ResponseDataRelationships) SetPromotion(v POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion) {
	o.Promotion = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *POSTSkuListPromotionRules201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret POSTAddresses201ResponseDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTSkuListPromotionRules201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *POSTSkuListPromotionRules201ResponseDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given POSTAddresses201ResponseDataRelationshipsVersions and assigns it to the Versions field.
func (o *POSTSkuListPromotionRules201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions) {
	o.Versions = &v
}

// GetSkuList returns the SkuList field value if set, zero value otherwise.
func (o *POSTSkuListPromotionRules201ResponseDataRelationships) GetSkuList() POSTBundles201ResponseDataRelationshipsSkuList {
	if o == nil || IsNil(o.SkuList) {
		var ret POSTBundles201ResponseDataRelationshipsSkuList
		return ret
	}
	return *o.SkuList
}

// GetSkuListOk returns a tuple with the SkuList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTSkuListPromotionRules201ResponseDataRelationships) GetSkuListOk() (*POSTBundles201ResponseDataRelationshipsSkuList, bool) {
	if o == nil || IsNil(o.SkuList) {
		return nil, false
	}
	return o.SkuList, true
}

// HasSkuList returns a boolean if a field has been set.
func (o *POSTSkuListPromotionRules201ResponseDataRelationships) HasSkuList() bool {
	if o != nil && !IsNil(o.SkuList) {
		return true
	}

	return false
}

// SetSkuList gets a reference to the given POSTBundles201ResponseDataRelationshipsSkuList and assigns it to the SkuList field.
func (o *POSTSkuListPromotionRules201ResponseDataRelationships) SetSkuList(v POSTBundles201ResponseDataRelationshipsSkuList) {
	o.SkuList = &v
}

// GetSkus returns the Skus field value if set, zero value otherwise.
func (o *POSTSkuListPromotionRules201ResponseDataRelationships) GetSkus() POSTBundles201ResponseDataRelationshipsSkus {
	if o == nil || IsNil(o.Skus) {
		var ret POSTBundles201ResponseDataRelationshipsSkus
		return ret
	}
	return *o.Skus
}

// GetSkusOk returns a tuple with the Skus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTSkuListPromotionRules201ResponseDataRelationships) GetSkusOk() (*POSTBundles201ResponseDataRelationshipsSkus, bool) {
	if o == nil || IsNil(o.Skus) {
		return nil, false
	}
	return o.Skus, true
}

// HasSkus returns a boolean if a field has been set.
func (o *POSTSkuListPromotionRules201ResponseDataRelationships) HasSkus() bool {
	if o != nil && !IsNil(o.Skus) {
		return true
	}

	return false
}

// SetSkus gets a reference to the given POSTBundles201ResponseDataRelationshipsSkus and assigns it to the Skus field.
func (o *POSTSkuListPromotionRules201ResponseDataRelationships) SetSkus(v POSTBundles201ResponseDataRelationshipsSkus) {
	o.Skus = &v
}

func (o POSTSkuListPromotionRules201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTSkuListPromotionRules201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Promotion) {
		toSerialize["promotion"] = o.Promotion
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	if !IsNil(o.SkuList) {
		toSerialize["sku_list"] = o.SkuList
	}
	if !IsNil(o.Skus) {
		toSerialize["skus"] = o.Skus
	}
	return toSerialize, nil
}

type NullablePOSTSkuListPromotionRules201ResponseDataRelationships struct {
	value *POSTSkuListPromotionRules201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTSkuListPromotionRules201ResponseDataRelationships) Get() *POSTSkuListPromotionRules201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTSkuListPromotionRules201ResponseDataRelationships) Set(val *POSTSkuListPromotionRules201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTSkuListPromotionRules201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTSkuListPromotionRules201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTSkuListPromotionRules201ResponseDataRelationships(val *POSTSkuListPromotionRules201ResponseDataRelationships) *NullablePOSTSkuListPromotionRules201ResponseDataRelationships {
	return &NullablePOSTSkuListPromotionRules201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTSkuListPromotionRules201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTSkuListPromotionRules201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
