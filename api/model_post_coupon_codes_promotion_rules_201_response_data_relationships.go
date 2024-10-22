/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the POSTCouponCodesPromotionRules201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTCouponCodesPromotionRules201ResponseDataRelationships{}

// POSTCouponCodesPromotionRules201ResponseDataRelationships struct for POSTCouponCodesPromotionRules201ResponseDataRelationships
type POSTCouponCodesPromotionRules201ResponseDataRelationships struct {
	Promotion *POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion `json:"promotion,omitempty"`
	Versions  *POSTAddresses201ResponseDataRelationshipsVersions                  `json:"versions,omitempty"`
	Coupons   *POSTBuyXPayYPromotions201ResponseDataRelationshipsCoupons          `json:"coupons,omitempty"`
}

// NewPOSTCouponCodesPromotionRules201ResponseDataRelationships instantiates a new POSTCouponCodesPromotionRules201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCouponCodesPromotionRules201ResponseDataRelationships() *POSTCouponCodesPromotionRules201ResponseDataRelationships {
	this := POSTCouponCodesPromotionRules201ResponseDataRelationships{}
	return &this
}

// NewPOSTCouponCodesPromotionRules201ResponseDataRelationshipsWithDefaults instantiates a new POSTCouponCodesPromotionRules201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCouponCodesPromotionRules201ResponseDataRelationshipsWithDefaults() *POSTCouponCodesPromotionRules201ResponseDataRelationships {
	this := POSTCouponCodesPromotionRules201ResponseDataRelationships{}
	return &this
}

// GetPromotion returns the Promotion field value if set, zero value otherwise.
func (o *POSTCouponCodesPromotionRules201ResponseDataRelationships) GetPromotion() POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion {
	if o == nil || IsNil(o.Promotion) {
		var ret POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion
		return ret
	}
	return *o.Promotion
}

// GetPromotionOk returns a tuple with the Promotion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCouponCodesPromotionRules201ResponseDataRelationships) GetPromotionOk() (*POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion, bool) {
	if o == nil || IsNil(o.Promotion) {
		return nil, false
	}
	return o.Promotion, true
}

// HasPromotion returns a boolean if a field has been set.
func (o *POSTCouponCodesPromotionRules201ResponseDataRelationships) HasPromotion() bool {
	if o != nil && !IsNil(o.Promotion) {
		return true
	}

	return false
}

// SetPromotion gets a reference to the given POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion and assigns it to the Promotion field.
func (o *POSTCouponCodesPromotionRules201ResponseDataRelationships) SetPromotion(v POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion) {
	o.Promotion = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *POSTCouponCodesPromotionRules201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret POSTAddresses201ResponseDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCouponCodesPromotionRules201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *POSTCouponCodesPromotionRules201ResponseDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given POSTAddresses201ResponseDataRelationshipsVersions and assigns it to the Versions field.
func (o *POSTCouponCodesPromotionRules201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions) {
	o.Versions = &v
}

// GetCoupons returns the Coupons field value if set, zero value otherwise.
func (o *POSTCouponCodesPromotionRules201ResponseDataRelationships) GetCoupons() POSTBuyXPayYPromotions201ResponseDataRelationshipsCoupons {
	if o == nil || IsNil(o.Coupons) {
		var ret POSTBuyXPayYPromotions201ResponseDataRelationshipsCoupons
		return ret
	}
	return *o.Coupons
}

// GetCouponsOk returns a tuple with the Coupons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCouponCodesPromotionRules201ResponseDataRelationships) GetCouponsOk() (*POSTBuyXPayYPromotions201ResponseDataRelationshipsCoupons, bool) {
	if o == nil || IsNil(o.Coupons) {
		return nil, false
	}
	return o.Coupons, true
}

// HasCoupons returns a boolean if a field has been set.
func (o *POSTCouponCodesPromotionRules201ResponseDataRelationships) HasCoupons() bool {
	if o != nil && !IsNil(o.Coupons) {
		return true
	}

	return false
}

// SetCoupons gets a reference to the given POSTBuyXPayYPromotions201ResponseDataRelationshipsCoupons and assigns it to the Coupons field.
func (o *POSTCouponCodesPromotionRules201ResponseDataRelationships) SetCoupons(v POSTBuyXPayYPromotions201ResponseDataRelationshipsCoupons) {
	o.Coupons = &v
}

func (o POSTCouponCodesPromotionRules201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTCouponCodesPromotionRules201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Promotion) {
		toSerialize["promotion"] = o.Promotion
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	if !IsNil(o.Coupons) {
		toSerialize["coupons"] = o.Coupons
	}
	return toSerialize, nil
}

type NullablePOSTCouponCodesPromotionRules201ResponseDataRelationships struct {
	value *POSTCouponCodesPromotionRules201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTCouponCodesPromotionRules201ResponseDataRelationships) Get() *POSTCouponCodesPromotionRules201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTCouponCodesPromotionRules201ResponseDataRelationships) Set(val *POSTCouponCodesPromotionRules201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCouponCodesPromotionRules201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCouponCodesPromotionRules201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCouponCodesPromotionRules201ResponseDataRelationships(val *POSTCouponCodesPromotionRules201ResponseDataRelationships) *NullablePOSTCouponCodesPromotionRules201ResponseDataRelationships {
	return &NullablePOSTCouponCodesPromotionRules201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTCouponCodesPromotionRules201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCouponCodesPromotionRules201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
