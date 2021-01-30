// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20201015preview

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
	case "azure-nextgen:eventgrid/v20201015preview:Domain":
		r, err = NewDomain(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:eventgrid/v20201015preview:DomainTopic":
		r, err = NewDomainTopic(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:eventgrid/v20201015preview:EventChannel":
		r, err = NewEventChannel(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:eventgrid/v20201015preview:EventSubscription":
		r, err = NewEventSubscription(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:eventgrid/v20201015preview:PartnerNamespace":
		r, err = NewPartnerNamespace(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:eventgrid/v20201015preview:PartnerRegistration":
		r, err = NewPartnerRegistration(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:eventgrid/v20201015preview:PartnerTopicEventSubscription":
		r, err = NewPartnerTopicEventSubscription(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:eventgrid/v20201015preview:PrivateEndpointConnection":
		r, err = NewPrivateEndpointConnection(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:eventgrid/v20201015preview:SystemTopic":
		r, err = NewSystemTopic(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:eventgrid/v20201015preview:SystemTopicEventSubscription":
		r, err = NewSystemTopicEventSubscription(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:eventgrid/v20201015preview:Topic":
		r, err = NewTopic(ctx, name, nil, pulumi.URN_(urn))
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
		"eventgrid/v20201015preview",
		&module{version},
	)
}
