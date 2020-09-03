// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The source user image virtual hard disk. The virtual hard disk will be copied before being attached to the virtual machine. If SourceImage is provided, the destination virtual hard drive must not exist.
 *
 * ## Example Usage
 * ### Create a virtual machine image from a blob with DiskEncryptionSet resource.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const image = new azurerm.compute.v20191201.Image("image", {
 *     imageName: "myImage",
 *     location: "West US",
 *     resourceGroupName: "myResourceGroup",
 *     storageProfile: {
 *         osDisk: {
 *             blobUri: "https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd",
 *             diskEncryptionSet: {
 *                 id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}",
 *             },
 *             osState: "Generalized",
 *             osType: "Linux",
 *         },
 *     },
 * });
 *
 * ```
 * ### Create a virtual machine image from a blob.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const image = new azurerm.compute.v20191201.Image("image", {
 *     imageName: "myImage",
 *     location: "West US",
 *     resourceGroupName: "myResourceGroup",
 *     storageProfile: {
 *         osDisk: {
 *             blobUri: "https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd",
 *             osState: "Generalized",
 *             osType: "Linux",
 *         },
 *         zoneResilient: true,
 *     },
 * });
 *
 * ```
 * ### Create a virtual machine image from a managed disk with DiskEncryptionSet resource.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const image = new azurerm.compute.v20191201.Image("image", {
 *     imageName: "myImage",
 *     location: "West US",
 *     resourceGroupName: "myResourceGroup",
 *     storageProfile: {
 *         osDisk: {
 *             diskEncryptionSet: {
 *                 id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}",
 *             },
 *             managedDisk: {
 *                 id: "subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk",
 *             },
 *             osState: "Generalized",
 *             osType: "Linux",
 *         },
 *     },
 * });
 *
 * ```
 * ### Create a virtual machine image from a managed disk.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const image = new azurerm.compute.v20191201.Image("image", {
 *     imageName: "myImage",
 *     location: "West US",
 *     resourceGroupName: "myResourceGroup",
 *     storageProfile: {
 *         osDisk: {
 *             managedDisk: {
 *                 id: "subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk",
 *             },
 *             osState: "Generalized",
 *             osType: "Linux",
 *         },
 *         zoneResilient: true,
 *     },
 * });
 *
 * ```
 * ### Create a virtual machine image from a snapshot with DiskEncryptionSet resource.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const image = new azurerm.compute.v20191201.Image("image", {
 *     imageName: "myImage",
 *     location: "West US",
 *     resourceGroupName: "myResourceGroup",
 *     storageProfile: {
 *         osDisk: {
 *             diskEncryptionSet: {
 *                 id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}",
 *             },
 *             osState: "Generalized",
 *             osType: "Linux",
 *             snapshot: {
 *                 id: "subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot",
 *             },
 *         },
 *     },
 * });
 *
 * ```
 * ### Create a virtual machine image from a snapshot.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const image = new azurerm.compute.v20191201.Image("image", {
 *     imageName: "myImage",
 *     location: "West US",
 *     resourceGroupName: "myResourceGroup",
 *     storageProfile: {
 *         osDisk: {
 *             osState: "Generalized",
 *             osType: "Linux",
 *             snapshot: {
 *                 id: "subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot",
 *             },
 *         },
 *         zoneResilient: false,
 *     },
 * });
 *
 * ```
 * ### Create a virtual machine image from an existing virtual machine.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const image = new azurerm.compute.v20191201.Image("image", {
 *     imageName: "myImage",
 *     location: "West US",
 *     resourceGroupName: "myResourceGroup",
 *     sourceVirtualMachine: {
 *         id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM",
 *     },
 * });
 *
 * ```
 * ### Create a virtual machine image that includes a data disk from a blob.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const image = new azurerm.compute.v20191201.Image("image", {
 *     imageName: "myImage",
 *     location: "West US",
 *     resourceGroupName: "myResourceGroup",
 *     storageProfile: {
 *         dataDisks: [{
 *             blobUri: "https://mystorageaccount.blob.core.windows.net/dataimages/dataimage.vhd",
 *             lun: 1,
 *         }],
 *         osDisk: {
 *             blobUri: "https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd",
 *             osState: "Generalized",
 *             osType: "Linux",
 *         },
 *         zoneResilient: false,
 *     },
 * });
 *
 * ```
 * ### Create a virtual machine image that includes a data disk from a managed disk.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const image = new azurerm.compute.v20191201.Image("image", {
 *     imageName: "myImage",
 *     location: "West US",
 *     resourceGroupName: "myResourceGroup",
 *     storageProfile: {
 *         dataDisks: [{
 *             lun: 1,
 *             managedDisk: {
 *                 id: "subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk2",
 *             },
 *         }],
 *         osDisk: {
 *             managedDisk: {
 *                 id: "subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk",
 *             },
 *             osState: "Generalized",
 *             osType: "Linux",
 *         },
 *         zoneResilient: false,
 *     },
 * });
 *
 * ```
 * ### Create a virtual machine image that includes a data disk from a snapshot.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const image = new azurerm.compute.v20191201.Image("image", {
 *     imageName: "myImage",
 *     location: "West US",
 *     resourceGroupName: "myResourceGroup",
 *     storageProfile: {
 *         dataDisks: [{
 *             lun: 1,
 *             snapshot: {
 *                 id: "subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot2",
 *             },
 *         }],
 *         osDisk: {
 *             osState: "Generalized",
 *             osType: "Linux",
 *             snapshot: {
 *                 id: "subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot",
 *             },
 *         },
 *         zoneResilient: true,
 *     },
 * });
 *
 * ```
 */
export class Image extends pulumi.CustomResource {
    /**
     * Get an existing Image resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Image {
        return new Image(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:compute/v20191201:Image';

    /**
     * Returns true if the given object is an instance of Image.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Image {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Image.__pulumiType;
    }

    /**
     * Gets the HyperVGenerationType of the VirtualMachine created from the image
     */
    public readonly hyperVGeneration!: pulumi.Output<string | undefined>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The provisioning state.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The source virtual machine from which Image is created.
     */
    public readonly sourceVirtualMachine!: pulumi.Output<outputs.compute.v20191201.SubResourceResponse | undefined>;
    /**
     * Specifies the storage settings for the virtual machine disks.
     */
    public readonly storageProfile!: pulumi.Output<outputs.compute.v20191201.ImageStorageProfileResponse | undefined>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Image resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ImageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ImageArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ImageArgs | undefined;
            if (!args || args.imageName === undefined) {
                throw new Error("Missing required property 'imageName'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["hyperVGeneration"] = args ? args.hyperVGeneration : undefined;
            inputs["imageName"] = args ? args.imageName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sourceVirtualMachine"] = args ? args.sourceVirtualMachine : undefined;
            inputs["storageProfile"] = args ? args.storageProfile : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:compute/latest:Image" }, { type: "azurerm:compute/v20160430preview:Image" }, { type: "azurerm:compute/v20170330:Image" }, { type: "azurerm:compute/v20171201:Image" }, { type: "azurerm:compute/v20180401:Image" }, { type: "azurerm:compute/v20180601:Image" }, { type: "azurerm:compute/v20181001:Image" }, { type: "azurerm:compute/v20190301:Image" }, { type: "azurerm:compute/v20190701:Image" }, { type: "azurerm:compute/v20200601:Image" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Image.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Image resource.
 */
export interface ImageArgs {
    /**
     * Gets the HyperVGenerationType of the VirtualMachine created from the image
     */
    readonly hyperVGeneration?: pulumi.Input<string>;
    /**
     * The name of the image.
     */
    readonly imageName: pulumi.Input<string>;
    /**
     * Resource location
     */
    readonly location: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The source virtual machine from which Image is created.
     */
    readonly sourceVirtualMachine?: pulumi.Input<inputs.compute.v20191201.SubResource>;
    /**
     * Specifies the storage settings for the virtual machine disks.
     */
    readonly storageProfile?: pulumi.Input<inputs.compute.v20191201.ImageStorageProfile>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
