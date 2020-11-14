// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * An instance failover group.
 */
export class InstanceFailoverGroup extends pulumi.CustomResource {
    /**
     * Get an existing InstanceFailoverGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): InstanceFailoverGroup {
        return new InstanceFailoverGroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:sql/v20171001preview:InstanceFailoverGroup';

    /**
     * Returns true if the given object is an instance of InstanceFailoverGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceFailoverGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceFailoverGroup.__pulumiType;
    }

    /**
     * List of managed instance pairs in the failover group.
     */
    public readonly managedInstancePairs!: pulumi.Output<outputs.sql.v20171001preview.ManagedInstancePairInfoResponse[]>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Partner region information for the failover group.
     */
    public readonly partnerRegions!: pulumi.Output<outputs.sql.v20171001preview.PartnerRegionInfoResponse[]>;
    /**
     * Read-only endpoint of the failover group instance.
     */
    public readonly readOnlyEndpoint!: pulumi.Output<outputs.sql.v20171001preview.InstanceFailoverGroupReadOnlyEndpointResponse | undefined>;
    /**
     * Read-write endpoint of the failover group instance.
     */
    public readonly readWriteEndpoint!: pulumi.Output<outputs.sql.v20171001preview.InstanceFailoverGroupReadWriteEndpointResponse>;
    /**
     * Local replication role of the failover group instance.
     */
    public /*out*/ readonly replicationRole!: pulumi.Output<string>;
    /**
     * Replication state of the failover group instance.
     */
    public /*out*/ readonly replicationState!: pulumi.Output<string>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a InstanceFailoverGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceFailoverGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.failoverGroupName === undefined) {
                throw new Error("Missing required property 'failoverGroupName'");
            }
            if (!args || args.locationName === undefined) {
                throw new Error("Missing required property 'locationName'");
            }
            if (!args || args.managedInstancePairs === undefined) {
                throw new Error("Missing required property 'managedInstancePairs'");
            }
            if (!args || args.partnerRegions === undefined) {
                throw new Error("Missing required property 'partnerRegions'");
            }
            if (!args || args.readWriteEndpoint === undefined) {
                throw new Error("Missing required property 'readWriteEndpoint'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["failoverGroupName"] = args ? args.failoverGroupName : undefined;
            inputs["locationName"] = args ? args.locationName : undefined;
            inputs["managedInstancePairs"] = args ? args.managedInstancePairs : undefined;
            inputs["partnerRegions"] = args ? args.partnerRegions : undefined;
            inputs["readOnlyEndpoint"] = args ? args.readOnlyEndpoint : undefined;
            inputs["readWriteEndpoint"] = args ? args.readWriteEndpoint : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["replicationRole"] = undefined /*out*/;
            inputs["replicationState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["managedInstancePairs"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["partnerRegions"] = undefined /*out*/;
            inputs["readOnlyEndpoint"] = undefined /*out*/;
            inputs["readWriteEndpoint"] = undefined /*out*/;
            inputs["replicationRole"] = undefined /*out*/;
            inputs["replicationState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(InstanceFailoverGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a InstanceFailoverGroup resource.
 */
export interface InstanceFailoverGroupArgs {
    /**
     * The name of the failover group.
     */
    readonly failoverGroupName: pulumi.Input<string>;
    /**
     * The name of the region where the resource is located.
     */
    readonly locationName: pulumi.Input<string>;
    /**
     * List of managed instance pairs in the failover group.
     */
    readonly managedInstancePairs: pulumi.Input<pulumi.Input<inputs.sql.v20171001preview.ManagedInstancePairInfo>[]>;
    /**
     * Partner region information for the failover group.
     */
    readonly partnerRegions: pulumi.Input<pulumi.Input<inputs.sql.v20171001preview.PartnerRegionInfo>[]>;
    /**
     * Read-only endpoint of the failover group instance.
     */
    readonly readOnlyEndpoint?: pulumi.Input<inputs.sql.v20171001preview.InstanceFailoverGroupReadOnlyEndpoint>;
    /**
     * Read-write endpoint of the failover group instance.
     */
    readonly readWriteEndpoint: pulumi.Input<inputs.sql.v20171001preview.InstanceFailoverGroupReadWriteEndpoint>;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
