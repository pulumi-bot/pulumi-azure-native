// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./appServiceEnvironment";
export * from "./appServicePlan";
export * from "./appServicePlanRouteForVnet";
export * from "./certificate";
export * from "./getAppServiceEnvironment";
export * from "./getAppServicePlan";
export * from "./getCertificate";
export * from "./getStaticSite";
export * from "./getWebApp";
export * from "./getWebAppDeployment";
export * from "./getWebAppDeploymentSlot";
export * from "./getWebAppDiagnosticLogsConfiguration";
export * from "./getWebAppDomainOwnershipIdentifier";
export * from "./getWebAppDomainOwnershipIdentifierSlot";
export * from "./getWebAppFunction";
export * from "./getWebAppHostNameBinding";
export * from "./getWebAppHostNameBindingSlot";
export * from "./getWebAppHybridConnection";
export * from "./getWebAppHybridConnectionSlot";
export * from "./getWebAppInstanceFunctionSlot";
export * from "./getWebAppPremierAddOn";
export * from "./getWebAppPremierAddOnSlot";
export * from "./getWebAppPrivateEndpointConnection";
export * from "./getWebAppPublicCertificate";
export * from "./getWebAppPublicCertificateSlot";
export * from "./getWebAppRelayServiceConnection";
export * from "./getWebAppRelayServiceConnectionSlot";
export * from "./getWebAppSiteExtension";
export * from "./getWebAppSiteExtensionSlot";
export * from "./getWebAppSlot";
export * from "./getWebAppSlotConfigurationNames";
export * from "./getWebAppSourceControl";
export * from "./getWebAppSourceControlSlot";
export * from "./getWebAppSwiftVirtualNetworkConnection";
export * from "./getWebAppSwiftVirtualNetworkConnectionSlot";
export * from "./getWebAppVnetConnection";
export * from "./getWebAppVnetConnectionSlot";
export * from "./listAppServicePlanHybridConnectionKeys";
export * from "./listSiteIdentifiersAssignedToHostName";
export * from "./listStaticSiteBuildFunctionAppSettings";
export * from "./listStaticSiteFunctionAppSettings";
export * from "./listStaticSiteSecrets";
export * from "./listStaticSiteUsers";
export * from "./listWebAppApplicationSettings";
export * from "./listWebAppApplicationSettingsSlot";
export * from "./listWebAppAuthSettings";
export * from "./listWebAppAuthSettingsSlot";
export * from "./listWebAppAzureStorageAccounts";
export * from "./listWebAppAzureStorageAccountsSlot";
export * from "./listWebAppBackupConfiguration";
export * from "./listWebAppBackupConfigurationSlot";
export * from "./listWebAppBackupStatusSecrets";
export * from "./listWebAppBackupStatusSecretsSlot";
export * from "./listWebAppConnectionStrings";
export * from "./listWebAppConnectionStringsSlot";
export * from "./listWebAppFunctionKeys";
export * from "./listWebAppFunctionKeysSlot";
export * from "./listWebAppFunctionSecrets";
export * from "./listWebAppFunctionSecretsSlot";
export * from "./listWebAppHostKeys";
export * from "./listWebAppHostKeysSlot";
export * from "./listWebAppMetadata";
export * from "./listWebAppMetadataSlot";
export * from "./listWebAppPublishingCredentials";
export * from "./listWebAppPublishingCredentialsSlot";
export * from "./listWebAppSiteBackups";
export * from "./listWebAppSiteBackupsSlot";
export * from "./listWebAppSitePushSettings";
export * from "./listWebAppSitePushSettingsSlot";
export * from "./listWebAppSyncFunctionTriggers";
export * from "./listWebAppSyncFunctionTriggersSlot";
export * from "./staticSite";
export * from "./webApp";
export * from "./webAppApplicationSettings";
export * from "./webAppApplicationSettingsSlot";
export * from "./webAppAuthSettings";
export * from "./webAppAuthSettingsSlot";
export * from "./webAppAuthSettingsV2";
export * from "./webAppAuthSettingsV2Slot";
export * from "./webAppAzureStorageAccounts";
export * from "./webAppAzureStorageAccountsSlot";
export * from "./webAppBackupConfiguration";
export * from "./webAppBackupConfigurationSlot";
export * from "./webAppConnectionStrings";
export * from "./webAppConnectionStringsSlot";
export * from "./webAppDeployment";
export * from "./webAppDeploymentSlot";
export * from "./webAppDiagnosticLogsConfiguration";
export * from "./webAppDomainOwnershipIdentifier";
export * from "./webAppDomainOwnershipIdentifierSlot";
export * from "./webAppFunction";
export * from "./webAppHostNameBinding";
export * from "./webAppHostNameBindingSlot";
export * from "./webAppHybridConnection";
export * from "./webAppHybridConnectionSlot";
export * from "./webAppInstanceFunctionSlot";
export * from "./webAppMetadata";
export * from "./webAppMetadataSlot";
export * from "./webAppPremierAddOn";
export * from "./webAppPremierAddOnSlot";
export * from "./webAppPrivateEndpointConnection";
export * from "./webAppPublicCertificate";
export * from "./webAppPublicCertificateSlot";
export * from "./webAppRelayServiceConnection";
export * from "./webAppRelayServiceConnectionSlot";
export * from "./webAppSiteExtension";
export * from "./webAppSiteExtensionSlot";
export * from "./webAppSitePushSettings";
export * from "./webAppSitePushSettingsSlot";
export * from "./webAppSlot";
export * from "./webAppSlotConfigurationNames";
export * from "./webAppSourceControl";
export * from "./webAppSourceControlSlot";
export * from "./webAppSwiftVirtualNetworkConnection";
export * from "./webAppSwiftVirtualNetworkConnectionSlot";
export * from "./webAppVnetConnection";
export * from "./webAppVnetConnectionSlot";

// Export enums:
export * from "../../types/enums/web/v20200901";

// Import resources to register:
import { AppServiceEnvironment } from "./appServiceEnvironment";
import { AppServicePlan } from "./appServicePlan";
import { AppServicePlanRouteForVnet } from "./appServicePlanRouteForVnet";
import { Certificate } from "./certificate";
import { StaticSite } from "./staticSite";
import { WebApp } from "./webApp";
import { WebAppApplicationSettings } from "./webAppApplicationSettings";
import { WebAppApplicationSettingsSlot } from "./webAppApplicationSettingsSlot";
import { WebAppAuthSettings } from "./webAppAuthSettings";
import { WebAppAuthSettingsSlot } from "./webAppAuthSettingsSlot";
import { WebAppAuthSettingsV2 } from "./webAppAuthSettingsV2";
import { WebAppAuthSettingsV2Slot } from "./webAppAuthSettingsV2Slot";
import { WebAppAzureStorageAccounts } from "./webAppAzureStorageAccounts";
import { WebAppAzureStorageAccountsSlot } from "./webAppAzureStorageAccountsSlot";
import { WebAppBackupConfiguration } from "./webAppBackupConfiguration";
import { WebAppBackupConfigurationSlot } from "./webAppBackupConfigurationSlot";
import { WebAppConnectionStrings } from "./webAppConnectionStrings";
import { WebAppConnectionStringsSlot } from "./webAppConnectionStringsSlot";
import { WebAppDeployment } from "./webAppDeployment";
import { WebAppDeploymentSlot } from "./webAppDeploymentSlot";
import { WebAppDiagnosticLogsConfiguration } from "./webAppDiagnosticLogsConfiguration";
import { WebAppDomainOwnershipIdentifier } from "./webAppDomainOwnershipIdentifier";
import { WebAppDomainOwnershipIdentifierSlot } from "./webAppDomainOwnershipIdentifierSlot";
import { WebAppFunction } from "./webAppFunction";
import { WebAppHostNameBinding } from "./webAppHostNameBinding";
import { WebAppHostNameBindingSlot } from "./webAppHostNameBindingSlot";
import { WebAppHybridConnection } from "./webAppHybridConnection";
import { WebAppHybridConnectionSlot } from "./webAppHybridConnectionSlot";
import { WebAppInstanceFunctionSlot } from "./webAppInstanceFunctionSlot";
import { WebAppMetadata } from "./webAppMetadata";
import { WebAppMetadataSlot } from "./webAppMetadataSlot";
import { WebAppPremierAddOn } from "./webAppPremierAddOn";
import { WebAppPremierAddOnSlot } from "./webAppPremierAddOnSlot";
import { WebAppPrivateEndpointConnection } from "./webAppPrivateEndpointConnection";
import { WebAppPublicCertificate } from "./webAppPublicCertificate";
import { WebAppPublicCertificateSlot } from "./webAppPublicCertificateSlot";
import { WebAppRelayServiceConnection } from "./webAppRelayServiceConnection";
import { WebAppRelayServiceConnectionSlot } from "./webAppRelayServiceConnectionSlot";
import { WebAppSiteExtension } from "./webAppSiteExtension";
import { WebAppSiteExtensionSlot } from "./webAppSiteExtensionSlot";
import { WebAppSitePushSettings } from "./webAppSitePushSettings";
import { WebAppSitePushSettingsSlot } from "./webAppSitePushSettingsSlot";
import { WebAppSlot } from "./webAppSlot";
import { WebAppSlotConfigurationNames } from "./webAppSlotConfigurationNames";
import { WebAppSourceControl } from "./webAppSourceControl";
import { WebAppSourceControlSlot } from "./webAppSourceControlSlot";
import { WebAppSwiftVirtualNetworkConnection } from "./webAppSwiftVirtualNetworkConnection";
import { WebAppSwiftVirtualNetworkConnectionSlot } from "./webAppSwiftVirtualNetworkConnectionSlot";
import { WebAppVnetConnection } from "./webAppVnetConnection";
import { WebAppVnetConnectionSlot } from "./webAppVnetConnectionSlot";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:web/v20200901:AppServiceEnvironment":
                return new AppServiceEnvironment(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:AppServicePlan":
                return new AppServicePlan(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:AppServicePlanRouteForVnet":
                return new AppServicePlanRouteForVnet(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:Certificate":
                return new Certificate(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:StaticSite":
                return new StaticSite(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebApp":
                return new WebApp(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppApplicationSettings":
                return new WebAppApplicationSettings(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppApplicationSettingsSlot":
                return new WebAppApplicationSettingsSlot(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppAuthSettings":
                return new WebAppAuthSettings(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppAuthSettingsSlot":
                return new WebAppAuthSettingsSlot(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppAuthSettingsV2":
                return new WebAppAuthSettingsV2(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppAuthSettingsV2Slot":
                return new WebAppAuthSettingsV2Slot(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppAzureStorageAccounts":
                return new WebAppAzureStorageAccounts(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppAzureStorageAccountsSlot":
                return new WebAppAzureStorageAccountsSlot(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppBackupConfiguration":
                return new WebAppBackupConfiguration(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppBackupConfigurationSlot":
                return new WebAppBackupConfigurationSlot(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppConnectionStrings":
                return new WebAppConnectionStrings(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppConnectionStringsSlot":
                return new WebAppConnectionStringsSlot(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppDeployment":
                return new WebAppDeployment(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppDeploymentSlot":
                return new WebAppDeploymentSlot(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppDiagnosticLogsConfiguration":
                return new WebAppDiagnosticLogsConfiguration(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppDomainOwnershipIdentifier":
                return new WebAppDomainOwnershipIdentifier(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppDomainOwnershipIdentifierSlot":
                return new WebAppDomainOwnershipIdentifierSlot(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppFunction":
                return new WebAppFunction(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppHostNameBinding":
                return new WebAppHostNameBinding(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppHostNameBindingSlot":
                return new WebAppHostNameBindingSlot(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppHybridConnection":
                return new WebAppHybridConnection(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppHybridConnectionSlot":
                return new WebAppHybridConnectionSlot(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppInstanceFunctionSlot":
                return new WebAppInstanceFunctionSlot(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppMetadata":
                return new WebAppMetadata(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppMetadataSlot":
                return new WebAppMetadataSlot(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppPremierAddOn":
                return new WebAppPremierAddOn(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppPremierAddOnSlot":
                return new WebAppPremierAddOnSlot(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppPrivateEndpointConnection":
                return new WebAppPrivateEndpointConnection(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppPublicCertificate":
                return new WebAppPublicCertificate(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppPublicCertificateSlot":
                return new WebAppPublicCertificateSlot(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppRelayServiceConnection":
                return new WebAppRelayServiceConnection(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppRelayServiceConnectionSlot":
                return new WebAppRelayServiceConnectionSlot(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppSiteExtension":
                return new WebAppSiteExtension(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppSiteExtensionSlot":
                return new WebAppSiteExtensionSlot(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppSitePushSettings":
                return new WebAppSitePushSettings(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppSitePushSettingsSlot":
                return new WebAppSitePushSettingsSlot(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppSlot":
                return new WebAppSlot(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppSlotConfigurationNames":
                return new WebAppSlotConfigurationNames(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppSourceControl":
                return new WebAppSourceControl(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppSourceControlSlot":
                return new WebAppSourceControlSlot(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppSwiftVirtualNetworkConnection":
                return new WebAppSwiftVirtualNetworkConnection(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppSwiftVirtualNetworkConnectionSlot":
                return new WebAppSwiftVirtualNetworkConnectionSlot(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppVnetConnection":
                return new WebAppVnetConnection(name, <any>undefined, { urn })
            case "azure-native:web/v20200901:WebAppVnetConnectionSlot":
                return new WebAppVnetConnectionSlot(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "web/v20200901", _module)
