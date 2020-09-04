// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The Managed Network Group resource
 */
export class ManagedNetworkGroup extends pulumi.CustomResource {
    /**
     * Get an existing ManagedNetworkGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ManagedNetworkGroup {
        return new ManagedNetworkGroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:managednetwork/v20190601preview:ManagedNetworkGroup';

    /**
     * Returns true if the given object is an instance of ManagedNetworkGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ManagedNetworkGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ManagedNetworkGroup.__pulumiType;
    }

    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Responsibility role under which this Managed Network Group will be created
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * The geo-location where the resource lives
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The collection of management groups covered by the Managed Network
     */
    public readonly managementGroups!: pulumi.Output<outputs.managednetwork.v20190601preview.ResourceIdResponse[] | undefined>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Provisioning state of the ManagedNetwork resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The collection of  subnets covered by the Managed Network
     */
    public readonly subnets!: pulumi.Output<outputs.managednetwork.v20190601preview.ResourceIdResponse[] | undefined>;
    /**
     * The collection of subscriptions covered by the Managed Network
     */
    public readonly subscriptions!: pulumi.Output<outputs.managednetwork.v20190601preview.ResourceIdResponse[] | undefined>;
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The collection of virtual nets covered by the Managed Network
     */
    public readonly virtualNetworks!: pulumi.Output<outputs.managednetwork.v20190601preview.ResourceIdResponse[] | undefined>;

    /**
     * Create a ManagedNetworkGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ManagedNetworkGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.managedNetworkGroupName === undefined) {
                throw new Error("Missing required property 'managedNetworkGroupName'");
            }
            if (!args || args.managedNetworkName === undefined) {
                throw new Error("Missing required property 'managedNetworkName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["kind"] = args ? args.kind : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["managedNetworkGroupName"] = args ? args.managedNetworkGroupName : undefined;
            inputs["managedNetworkName"] = args ? args.managedNetworkName : undefined;
            inputs["managementGroups"] = args ? args.managementGroups : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["subnets"] = args ? args.subnets : undefined;
            inputs["subscriptions"] = args ? args.subscriptions : undefined;
            inputs["virtualNetworks"] = args ? args.virtualNetworks : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["etag"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["managementGroups"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["subnets"] = undefined /*out*/;
            inputs["subscriptions"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["virtualNetworks"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ManagedNetworkGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ManagedNetworkGroup resource.
 */
export interface ManagedNetworkGroupArgs {
    /**
     * Responsibility role under which this Managed Network Group will be created
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * The geo-location where the resource lives
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the Managed Network Group.
     */
    readonly managedNetworkGroupName: pulumi.Input<string>;
    /**
     * The name of the Managed Network.
     */
    readonly managedNetworkName: pulumi.Input<string>;
    /**
     * The collection of management groups covered by the Managed Network
     */
    readonly managementGroups?: pulumi.Input<pulumi.Input<inputs.managednetwork.v20190601preview.ResourceId>[]>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The collection of  subnets covered by the Managed Network
     */
    readonly subnets?: pulumi.Input<pulumi.Input<inputs.managednetwork.v20190601preview.ResourceId>[]>;
    /**
     * The collection of subscriptions covered by the Managed Network
     */
    readonly subscriptions?: pulumi.Input<pulumi.Input<inputs.managednetwork.v20190601preview.ResourceId>[]>;
    /**
     * The collection of virtual nets covered by the Managed Network
     */
    readonly virtualNetworks?: pulumi.Input<pulumi.Input<inputs.managednetwork.v20190601preview.ResourceId>[]>;
}
