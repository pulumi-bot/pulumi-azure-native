// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupAppServicePublicCertificate(ctx *pulumi.Context, args *LookupAppServicePublicCertificateArgs, opts ...pulumi.InvokeOption) (*LookupAppServicePublicCertificateResult, error) {
	var rv LookupAppServicePublicCertificateResult
	err := ctx.Invoke("azurerm:web:getAppServicePublicCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAppServicePublicCertificateArgs struct {
	// Public certificate name.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Public certificate object
type LookupAppServicePublicCertificateResult struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// PublicCertificate resource specific properties
	Properties PublicCertificateResponseProperties `pulumi:"properties"`
	// Resource type.
	Type string `pulumi:"type"`
}
