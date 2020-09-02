// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Domain service.
 *
 * ## Example Usage
 * ### Create Domain Service
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const domainService = new azurerm.aad.latest.DomainService("domainService", {
 *     domainName: "TestDomainService.com",
 *     domainSecuritySettings: {
 *         ntlmV1: "Enabled",
 *         syncNtlmPasswords: "Enabled",
 *         tlsV1: "Disabled",
 *     },
 *     domainServiceName: "TestDomainService.com",
 *     filteredSync: "Enabled",
 *     ldapsSettings: {
 *         externalAccess: "Enabled",
 *         ldaps: "Enabled",
 *         pfxCertificate: "MIIDPDCCAiSgAwIBAgIQQUI9P6tq2p9OFIJa7DLNvTANBgkqhkiG9w0BAQsFADAgMR4w...",
 *         pfxCertificatePassword: "Password01",
 *     },
 *     location: "West US",
 *     notificationSettings: {
 *         additionalRecipients: [
 *             "jicha@microsoft.com",
 *             "caalmont@microsoft.com",
 *         ],
 *         notifyDcAdmins: "Enabled",
 *         notifyGlobalAdmins: "Enabled",
 *     },
 *     replicaSets: [{
 *         location: "West US",
 *         subnetId: "/subscriptions/1639790a-76a2-4ac4-98d9-8562f5dfcb4d/resourceGroups/TestNetworkResourceGroup/providers/Microsoft.Network/virtualNetworks/TestVnetWUS/subnets/TestSubnetWUS",
 *     }],
 *     resourceGroupName: "TestResourceGroup",
 * });
 *
 * ```
 */
export class DomainService extends pulumi.CustomResource {
    /**
     * Get an existing DomainService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DomainService {
        return new DomainService(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:aad/latest:DomainService';

    /**
     * Returns true if the given object is an instance of DomainService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DomainService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DomainService.__pulumiType;
    }

    /**
     * Deployment Id
     */
    public /*out*/ readonly deploymentId!: pulumi.Output<string>;
    /**
     * The name of the Azure domain that the user would like to deploy Domain Services to.
     */
    public readonly domainName!: pulumi.Output<string | undefined>;
    /**
     * DomainSecurity Settings
     */
    public readonly domainSecuritySettings!: pulumi.Output<outputs.aad.latest.DomainSecuritySettingsResponse | undefined>;
    /**
     * Resource etag
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * Enabled or Disabled flag to turn on Group-based filtered sync
     */
    public readonly filteredSync!: pulumi.Output<string | undefined>;
    /**
     * Secure LDAP Settings
     */
    public readonly ldapsSettings!: pulumi.Output<outputs.aad.latest.LdapsSettingsResponse | undefined>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Notification Settings
     */
    public readonly notificationSettings!: pulumi.Output<outputs.aad.latest.NotificationSettingsResponse | undefined>;
    /**
     * the current deployment or provisioning state, which only appears in the response.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * List of ReplicaSets
     */
    public readonly replicaSets!: pulumi.Output<outputs.aad.latest.ReplicaSetResponse[] | undefined>;
    /**
     * SyncOwner ReplicaSet Id
     */
    public /*out*/ readonly syncOwner!: pulumi.Output<string>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Azure Active Directory Tenant Id
     */
    public /*out*/ readonly tenantId!: pulumi.Output<string>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Data Model Version
     */
    public /*out*/ readonly version!: pulumi.Output<number>;

    /**
     * Create a DomainService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DomainServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DomainServiceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as DomainServiceArgs | undefined;
            if (!args || args.domainServiceName === undefined) {
                throw new Error("Missing required property 'domainServiceName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["domainName"] = args ? args.domainName : undefined;
            inputs["domainSecuritySettings"] = args ? args.domainSecuritySettings : undefined;
            inputs["domainServiceName"] = args ? args.domainServiceName : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["filteredSync"] = args ? args.filteredSync : undefined;
            inputs["ldapsSettings"] = args ? args.ldapsSettings : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["notificationSettings"] = args ? args.notificationSettings : undefined;
            inputs["replicaSets"] = args ? args.replicaSets : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["deploymentId"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["syncOwner"] = undefined /*out*/;
            inputs["tenantId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:aad/v20170101:DomainService" }, { type: "azurerm:aad/v20170601:DomainService" }, { type: "azurerm:aad/v20200101:DomainService" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(DomainService.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DomainService resource.
 */
export interface DomainServiceArgs {
    /**
     * The name of the Azure domain that the user would like to deploy Domain Services to.
     */
    readonly domainName?: pulumi.Input<string>;
    /**
     * DomainSecurity Settings
     */
    readonly domainSecuritySettings?: pulumi.Input<inputs.aad.latest.DomainSecuritySettings>;
    /**
     * The name of the domain service.
     */
    readonly domainServiceName: pulumi.Input<string>;
    /**
     * Resource etag
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Enabled or Disabled flag to turn on Group-based filtered sync
     */
    readonly filteredSync?: pulumi.Input<string>;
    /**
     * Secure LDAP Settings
     */
    readonly ldapsSettings?: pulumi.Input<inputs.aad.latest.LdapsSettings>;
    /**
     * Resource location
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Notification Settings
     */
    readonly notificationSettings?: pulumi.Input<inputs.aad.latest.NotificationSettings>;
    /**
     * List of ReplicaSets
     */
    readonly replicaSets?: pulumi.Input<pulumi.Input<inputs.aad.latest.ReplicaSet>[]>;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
