// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storsimple

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Challenge-Handshake Authentication Protocol (CHAP) setting
// API Version: 2016-10-01.
func LookupChapSetting(ctx *pulumi.Context, args *LookupChapSettingArgs, opts ...pulumi.InvokeOption) (*LookupChapSettingResult, error) {
	var rv LookupChapSettingResult
	err := ctx.Invoke("azure-native:storsimple:getChapSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupChapSettingArgs struct {
	// The user name of chap to be fetched.
	ChapUserName string `pulumi:"chapUserName"`
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// The manager name
	ManagerName string `pulumi:"managerName"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Challenge-Handshake Authentication Protocol (CHAP) setting
type LookupChapSettingResult struct {
	// The identifier.
	Id string `pulumi:"id"`
	// The name.
	Name string `pulumi:"name"`
	// The chap password.
	Password AsymmetricEncryptedSecretResponse `pulumi:"password"`
	// The type.
	Type string `pulumi:"type"`
}
