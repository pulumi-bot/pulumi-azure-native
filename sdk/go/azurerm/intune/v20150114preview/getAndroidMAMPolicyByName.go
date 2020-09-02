// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150114preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupAndroidMAMPolicyByName(ctx *pulumi.Context, args *LookupAndroidMAMPolicyByNameArgs, opts ...pulumi.InvokeOption) (*LookupAndroidMAMPolicyByNameResult, error) {
	var rv LookupAndroidMAMPolicyByNameResult
	err := ctx.Invoke("azurerm:intune/v20150114preview:getAndroidMAMPolicyByName", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAndroidMAMPolicyByNameArgs struct {
	// Location hostName for the tenant
	HostName string `pulumi:"hostName"`
	// Unique name for the policy
	PolicyName string `pulumi:"policyName"`
	// select specific fields in entity.
	Select *string `pulumi:"select"`
}

// Android Policy entity for Intune MAM.
type LookupAndroidMAMPolicyByNameResult struct {
	AccessRecheckOfflineTimeout *string `pulumi:"accessRecheckOfflineTimeout"`
	AccessRecheckOnlineTimeout  *string `pulumi:"accessRecheckOnlineTimeout"`
	AppSharingFromLevel         *string `pulumi:"appSharingFromLevel"`
	AppSharingToLevel           *string `pulumi:"appSharingToLevel"`
	Authentication              *string `pulumi:"authentication"`
	ClipboardSharingLevel       *string `pulumi:"clipboardSharingLevel"`
	DataBackup                  *string `pulumi:"dataBackup"`
	Description                 *string `pulumi:"description"`
	DeviceCompliance            *string `pulumi:"deviceCompliance"`
	FileEncryption              *string `pulumi:"fileEncryption"`
	FileSharingSaveAs           *string `pulumi:"fileSharingSaveAs"`
	FriendlyName                string  `pulumi:"friendlyName"`
	GroupStatus                 string  `pulumi:"groupStatus"`
	LastModifiedTime            string  `pulumi:"lastModifiedTime"`
	// Resource Location
	Location       *string `pulumi:"location"`
	ManagedBrowser *string `pulumi:"managedBrowser"`
	// Resource name
	Name               string  `pulumi:"name"`
	NumOfApps          int     `pulumi:"numOfApps"`
	OfflineWipeTimeout *string `pulumi:"offlineWipeTimeout"`
	Pin                *string `pulumi:"pin"`
	PinNumRetry        *int    `pulumi:"pinNumRetry"`
	ScreenCapture      *string `pulumi:"screenCapture"`
	// Resource Tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}
