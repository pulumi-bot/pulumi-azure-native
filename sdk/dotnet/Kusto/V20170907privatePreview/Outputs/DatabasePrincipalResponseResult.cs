// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Kusto.V20170907privatePreview.Outputs
{

    [OutputType]
    public sealed class DatabasePrincipalResponseResult
    {
        /// <summary>
        /// Application id - relevant only for application principal type.
        /// </summary>
        public readonly string? AppId;
        /// <summary>
        /// Database principal email if exists.
        /// </summary>
        public readonly string? Email;
        /// <summary>
        /// Database principal fully qualified name.
        /// </summary>
        public readonly string? Fqn;
        /// <summary>
        /// Database principal name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Database principal role.
        /// </summary>
        public readonly string Role;
        /// <summary>
        /// Database principal type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private DatabasePrincipalResponseResult(
            string? appId,

            string? email,

            string? fqn,

            string name,

            string role,

            string type)
        {
            AppId = appId;
            Email = email;
            Fqn = fqn;
            Name = name;
            Role = role;
            Type = type;
        }
    }
}
