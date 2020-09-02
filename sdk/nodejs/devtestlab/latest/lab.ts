// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A lab.
 */
export class Lab extends pulumi.CustomResource {
    /**
     * Get an existing Lab resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Lab {
        return new Lab(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:devtestlab/latest:Lab';

    /**
     * Returns true if the given object is an instance of Lab.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Lab {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Lab.__pulumiType;
    }

    /**
     * The properties of any lab announcement associated with this lab
     */
    public readonly announcement!: pulumi.Output<outputs.devtestlab.latest.LabAnnouncementPropertiesResponse | undefined>;
    /**
     * The lab's artifact storage account.
     */
    public /*out*/ readonly artifactsStorageAccount!: pulumi.Output<string>;
    /**
     * The creation date of the lab.
     */
    public /*out*/ readonly createdDate!: pulumi.Output<string>;
    /**
     * The lab's default premium storage account.
     */
    public /*out*/ readonly defaultPremiumStorageAccount!: pulumi.Output<string>;
    /**
     * The lab's default storage account.
     */
    public /*out*/ readonly defaultStorageAccount!: pulumi.Output<string>;
    /**
     * The access rights to be granted to the user when provisioning an environment
     */
    public readonly environmentPermission!: pulumi.Output<string | undefined>;
    /**
     * Extended properties of the lab used for experimental features
     */
    public readonly extendedProperties!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Type of storage used by the lab. It can be either Premium or Standard. Default is Premium.
     */
    public readonly labStorageType!: pulumi.Output<string | undefined>;
    /**
     * The load balancer used to for lab VMs that use shared IP address.
     */
    public /*out*/ readonly loadBalancerId!: pulumi.Output<string>;
    /**
     * The location of the resource.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The ordered list of artifact resource IDs that should be applied on all Linux VM creations by default, prior to the artifacts specified by the user.
     */
    public readonly mandatoryArtifactsResourceIdsLinux!: pulumi.Output<string[] | undefined>;
    /**
     * The ordered list of artifact resource IDs that should be applied on all Windows VM creations by default, prior to the artifacts specified by the user.
     */
    public readonly mandatoryArtifactsResourceIdsWindows!: pulumi.Output<string[] | undefined>;
    /**
     * The name of the resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The Network Security Group attached to the lab VMs Network interfaces to restrict open ports.
     */
    public /*out*/ readonly networkSecurityGroupId!: pulumi.Output<string>;
    /**
     * The lab's premium data disk storage account.
     */
    public /*out*/ readonly premiumDataDiskStorageAccount!: pulumi.Output<string>;
    /**
     * The setting to enable usage of premium data disks.
     * When its value is 'Enabled', creation of standard or premium data disks is allowed.
     * When its value is 'Disabled', only creation of standard data disks is allowed.
     */
    public readonly premiumDataDisks!: pulumi.Output<string | undefined>;
    /**
     * The provisioning status of the resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The public IP address for the lab's load balancer.
     */
    public /*out*/ readonly publicIpId!: pulumi.Output<string>;
    /**
     * The properties of any lab support message associated with this lab
     */
    public readonly support!: pulumi.Output<outputs.devtestlab.latest.LabSupportPropertiesResponse | undefined>;
    /**
     * The tags of the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The unique immutable identifier of a resource (Guid).
     */
    public /*out*/ readonly uniqueIdentifier!: pulumi.Output<string>;
    /**
     * The lab's Key vault.
     */
    public /*out*/ readonly vaultName!: pulumi.Output<string>;
    /**
     * The resource group in which all new lab virtual machines will be created. To let DevTest Labs manage resource group creation, set this value to null.
     */
    public /*out*/ readonly vmCreationResourceGroup!: pulumi.Output<string>;

    /**
     * Create a Lab resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LabArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LabArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as LabArgs | undefined;
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["announcement"] = args ? args.announcement : undefined;
            inputs["environmentPermission"] = args ? args.environmentPermission : undefined;
            inputs["extendedProperties"] = args ? args.extendedProperties : undefined;
            inputs["labStorageType"] = args ? args.labStorageType : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["mandatoryArtifactsResourceIdsLinux"] = args ? args.mandatoryArtifactsResourceIdsLinux : undefined;
            inputs["mandatoryArtifactsResourceIdsWindows"] = args ? args.mandatoryArtifactsResourceIdsWindows : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["premiumDataDisks"] = args ? args.premiumDataDisks : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["support"] = args ? args.support : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["artifactsStorageAccount"] = undefined /*out*/;
            inputs["createdDate"] = undefined /*out*/;
            inputs["defaultPremiumStorageAccount"] = undefined /*out*/;
            inputs["defaultStorageAccount"] = undefined /*out*/;
            inputs["loadBalancerId"] = undefined /*out*/;
            inputs["networkSecurityGroupId"] = undefined /*out*/;
            inputs["premiumDataDiskStorageAccount"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["publicIpId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["uniqueIdentifier"] = undefined /*out*/;
            inputs["vaultName"] = undefined /*out*/;
            inputs["vmCreationResourceGroup"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:devtestlab/v20150521preview:Lab" }, { type: "azurerm:devtestlab/v20160515:Lab" }, { type: "azurerm:devtestlab/v20180915:Lab" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Lab.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Lab resource.
 */
export interface LabArgs {
    /**
     * The properties of any lab announcement associated with this lab
     */
    readonly announcement?: pulumi.Input<inputs.devtestlab.latest.LabAnnouncementProperties>;
    /**
     * The access rights to be granted to the user when provisioning an environment
     */
    readonly environmentPermission?: pulumi.Input<string>;
    /**
     * Extended properties of the lab used for experimental features
     */
    readonly extendedProperties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Type of storage used by the lab. It can be either Premium or Standard. Default is Premium.
     */
    readonly labStorageType?: pulumi.Input<string>;
    /**
     * The location of the resource.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The ordered list of artifact resource IDs that should be applied on all Linux VM creations by default, prior to the artifacts specified by the user.
     */
    readonly mandatoryArtifactsResourceIdsLinux?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ordered list of artifact resource IDs that should be applied on all Windows VM creations by default, prior to the artifacts specified by the user.
     */
    readonly mandatoryArtifactsResourceIdsWindows?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the lab.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The setting to enable usage of premium data disks.
     * When its value is 'Enabled', creation of standard or premium data disks is allowed.
     * When its value is 'Disabled', only creation of standard data disks is allowed.
     */
    readonly premiumDataDisks?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The properties of any lab support message associated with this lab
     */
    readonly support?: pulumi.Input<inputs.devtestlab.latest.LabSupportProperties>;
    /**
     * The tags of the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
