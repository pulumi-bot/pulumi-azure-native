// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Specifies information about the Dedicated host.
 */
export class DedicatedHost extends pulumi.CustomResource {
    /**
     * Get an existing DedicatedHost resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DedicatedHost {
        return new DedicatedHost(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:compute/v20191201:DedicatedHost';

    /**
     * Returns true if the given object is an instance of DedicatedHost.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DedicatedHost {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DedicatedHost.__pulumiType;
    }

    /**
     * Specifies whether the dedicated host should be replaced automatically in case of a failure. The value is defaulted to 'true' when not provided.
     */
    public readonly autoReplaceOnFailure!: pulumi.Output<boolean | undefined>;
    /**
     * A unique id generated and assigned to the dedicated host by the platform. <br><br> Does not change throughout the lifetime of the host.
     */
    public /*out*/ readonly hostId!: pulumi.Output<string>;
    /**
     * The dedicated host instance view.
     */
    public /*out*/ readonly instanceView!: pulumi.Output<outputs.compute.v20191201.DedicatedHostInstanceViewResponse>;
    /**
     * Specifies the software license type that will be applied to the VMs deployed on the dedicated host. <br><br> Possible values are: <br><br> **None** <br><br> **Windows_Server_Hybrid** <br><br> **Windows_Server_Perpetual** <br><br> Default: **None**
     */
    public readonly licenseType!: pulumi.Output<string | undefined>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Fault domain of the dedicated host within a dedicated host group.
     */
    public readonly platformFaultDomain!: pulumi.Output<number | undefined>;
    /**
     * The provisioning state, which only appears in the response.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The date when the host was first provisioned.
     */
    public /*out*/ readonly provisioningTime!: pulumi.Output<string>;
    /**
     * SKU of the dedicated host for Hardware Generation and VM family. Only name is required to be set. List Microsoft.Compute SKUs for a list of possible values.
     */
    public readonly sku!: pulumi.Output<outputs.compute.v20191201.SkuResponse>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * A list of references to all virtual machines in the Dedicated Host.
     */
    public /*out*/ readonly virtualMachines!: pulumi.Output<outputs.compute.v20191201.SubResourceReadOnlyResponse[]>;

    /**
     * Create a DedicatedHost resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DedicatedHostArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.hostGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'hostGroupName'");
            }
            if ((!args || args.hostName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'hostName'");
            }
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.sku === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'sku'");
            }
            inputs["autoReplaceOnFailure"] = args ? args.autoReplaceOnFailure : undefined;
            inputs["hostGroupName"] = args ? args.hostGroupName : undefined;
            inputs["hostName"] = args ? args.hostName : undefined;
            inputs["licenseType"] = args ? args.licenseType : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["platformFaultDomain"] = args ? args.platformFaultDomain : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["hostId"] = undefined /*out*/;
            inputs["instanceView"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["provisioningTime"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["virtualMachines"] = undefined /*out*/;
        } else {
            inputs["autoReplaceOnFailure"] = undefined /*out*/;
            inputs["hostId"] = undefined /*out*/;
            inputs["instanceView"] = undefined /*out*/;
            inputs["licenseType"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["platformFaultDomain"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["provisioningTime"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["virtualMachines"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:compute/latest:DedicatedHost" }, { type: "azure-nextgen:compute/v20190301:DedicatedHost" }, { type: "azure-nextgen:compute/v20190701:DedicatedHost" }, { type: "azure-nextgen:compute/v20200601:DedicatedHost" }, { type: "azure-nextgen:compute/v20201201:DedicatedHost" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(DedicatedHost.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DedicatedHost resource.
 */
export interface DedicatedHostArgs {
    /**
     * Specifies whether the dedicated host should be replaced automatically in case of a failure. The value is defaulted to 'true' when not provided.
     */
    readonly autoReplaceOnFailure?: pulumi.Input<boolean>;
    /**
     * The name of the dedicated host group.
     */
    readonly hostGroupName: pulumi.Input<string>;
    /**
     * The name of the dedicated host .
     */
    readonly hostName: pulumi.Input<string>;
    /**
     * Specifies the software license type that will be applied to the VMs deployed on the dedicated host. <br><br> Possible values are: <br><br> **None** <br><br> **Windows_Server_Hybrid** <br><br> **Windows_Server_Perpetual** <br><br> Default: **None**
     */
    readonly licenseType?: pulumi.Input<enums.compute.v20191201.DedicatedHostLicenseTypes>;
    /**
     * Resource location
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Fault domain of the dedicated host within a dedicated host group.
     */
    readonly platformFaultDomain?: pulumi.Input<number>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * SKU of the dedicated host for Hardware Generation and VM family. Only name is required to be set. List Microsoft.Compute SKUs for a list of possible values.
     */
    readonly sku: pulumi.Input<inputs.compute.v20191201.Sku>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
