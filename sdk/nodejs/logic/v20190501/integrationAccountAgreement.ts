// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The integration account agreement.
 */
export class IntegrationAccountAgreement extends pulumi.CustomResource {
    /**
     * Get an existing IntegrationAccountAgreement resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): IntegrationAccountAgreement {
        return new IntegrationAccountAgreement(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:logic/v20190501:IntegrationAccountAgreement';

    /**
     * Returns true if the given object is an instance of IntegrationAccountAgreement.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IntegrationAccountAgreement {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IntegrationAccountAgreement.__pulumiType;
    }

    /**
     * The resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Gets the resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The integration account agreement properties.
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.logic.v20190501.IntegrationAccountAgreementPropertiesResponse>;
    /**
     * The resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Gets the resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a IntegrationAccountAgreement resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IntegrationAccountAgreementArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IntegrationAccountAgreementArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as IntegrationAccountAgreementArgs | undefined;
            if (!args || args.agreementType === undefined) {
                throw new Error("Missing required property 'agreementType'");
            }
            if (!args || args.content === undefined) {
                throw new Error("Missing required property 'content'");
            }
            if (!args || args.guestIdentity === undefined) {
                throw new Error("Missing required property 'guestIdentity'");
            }
            if (!args || args.guestPartner === undefined) {
                throw new Error("Missing required property 'guestPartner'");
            }
            if (!args || args.hostIdentity === undefined) {
                throw new Error("Missing required property 'hostIdentity'");
            }
            if (!args || args.hostPartner === undefined) {
                throw new Error("Missing required property 'hostPartner'");
            }
            if (!args || args.integrationAccountName === undefined) {
                throw new Error("Missing required property 'integrationAccountName'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["agreementType"] = args ? args.agreementType : undefined;
            inputs["content"] = args ? args.content : undefined;
            inputs["guestIdentity"] = args ? args.guestIdentity : undefined;
            inputs["guestPartner"] = args ? args.guestPartner : undefined;
            inputs["hostIdentity"] = args ? args.hostIdentity : undefined;
            inputs["hostPartner"] = args ? args.hostPartner : undefined;
            inputs["integrationAccountName"] = args ? args.integrationAccountName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["name"] = args ? args.name : undefined;
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
        super(IntegrationAccountAgreement.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a IntegrationAccountAgreement resource.
 */
export interface IntegrationAccountAgreementArgs {
    /**
     * The agreement type.
     */
    readonly agreementType: pulumi.Input<string>;
    /**
     * The agreement content.
     */
    readonly content: pulumi.Input<inputs.logic.v20190501.AgreementContent>;
    /**
     * The business identity of the guest partner.
     */
    readonly guestIdentity: pulumi.Input<inputs.logic.v20190501.BusinessIdentity>;
    /**
     * The integration account partner that is set as guest partner for this agreement.
     */
    readonly guestPartner: pulumi.Input<string>;
    /**
     * The business identity of the host partner.
     */
    readonly hostIdentity: pulumi.Input<inputs.logic.v20190501.BusinessIdentity>;
    /**
     * The integration account partner that is set as host partner for this agreement.
     */
    readonly hostPartner: pulumi.Input<string>;
    /**
     * The integration account name.
     */
    readonly integrationAccountName: pulumi.Input<string>;
    /**
     * The resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The metadata.
     */
    readonly metadata?: pulumi.Input<{[key: string]: any}>;
    /**
     * The integration account agreement name.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
