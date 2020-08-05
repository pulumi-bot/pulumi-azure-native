// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A export resource.
 */
export class Export extends pulumi.CustomResource {
    /**
     * Get an existing Export resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Export {
        return new Export(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:costmanagement/v20190101:Export';

    /**
     * Returns true if the given object is an instance of Export.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Export {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Export.__pulumiType;
    }

    /**
     * Resource name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The properties of the export.
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.costmanagement.v20190101.ExportPropertiesResponse>;
    /**
     * Resource tags.
     */
    public /*out*/ readonly tags!: pulumi.Output<{[key: string]: string}>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Export resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ExportArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ExportArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ExportArgs | undefined;
            if (!args || args.definition === undefined) {
                throw new Error("Missing required property 'definition'");
            }
            if (!args || args.deliveryInfo === undefined) {
                throw new Error("Missing required property 'deliveryInfo'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.scope === undefined) {
                throw new Error("Missing required property 'scope'");
            }
            inputs["definition"] = args ? args.definition : undefined;
            inputs["deliveryInfo"] = args ? args.deliveryInfo : undefined;
            inputs["format"] = args ? args.format : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["schedule"] = args ? args.schedule : undefined;
            inputs["scope"] = args ? args.scope : undefined;
            inputs["properties"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Export.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Export resource.
 */
export interface ExportArgs {
    /**
     * Has definition for the export.
     */
    readonly definition: pulumi.Input<inputs.costmanagement.v20190101.QueryDefinition>;
    /**
     * Has delivery information for the export.
     */
    readonly deliveryInfo: pulumi.Input<inputs.costmanagement.v20190101.ExportDeliveryInfo>;
    /**
     * The format of the export being delivered.
     */
    readonly format?: pulumi.Input<string>;
    /**
     * Export Name.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Has schedule information for the export.
     */
    readonly schedule?: pulumi.Input<inputs.costmanagement.v20190101.ExportSchedule>;
    /**
     * The scope associated with export operations. This includes '/subscriptions/{subscriptionId}' for subscription scope, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resourceGroup scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope and '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount scope.
     */
    readonly scope: pulumi.Input<string>;
}
