// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Virtual machine model
 *
 * ## Example Usage
 * ### CreateVirtualMachine
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const virtualMachine = new azurerm.vmwarecloudsimple.v20190401.VirtualMachine("virtualMachine", {
 *     amountOfRam: 4096,
 *     disks: [{
 *         controllerId: "1000",
 *         independenceMode: "persistent",
 *         totalSize: 10485760,
 *         virtualDiskId: "2000",
 *     }],
 *     location: "westus2",
 *     nics: [{
 *         network: {
 *             id: "/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/virtualNetworks/dvportgroup-19",
 *         },
 *         nicType: "E1000",
 *         powerOnBoot: true,
 *         virtualNicId: "4000",
 *     }],
 *     numberOfCores: 2,
 *     privateCloudId: "/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud",
 *     resourceGroupName: "myResourceGroup",
 *     resourcePool: {
 *         id: "/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/resourcePools/resgroup-26",
 *     },
 *     templateId: "/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud/virtualMachineTemplates/vm-34",
 *     virtualMachineName: "myVirtualMachine",
 * });
 *
 * ```
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
    public static readonly __pulumiType = 'azurerm:vmwarecloudsimple/v20190401:VirtualMachine';

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
     * The amount of memory
     */
    public readonly amountOfRam!: pulumi.Output<number>;
    /**
     * The list of Virtual Disks' Controllers
     */
    public /*out*/ readonly controllers!: pulumi.Output<outputs.vmwarecloudsimple.v20190401.VirtualDiskControllerResponse[]>;
    /**
     * Virtual machine properties
     */
    public readonly customization!: pulumi.Output<outputs.vmwarecloudsimple.v20190401.GuestOSCustomizationResponse | undefined>;
    /**
     * The list of Virtual Disks
     */
    public readonly disks!: pulumi.Output<outputs.vmwarecloudsimple.v20190401.VirtualDiskResponse[] | undefined>;
    /**
     * The DNS name of Virtual Machine in VCenter
     */
    public /*out*/ readonly dnsname!: pulumi.Output<string>;
    /**
     * Expose Guest OS or not
     */
    public readonly exposeToGuestVM!: pulumi.Output<boolean | undefined>;
    /**
     * The path to virtual machine folder in VCenter
     */
    public /*out*/ readonly folder!: pulumi.Output<string>;
    /**
     * The name of Guest OS
     */
    public /*out*/ readonly guestOS!: pulumi.Output<string>;
    /**
     * The Guest OS type
     */
    public /*out*/ readonly guestOSType!: pulumi.Output<string>;
    /**
     * Azure region
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * {virtualMachineName}
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The list of Virtual NICs
     */
    public readonly nics!: pulumi.Output<outputs.vmwarecloudsimple.v20190401.VirtualNicResponse[] | undefined>;
    /**
     * The number of CPU cores
     */
    public readonly numberOfCores!: pulumi.Output<number>;
    /**
     * Password for login. Deprecated - use customization property
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * Private Cloud Id
     */
    public readonly privateCloudId!: pulumi.Output<string>;
    /**
     * The provisioning status of the resource
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The public ip of Virtual Machine
     */
    public /*out*/ readonly publicIP!: pulumi.Output<string>;
    /**
     * Virtual Machines Resource Pool
     */
    public readonly resourcePool!: pulumi.Output<outputs.vmwarecloudsimple.v20190401.ResourcePoolResponse | undefined>;
    /**
     * The status of Virtual machine
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The list of tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Virtual Machine Template Id
     */
    public readonly templateId!: pulumi.Output<string | undefined>;
    /**
     * {resourceProviderNamespace}/{resourceType}
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Username for login. Deprecated - use customization property
     */
    public readonly username!: pulumi.Output<string | undefined>;
    /**
     * The list of Virtual VSphere Networks
     */
    public readonly vSphereNetworks!: pulumi.Output<string[] | undefined>;
    /**
     * The internal id of Virtual Machine in VCenter
     */
    public /*out*/ readonly vmId!: pulumi.Output<string>;
    /**
     * VMware tools version
     */
    public /*out*/ readonly vmwaretools!: pulumi.Output<string>;

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
            if (!args || args.amountOfRam === undefined) {
                throw new Error("Missing required property 'amountOfRam'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.numberOfCores === undefined) {
                throw new Error("Missing required property 'numberOfCores'");
            }
            if (!args || args.privateCloudId === undefined) {
                throw new Error("Missing required property 'privateCloudId'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.virtualMachineName === undefined) {
                throw new Error("Missing required property 'virtualMachineName'");
            }
            inputs["amountOfRam"] = args ? args.amountOfRam : undefined;
            inputs["customization"] = args ? args.customization : undefined;
            inputs["disks"] = args ? args.disks : undefined;
            inputs["exposeToGuestVM"] = args ? args.exposeToGuestVM : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["nics"] = args ? args.nics : undefined;
            inputs["numberOfCores"] = args ? args.numberOfCores : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["privateCloudId"] = args ? args.privateCloudId : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourcePool"] = args ? args.resourcePool : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["templateId"] = args ? args.templateId : undefined;
            inputs["username"] = args ? args.username : undefined;
            inputs["vSphereNetworks"] = args ? args.vSphereNetworks : undefined;
            inputs["virtualMachineName"] = args ? args.virtualMachineName : undefined;
            inputs["controllers"] = undefined /*out*/;
            inputs["dnsname"] = undefined /*out*/;
            inputs["folder"] = undefined /*out*/;
            inputs["guestOS"] = undefined /*out*/;
            inputs["guestOSType"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["publicIP"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["vmId"] = undefined /*out*/;
            inputs["vmwaretools"] = undefined /*out*/;
        } else {
            inputs["amountOfRam"] = undefined /*out*/;
            inputs["controllers"] = undefined /*out*/;
            inputs["customization"] = undefined /*out*/;
            inputs["disks"] = undefined /*out*/;
            inputs["dnsname"] = undefined /*out*/;
            inputs["exposeToGuestVM"] = undefined /*out*/;
            inputs["folder"] = undefined /*out*/;
            inputs["guestOS"] = undefined /*out*/;
            inputs["guestOSType"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["nics"] = undefined /*out*/;
            inputs["numberOfCores"] = undefined /*out*/;
            inputs["password"] = undefined /*out*/;
            inputs["privateCloudId"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["publicIP"] = undefined /*out*/;
            inputs["resourcePool"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["templateId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["username"] = undefined /*out*/;
            inputs["vSphereNetworks"] = undefined /*out*/;
            inputs["vmId"] = undefined /*out*/;
            inputs["vmwaretools"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:vmwarecloudsimple/latest:VirtualMachine" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(VirtualMachine.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a VirtualMachine resource.
 */
export interface VirtualMachineArgs {
    /**
     * The amount of memory
     */
    readonly amountOfRam: pulumi.Input<number>;
    /**
     * Virtual machine properties
     */
    readonly customization?: pulumi.Input<inputs.vmwarecloudsimple.v20190401.GuestOSCustomization>;
    /**
     * The list of Virtual Disks
     */
    readonly disks?: pulumi.Input<pulumi.Input<inputs.vmwarecloudsimple.v20190401.VirtualDisk>[]>;
    /**
     * Expose Guest OS or not
     */
    readonly exposeToGuestVM?: pulumi.Input<boolean>;
    /**
     * Azure region
     */
    readonly location: pulumi.Input<string>;
    /**
     * The list of Virtual NICs
     */
    readonly nics?: pulumi.Input<pulumi.Input<inputs.vmwarecloudsimple.v20190401.VirtualNic>[]>;
    /**
     * The number of CPU cores
     */
    readonly numberOfCores: pulumi.Input<number>;
    /**
     * Password for login. Deprecated - use customization property
     */
    readonly password?: pulumi.Input<string>;
    /**
     * Private Cloud Id
     */
    readonly privateCloudId: pulumi.Input<string>;
    /**
     * The name of the resource group
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Virtual Machines Resource Pool
     */
    readonly resourcePool?: pulumi.Input<inputs.vmwarecloudsimple.v20190401.ResourcePool>;
    /**
     * The list of tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Virtual Machine Template Id
     */
    readonly templateId?: pulumi.Input<string>;
    /**
     * Username for login. Deprecated - use customization property
     */
    readonly username?: pulumi.Input<string>;
    /**
     * The list of Virtual VSphere Networks
     */
    readonly vSphereNetworks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * virtual machine name
     */
    readonly virtualMachineName: pulumi.Input<string>;
}
