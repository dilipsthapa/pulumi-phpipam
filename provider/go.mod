module github.com/dilipsthapa/pulumi-phpipam/provider

go 1.16

replace github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0

require (
	github.com/hashicorp/terraform-plugin-sdk v1.17.2 // indirect
	github.com/lord-kyron/terraform-provider-phpipam v1.2.8
	github.com/pulumi/pulumi-terraform v1.1.0
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.2.1
	github.com/pulumi/pulumi/sdk/v3 v3.4.0
	github.com/dilipsthapa/pulumi-phpipam v0.0.1
)
