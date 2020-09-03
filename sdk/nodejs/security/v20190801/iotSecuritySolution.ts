// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * IoT Security solution configuration and resource information.
 *
 * ## Example Usage
 * ### Create or update a IoT security solution
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const iotSecuritySolution = new azurerm.security.v20190801.IotSecuritySolution("iotSecuritySolution", {
 *     disabledDataSources: [],
 *     displayName: "Solution Default",
 *     "export": [],
 *     iotHubs: ["/subscriptions/075423e9-7d33-4166-8bdf-3920b04e3735/resourceGroups/myRg/providers/Microsoft.Devices/IotHubs/FirstIotHub"],
 *     location: "East Us",
 *     recommendationsConfiguration: [
 *         {
 *             recommendationType: "IoT_OpenPorts",
 *             status: "Disabled",
 *         },
 *         {
 *             recommendationType: "IoT_SharedCredentials",
 *             status: "Disabled",
 *         },
 *     ],
 *     resourceGroupName: "MyGroup",
 *     solutionName: "default",
 *     status: "Enabled",
 *     tags: {},
 *     unmaskedIpLoggingStatus: "Enabled",
 *     userDefinedResources: {
 *         query: "where type != \"microsoft.devices/iothubs\" | where name contains \"iot\"",
 *         querySubscriptions: ["075423e9-7d33-4166-8bdf-3920b04e3735"],
 *     },
 *     workspace: "/subscriptions/c4930e90-cd72-4aa5-93e9-2d081d129569/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace1",
 * });
 *
 * ```
 */
export class IotSecuritySolution extends pulumi.CustomResource {
    /**
     * Get an existing IotSecuritySolution resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): IotSecuritySolution {
        return new IotSecuritySolution(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:security/v20190801:IotSecuritySolution';

    /**
     * Returns true if the given object is an instance of IotSecuritySolution.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IotSecuritySolution {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IotSecuritySolution.__pulumiType;
    }

    /**
     * List of resources that were automatically discovered as relevant to the security solution.
     */
    public /*out*/ readonly autoDiscoveredResources!: pulumi.Output<string[]>;
    /**
     * Disabled data sources. Disabling these data sources compromises the system.
     */
    public readonly disabledDataSources!: pulumi.Output<string[] | undefined>;
    /**
     * Resource display name.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * List of additional options for exporting to workspace data.
     */
    public readonly export!: pulumi.Output<string[] | undefined>;
    /**
     * IoT Hub resource IDs
     */
    public readonly iotHubs!: pulumi.Output<string[]>;
    /**
     * The resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * List of the configuration status for each recommendation type.
     */
    public readonly recommendationsConfiguration!: pulumi.Output<outputs.security.v20190801.RecommendationConfigurationPropertiesResponse[] | undefined>;
    /**
     * Status of the IoT Security solution.
     */
    public readonly status!: pulumi.Output<string | undefined>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Unmasked IP address logging status
     */
    public readonly unmaskedIpLoggingStatus!: pulumi.Output<string | undefined>;
    /**
     * Properties of the IoT Security solution's user defined resources.
     */
    public readonly userDefinedResources!: pulumi.Output<outputs.security.v20190801.UserDefinedResourcesPropertiesResponse | undefined>;
    /**
     * Workspace resource ID
     */
    public readonly workspace!: pulumi.Output<string | undefined>;

    /**
     * Create a IotSecuritySolution resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IotSecuritySolutionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IotSecuritySolutionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as IotSecuritySolutionArgs | undefined;
            if (!args || args.displayName === undefined) {
                throw new Error("Missing required property 'displayName'");
            }
            if (!args || args.iotHubs === undefined) {
                throw new Error("Missing required property 'iotHubs'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.solutionName === undefined) {
                throw new Error("Missing required property 'solutionName'");
            }
            inputs["disabledDataSources"] = args ? args.disabledDataSources : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["export"] = args ? args.export : undefined;
            inputs["iotHubs"] = args ? args.iotHubs : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["recommendationsConfiguration"] = args ? args.recommendationsConfiguration : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["solutionName"] = args ? args.solutionName : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["unmaskedIpLoggingStatus"] = args ? args.unmaskedIpLoggingStatus : undefined;
            inputs["userDefinedResources"] = args ? args.userDefinedResources : undefined;
            inputs["workspace"] = args ? args.workspace : undefined;
            inputs["autoDiscoveredResources"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:security/latest:IotSecuritySolution" }, { type: "azurerm:security/v20170801preview:IotSecuritySolution" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(IotSecuritySolution.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a IotSecuritySolution resource.
 */
export interface IotSecuritySolutionArgs {
    /**
     * Disabled data sources. Disabling these data sources compromises the system.
     */
    readonly disabledDataSources?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Resource display name.
     */
    readonly displayName: pulumi.Input<string>;
    /**
     * List of additional options for exporting to workspace data.
     */
    readonly export?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * IoT Hub resource IDs
     */
    readonly iotHubs: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * List of the configuration status for each recommendation type.
     */
    readonly recommendationsConfiguration?: pulumi.Input<pulumi.Input<inputs.security.v20190801.RecommendationConfigurationProperties>[]>;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the IoT Security solution.
     */
    readonly solutionName: pulumi.Input<string>;
    /**
     * Status of the IoT Security solution.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Unmasked IP address logging status
     */
    readonly unmaskedIpLoggingStatus?: pulumi.Input<string>;
    /**
     * Properties of the IoT Security solution's user defined resources.
     */
    readonly userDefinedResources?: pulumi.Input<inputs.security.v20190801.UserDefinedResourcesProperties>;
    /**
     * Workspace resource ID
     */
    readonly workspace?: pulumi.Input<string>;
}
