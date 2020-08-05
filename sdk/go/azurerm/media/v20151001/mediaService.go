// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20151001

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The properties of a Media Service resource.
type MediaService struct {
	pulumi.CustomResourceState

	// The geographic location of the resource. This must be one of the supported and registered Azure Geo Regions (for example, West US, East US, Southeast Asia, and so forth).
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The additional properties of a Media Service resource.
	Properties MediaServicePropertiesResponseOutput `pulumi:"properties"`
	// Tags to help categorize the resource in the Azure portal.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMediaService registers a new resource with the given unique name, arguments, and options.
func NewMediaService(ctx *pulumi.Context,
	name string, args *MediaServiceArgs, opts ...pulumi.ResourceOption) (*MediaService, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &MediaServiceArgs{}
	}
	var resource MediaService
	err := ctx.RegisterResource("azurerm:media/v20151001:MediaService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMediaService gets an existing MediaService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMediaService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MediaServiceState, opts ...pulumi.ResourceOption) (*MediaService, error) {
	var resource MediaService
	err := ctx.ReadResource("azurerm:media/v20151001:MediaService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MediaService resources.
type mediaServiceState struct {
	// The geographic location of the resource. This must be one of the supported and registered Azure Geo Regions (for example, West US, East US, Southeast Asia, and so forth).
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The additional properties of a Media Service resource.
	Properties *MediaServicePropertiesResponse `pulumi:"properties"`
	// Tags to help categorize the resource in the Azure portal.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource
	Type *string `pulumi:"type"`
}

type MediaServiceState struct {
	// The geographic location of the resource. This must be one of the supported and registered Azure Geo Regions (for example, West US, East US, Southeast Asia, and so forth).
	Location pulumi.StringPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The additional properties of a Media Service resource.
	Properties MediaServicePropertiesResponsePtrInput
	// Tags to help categorize the resource in the Azure portal.
	Tags pulumi.StringMapInput
	// The type of the resource
	Type pulumi.StringPtrInput
}

func (MediaServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*mediaServiceState)(nil)).Elem()
}

type mediaServiceArgs struct {
	// The geographic location of the resource. This must be one of the supported and registered Azure Geo Regions (for example, West US, East US, Southeast Asia, and so forth).
	Location *string `pulumi:"location"`
	// Name of the Media Service.
	Name string `pulumi:"name"`
	// Name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The storage accounts for this resource.
	StorageAccounts []StorageAccount `pulumi:"storageAccounts"`
	// Tags to help categorize the resource in the Azure portal.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a MediaService resource.
type MediaServiceArgs struct {
	// The geographic location of the resource. This must be one of the supported and registered Azure Geo Regions (for example, West US, East US, Southeast Asia, and so forth).
	Location pulumi.StringPtrInput
	// Name of the Media Service.
	Name pulumi.StringInput
	// Name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// The storage accounts for this resource.
	StorageAccounts StorageAccountArrayInput
	// Tags to help categorize the resource in the Azure portal.
	Tags pulumi.StringMapInput
}

func (MediaServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mediaServiceArgs)(nil)).Elem()
}
