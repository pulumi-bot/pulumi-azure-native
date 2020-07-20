import * as pulumi from "@pulumi/pulumi";
import * as azurerm from "../../sdk/nodejs";
import * as cosmosdb from "./component";

const resourceGroup = new azurerm.core.ResourceGroup("rg", {
    name: "azurerm-cosmos",
    location: "westeurope",
    tags: {
        Owner: "mikhailshilkov",
    },
});

export const cosmosdbAccount = new cosmosdb.DatabaseAccount("pulumicosmosdb", {
    resourceGroup,
    api: "Sql",
    consisencyPolicy: { defaultConsistencyLevel: "Session" },
    location: { 
        type: "multiple",
        regions: ["westeurope", "northeurope"],
        enableMultipleWriteLocations: true,
        enableAutomaticFailover: true,
    },    
});

// const cosmosdbAccount2 = new azurerm.documentdb.DatabaseAccount("pulumicosmosdb", {
//     resourceGroupName: resourceGroup.name,
//     name: "pulumicosmosdb",
//     location: resourceGroup.location,
//     kind: "GlobalDocumentDB",
//     properties: {
//         consistencyPolicy: {
//             defaultConsistencyLevel: "Session",
//         },        
//         locations: [{
//             locationName: "westeurope",
//             failoverPriority: 0,
//             isZoneRedundant: false
//         }, {
//             locationName: "northeurope",
//             failoverPriority: 1,
//             isZoneRedundant: false
//         }],
//         databaseAccountOfferType: "Standard",
//         enableAutomaticFailover: true,
//         capabilities: [],
//         enableMultipleWriteLocations: true,
//     }
// });