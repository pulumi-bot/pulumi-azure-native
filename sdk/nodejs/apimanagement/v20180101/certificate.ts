// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Certificate details.
 *
 * ## Example Usage
 * ### ApiManagementCreateCertificate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const certificate = new azurerm.apimanagement.v20180101.Certificate("certificate", {
 *     certificateId: "tempcert",
 *     data: "****************Base 64 Encoded Certificate *******************************",
 *     password: "****Certificate Password******",
 *     resourceGroupName: "rg1",
 *     serviceName: "apimService1",
 * });
 *
 * ```
 */
export class Certificate extends pulumi.CustomResource {
    /**
     * Get an existing Certificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Certificate {
        return new Certificate(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:apimanagement/v20180101:Certificate';

    /**
     * Returns true if the given object is an instance of Certificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Certificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Certificate.__pulumiType;
    }

    /**
     * Expiration date of the certificate. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
     */
    public /*out*/ readonly expirationDate!: pulumi.Output<string>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Subject attribute of the certificate.
     */
    public /*out*/ readonly subject!: pulumi.Output<string>;
    /**
     * Thumbprint of the certificate.
     */
    public /*out*/ readonly thumbprint!: pulumi.Output<string>;
    /**
     * Resource type for API Management resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Certificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CertificateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CertificateArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as CertificateArgs | undefined;
            if (!args || args.certificateId === undefined) {
                throw new Error("Missing required property 'certificateId'");
            }
            if (!args || args.data === undefined) {
                throw new Error("Missing required property 'data'");
            }
            if (!args || args.password === undefined) {
                throw new Error("Missing required property 'password'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serviceName === undefined) {
                throw new Error("Missing required property 'serviceName'");
            }
            inputs["certificateId"] = args ? args.certificateId : undefined;
            inputs["data"] = args ? args.data : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["expirationDate"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["subject"] = undefined /*out*/;
            inputs["thumbprint"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:apimanagement/latest:Certificate" }, { type: "azurerm:apimanagement/v20160707:Certificate" }, { type: "azurerm:apimanagement/v20161010:Certificate" }, { type: "azurerm:apimanagement/v20170301:Certificate" }, { type: "azurerm:apimanagement/v20190101:Certificate" }, { type: "azurerm:apimanagement/v20191201:Certificate" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Certificate.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Certificate resource.
 */
export interface CertificateArgs {
    /**
     * Identifier of the certificate entity. Must be unique in the current API Management service instance.
     */
    readonly certificateId: pulumi.Input<string>;
    /**
     * Base 64 encoded certificate using the application/x-pkcs12 representation.
     */
    readonly data: pulumi.Input<string>;
    /**
     * Password for the Certificate
     */
    readonly password: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the API Management service.
     */
    readonly serviceName: pulumi.Input<string>;
}
