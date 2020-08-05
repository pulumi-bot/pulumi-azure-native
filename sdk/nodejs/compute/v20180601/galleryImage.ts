// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Specifies information about the gallery Image Definition that you want to create or update.
 */
export class GalleryImage extends pulumi.CustomResource {
    /**
     * Get an existing GalleryImage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): GalleryImage {
        return new GalleryImage(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:compute/v20180601:GalleryImage';

    /**
     * Returns true if the given object is an instance of GalleryImage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GalleryImage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GalleryImage.__pulumiType;
    }

    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Describes the properties of a gallery Image Definition.
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.compute.v20180601.GalleryImagePropertiesResponse>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a GalleryImage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GalleryImageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GalleryImageArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as GalleryImageArgs | undefined;
            if (!args || args.galleryName === undefined) {
                throw new Error("Missing required property 'galleryName'");
            }
            if (!args || args.identifier === undefined) {
                throw new Error("Missing required property 'identifier'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.osState === undefined) {
                throw new Error("Missing required property 'osState'");
            }
            if (!args || args.osType === undefined) {
                throw new Error("Missing required property 'osType'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["disallowed"] = args ? args.disallowed : undefined;
            inputs["endOfLifeDate"] = args ? args.endOfLifeDate : undefined;
            inputs["eula"] = args ? args.eula : undefined;
            inputs["galleryName"] = args ? args.galleryName : undefined;
            inputs["identifier"] = args ? args.identifier : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["osState"] = args ? args.osState : undefined;
            inputs["osType"] = args ? args.osType : undefined;
            inputs["privacyStatementUri"] = args ? args.privacyStatementUri : undefined;
            inputs["purchasePlan"] = args ? args.purchasePlan : undefined;
            inputs["recommended"] = args ? args.recommended : undefined;
            inputs["releaseNoteUri"] = args ? args.releaseNoteUri : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(GalleryImage.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a GalleryImage resource.
 */
export interface GalleryImageArgs {
    /**
     * The description of this gallery Image Definition resource. This property is updatable.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Describes the disallowed disk types.
     */
    readonly disallowed?: pulumi.Input<inputs.compute.v20180601.Disallowed>;
    /**
     * The end of life date of the gallery Image Definition. This property can be used for decommissioning purposes. This property is updatable.
     */
    readonly endOfLifeDate?: pulumi.Input<string>;
    /**
     * The Eula agreement for the gallery Image Definition.
     */
    readonly eula?: pulumi.Input<string>;
    /**
     * The name of the Shared Image Gallery in which the Image Definition is to be created.
     */
    readonly galleryName: pulumi.Input<string>;
    /**
     * This is the gallery Image Definition identifier.
     */
    readonly identifier: pulumi.Input<inputs.compute.v20180601.GalleryImageIdentifier>;
    /**
     * Resource location
     */
    readonly location: pulumi.Input<string>;
    /**
     * The name of the gallery Image Definition to be created or updated. The allowed characters are alphabets and numbers with dots, dashes, and periods allowed in the middle. The maximum length is 80 characters.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The allowed values for OS State are 'Generalized'.
     */
    readonly osState: pulumi.Input<string>;
    /**
     * This property allows you to specify the type of the OS that is included in the disk when creating a VM from a managed image. <br><br> Possible values are: <br><br> **Windows** <br><br> **Linux**
     */
    readonly osType: pulumi.Input<string>;
    /**
     * The privacy statement uri.
     */
    readonly privacyStatementUri?: pulumi.Input<string>;
    /**
     * Describes the gallery Image Definition purchase plan. This is used by marketplace images.
     */
    readonly purchasePlan?: pulumi.Input<inputs.compute.v20180601.ImagePurchasePlan>;
    /**
     * The properties describe the recommended machine configuration for this Image Definition. These properties are updatable.
     */
    readonly recommended?: pulumi.Input<inputs.compute.v20180601.RecommendedMachineConfiguration>;
    /**
     * The release note uri.
     */
    readonly releaseNoteUri?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
