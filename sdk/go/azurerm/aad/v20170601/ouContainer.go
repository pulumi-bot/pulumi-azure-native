// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170601

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Resource for OuContainer.
//
// ## Example Usage
// ### Create Domain Service
//
// ```go
// package main
//
// import (
// 	aad "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/aad/v20170601"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := aad.NewOuContainer(ctx, "ouContainer", &aad.OuContainerArgs{
// 			AccountName:       pulumi.String("AccountName1"),
// 			DomainServiceName: pulumi.String("OuContainer.com"),
// 			OuContainerName:   pulumi.String("OuContainer1"),
// 			Password:          pulumi.String("Password1"),
// 			ResourceGroupName: pulumi.String("OuContainerResourceGroup"),
// 			Spn:               pulumi.String("Spn1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type OuContainer struct {
	pulumi.CustomResourceState

	// The list of container accounts
	Accounts ContainerAccountResponseArrayOutput `pulumi:"accounts"`
	// The OuContainer name
	ContainerId pulumi.StringOutput `pulumi:"containerId"`
	// The Deployment id
	DeploymentId pulumi.StringOutput `pulumi:"deploymentId"`
	// The domain name of Domain Services.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// Resource etag
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Resource location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The current deployment or provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Status of OuContainer instance
	ServiceStatus pulumi.StringOutput `pulumi:"serviceStatus"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Azure Active Directory tenant id
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewOuContainer registers a new resource with the given unique name, arguments, and options.
func NewOuContainer(ctx *pulumi.Context,
	name string, args *OuContainerArgs, opts ...pulumi.ResourceOption) (*OuContainer, error) {
	if args == nil || args.DomainServiceName == nil {
		return nil, errors.New("missing required argument 'DomainServiceName'")
	}
	if args == nil || args.OuContainerName == nil {
		return nil, errors.New("missing required argument 'OuContainerName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &OuContainerArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:aad/latest:OuContainer"),
		},
		{
			Type: pulumi.String("azurerm:aad/v20200101:OuContainer"),
		},
	})
	opts = append(opts, aliases)
	var resource OuContainer
	err := ctx.RegisterResource("azurerm:aad/v20170601:OuContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOuContainer gets an existing OuContainer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOuContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OuContainerState, opts ...pulumi.ResourceOption) (*OuContainer, error) {
	var resource OuContainer
	err := ctx.ReadResource("azurerm:aad/v20170601:OuContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OuContainer resources.
type ouContainerState struct {
	// The list of container accounts
	Accounts []ContainerAccountResponse `pulumi:"accounts"`
	// The OuContainer name
	ContainerId *string `pulumi:"containerId"`
	// The Deployment id
	DeploymentId *string `pulumi:"deploymentId"`
	// The domain name of Domain Services.
	DomainName *string `pulumi:"domainName"`
	// Resource etag
	Etag *string `pulumi:"etag"`
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// The current deployment or provisioning state, which only appears in the response.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Status of OuContainer instance
	ServiceStatus *string `pulumi:"serviceStatus"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Azure Active Directory tenant id
	TenantId *string `pulumi:"tenantId"`
	// Resource type
	Type *string `pulumi:"type"`
}

type OuContainerState struct {
	// The list of container accounts
	Accounts ContainerAccountResponseArrayInput
	// The OuContainer name
	ContainerId pulumi.StringPtrInput
	// The Deployment id
	DeploymentId pulumi.StringPtrInput
	// The domain name of Domain Services.
	DomainName pulumi.StringPtrInput
	// Resource etag
	Etag pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// The current deployment or provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringPtrInput
	// Status of OuContainer instance
	ServiceStatus pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Azure Active Directory tenant id
	TenantId pulumi.StringPtrInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (OuContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*ouContainerState)(nil)).Elem()
}

type ouContainerArgs struct {
	// The account name
	AccountName *string `pulumi:"accountName"`
	// The name of the domain service.
	DomainServiceName string `pulumi:"domainServiceName"`
	// The name of the OuContainer.
	OuContainerName string `pulumi:"ouContainerName"`
	// The account password
	Password *string `pulumi:"password"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The account spn
	Spn *string `pulumi:"spn"`
}

// The set of arguments for constructing a OuContainer resource.
type OuContainerArgs struct {
	// The account name
	AccountName pulumi.StringPtrInput
	// The name of the domain service.
	DomainServiceName pulumi.StringInput
	// The name of the OuContainer.
	OuContainerName pulumi.StringInput
	// The account password
	Password pulumi.StringPtrInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The account spn
	Spn pulumi.StringPtrInput
}

func (OuContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ouContainerArgs)(nil)).Elem()
}
