// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191101preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The resource that defines the source location where the artifacts are located.
//
// ## Example Usage
// ### Create artifact source
//
// ```go
// package main
//
// import (
// 	deploymentmanager "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/deploymentmanager/v20191101preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := deploymentmanager.NewArtifactSource(ctx, "artifactSource", &deploymentmanager.ArtifactSourceArgs{
// 			ArtifactSourceName: pulumi.String("myArtifactSource"),
// 			Authentication: &deploymentmanager.AuthenticationArgs{
// 				Type: pulumi.String("Sas"),
// 			},
// 			Location:          pulumi.String("centralus"),
// 			ResourceGroupName: pulumi.String("myResourceGroup"),
// 			SourceType:        pulumi.String("AzureStorage"),
// 			Tags:              nil,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### Create artifact source with artifact root, an offset into the storage container
//
// ```go
// package main
//
// import (
// 	deploymentmanager "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/deploymentmanager/v20191101preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := deploymentmanager.NewArtifactSource(ctx, "artifactSource", &deploymentmanager.ArtifactSourceArgs{
// 			ArtifactRoot:       pulumi.String("1.0.0.0"),
// 			ArtifactSourceName: pulumi.String("myArtifactSource"),
// 			Authentication: &deploymentmanager.AuthenticationArgs{
// 				Type: pulumi.String("Sas"),
// 			},
// 			Location:          pulumi.String("centralus"),
// 			ResourceGroupName: pulumi.String("myResourceGroup"),
// 			SourceType:        pulumi.String("AzureStorage"),
// 			Tags:              nil,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type ArtifactSource struct {
	pulumi.CustomResourceState

	// The path from the location that the 'authentication' property [say, a SAS URI to the blob container] refers to, to the location of the artifacts. This can be used to differentiate different versions of the artifacts. Or, different types of artifacts like binaries or templates. The location referenced by the authentication property concatenated with this optional artifactRoot path forms the artifact source location where the artifacts are expected to be found.
	ArtifactRoot pulumi.StringPtrOutput `pulumi:"artifactRoot"`
	// The authentication method to use to access the artifact source.
	Authentication SasAuthenticationResponseOutput `pulumi:"authentication"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of artifact source used.
	SourceType pulumi.StringOutput `pulumi:"sourceType"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewArtifactSource registers a new resource with the given unique name, arguments, and options.
func NewArtifactSource(ctx *pulumi.Context,
	name string, args *ArtifactSourceArgs, opts ...pulumi.ResourceOption) (*ArtifactSource, error) {
	if args == nil || args.ArtifactSourceName == nil {
		return nil, errors.New("missing required argument 'ArtifactSourceName'")
	}
	if args == nil || args.Authentication == nil {
		return nil, errors.New("missing required argument 'Authentication'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.SourceType == nil {
		return nil, errors.New("missing required argument 'SourceType'")
	}
	if args == nil {
		args = &ArtifactSourceArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:deploymentmanager/v20180901preview:ArtifactSource"),
		},
	})
	opts = append(opts, aliases)
	var resource ArtifactSource
	err := ctx.RegisterResource("azurerm:deploymentmanager/v20191101preview:ArtifactSource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetArtifactSource gets an existing ArtifactSource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetArtifactSource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ArtifactSourceState, opts ...pulumi.ResourceOption) (*ArtifactSource, error) {
	var resource ArtifactSource
	err := ctx.ReadResource("azurerm:deploymentmanager/v20191101preview:ArtifactSource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ArtifactSource resources.
type artifactSourceState struct {
	// The path from the location that the 'authentication' property [say, a SAS URI to the blob container] refers to, to the location of the artifacts. This can be used to differentiate different versions of the artifacts. Or, different types of artifacts like binaries or templates. The location referenced by the authentication property concatenated with this optional artifactRoot path forms the artifact source location where the artifacts are expected to be found.
	ArtifactRoot *string `pulumi:"artifactRoot"`
	// The authentication method to use to access the artifact source.
	Authentication *SasAuthenticationResponse `pulumi:"authentication"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The type of artifact source used.
	SourceType *string `pulumi:"sourceType"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
}

type ArtifactSourceState struct {
	// The path from the location that the 'authentication' property [say, a SAS URI to the blob container] refers to, to the location of the artifacts. This can be used to differentiate different versions of the artifacts. Or, different types of artifacts like binaries or templates. The location referenced by the authentication property concatenated with this optional artifactRoot path forms the artifact source location where the artifacts are expected to be found.
	ArtifactRoot pulumi.StringPtrInput
	// The authentication method to use to access the artifact source.
	Authentication SasAuthenticationResponsePtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// The type of artifact source used.
	SourceType pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
}

func (ArtifactSourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*artifactSourceState)(nil)).Elem()
}

type artifactSourceArgs struct {
	// The path from the location that the 'authentication' property [say, a SAS URI to the blob container] refers to, to the location of the artifacts. This can be used to differentiate different versions of the artifacts. Or, different types of artifacts like binaries or templates. The location referenced by the authentication property concatenated with this optional artifactRoot path forms the artifact source location where the artifacts are expected to be found.
	ArtifactRoot *string `pulumi:"artifactRoot"`
	// The name of the artifact source.
	ArtifactSourceName string `pulumi:"artifactSourceName"`
	// The authentication method to use to access the artifact source.
	Authentication SasAuthentication `pulumi:"authentication"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The type of artifact source used.
	SourceType string `pulumi:"sourceType"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ArtifactSource resource.
type ArtifactSourceArgs struct {
	// The path from the location that the 'authentication' property [say, a SAS URI to the blob container] refers to, to the location of the artifacts. This can be used to differentiate different versions of the artifacts. Or, different types of artifacts like binaries or templates. The location referenced by the authentication property concatenated with this optional artifactRoot path forms the artifact source location where the artifacts are expected to be found.
	ArtifactRoot pulumi.StringPtrInput
	// The name of the artifact source.
	ArtifactSourceName pulumi.StringInput
	// The authentication method to use to access the artifact source.
	Authentication SasAuthenticationInput
	// The geo-location where the resource lives
	Location pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The type of artifact source used.
	SourceType pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ArtifactSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*artifactSourceArgs)(nil)).Elem()
}
