// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Virtual machine guest diagnostic settings resource.
 */
export class GuestDiagnosticsSettingsAssociation extends pulumi.CustomResource {
    /**
     * Get an existing GuestDiagnosticsSettingsAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): GuestDiagnosticsSettingsAssociation {
        return new GuestDiagnosticsSettingsAssociation(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:insights/v20180601preview:GuestDiagnosticsSettingsAssociation';

    /**
     * Returns true if the given object is an instance of GuestDiagnosticsSettingsAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GuestDiagnosticsSettingsAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GuestDiagnosticsSettingsAssociation.__pulumiType;
    }

    /**
     * The guest diagnostic settings name.
     */
    public readonly guestDiagnosticSettingsName!: pulumi.Output<string>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Azure resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Azure resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a GuestDiagnosticsSettingsAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GuestDiagnosticsSettingsAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GuestDiagnosticsSettingsAssociationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as GuestDiagnosticsSettingsAssociationArgs | undefined;
            if (!args || args.associationName === undefined) {
                throw new Error("Missing required property 'associationName'");
            }
            if (!args || args.guestDiagnosticSettingsName === undefined) {
                throw new Error("Missing required property 'guestDiagnosticSettingsName'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceUri === undefined) {
                throw new Error("Missing required property 'resourceUri'");
            }
            inputs["associationName"] = args ? args.associationName : undefined;
            inputs["guestDiagnosticSettingsName"] = args ? args.guestDiagnosticSettingsName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceUri"] = args ? args.resourceUri : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(GuestDiagnosticsSettingsAssociation.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a GuestDiagnosticsSettingsAssociation resource.
 */
export interface GuestDiagnosticsSettingsAssociationArgs {
    /**
     * The name of the diagnostic settings association.
     */
    readonly associationName: pulumi.Input<string>;
    /**
     * The guest diagnostic settings name.
     */
    readonly guestDiagnosticSettingsName: pulumi.Input<string>;
    /**
     * Resource location
     */
    readonly location: pulumi.Input<string>;
    /**
     * The fully qualified ID of the resource, including the resource name and resource type.
     */
    readonly resourceUri: pulumi.Input<string>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
