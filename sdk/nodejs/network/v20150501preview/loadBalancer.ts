// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * LoadBalancer resource
 */
export class LoadBalancer extends pulumi.CustomResource {
    /**
     * Get an existing LoadBalancer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): LoadBalancer {
        return new LoadBalancer(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:network/v20150501preview:LoadBalancer';

    /**
     * Returns true if the given object is an instance of LoadBalancer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LoadBalancer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LoadBalancer.__pulumiType;
    }

    /**
     * Gets or sets Pools of backend IP addresses
     */
    public readonly backendAddressPools!: pulumi.Output<outputs.network.v20150501preview.BackendAddressPoolResponse[] | undefined>;
    /**
     * Gets a unique read-only string that changes whenever the resource is updated
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * Gets or sets frontend IP addresses of the load balancer
     */
    public readonly frontendIPConfigurations!: pulumi.Output<outputs.network.v20150501preview.FrontendIpConfigurationResponse[] | undefined>;
    /**
     * Gets or sets inbound NAT pools
     */
    public readonly inboundNatPools!: pulumi.Output<outputs.network.v20150501preview.InboundNatPoolResponse[] | undefined>;
    /**
     * Gets or sets list of inbound rules
     */
    public readonly inboundNatRules!: pulumi.Output<outputs.network.v20150501preview.InboundNatRuleResponse[] | undefined>;
    /**
     * Gets or sets load balancing rules
     */
    public readonly loadBalancingRules!: pulumi.Output<outputs.network.v20150501preview.LoadBalancingRuleResponse[] | undefined>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Gets or sets outbound NAT rules
     */
    public readonly outboundNatRules!: pulumi.Output<outputs.network.v20150501preview.OutboundNatRuleResponse[] | undefined>;
    /**
     * Gets or sets list of Load balancer probes
     */
    public readonly probes!: pulumi.Output<outputs.network.v20150501preview.ProbeResponse[] | undefined>;
    /**
     * Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
     */
    public readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * Gets or sets resource guid property of the Load balancer resource
     */
    public readonly resourceGuid!: pulumi.Output<string | undefined>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a LoadBalancer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LoadBalancerArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.loadBalancerName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'loadBalancerName'");
            }
            if ((!args || args.location === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["backendAddressPools"] = args ? args.backendAddressPools : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["frontendIPConfigurations"] = args ? args.frontendIPConfigurations : undefined;
            inputs["inboundNatPools"] = args ? args.inboundNatPools : undefined;
            inputs["inboundNatRules"] = args ? args.inboundNatRules : undefined;
            inputs["loadBalancerName"] = args ? args.loadBalancerName : undefined;
            inputs["loadBalancingRules"] = args ? args.loadBalancingRules : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["outboundNatRules"] = args ? args.outboundNatRules : undefined;
            inputs["probes"] = args ? args.probes : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceGuid"] = args ? args.resourceGuid : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["backendAddressPools"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["frontendIPConfigurations"] = undefined /*out*/;
            inputs["inboundNatPools"] = undefined /*out*/;
            inputs["inboundNatRules"] = undefined /*out*/;
            inputs["loadBalancingRules"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["outboundNatRules"] = undefined /*out*/;
            inputs["probes"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["resourceGuid"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:network/latest:LoadBalancer" }, { type: "azure-nextgen:network/v20150615:LoadBalancer" }, { type: "azure-nextgen:network/v20160330:LoadBalancer" }, { type: "azure-nextgen:network/v20160601:LoadBalancer" }, { type: "azure-nextgen:network/v20160901:LoadBalancer" }, { type: "azure-nextgen:network/v20161201:LoadBalancer" }, { type: "azure-nextgen:network/v20170301:LoadBalancer" }, { type: "azure-nextgen:network/v20170601:LoadBalancer" }, { type: "azure-nextgen:network/v20170801:LoadBalancer" }, { type: "azure-nextgen:network/v20170901:LoadBalancer" }, { type: "azure-nextgen:network/v20171001:LoadBalancer" }, { type: "azure-nextgen:network/v20171101:LoadBalancer" }, { type: "azure-nextgen:network/v20180101:LoadBalancer" }, { type: "azure-nextgen:network/v20180201:LoadBalancer" }, { type: "azure-nextgen:network/v20180401:LoadBalancer" }, { type: "azure-nextgen:network/v20180601:LoadBalancer" }, { type: "azure-nextgen:network/v20180701:LoadBalancer" }, { type: "azure-nextgen:network/v20180801:LoadBalancer" }, { type: "azure-nextgen:network/v20181001:LoadBalancer" }, { type: "azure-nextgen:network/v20181101:LoadBalancer" }, { type: "azure-nextgen:network/v20181201:LoadBalancer" }, { type: "azure-nextgen:network/v20190201:LoadBalancer" }, { type: "azure-nextgen:network/v20190401:LoadBalancer" }, { type: "azure-nextgen:network/v20190601:LoadBalancer" }, { type: "azure-nextgen:network/v20190701:LoadBalancer" }, { type: "azure-nextgen:network/v20190801:LoadBalancer" }, { type: "azure-nextgen:network/v20190901:LoadBalancer" }, { type: "azure-nextgen:network/v20191101:LoadBalancer" }, { type: "azure-nextgen:network/v20191201:LoadBalancer" }, { type: "azure-nextgen:network/v20200301:LoadBalancer" }, { type: "azure-nextgen:network/v20200401:LoadBalancer" }, { type: "azure-nextgen:network/v20200501:LoadBalancer" }, { type: "azure-nextgen:network/v20200601:LoadBalancer" }, { type: "azure-nextgen:network/v20200701:LoadBalancer" }, { type: "azure-nextgen:network/v20200801:LoadBalancer" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(LoadBalancer.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a LoadBalancer resource.
 */
export interface LoadBalancerArgs {
    /**
     * Gets or sets Pools of backend IP addresses
     */
    readonly backendAddressPools?: pulumi.Input<pulumi.Input<inputs.network.v20150501preview.BackendAddressPool>[]>;
    /**
     * Gets a unique read-only string that changes whenever the resource is updated
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Gets or sets frontend IP addresses of the load balancer
     */
    readonly frontendIPConfigurations?: pulumi.Input<pulumi.Input<inputs.network.v20150501preview.FrontendIpConfiguration>[]>;
    /**
     * Gets or sets inbound NAT pools
     */
    readonly inboundNatPools?: pulumi.Input<pulumi.Input<inputs.network.v20150501preview.InboundNatPool>[]>;
    /**
     * Gets or sets list of inbound rules
     */
    readonly inboundNatRules?: pulumi.Input<pulumi.Input<inputs.network.v20150501preview.InboundNatRule>[]>;
    /**
     * The name of the loadBalancer.
     */
    readonly loadBalancerName: pulumi.Input<string>;
    /**
     * Gets or sets load balancing rules
     */
    readonly loadBalancingRules?: pulumi.Input<pulumi.Input<inputs.network.v20150501preview.LoadBalancingRule>[]>;
    /**
     * Resource location
     */
    readonly location: pulumi.Input<string>;
    /**
     * Gets or sets outbound NAT rules
     */
    readonly outboundNatRules?: pulumi.Input<pulumi.Input<inputs.network.v20150501preview.OutboundNatRule>[]>;
    /**
     * Gets or sets list of Load balancer probes
     */
    readonly probes?: pulumi.Input<pulumi.Input<inputs.network.v20150501preview.Probe>[]>;
    /**
     * Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
     */
    readonly provisioningState?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Gets or sets resource guid property of the Load balancer resource
     */
    readonly resourceGuid?: pulumi.Input<string>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
