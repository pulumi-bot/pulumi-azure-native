// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Single item in List or Get Alias(Disaster Recovery configuration) operation
 */
export class DisasterRecoveryConfig extends pulumi.CustomResource {
    /**
     * Get an existing DisasterRecoveryConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DisasterRecoveryConfig {
        return new DisasterRecoveryConfig(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:eventhub/v20170401:DisasterRecoveryConfig';

    /**
     * Returns true if the given object is an instance of DisasterRecoveryConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DisasterRecoveryConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DisasterRecoveryConfig.__pulumiType;
    }

    /**
     * Alternate name specified when alias and namespace names are same.
     */
    public readonly alternateName!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * ARM Id of the Primary/Secondary eventhub namespace name, which is part of GEO DR pairing
     */
    public readonly partnerNamespace!: pulumi.Output<string | undefined>;
    /**
     * Number of entities pending to be replicated.
     */
    public /*out*/ readonly pendingReplicationOperationsCount!: pulumi.Output<number>;
    /**
     * Provisioning state of the Alias(Disaster Recovery configuration) - possible values 'Accepted' or 'Succeeded' or 'Failed'
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<ProvisioningStateDR>;
    /**
     * role of namespace in GEO DR - possible values 'Primary' or 'PrimaryNotReplicating' or 'Secondary'
     */
    public /*out*/ readonly role!: pulumi.Output<RoleDisasterRecovery>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a DisasterRecoveryConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DisasterRecoveryConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DisasterRecoveryConfigArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as DisasterRecoveryConfigArgs | undefined;
            if (!args || args.alias === undefined) {
                throw new Error("Missing required property 'alias'");
            }
            if (!args || args.namespaceName === undefined) {
                throw new Error("Missing required property 'namespaceName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["alias"] = args ? args.alias : undefined;
            inputs["alternateName"] = args ? args.alternateName : undefined;
            inputs["namespaceName"] = args ? args.namespaceName : undefined;
            inputs["partnerNamespace"] = args ? args.partnerNamespace : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["pendingReplicationOperationsCount"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["role"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DisasterRecoveryConfig.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DisasterRecoveryConfig resource.
 */
export interface DisasterRecoveryConfigArgs {
    /**
     * The Disaster Recovery configuration name
     */
    readonly alias: pulumi.Input<string>;
    /**
     * Alternate name specified when alias and namespace names are same.
     */
    readonly alternateName?: pulumi.Input<string>;
    /**
     * The Namespace name
     */
    readonly namespaceName: pulumi.Input<string>;
    /**
     * ARM Id of the Primary/Secondary eventhub namespace name, which is part of GEO DR pairing
     */
    readonly partnerNamespace?: pulumi.Input<string>;
    /**
     * Name of the resource group within the azure subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
