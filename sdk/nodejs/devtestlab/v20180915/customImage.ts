// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A custom image.
 */
export class CustomImage extends pulumi.CustomResource {
    /**
     * Get an existing CustomImage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CustomImage {
        return new CustomImage(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:devtestlab/v20180915:CustomImage';

    /**
     * Returns true if the given object is an instance of CustomImage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomImage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomImage.__pulumiType;
    }

    /**
     * The author of the custom image.
     */
    public readonly author!: pulumi.Output<string | undefined>;
    /**
     * The creation date of the custom image.
     */
    public /*out*/ readonly creationDate!: pulumi.Output<string>;
    /**
     * Storage information about the plan related to this custom image
     */
    public readonly customImagePlan!: pulumi.Output<outputs.devtestlab.v20180915.CustomImagePropertiesFromPlanResponse | undefined>;
    /**
     * Storage information about the data disks present in the custom image
     */
    public readonly dataDiskStorageInfo!: pulumi.Output<outputs.devtestlab.v20180915.DataDiskStorageTypeInfoResponse[] | undefined>;
    /**
     * The description of the custom image.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Whether or not the custom images underlying offer/plan has been enabled for programmatic deployment
     */
    public readonly isPlanAuthorized!: pulumi.Output<boolean | undefined>;
    /**
     * The location of the resource.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The Managed Image Id backing the custom image.
     */
    public readonly managedImageId!: pulumi.Output<string | undefined>;
    /**
     * The Managed Snapshot Id backing the custom image.
     */
    public readonly managedSnapshotId!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The provisioning status of the resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
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
     * The VHD from which the image is to be created.
     */
    public readonly vhd!: pulumi.Output<outputs.devtestlab.v20180915.CustomImagePropertiesCustomResponse | undefined>;
    /**
     * The virtual machine from which the image is to be created.
     */
    public readonly vm!: pulumi.Output<outputs.devtestlab.v20180915.CustomImagePropertiesFromVmResponse | undefined>;

    /**
     * Create a CustomImage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomImageArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.labName === undefined) {
                throw new Error("Missing required property 'labName'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["author"] = args ? args.author : undefined;
            inputs["customImagePlan"] = args ? args.customImagePlan : undefined;
            inputs["dataDiskStorageInfo"] = args ? args.dataDiskStorageInfo : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["isPlanAuthorized"] = args ? args.isPlanAuthorized : undefined;
            inputs["labName"] = args ? args.labName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["managedImageId"] = args ? args.managedImageId : undefined;
            inputs["managedSnapshotId"] = args ? args.managedSnapshotId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vhd"] = args ? args.vhd : undefined;
            inputs["vm"] = args ? args.vm : undefined;
            inputs["creationDate"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["uniqueIdentifier"] = undefined /*out*/;
        } else {
            inputs["author"] = undefined /*out*/;
            inputs["creationDate"] = undefined /*out*/;
            inputs["customImagePlan"] = undefined /*out*/;
            inputs["dataDiskStorageInfo"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["isPlanAuthorized"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["managedImageId"] = undefined /*out*/;
            inputs["managedSnapshotId"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["uniqueIdentifier"] = undefined /*out*/;
            inputs["vhd"] = undefined /*out*/;
            inputs["vm"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:devtestlab/latest:CustomImage" }, { type: "azurerm:devtestlab/v20150521preview:CustomImage" }, { type: "azurerm:devtestlab/v20160515:CustomImage" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(CustomImage.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a CustomImage resource.
 */
export interface CustomImageArgs {
    /**
     * The author of the custom image.
     */
    readonly author?: pulumi.Input<string>;
    /**
     * Storage information about the plan related to this custom image
     */
    readonly customImagePlan?: pulumi.Input<inputs.devtestlab.v20180915.CustomImagePropertiesFromPlan>;
    /**
     * Storage information about the data disks present in the custom image
     */
    readonly dataDiskStorageInfo?: pulumi.Input<pulumi.Input<inputs.devtestlab.v20180915.DataDiskStorageTypeInfo>[]>;
    /**
     * The description of the custom image.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Whether or not the custom images underlying offer/plan has been enabled for programmatic deployment
     */
    readonly isPlanAuthorized?: pulumi.Input<boolean>;
    /**
     * The name of the lab.
     */
    readonly labName: pulumi.Input<string>;
    /**
     * The location of the resource.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The Managed Image Id backing the custom image.
     */
    readonly managedImageId?: pulumi.Input<string>;
    /**
     * The Managed Snapshot Id backing the custom image.
     */
    readonly managedSnapshotId?: pulumi.Input<string>;
    /**
     * The name of the custom image.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The tags of the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The VHD from which the image is to be created.
     */
    readonly vhd?: pulumi.Input<inputs.devtestlab.v20180915.CustomImagePropertiesCustom>;
    /**
     * The virtual machine from which the image is to be created.
     */
    readonly vm?: pulumi.Input<inputs.devtestlab.v20180915.CustomImagePropertiesFromVm>;
}
