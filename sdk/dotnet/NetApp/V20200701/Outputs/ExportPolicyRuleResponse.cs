// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.NetApp.V20200701.Outputs
{

    [OutputType]
    public sealed class ExportPolicyRuleResponse
    {
        /// <summary>
        /// Client ingress specification as comma separated string with IPv4 CIDRs, IPv4 host addresses and host names
        /// </summary>
        public readonly string? AllowedClients;
        /// <summary>
        /// Allows CIFS protocol
        /// </summary>
        public readonly bool? Cifs;
        /// <summary>
        /// Has root access to volume
        /// </summary>
        public readonly bool? HasRootAccess;
        /// <summary>
        /// Kerberos5 Read only access. To be use with swagger version 2020-05-01 or later
        /// </summary>
        public readonly bool? Kerberos5ReadOnly;
        /// <summary>
        /// Kerberos5 Read and write access. To be use with swagger version 2020-05-01 or later
        /// </summary>
        public readonly bool? Kerberos5ReadWrite;
        /// <summary>
        /// Kerberos5i Read only access. To be use with swagger version 2020-05-01 or later
        /// </summary>
        public readonly bool? Kerberos5iReadOnly;
        /// <summary>
        /// Kerberos5i Read and write access. To be use with swagger version 2020-05-01 or later
        /// </summary>
        public readonly bool? Kerberos5iReadWrite;
        /// <summary>
        /// Kerberos5p Read only access. To be use with swagger version 2020-05-01 or later
        /// </summary>
        public readonly bool? Kerberos5pReadOnly;
        /// <summary>
        /// Kerberos5p Read and write access. To be use with swagger version 2020-05-01 or later
        /// </summary>
        public readonly bool? Kerberos5pReadWrite;
        /// <summary>
        /// Allows NFSv3 protocol. Enable only for NFSv3 type volumes
        /// </summary>
        public readonly bool? Nfsv3;
        /// <summary>
        /// Allows NFSv4.1 protocol. Enable only for NFSv4.1 type volumes
        /// </summary>
        public readonly bool? Nfsv41;
        /// <summary>
        /// Order index
        /// </summary>
        public readonly int? RuleIndex;
        /// <summary>
        /// Read only access
        /// </summary>
        public readonly bool? UnixReadOnly;
        /// <summary>
        /// Read and write access
        /// </summary>
        public readonly bool? UnixReadWrite;

        [OutputConstructor]
        private ExportPolicyRuleResponse(
            string? allowedClients,

            bool? cifs,

            bool? hasRootAccess,

            bool? kerberos5ReadOnly,

            bool? kerberos5ReadWrite,

            bool? kerberos5iReadOnly,

            bool? kerberos5iReadWrite,

            bool? kerberos5pReadOnly,

            bool? kerberos5pReadWrite,

            bool? nfsv3,

            bool? nfsv41,

            int? ruleIndex,

            bool? unixReadOnly,

            bool? unixReadWrite)
        {
            AllowedClients = allowedClients;
            Cifs = cifs;
            HasRootAccess = hasRootAccess;
            Kerberos5ReadOnly = kerberos5ReadOnly;
            Kerberos5ReadWrite = kerberos5ReadWrite;
            Kerberos5iReadOnly = kerberos5iReadOnly;
            Kerberos5iReadWrite = kerberos5iReadWrite;
            Kerberos5pReadOnly = kerberos5pReadOnly;
            Kerberos5pReadWrite = kerberos5pReadWrite;
            Nfsv3 = nfsv3;
            Nfsv41 = nfsv41;
            RuleIndex = ruleIndex;
            UnixReadOnly = unixReadOnly;
            UnixReadWrite = unixReadWrite;
        }
    }
}
