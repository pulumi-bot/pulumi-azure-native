// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataBox.V20200401.Outputs
{

    [OutputType]
    public sealed class DataImportDetailsResponseResult
    {
        /// <summary>
        /// Account details of the data to be transferred
        /// </summary>
        public readonly Outputs.DataAccountDetailsResponseResult AccountDetails;

        [OutputConstructor]
        private DataImportDetailsResponseResult(Outputs.DataAccountDetailsResponseResult accountDetails)
        {
            AccountDetails = accountDetails;
        }
    }
}