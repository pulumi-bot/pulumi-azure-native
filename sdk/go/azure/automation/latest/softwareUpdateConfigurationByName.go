// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Software update configuration properties.
// Latest API Version: 2019-06-01.
type SoftwareUpdateConfigurationByName struct {
	pulumi.CustomResourceState

	// CreatedBy property, which only appears in the response.
	CreatedBy pulumi.StringOutput `pulumi:"createdBy"`
	// Creation time of the resource, which only appears in the response.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// Details of provisioning error
	Error ErrorResponseResponsePtrOutput `pulumi:"error"`
	// LastModifiedBy property, which only appears in the response.
	LastModifiedBy pulumi.StringOutput `pulumi:"lastModifiedBy"`
	// Last time resource was modified, which only appears in the response.
	LastModifiedTime pulumi.StringOutput `pulumi:"lastModifiedTime"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state for the software update configuration, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Schedule information for the Software update configuration
	ScheduleInfo SUCSchedulePropertiesResponseOutput `pulumi:"scheduleInfo"`
	// Tasks information for the Software update configuration.
	Tasks SoftwareUpdateConfigurationTasksResponsePtrOutput `pulumi:"tasks"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// update specific properties for the Software update configuration
	UpdateConfiguration UpdateConfigurationResponseOutput `pulumi:"updateConfiguration"`
}

// NewSoftwareUpdateConfigurationByName registers a new resource with the given unique name, arguments, and options.
func NewSoftwareUpdateConfigurationByName(ctx *pulumi.Context,
	name string, args *SoftwareUpdateConfigurationByNameArgs, opts ...pulumi.ResourceOption) (*SoftwareUpdateConfigurationByName, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ScheduleInfo == nil {
		return nil, errors.New("invalid value for required argument 'ScheduleInfo'")
	}
	if args.SoftwareUpdateConfigurationName == nil {
		return nil, errors.New("invalid value for required argument 'SoftwareUpdateConfigurationName'")
	}
	if args.UpdateConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'UpdateConfiguration'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:automation/v20170515preview:SoftwareUpdateConfigurationByName"),
		},
		{
			Type: pulumi.String("azure-nextgen:automation/v20190601:SoftwareUpdateConfigurationByName"),
		},
	})
	opts = append(opts, aliases)
	var resource SoftwareUpdateConfigurationByName
	err := ctx.RegisterResource("azure-nextgen:automation/latest:SoftwareUpdateConfigurationByName", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSoftwareUpdateConfigurationByName gets an existing SoftwareUpdateConfigurationByName resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSoftwareUpdateConfigurationByName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SoftwareUpdateConfigurationByNameState, opts ...pulumi.ResourceOption) (*SoftwareUpdateConfigurationByName, error) {
	var resource SoftwareUpdateConfigurationByName
	err := ctx.ReadResource("azure-nextgen:automation/latest:SoftwareUpdateConfigurationByName", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SoftwareUpdateConfigurationByName resources.
type softwareUpdateConfigurationByNameState struct {
	// CreatedBy property, which only appears in the response.
	CreatedBy *string `pulumi:"createdBy"`
	// Creation time of the resource, which only appears in the response.
	CreationTime *string `pulumi:"creationTime"`
	// Details of provisioning error
	Error *ErrorResponseResponse `pulumi:"error"`
	// LastModifiedBy property, which only appears in the response.
	LastModifiedBy *string `pulumi:"lastModifiedBy"`
	// Last time resource was modified, which only appears in the response.
	LastModifiedTime *string `pulumi:"lastModifiedTime"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Provisioning state for the software update configuration, which only appears in the response.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Schedule information for the Software update configuration
	ScheduleInfo *SUCSchedulePropertiesResponse `pulumi:"scheduleInfo"`
	// Tasks information for the Software update configuration.
	Tasks *SoftwareUpdateConfigurationTasksResponse `pulumi:"tasks"`
	// Resource type
	Type *string `pulumi:"type"`
	// update specific properties for the Software update configuration
	UpdateConfiguration *UpdateConfigurationResponse `pulumi:"updateConfiguration"`
}

type SoftwareUpdateConfigurationByNameState struct {
	// CreatedBy property, which only appears in the response.
	CreatedBy pulumi.StringPtrInput
	// Creation time of the resource, which only appears in the response.
	CreationTime pulumi.StringPtrInput
	// Details of provisioning error
	Error ErrorResponseResponsePtrInput
	// LastModifiedBy property, which only appears in the response.
	LastModifiedBy pulumi.StringPtrInput
	// Last time resource was modified, which only appears in the response.
	LastModifiedTime pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Provisioning state for the software update configuration, which only appears in the response.
	ProvisioningState pulumi.StringPtrInput
	// Schedule information for the Software update configuration
	ScheduleInfo SUCSchedulePropertiesResponsePtrInput
	// Tasks information for the Software update configuration.
	Tasks SoftwareUpdateConfigurationTasksResponsePtrInput
	// Resource type
	Type pulumi.StringPtrInput
	// update specific properties for the Software update configuration
	UpdateConfiguration UpdateConfigurationResponsePtrInput
}

func (SoftwareUpdateConfigurationByNameState) ElementType() reflect.Type {
	return reflect.TypeOf((*softwareUpdateConfigurationByNameState)(nil)).Elem()
}

type softwareUpdateConfigurationByNameArgs struct {
	// The name of the automation account.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// Details of provisioning error
	Error *ErrorResponse `pulumi:"error"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Schedule information for the Software update configuration
	ScheduleInfo SUCScheduleProperties `pulumi:"scheduleInfo"`
	// The name of the software update configuration to be created.
	SoftwareUpdateConfigurationName string `pulumi:"softwareUpdateConfigurationName"`
	// Tasks information for the Software update configuration.
	Tasks *SoftwareUpdateConfigurationTasks `pulumi:"tasks"`
	// update specific properties for the Software update configuration
	UpdateConfiguration UpdateConfiguration `pulumi:"updateConfiguration"`
}

// The set of arguments for constructing a SoftwareUpdateConfigurationByName resource.
type SoftwareUpdateConfigurationByNameArgs struct {
	// The name of the automation account.
	AutomationAccountName pulumi.StringInput
	// Details of provisioning error
	Error ErrorResponsePtrInput
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput
	// Schedule information for the Software update configuration
	ScheduleInfo SUCSchedulePropertiesInput
	// The name of the software update configuration to be created.
	SoftwareUpdateConfigurationName pulumi.StringInput
	// Tasks information for the Software update configuration.
	Tasks SoftwareUpdateConfigurationTasksPtrInput
	// update specific properties for the Software update configuration
	UpdateConfiguration UpdateConfigurationInput
}

func (SoftwareUpdateConfigurationByNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*softwareUpdateConfigurationByNameArgs)(nil)).Elem()
}

type SoftwareUpdateConfigurationByNameInput interface {
	pulumi.Input

	ToSoftwareUpdateConfigurationByNameOutput() SoftwareUpdateConfigurationByNameOutput
	ToSoftwareUpdateConfigurationByNameOutputWithContext(ctx context.Context) SoftwareUpdateConfigurationByNameOutput
}

func (*SoftwareUpdateConfigurationByName) ElementType() reflect.Type {
	return reflect.TypeOf((*SoftwareUpdateConfigurationByName)(nil))
}

func (i *SoftwareUpdateConfigurationByName) ToSoftwareUpdateConfigurationByNameOutput() SoftwareUpdateConfigurationByNameOutput {
	return i.ToSoftwareUpdateConfigurationByNameOutputWithContext(context.Background())
}

func (i *SoftwareUpdateConfigurationByName) ToSoftwareUpdateConfigurationByNameOutputWithContext(ctx context.Context) SoftwareUpdateConfigurationByNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SoftwareUpdateConfigurationByNameOutput)
}

type SoftwareUpdateConfigurationByNameOutput struct {
	*pulumi.OutputState
}

func (SoftwareUpdateConfigurationByNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SoftwareUpdateConfigurationByName)(nil))
}

func (o SoftwareUpdateConfigurationByNameOutput) ToSoftwareUpdateConfigurationByNameOutput() SoftwareUpdateConfigurationByNameOutput {
	return o
}

func (o SoftwareUpdateConfigurationByNameOutput) ToSoftwareUpdateConfigurationByNameOutputWithContext(ctx context.Context) SoftwareUpdateConfigurationByNameOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SoftwareUpdateConfigurationByNameOutput{})
}
