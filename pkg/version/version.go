// Licensed to Alexandre VILAIN under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Alexandre VILAIN licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

//+kubebuilder:object:generate=true
package version

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Masterminds/semver/v3"
)

var (
	// SupportedVersionsRange holds all supported temporal versions.
	SupportedVersionsRange = mustNewConstraint(">= 1.14.0 < 1.19.0")
	V1_18_0                = MustNewVersionFromString("1.18.0")
)

// +kubebuilder:validation:Type=string
// Version is a wrapper around semver.Version which supports correct
// marshaling to YAML and JSON. In particular, it marshals into strings.
type Version struct {
	*semver.Version
}

// Validate checks if the current version is in the supported temporal cluster
// version range.
func (v *Version) Validate() error {
	inRange := SupportedVersionsRange.Check(v.Version)
	if !inRange {
		return errors.New("provided version not in the supported range")
	}
	return nil
}

// ToUnstructured implements the value.UnstructuredConverter interface.
func (v Version) ToUnstructured() interface{} {
	return v.Version.String()
}

// UnmarshalJSON implements the json.Unmarshaller interface.
func (v *Version) UnmarshalJSON(b []byte) error {
	var str string
	err := json.Unmarshal(b, &str)
	if err != nil {
		return err
	}

	parsed, err := semver.NewVersion(str)
	if err != nil {
		return err
	}
	v.Version = parsed
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (v Version) MarshalJSON() ([]byte, error) {
	if v.Version == nil {
		return []byte("0.0.0"), nil
	}
	return json.Marshal(v.Version.String())
}

func (v *Version) GreaterOrEqual(compare *Version) bool {
	str := fmt.Sprintf(">= %s", compare.String())
	c, _ := semver.NewConstraint(str)
	return c.Check(v.Version)
}

// OpenAPISchemaType is used by the kube-openapi generator when constructing
// the OpenAPI spec of this type.
//
// See: https://github.com/kubernetes/kube-openapi/tree/master/pkg/generators
func (Version) OpenAPISchemaType() []string { return []string{"string"} }

// OpenAPISchemaFormat is used by the kube-openapi generator when constructing
// the OpenAPI spec of this type.
func (Version) OpenAPISchemaFormat() string { return "" }

func NewVersionFromString(v string) (*Version, error) {
	version, err := semver.NewVersion(v)
	return &Version{Version: version}, err
}

func MustNewVersionFromString(v string) *Version {
	version, err := NewVersionFromString(v)
	if err != nil {
		panic(err)
	}
	return version
}

func mustNewConstraint(constraint string) *semver.Constraints {
	c, err := semver.NewConstraint(constraint)
	if err != nil {
		panic(err)
	}
	return c
}
