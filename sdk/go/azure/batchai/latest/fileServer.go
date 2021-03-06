// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// File Server information.
// Latest API Version: 2018-05-01.
//
// Deprecated: The 'latest' version is deprecated. Please migrate to the resource in the top-level module: 'azure-native:batchai:FileServer'.
type FileServer struct {
	pulumi.CustomResourceState

	// Time when the FileServer was created.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// Information about disks attached to File Server VM.
	DataDisks DataDisksResponsePtrOutput `pulumi:"dataDisks"`
	// File Server mount settings.
	MountSettings MountSettingsResponseOutput `pulumi:"mountSettings"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the File Server. Possible values: creating - The File Server is getting created; updating - The File Server creation has been accepted and it is getting updated; deleting - The user has requested that the File Server be deleted, and it is in the process of being deleted; failed - The File Server creation has failed with the specified error code. Details about the error code are specified in the message field; succeeded - The File Server creation has succeeded.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Time when the provisioning state was changed.
	ProvisioningStateTransitionTime pulumi.StringOutput `pulumi:"provisioningStateTransitionTime"`
	// SSH configuration for accessing the File Server node.
	SshConfiguration SshConfigurationResponsePtrOutput `pulumi:"sshConfiguration"`
	// File Server virtual network subnet resource ID.
	Subnet ResourceIdResponsePtrOutput `pulumi:"subnet"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// VM size of the File Server.
	VmSize pulumi.StringPtrOutput `pulumi:"vmSize"`
}

// NewFileServer registers a new resource with the given unique name, arguments, and options.
func NewFileServer(ctx *pulumi.Context,
	name string, args *FileServerArgs, opts ...pulumi.ResourceOption) (*FileServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataDisks == nil {
		return nil, errors.New("invalid value for required argument 'DataDisks'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SshConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'SshConfiguration'")
	}
	if args.VmSize == nil {
		return nil, errors.New("invalid value for required argument 'VmSize'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:batchai/latest:FileServer"),
		},
		{
			Type: pulumi.String("azure-native:batchai:FileServer"),
		},
		{
			Type: pulumi.String("azure-nextgen:batchai:FileServer"),
		},
		{
			Type: pulumi.String("azure-native:batchai/v20180501:FileServer"),
		},
		{
			Type: pulumi.String("azure-nextgen:batchai/v20180501:FileServer"),
		},
	})
	opts = append(opts, aliases)
	var resource FileServer
	err := ctx.RegisterResource("azure-native:batchai/latest:FileServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFileServer gets an existing FileServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFileServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FileServerState, opts ...pulumi.ResourceOption) (*FileServer, error) {
	var resource FileServer
	err := ctx.ReadResource("azure-native:batchai/latest:FileServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FileServer resources.
type fileServerState struct {
	// Time when the FileServer was created.
	CreationTime *string `pulumi:"creationTime"`
	// Information about disks attached to File Server VM.
	DataDisks *DataDisksResponse `pulumi:"dataDisks"`
	// File Server mount settings.
	MountSettings *MountSettingsResponse `pulumi:"mountSettings"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// Provisioning state of the File Server. Possible values: creating - The File Server is getting created; updating - The File Server creation has been accepted and it is getting updated; deleting - The user has requested that the File Server be deleted, and it is in the process of being deleted; failed - The File Server creation has failed with the specified error code. Details about the error code are specified in the message field; succeeded - The File Server creation has succeeded.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Time when the provisioning state was changed.
	ProvisioningStateTransitionTime *string `pulumi:"provisioningStateTransitionTime"`
	// SSH configuration for accessing the File Server node.
	SshConfiguration *SshConfigurationResponse `pulumi:"sshConfiguration"`
	// File Server virtual network subnet resource ID.
	Subnet *ResourceIdResponse `pulumi:"subnet"`
	// The type of the resource.
	Type *string `pulumi:"type"`
	// VM size of the File Server.
	VmSize *string `pulumi:"vmSize"`
}

type FileServerState struct {
	// Time when the FileServer was created.
	CreationTime pulumi.StringPtrInput
	// Information about disks attached to File Server VM.
	DataDisks DataDisksResponsePtrInput
	// File Server mount settings.
	MountSettings MountSettingsResponsePtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// Provisioning state of the File Server. Possible values: creating - The File Server is getting created; updating - The File Server creation has been accepted and it is getting updated; deleting - The user has requested that the File Server be deleted, and it is in the process of being deleted; failed - The File Server creation has failed with the specified error code. Details about the error code are specified in the message field; succeeded - The File Server creation has succeeded.
	ProvisioningState pulumi.StringPtrInput
	// Time when the provisioning state was changed.
	ProvisioningStateTransitionTime pulumi.StringPtrInput
	// SSH configuration for accessing the File Server node.
	SshConfiguration SshConfigurationResponsePtrInput
	// File Server virtual network subnet resource ID.
	Subnet ResourceIdResponsePtrInput
	// The type of the resource.
	Type pulumi.StringPtrInput
	// VM size of the File Server.
	VmSize pulumi.StringPtrInput
}

func (FileServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*fileServerState)(nil)).Elem()
}

type fileServerArgs struct {
	// Settings for the data disks which will be created for the File Server.
	DataDisks DataDisks `pulumi:"dataDisks"`
	// The name of the file server within the specified resource group. File server names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
	FileServerName *string `pulumi:"fileServerName"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// SSH configuration for the File Server node.
	SshConfiguration SshConfiguration `pulumi:"sshConfiguration"`
	// Identifier of an existing virtual network subnet to put the File Server in. If not provided, a new virtual network and subnet will be created.
	Subnet *ResourceId `pulumi:"subnet"`
	// The size of the virtual machine for the File Server. For information about available VM sizes from the Virtual Machines Marketplace, see Sizes for Virtual Machines (Linux).
	VmSize string `pulumi:"vmSize"`
	// The name of the workspace. Workspace names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a FileServer resource.
type FileServerArgs struct {
	// Settings for the data disks which will be created for the File Server.
	DataDisks DataDisksInput
	// The name of the file server within the specified resource group. File server names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
	FileServerName pulumi.StringPtrInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// SSH configuration for the File Server node.
	SshConfiguration SshConfigurationInput
	// Identifier of an existing virtual network subnet to put the File Server in. If not provided, a new virtual network and subnet will be created.
	Subnet ResourceIdPtrInput
	// The size of the virtual machine for the File Server. For information about available VM sizes from the Virtual Machines Marketplace, see Sizes for Virtual Machines (Linux).
	VmSize pulumi.StringInput
	// The name of the workspace. Workspace names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
	WorkspaceName pulumi.StringInput
}

func (FileServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fileServerArgs)(nil)).Elem()
}

type FileServerInput interface {
	pulumi.Input

	ToFileServerOutput() FileServerOutput
	ToFileServerOutputWithContext(ctx context.Context) FileServerOutput
}

func (*FileServer) ElementType() reflect.Type {
	return reflect.TypeOf((*FileServer)(nil))
}

func (i *FileServer) ToFileServerOutput() FileServerOutput {
	return i.ToFileServerOutputWithContext(context.Background())
}

func (i *FileServer) ToFileServerOutputWithContext(ctx context.Context) FileServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileServerOutput)
}

type FileServerOutput struct {
	*pulumi.OutputState
}

func (FileServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FileServer)(nil))
}

func (o FileServerOutput) ToFileServerOutput() FileServerOutput {
	return o
}

func (o FileServerOutput) ToFileServerOutputWithContext(ctx context.Context) FileServerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(FileServerOutput{})
}
