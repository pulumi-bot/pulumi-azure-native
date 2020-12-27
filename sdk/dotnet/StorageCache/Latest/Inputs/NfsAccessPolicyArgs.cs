// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.StorageCache.Latest.Inputs
{

    /// <summary>
    /// A set of rules describing access policies applied to NFSv3 clients of the cache.
    /// </summary>
    public sealed class NfsAccessPolicyArgs : Pulumi.ResourceArgs
    {
        [Input("accessRules", required: true)]
        private InputList<Inputs.NfsAccessRuleArgs>? _accessRules;

        /// <summary>
        /// The set of rules describing client accesses allowed under this policy.
        /// </summary>
        public InputList<Inputs.NfsAccessRuleArgs> AccessRules
        {
            get => _accessRules ?? (_accessRules = new InputList<Inputs.NfsAccessRuleArgs>());
            set => _accessRules = value;
        }

        /// <summary>
        /// Name identifying this policy. Access Policy names are not case sensitive.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public NfsAccessPolicyArgs()
        {
        }
    }
}
