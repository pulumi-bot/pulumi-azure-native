// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * P2SVpnServerConfiguration Resource.
 */
export class P2sVpnServerConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing P2sVpnServerConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): P2sVpnServerConfiguration {
        return new P2sVpnServerConfiguration(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20181201:P2sVpnServerConfiguration';

    /**
     * Returns true if the given object is an instance of P2sVpnServerConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is P2sVpnServerConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === P2sVpnServerConfiguration.__pulumiType;
    }

    /**
     * Gets a unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * Parameters for P2SVpnServerConfiguration
     */
    public readonly properties!: pulumi.Output<outputs.network.v20181201.P2SVpnServerConfigurationPropertiesResponse>;

    /**
     * Create a P2sVpnServerConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: P2sVpnServerConfigurationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: P2sVpnServerConfigurationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as P2sVpnServerConfigurationArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.virtualWanName === undefined) {
                throw new Error("Missing required property 'virtualWanName'");
            }
            inputs["id"] = args ? args.id : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["virtualWanName"] = args ? args.virtualWanName : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(P2sVpnServerConfiguration.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a P2sVpnServerConfiguration resource.
 */
export interface P2sVpnServerConfigurationArgs {
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * The name of the P2SVpnServerConfiguration.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Parameters for P2SVpnServerConfiguration
     */
    readonly properties?: pulumi.Input<inputs.network.v20181201.P2SVpnServerConfigurationProperties>;
    /**
     * The resource group name of the VirtualWan.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the VirtualWan.
     */
    readonly virtualWanName: pulumi.Input<string>;
}
