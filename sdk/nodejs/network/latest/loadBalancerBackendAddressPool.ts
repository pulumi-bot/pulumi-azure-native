// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Pool of backend IP addresses.
 * Latest API Version: 2020-08-01.
 */
export class LoadBalancerBackendAddressPool extends pulumi.CustomResource {
    /**
     * Get an existing LoadBalancerBackendAddressPool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): LoadBalancerBackendAddressPool {
        return new LoadBalancerBackendAddressPool(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:network/latest:LoadBalancerBackendAddressPool';

    /**
     * Returns true if the given object is an instance of LoadBalancerBackendAddressPool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LoadBalancerBackendAddressPool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LoadBalancerBackendAddressPool.__pulumiType;
    }

    /**
     * An array of references to IP addresses defined in network interfaces.
     */
    public /*out*/ readonly backendIPConfigurations!: pulumi.Output<outputs.network.latest.NetworkInterfaceIPConfigurationResponse[]>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * An array of backend addresses.
     */
    public readonly loadBalancerBackendAddresses!: pulumi.Output<outputs.network.latest.LoadBalancerBackendAddressResponse[] | undefined>;
    /**
     * An array of references to load balancing rules that use this backend address pool.
     */
    public /*out*/ readonly loadBalancingRules!: pulumi.Output<outputs.network.latest.SubResourceResponse[]>;
    /**
     * The location of the backend address pool.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource that is unique within the set of backend address pools used by the load balancer. This name can be used to access the resource.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * A reference to an outbound rule that uses this backend address pool.
     */
    public /*out*/ readonly outboundRule!: pulumi.Output<outputs.network.latest.SubResourceResponse>;
    /**
     * An array of references to outbound rules that use this backend address pool.
     */
    public /*out*/ readonly outboundRules!: pulumi.Output<outputs.network.latest.SubResourceResponse[]>;
    /**
     * The provisioning state of the backend address pool resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a LoadBalancerBackendAddressPool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LoadBalancerBackendAddressPoolArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.backendAddressPoolName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'backendAddressPoolName'");
            }
            if ((!args || args.loadBalancerName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'loadBalancerName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["backendAddressPoolName"] = args ? args.backendAddressPoolName : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["loadBalancerBackendAddresses"] = args ? args.loadBalancerBackendAddresses : undefined;
            inputs["loadBalancerName"] = args ? args.loadBalancerName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["backendIPConfigurations"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["loadBalancingRules"] = undefined /*out*/;
            inputs["outboundRule"] = undefined /*out*/;
            inputs["outboundRules"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["backendIPConfigurations"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["loadBalancerBackendAddresses"] = undefined /*out*/;
            inputs["loadBalancingRules"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["outboundRule"] = undefined /*out*/;
            inputs["outboundRules"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:network/v20200401:LoadBalancerBackendAddressPool" }, { type: "azure-nextgen:network/v20200501:LoadBalancerBackendAddressPool" }, { type: "azure-nextgen:network/v20200601:LoadBalancerBackendAddressPool" }, { type: "azure-nextgen:network/v20200701:LoadBalancerBackendAddressPool" }, { type: "azure-nextgen:network/v20200801:LoadBalancerBackendAddressPool" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(LoadBalancerBackendAddressPool.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a LoadBalancerBackendAddressPool resource.
 */
export interface LoadBalancerBackendAddressPoolArgs {
    /**
     * The name of the backend address pool.
     */
    readonly backendAddressPoolName: pulumi.Input<string>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * An array of backend addresses.
     */
    readonly loadBalancerBackendAddresses?: pulumi.Input<pulumi.Input<inputs.network.latest.LoadBalancerBackendAddress>[]>;
    /**
     * The name of the load balancer.
     */
    readonly loadBalancerName: pulumi.Input<string>;
    /**
     * The location of the backend address pool.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the resource that is unique within the set of backend address pools used by the load balancer. This name can be used to access the resource.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
