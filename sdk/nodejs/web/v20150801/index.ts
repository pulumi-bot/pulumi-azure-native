// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./certificate";
export * from "./certificateCsr";
export * from "./getCertificate";
export * from "./getCertificateCsr";
export * from "./getHostingEnvironment";
export * from "./getManagedHostingEnvironment";
export * from "./getServerFarm";
export * from "./getSite";
export * from "./getSiteDeployment";
export * from "./getSiteDeploymentSlot";
export * from "./getSiteHostNameBinding";
export * from "./getSiteHostNameBindingSlot";
export * from "./getSiteInstanceDeployment";
export * from "./getSiteInstanceDeploymentSlot";
export * from "./getSiteLogsConfig";
export * from "./getSiteRelayServiceConnection";
export * from "./getSiteRelayServiceConnectionSlot";
export * from "./getSiteSlot";
export * from "./getSiteSlotConfigNames";
export * from "./getSiteSourceControl";
export * from "./getSiteSourceControlSlot";
export * from "./getSiteVNETConnection";
export * from "./getSiteVNETConnectionSlot";
export * from "./hostingEnvironment";
export * from "./listSiteAppSettings";
export * from "./listSiteAppSettingsSlot";
export * from "./listSiteAuthSettings";
export * from "./listSiteAuthSettingsSlot";
export * from "./listSiteBackupConfiguration";
export * from "./listSiteBackupConfigurationSlot";
export * from "./listSiteBackupStatusSecrets";
export * from "./listSiteBackupStatusSecretsSlot";
export * from "./listSiteConnectionStrings";
export * from "./listSiteConnectionStringsSlot";
export * from "./listSiteMetadata";
export * from "./listSiteMetadataSlot";
export * from "./listSitePublishingCredentials";
export * from "./listSitePublishingCredentialsSlot";
export * from "./managedHostingEnvironment";
export * from "./serverFarm";
export * from "./serverFarmRouteForVnet";
export * from "./site";
export * from "./siteAppSettings";
export * from "./siteAppSettingsSlot";
export * from "./siteAuthSettings";
export * from "./siteAuthSettingsSlot";
export * from "./siteBackupConfiguration";
export * from "./siteBackupConfigurationSlot";
export * from "./siteConnectionStrings";
export * from "./siteConnectionStringsSlot";
export * from "./siteDeployment";
export * from "./siteDeploymentSlot";
export * from "./siteHostNameBinding";
export * from "./siteHostNameBindingSlot";
export * from "./siteInstanceDeployment";
export * from "./siteInstanceDeploymentSlot";
export * from "./siteLogsConfig";
export * from "./siteMetadata";
export * from "./siteMetadataSlot";
export * from "./siteRelayServiceConnection";
export * from "./siteRelayServiceConnectionSlot";
export * from "./siteSlot";
export * from "./siteSlotConfigNames";
export * from "./siteSourceControl";
export * from "./siteSourceControlSlot";
export * from "./siteVNETConnection";
export * from "./siteVNETConnectionSlot";

// Export enums:
export * from "../../types/enums/web/v20150801";

// Import resources to register:
import { Certificate } from "./certificate";
import { CertificateCsr } from "./certificateCsr";
import { HostingEnvironment } from "./hostingEnvironment";
import { ManagedHostingEnvironment } from "./managedHostingEnvironment";
import { ServerFarm } from "./serverFarm";
import { ServerFarmRouteForVnet } from "./serverFarmRouteForVnet";
import { Site } from "./site";
import { SiteAppSettings } from "./siteAppSettings";
import { SiteAppSettingsSlot } from "./siteAppSettingsSlot";
import { SiteAuthSettings } from "./siteAuthSettings";
import { SiteAuthSettingsSlot } from "./siteAuthSettingsSlot";
import { SiteBackupConfiguration } from "./siteBackupConfiguration";
import { SiteBackupConfigurationSlot } from "./siteBackupConfigurationSlot";
import { SiteConnectionStrings } from "./siteConnectionStrings";
import { SiteConnectionStringsSlot } from "./siteConnectionStringsSlot";
import { SiteDeployment } from "./siteDeployment";
import { SiteDeploymentSlot } from "./siteDeploymentSlot";
import { SiteHostNameBinding } from "./siteHostNameBinding";
import { SiteHostNameBindingSlot } from "./siteHostNameBindingSlot";
import { SiteInstanceDeployment } from "./siteInstanceDeployment";
import { SiteInstanceDeploymentSlot } from "./siteInstanceDeploymentSlot";
import { SiteLogsConfig } from "./siteLogsConfig";
import { SiteMetadata } from "./siteMetadata";
import { SiteMetadataSlot } from "./siteMetadataSlot";
import { SiteRelayServiceConnection } from "./siteRelayServiceConnection";
import { SiteRelayServiceConnectionSlot } from "./siteRelayServiceConnectionSlot";
import { SiteSlot } from "./siteSlot";
import { SiteSlotConfigNames } from "./siteSlotConfigNames";
import { SiteSourceControl } from "./siteSourceControl";
import { SiteSourceControlSlot } from "./siteSourceControlSlot";
import { SiteVNETConnection } from "./siteVNETConnection";
import { SiteVNETConnectionSlot } from "./siteVNETConnectionSlot";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:web/v20150801:Certificate":
                return new Certificate(name, <any>undefined, { urn })
            case "azure-native:web/v20150801:CertificateCsr":
                return new CertificateCsr(name, <any>undefined, { urn })
            case "azure-native:web/v20150801:HostingEnvironment":
                return new HostingEnvironment(name, <any>undefined, { urn })
            case "azure-native:web/v20150801:ManagedHostingEnvironment":
                return new ManagedHostingEnvironment(name, <any>undefined, { urn })
            case "azure-native:web/v20150801:ServerFarm":
                return new ServerFarm(name, <any>undefined, { urn })
            case "azure-native:web/v20150801:ServerFarmRouteForVnet":
                return new ServerFarmRouteForVnet(name, <any>undefined, { urn })
            case "azure-native:web/v20150801:Site":
                return new Site(name, <any>undefined, { urn })
            case "azure-native:web/v20150801:SiteAppSettings":
                return new SiteAppSettings(name, <any>undefined, { urn })
            case "azure-native:web/v20150801:SiteAppSettingsSlot":
                return new SiteAppSettingsSlot(name, <any>undefined, { urn })
            case "azure-native:web/v20150801:SiteAuthSettings":
                return new SiteAuthSettings(name, <any>undefined, { urn })
            case "azure-native:web/v20150801:SiteAuthSettingsSlot":
                return new SiteAuthSettingsSlot(name, <any>undefined, { urn })
            case "azure-native:web/v20150801:SiteBackupConfiguration":
                return new SiteBackupConfiguration(name, <any>undefined, { urn })
            case "azure-native:web/v20150801:SiteBackupConfigurationSlot":
                return new SiteBackupConfigurationSlot(name, <any>undefined, { urn })
            case "azure-native:web/v20150801:SiteConnectionStrings":
                return new SiteConnectionStrings(name, <any>undefined, { urn })
            case "azure-native:web/v20150801:SiteConnectionStringsSlot":
                return new SiteConnectionStringsSlot(name, <any>undefined, { urn })
            case "azure-native:web/v20150801:SiteDeployment":
                return new SiteDeployment(name, <any>undefined, { urn })
            case "azure-native:web/v20150801:SiteDeploymentSlot":
                return new SiteDeploymentSlot(name, <any>undefined, { urn })
            case "azure-native:web/v20150801:SiteHostNameBinding":
                return new SiteHostNameBinding(name, <any>undefined, { urn })
            case "azure-native:web/v20150801:SiteHostNameBindingSlot":
                return new SiteHostNameBindingSlot(name, <any>undefined, { urn })
            case "azure-native:web/v20150801:SiteInstanceDeployment":
                return new SiteInstanceDeployment(name, <any>undefined, { urn })
            case "azure-native:web/v20150801:SiteInstanceDeploymentSlot":
                return new SiteInstanceDeploymentSlot(name, <any>undefined, { urn })
            case "azure-native:web/v20150801:SiteLogsConfig":
                return new SiteLogsConfig(name, <any>undefined, { urn })
            case "azure-native:web/v20150801:SiteMetadata":
                return new SiteMetadata(name, <any>undefined, { urn })
            case "azure-native:web/v20150801:SiteMetadataSlot":
                return new SiteMetadataSlot(name, <any>undefined, { urn })
            case "azure-native:web/v20150801:SiteRelayServiceConnection":
                return new SiteRelayServiceConnection(name, <any>undefined, { urn })
            case "azure-native:web/v20150801:SiteRelayServiceConnectionSlot":
                return new SiteRelayServiceConnectionSlot(name, <any>undefined, { urn })
            case "azure-native:web/v20150801:SiteSlot":
                return new SiteSlot(name, <any>undefined, { urn })
            case "azure-native:web/v20150801:SiteSlotConfigNames":
                return new SiteSlotConfigNames(name, <any>undefined, { urn })
            case "azure-native:web/v20150801:SiteSourceControl":
                return new SiteSourceControl(name, <any>undefined, { urn })
            case "azure-native:web/v20150801:SiteSourceControlSlot":
                return new SiteSourceControlSlot(name, <any>undefined, { urn })
            case "azure-native:web/v20150801:SiteVNETConnection":
                return new SiteVNETConnection(name, <any>undefined, { urn })
            case "azure-native:web/v20150801:SiteVNETConnectionSlot":
                return new SiteVNETConnectionSlot(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "web/v20150801", _module)
