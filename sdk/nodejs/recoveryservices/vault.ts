// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Resource information, as returned by the resource provider.
 */
export class Vault extends pulumi.CustomResource {
    /**
     * Get an existing Vault resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Vault {
        return new Vault(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:recoveryservices:Vault';

    /**
     * Returns true if the given object is an instance of Vault.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Vault {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Vault.__pulumiType;
    }

    /**
     * Optional ETag.
     */
    public readonly eTag!: pulumi.Output<string | undefined>;
    /**
     * Identity for the resource.
     */
    public readonly identity!: pulumi.Output<outputs.recoveryservices.IdentityDataResponse | undefined>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource name associated with the resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Properties of the vault.
     */
    public readonly properties!: pulumi.Output<outputs.recoveryservices.VaultPropertiesResponse>;
    /**
     * Identifies the unique system identifier for each Azure resource.
     */
    public readonly sku!: pulumi.Output<outputs.recoveryservices.SkuResponse | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Vault resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: VaultArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
        inputs["eTag"] = args ? args.eTag : undefined;
        inputs["identity"] = args ? args.identity : undefined;
        inputs["location"] = args ? args.location : undefined;
        inputs["name"] = args ? args.name : undefined;
        inputs["properties"] = args ? args.properties : undefined;
        inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        inputs["sku"] = args ? args.sku : undefined;
        inputs["tags"] = args ? args.tags : undefined;
        inputs["type"] = undefined /*out*/;
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Vault.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Vault resource.
 */
export interface VaultArgs {
    /**
     * Optional ETag.
     */
    readonly eTag?: pulumi.Input<string>;
    /**
     * Identity for the resource.
     */
    readonly identity?: pulumi.Input<inputs.recoveryservices.IdentityData>;
    /**
     * Resource location.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The name of the recovery services vault.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Properties of the vault.
     */
    readonly properties?: pulumi.Input<inputs.recoveryservices.VaultProperties>;
    /**
     * The name of the resource group where the recovery services vault is present.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Identifies the unique system identifier for each Azure resource.
     */
    readonly sku?: pulumi.Input<inputs.recoveryservices.Sku>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
