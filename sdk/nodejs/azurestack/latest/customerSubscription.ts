// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Customer subscription.
 *
 * ## Example Usage
 * ### Creates a new customer subscription under a registration.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const customerSubscription = new azurerm.azurestack.latest.CustomerSubscription("customerSubscription", {
 *     customerSubscriptionName: "E09A4E93-29A7-4EBA-A6D4-76202383F07F",
 *     registrationName: "testregistration",
 *     resourceGroup: "azurestack",
 *     tenantId: "dbab3982-796f-4d03-9908-044c08aef8a2",
 * });
 *
 * ```
 */
export class CustomerSubscription extends pulumi.CustomResource {
    /**
     * Get an existing CustomerSubscription resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CustomerSubscription {
        return new CustomerSubscription(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:azurestack/latest:CustomerSubscription';

    /**
     * Returns true if the given object is an instance of CustomerSubscription.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomerSubscription {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomerSubscription.__pulumiType;
    }

    /**
     * The entity tag used for optimistic concurrency when modifying the resource.
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * Name of the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Tenant Id.
     */
    public readonly tenantId!: pulumi.Output<string | undefined>;
    /**
     * Type of Resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a CustomerSubscription resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomerSubscriptionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CustomerSubscriptionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as CustomerSubscriptionArgs | undefined;
            if (!args || args.customerSubscriptionName === undefined) {
                throw new Error("Missing required property 'customerSubscriptionName'");
            }
            if (!args || args.registrationName === undefined) {
                throw new Error("Missing required property 'registrationName'");
            }
            if (!args || args.resourceGroup === undefined) {
                throw new Error("Missing required property 'resourceGroup'");
            }
            inputs["customerSubscriptionName"] = args ? args.customerSubscriptionName : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["registrationName"] = args ? args.registrationName : undefined;
            inputs["resourceGroup"] = args ? args.resourceGroup : undefined;
            inputs["tenantId"] = args ? args.tenantId : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:azurestack/v20170601:CustomerSubscription" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(CustomerSubscription.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a CustomerSubscription resource.
 */
export interface CustomerSubscriptionArgs {
    /**
     * Name of the product.
     */
    readonly customerSubscriptionName: pulumi.Input<string>;
    /**
     * The entity tag used for optimistic concurrency when modifying the resource.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Name of the Azure Stack registration.
     */
    readonly registrationName: pulumi.Input<string>;
    /**
     * Name of the resource group.
     */
    readonly resourceGroup: pulumi.Input<string>;
    /**
     * Tenant Id.
     */
    readonly tenantId?: pulumi.Input<string>;
}
