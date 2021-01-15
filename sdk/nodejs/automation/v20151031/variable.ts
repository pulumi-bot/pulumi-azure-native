// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Definition of the variable.
 */
export class Variable extends pulumi.CustomResource {
    /**
     * Get an existing Variable resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Variable {
        return new Variable(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:automation/v20151031:Variable';

    /**
     * Returns true if the given object is an instance of Variable.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Variable {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Variable.__pulumiType;
    }

    /**
     * Gets or sets the creation time.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string | undefined>;
    /**
     * Gets or sets the description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Gets or sets the encrypted flag of the variable.
     */
    public readonly isEncrypted!: pulumi.Output<boolean | undefined>;
    /**
     * Gets or sets the last modified time.
     */
    public /*out*/ readonly lastModifiedTime!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Gets or sets the value of the variable.
     */
    public readonly value!: pulumi.Output<string | undefined>;

    /**
     * Create a Variable resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VariableArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if ((!args || args.automationAccountName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'automationAccountName'");
            }
            if ((!args || args.name === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.resourceGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.variableName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'variableName'");
            }
            inputs["automationAccountName"] = args ? args.automationAccountName : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["isEncrypted"] = args ? args.isEncrypted : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["value"] = args ? args.value : undefined;
            inputs["variableName"] = args ? args.variableName : undefined;
            inputs["creationTime"] = undefined /*out*/;
            inputs["lastModifiedTime"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["creationTime"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["isEncrypted"] = undefined /*out*/;
            inputs["lastModifiedTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["value"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:automation/latest:Variable" }, { type: "azure-nextgen:automation/v20190601:Variable" }, { type: "azure-nextgen:automation/v20200113preview:Variable" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Variable.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Variable resource.
 */
export interface VariableArgs {
    /**
     * The name of the automation account.
     */
    readonly automationAccountName: pulumi.Input<string>;
    /**
     * Gets or sets the description of the variable.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Gets or sets the encrypted flag of the variable.
     */
    readonly isEncrypted?: pulumi.Input<boolean>;
    /**
     * Gets or sets the name of the variable.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Name of an Azure Resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Gets or sets the value of the variable.
     */
    readonly value?: pulumi.Input<string>;
    /**
     * The variable name.
     */
    readonly variableName: pulumi.Input<string>;
}
