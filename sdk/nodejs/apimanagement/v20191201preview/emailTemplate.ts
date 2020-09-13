// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Email Template details.
 */
export class EmailTemplate extends pulumi.CustomResource {
    /**
     * Get an existing EmailTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): EmailTemplate {
        return new EmailTemplate(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:apimanagement/v20191201preview:EmailTemplate';

    /**
     * Returns true if the given object is an instance of EmailTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EmailTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EmailTemplate.__pulumiType;
    }

    /**
     * Email Template Body. This should be a valid XDocument
     */
    public readonly body!: pulumi.Output<string>;
    /**
     * Description of the Email Template.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Whether the template is the default template provided by Api Management or has been edited.
     */
    public /*out*/ readonly isDefault!: pulumi.Output<boolean>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Email Template Parameter values.
     */
    public readonly parameters!: pulumi.Output<outputs.apimanagement.v20191201preview.EmailTemplateParametersContractPropertiesResponse[] | undefined>;
    /**
     * Subject of the Template.
     */
    public readonly subject!: pulumi.Output<string>;
    /**
     * Title of the Template.
     */
    public readonly title!: pulumi.Output<string | undefined>;
    /**
     * Resource type for API Management resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a EmailTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EmailTemplateArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serviceName === undefined) {
                throw new Error("Missing required property 'serviceName'");
            }
            if (!args || args.templateName === undefined) {
                throw new Error("Missing required property 'templateName'");
            }
            inputs["body"] = args ? args.body : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["subject"] = args ? args.subject : undefined;
            inputs["templateName"] = args ? args.templateName : undefined;
            inputs["title"] = args ? args.title : undefined;
            inputs["isDefault"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["body"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["isDefault"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["parameters"] = undefined /*out*/;
            inputs["subject"] = undefined /*out*/;
            inputs["title"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:apimanagement/latest:EmailTemplate" }, { type: "azurerm:apimanagement/v20170301:EmailTemplate" }, { type: "azurerm:apimanagement/v20180101:EmailTemplate" }, { type: "azurerm:apimanagement/v20180601preview:EmailTemplate" }, { type: "azurerm:apimanagement/v20190101:EmailTemplate" }, { type: "azurerm:apimanagement/v20191201:EmailTemplate" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(EmailTemplate.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a EmailTemplate resource.
 */
export interface EmailTemplateArgs {
    /**
     * Email Template Body. This should be a valid XDocument
     */
    readonly body?: pulumi.Input<string>;
    /**
     * Description of the Email Template.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Email Template Parameter values.
     */
    readonly parameters?: pulumi.Input<pulumi.Input<inputs.apimanagement.v20191201preview.EmailTemplateParametersContractProperties>[]>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the API Management service.
     */
    readonly serviceName: pulumi.Input<string>;
    /**
     * Subject of the Template.
     */
    readonly subject?: pulumi.Input<string>;
    /**
     * Email Template Name Identifier.
     */
    readonly templateName: pulumi.Input<string>;
    /**
     * Title of the Template.
     */
    readonly title?: pulumi.Input<string>;
}
