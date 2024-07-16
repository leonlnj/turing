/*
 * Turing Minimal Openapi Spec for SDK
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// BigQueryDataset struct for BigQueryDataset
type BigQueryDataset struct {
	Type string `json:"type"`
	BqConfig BigQueryDatasetConfig `json:"bq_config"`
}

// NewBigQueryDataset instantiates a new BigQueryDataset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBigQueryDataset(type_ string, bqConfig BigQueryDatasetConfig) *BigQueryDataset {
	this := BigQueryDataset{}
	this.Type = type_
	this.BqConfig = bqConfig
	return &this
}

// NewBigQueryDatasetWithDefaults instantiates a new BigQueryDataset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBigQueryDatasetWithDefaults() *BigQueryDataset {
	this := BigQueryDataset{}
	var type_ string = "BQ"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *BigQueryDataset) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BigQueryDataset) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BigQueryDataset) SetType(v string) {
	o.Type = v
}

// GetBqConfig returns the BqConfig field value
func (o *BigQueryDataset) GetBqConfig() BigQueryDatasetConfig {
	if o == nil {
		var ret BigQueryDatasetConfig
		return ret
	}

	return o.BqConfig
}

// GetBqConfigOk returns a tuple with the BqConfig field value
// and a boolean to check if the value has been set.
func (o *BigQueryDataset) GetBqConfigOk() (*BigQueryDatasetConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BqConfig, true
}

// SetBqConfig sets field value
func (o *BigQueryDataset) SetBqConfig(v BigQueryDatasetConfig) {
	o.BqConfig = v
}

func (o BigQueryDataset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["bq_config"] = o.BqConfig
	}
	return json.Marshal(toSerialize)
}

type NullableBigQueryDataset struct {
	value *BigQueryDataset
	isSet bool
}

func (v NullableBigQueryDataset) Get() *BigQueryDataset {
	return v.value
}

func (v *NullableBigQueryDataset) Set(val *BigQueryDataset) {
	v.value = val
	v.isSet = true
}

func (v NullableBigQueryDataset) IsSet() bool {
	return v.isSet
}

func (v *NullableBigQueryDataset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBigQueryDataset(val *BigQueryDataset) *NullableBigQueryDataset {
	return &NullableBigQueryDataset{value: val, isSet: true}
}

func (v NullableBigQueryDataset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBigQueryDataset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


