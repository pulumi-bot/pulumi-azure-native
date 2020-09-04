// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
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
    public static readonly __pulumiType = 'azurerm:network/v20181101:LoadBalancer';

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
     * Collection of backend address pools used by a load balancer
     */
    public readonly backendAddressPools!: pulumi.Output<outputs.network.v20181101.BackendAddressPoolResponse[] | undefined>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * Object representing the frontend IPs to be used for the load balancer
     */
    public readonly frontendIPConfigurations!: pulumi.Output<outputs.network.v20181101.FrontendIPConfigurationResponse[] | undefined>;
    /**
     * Defines an external port range for inbound NAT to a single backend port on NICs associated with a load balancer. Inbound NAT rules are created automatically for each NIC associated with the Load Balancer using an external port from this range. Defining an Inbound NAT pool on your Load Balancer is mutually exclusive with defining inbound Nat rules. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an inbound NAT pool. They have to reference individual inbound NAT rules.
     */
    public readonly inboundNatPools!: pulumi.Output<outputs.network.v20181101.InboundNatPoolResponse[] | undefined>;
    /**
     * Collection of inbound NAT Rules used by a load balancer. Defining inbound NAT rules on your load balancer is mutually exclusive with defining an inbound NAT pool. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an Inbound NAT pool. They have to reference individual inbound NAT rules.
     */
    public readonly inboundNatRules!: pulumi.Output<outputs.network.v20181101.InboundNatRuleResponse[] | undefined>;
    /**
     * Object collection representing the load balancing rules Gets the provisioning 
     */
    public readonly loadBalancingRules!: pulumi.Output<outputs.network.v20181101.LoadBalancingRuleResponse[] | undefined>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The outbound rules.
     */
    public readonly outboundRules!: pulumi.Output<outputs.network.v20181101.OutboundRuleResponse[] | undefined>;
    /**
     * Collection of probe objects used in the load balancer
     */
    public readonly probes!: pulumi.Output<outputs.network.v20181101.ProbeResponse[] | undefined>;
    /**
     * Gets the provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
     */
    public readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * The resource GUID property of the load balancer resource.
     */
    public readonly resourceGuid!: pulumi.Output<string | undefined>;
    /**
     * The load balancer SKU.
     */
    public readonly sku!: pulumi.Output<outputs.network.v20181101.LoadBalancerSkuResponse | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
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
            if (!args || args.loadBalancerName === undefined) {
                throw new Error("Missing required property 'loadBalancerName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["backendAddressPools"] = args ? args.backendAddressPools : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["frontendIPConfigurations"] = args ? args.frontendIPConfigurations : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["inboundNatPools"] = args ? args.inboundNatPools : undefined;
            inputs["inboundNatRules"] = args ? args.inboundNatRules : undefined;
            inputs["loadBalancerName"] = args ? args.loadBalancerName : undefined;
            inputs["loadBalancingRules"] = args ? args.loadBalancingRules : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["outboundRules"] = args ? args.outboundRules : undefined;
            inputs["probes"] = args ? args.probes : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceGuid"] = args ? args.resourceGuid : undefined;
            inputs["sku"] = args ? args.sku : undefined;
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
            inputs["outboundRules"] = undefined /*out*/;
            inputs["probes"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["resourceGuid"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/latest:LoadBalancer" }, { type: "azurerm:network/v20150501preview:LoadBalancer" }, { type: "azurerm:network/v20150615:LoadBalancer" }, { type: "azurerm:network/v20160330:LoadBalancer" }, { type: "azurerm:network/v20160601:LoadBalancer" }, { type: "azurerm:network/v20160901:LoadBalancer" }, { type: "azurerm:network/v20161201:LoadBalancer" }, { type: "azurerm:network/v20170301:LoadBalancer" }, { type: "azurerm:network/v20170601:LoadBalancer" }, { type: "azurerm:network/v20170801:LoadBalancer" }, { type: "azurerm:network/v20170901:LoadBalancer" }, { type: "azurerm:network/v20171001:LoadBalancer" }, { type: "azurerm:network/v20171101:LoadBalancer" }, { type: "azurerm:network/v20180101:LoadBalancer" }, { type: "azurerm:network/v20180201:LoadBalancer" }, { type: "azurerm:network/v20180401:LoadBalancer" }, { type: "azurerm:network/v20180601:LoadBalancer" }, { type: "azurerm:network/v20180701:LoadBalancer" }, { type: "azurerm:network/v20180801:LoadBalancer" }, { type: "azurerm:network/v20181001:LoadBalancer" }, { type: "azurerm:network/v20181201:LoadBalancer" }, { type: "azurerm:network/v20190201:LoadBalancer" }, { type: "azurerm:network/v20190401:LoadBalancer" }, { type: "azurerm:network/v20190601:LoadBalancer" }, { type: "azurerm:network/v20190701:LoadBalancer" }, { type: "azurerm:network/v20190801:LoadBalancer" }, { type: "azurerm:network/v20190901:LoadBalancer" }, { type: "azurerm:network/v20191101:LoadBalancer" }, { type: "azurerm:network/v20191201:LoadBalancer" }, { type: "azurerm:network/v20200301:LoadBalancer" }, { type: "azurerm:network/v20200401:LoadBalancer" }, { type: "azurerm:network/v20200501:LoadBalancer" }, { type: "azurerm:network/v20200601:LoadBalancer" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(LoadBalancer.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a LoadBalancer resource.
 */
export interface LoadBalancerArgs {
    /**
     * Collection of backend address pools used by a load balancer
     */
    readonly backendAddressPools?: pulumi.Input<pulumi.Input<inputs.network.v20181101.BackendAddressPool>[]>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Object representing the frontend IPs to be used for the load balancer
     */
    readonly frontendIPConfigurations?: pulumi.Input<pulumi.Input<inputs.network.v20181101.FrontendIPConfiguration>[]>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * Defines an external port range for inbound NAT to a single backend port on NICs associated with a load balancer. Inbound NAT rules are created automatically for each NIC associated with the Load Balancer using an external port from this range. Defining an Inbound NAT pool on your Load Balancer is mutually exclusive with defining inbound Nat rules. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an inbound NAT pool. They have to reference individual inbound NAT rules.
     */
    readonly inboundNatPools?: pulumi.Input<pulumi.Input<inputs.network.v20181101.InboundNatPool>[]>;
    /**
     * Collection of inbound NAT Rules used by a load balancer. Defining inbound NAT rules on your load balancer is mutually exclusive with defining an inbound NAT pool. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual virtual machines cannot reference an Inbound NAT pool. They have to reference individual inbound NAT rules.
     */
    readonly inboundNatRules?: pulumi.Input<pulumi.Input<inputs.network.v20181101.InboundNatRule>[]>;
    /**
     * The name of the load balancer.
     */
    readonly loadBalancerName: pulumi.Input<string>;
    /**
     * Object collection representing the load balancing rules Gets the provisioning 
     */
    readonly loadBalancingRules?: pulumi.Input<pulumi.Input<inputs.network.v20181101.LoadBalancingRule>[]>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The outbound rules.
     */
    readonly outboundRules?: pulumi.Input<pulumi.Input<inputs.network.v20181101.OutboundRule>[]>;
    /**
     * Collection of probe objects used in the load balancer
     */
    readonly probes?: pulumi.Input<pulumi.Input<inputs.network.v20181101.Probe>[]>;
    /**
     * Gets the provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
     */
    readonly provisioningState?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The resource GUID property of the load balancer resource.
     */
    readonly resourceGuid?: pulumi.Input<string>;
    /**
     * The load balancer SKU.
     */
    readonly sku?: pulumi.Input<inputs.network.v20181101.LoadBalancerSku>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
