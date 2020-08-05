// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Certificate purchase order
 */
export class CertificateOrder extends pulumi.CustomResource {
    /**
     * Get an existing CertificateOrder resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CertificateOrder {
        return new CertificateOrder(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:certificateregistration/v20150801:CertificateOrder';

    /**
     * Returns true if the given object is an instance of CertificateOrder.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CertificateOrder {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CertificateOrder.__pulumiType;
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
    public /*out*/ readonly properties!: pulumi.Output<outputs.certificateregistration.v20150801.CertificateOrderResponseProperties>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a CertificateOrder resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CertificateOrderArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CertificateOrderArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as CertificateOrderArgs | undefined;
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["autoRenew"] = args ? args.autoRenew : undefined;
            inputs["certificates"] = args ? args.certificates : undefined;
            inputs["csr"] = args ? args.csr : undefined;
            inputs["distinguishedName"] = args ? args.distinguishedName : undefined;
            inputs["domainVerificationToken"] = args ? args.domainVerificationToken : undefined;
            inputs["expirationTime"] = args ? args.expirationTime : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["intermediate"] = args ? args.intermediate : undefined;
            inputs["keySize"] = args ? args.keySize : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["lastCertificateIssuanceTime"] = args ? args.lastCertificateIssuanceTime : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["productType"] = args ? args.productType : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["root"] = args ? args.root : undefined;
            inputs["serialNumber"] = args ? args.serialNumber : undefined;
            inputs["signedCertificate"] = args ? args.signedCertificate : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["validityInYears"] = args ? args.validityInYears : undefined;
            inputs["properties"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(CertificateOrder.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a CertificateOrder resource.
 */
export interface CertificateOrderArgs {
    /**
     * Auto renew
     */
    readonly autoRenew?: pulumi.Input<boolean>;
    /**
     * State of the Key Vault secret
     */
    readonly certificates?: pulumi.Input<{[key: string]: pulumi.Input<inputs.certificateregistration.v20150801.CertificateOrderCertificate>}>;
    /**
     * Last CSR that was created for this order
     */
    readonly csr?: pulumi.Input<string>;
    /**
     * Certificate distinguished name
     */
    readonly distinguishedName?: pulumi.Input<string>;
    /**
     * Domain Verification Token
     */
    readonly domainVerificationToken?: pulumi.Input<string>;
    /**
     * Certificate expiration time
     */
    readonly expirationTime?: pulumi.Input<string>;
    /**
     * Resource Id
     */
    readonly id?: pulumi.Input<string>;
    /**
     * Intermediate certificate
     */
    readonly intermediate?: pulumi.Input<inputs.certificateregistration.v20150801.CertificateDetails>;
    /**
     * Certificate Key Size
     */
    readonly keySize?: pulumi.Input<number>;
    /**
     * Kind of resource
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Certificate last issuance time
     */
    readonly lastCertificateIssuanceTime?: pulumi.Input<string>;
    /**
     * Resource Location
     */
    readonly location: pulumi.Input<string>;
    /**
     * Resource Name
     */
    readonly name: pulumi.Input<string>;
    /**
     * Certificate product type
     */
    readonly productType?: pulumi.Input<string>;
    /**
     * Status of certificate order
     */
    readonly provisioningState?: pulumi.Input<string>;
    /**
     * Azure resource group name
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Root certificate
     */
    readonly root?: pulumi.Input<inputs.certificateregistration.v20150801.CertificateDetails>;
    /**
     * Current serial number of the certificate
     */
    readonly serialNumber?: pulumi.Input<string>;
    /**
     * Signed certificate
     */
    readonly signedCertificate?: pulumi.Input<inputs.certificateregistration.v20150801.CertificateDetails>;
    /**
     * Current order status
     */
    readonly status?: pulumi.Input<string>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Resource type
     */
    readonly type?: pulumi.Input<string>;
    /**
     * Duration in years (must be between 1 and 3)
     */
    readonly validityInYears?: pulumi.Input<number>;
}
