// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.V20150801.Outputs
{

    [OutputType]
    public sealed class SiteConfigResponsePropertiesResult
    {
        /// <summary>
        /// Always On
        /// </summary>
        public readonly bool? AlwaysOn;
        /// <summary>
        /// Information about the formal API definition for the web app.
        /// </summary>
        public readonly Outputs.ApiDefinitionInfoResponseResult? ApiDefinition;
        /// <summary>
        /// App Command Line to launch
        /// </summary>
        public readonly string? AppCommandLine;
        /// <summary>
        /// Application Settings
        /// </summary>
        public readonly ImmutableArray<Outputs.NameValuePairResponseResult> AppSettings;
        /// <summary>
        /// Auto heal enabled
        /// </summary>
        public readonly bool? AutoHealEnabled;
        /// <summary>
        /// Auto heal rules
        /// </summary>
        public readonly Outputs.AutoHealRulesResponseResult? AutoHealRules;
        /// <summary>
        /// Auto swap slot name
        /// </summary>
        public readonly string? AutoSwapSlotName;
        /// <summary>
        /// Connection strings
        /// </summary>
        public readonly ImmutableArray<Outputs.ConnStringInfoResponseResult> ConnectionStrings;
        /// <summary>
        /// Cross-Origin Resource Sharing (CORS) settings.
        /// </summary>
        public readonly Outputs.CorsSettingsResponseResult? Cors;
        /// <summary>
        /// Default documents
        /// </summary>
        public readonly ImmutableArray<string> DefaultDocuments;
        /// <summary>
        /// Detailed error logging enabled
        /// </summary>
        public readonly bool? DetailedErrorLoggingEnabled;
        /// <summary>
        /// Document root
        /// </summary>
        public readonly string? DocumentRoot;
        /// <summary>
        /// This is work around for polymorphic types
        /// </summary>
        public readonly Outputs.ExperimentsResponseResult? Experiments;
        /// <summary>
        /// Handler mappings
        /// </summary>
        public readonly ImmutableArray<Outputs.HandlerMappingResponseResult> HandlerMappings;
        /// <summary>
        /// HTTP logging Enabled
        /// </summary>
        public readonly bool? HttpLoggingEnabled;
        /// <summary>
        /// Ip Security restrictions
        /// </summary>
        public readonly ImmutableArray<Outputs.IpSecurityRestrictionResponseResult> IpSecurityRestrictions;
        /// <summary>
        /// Java container
        /// </summary>
        public readonly string? JavaContainer;
        /// <summary>
        /// Java container version
        /// </summary>
        public readonly string? JavaContainerVersion;
        /// <summary>
        /// Java version
        /// </summary>
        public readonly string? JavaVersion;
        /// <summary>
        /// Site limits
        /// </summary>
        public readonly Outputs.SiteLimitsResponseResult? Limits;
        /// <summary>
        /// Site load balancing
        /// </summary>
        public readonly string? LoadBalancing;
        /// <summary>
        /// Local mysql enabled
        /// </summary>
        public readonly bool? LocalMySqlEnabled;
        /// <summary>
        /// HTTP Logs Directory size limit
        /// </summary>
        public readonly int? LogsDirectorySizeLimit;
        /// <summary>
        /// Managed pipeline mode
        /// </summary>
        public readonly string? ManagedPipelineMode;
        /// <summary>
        /// Site Metadata
        /// </summary>
        public readonly ImmutableArray<Outputs.NameValuePairResponseResult> Metadata;
        /// <summary>
        /// Net Framework Version
        /// </summary>
        public readonly string? NetFrameworkVersion;
        /// <summary>
        /// Version of Node
        /// </summary>
        public readonly string? NodeVersion;
        /// <summary>
        /// Number of workers
        /// </summary>
        public readonly int? NumberOfWorkers;
        /// <summary>
        /// Version of PHP
        /// </summary>
        public readonly string? PhpVersion;
        /// <summary>
        /// Publishing password
        /// </summary>
        public readonly string? PublishingPassword;
        /// <summary>
        /// Publishing user name
        /// </summary>
        public readonly string? PublishingUsername;
        /// <summary>
        /// Version of Python
        /// </summary>
        public readonly string? PythonVersion;
        /// <summary>
        /// Remote Debugging Enabled
        /// </summary>
        public readonly bool? RemoteDebuggingEnabled;
        /// <summary>
        /// Remote Debugging Version
        /// </summary>
        public readonly string? RemoteDebuggingVersion;
        /// <summary>
        /// Enable request tracing
        /// </summary>
        public readonly bool? RequestTracingEnabled;
        /// <summary>
        /// Request tracing expiration time
        /// </summary>
        public readonly string? RequestTracingExpirationTime;
        /// <summary>
        /// SCM type
        /// </summary>
        public readonly string? ScmType;
        /// <summary>
        /// Tracing options
        /// </summary>
        public readonly string? TracingOptions;
        /// <summary>
        /// Use 32 bit worker process
        /// </summary>
        public readonly bool? Use32BitWorkerProcess;
        /// <summary>
        /// Virtual applications
        /// </summary>
        public readonly ImmutableArray<Outputs.VirtualApplicationResponseResult> VirtualApplications;
        /// <summary>
        /// Vnet name
        /// </summary>
        public readonly string? VnetName;
        /// <summary>
        /// Web socket enabled.
        /// </summary>
        public readonly bool? WebSocketsEnabled;

        [OutputConstructor]
        private SiteConfigResponsePropertiesResult(
            bool? alwaysOn,

            Outputs.ApiDefinitionInfoResponseResult? apiDefinition,

            string? appCommandLine,

            ImmutableArray<Outputs.NameValuePairResponseResult> appSettings,

            bool? autoHealEnabled,

            Outputs.AutoHealRulesResponseResult? autoHealRules,

            string? autoSwapSlotName,

            ImmutableArray<Outputs.ConnStringInfoResponseResult> connectionStrings,

            Outputs.CorsSettingsResponseResult? cors,

            ImmutableArray<string> defaultDocuments,

            bool? detailedErrorLoggingEnabled,

            string? documentRoot,

            Outputs.ExperimentsResponseResult? experiments,

            ImmutableArray<Outputs.HandlerMappingResponseResult> handlerMappings,

            bool? httpLoggingEnabled,

            ImmutableArray<Outputs.IpSecurityRestrictionResponseResult> ipSecurityRestrictions,

            string? javaContainer,

            string? javaContainerVersion,

            string? javaVersion,

            Outputs.SiteLimitsResponseResult? limits,

            string? loadBalancing,

            bool? localMySqlEnabled,

            int? logsDirectorySizeLimit,

            string? managedPipelineMode,

            ImmutableArray<Outputs.NameValuePairResponseResult> metadata,

            string? netFrameworkVersion,

            string? nodeVersion,

            int? numberOfWorkers,

            string? phpVersion,

            string? publishingPassword,

            string? publishingUsername,

            string? pythonVersion,

            bool? remoteDebuggingEnabled,

            string? remoteDebuggingVersion,

            bool? requestTracingEnabled,

            string? requestTracingExpirationTime,

            string? scmType,

            string? tracingOptions,

            bool? use32BitWorkerProcess,

            ImmutableArray<Outputs.VirtualApplicationResponseResult> virtualApplications,

            string? vnetName,

            bool? webSocketsEnabled)
        {
            AlwaysOn = alwaysOn;
            ApiDefinition = apiDefinition;
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
            HandlerMappings = handlerMappings;
            HttpLoggingEnabled = httpLoggingEnabled;
            IpSecurityRestrictions = ipSecurityRestrictions;
            JavaContainer = javaContainer;
            JavaContainerVersion = javaContainerVersion;
            JavaVersion = javaVersion;
            Limits = limits;
            LoadBalancing = loadBalancing;
            LocalMySqlEnabled = localMySqlEnabled;
            LogsDirectorySizeLimit = logsDirectorySizeLimit;
            ManagedPipelineMode = managedPipelineMode;
            Metadata = metadata;
            NetFrameworkVersion = netFrameworkVersion;
            NodeVersion = nodeVersion;
            NumberOfWorkers = numberOfWorkers;
            PhpVersion = phpVersion;
            PublishingPassword = publishingPassword;
            PublishingUsername = publishingUsername;
            PythonVersion = pythonVersion;
            RemoteDebuggingEnabled = remoteDebuggingEnabled;
            RemoteDebuggingVersion = remoteDebuggingVersion;
            RequestTracingEnabled = requestTracingEnabled;
            RequestTracingExpirationTime = requestTracingExpirationTime;
            ScmType = scmType;
            TracingOptions = tracingOptions;
            Use32BitWorkerProcess = use32BitWorkerProcess;
            VirtualApplications = virtualApplications;
            VnetName = vnetName;
            WebSocketsEnabled = webSocketsEnabled;
        }
    }
}