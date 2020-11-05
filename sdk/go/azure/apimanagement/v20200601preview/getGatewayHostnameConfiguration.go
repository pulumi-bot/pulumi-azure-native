// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200601preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupGatewayHostnameConfiguration(ctx *pulumi.Context, args *LookupGatewayHostnameConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupGatewayHostnameConfigurationResult, error) {
	var rv LookupGatewayHostnameConfigurationResult
	err := ctx.Invoke("azure-nextgen:apimanagement/v20200601preview:getGatewayHostnameConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGatewayHostnameConfigurationArgs struct {
	// Gateway entity identifier. Must be unique in the current API Management service instance. Must not have value 'managed'
	GatewayId string `pulumi:"gatewayId"`
	// Gateway hostname configuration identifier. Must be unique in the scope of parent Gateway entity.
	HcId string `pulumi:"hcId"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Gateway hostname configuration details.
type LookupGatewayHostnameConfigurationResult struct {
	// Identifier of Certificate entity that will be used for TLS connection establishment
	CertificateId *string `pulumi:"certificateId"`
	// Hostname value. Supports valid domain name, partial or full wildcard
	Hostname *string `pulumi:"hostname"`
	// Specifies if HTTP/2.0 is supported
	Http2Enabled *bool `pulumi:"http2Enabled"`
	// Resource name.
	Name string `pulumi:"name"`
	// Determines whether gateway requests client certificate
	NegotiateClientCertificate *bool `pulumi:"negotiateClientCertificate"`
	// Specifies if TLS 1.0 is supported
	Tls10Enabled *bool `pulumi:"tls10Enabled"`
	// Specifies if TLS 1.1 is supported
	Tls11Enabled *bool `pulumi:"tls11Enabled"`
	// Resource type for API Management resource.
	Type string `pulumi:"type"`
}
