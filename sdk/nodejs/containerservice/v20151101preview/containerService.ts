// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Container service
 */
export class ContainerService extends pulumi.CustomResource {
    /**
     * Get an existing ContainerService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ContainerService {
        return new ContainerService(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:containerservice/v20151101preview:ContainerService';

    /**
     * Returns true if the given object is an instance of ContainerService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ContainerService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ContainerService.__pulumiType;
    }

    /**
     * Properties of agent pools
     */
    public readonly agentPoolProfiles!: pulumi.Output<outputs.containerservice.v20151101preview.ContainerServiceAgentPoolProfileResponse[]>;
    /**
     * Properties for Diagnostic Agent
     */
    public readonly diagnosticsProfile!: pulumi.Output<outputs.containerservice.v20151101preview.ContainerServiceDiagnosticsProfileResponse | undefined>;
    /**
     * Properties for Linux VMs
     */
    public readonly linuxProfile!: pulumi.Output<outputs.containerservice.v20151101preview.ContainerServiceLinuxProfileResponse>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Properties of master agents
     */
    public readonly masterProfile!: pulumi.Output<outputs.containerservice.v20151101preview.ContainerServiceMasterProfileResponse>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Properties of orchestrator
     */
    public readonly orchestratorProfile!: pulumi.Output<outputs.containerservice.v20151101preview.ContainerServiceOrchestratorProfileResponse | undefined>;
    /**
     * Gets the provisioning state, which only appears in the response.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Properties of Windows VMs
     */
    public readonly windowsProfile!: pulumi.Output<outputs.containerservice.v20151101preview.ContainerServiceWindowsProfileResponse | undefined>;

    /**
     * Create a ContainerService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ContainerServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ContainerServiceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ContainerServiceArgs | undefined;
            if (!args || args.agentPoolProfiles === undefined) {
                throw new Error("Missing required property 'agentPoolProfiles'");
            }
            if (!args || args.containerServiceName === undefined) {
                throw new Error("Missing required property 'containerServiceName'");
            }
            if (!args || args.linuxProfile === undefined) {
                throw new Error("Missing required property 'linuxProfile'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.masterProfile === undefined) {
                throw new Error("Missing required property 'masterProfile'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["agentPoolProfiles"] = args ? args.agentPoolProfiles : undefined;
            inputs["containerServiceName"] = args ? args.containerServiceName : undefined;
            inputs["diagnosticsProfile"] = args ? args.diagnosticsProfile : undefined;
            inputs["linuxProfile"] = args ? args.linuxProfile : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["masterProfile"] = args ? args.masterProfile : undefined;
            inputs["orchestratorProfile"] = args ? args.orchestratorProfile : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["windowsProfile"] = args ? args.windowsProfile : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:containerservice/latest:ContainerService" }, { type: "azurerm:containerservice/v20160330:ContainerService" }, { type: "azurerm:containerservice/v20160930:ContainerService" }, { type: "azurerm:containerservice/v20170131:ContainerService" }, { type: "azurerm:containerservice/v20170701:ContainerService" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ContainerService.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ContainerService resource.
 */
export interface ContainerServiceArgs {
    /**
     * Properties of agent pools
     */
    readonly agentPoolProfiles: pulumi.Input<pulumi.Input<inputs.containerservice.v20151101preview.ContainerServiceAgentPoolProfile>[]>;
    /**
     * The name of the container service within the given subscription and resource group.
     */
    readonly containerServiceName: pulumi.Input<string>;
    /**
     * Properties for Diagnostic Agent
     */
    readonly diagnosticsProfile?: pulumi.Input<inputs.containerservice.v20151101preview.ContainerServiceDiagnosticsProfile>;
    /**
     * Properties for Linux VMs
     */
    readonly linuxProfile: pulumi.Input<inputs.containerservice.v20151101preview.ContainerServiceLinuxProfile>;
    /**
     * Resource location
     */
    readonly location: pulumi.Input<string>;
    /**
     * Properties of master agents
     */
    readonly masterProfile: pulumi.Input<inputs.containerservice.v20151101preview.ContainerServiceMasterProfile>;
    /**
     * Properties of orchestrator
     */
    readonly orchestratorProfile?: pulumi.Input<inputs.containerservice.v20151101preview.ContainerServiceOrchestratorProfile>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Properties of Windows VMs
     */
    readonly windowsProfile?: pulumi.Input<inputs.containerservice.v20151101preview.ContainerServiceWindowsProfile>;
}
