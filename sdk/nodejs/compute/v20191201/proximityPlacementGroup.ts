// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Specifies information about the proximity placement group.
 */
export class ProximityPlacementGroup extends pulumi.CustomResource {
    /**
     * Get an existing ProximityPlacementGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ProximityPlacementGroup {
        return new ProximityPlacementGroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:compute/v20191201:ProximityPlacementGroup';

    /**
     * Returns true if the given object is an instance of ProximityPlacementGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProximityPlacementGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProximityPlacementGroup.__pulumiType;
    }

    /**
     * A list of references to all availability sets in the proximity placement group.
     */
    public /*out*/ readonly availabilitySets!: pulumi.Output<outputs.compute.v20191201.SubResourceWithColocationStatusResponse[]>;
    /**
     * Describes colocation status of the Proximity Placement Group.
     */
    public readonly colocationStatus!: pulumi.Output<outputs.compute.v20191201.InstanceViewStatusResponse | undefined>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Specifies the type of the proximity placement group. <br><br> Possible values are: <br><br> **Standard** : Co-locate resources within an Azure region or Availability Zone. <br><br> **Ultra** : For future use.
     */
    public readonly proximityPlacementGroupType!: pulumi.Output<string | undefined>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * A list of references to all virtual machine scale sets in the proximity placement group.
     */
    public /*out*/ readonly virtualMachineScaleSets!: pulumi.Output<outputs.compute.v20191201.SubResourceWithColocationStatusResponse[]>;
    /**
     * A list of references to all virtual machines in the proximity placement group.
     */
    public /*out*/ readonly virtualMachines!: pulumi.Output<outputs.compute.v20191201.SubResourceWithColocationStatusResponse[]>;

    /**
     * Create a ProximityPlacementGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProximityPlacementGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.proximityPlacementGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'proximityPlacementGroupName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["colocationStatus"] = args ? args.colocationStatus : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["proximityPlacementGroupName"] = args ? args.proximityPlacementGroupName : undefined;
            inputs["proximityPlacementGroupType"] = args ? args.proximityPlacementGroupType : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["availabilitySets"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["virtualMachineScaleSets"] = undefined /*out*/;
            inputs["virtualMachines"] = undefined /*out*/;
        } else {
            inputs["availabilitySets"] = undefined /*out*/;
            inputs["colocationStatus"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["proximityPlacementGroupType"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["virtualMachineScaleSets"] = undefined /*out*/;
            inputs["virtualMachines"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:compute/latest:ProximityPlacementGroup" }, { type: "azure-nextgen:compute/v20180401:ProximityPlacementGroup" }, { type: "azure-nextgen:compute/v20180601:ProximityPlacementGroup" }, { type: "azure-nextgen:compute/v20181001:ProximityPlacementGroup" }, { type: "azure-nextgen:compute/v20190301:ProximityPlacementGroup" }, { type: "azure-nextgen:compute/v20190701:ProximityPlacementGroup" }, { type: "azure-nextgen:compute/v20200601:ProximityPlacementGroup" }, { type: "azure-nextgen:compute/v20201201:ProximityPlacementGroup" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ProximityPlacementGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ProximityPlacementGroup resource.
 */
export interface ProximityPlacementGroupArgs {
    /**
     * Describes colocation status of the Proximity Placement Group.
     */
    readonly colocationStatus?: pulumi.Input<inputs.compute.v20191201.InstanceViewStatus>;
    /**
     * Resource location
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the proximity placement group.
     */
    readonly proximityPlacementGroupName: pulumi.Input<string>;
    /**
     * Specifies the type of the proximity placement group. <br><br> Possible values are: <br><br> **Standard** : Co-locate resources within an Azure region or Availability Zone. <br><br> **Ultra** : For future use.
     */
    readonly proximityPlacementGroupType?: pulumi.Input<string | enums.compute.v20191201.ProximityPlacementGroupType>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
