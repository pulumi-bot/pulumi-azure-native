// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20181201

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Contains information about an application in a Batch account.
type Application struct {
	pulumi.CustomResourceState

	// The ETag of the resource, used for concurrency statements.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties associated with the Application.
	Properties ApplicationPropertiesResponseOutput `pulumi:"properties"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ApplicationArgs{}
	}
	var resource Application
	err := ctx.RegisterResource("azurerm:batch/v20181201:Application", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:batch/v20181201:Application", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Application resources.
type applicationState struct {
	// The ETag of the resource, used for concurrency statements.
	Etag *string `pulumi:"etag"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The properties associated with the Application.
	Properties *ApplicationPropertiesResponse `pulumi:"properties"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type ApplicationState struct {
	// The ETag of the resource, used for concurrency statements.
	Etag pulumi.StringPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The properties associated with the Application.
	Properties ApplicationPropertiesResponsePtrInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (ApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationState)(nil)).Elem()
}

type applicationArgs struct {
	// The name of the Batch account.
	AccountName string `pulumi:"accountName"`
	// A value indicating whether packages within the application may be overwritten using the same version string.
	AllowUpdates *bool `pulumi:"allowUpdates"`
	// The package to use if a client requests the application but does not specify a version. This property can only be set to the name of an existing package.
	DefaultVersion *string `pulumi:"defaultVersion"`
	// The display name for the application.
	DisplayName *string `pulumi:"displayName"`
	// The name of the application. This must be unique within the account.
	Name string `pulumi:"name"`
	// The name of the resource group that contains the Batch account.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	// The name of the Batch account.
	AccountName pulumi.StringInput
	// A value indicating whether packages within the application may be overwritten using the same version string.
	AllowUpdates pulumi.BoolPtrInput
	// The package to use if a client requests the application but does not specify a version. This property can only be set to the name of an existing package.
	DefaultVersion pulumi.StringPtrInput
	// The display name for the application.
	DisplayName pulumi.StringPtrInput
	// The name of the application. This must be unique within the account.
	Name pulumi.StringInput
	// The name of the resource group that contains the Batch account.
	ResourceGroupName pulumi.StringInput
}

func (ApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationArgs)(nil)).Elem()
}
