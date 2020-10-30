// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Represents a HostPool definition.
 */
export class HostPool extends pulumi.CustomResource {
    /**
     * Get an existing HostPool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): HostPool {
        return new HostPool(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:desktopvirtualization/v20201019preview:HostPool';

    /**
     * Returns true if the given object is an instance of HostPool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HostPool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HostPool.__pulumiType;
    }

    /**
     * List of applicationGroup links.
     */
    public /*out*/ readonly applicationGroupReferences!: pulumi.Output<string[]>;
    /**
     * Custom rdp property of HostPool.
     */
    public readonly customRdpProperty!: pulumi.Output<string | undefined>;
    /**
     * Description of HostPool.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Friendly name of HostPool.
     */
    public readonly friendlyName!: pulumi.Output<string | undefined>;
    /**
     * HostPool type for desktop.
     */
    public readonly hostPoolType!: pulumi.Output<string>;
    /**
     * The type of the load balancer.
     */
    public readonly loadBalancerType!: pulumi.Output<string>;
    /**
     * The geo-location where the resource lives
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The max session limit of HostPool.
     */
    public readonly maxSessionLimit!: pulumi.Output<number | undefined>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * PersonalDesktopAssignment type for HostPool.
     */
    public readonly personalDesktopAssignmentType!: pulumi.Output<string | undefined>;
    /**
     * The type of preferred application group type, default to Desktop Application Group
     */
    public readonly preferredAppGroupType!: pulumi.Output<string>;
    /**
     * The registration info of HostPool.
     */
    public readonly registrationInfo!: pulumi.Output<outputs.desktopvirtualization.v20201019preview.RegistrationInfoResponse | undefined>;
    /**
     * The ring number of HostPool.
     */
    public readonly ring!: pulumi.Output<number | undefined>;
    /**
     * ClientId for the registered Relying Party used to issue WVD SSO certificates.
     */
    public readonly ssoClientId!: pulumi.Output<string | undefined>;
    /**
     * Path to Azure KeyVault storing the secret used for communication to ADFS.
     */
    public readonly ssoClientSecretKeyVaultPath!: pulumi.Output<string | undefined>;
    /**
     * Path to keyvault containing ssoContext secret.
     */
    public readonly ssoContext!: pulumi.Output<string | undefined>;
    /**
     * The type of single sign on Secret Type.
     */
    public readonly ssoSecretType!: pulumi.Output<string | undefined>;
    /**
     * URL to customer ADFS server for signing WVD SSO certificates.
     */
    public readonly ssoadfsAuthority!: pulumi.Output<string | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Is validation environment.
     */
    public readonly validationEnvironment!: pulumi.Output<boolean | undefined>;
    /**
     * VM template for sessionhosts configuration within hostpool.
     */
    public readonly vmTemplate!: pulumi.Output<string | undefined>;

    /**
     * Create a HostPool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HostPoolArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.hostPoolName === undefined) {
                throw new Error("Missing required property 'hostPoolName'");
            }
            if (!args || args.hostPoolType === undefined) {
                throw new Error("Missing required property 'hostPoolType'");
            }
            if (!args || args.loadBalancerType === undefined) {
                throw new Error("Missing required property 'loadBalancerType'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.preferredAppGroupType === undefined) {
                throw new Error("Missing required property 'preferredAppGroupType'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["customRdpProperty"] = args ? args.customRdpProperty : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["friendlyName"] = args ? args.friendlyName : undefined;
            inputs["hostPoolName"] = args ? args.hostPoolName : undefined;
            inputs["hostPoolType"] = args ? args.hostPoolType : undefined;
            inputs["loadBalancerType"] = args ? args.loadBalancerType : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["maxSessionLimit"] = args ? args.maxSessionLimit : undefined;
            inputs["personalDesktopAssignmentType"] = args ? args.personalDesktopAssignmentType : undefined;
            inputs["preferredAppGroupType"] = args ? args.preferredAppGroupType : undefined;
            inputs["registrationInfo"] = args ? args.registrationInfo : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["ring"] = args ? args.ring : undefined;
            inputs["ssoClientId"] = args ? args.ssoClientId : undefined;
            inputs["ssoClientSecretKeyVaultPath"] = args ? args.ssoClientSecretKeyVaultPath : undefined;
            inputs["ssoContext"] = args ? args.ssoContext : undefined;
            inputs["ssoSecretType"] = args ? args.ssoSecretType : undefined;
            inputs["ssoadfsAuthority"] = args ? args.ssoadfsAuthority : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["validationEnvironment"] = args ? args.validationEnvironment : undefined;
            inputs["vmTemplate"] = args ? args.vmTemplate : undefined;
            inputs["applicationGroupReferences"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["applicationGroupReferences"] = undefined /*out*/;
            inputs["customRdpProperty"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["friendlyName"] = undefined /*out*/;
            inputs["hostPoolType"] = undefined /*out*/;
            inputs["loadBalancerType"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["maxSessionLimit"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["personalDesktopAssignmentType"] = undefined /*out*/;
            inputs["preferredAppGroupType"] = undefined /*out*/;
            inputs["registrationInfo"] = undefined /*out*/;
            inputs["ring"] = undefined /*out*/;
            inputs["ssoClientId"] = undefined /*out*/;
            inputs["ssoClientSecretKeyVaultPath"] = undefined /*out*/;
            inputs["ssoContext"] = undefined /*out*/;
            inputs["ssoSecretType"] = undefined /*out*/;
            inputs["ssoadfsAuthority"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["validationEnvironment"] = undefined /*out*/;
            inputs["vmTemplate"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:desktopvirtualization/v20190123preview:HostPool" }, { type: "azure-nextgen:desktopvirtualization/v20190924preview:HostPool" }, { type: "azure-nextgen:desktopvirtualization/v20191210preview:HostPool" }, { type: "azure-nextgen:desktopvirtualization/v20200921preview:HostPool" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(HostPool.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a HostPool resource.
 */
export interface HostPoolArgs {
    /**
     * Custom rdp property of HostPool.
     */
    readonly customRdpProperty?: pulumi.Input<string>;
    /**
     * Description of HostPool.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Friendly name of HostPool.
     */
    readonly friendlyName?: pulumi.Input<string>;
    /**
     * The name of the host pool within the specified resource group
     */
    readonly hostPoolName: pulumi.Input<string>;
    /**
     * HostPool type for desktop.
     */
    readonly hostPoolType: pulumi.Input<string>;
    /**
     * The type of the load balancer.
     */
    readonly loadBalancerType: pulumi.Input<string>;
    /**
     * The geo-location where the resource lives
     */
    readonly location: pulumi.Input<string>;
    /**
     * The max session limit of HostPool.
     */
    readonly maxSessionLimit?: pulumi.Input<number>;
    /**
     * PersonalDesktopAssignment type for HostPool.
     */
    readonly personalDesktopAssignmentType?: pulumi.Input<string>;
    /**
     * The type of preferred application group type, default to Desktop Application Group
     */
    readonly preferredAppGroupType: pulumi.Input<string>;
    /**
     * The registration info of HostPool.
     */
    readonly registrationInfo?: pulumi.Input<inputs.desktopvirtualization.v20201019preview.RegistrationInfo>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The ring number of HostPool.
     */
    readonly ring?: pulumi.Input<number>;
    /**
     * ClientId for the registered Relying Party used to issue WVD SSO certificates.
     */
    readonly ssoClientId?: pulumi.Input<string>;
    /**
     * Path to Azure KeyVault storing the secret used for communication to ADFS.
     */
    readonly ssoClientSecretKeyVaultPath?: pulumi.Input<string>;
    /**
     * Path to keyvault containing ssoContext secret.
     */
    readonly ssoContext?: pulumi.Input<string>;
    /**
     * The type of single sign on Secret Type.
     */
    readonly ssoSecretType?: pulumi.Input<string>;
    /**
     * URL to customer ADFS server for signing WVD SSO certificates.
     */
    readonly ssoadfsAuthority?: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Is validation environment.
     */
    readonly validationEnvironment?: pulumi.Input<boolean>;
    /**
     * VM template for sessionhosts configuration within hostpool.
     */
    readonly vmTemplate?: pulumi.Input<string>;
}
