// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupCertificate(ctx *pulumi.Context, args *LookupCertificateArgs, opts ...pulumi.InvokeOption) (*LookupCertificateResult, error) {
	var rv LookupCertificateResult
	err := ctx.Invoke("azurerm:apimanagement/v20180601preview:getCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCertificateArgs struct {
	// Identifier of the certificate entity. Must be unique in the current API Management service instance.
	CertificateId string `pulumi:"certificateId"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Certificate details.
type LookupCertificateResult struct {
	// Expiration date of the certificate. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	ExpirationDate string `pulumi:"expirationDate"`
	// Resource name.
	Name string `pulumi:"name"`
	// Subject attribute of the certificate.
	Subject string `pulumi:"subject"`
	// Thumbprint of the certificate.
	Thumbprint string `pulumi:"thumbprint"`
	// Resource type for API Management resource.
	Type string `pulumi:"type"`
}
