// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20180630Preview.Outputs
{

    [OutputType]
    public sealed class ConfigurationParameterResponseResult
    {
        /// <summary>
        /// Name of the configuration parameter.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Value of the configuration parameter.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private ConfigurationParameterResponseResult(
            string name,

            string value)
        {
            Name = name;
            Value = value;
        }
    }
}
