// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A virtual machine.
 */
export class VirtualMachine extends pulumi.CustomResource {
    /**
     * Get an existing VirtualMachine resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VirtualMachine {
        return new VirtualMachine(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:devtestlab/v20160515:VirtualMachine';

    /**
     * Returns true if the given object is an instance of VirtualMachine.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualMachine {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualMachine.__pulumiType;
    }

    /**
     * Indicates whether another user can take ownership of the virtual machine
     */
    public readonly allowClaim!: pulumi.Output<boolean | undefined>;
    /**
     * The applicable schedule for the virtual machine.
     */
    public readonly applicableSchedule!: pulumi.Output<outputs.devtestlab.v20160515.ApplicableScheduleResponse | undefined>;
    /**
     * The artifact deployment status for the virtual machine.
     */
    public readonly artifactDeploymentStatus!: pulumi.Output<outputs.devtestlab.v20160515.ArtifactDeploymentStatusPropertiesResponse | undefined>;
    /**
     * The artifacts to be installed on the virtual machine.
     */
    public readonly artifacts!: pulumi.Output<outputs.devtestlab.v20160515.ArtifactInstallPropertiesResponse[] | undefined>;
    /**
     * The resource identifier (Microsoft.Compute) of the virtual machine.
     */
    public /*out*/ readonly computeId!: pulumi.Output<string>;
    /**
     * The compute virtual machine properties.
     */
    public readonly computeVm!: pulumi.Output<outputs.devtestlab.v20160515.ComputeVmPropertiesResponse | undefined>;
    /**
     * The email address of creator of the virtual machine.
     */
    public readonly createdByUser!: pulumi.Output<string | undefined>;
    /**
     * The object identifier of the creator of the virtual machine.
     */
    public readonly createdByUserId!: pulumi.Output<string | undefined>;
    /**
     * The creation date of the virtual machine.
     */
    public readonly createdDate!: pulumi.Output<string | undefined>;
    /**
     * The custom image identifier of the virtual machine.
     */
    public readonly customImageId!: pulumi.Output<string | undefined>;
    /**
     * Indicates whether the virtual machine is to be created without a public IP address.
     */
    public readonly disallowPublicIpAddress!: pulumi.Output<boolean | undefined>;
    /**
     * The resource ID of the environment that contains this virtual machine, if any.
     */
    public readonly environmentId!: pulumi.Output<string | undefined>;
    /**
     * The expiration date for VM.
     */
    public readonly expirationDate!: pulumi.Output<string | undefined>;
    /**
     * The fully-qualified domain name of the virtual machine.
     */
    public readonly fqdn!: pulumi.Output<string | undefined>;
    /**
     * The Microsoft Azure Marketplace image reference of the virtual machine.
     */
    public readonly galleryImageReference!: pulumi.Output<outputs.devtestlab.v20160515.GalleryImageReferenceResponse | undefined>;
    /**
     * Indicates whether this virtual machine uses an SSH key for authentication.
     */
    public readonly isAuthenticationWithSshKey!: pulumi.Output<boolean | undefined>;
    /**
     * The lab subnet name of the virtual machine.
     */
    public readonly labSubnetName!: pulumi.Output<string | undefined>;
    /**
     * The lab virtual network identifier of the virtual machine.
     */
    public readonly labVirtualNetworkId!: pulumi.Output<string | undefined>;
    /**
     * The location of the resource.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The network interface properties.
     */
    public readonly networkInterface!: pulumi.Output<outputs.devtestlab.v20160515.NetworkInterfacePropertiesResponse | undefined>;
    /**
     * The notes of the virtual machine.
     */
    public readonly notes!: pulumi.Output<string | undefined>;
    /**
     * The OS type of the virtual machine.
     */
    public readonly osType!: pulumi.Output<string | undefined>;
    /**
     * The object identifier of the owner of the virtual machine.
     */
    public readonly ownerObjectId!: pulumi.Output<string | undefined>;
    /**
     * The user principal name of the virtual machine owner.
     */
    public readonly ownerUserPrincipalName!: pulumi.Output<string | undefined>;
    /**
     * The password of the virtual machine administrator.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * The provisioning status of the resource.
     */
    public readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * The size of the virtual machine.
     */
    public readonly size!: pulumi.Output<string | undefined>;
    /**
     * The SSH key of the virtual machine administrator.
     */
    public readonly sshKey!: pulumi.Output<string | undefined>;
    /**
     * Storage type to use for virtual machine (i.e. Standard, Premium).
     */
    public readonly storageType!: pulumi.Output<string | undefined>;
    /**
     * The tags of the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The unique immutable identifier of a resource (Guid).
     */
    public readonly uniqueIdentifier!: pulumi.Output<string | undefined>;
    /**
     * The user name of the virtual machine.
     */
    public readonly userName!: pulumi.Output<string | undefined>;
    /**
     * Tells source of creation of lab virtual machine. Output property only.
     */
    public readonly virtualMachineCreationSource!: pulumi.Output<string | undefined>;

    /**
     * Create a VirtualMachine resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualMachineArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.labName === undefined) {
                throw new Error("Missing required property 'labName'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["allowClaim"] = args ? args.allowClaim : undefined;
            inputs["applicableSchedule"] = args ? args.applicableSchedule : undefined;
            inputs["artifactDeploymentStatus"] = args ? args.artifactDeploymentStatus : undefined;
            inputs["artifacts"] = args ? args.artifacts : undefined;
            inputs["computeVm"] = args ? args.computeVm : undefined;
            inputs["createdByUser"] = args ? args.createdByUser : undefined;
            inputs["createdByUserId"] = args ? args.createdByUserId : undefined;
            inputs["createdDate"] = args ? args.createdDate : undefined;
            inputs["customImageId"] = args ? args.customImageId : undefined;
            inputs["disallowPublicIpAddress"] = args ? args.disallowPublicIpAddress : undefined;
            inputs["environmentId"] = args ? args.environmentId : undefined;
            inputs["expirationDate"] = args ? args.expirationDate : undefined;
            inputs["fqdn"] = args ? args.fqdn : undefined;
            inputs["galleryImageReference"] = args ? args.galleryImageReference : undefined;
            inputs["isAuthenticationWithSshKey"] = args ? args.isAuthenticationWithSshKey : undefined;
            inputs["labName"] = args ? args.labName : undefined;
            inputs["labSubnetName"] = args ? args.labSubnetName : undefined;
            inputs["labVirtualNetworkId"] = args ? args.labVirtualNetworkId : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkInterface"] = args ? args.networkInterface : undefined;
            inputs["notes"] = args ? args.notes : undefined;
            inputs["osType"] = args ? args.osType : undefined;
            inputs["ownerObjectId"] = args ? args.ownerObjectId : undefined;
            inputs["ownerUserPrincipalName"] = args ? args.ownerUserPrincipalName : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["size"] = args ? args.size : undefined;
            inputs["sshKey"] = args ? args.sshKey : undefined;
            inputs["storageType"] = args ? args.storageType : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["uniqueIdentifier"] = args ? args.uniqueIdentifier : undefined;
            inputs["userName"] = args ? args.userName : undefined;
            inputs["virtualMachineCreationSource"] = args ? args.virtualMachineCreationSource : undefined;
            inputs["computeId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["allowClaim"] = undefined /*out*/;
            inputs["applicableSchedule"] = undefined /*out*/;
            inputs["artifactDeploymentStatus"] = undefined /*out*/;
            inputs["artifacts"] = undefined /*out*/;
            inputs["computeId"] = undefined /*out*/;
            inputs["computeVm"] = undefined /*out*/;
            inputs["createdByUser"] = undefined /*out*/;
            inputs["createdByUserId"] = undefined /*out*/;
            inputs["createdDate"] = undefined /*out*/;
            inputs["customImageId"] = undefined /*out*/;
            inputs["disallowPublicIpAddress"] = undefined /*out*/;
            inputs["environmentId"] = undefined /*out*/;
            inputs["expirationDate"] = undefined /*out*/;
            inputs["fqdn"] = undefined /*out*/;
            inputs["galleryImageReference"] = undefined /*out*/;
            inputs["isAuthenticationWithSshKey"] = undefined /*out*/;
            inputs["labSubnetName"] = undefined /*out*/;
            inputs["labVirtualNetworkId"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["networkInterface"] = undefined /*out*/;
            inputs["notes"] = undefined /*out*/;
            inputs["osType"] = undefined /*out*/;
            inputs["ownerObjectId"] = undefined /*out*/;
            inputs["ownerUserPrincipalName"] = undefined /*out*/;
            inputs["password"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["size"] = undefined /*out*/;
            inputs["sshKey"] = undefined /*out*/;
            inputs["storageType"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["uniqueIdentifier"] = undefined /*out*/;
            inputs["userName"] = undefined /*out*/;
            inputs["virtualMachineCreationSource"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:devtestlab/latest:VirtualMachine" }, { type: "azurerm:devtestlab/v20150521preview:VirtualMachine" }, { type: "azurerm:devtestlab/v20180915:VirtualMachine" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(VirtualMachine.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a VirtualMachine resource.
 */
export interface VirtualMachineArgs {
    /**
     * Indicates whether another user can take ownership of the virtual machine
     */
    readonly allowClaim?: pulumi.Input<boolean>;
    /**
     * The applicable schedule for the virtual machine.
     */
    readonly applicableSchedule?: pulumi.Input<inputs.devtestlab.v20160515.ApplicableSchedule>;
    /**
     * The artifact deployment status for the virtual machine.
     */
    readonly artifactDeploymentStatus?: pulumi.Input<inputs.devtestlab.v20160515.ArtifactDeploymentStatusProperties>;
    /**
     * The artifacts to be installed on the virtual machine.
     */
    readonly artifacts?: pulumi.Input<pulumi.Input<inputs.devtestlab.v20160515.ArtifactInstallProperties>[]>;
    /**
     * The compute virtual machine properties.
     */
    readonly computeVm?: pulumi.Input<inputs.devtestlab.v20160515.ComputeVmProperties>;
    /**
     * The email address of creator of the virtual machine.
     */
    readonly createdByUser?: pulumi.Input<string>;
    /**
     * The object identifier of the creator of the virtual machine.
     */
    readonly createdByUserId?: pulumi.Input<string>;
    /**
     * The creation date of the virtual machine.
     */
    readonly createdDate?: pulumi.Input<string>;
    /**
     * The custom image identifier of the virtual machine.
     */
    readonly customImageId?: pulumi.Input<string>;
    /**
     * Indicates whether the virtual machine is to be created without a public IP address.
     */
    readonly disallowPublicIpAddress?: pulumi.Input<boolean>;
    /**
     * The resource ID of the environment that contains this virtual machine, if any.
     */
    readonly environmentId?: pulumi.Input<string>;
    /**
     * The expiration date for VM.
     */
    readonly expirationDate?: pulumi.Input<string>;
    /**
     * The fully-qualified domain name of the virtual machine.
     */
    readonly fqdn?: pulumi.Input<string>;
    /**
     * The Microsoft Azure Marketplace image reference of the virtual machine.
     */
    readonly galleryImageReference?: pulumi.Input<inputs.devtestlab.v20160515.GalleryImageReference>;
    /**
     * Indicates whether this virtual machine uses an SSH key for authentication.
     */
    readonly isAuthenticationWithSshKey?: pulumi.Input<boolean>;
    /**
     * The name of the lab.
     */
    readonly labName: pulumi.Input<string>;
    /**
     * The lab subnet name of the virtual machine.
     */
    readonly labSubnetName?: pulumi.Input<string>;
    /**
     * The lab virtual network identifier of the virtual machine.
     */
    readonly labVirtualNetworkId?: pulumi.Input<string>;
    /**
     * The location of the resource.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the virtual machine.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The network interface properties.
     */
    readonly networkInterface?: pulumi.Input<inputs.devtestlab.v20160515.NetworkInterfaceProperties>;
    /**
     * The notes of the virtual machine.
     */
    readonly notes?: pulumi.Input<string>;
    /**
     * The OS type of the virtual machine.
     */
    readonly osType?: pulumi.Input<string>;
    /**
     * The object identifier of the owner of the virtual machine.
     */
    readonly ownerObjectId?: pulumi.Input<string>;
    /**
     * The user principal name of the virtual machine owner.
     */
    readonly ownerUserPrincipalName?: pulumi.Input<string>;
    /**
     * The password of the virtual machine administrator.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * The provisioning status of the resource.
     */
    readonly provisioningState?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The size of the virtual machine.
     */
    readonly size?: pulumi.Input<string>;
    /**
     * The SSH key of the virtual machine administrator.
     */
    readonly sshKey?: pulumi.Input<string>;
    /**
     * Storage type to use for virtual machine (i.e. Standard, Premium).
     */
    readonly storageType?: pulumi.Input<string>;
    /**
     * The tags of the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The unique immutable identifier of a resource (Guid).
     */
    readonly uniqueIdentifier?: pulumi.Input<string>;
    /**
     * The user name of the virtual machine.
     */
    readonly userName?: pulumi.Input<string>;
    /**
     * Tells source of creation of lab virtual machine. Output property only.
     */
    readonly virtualMachineCreationSource?: pulumi.Input<string>;
}
