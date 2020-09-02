// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170301preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupManagedInstanceAdministrator(ctx *pulumi.Context, args *LookupManagedInstanceAdministratorArgs, opts ...pulumi.InvokeOption) (*LookupManagedInstanceAdministratorResult, error) {
	var rv LookupManagedInstanceAdministratorResult
	err := ctx.Invoke("azurerm:sql/v20170301preview:getManagedInstanceAdministrator", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedInstanceAdministratorArgs struct {
	// The administrator name.
	AdministratorName string `pulumi:"administratorName"`
	// The name of the managed instance.
	ManagedInstanceName string `pulumi:"managedInstanceName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An Azure SQL managed instance administrator.
type LookupManagedInstanceAdministratorResult struct {
	// Type of the managed instance administrator.
	AdministratorType string `pulumi:"administratorType"`
	// Login name of the managed instance administrator.
	Login string `pulumi:"login"`
	// Resource name.
	Name string `pulumi:"name"`
	// SID (object ID) of the managed instance administrator.
	Sid string `pulumi:"sid"`
	// Tenant ID of the managed instance administrator.
	TenantId *string `pulumi:"tenantId"`
	// Resource type.
	Type string `pulumi:"type"`
}
