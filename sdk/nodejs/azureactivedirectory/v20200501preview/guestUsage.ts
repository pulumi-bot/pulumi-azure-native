// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Guest Usages Resource
 */
export class GuestUsage extends pulumi.CustomResource {
    /**
     * Get an existing GuestUsage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): GuestUsage {
        return new GuestUsage(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:azureactivedirectory/v20200501preview:GuestUsage';

    /**
     * Returns true if the given object is an instance of GuestUsage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GuestUsage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GuestUsage.__pulumiType;
    }

    /**
     * Location of the Guest Usages resource.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The name of the Guest Usages resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Key-value pairs of additional resource provisioning properties.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * An identifier for the tenant for which the resource is being created
     */
    public readonly tenantId!: pulumi.Output<string | undefined>;
    /**
     * The type of the Guest Usages resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a GuestUsage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GuestUsageArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.resourceName === undefined) {
                throw new Error("Missing required property 'resourceName'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceName"] = args ? args.resourceName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tenantId"] = args ? args.tenantId : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["tenantId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(GuestUsage.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a GuestUsage resource.
 */
export interface GuestUsageArgs {
    /**
     * Location of the Guest Usages resource.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The initial domain name of the AAD tenant.
     */
    readonly resourceName: pulumi.Input<string>;
    /**
     * Key-value pairs of additional resource provisioning properties.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * An identifier for the tenant for which the resource is being created
     */
    readonly tenantId?: pulumi.Input<string>;
}
