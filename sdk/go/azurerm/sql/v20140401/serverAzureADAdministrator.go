// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20140401

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An server Active Directory Administrator.
type ServerAzureADAdministrator struct {
	pulumi.CustomResourceState

	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of the resource.
	Properties ServerAdministratorPropertiesResponseOutput `pulumi:"properties"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewServerAzureADAdministrator registers a new resource with the given unique name, arguments, and options.
func NewServerAzureADAdministrator(ctx *pulumi.Context,
	name string, args *ServerAzureADAdministratorArgs, opts ...pulumi.ResourceOption) (*ServerAzureADAdministrator, error) {
	if args == nil || args.AdministratorType == nil {
		return nil, errors.New("missing required argument 'AdministratorType'")
	}
	if args == nil || args.Login == nil {
		return nil, errors.New("missing required argument 'Login'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServerName == nil {
		return nil, errors.New("missing required argument 'ServerName'")
	}
	if args == nil || args.Sid == nil {
		return nil, errors.New("missing required argument 'Sid'")
	}
	if args == nil || args.TenantId == nil {
		return nil, errors.New("missing required argument 'TenantId'")
	}
	if args == nil {
		args = &ServerAzureADAdministratorArgs{}
	}
	var resource ServerAzureADAdministrator
	err := ctx.RegisterResource("azurerm:sql/v20140401:ServerAzureADAdministrator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerAzureADAdministrator gets an existing ServerAzureADAdministrator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerAzureADAdministrator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerAzureADAdministratorState, opts ...pulumi.ResourceOption) (*ServerAzureADAdministrator, error) {
	var resource ServerAzureADAdministrator
	err := ctx.ReadResource("azurerm:sql/v20140401:ServerAzureADAdministrator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerAzureADAdministrator resources.
type serverAzureADAdministratorState struct {
	// Resource name.
	Name *string `pulumi:"name"`
	// The properties of the resource.
	Properties *ServerAdministratorPropertiesResponse `pulumi:"properties"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type ServerAzureADAdministratorState struct {
	// Resource name.
	Name pulumi.StringPtrInput
	// The properties of the resource.
	Properties ServerAdministratorPropertiesResponsePtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (ServerAzureADAdministratorState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverAzureADAdministratorState)(nil)).Elem()
}

type serverAzureADAdministratorArgs struct {
	// The type of administrator.
	AdministratorType string `pulumi:"administratorType"`
	// The server administrator login value.
	Login string `pulumi:"login"`
	// Name of the server administrator resource.
	Name string `pulumi:"name"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The server administrator Sid (Secure ID).
	Sid string `pulumi:"sid"`
	// The server Active Directory Administrator tenant id.
	TenantId string `pulumi:"tenantId"`
}

// The set of arguments for constructing a ServerAzureADAdministrator resource.
type ServerAzureADAdministratorArgs struct {
	// The type of administrator.
	AdministratorType pulumi.StringInput
	// The server administrator login value.
	Login pulumi.StringInput
	// Name of the server administrator resource.
	Name pulumi.StringInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
	// The server administrator Sid (Secure ID).
	Sid pulumi.StringInput
	// The server Active Directory Administrator tenant id.
	TenantId pulumi.StringInput
}

func (ServerAzureADAdministratorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverAzureADAdministratorArgs)(nil)).Elem()
}
