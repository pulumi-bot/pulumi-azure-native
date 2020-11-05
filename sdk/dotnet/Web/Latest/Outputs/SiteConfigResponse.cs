// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Web.Latest.Outputs
{

    [OutputType]
    public sealed class SiteConfigResponse
    {
        /// <summary>
        /// &lt;code&gt;true&lt;/code&gt; if Always On is enabled; otherwise, &lt;code&gt;false&lt;/code&gt;.
        /// </summary>
        public readonly bool? AlwaysOn;
        /// <summary>
        /// Information about the formal API definition for the app.
        /// </summary>
        public readonly Outputs.ApiDefinitionInfoResponse? ApiDefinition;
        /// <summary>
        /// Azure API management settings linked to the app.
        /// </summary>
        public readonly Outputs.ApiManagementConfigResponse? ApiManagementConfig;
        /// <summary>
        /// App command line to launch.
        /// </summary>
        public readonly string? AppCommandLine;
        /// <summary>
        /// Application settings.
        /// </summary>
        public readonly ImmutableArray<Outputs.NameValuePairResponse> AppSettings;
        /// <summary>
        /// &lt;code&gt;true&lt;/code&gt; if Auto Heal is enabled; otherwise, &lt;code&gt;false&lt;/code&gt;.
        /// </summary>
        public readonly bool? AutoHealEnabled;
        /// <summary>
        /// Auto Heal rules.
        /// </summary>
        public readonly Outputs.AutoHealRulesResponse? AutoHealRules;
        /// <summary>
        /// Auto-swap slot name.
        /// </summary>
        public readonly string? AutoSwapSlotName;
        /// <summary>
        /// Connection strings.
        /// </summary>
        public readonly ImmutableArray<Outputs.ConnStringInfoResponse> ConnectionStrings;
        /// <summary>
        /// Cross-Origin Resource Sharing (CORS) settings.
        /// </summary>
        public readonly Outputs.CorsSettingsResponse? Cors;
        /// <summary>
        /// Default documents.
        /// </summary>
        public readonly ImmutableArray<string> DefaultDocuments;
        /// <summary>
        /// &lt;code&gt;true&lt;/code&gt; if detailed error logging is enabled; otherwise, &lt;code&gt;false&lt;/code&gt;.
        /// </summary>
        public readonly bool? DetailedErrorLoggingEnabled;
        /// <summary>
        /// Document root.
        /// </summary>
        public readonly string? DocumentRoot;
        /// <summary>
        /// This is work around for polymorphic types.
        /// </summary>
        public readonly Outputs.ExperimentsResponse? Experiments;
        /// <summary>
        /// State of FTP / FTPS service
        /// </summary>
        public readonly string? FtpsState;
        /// <summary>
        /// Handler mappings.
        /// </summary>
        public readonly ImmutableArray<Outputs.HandlerMappingResponse> HandlerMappings;
        /// <summary>
        /// Health check path
        /// </summary>
        public readonly string? HealthCheckPath;
        /// <summary>
        /// Http20Enabled: configures a web site to allow clients to connect over http2.0
        /// </summary>
        public readonly bool? Http20Enabled;
        /// <summary>
        /// &lt;code&gt;true&lt;/code&gt; if HTTP logging is enabled; otherwise, &lt;code&gt;false&lt;/code&gt;.
        /// </summary>
        public readonly bool? HttpLoggingEnabled;
        /// <summary>
        /// IP security restrictions for main.
        /// </summary>
        public readonly ImmutableArray<Outputs.IpSecurityRestrictionResponse> IpSecurityRestrictions;
        /// <summary>
        /// Java container.
        /// </summary>
        public readonly string? JavaContainer;
        /// <summary>
        /// Java container version.
        /// </summary>
        public readonly string? JavaContainerVersion;
        /// <summary>
        /// Java version.
        /// </summary>
        public readonly string? JavaVersion;
        /// <summary>
        /// Site limits.
        /// </summary>
        public readonly Outputs.SiteLimitsResponse? Limits;
        /// <summary>
        /// Linux App Framework and version
        /// </summary>
        public readonly string? LinuxFxVersion;
        /// <summary>
        /// Site load balancing.
        /// </summary>
        public readonly string? LoadBalancing;
        /// <summary>
        /// &lt;code&gt;true&lt;/code&gt; to enable local MySQL; otherwise, &lt;code&gt;false&lt;/code&gt;.
        /// </summary>
        public readonly bool? LocalMySqlEnabled;
        /// <summary>
        /// HTTP logs directory size limit.
        /// </summary>
        public readonly int? LogsDirectorySizeLimit;
        /// <summary>
        /// Site MachineKey.
        /// </summary>
        public readonly Outputs.SiteMachineKeyResponse MachineKey;
        /// <summary>
        /// Managed pipeline mode.
        /// </summary>
        public readonly string? ManagedPipelineMode;
        /// <summary>
        /// Managed Service Identity Id
        /// </summary>
        public readonly int? ManagedServiceIdentityId;
        /// <summary>
        /// MinTlsVersion: configures the minimum version of TLS required for SSL requests
        /// </summary>
        public readonly string? MinTlsVersion;
        /// <summary>
        /// .NET Framework version.
        /// </summary>
        public readonly string? NetFrameworkVersion;
        /// <summary>
        /// Version of Node.js.
        /// </summary>
        public readonly string? NodeVersion;
        /// <summary>
        /// Number of workers.
        /// </summary>
        public readonly int? NumberOfWorkers;
        /// <summary>
        /// Version of PHP.
        /// </summary>
        public readonly string? PhpVersion;
        /// <summary>
        /// Version of PowerShell.
        /// </summary>
        public readonly string? PowerShellVersion;
        /// <summary>
        /// Number of preWarmed instances.
        /// This setting only applies to the Consumption and Elastic Plans
        /// </summary>
        public readonly int? PreWarmedInstanceCount;
        /// <summary>
        /// Publishing user name.
        /// </summary>
        public readonly string? PublishingUsername;
        /// <summary>
        /// Push endpoint settings.
        /// </summary>
        public readonly Outputs.PushSettingsResponse? Push;
        /// <summary>
        /// Version of Python.
        /// </summary>
        public readonly string? PythonVersion;
        /// <summary>
        /// &lt;code&gt;true&lt;/code&gt; if remote debugging is enabled; otherwise, &lt;code&gt;false&lt;/code&gt;.
        /// </summary>
        public readonly bool? RemoteDebuggingEnabled;
        /// <summary>
        /// Remote debugging version.
        /// </summary>
        public readonly string? RemoteDebuggingVersion;
        /// <summary>
        /// &lt;code&gt;true&lt;/code&gt; if request tracing is enabled; otherwise, &lt;code&gt;false&lt;/code&gt;.
        /// </summary>
        public readonly bool? RequestTracingEnabled;
        /// <summary>
        /// Request tracing expiration time.
        /// </summary>
        public readonly string? RequestTracingExpirationTime;
        /// <summary>
        /// IP security restrictions for scm.
        /// </summary>
        public readonly ImmutableArray<Outputs.IpSecurityRestrictionResponse> ScmIpSecurityRestrictions;
        /// <summary>
        /// IP security restrictions for scm to use main.
        /// </summary>
        public readonly bool? ScmIpSecurityRestrictionsUseMain;
        /// <summary>
        /// ScmMinTlsVersion: configures the minimum version of TLS required for SSL requests for SCM site
        /// </summary>
        public readonly string? ScmMinTlsVersion;
        /// <summary>
        /// SCM type.
        /// </summary>
        public readonly string? ScmType;
        /// <summary>
        /// Tracing options.
        /// </summary>
        public readonly string? TracingOptions;
        /// <summary>
        /// &lt;code&gt;true&lt;/code&gt; to use 32-bit worker process; otherwise, &lt;code&gt;false&lt;/code&gt;.
        /// </summary>
        public readonly bool? Use32BitWorkerProcess;
        /// <summary>
        /// Virtual applications.
        /// </summary>
        public readonly ImmutableArray<Outputs.VirtualApplicationResponse> VirtualApplications;
        /// <summary>
        /// Virtual Network name.
        /// </summary>
        public readonly string? VnetName;
        /// <summary>
        /// The number of private ports assigned to this app. These will be assigned dynamically on runtime.
        /// </summary>
        public readonly int? VnetPrivatePortsCount;
        /// <summary>
        /// Virtual Network Route All enabled. This causes all outbound traffic to have Virtual Network Security Groups and User Defined Routes applied.
        /// </summary>
        public readonly bool? VnetRouteAllEnabled;
        /// <summary>
        /// &lt;code&gt;true&lt;/code&gt; if WebSocket is enabled; otherwise, &lt;code&gt;false&lt;/code&gt;.
        /// </summary>
        public readonly bool? WebSocketsEnabled;
        /// <summary>
        /// Xenon App Framework and version
        /// </summary>
        public readonly string? WindowsFxVersion;
        /// <summary>
        /// Explicit Managed Service Identity Id
        /// </summary>
        public readonly int? XManagedServiceIdentityId;

        [OutputConstructor]
        private SiteConfigResponse(
            bool? alwaysOn,

            Outputs.ApiDefinitionInfoResponse? apiDefinition,

            Outputs.ApiManagementConfigResponse? apiManagementConfig,

            string? appCommandLine,

            ImmutableArray<Outputs.NameValuePairResponse> appSettings,

            bool? autoHealEnabled,

            Outputs.AutoHealRulesResponse? autoHealRules,

            string? autoSwapSlotName,

            ImmutableArray<Outputs.ConnStringInfoResponse> connectionStrings,

            Outputs.CorsSettingsResponse? cors,

            ImmutableArray<string> defaultDocuments,

            bool? detailedErrorLoggingEnabled,

            string? documentRoot,

            Outputs.ExperimentsResponse? experiments,

            string? ftpsState,

            ImmutableArray<Outputs.HandlerMappingResponse> handlerMappings,

            string? healthCheckPath,

            bool? http20Enabled,

            bool? httpLoggingEnabled,

            ImmutableArray<Outputs.IpSecurityRestrictionResponse> ipSecurityRestrictions,

            string? javaContainer,

            string? javaContainerVersion,

            string? javaVersion,

            Outputs.SiteLimitsResponse? limits,

            string? linuxFxVersion,

            string? loadBalancing,

            bool? localMySqlEnabled,

            int? logsDirectorySizeLimit,

            Outputs.SiteMachineKeyResponse machineKey,

            string? managedPipelineMode,

            int? managedServiceIdentityId,

            string? minTlsVersion,

            string? netFrameworkVersion,

            string? nodeVersion,

            int? numberOfWorkers,

            string? phpVersion,

            string? powerShellVersion,

            int? preWarmedInstanceCount,

            string? publishingUsername,

            Outputs.PushSettingsResponse? push,

            string? pythonVersion,

            bool? remoteDebuggingEnabled,

            string? remoteDebuggingVersion,

            bool? requestTracingEnabled,

            string? requestTracingExpirationTime,

            ImmutableArray<Outputs.IpSecurityRestrictionResponse> scmIpSecurityRestrictions,

            bool? scmIpSecurityRestrictionsUseMain,

            string? scmMinTlsVersion,

            string? scmType,

            string? tracingOptions,

            bool? use32BitWorkerProcess,

            ImmutableArray<Outputs.VirtualApplicationResponse> virtualApplications,

            string? vnetName,

            int? vnetPrivatePortsCount,

            bool? vnetRouteAllEnabled,

            bool? webSocketsEnabled,

            string? windowsFxVersion,

            int? xManagedServiceIdentityId)
        {
            AlwaysOn = alwaysOn;
            ApiDefinition = apiDefinition;
            ApiManagementConfig = apiManagementConfig;
            AppCommandLine = appCommandLine;
            AppSettings = appSettings;
            AutoHealEnabled = autoHealEnabled;
            AutoHealRules = autoHealRules;
            AutoSwapSlotName = autoSwapSlotName;
            ConnectionStrings = connectionStrings;
            Cors = cors;
            DefaultDocuments = defaultDocuments;
            DetailedErrorLoggingEnabled = detailedErrorLoggingEnabled;
            DocumentRoot = documentRoot;
            Experiments = experiments;
            FtpsState = ftpsState;
            HandlerMappings = handlerMappings;
            HealthCheckPath = healthCheckPath;
            Http20Enabled = http20Enabled;
            HttpLoggingEnabled = httpLoggingEnabled;
            IpSecurityRestrictions = ipSecurityRestrictions;
            JavaContainer = javaContainer;
            JavaContainerVersion = javaContainerVersion;
            JavaVersion = javaVersion;
            Limits = limits;
            LinuxFxVersion = linuxFxVersion;
            LoadBalancing = loadBalancing;
            LocalMySqlEnabled = localMySqlEnabled;
            LogsDirectorySizeLimit = logsDirectorySizeLimit;
            MachineKey = machineKey;
            ManagedPipelineMode = managedPipelineMode;
            ManagedServiceIdentityId = managedServiceIdentityId;
            MinTlsVersion = minTlsVersion;
            NetFrameworkVersion = netFrameworkVersion;
            NodeVersion = nodeVersion;
            NumberOfWorkers = numberOfWorkers;
            PhpVersion = phpVersion;
            PowerShellVersion = powerShellVersion;
            PreWarmedInstanceCount = preWarmedInstanceCount;
            PublishingUsername = publishingUsername;
            Push = push;
            PythonVersion = pythonVersion;
            RemoteDebuggingEnabled = remoteDebuggingEnabled;
            RemoteDebuggingVersion = remoteDebuggingVersion;
            RequestTracingEnabled = requestTracingEnabled;
            RequestTracingExpirationTime = requestTracingExpirationTime;
            ScmIpSecurityRestrictions = scmIpSecurityRestrictions;
            ScmIpSecurityRestrictionsUseMain = scmIpSecurityRestrictionsUseMain;
            ScmMinTlsVersion = scmMinTlsVersion;
            ScmType = scmType;
            TracingOptions = tracingOptions;
            Use32BitWorkerProcess = use32BitWorkerProcess;
            VirtualApplications = virtualApplications;
            VnetName = vnetName;
            VnetPrivatePortsCount = vnetPrivatePortsCount;
            VnetRouteAllEnabled = vnetRouteAllEnabled;
            WebSocketsEnabled = webSocketsEnabled;
            WindowsFxVersion = windowsFxVersion;
            XManagedServiceIdentityId = xManagedServiceIdentityId;
        }
    }
}
