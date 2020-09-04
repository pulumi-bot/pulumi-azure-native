// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190924preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Schema for Application properties.
type Application struct {
	pulumi.CustomResourceState

	// Command Line Arguments for Application.
	CommandLineArguments pulumi.StringPtrOutput `pulumi:"commandLineArguments"`
	// Specifies whether this published application can be launched with command line arguments provided by the client, command line arguments specified at publish time, or no command line arguments at all.
	CommandLineSetting pulumi.StringOutput `pulumi:"commandLineSetting"`
	// Description of Application.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies a path for the executable file for the application.
	FilePath pulumi.StringPtrOutput `pulumi:"filePath"`
	// Friendly name of Application.
	FriendlyName pulumi.StringPtrOutput `pulumi:"friendlyName"`
	// the icon a 64 bit string as a byte array.
	IconContent pulumi.StringOutput `pulumi:"iconContent"`
	// Hash of the icon.
	IconHash pulumi.StringOutput `pulumi:"iconHash"`
	// Index of the icon.
	IconIndex pulumi.IntPtrOutput `pulumi:"iconIndex"`
	// Path to icon.
	IconPath pulumi.StringPtrOutput `pulumi:"iconPath"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies whether to show the RemoteApp program in the RD Web Access server.
	ShowInPortal pulumi.BoolPtrOutput `pulumi:"showInPortal"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil || args.ApplicationGroupName == nil {
		return nil, errors.New("missing required argument 'ApplicationGroupName'")
	}
	if args == nil || args.ApplicationName == nil {
		return nil, errors.New("missing required argument 'ApplicationName'")
	}
	if args == nil || args.CommandLineSetting == nil {
		return nil, errors.New("missing required argument 'CommandLineSetting'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ApplicationArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:desktopvirtualization/v20190123preview:Application"),
		},
		{
			Type: pulumi.String("azurerm:desktopvirtualization/v20191210preview:Application"),
		},
	})
	opts = append(opts, aliases)
	var resource Application
	err := ctx.RegisterResource("azurerm:desktopvirtualization/v20190924preview:Application", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplication gets an existing Application resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationState, opts ...pulumi.ResourceOption) (*Application, error) {
	var resource Application
	err := ctx.ReadResource("azurerm:desktopvirtualization/v20190924preview:Application", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Application resources.
type applicationState struct {
	// Command Line Arguments for Application.
	CommandLineArguments *string `pulumi:"commandLineArguments"`
	// Specifies whether this published application can be launched with command line arguments provided by the client, command line arguments specified at publish time, or no command line arguments at all.
	CommandLineSetting *string `pulumi:"commandLineSetting"`
	// Description of Application.
	Description *string `pulumi:"description"`
	// Specifies a path for the executable file for the application.
	FilePath *string `pulumi:"filePath"`
	// Friendly name of Application.
	FriendlyName *string `pulumi:"friendlyName"`
	// the icon a 64 bit string as a byte array.
	IconContent *string `pulumi:"iconContent"`
	// Hash of the icon.
	IconHash *string `pulumi:"iconHash"`
	// Index of the icon.
	IconIndex *int `pulumi:"iconIndex"`
	// Path to icon.
	IconPath *string `pulumi:"iconPath"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// Specifies whether to show the RemoteApp program in the RD Web Access server.
	ShowInPortal *bool `pulumi:"showInPortal"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
}

type ApplicationState struct {
	// Command Line Arguments for Application.
	CommandLineArguments pulumi.StringPtrInput
	// Specifies whether this published application can be launched with command line arguments provided by the client, command line arguments specified at publish time, or no command line arguments at all.
	CommandLineSetting pulumi.StringPtrInput
	// Description of Application.
	Description pulumi.StringPtrInput
	// Specifies a path for the executable file for the application.
	FilePath pulumi.StringPtrInput
	// Friendly name of Application.
	FriendlyName pulumi.StringPtrInput
	// the icon a 64 bit string as a byte array.
	IconContent pulumi.StringPtrInput
	// Hash of the icon.
	IconHash pulumi.StringPtrInput
	// Index of the icon.
	IconIndex pulumi.IntPtrInput
	// Path to icon.
	IconPath pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// Specifies whether to show the RemoteApp program in the RD Web Access server.
	ShowInPortal pulumi.BoolPtrInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
}

func (ApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationState)(nil)).Elem()
}

type applicationArgs struct {
	// The name of the application group
	ApplicationGroupName string `pulumi:"applicationGroupName"`
	// The name of the application within the specified application group
	ApplicationName string `pulumi:"applicationName"`
	// Command Line Arguments for Application.
	CommandLineArguments *string `pulumi:"commandLineArguments"`
	// Specifies whether this published application can be launched with command line arguments provided by the client, command line arguments specified at publish time, or no command line arguments at all.
	CommandLineSetting string `pulumi:"commandLineSetting"`
	// Description of Application.
	Description *string `pulumi:"description"`
	// Specifies a path for the executable file for the application.
	FilePath *string `pulumi:"filePath"`
	// Friendly name of Application.
	FriendlyName *string `pulumi:"friendlyName"`
	// Index of the icon.
	IconIndex *int `pulumi:"iconIndex"`
	// Path to icon.
	IconPath *string `pulumi:"iconPath"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies whether to show the RemoteApp program in the RD Web Access server.
	ShowInPortal *bool `pulumi:"showInPortal"`
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	// The name of the application group
	ApplicationGroupName pulumi.StringInput
	// The name of the application within the specified application group
	ApplicationName pulumi.StringInput
	// Command Line Arguments for Application.
	CommandLineArguments pulumi.StringPtrInput
	// Specifies whether this published application can be launched with command line arguments provided by the client, command line arguments specified at publish time, or no command line arguments at all.
	CommandLineSetting pulumi.StringInput
	// Description of Application.
	Description pulumi.StringPtrInput
	// Specifies a path for the executable file for the application.
	FilePath pulumi.StringPtrInput
	// Friendly name of Application.
	FriendlyName pulumi.StringPtrInput
	// Index of the icon.
	IconIndex pulumi.IntPtrInput
	// Path to icon.
	IconPath pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Specifies whether to show the RemoteApp program in the RD Web Access server.
	ShowInPortal pulumi.BoolPtrInput
}

func (ApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationArgs)(nil)).Elem()
}
