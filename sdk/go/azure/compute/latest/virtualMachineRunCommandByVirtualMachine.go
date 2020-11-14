// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Describes a Virtual Machine run command.
type VirtualMachineRunCommandByVirtualMachine struct {
	pulumi.CustomResourceState

	// Optional. If set to true, provisioning will complete as soon as the script starts and will not wait for script to complete.
	AsyncExecution pulumi.BoolPtrOutput `pulumi:"asyncExecution"`
	// Specifies the Azure storage blob where script error stream will be uploaded.
	ErrorBlobUri pulumi.StringPtrOutput `pulumi:"errorBlobUri"`
	// The virtual machine run command instance view.
	InstanceView VirtualMachineRunCommandInstanceViewResponseOutput `pulumi:"instanceView"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the Azure storage blob where script output stream will be uploaded.
	OutputBlobUri pulumi.StringPtrOutput `pulumi:"outputBlobUri"`
	// The parameters used by the script.
	Parameters RunCommandInputParameterResponseArrayOutput `pulumi:"parameters"`
	// The parameters used by the script.
	ProtectedParameters RunCommandInputParameterResponseArrayOutput `pulumi:"protectedParameters"`
	// The provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Specifies the user account password on the VM when executing the run command.
	RunAsPassword pulumi.StringPtrOutput `pulumi:"runAsPassword"`
	// Specifies the user account on the VM when executing the run command.
	RunAsUser pulumi.StringPtrOutput `pulumi:"runAsUser"`
	// The source of the run command script.
	Source VirtualMachineRunCommandScriptSourceResponsePtrOutput `pulumi:"source"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The timeout in seconds to execute the run command.
	TimeoutInSeconds pulumi.IntPtrOutput `pulumi:"timeoutInSeconds"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewVirtualMachineRunCommandByVirtualMachine registers a new resource with the given unique name, arguments, and options.
func NewVirtualMachineRunCommandByVirtualMachine(ctx *pulumi.Context,
	name string, args *VirtualMachineRunCommandByVirtualMachineArgs, opts ...pulumi.ResourceOption) (*VirtualMachineRunCommandByVirtualMachine, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.RunCommandName == nil {
		return nil, errors.New("missing required argument 'RunCommandName'")
	}
	if args == nil || args.VmName == nil {
		return nil, errors.New("missing required argument 'VmName'")
	}
	if args == nil {
		args = &VirtualMachineRunCommandByVirtualMachineArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:compute/v20200601:VirtualMachineRunCommandByVirtualMachine"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualMachineRunCommandByVirtualMachine
	err := ctx.RegisterResource("azure-nextgen:compute/latest:VirtualMachineRunCommandByVirtualMachine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualMachineRunCommandByVirtualMachine gets an existing VirtualMachineRunCommandByVirtualMachine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualMachineRunCommandByVirtualMachine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineRunCommandByVirtualMachineState, opts ...pulumi.ResourceOption) (*VirtualMachineRunCommandByVirtualMachine, error) {
	var resource VirtualMachineRunCommandByVirtualMachine
	err := ctx.ReadResource("azure-nextgen:compute/latest:VirtualMachineRunCommandByVirtualMachine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualMachineRunCommandByVirtualMachine resources.
type virtualMachineRunCommandByVirtualMachineState struct {
	// Optional. If set to true, provisioning will complete as soon as the script starts and will not wait for script to complete.
	AsyncExecution *bool `pulumi:"asyncExecution"`
	// Specifies the Azure storage blob where script error stream will be uploaded.
	ErrorBlobUri *string `pulumi:"errorBlobUri"`
	// The virtual machine run command instance view.
	InstanceView *VirtualMachineRunCommandInstanceViewResponse `pulumi:"instanceView"`
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// Specifies the Azure storage blob where script output stream will be uploaded.
	OutputBlobUri *string `pulumi:"outputBlobUri"`
	// The parameters used by the script.
	Parameters []RunCommandInputParameterResponse `pulumi:"parameters"`
	// The parameters used by the script.
	ProtectedParameters []RunCommandInputParameterResponse `pulumi:"protectedParameters"`
	// The provisioning state, which only appears in the response.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Specifies the user account password on the VM when executing the run command.
	RunAsPassword *string `pulumi:"runAsPassword"`
	// Specifies the user account on the VM when executing the run command.
	RunAsUser *string `pulumi:"runAsUser"`
	// The source of the run command script.
	Source *VirtualMachineRunCommandScriptSourceResponse `pulumi:"source"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// The timeout in seconds to execute the run command.
	TimeoutInSeconds *int `pulumi:"timeoutInSeconds"`
	// Resource type
	Type *string `pulumi:"type"`
}

type VirtualMachineRunCommandByVirtualMachineState struct {
	// Optional. If set to true, provisioning will complete as soon as the script starts and will not wait for script to complete.
	AsyncExecution pulumi.BoolPtrInput
	// Specifies the Azure storage blob where script error stream will be uploaded.
	ErrorBlobUri pulumi.StringPtrInput
	// The virtual machine run command instance view.
	InstanceView VirtualMachineRunCommandInstanceViewResponsePtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Specifies the Azure storage blob where script output stream will be uploaded.
	OutputBlobUri pulumi.StringPtrInput
	// The parameters used by the script.
	Parameters RunCommandInputParameterResponseArrayInput
	// The parameters used by the script.
	ProtectedParameters RunCommandInputParameterResponseArrayInput
	// The provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringPtrInput
	// Specifies the user account password on the VM when executing the run command.
	RunAsPassword pulumi.StringPtrInput
	// Specifies the user account on the VM when executing the run command.
	RunAsUser pulumi.StringPtrInput
	// The source of the run command script.
	Source VirtualMachineRunCommandScriptSourceResponsePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// The timeout in seconds to execute the run command.
	TimeoutInSeconds pulumi.IntPtrInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (VirtualMachineRunCommandByVirtualMachineState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineRunCommandByVirtualMachineState)(nil)).Elem()
}

type virtualMachineRunCommandByVirtualMachineArgs struct {
	// Optional. If set to true, provisioning will complete as soon as the script starts and will not wait for script to complete.
	AsyncExecution *bool `pulumi:"asyncExecution"`
	// Specifies the Azure storage blob where script error stream will be uploaded.
	ErrorBlobUri *string `pulumi:"errorBlobUri"`
	// Resource location
	Location string `pulumi:"location"`
	// Specifies the Azure storage blob where script output stream will be uploaded.
	OutputBlobUri *string `pulumi:"outputBlobUri"`
	// The parameters used by the script.
	Parameters []RunCommandInputParameter `pulumi:"parameters"`
	// The parameters used by the script.
	ProtectedParameters []RunCommandInputParameter `pulumi:"protectedParameters"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the user account password on the VM when executing the run command.
	RunAsPassword *string `pulumi:"runAsPassword"`
	// Specifies the user account on the VM when executing the run command.
	RunAsUser *string `pulumi:"runAsUser"`
	// The name of the virtual machine run command.
	RunCommandName string `pulumi:"runCommandName"`
	// The source of the run command script.
	Source *VirtualMachineRunCommandScriptSource `pulumi:"source"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// The timeout in seconds to execute the run command.
	TimeoutInSeconds *int `pulumi:"timeoutInSeconds"`
	// The name of the virtual machine where the run command should be created or updated.
	VmName string `pulumi:"vmName"`
}

// The set of arguments for constructing a VirtualMachineRunCommandByVirtualMachine resource.
type VirtualMachineRunCommandByVirtualMachineArgs struct {
	// Optional. If set to true, provisioning will complete as soon as the script starts and will not wait for script to complete.
	AsyncExecution pulumi.BoolPtrInput
	// Specifies the Azure storage blob where script error stream will be uploaded.
	ErrorBlobUri pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringInput
	// Specifies the Azure storage blob where script output stream will be uploaded.
	OutputBlobUri pulumi.StringPtrInput
	// The parameters used by the script.
	Parameters RunCommandInputParameterArrayInput
	// The parameters used by the script.
	ProtectedParameters RunCommandInputParameterArrayInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Specifies the user account password on the VM when executing the run command.
	RunAsPassword pulumi.StringPtrInput
	// Specifies the user account on the VM when executing the run command.
	RunAsUser pulumi.StringPtrInput
	// The name of the virtual machine run command.
	RunCommandName pulumi.StringInput
	// The source of the run command script.
	Source VirtualMachineRunCommandScriptSourcePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// The timeout in seconds to execute the run command.
	TimeoutInSeconds pulumi.IntPtrInput
	// The name of the virtual machine where the run command should be created or updated.
	VmName pulumi.StringInput
}

func (VirtualMachineRunCommandByVirtualMachineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineRunCommandByVirtualMachineArgs)(nil)).Elem()
}

type VirtualMachineRunCommandByVirtualMachineInput interface {
	pulumi.Input

	ToVirtualMachineRunCommandByVirtualMachineOutput() VirtualMachineRunCommandByVirtualMachineOutput
	ToVirtualMachineRunCommandByVirtualMachineOutputWithContext(ctx context.Context) VirtualMachineRunCommandByVirtualMachineOutput
}

func (VirtualMachineRunCommandByVirtualMachine) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineRunCommandByVirtualMachine)(nil)).Elem()
}

func (i VirtualMachineRunCommandByVirtualMachine) ToVirtualMachineRunCommandByVirtualMachineOutput() VirtualMachineRunCommandByVirtualMachineOutput {
	return i.ToVirtualMachineRunCommandByVirtualMachineOutputWithContext(context.Background())
}

func (i VirtualMachineRunCommandByVirtualMachine) ToVirtualMachineRunCommandByVirtualMachineOutputWithContext(ctx context.Context) VirtualMachineRunCommandByVirtualMachineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineRunCommandByVirtualMachineOutput)
}

type VirtualMachineRunCommandByVirtualMachineOutput struct {
	*pulumi.OutputState
}

func (VirtualMachineRunCommandByVirtualMachineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineRunCommandByVirtualMachineOutput)(nil)).Elem()
}

func (o VirtualMachineRunCommandByVirtualMachineOutput) ToVirtualMachineRunCommandByVirtualMachineOutput() VirtualMachineRunCommandByVirtualMachineOutput {
	return o
}

func (o VirtualMachineRunCommandByVirtualMachineOutput) ToVirtualMachineRunCommandByVirtualMachineOutputWithContext(ctx context.Context) VirtualMachineRunCommandByVirtualMachineOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineRunCommandByVirtualMachineOutput{})
}
