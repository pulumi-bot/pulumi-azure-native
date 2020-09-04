// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Specifies information about the gallery Image Version that you want to create or update.
 *
 * ## Example Usage
 * ### Create or update a simple gallery Image Version.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const galleryImageVersion = new azurerm.compute.v20180601.GalleryImageVersion("galleryImageVersion", {
 *     galleryImageName: "myGalleryImageName",
 *     galleryImageVersionName: "1.0.0",
 *     galleryName: "myGalleryName",
 *     location: "West US",
 *     publishingProfile: {
 *         source: {
 *             managedImage: {
 *                 id: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/images/{imageName}",
 *             },
 *         },
 *         targetRegions: [
 *             {
 *                 name: "West US",
 *                 regionalReplicaCount: 1,
 *             },
 *             {
 *                 name: "East US",
 *                 regionalReplicaCount: 2,
 *             },
 *         ],
 *     },
 *     resourceGroupName: "myResourceGroup",
 * });
 *
 * ```
 */
export class GalleryImageVersion extends pulumi.CustomResource {
    /**
     * Get an existing GalleryImageVersion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): GalleryImageVersion {
        return new GalleryImageVersion(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:compute/v20180601:GalleryImageVersion';

    /**
     * Returns true if the given object is an instance of GalleryImageVersion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GalleryImageVersion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GalleryImageVersion.__pulumiType;
    }

    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The provisioning state, which only appears in the response.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The publishing profile of a gallery Image Version.
     */
    public readonly publishingProfile!: pulumi.Output<outputs.compute.v20180601.GalleryImageVersionPublishingProfileResponse>;
    /**
     * This is the replication status of the gallery Image Version.
     */
    public /*out*/ readonly replicationStatus!: pulumi.Output<outputs.compute.v20180601.ReplicationStatusResponse>;
    /**
     * This is the storage profile of a gallery Image Version.
     */
    public /*out*/ readonly storageProfile!: pulumi.Output<outputs.compute.v20180601.GalleryImageVersionStorageProfileResponse>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a GalleryImageVersion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GalleryImageVersionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.galleryImageName === undefined) {
                throw new Error("Missing required property 'galleryImageName'");
            }
            if (!args || args.galleryImageVersionName === undefined) {
                throw new Error("Missing required property 'galleryImageVersionName'");
            }
            if (!args || args.galleryName === undefined) {
                throw new Error("Missing required property 'galleryName'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.publishingProfile === undefined) {
                throw new Error("Missing required property 'publishingProfile'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["galleryImageName"] = args ? args.galleryImageName : undefined;
            inputs["galleryImageVersionName"] = args ? args.galleryImageVersionName : undefined;
            inputs["galleryName"] = args ? args.galleryName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["publishingProfile"] = args ? args.publishingProfile : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["replicationStatus"] = undefined /*out*/;
            inputs["storageProfile"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["publishingProfile"] = undefined /*out*/;
            inputs["replicationStatus"] = undefined /*out*/;
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
        const aliasOpts = { aliases: [{ type: "azurerm:compute/latest:GalleryImageVersion" }, { type: "azurerm:compute/v20190301:GalleryImageVersion" }, { type: "azurerm:compute/v20190701:GalleryImageVersion" }, { type: "azurerm:compute/v20191201:GalleryImageVersion" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(GalleryImageVersion.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a GalleryImageVersion resource.
 */
export interface GalleryImageVersionArgs {
    /**
     * The name of the gallery Image Definition in which the Image Version is to be created.
     */
    readonly galleryImageName: pulumi.Input<string>;
    /**
     * The name of the gallery Image Version to be created. Needs to follow semantic version name pattern: The allowed characters are digit and period. Digits must be within the range of a 32-bit integer. Format: <MajorVersion>.<MinorVersion>.<Patch>
     */
    readonly galleryImageVersionName: pulumi.Input<string>;
    /**
     * The name of the Shared Image Gallery in which the Image Definition resides.
     */
    readonly galleryName: pulumi.Input<string>;
    /**
     * Resource location
     */
    readonly location: pulumi.Input<string>;
    /**
     * The publishing profile of a gallery Image Version.
     */
    readonly publishingProfile: pulumi.Input<inputs.compute.v20180601.GalleryImageVersionPublishingProfile>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
