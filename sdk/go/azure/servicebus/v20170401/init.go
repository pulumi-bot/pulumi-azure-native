// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170401

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure-nextgen/sdk/go/azure"
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
	case "azure-nextgen:servicebus/v20170401:DisasterRecoveryConfig":
		r, err = NewDisasterRecoveryConfig(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:servicebus/v20170401:MigrationConfig":
		r, err = NewMigrationConfig(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:servicebus/v20170401:Namespace":
		r, err = NewNamespace(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:servicebus/v20170401:NamespaceAuthorizationRule":
		r, err = NewNamespaceAuthorizationRule(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:servicebus/v20170401:NamespaceNetworkRuleSet":
		r, err = NewNamespaceNetworkRuleSet(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:servicebus/v20170401:Queue":
		r, err = NewQueue(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:servicebus/v20170401:QueueAuthorizationRule":
		r, err = NewQueueAuthorizationRule(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:servicebus/v20170401:Rule":
		r, err = NewRule(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:servicebus/v20170401:Subscription":
		r, err = NewSubscription(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:servicebus/v20170401:Topic":
		r, err = NewTopic(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:servicebus/v20170401:TopicAuthorizationRule":
		r, err = NewTopicAuthorizationRule(ctx, name, nil, pulumi.URN_(urn))
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
		"azure-nextgen",
		"servicebus/v20170401",
		&module{version},
	)
}
