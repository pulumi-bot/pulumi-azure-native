import * as pulumi from "@pulumi/pulumi";
import * as random from "@pulumi/random";
import * as authorization from "@pulumi/azure-nextgen/authorization/latest";
// TODO change to latest when https://github.com/Azure/azure-rest-api-specs/issues/11634 is fixed
import * as containerinstance from "@pulumi/azure-nextgen/containerinstance/v20191201";
import * as keyvault from "@pulumi/azure-nextgen/keyvault/latest";
import * as managedidentity from "@pulumi/azure-nextgen/managedidentity/latest";
import * as resources from "@pulumi/azure-nextgen/resources/latest";

const randomString = new random.RandomString("random", {
    length: 12,
    special: false,
    upper: false,
    number: false,
}).result;

const resourceGroup = new resources.ResourceGroup("rg", {
    resourceGroupName: randomString,
    location: "westus2",
});

const userIdentity = new managedidentity.UserAssignedIdentity("uai", {
    resourceGroupName: resourceGroup.name,
    location: resourceGroup.location,
    resourceName: randomString,
});

const container = new containerinstance.ContainerGroup("containergroup", {
    resourceGroupName: resourceGroup.name,
    containerGroupName: randomString,
    location: "westus2",
    osType: containerinstance.OperatingSystemTypes.Linux,
    containers: [{
        name: "foo",
        image: "nginx",
        resources: {
            requests: {
                memoryInGB: 1,
                cpu: 1,
            },
        },
    }],
    identity: {
        type: containerinstance.ResourceIdentityType.UserAssigned,
        userAssignedIdentities: userIdentity.id.apply(id => {
            const dict: { [key: string] : object } = {};
            dict[id] = {};
            return dict;
        }),
    },
});

const config = pulumi.output(authorization.getClientConfig());

const vault = new keyvault.Vault("vault", {
    resourceGroupName: resourceGroup.name,
    location: resourceGroup.location,
    vaultName: randomString,
    properties: {
        accessPolicies: [{
            objectId: config.objectId,
            permissions: {
                keys: [
                    keyvault.KeyPermissions.Get,
                    keyvault.KeyPermissions.Create,
                    keyvault.KeyPermissions.Delete,
                    keyvault.KeyPermissions.List,
                    keyvault.KeyPermissions.Recover,
                    keyvault.KeyPermissions.Purge,
                ],
                secrets: [
                    keyvault.SecretPermissions.Get,
                    keyvault.SecretPermissions.List,
                    keyvault.SecretPermissions.Set,
                    keyvault.SecretPermissions.Delete,
                    keyvault.SecretPermissions.Recover,
                    keyvault.SecretPermissions.Purge,
                ],
            },
            tenantId: config.tenantId,
        }, {
            objectId: container.identity.apply(i => i?.userAssignedIdentities!).apply(uai => Object.values(uai)[0].principalId),
            permissions: {
                secrets: [keyvault.SecretPermissions.Get],
            },
            tenantId: config.tenantId,
        }],
        sku: {
            family: keyvault.SkuFamily.A,
            name: keyvault.SkuName.Standard,
        },
        tenantId: config.tenantId,
    },
});

const secret = new keyvault.Secret("mysecret", {
    resourceGroupName: resourceGroup.name,
    vaultName: vault.name,
    secretName: "mysecret",
    properties: {
        value: "myvalue",
    },
});

const key = new keyvault.Key("mykey", {
    resourceGroupName: resourceGroup.name,
    vaultName: vault.name,
    keyName: "mykey",
    properties: {
        kty: "RSA",
        keySize: 2048,
        keyOps: [
            "decrypt",
            "encrypt",
            "sign",
            "unwrapKey",
            "verify",
            "wrapKey",
        ],
    },
});
