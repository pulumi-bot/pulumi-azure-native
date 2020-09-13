// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An Application Insights workbook definition.
//
// ## Example Usage
// ### WorkbookAdd
//
// ```go
//
// ```
type Workbook struct {
	pulumi.CustomResourceState

	// Workbook category, as defined by the user at creation time.
	Category pulumi.StringOutput `pulumi:"category"`
	// The kind of workbook. Choices are user and shared.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Azure resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Configuration of this particular workbook. Configuration data is a string containing valid JSON
	SerializedData pulumi.StringOutput `pulumi:"serializedData"`
	// Enum indicating if this workbook definition is owned by a specific user or is shared between all users with access to the Application Insights component.
	SharedTypeKind pulumi.StringOutput `pulumi:"sharedTypeKind"`
	// Optional resourceId for a source resource.
	SourceResourceId pulumi.StringPtrOutput `pulumi:"sourceResourceId"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Date and time in UTC of the last modification that was made to this workbook definition.
	TimeModified pulumi.StringOutput `pulumi:"timeModified"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// Unique user id of the specific user that owns this workbook.
	UserId pulumi.StringOutput `pulumi:"userId"`
	// This instance's version of the data model. This can change as new features are added that can be marked workbook.
	Version pulumi.StringPtrOutput `pulumi:"version"`
	// Internally assigned unique id of the workbook definition.
	WorkbookId pulumi.StringOutput `pulumi:"workbookId"`
}

// NewWorkbook registers a new resource with the given unique name, arguments, and options.
func NewWorkbook(ctx *pulumi.Context,
	name string, args *WorkbookArgs, opts ...pulumi.ResourceOption) (*Workbook, error) {
	if args == nil || args.Category == nil {
		return nil, errors.New("missing required argument 'Category'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ResourceName == nil {
		return nil, errors.New("missing required argument 'ResourceName'")
	}
	if args == nil || args.SerializedData == nil {
		return nil, errors.New("missing required argument 'SerializedData'")
	}
	if args == nil || args.SharedTypeKind == nil {
		return nil, errors.New("missing required argument 'SharedTypeKind'")
	}
	if args == nil || args.UserId == nil {
		return nil, errors.New("missing required argument 'UserId'")
	}
	if args == nil || args.WorkbookId == nil {
		return nil, errors.New("missing required argument 'WorkbookId'")
	}
	if args == nil {
		args = &WorkbookArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:insights/v20150501:Workbook"),
		},
		{
			Type: pulumi.String("azurerm:insights/v20180617preview:Workbook"),
		},
	})
	opts = append(opts, aliases)
	var resource Workbook
	err := ctx.RegisterResource("azurerm:insights/latest:Workbook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkbook gets an existing Workbook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkbook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkbookState, opts ...pulumi.ResourceOption) (*Workbook, error) {
	var resource Workbook
	err := ctx.ReadResource("azurerm:insights/latest:Workbook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Workbook resources.
type workbookState struct {
	// Workbook category, as defined by the user at creation time.
	Category *string `pulumi:"category"`
	// The kind of workbook. Choices are user and shared.
	Kind *string `pulumi:"kind"`
	// Resource location
	Location *string `pulumi:"location"`
	// Azure resource name
	Name *string `pulumi:"name"`
	// Configuration of this particular workbook. Configuration data is a string containing valid JSON
	SerializedData *string `pulumi:"serializedData"`
	// Enum indicating if this workbook definition is owned by a specific user or is shared between all users with access to the Application Insights component.
	SharedTypeKind *string `pulumi:"sharedTypeKind"`
	// Optional resourceId for a source resource.
	SourceResourceId *string `pulumi:"sourceResourceId"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Date and time in UTC of the last modification that was made to this workbook definition.
	TimeModified *string `pulumi:"timeModified"`
	// Azure resource type
	Type *string `pulumi:"type"`
	// Unique user id of the specific user that owns this workbook.
	UserId *string `pulumi:"userId"`
	// This instance's version of the data model. This can change as new features are added that can be marked workbook.
	Version *string `pulumi:"version"`
	// Internally assigned unique id of the workbook definition.
	WorkbookId *string `pulumi:"workbookId"`
}

type WorkbookState struct {
	// Workbook category, as defined by the user at creation time.
	Category pulumi.StringPtrInput
	// The kind of workbook. Choices are user and shared.
	Kind pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Azure resource name
	Name pulumi.StringPtrInput
	// Configuration of this particular workbook. Configuration data is a string containing valid JSON
	SerializedData pulumi.StringPtrInput
	// Enum indicating if this workbook definition is owned by a specific user or is shared between all users with access to the Application Insights component.
	SharedTypeKind pulumi.StringPtrInput
	// Optional resourceId for a source resource.
	SourceResourceId pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Date and time in UTC of the last modification that was made to this workbook definition.
	TimeModified pulumi.StringPtrInput
	// Azure resource type
	Type pulumi.StringPtrInput
	// Unique user id of the specific user that owns this workbook.
	UserId pulumi.StringPtrInput
	// This instance's version of the data model. This can change as new features are added that can be marked workbook.
	Version pulumi.StringPtrInput
	// Internally assigned unique id of the workbook definition.
	WorkbookId pulumi.StringPtrInput
}

func (WorkbookState) ElementType() reflect.Type {
	return reflect.TypeOf((*workbookState)(nil)).Elem()
}

type workbookArgs struct {
	// Workbook category, as defined by the user at creation time.
	Category string `pulumi:"category"`
	// The kind of workbook. Choices are user and shared.
	Kind *string `pulumi:"kind"`
	// Resource location
	Location *string `pulumi:"location"`
	// The user-defined name of the workbook.
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Application Insights component resource.
	ResourceName string `pulumi:"resourceName"`
	// Configuration of this particular workbook. Configuration data is a string containing valid JSON
	SerializedData string `pulumi:"serializedData"`
	// Enum indicating if this workbook definition is owned by a specific user or is shared between all users with access to the Application Insights component.
	SharedTypeKind string `pulumi:"sharedTypeKind"`
	// Optional resourceId for a source resource.
	SourceResourceId *string `pulumi:"sourceResourceId"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Unique user id of the specific user that owns this workbook.
	UserId string `pulumi:"userId"`
	// This instance's version of the data model. This can change as new features are added that can be marked workbook.
	Version *string `pulumi:"version"`
	// Internally assigned unique id of the workbook definition.
	WorkbookId string `pulumi:"workbookId"`
}

// The set of arguments for constructing a Workbook resource.
type WorkbookArgs struct {
	// Workbook category, as defined by the user at creation time.
	Category pulumi.StringInput
	// The kind of workbook. Choices are user and shared.
	Kind pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// The user-defined name of the workbook.
	Name pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the Application Insights component resource.
	ResourceName pulumi.StringInput
	// Configuration of this particular workbook. Configuration data is a string containing valid JSON
	SerializedData pulumi.StringInput
	// Enum indicating if this workbook definition is owned by a specific user or is shared between all users with access to the Application Insights component.
	SharedTypeKind pulumi.StringInput
	// Optional resourceId for a source resource.
	SourceResourceId pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Unique user id of the specific user that owns this workbook.
	UserId pulumi.StringInput
	// This instance's version of the data model. This can change as new features are added that can be marked workbook.
	Version pulumi.StringPtrInput
	// Internally assigned unique id of the workbook definition.
	WorkbookId pulumi.StringInput
}

func (WorkbookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workbookArgs)(nil)).Elem()
}
