// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The link resource format.
 */
export class Link extends pulumi.CustomResource {
    /**
     * Get an existing Link resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Link {
        return new Link(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:customerinsights/v20170101:Link';

    /**
     * Returns true if the given object is an instance of Link.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Link {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Link.__pulumiType;
    }

    /**
     * Resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The definition of Link.
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.customerinsights.v20170101.LinkDefinitionResponse>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Link resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LinkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LinkArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as LinkArgs | undefined;
            if (!args || args.hubName === undefined) {
                throw new Error("Missing required property 'hubName'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.participantPropertyReferences === undefined) {
                throw new Error("Missing required property 'participantPropertyReferences'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sourceInteractionType === undefined) {
                throw new Error("Missing required property 'sourceInteractionType'");
            }
            if (!args || args.targetProfileType === undefined) {
                throw new Error("Missing required property 'targetProfileType'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["hubName"] = args ? args.hubName : undefined;
            inputs["mappings"] = args ? args.mappings : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["operationType"] = args ? args.operationType : undefined;
            inputs["participantPropertyReferences"] = args ? args.participantPropertyReferences : undefined;
            inputs["referenceOnly"] = args ? args.referenceOnly : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sourceInteractionType"] = args ? args.sourceInteractionType : undefined;
            inputs["targetProfileType"] = args ? args.targetProfileType : undefined;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Link.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Link resource.
 */
export interface LinkArgs {
    /**
     * Localized descriptions for the Link.
     */
    readonly description?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Localized display name for the Link.
     */
    readonly displayName?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the hub.
     */
    readonly hubName: pulumi.Input<string>;
    /**
     * The set of properties mappings between the source and target Types.
     */
    readonly mappings?: pulumi.Input<pulumi.Input<inputs.customerinsights.v20170101.TypePropertiesMapping>[]>;
    /**
     * The name of the link.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Determines whether this link is supposed to create or delete instances if Link is NOT Reference Only.
     */
    readonly operationType?: pulumi.Input<string>;
    /**
     * The properties that represent the participating profile.
     */
    readonly participantPropertyReferences: pulumi.Input<pulumi.Input<inputs.customerinsights.v20170101.ParticipantPropertyReference>[]>;
    /**
     * Indicating whether the link is reference only link. This flag is ignored if the Mappings are defined. If the mappings are not defined and it is set to true, links processing will not create or update profiles.
     */
    readonly referenceOnly?: pulumi.Input<boolean>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Name of the source Interaction Type.
     */
    readonly sourceInteractionType: pulumi.Input<string>;
    /**
     * Name of the target Profile Type.
     */
    readonly targetProfileType: pulumi.Input<string>;
}
