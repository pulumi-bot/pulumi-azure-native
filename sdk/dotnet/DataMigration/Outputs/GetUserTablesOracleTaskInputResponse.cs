// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataMigration.Outputs
{

    [OutputType]
    public sealed class GetUserTablesOracleTaskInputResponse
    {
        /// <summary>
        /// Information for connecting to Oracle source
        /// </summary>
        public readonly Outputs.OracleConnectionInfoResponse ConnectionInfo;
        /// <summary>
        /// List of Oracle schemas for which to collect tables
        /// </summary>
        public readonly ImmutableArray<string> SelectedSchemas;

        [OutputConstructor]
        private GetUserTablesOracleTaskInputResponse(
            Outputs.OracleConnectionInfoResponse connectionInfo,

            ImmutableArray<string> selectedSchemas)
        {
            ConnectionInfo = connectionInfo;
            SelectedSchemas = selectedSchemas;
        }
    }
}
