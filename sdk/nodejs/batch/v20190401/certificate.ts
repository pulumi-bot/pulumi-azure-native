// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Contains information about a certificate.
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
        return new Certificate(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:batch/v20190401:Certificate';

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
     * This is only returned when the certificate provisioningState is 'Failed'.
     */
    public /*out*/ readonly deleteCertificateError!: pulumi.Output<outputs.batch.v20190401.DeleteCertificateErrorResponse>;
    /**
     * The ETag of the resource, used for concurrency statements.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The format of the certificate - either Pfx or Cer. If omitted, the default is Pfx.
     */
    public readonly format!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The previous provisioned state of the resource
     */
    public /*out*/ readonly previousProvisioningState!: pulumi.Output<string>;
    public /*out*/ readonly previousProvisioningStateTransitionTime!: pulumi.Output<string>;
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    public /*out*/ readonly provisioningStateTransitionTime!: pulumi.Output<string>;
    /**
     * The public key of the certificate.
     */
    public /*out*/ readonly publicData!: pulumi.Output<string>;
    /**
     * This must match the thumbprint from the name.
     */
    public readonly thumbprint!: pulumi.Output<string | undefined>;
    /**
     * This must match the first portion of the certificate name. Currently required to be 'SHA1'.
     */
    public readonly thumbprintAlgorithm!: pulumi.Output<string | undefined>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Certificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CertificateArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.accountName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'accountName'");
            }
            if ((!args || args.certificateName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'certificateName'");
            }
            if ((!args || args.data === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'data'");
            }
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["accountName"] = args ? args.accountName : undefined;
            inputs["certificateName"] = args ? args.certificateName : undefined;
            inputs["data"] = args ? args.data : undefined;
            inputs["format"] = args ? args.format : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["thumbprint"] = args ? args.thumbprint : undefined;
            inputs["thumbprintAlgorithm"] = args ? args.thumbprintAlgorithm : undefined;
            inputs["deleteCertificateError"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["previousProvisioningState"] = undefined /*out*/;
            inputs["previousProvisioningStateTransitionTime"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["provisioningStateTransitionTime"] = undefined /*out*/;
            inputs["publicData"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["deleteCertificateError"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["format"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["previousProvisioningState"] = undefined /*out*/;
            inputs["previousProvisioningStateTransitionTime"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["provisioningStateTransitionTime"] = undefined /*out*/;
            inputs["publicData"] = undefined /*out*/;
            inputs["thumbprint"] = undefined /*out*/;
            inputs["thumbprintAlgorithm"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:batch/latest:Certificate" }, { type: "azure-nextgen:batch/v20170901:Certificate" }, { type: "azure-nextgen:batch/v20181201:Certificate" }, { type: "azure-nextgen:batch/v20190801:Certificate" }, { type: "azure-nextgen:batch/v20200301:Certificate" }, { type: "azure-nextgen:batch/v20200501:Certificate" }, { type: "azure-nextgen:batch/v20200901:Certificate" }, { type: "azure-nextgen:batch/v20210101:Certificate" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Certificate.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Certificate resource.
 */
export interface CertificateArgs {
    /**
     * The name of the Batch account.
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * The identifier for the certificate. This must be made up of algorithm and thumbprint separated by a dash, and must match the certificate data in the request. For example SHA1-a3d1c5.
     */
    readonly certificateName: pulumi.Input<string>;
    /**
     * The maximum size is 10KB.
     */
    readonly data: pulumi.Input<string>;
    /**
     * The format of the certificate - either Pfx or Cer. If omitted, the default is Pfx.
     */
    readonly format?: pulumi.Input<enums.batch.v20190401.CertificateFormat>;
    /**
     * This is required if the certificate format is pfx and must be omitted if the certificate format is cer.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * The name of the resource group that contains the Batch account.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * This must match the thumbprint from the name.
     */
    readonly thumbprint?: pulumi.Input<string>;
    /**
     * This must match the first portion of the certificate name. Currently required to be 'SHA1'.
     */
    readonly thumbprintAlgorithm?: pulumi.Input<string>;
}
