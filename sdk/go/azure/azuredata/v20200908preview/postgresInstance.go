// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200908preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A Postgres Instance.
type PostgresInstance struct {
	pulumi.CustomResourceState

	// The instance admin
	Admin pulumi.StringPtrOutput `pulumi:"admin"`
	// The data controller id
	DataControllerId pulumi.StringPtrOutput `pulumi:"dataControllerId"`
	// The raw kubernetes information
	K8sRaw pulumi.MapOutput `pulumi:"k8sRaw"`
	// Last uploaded date from on premise cluster. Defaults to current date time
	LastUploadedDate pulumi.StringPtrOutput `pulumi:"lastUploadedDate"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Read only system data
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPostgresInstance registers a new resource with the given unique name, arguments, and options.
func NewPostgresInstance(ctx *pulumi.Context,
	name string, args *PostgresInstanceArgs, opts ...pulumi.ResourceOption) (*PostgresInstance, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.PostgresInstanceName == nil {
		return nil, errors.New("missing required argument 'PostgresInstanceName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &PostgresInstanceArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:azuredata/v20190724preview:PostgresInstance"),
		},
	})
	opts = append(opts, aliases)
	var resource PostgresInstance
	err := ctx.RegisterResource("azure-nextgen:azuredata/v20200908preview:PostgresInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPostgresInstance gets an existing PostgresInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPostgresInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PostgresInstanceState, opts ...pulumi.ResourceOption) (*PostgresInstance, error) {
	var resource PostgresInstance
	err := ctx.ReadResource("azure-nextgen:azuredata/v20200908preview:PostgresInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PostgresInstance resources.
type postgresInstanceState struct {
	// The instance admin
	Admin *string `pulumi:"admin"`
	// The data controller id
	DataControllerId *string `pulumi:"dataControllerId"`
	// The raw kubernetes information
	K8sRaw map[string]interface{} `pulumi:"k8sRaw"`
	// Last uploaded date from on premise cluster. Defaults to current date time
	LastUploadedDate *string `pulumi:"lastUploadedDate"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// Read only system data
	SystemData *SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
}

type PostgresInstanceState struct {
	// The instance admin
	Admin pulumi.StringPtrInput
	// The data controller id
	DataControllerId pulumi.StringPtrInput
	// The raw kubernetes information
	K8sRaw pulumi.MapInput
	// Last uploaded date from on premise cluster. Defaults to current date time
	LastUploadedDate pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// Read only system data
	SystemData SystemDataResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
}

func (PostgresInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*postgresInstanceState)(nil)).Elem()
}

type postgresInstanceArgs struct {
	// The instance admin
	Admin *string `pulumi:"admin"`
	// The data controller id
	DataControllerId *string `pulumi:"dataControllerId"`
	// The raw kubernetes information
	K8sRaw map[string]interface{} `pulumi:"k8sRaw"`
	// Last uploaded date from on premise cluster. Defaults to current date time
	LastUploadedDate *string `pulumi:"lastUploadedDate"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Name of PostgresInstance
	PostgresInstanceName string `pulumi:"postgresInstanceName"`
	// The name of the Azure resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a PostgresInstance resource.
type PostgresInstanceArgs struct {
	// The instance admin
	Admin pulumi.StringPtrInput
	// The data controller id
	DataControllerId pulumi.StringPtrInput
	// The raw kubernetes information
	K8sRaw pulumi.MapInput
	// Last uploaded date from on premise cluster. Defaults to current date time
	LastUploadedDate pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringInput
	// Name of PostgresInstance
	PostgresInstanceName pulumi.StringInput
	// The name of the Azure resource group
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (PostgresInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*postgresInstanceArgs)(nil)).Elem()
}
