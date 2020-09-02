// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Container service.
 *
 * ## Example Usage
 * ### Create/Update Container Service
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const containerService = new azurerm.containerservice.v20170131.ContainerService("containerService", {
 *     containerServiceName: "acs1",
 *     location: "location1",
 *     resourceGroupName: "rg1",
 * });
 *
 * ```
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
    public static readonly __pulumiType = 'azurerm:containerservice/v20170131:ContainerService';

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
     * Properties of the agent pool.
     */
    public readonly agentPoolProfiles!: pulumi.Output<outputs.containerservice.v20170131.ContainerServiceAgentPoolProfileResponse[]>;
    /**
     * Properties for custom clusters.
     */
    public readonly customProfile!: pulumi.Output<outputs.containerservice.v20170131.ContainerServiceCustomProfileResponse | undefined>;
    /**
     * Properties of the diagnostic agent.
     */
    public readonly diagnosticsProfile!: pulumi.Output<outputs.containerservice.v20170131.ContainerServiceDiagnosticsProfileResponse | undefined>;
    /**
     * Properties of Linux VMs.
     */
    public readonly linuxProfile!: pulumi.Output<outputs.containerservice.v20170131.ContainerServiceLinuxProfileResponse>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Properties of master agents.
     */
    public readonly masterProfile!: pulumi.Output<outputs.containerservice.v20170131.ContainerServiceMasterProfileResponse>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Properties of the orchestrator.
     */
    public readonly orchestratorProfile!: pulumi.Output<outputs.containerservice.v20170131.ContainerServiceOrchestratorProfileResponse | undefined>;
    /**
     * the current deployment or provisioning state, which only appears in the response.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Properties for cluster service principals.
     */
    public readonly servicePrincipalProfile!: pulumi.Output<outputs.containerservice.v20170131.ContainerServiceServicePrincipalProfileResponse | undefined>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Properties of Windows VMs.
     */
    public readonly windowsProfile!: pulumi.Output<outputs.containerservice.v20170131.ContainerServiceWindowsProfileResponse | undefined>;

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
            inputs["customProfile"] = args ? args.customProfile : undefined;
            inputs["diagnosticsProfile"] = args ? args.diagnosticsProfile : undefined;
            inputs["linuxProfile"] = args ? args.linuxProfile : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["masterProfile"] = args ? args.masterProfile : undefined;
            inputs["orchestratorProfile"] = args ? args.orchestratorProfile : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["servicePrincipalProfile"] = args ? args.servicePrincipalProfile : undefined;
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
        const aliasOpts = { aliases: [{ type: "azurerm:containerservice/latest:ContainerService" }, { type: "azurerm:containerservice/v20160330:ContainerService" }, { type: "azurerm:containerservice/v20160930:ContainerService" }, { type: "azurerm:containerservice/v20170701:ContainerService" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ContainerService.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ContainerService resource.
 */
export interface ContainerServiceArgs {
    /**
     * Properties of the agent pool.
     */
    readonly agentPoolProfiles: pulumi.Input<pulumi.Input<inputs.containerservice.v20170131.ContainerServiceAgentPoolProfile>[]>;
    /**
     * The name of the container service in the specified subscription and resource group.
     */
    readonly containerServiceName: pulumi.Input<string>;
    /**
     * Properties for custom clusters.
     */
    readonly customProfile?: pulumi.Input<inputs.containerservice.v20170131.ContainerServiceCustomProfile>;
    /**
     * Properties of the diagnostic agent.
     */
    readonly diagnosticsProfile?: pulumi.Input<inputs.containerservice.v20170131.ContainerServiceDiagnosticsProfile>;
    /**
     * Properties of Linux VMs.
     */
    readonly linuxProfile: pulumi.Input<inputs.containerservice.v20170131.ContainerServiceLinuxProfile>;
    /**
     * Resource location
     */
    readonly location: pulumi.Input<string>;
    /**
     * Properties of master agents.
     */
    readonly masterProfile: pulumi.Input<inputs.containerservice.v20170131.ContainerServiceMasterProfile>;
    /**
     * Properties of the orchestrator.
     */
    readonly orchestratorProfile?: pulumi.Input<inputs.containerservice.v20170131.ContainerServiceOrchestratorProfile>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Properties for cluster service principals.
     */
    readonly servicePrincipalProfile?: pulumi.Input<inputs.containerservice.v20170131.ContainerServiceServicePrincipalProfile>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Properties of Windows VMs.
     */
    readonly windowsProfile?: pulumi.Input<inputs.containerservice.v20170131.ContainerServiceWindowsProfile>;
}
