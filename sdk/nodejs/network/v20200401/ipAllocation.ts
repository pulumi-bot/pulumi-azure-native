// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * IpAllocation resource.
 */
export class IpAllocation extends pulumi.CustomResource {
    /**
     * Get an existing IpAllocation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): IpAllocation {
        return new IpAllocation(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20200401:IpAllocation';

    /**
     * Returns true if the given object is an instance of IpAllocation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IpAllocation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IpAllocation.__pulumiType;
    }

    /**
     * IpAllocation tags.
     */
    public readonly allocationTags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The IPAM allocation ID.
     */
    public readonly ipamAllocationId!: pulumi.Output<string | undefined>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The address prefix for the IpAllocation.
     */
    public readonly prefix!: pulumi.Output<string | undefined>;
    /**
     * The address prefix length for the IpAllocation.
     */
    public readonly prefixLength!: pulumi.Output<number | undefined>;
    /**
     * The address prefix Type for the IpAllocation.
     */
    public readonly prefixType!: pulumi.Output<string | undefined>;
    /**
     * The Subnet that using the prefix of this IpAllocation resource.
     */
    public /*out*/ readonly subnet!: pulumi.Output<outputs.network.v20200401.SubResourceResponse>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The VirtualNetwork that using the prefix of this IpAllocation resource.
     */
    public /*out*/ readonly virtualNetwork!: pulumi.Output<outputs.network.v20200401.SubResourceResponse>;

    /**
     * Create a IpAllocation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IpAllocationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.ipAllocationName === undefined) {
                throw new Error("Missing required property 'ipAllocationName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["allocationTags"] = args ? args.allocationTags : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["ipAllocationName"] = args ? args.ipAllocationName : undefined;
            inputs["ipamAllocationId"] = args ? args.ipamAllocationId : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["prefix"] = args ? args.prefix : undefined;
            inputs["prefixLength"] = args ? args.prefixLength : undefined;
            inputs["prefixType"] = args ? args.prefixType : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["subnet"] = undefined /*out*/;
            inputs["virtualNetwork"] = undefined /*out*/;
        } else {
            inputs["allocationTags"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["ipamAllocationId"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["prefix"] = undefined /*out*/;
            inputs["prefixLength"] = undefined /*out*/;
            inputs["prefixType"] = undefined /*out*/;
            inputs["subnet"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["virtualNetwork"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/latest:IpAllocation" }, { type: "azurerm:network/v20200301:IpAllocation" }, { type: "azurerm:network/v20200501:IpAllocation" }, { type: "azurerm:network/v20200601:IpAllocation" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(IpAllocation.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a IpAllocation resource.
 */
export interface IpAllocationArgs {
    /**
     * IpAllocation tags.
     */
    readonly allocationTags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * The name of the IpAllocation.
     */
    readonly ipAllocationName: pulumi.Input<string>;
    /**
     * The IPAM allocation ID.
     */
    readonly ipamAllocationId?: pulumi.Input<string>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The address prefix for the IpAllocation.
     */
    readonly prefix?: pulumi.Input<string>;
    /**
     * The address prefix length for the IpAllocation.
     */
    readonly prefixLength?: pulumi.Input<number>;
    /**
     * The address prefix Type for the IpAllocation.
     */
    readonly prefixType?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type for the IpAllocation.
     */
    readonly type?: pulumi.Input<string>;
}
