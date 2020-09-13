// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Specifies information about the gallery Application Definition that you want to create or update.
 */
export class GalleryApplication extends pulumi.CustomResource {
    /**
     * Get an existing GalleryApplication resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): GalleryApplication {
        return new GalleryApplication(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:compute/v20190701:GalleryApplication';

    /**
     * Returns true if the given object is an instance of GalleryApplication.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GalleryApplication {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GalleryApplication.__pulumiType;
    }

    /**
     * The description of this gallery Application Definition resource. This property is updatable.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The end of life date of the gallery Application Definition. This property can be used for decommissioning purposes. This property is updatable.
     */
    public readonly endOfLifeDate!: pulumi.Output<string | undefined>;
    /**
     * The Eula agreement for the gallery Application Definition.
     */
    public readonly eula!: pulumi.Output<string | undefined>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The privacy statement uri.
     */
    public readonly privacyStatementUri!: pulumi.Output<string | undefined>;
    /**
     * The release note uri.
     */
    public readonly releaseNoteUri!: pulumi.Output<string | undefined>;
    /**
     * This property allows you to specify the supported type of the OS that application is built for. <br><br> Possible values are: <br><br> **Windows** <br><br> **Linux**
     */
    public readonly supportedOSType!: pulumi.Output<string>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a GalleryApplication resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GalleryApplicationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.galleryApplicationName === undefined) {
                throw new Error("Missing required property 'galleryApplicationName'");
            }
            if (!args || args.galleryName === undefined) {
                throw new Error("Missing required property 'galleryName'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.supportedOSType === undefined) {
                throw new Error("Missing required property 'supportedOSType'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["endOfLifeDate"] = args ? args.endOfLifeDate : undefined;
            inputs["eula"] = args ? args.eula : undefined;
            inputs["galleryApplicationName"] = args ? args.galleryApplicationName : undefined;
            inputs["galleryName"] = args ? args.galleryName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["privacyStatementUri"] = args ? args.privacyStatementUri : undefined;
            inputs["releaseNoteUri"] = args ? args.releaseNoteUri : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["supportedOSType"] = args ? args.supportedOSType : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["description"] = undefined /*out*/;
            inputs["endOfLifeDate"] = undefined /*out*/;
            inputs["eula"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["privacyStatementUri"] = undefined /*out*/;
            inputs["releaseNoteUri"] = undefined /*out*/;
            inputs["supportedOSType"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:compute/latest:GalleryApplication" }, { type: "azurerm:compute/v20190301:GalleryApplication" }, { type: "azurerm:compute/v20191201:GalleryApplication" }, { type: "azurerm:compute/v20200930:GalleryApplication" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(GalleryApplication.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a GalleryApplication resource.
 */
export interface GalleryApplicationArgs {
    /**
     * The description of this gallery Application Definition resource. This property is updatable.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The end of life date of the gallery Application Definition. This property can be used for decommissioning purposes. This property is updatable.
     */
    readonly endOfLifeDate?: pulumi.Input<string>;
    /**
     * The Eula agreement for the gallery Application Definition.
     */
    readonly eula?: pulumi.Input<string>;
    /**
     * The name of the gallery Application Definition to be created or updated. The allowed characters are alphabets and numbers with dots, dashes, and periods allowed in the middle. The maximum length is 80 characters.
     */
    readonly galleryApplicationName: pulumi.Input<string>;
    /**
     * The name of the Shared Application Gallery in which the Application Definition is to be created.
     */
    readonly galleryName: pulumi.Input<string>;
    /**
     * Resource location
     */
    readonly location: pulumi.Input<string>;
    /**
     * The privacy statement uri.
     */
    readonly privacyStatementUri?: pulumi.Input<string>;
    /**
     * The release note uri.
     */
    readonly releaseNoteUri?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * This property allows you to specify the supported type of the OS that application is built for. <br><br> Possible values are: <br><br> **Windows** <br><br> **Linux**
     */
    readonly supportedOSType: pulumi.Input<string>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
