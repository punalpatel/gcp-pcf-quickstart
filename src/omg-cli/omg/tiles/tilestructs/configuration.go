/*
 * Copyright 2017 Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package tilestructs

import (
	"omg-cli/config"
)

type AvalibilityZone struct {
	Name string `json:"name"`
}

type NetworkName struct {
	Name string `json:"name"`
}

type Network struct {
	SingletonAvalibilityZone AvalibilityZone   `json:"singleton_availability_zone"`
	OtherAvailabilityZones   []AvalibilityZone `json:"other_availability_zones"`
	Network                  NetworkName       `json:"network"`
	ODBNetwork               NetworkName       `json:"service_network"`
}

func NetworkConfig(subnetName string, cfg *config.Config) Network {
	return Network{
		SingletonAvalibilityZone: AvalibilityZone{cfg.Zone1},
		OtherAvailabilityZones:   []AvalibilityZone{{cfg.Zone1}, {cfg.Zone2}, {cfg.Zone3}},
		Network:                  NetworkName{subnetName},
	}
}

func NetworkODBConfig(subnetName string, cfg *config.Config, odbNetworkName string) Network {
	return Network{
		SingletonAvalibilityZone: AvalibilityZone{cfg.Zone1},
		OtherAvailabilityZones:   []AvalibilityZone{{cfg.Zone1}, {cfg.Zone2}, {cfg.Zone3}},
		Network:                  NetworkName{subnetName},
		ODBNetwork:               NetworkName{odbNetworkName},
	}
}

type Value struct {
	Value string `json:"value"`
}

type IntegerValue struct {
	Value int `json:"value"`
}

type BooleanValue struct {
	Value bool `json:"value"`
}

type Secret struct {
	Value string `json:"secret"`
}

type SecretValue struct {
	Sec Secret `json:"value"`
}

type Certificate struct {
	PublicKey  string `json:"cert_pem"`
	PrivateKey string `json:"private_key_pem"`
}

type CertificateConstruct struct {
	Certificate Certificate `json:"certificate"`
	Name        string      `json:"name"`
}

type CertificateValue struct {
	Value []CertificateConstruct `json:"value"`
}

type OldCertificateValue struct {
	Value Certificate `json:"value"`
}

type KeyStruct struct {
	Secret string `json:"secret"`
}
type EncryptionKey struct {
	Name    string    `json:"name"`
	Key     KeyStruct `json:"key"`
	Primary bool      `json:"primary"`
}
type EncryptionKeyValue struct {
	Value []EncryptionKey `json:"value"`
}

type Resource struct {
	RouterNames       []string `json:"elb_names,omitempty"`
	Instances         *int     `json:"instances,omitempty"`
	InternetConnected bool     `json:"internet_connected"`
	VmTypeId          string   `json:"vm_type_id,omitempty"`
}
