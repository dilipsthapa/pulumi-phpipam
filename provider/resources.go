// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package phpipam

import (
	"unicode"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

// all of the Enterprise Cloud token components used below.
const (
	// packages:
	ipamPkg = "ipam"
)

func ipamMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(ipamPkg + ":" + mod + ":" + mem)
}

func ipamType(mod string, typ string) tokens.Type {
	return tokens.Type(ipamMember(mod, typ))
}

func ipamDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return ipamMember(mod+"/"+fn, res)
}

func ipamResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return ipamType(mod+"/"+fn, res)
}

// Provider returns additional overlaid schema and metadata associated with the ipam package.
func Provider() tfbridge.ProviderInfo {
	p := ipam.Provider().(*schema.Provider)

	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "terraform-provider-phpipam",
		GitHubOrg:   "lord-kyron",
		Description: "A Pulumi package for creating phpipam resources.",
		Keywords:    []string{"pulumi", "ipam"},
		Homepage:    "https://pulumi.io",
		License:     "Apache-2.0",
		Repository:  "https://github.com/lord-kyron/terraform-provider-phpipam",
		Config: map[string]*tfbridge.SchemaInfo{
			"AppID": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"None"},
				},
			},
			"Endpoint": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"http://localhost/api"},
				},
			},
			"Username": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OS_USERNAME"},
				},
			},
			"Password": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OS_PASSWORD"},
				},
			},
			"Insecure": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []bool{false},
				},
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"phpipam_address":  {Tok: ipamDataSource(ipamPkg, "getAddress")},
			"phpipam_addresses": {Tok: ipamDataSource(ipamPkg, "getAddresses")},
			"phpipam_first_free_address": {Tok: ipamDataSource(ipamPkg, "getFirst_free_address")},
			"phpipam_section": {Tok: ipamDataSource(ipamPkg, "getSection")},
			"phpipam_subnet": {Tok: ipamDataSource(ipamPkg, "getSubnet")},
			"phpipam_subnets":       {Tok: ipamDataSource(ipamPkg, "getSubnets")},
			"phpipam_vlan":        {Tok: ipamDataSource(ipamPkg, "getVlan")},
			"phpipam_first_free_subnet":        {Tok: ipamDataSource(ipamPkg, "getFirst_free_subnet")},
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			"phpipam_address":      {Tok: ipamResource(ipamPkg, "address")},
			"phpipam_section":       {Tok: ipamResource(ipamPkg, "section")},
			"phpipam_subnet": {Tok: ipamResource(ipamPkg, "subnet")},
			"phpipam_first_free_address":        {Tok: ipamResource(ipamPkg, "first_free_address")},
			"phpipam_first_free_subnet": {Tok: ipamResource(ipamPkg, "first_free_subnet")},
		},

		JavaScript: &tfbridge.JavaScriptInfo{
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25", // so we can access strongly typed node definitions.
			},
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^0.17.0",
			},
			Overlay: &tfbridge.OverlayInfo{
				Files:   []string{},
				Modules: map[string]*tfbridge.OverlayInfo{},
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=0.17.0,<0.18.0",
			},
		},
	}

	// For all resources with name properties, we will add an auto-name property.  Make sure to skip those that
	// already have a name mapping entry, since those may have custom overrides set above (e.g., for length).
	const ipamName = "name"
	for resname, res := range prov.Resources {
		if schema := p.ResourcesMap[resname]; schema != nil {
			// Only apply auto-name to input properties (Optional || Required) named `name`
			if tfs, has := schema.Schema[ipamName]; has && (tfs.Optional || tfs.Required) {
				if _, hasfield := res.Fields[ipamName]; !hasfield {
					if res.Fields == nil {
						res.Fields = make(map[string]*tfbridge.SchemaInfo)
					}
					// Use conservative options that apply broadly for Enterprise Cloud.
					res.Fields[ipamName] = tfbridge.AutoName(ipamName, 255)
				}
			}
		}
	}

	return prov
}
