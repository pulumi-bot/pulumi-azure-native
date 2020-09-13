// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200101preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The connector setting
//
// ## Example Usage
// ### AwsAssumeRole - Create a cloud account connector for a subscription
//
// ```go
// package main
//
// import (
// 	security "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/security/v20200101preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := security.NewConnector(ctx, "connector", &security.ConnectorArgs{
// 			AuthenticationDetails: &security.AuthenticationDetailsPropertiesArgs{
// 				AuthenticationType: pulumi.String("awsAssumeRole"),
// 				AwsAssumeRoleArn:   pulumi.String("arn:aws:iam::81231569658:role/AscConnector"),
// 				AwsExternalId:      pulumi.String("20ff7fc3-e762-44dd-bd96-b71116dcdc23"),
// 			},
// 			ConnectorName: pulumi.String("aws_dev2"),
// 			HybridComputeSettings: &security.HybridComputeSettingsPropertiesArgs{
// 				AutoProvision: pulumi.String("On"),
// 				ProxyServer: &security.ProxyServerPropertiesArgs{
// 					Ip:   pulumi.String("167.220.197.140"),
// 					Port: pulumi.String("34"),
// 				},
// 				Region:            pulumi.String("West US 2"),
// 				ResourceGroupName: pulumi.String("AwsConnectorRG"),
// 				ServicePrincipal: &security.ServicePrincipalPropertiesArgs{
// 					ApplicationId: pulumi.String("ad9bcd79-be9c-45ab-abd8-80ca1654a7d1"),
// 					Secret:        pulumi.String("x2yS:FnCHssRkH0@CJY5pATzlEs@r5m."),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### AwsCred -  Create a cloud account connector for a subscription
//
// ```go
// package main
//
// import (
// 	security "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/security/v20200101preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := security.NewConnector(ctx, "connector", &security.ConnectorArgs{
// 			AuthenticationDetails: &security.AuthenticationDetailsPropertiesArgs{
// 				AuthenticationType: pulumi.String("awsCreds"),
// 				AwsAccessKeyId:     pulumi.String("AKIARPZCNODDNAEQFSOE"),
// 				AwsSecretAccessKey: pulumi.String("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"),
// 			},
// 			ConnectorName: pulumi.String("aws_dev1"),
// 			HybridComputeSettings: &security.HybridComputeSettingsPropertiesArgs{
// 				AutoProvision: pulumi.String("On"),
// 				ProxyServer: &security.ProxyServerPropertiesArgs{
// 					Ip:   pulumi.String("167.220.197.140"),
// 					Port: pulumi.String("34"),
// 				},
// 				Region:            pulumi.String("West US 2"),
// 				ResourceGroupName: pulumi.String("AwsConnectorRG"),
// 				ServicePrincipal: &security.ServicePrincipalPropertiesArgs{
// 					ApplicationId: pulumi.String("ad9bcd79-be9c-45ab-abd8-80ca1654a7d1"),
// 					Secret:        pulumi.String("x2yS:FnCHssRkH0@CJY5pATzlEs@r5m."),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### gcpCredentials - Create a cloud account connector for a subscription
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	security "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/security/v20200101preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := security.NewConnector(ctx, "connector", &security.ConnectorArgs{
// 			AuthenticationDetails: &security.AuthenticationDetailsPropertiesArgs{
// 				AuthProviderX509CertUrl: pulumi.String("https://www.googleapis.com/oauth2/v1/certs"),
// 				AuthUri:                 pulumi.String("https://accounts.google.com/o/oauth2/auth"),
// 				AuthenticationType:      pulumi.String("gcpCredentials"),
// 				ClientEmail:             pulumi.String("asc-135@asc-project-1234.iam.gserviceaccount.com"),
// 				ClientId:                pulumi.String("105889053725632919854"),
// 				ClientX509CertUrl:       pulumi.String(fmt.Sprintf("%v%v%v", "https://www.googleapis.com/robot/v1/metadata/x509/asc-135", "%", "40asc-project-1234.iam.gserviceaccount.com")),
// 				OrganizationId:          pulumi.String("AscDemoOrg"),
// 				PrivateKey:              pulumi.String("-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCpxYHcLzcDZ6/Q\nAeQZnQXM5GTb3p09Xsbjo2T2F61b6I7FZiQXBrbw3Zf0CUCkkqTTpD5xifl82yQ6\n89V7SAe8hxI7esAcVDhm/aJMqzVjHLISAU2L3li1sn0jjY2oYtndwN6bRivP8O6t\n9F+W6E0zMlbCxtpZEHLbb6WxlJJrwEQ0MPH2yOCwZUQi6NHksAtEzX2nNKJNyUC7\nQyBVHHMm34H2bmZwsuQp3y2otpcJ9tJnVmYfC3k/w4x2L+DIK7JnQP/C1wQqu2du\nc0w6sydF6RhLoHButrVdYRJTdfK4k03SsSTyMqZ+f7LNnKw3xenzw1VmEpk8mvoQ\nt08tCBOrAgMBAAECggEAByzz6iyMtLYjNjV+QJ7kad6VbL2iA8AHxANZ9xTVHPdd\nYXaJu/dqsA+NpqDlfI8+LDva782XH/HbPCqmMUnAGfXTjXQIvqnIoIHD5F2wKfpC\nhIRNlMXXFgbvRxtqi11yO+80+XcjzuwuCmgzyhsTeEB+bkkdXXpWgHPdmv3emnM6\nMQM9Zgrug0UndPmiUwKOcJSU4PlmlTpHEV4vA6JfA4bvphy9m1jxO5qWeah5yym2\n6FP5BRIDF98kFrDnSXJjajwgLCQ+MypFQXyax6XkxDxuKXbng1bv7eZDjqazIChk\nm0y14X0s0jnWc+AX8vfeSf7d+EsGdVinEwR1aAawEQKBgQDqDB0qxcIQ1oI1Kww8\n9vXefTiuWsf47F+fJ/DIOEbiRfE8IdCgmOABvcqJIoxW/DFMBEdLCcx73Km7pOmd\nKg1ddScnaO8cOj2v/Ub+fAqVrA4ki4ViYP0A7/Nogga3Jr/x3ey5bitrIfFImteS\nCgBHBzZvoQpvO4lB2tKVgo2P9wKBgQC5sgTEq4sasRGSAY6lIoJno0I8w28a/16D\nes60XQeY1ger8uTGwlT02v/u/arDUmRLPClpujXq6gK29KvtRCHy7JkpGbqW2bZs\nPFKKWR7Tk3XPKYyjv94AIi5/xoFeDhS4lpAvy3Z5tQhYS6wqWKvT6yZQ3kM+Hfxs\npHgvu3mU7QKBgQC9/E1k3hj1cBtMK4CIsHPPQljTd4+iacYJPPPAo6YuoVX8WPqw\nksgrwbN59Fh1d8xQh5yTtgWOegYx8uFMGcm1lpbM7+pBQKm4hWGuzGQPMRZd5f/F\nZzOZIi61I+9tlv/yxxIVR+/ozCm/pSneO04UWi9/F/uPZYW6tnWAtfRR6wKBgGsZ\n8MQaCK4JaI/klAhMghgSQnbXZXKVzUZaA3Rln6cX8u7KtgapOOTMlwaZie8Dy1LV\nTTFstAJcm9o3/h1nyYjZy3C4JTUyNpPwqs6enjf7edxVI4eidwFutZD+xcigqHTa\naikW2atSrZB3fMIjyF7+5meH+hKOqvNiXOty3qn1AoGAZuVxYQy5FVq3YZxzr3Aa\nAm0ShoXTF6QYIbsaUiUGoa/NlHcw9V/lj4AqBRbxbaYMD+hz2J/od9cb268eJKY8\n3b6MvaUqdNhNnWodJXLhgtmGEHDKmTppz2JSTx/tVzCfhFdcOC79StZvcKLhtoFQ\n+/3lEw6NCIXzm5E4+dtJG4k=\n-----END PRIVATE KEY-----\n"),
// 				PrivateKeyId:            pulumi.String("6efg587hra2568as34d22326b044cc20dc2af"),
// 				ProjectId:               pulumi.String("asc-project-1234"),
// 				TokenUri:                pulumi.String("https://oauth2.googleapis.com/token"),
// 				Type:                    pulumi.String("service_account"),
// 			},
// 			ConnectorName: pulumi.String("gcp_dev"),
// 			HybridComputeSettings: &security.HybridComputeSettingsPropertiesArgs{
// 				AutoProvision: pulumi.String("Off"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type Connector struct {
	pulumi.CustomResourceState

	// Settings for authentication management, these settings are relevant only for the cloud connector.
	AuthenticationDetails pulumi.AnyOutput `pulumi:"authenticationDetails"`
	// Settings for hybrid compute management, these settings are relevant only Arc autoProvision (Hybrid Compute).
	HybridComputeSettings HybridComputeSettingsPropertiesResponsePtrOutput `pulumi:"hybridComputeSettings"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewConnector registers a new resource with the given unique name, arguments, and options.
func NewConnector(ctx *pulumi.Context,
	name string, args *ConnectorArgs, opts ...pulumi.ResourceOption) (*Connector, error) {
	if args == nil || args.ConnectorName == nil {
		return nil, errors.New("missing required argument 'ConnectorName'")
	}
	if args == nil {
		args = &ConnectorArgs{}
	}
	var resource Connector
	err := ctx.RegisterResource("azurerm:security/v20200101preview:Connector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnector gets an existing Connector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectorState, opts ...pulumi.ResourceOption) (*Connector, error) {
	var resource Connector
	err := ctx.ReadResource("azurerm:security/v20200101preview:Connector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Connector resources.
type connectorState struct {
	// Settings for authentication management, these settings are relevant only for the cloud connector.
	AuthenticationDetails interface{} `pulumi:"authenticationDetails"`
	// Settings for hybrid compute management, these settings are relevant only Arc autoProvision (Hybrid Compute).
	HybridComputeSettings *HybridComputeSettingsPropertiesResponse `pulumi:"hybridComputeSettings"`
	// Resource name
	Name *string `pulumi:"name"`
	// Resource type
	Type *string `pulumi:"type"`
}

type ConnectorState struct {
	// Settings for authentication management, these settings are relevant only for the cloud connector.
	AuthenticationDetails pulumi.Input
	// Settings for hybrid compute management, these settings are relevant only Arc autoProvision (Hybrid Compute).
	HybridComputeSettings HybridComputeSettingsPropertiesResponsePtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (ConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectorState)(nil)).Elem()
}

type connectorArgs struct {
	// Settings for authentication management, these settings are relevant only for the cloud connector.
	AuthenticationDetails interface{} `pulumi:"authenticationDetails"`
	// Name of the cloud account connector
	ConnectorName string `pulumi:"connectorName"`
	// Settings for hybrid compute management, these settings are relevant only Arc autoProvision (Hybrid Compute).
	HybridComputeSettings *HybridComputeSettingsProperties `pulumi:"hybridComputeSettings"`
}

// The set of arguments for constructing a Connector resource.
type ConnectorArgs struct {
	// Settings for authentication management, these settings are relevant only for the cloud connector.
	AuthenticationDetails pulumi.Input
	// Name of the cloud account connector
	ConnectorName pulumi.StringInput
	// Settings for hybrid compute management, these settings are relevant only Arc autoProvision (Hybrid Compute).
	HybridComputeSettings HybridComputeSettingsPropertiesPtrInput
}

func (ConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectorArgs)(nil)).Elem()
}
