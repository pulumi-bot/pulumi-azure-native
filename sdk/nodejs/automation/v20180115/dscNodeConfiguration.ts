// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Definition of the dsc node configuration.
 *
 * ## Example Usage
 * ### Create node configuration
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const dscNodeConfiguration = new azurerm.automation.v20180115.DscNodeConfiguration("dscNodeConfiguration", {
 *     automationAccountName: "myAutomationAccount20",
 *     configuration: {
 *         name: "configName",
 *     },
 *     incrementNodeConfigurationBuild: true,
 *     name: "configName.nodeConfigName",
 *     nodeConfigurationName: "configName.nodeConfigName",
 *     resourceGroupName: "rg",
 *     source: {
 *         hash: {
 *             algorithm: "sha256",
 *             value: "6DE256A57F01BFA29B88696D5E77A383D6E61484C7686E8DB955FA10ACE9FFE5",
 *         },
 *         type: "embeddedContent",
 *         value: `
 * instance of MSFT_RoleResource as $MSFT_RoleResource1ref
 * {
 * ResourceID = "[WindowsFeature]IIS";
 *  Ensure = "Present";
 *  SourceInfo = "::3::32::WindowsFeature";
 *  Name = "Web-Server";
 *  ModuleName = "PsDesiredStateConfiguration";
 * 
 * ModuleVersion = "1.0";
 *  ConfigurationName = "configName";
 * };
 * instance of OMI_ConfigurationDocument
 * 
 *                     {
 *  Version="2.0.0";
 *  
 *                         MinimumCompatibleVersion = "1.0.0";
 *  
 *                         CompatibleVersionAdditionalProperties= {"Omi_BaseResource:ConfigurationName"};
 *  
 *                         Author="weijiel";
 *  
 *                         GenerationDate="03/30/2017 13:40:25";
 *  
 *                         GenerationHost="TEST-BACKEND";
 *  
 *                         Name="configName";
 * 
 *                     };
 * `,
 *         version: "1.0",
 *     },
 * });
 *
 * ```
 */
export class DscNodeConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing DscNodeConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DscNodeConfiguration {
        return new DscNodeConfiguration(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:automation/v20180115:DscNodeConfiguration';

    /**
     * Returns true if the given object is an instance of DscNodeConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DscNodeConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DscNodeConfiguration.__pulumiType;
    }

    /**
     * Gets or sets the configuration of the node.
     */
    public readonly configuration!: pulumi.Output<outputs.automation.v20180115.DscConfigurationAssociationPropertyResponse | undefined>;
    /**
     * Gets or sets creation time.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string | undefined>;
    /**
     * If a new build version of NodeConfiguration is required.
     */
    public readonly incrementNodeConfigurationBuild!: pulumi.Output<boolean | undefined>;
    /**
     * Gets or sets the last modified time.
     */
    public /*out*/ readonly lastModifiedTime!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Number of nodes with this node configuration assigned
     */
    public /*out*/ readonly nodeCount!: pulumi.Output<number | undefined>;
    /**
     * Source of node configuration.
     */
    public readonly source!: pulumi.Output<string | undefined>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a DscNodeConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DscNodeConfigurationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.automationAccountName === undefined) {
                throw new Error("Missing required property 'automationAccountName'");
            }
            if (!args || args.configuration === undefined) {
                throw new Error("Missing required property 'configuration'");
            }
            if (!args || args.nodeConfigurationName === undefined) {
                throw new Error("Missing required property 'nodeConfigurationName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.source === undefined) {
                throw new Error("Missing required property 'source'");
            }
            inputs["automationAccountName"] = args ? args.automationAccountName : undefined;
            inputs["configuration"] = args ? args.configuration : undefined;
            inputs["incrementNodeConfigurationBuild"] = args ? args.incrementNodeConfigurationBuild : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["nodeConfigurationName"] = args ? args.nodeConfigurationName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["source"] = args ? args.source : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["creationTime"] = undefined /*out*/;
            inputs["lastModifiedTime"] = undefined /*out*/;
            inputs["nodeCount"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["configuration"] = undefined /*out*/;
            inputs["creationTime"] = undefined /*out*/;
            inputs["incrementNodeConfigurationBuild"] = undefined /*out*/;
            inputs["lastModifiedTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["nodeCount"] = undefined /*out*/;
            inputs["source"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:automation/latest:DscNodeConfiguration" }, { type: "azurerm:automation/v20151031:DscNodeConfiguration" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(DscNodeConfiguration.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DscNodeConfiguration resource.
 */
export interface DscNodeConfigurationArgs {
    /**
     * The name of the automation account.
     */
    readonly automationAccountName: pulumi.Input<string>;
    /**
     * Gets or sets the configuration of the node.
     */
    readonly configuration: pulumi.Input<inputs.automation.v20180115.DscConfigurationAssociationProperty>;
    /**
     * If a new build version of NodeConfiguration is required.
     */
    readonly incrementNodeConfigurationBuild?: pulumi.Input<boolean>;
    /**
     * Name of the node configuration.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The Dsc node configuration name.
     */
    readonly nodeConfigurationName: pulumi.Input<string>;
    /**
     * Name of an Azure Resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Gets or sets the source.
     */
    readonly source: pulumi.Input<inputs.automation.v20180115.ContentSource>;
    /**
     * Gets or sets the tags attached to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
