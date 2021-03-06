// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20151031

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure-native/sdk/go/azure"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "azure-native:automation/v20151031:AutomationAccount":
		r, err = NewAutomationAccount(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:automation/v20151031:Certificate":
		r, err = NewCertificate(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:automation/v20151031:Connection":
		r, err = NewConnection(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:automation/v20151031:ConnectionType":
		r, err = NewConnectionType(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:automation/v20151031:Credential":
		r, err = NewCredential(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:automation/v20151031:DscConfiguration":
		r, err = NewDscConfiguration(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:automation/v20151031:DscNodeConfiguration":
		r, err = NewDscNodeConfiguration(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:automation/v20151031:JobSchedule":
		r, err = NewJobSchedule(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:automation/v20151031:Module":
		r, err = NewModule(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:automation/v20151031:Runbook":
		r, err = NewRunbook(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:automation/v20151031:Schedule":
		r, err = NewSchedule(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:automation/v20151031:Variable":
		r, err = NewVariable(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:automation/v20151031:Watcher":
		r, err = NewWatcher(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:automation/v20151031:Webhook":
		r, err = NewWebhook(ctx, name, nil, pulumi.URN_(urn))
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	return
}

func init() {
	version, err := azure.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"azure-native",
		"automation/v20151031",
		&module{version},
	)
}
