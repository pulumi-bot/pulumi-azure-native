// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Class representing the Key Vault container for certificate purchased through Azure
 */
export class CertificateOrderCertificate extends pulumi.CustomResource {
    /**
     * Get an existing CertificateOrderCertificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CertificateOrderCertificate {
        return new CertificateOrderCertificate(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:certificateregistration/v20150801:CertificateOrderCertificate';

    /**
     * Returns true if the given object is an instance of CertificateOrderCertificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CertificateOrderCertificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CertificateOrderCertificate.__pulumiType;
    }

    /**
     * Kind of resource
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * Resource Location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource Name
     */
    public readonly name!: pulumi.Output<string | undefined>;
    public /*out*/ readonly properties!: pulumi.Output<outputs.certificateregistration.v20150801.CertificateOrderCertificateResponseProperties>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a CertificateOrderCertificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CertificateOrderCertificateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CertificateOrderCertificateArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as CertificateOrderCertificateArgs | undefined;
            if (!args || args.certificateOrderName === undefined) {
                throw new Error("Missing required property 'certificateOrderName'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["certificateOrderName"] = args ? args.certificateOrderName : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["keyVaultId"] = args ? args.keyVaultId : undefined;
            inputs["keyVaultSecretName"] = args ? args.keyVaultSecretName : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["properties"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(CertificateOrderCertificate.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a CertificateOrderCertificate resource.
 */
export interface CertificateOrderCertificateArgs {
    /**
     * Certificate name
     */
    readonly certificateOrderName: pulumi.Input<string>;
    /**
     * Resource Id
     */
    readonly id?: pulumi.Input<string>;
    /**
     * Key Vault Csm resource Id
     */
    readonly keyVaultId?: pulumi.Input<string>;
    /**
     * Key Vault secret name
     */
    readonly keyVaultSecretName?: pulumi.Input<string>;
    /**
     * Kind of resource
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Resource Location
     */
    readonly location: pulumi.Input<string>;
    /**
     * Resource Name
     */
    readonly name: pulumi.Input<string>;
    /**
     * Status of the Key Vault secret
     */
    readonly provisioningState?: pulumi.Input<string>;
    /**
     * Azure resource group name
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Resource type
     */
    readonly type?: pulumi.Input<string>;
}
