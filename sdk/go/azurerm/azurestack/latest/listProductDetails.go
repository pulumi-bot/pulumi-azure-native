// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func ListProductDetails(ctx *pulumi.Context, args *ListProductDetailsArgs, opts ...pulumi.InvokeOption) (*ListProductDetailsResult, error) {
	var rv ListProductDetailsResult
	err := ctx.Invoke("azurerm:azurestack/latest:listProductDetails", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListProductDetailsArgs struct {
	// Name of the product.
	ProductName string `pulumi:"productName"`
	// Name of the Azure Stack registration.
	RegistrationName string `pulumi:"registrationName"`
	// Name of the resource group.
	ResourceGroup string `pulumi:"resourceGroup"`
}

// Extended description about the product required for installing it into Azure Stack.
type ListProductDetailsResult struct {
	// Specifies kind of compute role included in the package.
	ComputeRole string `pulumi:"computeRole"`
	// List of attached data disks.
	DataDiskImages []DataDiskImageResponse `pulumi:"dataDiskImages"`
	// The URI to the .azpkg file that provides information required for showing product in the gallery.
	GalleryPackageBlobSasUri string `pulumi:"galleryPackageBlobSasUri"`
	// Specifies if product is a Virtual Machine Extension.
	IsSystemExtension bool `pulumi:"isSystemExtension"`
	// OS disk image used by product.
	OsDiskImage OsDiskImageResponse `pulumi:"osDiskImage"`
	// Specifies the kind of the product (virtualMachine or virtualMachineExtension).
	ProductKind string `pulumi:"productKind"`
	// Indicates if specified product supports multiple extensions.
	SupportMultipleExtensions bool `pulumi:"supportMultipleExtensions"`
	// The URI.
	Uri string `pulumi:"uri"`
	// Specifies product version.
	Version string `pulumi:"version"`
	// Specifies operating system used by the product.
	VmOsType string `pulumi:"vmOsType"`
	// Indicates if virtual machine Scale Set is enabled in the specified product.
	VmScaleSetEnabled bool `pulumi:"vmScaleSetEnabled"`
}