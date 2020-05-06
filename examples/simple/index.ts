import * as pulumi from "@pulumi/pulumi";
import { networkInterfaces } from "os";

class ResourceGroup extends pulumi.CustomResource {
    constructor(name: string, args: any, opts?: pulumi.CustomResourceOptions) {
        super("azurerm:core:ResourceGroup", name, args, opts);
    }
}

class ContainerGroup extends pulumi.CustomResource {
    constructor(name: string, args: any, opts?: pulumi.CustomResourceOptions) {
        super("azurerm:containerinstance:ContainerGroup", name, args, opts);
    }
}

class VirtualMachine extends pulumi.CustomResource {
    constructor(name: string, args: any, opts?: pulumi.CustomResourceOptions) {
        super("azurerm:compute:VirtualMachine", name, args, opts);
    }
}

class NetworkInterface extends pulumi.CustomResource {
    id: pulumi.Output<string>;
    constructor(name: string, args: any, opts?: pulumi.CustomResourceOptions) {
        super("azurerm:network:NetworkInterface", name, args, opts);
    }
}


class VirtualNetwork extends pulumi.CustomResource {
    id: pulumi.Output<string>;
    constructor(name: string, args: any, opts?: pulumi.CustomResourceOptions) {
        super("azurerm:network:VirtualNetwork", name, args, opts);
    }
}

const resourceGroupName = "azurerm";

const resourceGroup = new ResourceGroup("azurerm", {
    resourceGroupName: resourceGroupName,
    parameters: {
        location: "westus2",
        tags: {
            Owner: "mikhailshilkov",
        },
    },
});

const containerinstance = new ContainerGroup("abc", {
    resourceGroupName: resourceGroupName,
    // should be autonamed?
    containerGroupName: "abc-1234",
    // should be inlined via use of `"in": "body"`?
    containerGroup: {
        location: "westus2",
        // should be inlined via 'x-ms-client-flatten'
        properties: {
            osType: "Linux",
            containers: [{
                name: "foo",
                // should be inlined via 'x-ms-client-flatten'
                properties: {
                    image: "nginx",
                    resources: {
                        requests: {
                            memoryInGB: "1",
                            cpu: "1",
                        },
                    },
                },
            }],
        },
    },
}, { dependsOn: [resourceGroup]});

const vnet = new VirtualNetwork("vnet", {
    resourceGroupName: resourceGroupName,
    virtualNetworkName: "vnet-1234",
    parameters: {
        location: "westus2",
        properties: {
            addressSpace: {
                addressPrefixes: ["10.1.0.0/24"],
            },
            subnets: [{
                name: "default",
                properties: {
                    addressPrefix: "10.1.0.0/24",
                },
            }],
        },
    },
}, { dependsOn: [resourceGroup]});

const networkInterface = new NetworkInterface("nic", {
    resourceGroupName: "aks-rg70afafca",
    networkInterfaceName: "nic-1234",
    parameters: {
        location: "westus2",
        properties: {
            ipConfigurations: [{
                name: "ipconfig1",
                properties: {
                    subnet: {
                        id: pulumi.interpolate`${vnet.id}/subnets/default`,
                    },
                    privateIPAllocationMethod: "Dynamic",
                },
            }],
        },
    },
});

const virtualmachine  = new VirtualMachine("vm", {
    resourceGroupName: resourceGroupName,
    vmName: "abc-1234",
    parameters: {
        location: "westus2",
        properties: {
            hardwareProfile: {
                vmSize: "Standard_A0",
            },
            networkProfile: {
                networkInterfaces: [{
                    id: networkInterface.id,
                }],
            },
            storageProfile: {
                imageReference: {
                    publisher: "Canonical",
                    offer: "UbuntuServer",
                    sku: "18.04-LTS",
                    version: "latest"
                }
            },
            osProfile: {
                computerName: "foo",
                adminUsername: "someusername",
                adminPassword: "someFancyp@wd2!",
            },
        },
    },
});
