// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Sql.V20140401.Outputs
{

    [OutputType]
    public sealed class ServerPropertiesResponseResult
    {
        /// <summary>
        /// Administrator username for the server. Can only be specified when the server is being created (and is required for creation).
        /// </summary>
        public readonly string? AdministratorLogin;
        /// <summary>
        /// The administrator login password (required for server creation).
        /// </summary>
        public readonly string? AdministratorLoginPassword;
        /// <summary>
        /// The display name of the Azure Active Directory object with admin permissions on this server. Legacy parameter, always null. To check for Active Directory admin, query .../servers/{serverName}/administrators
        /// </summary>
        public readonly string ExternalAdministratorLogin;
        /// <summary>
        /// The ID of the Active Azure Directory object with admin permissions on this server. Legacy parameter, always null. To check for Active Directory admin, query .../servers/{serverName}/administrators.
        /// </summary>
        public readonly string ExternalAdministratorSid;
        /// <summary>
        /// The fully qualified domain name of the server.
        /// </summary>
        public readonly string FullyQualifiedDomainName;
        /// <summary>
        /// The state of the server.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The version of the server.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private ServerPropertiesResponseResult(
            string? administratorLogin,

            string? administratorLoginPassword,

            string externalAdministratorLogin,

            string externalAdministratorSid,

            string fullyQualifiedDomainName,

            string state,

            string? version)
        {
            AdministratorLogin = administratorLogin;
            AdministratorLoginPassword = administratorLoginPassword;
            ExternalAdministratorLogin = externalAdministratorLogin;
            ExternalAdministratorSid = externalAdministratorSid;
            FullyQualifiedDomainName = fullyQualifiedDomainName;
            State = state;
            Version = version;
        }
    }
}