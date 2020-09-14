// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * HANA instance info on Azure (ARM properties and HANA properties)
 */
export class HanaInstance extends pulumi.CustomResource {
    /**
     * Get an existing HanaInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): HanaInstance {
        return new HanaInstance(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:hanaonazure/v20171103preview:HanaInstance';

    /**
     * Returns true if the given object is an instance of HanaInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HanaInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HanaInstance.__pulumiType;
    }

    /**
     * Specifies the HANA instance unique ID.
     */
    public /*out*/ readonly hanaInstanceId!: pulumi.Output<string>;
    /**
     * Specifies the hardware settings for the HANA instance.
     */
    public /*out*/ readonly hardwareProfile!: pulumi.Output<outputs.hanaonazure.v20171103preview.HardwareProfileResponse | undefined>;
    /**
     * Hardware revision of a HANA instance
     */
    public /*out*/ readonly hwRevision!: pulumi.Output<string>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Specifies the network settings for the HANA instance.
     */
    public readonly networkProfile!: pulumi.Output<outputs.hanaonazure.v20171103preview.NetworkProfileResponse | undefined>;
    /**
     * Specifies the operating system settings for the HANA instance.
     */
    public readonly osProfile!: pulumi.Output<outputs.hanaonazure.v20171103preview.OSProfileResponse | undefined>;
    /**
     * ARM ID of another HanaInstance that will share a network with this HanaInstance
     */
    public readonly partnerNodeId!: pulumi.Output<string | undefined>;
    /**
     * Resource power state
     */
    public /*out*/ readonly powerState!: pulumi.Output<string>;
    /**
     * State of provisioning of the HanaInstance
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Resource proximity placement group
     */
    public /*out*/ readonly proximityPlacementGroup!: pulumi.Output<string>;
    /**
     * Specifies the storage settings for the HANA instance disks.
     */
    public readonly storageProfile!: pulumi.Output<outputs.hanaonazure.v20171103preview.StorageProfileResponse | undefined>;
    /**
     * Resource tags
     */
    public /*out*/ readonly tags!: pulumi.Output<{[key: string]: string}>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a HanaInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HanaInstanceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.hanaInstanceName === undefined) {
                throw new Error("Missing required property 'hanaInstanceName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["hanaInstanceName"] = args ? args.hanaInstanceName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["networkProfile"] = args ? args.networkProfile : undefined;
            inputs["osProfile"] = args ? args.osProfile : undefined;
            inputs["partnerNodeId"] = args ? args.partnerNodeId : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["storageProfile"] = args ? args.storageProfile : undefined;
            inputs["hanaInstanceId"] = undefined /*out*/;
            inputs["hardwareProfile"] = undefined /*out*/;
            inputs["hwRevision"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["powerState"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["proximityPlacementGroup"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["hanaInstanceId"] = undefined /*out*/;
            inputs["hardwareProfile"] = undefined /*out*/;
            inputs["hwRevision"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["networkProfile"] = undefined /*out*/;
            inputs["osProfile"] = undefined /*out*/;
            inputs["partnerNodeId"] = undefined /*out*/;
            inputs["powerState"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["proximityPlacementGroup"] = undefined /*out*/;
            inputs["storageProfile"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:hanaonazure/latest:HanaInstance" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(HanaInstance.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a HanaInstance resource.
 */
export interface HanaInstanceArgs {
    /**
     * Name of the SAP HANA on Azure instance.
     */
    readonly hanaInstanceName: pulumi.Input<string>;
    /**
     * Resource location
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the network settings for the HANA instance.
     */
    readonly networkProfile?: pulumi.Input<inputs.hanaonazure.v20171103preview.NetworkProfile>;
    /**
     * Specifies the operating system settings for the HANA instance.
     */
    readonly osProfile?: pulumi.Input<inputs.hanaonazure.v20171103preview.OSProfile>;
    /**
     * ARM ID of another HanaInstance that will share a network with this HanaInstance
     */
    readonly partnerNodeId?: pulumi.Input<string>;
    /**
     * Name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Specifies the storage settings for the HANA instance disks.
     */
    readonly storageProfile?: pulumi.Input<inputs.hanaonazure.v20171103preview.StorageProfile>;
}
