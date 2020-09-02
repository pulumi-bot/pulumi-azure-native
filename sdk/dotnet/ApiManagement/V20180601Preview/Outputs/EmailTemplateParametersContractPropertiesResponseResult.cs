// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20180601Preview.Outputs
{

    [OutputType]
    public sealed class EmailTemplateParametersContractPropertiesResponseResult
    {
        /// <summary>
        /// Template parameter description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Template parameter name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Template parameter title.
        /// </summary>
        public readonly string? Title;

        [OutputConstructor]
        private EmailTemplateParametersContractPropertiesResponseResult(
            string? description,

            string? name,

            string? title)
        {
            Description = description;
            Name = name;
            Title = title;
        }
    }
}
