// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190701

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Specifies information about the gallery Image Definition that you want to create or update.
//
// ## Example Usage
// ### Create or update a simple gallery image.
//
// ```go
// package main
//
// import (
// 	compute "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/compute/v20190701"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewGalleryImage(ctx, "galleryImage", &compute.GalleryImageArgs{
// 			GalleryImageName: pulumi.String("myGalleryImageName"),
// 			GalleryName:      pulumi.String("myGalleryName"),
// 			HyperVGeneration: pulumi.String("V1"),
// 			Identifier: &compute.GalleryImageIdentifierArgs{
// 				Offer:     pulumi.String("myOfferName"),
// 				Publisher: pulumi.String("myPublisherName"),
// 				Sku:       pulumi.String("mySkuName"),
// 			},
// 			Location:          pulumi.String("West US"),
// 			OsState:           pulumi.String("Generalized"),
// 			OsType:            pulumi.String("Windows"),
// 			ResourceGroupName: pulumi.String("myResourceGroup"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type GalleryImage struct {
	pulumi.CustomResourceState

	// The description of this gallery Image Definition resource. This property is updatable.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Describes the disallowed disk types.
	Disallowed DisallowedResponsePtrOutput `pulumi:"disallowed"`
	// The end of life date of the gallery Image Definition. This property can be used for decommissioning purposes. This property is updatable.
	EndOfLifeDate pulumi.StringPtrOutput `pulumi:"endOfLifeDate"`
	// The Eula agreement for the gallery Image Definition.
	Eula pulumi.StringPtrOutput `pulumi:"eula"`
	// The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
	HyperVGeneration pulumi.StringPtrOutput `pulumi:"hyperVGeneration"`
	// This is the gallery Image Definition identifier.
	Identifier GalleryImageIdentifierResponseOutput `pulumi:"identifier"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// This property allows the user to specify whether the virtual machines created under this image are 'Generalized' or 'Specialized'.
	OsState pulumi.StringOutput `pulumi:"osState"`
	// This property allows you to specify the type of the OS that is included in the disk when creating a VM from a managed image. <br><br> Possible values are: <br><br> **Windows** <br><br> **Linux**
	OsType pulumi.StringOutput `pulumi:"osType"`
	// The privacy statement uri.
	PrivacyStatementUri pulumi.StringPtrOutput `pulumi:"privacyStatementUri"`
	// The provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Describes the gallery Image Definition purchase plan. This is used by marketplace images.
	PurchasePlan ImagePurchasePlanResponsePtrOutput `pulumi:"purchasePlan"`
	// The properties describe the recommended machine configuration for this Image Definition. These properties are updatable.
	Recommended RecommendedMachineConfigurationResponsePtrOutput `pulumi:"recommended"`
	// The release note uri.
	ReleaseNoteUri pulumi.StringPtrOutput `pulumi:"releaseNoteUri"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGalleryImage registers a new resource with the given unique name, arguments, and options.
func NewGalleryImage(ctx *pulumi.Context,
	name string, args *GalleryImageArgs, opts ...pulumi.ResourceOption) (*GalleryImage, error) {
	if args == nil || args.GalleryImageName == nil {
		return nil, errors.New("missing required argument 'GalleryImageName'")
	}
	if args == nil || args.GalleryName == nil {
		return nil, errors.New("missing required argument 'GalleryName'")
	}
	if args == nil || args.Identifier == nil {
		return nil, errors.New("missing required argument 'Identifier'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.OsState == nil {
		return nil, errors.New("missing required argument 'OsState'")
	}
	if args == nil || args.OsType == nil {
		return nil, errors.New("missing required argument 'OsType'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &GalleryImageArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:compute/latest:GalleryImage"),
		},
		{
			Type: pulumi.String("azurerm:compute/v20180601:GalleryImage"),
		},
		{
			Type: pulumi.String("azurerm:compute/v20190301:GalleryImage"),
		},
		{
			Type: pulumi.String("azurerm:compute/v20191201:GalleryImage"),
		},
		{
			Type: pulumi.String("azurerm:compute/v20200930:GalleryImage"),
		},
	})
	opts = append(opts, aliases)
	var resource GalleryImage
	err := ctx.RegisterResource("azurerm:compute/v20190701:GalleryImage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGalleryImage gets an existing GalleryImage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGalleryImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GalleryImageState, opts ...pulumi.ResourceOption) (*GalleryImage, error) {
	var resource GalleryImage
	err := ctx.ReadResource("azurerm:compute/v20190701:GalleryImage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GalleryImage resources.
type galleryImageState struct {
	// The description of this gallery Image Definition resource. This property is updatable.
	Description *string `pulumi:"description"`
	// Describes the disallowed disk types.
	Disallowed *DisallowedResponse `pulumi:"disallowed"`
	// The end of life date of the gallery Image Definition. This property can be used for decommissioning purposes. This property is updatable.
	EndOfLifeDate *string `pulumi:"endOfLifeDate"`
	// The Eula agreement for the gallery Image Definition.
	Eula *string `pulumi:"eula"`
	// The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
	HyperVGeneration *string `pulumi:"hyperVGeneration"`
	// This is the gallery Image Definition identifier.
	Identifier *GalleryImageIdentifierResponse `pulumi:"identifier"`
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// This property allows the user to specify whether the virtual machines created under this image are 'Generalized' or 'Specialized'.
	OsState *string `pulumi:"osState"`
	// This property allows you to specify the type of the OS that is included in the disk when creating a VM from a managed image. <br><br> Possible values are: <br><br> **Windows** <br><br> **Linux**
	OsType *string `pulumi:"osType"`
	// The privacy statement uri.
	PrivacyStatementUri *string `pulumi:"privacyStatementUri"`
	// The provisioning state, which only appears in the response.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Describes the gallery Image Definition purchase plan. This is used by marketplace images.
	PurchasePlan *ImagePurchasePlanResponse `pulumi:"purchasePlan"`
	// The properties describe the recommended machine configuration for this Image Definition. These properties are updatable.
	Recommended *RecommendedMachineConfigurationResponse `pulumi:"recommended"`
	// The release note uri.
	ReleaseNoteUri *string `pulumi:"releaseNoteUri"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type GalleryImageState struct {
	// The description of this gallery Image Definition resource. This property is updatable.
	Description pulumi.StringPtrInput
	// Describes the disallowed disk types.
	Disallowed DisallowedResponsePtrInput
	// The end of life date of the gallery Image Definition. This property can be used for decommissioning purposes. This property is updatable.
	EndOfLifeDate pulumi.StringPtrInput
	// The Eula agreement for the gallery Image Definition.
	Eula pulumi.StringPtrInput
	// The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
	HyperVGeneration pulumi.StringPtrInput
	// This is the gallery Image Definition identifier.
	Identifier GalleryImageIdentifierResponsePtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// This property allows the user to specify whether the virtual machines created under this image are 'Generalized' or 'Specialized'.
	OsState pulumi.StringPtrInput
	// This property allows you to specify the type of the OS that is included in the disk when creating a VM from a managed image. <br><br> Possible values are: <br><br> **Windows** <br><br> **Linux**
	OsType pulumi.StringPtrInput
	// The privacy statement uri.
	PrivacyStatementUri pulumi.StringPtrInput
	// The provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringPtrInput
	// Describes the gallery Image Definition purchase plan. This is used by marketplace images.
	PurchasePlan ImagePurchasePlanResponsePtrInput
	// The properties describe the recommended machine configuration for this Image Definition. These properties are updatable.
	Recommended RecommendedMachineConfigurationResponsePtrInput
	// The release note uri.
	ReleaseNoteUri pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (GalleryImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*galleryImageState)(nil)).Elem()
}

type galleryImageArgs struct {
	// The description of this gallery Image Definition resource. This property is updatable.
	Description *string `pulumi:"description"`
	// Describes the disallowed disk types.
	Disallowed *Disallowed `pulumi:"disallowed"`
	// The end of life date of the gallery Image Definition. This property can be used for decommissioning purposes. This property is updatable.
	EndOfLifeDate *string `pulumi:"endOfLifeDate"`
	// The Eula agreement for the gallery Image Definition.
	Eula *string `pulumi:"eula"`
	// The name of the gallery Image Definition to be created or updated. The allowed characters are alphabets and numbers with dots, dashes, and periods allowed in the middle. The maximum length is 80 characters.
	GalleryImageName string `pulumi:"galleryImageName"`
	// The name of the Shared Image Gallery in which the Image Definition is to be created.
	GalleryName string `pulumi:"galleryName"`
	// The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
	HyperVGeneration *string `pulumi:"hyperVGeneration"`
	// This is the gallery Image Definition identifier.
	Identifier GalleryImageIdentifier `pulumi:"identifier"`
	// Resource location
	Location string `pulumi:"location"`
	// This property allows the user to specify whether the virtual machines created under this image are 'Generalized' or 'Specialized'.
	OsState string `pulumi:"osState"`
	// This property allows you to specify the type of the OS that is included in the disk when creating a VM from a managed image. <br><br> Possible values are: <br><br> **Windows** <br><br> **Linux**
	OsType string `pulumi:"osType"`
	// The privacy statement uri.
	PrivacyStatementUri *string `pulumi:"privacyStatementUri"`
	// Describes the gallery Image Definition purchase plan. This is used by marketplace images.
	PurchasePlan *ImagePurchasePlan `pulumi:"purchasePlan"`
	// The properties describe the recommended machine configuration for this Image Definition. These properties are updatable.
	Recommended *RecommendedMachineConfiguration `pulumi:"recommended"`
	// The release note uri.
	ReleaseNoteUri *string `pulumi:"releaseNoteUri"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a GalleryImage resource.
type GalleryImageArgs struct {
	// The description of this gallery Image Definition resource. This property is updatable.
	Description pulumi.StringPtrInput
	// Describes the disallowed disk types.
	Disallowed DisallowedPtrInput
	// The end of life date of the gallery Image Definition. This property can be used for decommissioning purposes. This property is updatable.
	EndOfLifeDate pulumi.StringPtrInput
	// The Eula agreement for the gallery Image Definition.
	Eula pulumi.StringPtrInput
	// The name of the gallery Image Definition to be created or updated. The allowed characters are alphabets and numbers with dots, dashes, and periods allowed in the middle. The maximum length is 80 characters.
	GalleryImageName pulumi.StringInput
	// The name of the Shared Image Gallery in which the Image Definition is to be created.
	GalleryName pulumi.StringInput
	// The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
	HyperVGeneration pulumi.StringPtrInput
	// This is the gallery Image Definition identifier.
	Identifier GalleryImageIdentifierInput
	// Resource location
	Location pulumi.StringInput
	// This property allows the user to specify whether the virtual machines created under this image are 'Generalized' or 'Specialized'.
	OsState pulumi.StringInput
	// This property allows you to specify the type of the OS that is included in the disk when creating a VM from a managed image. <br><br> Possible values are: <br><br> **Windows** <br><br> **Linux**
	OsType pulumi.StringInput
	// The privacy statement uri.
	PrivacyStatementUri pulumi.StringPtrInput
	// Describes the gallery Image Definition purchase plan. This is used by marketplace images.
	PurchasePlan ImagePurchasePlanPtrInput
	// The properties describe the recommended machine configuration for this Image Definition. These properties are updatable.
	Recommended RecommendedMachineConfigurationPtrInput
	// The release note uri.
	ReleaseNoteUri pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (GalleryImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*galleryImageArgs)(nil)).Elem()
}
