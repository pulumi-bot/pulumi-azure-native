// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Insights.V20150501.Outputs
{

    [OutputType]
    public sealed class WorkbookPropertiesResponseResult
    {
        /// <summary>
        /// Workbook category, as defined by the user at creation time.
        /// </summary>
        public readonly string Category;
        /// <summary>
        /// Enum indicating if this workbook definition is owned by a specific user or is shared between all users with access to the Application Insights component.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The user-defined name of the workbook.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Configuration of this particular workbook. Configuration data is a string containing valid JSON
        /// </summary>
        public readonly string SerializedData;
        /// <summary>
        /// Optional resourceId for a source resource.
        /// </summary>
        public readonly string? SourceResourceId;
        /// <summary>
        /// A list of 0 or more tags that are associated with this workbook definition
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// Date and time in UTC of the last modification that was made to this workbook definition.
        /// </summary>
        public readonly string TimeModified;
        /// <summary>
        /// Unique user id of the specific user that owns this workbook.
        /// </summary>
        public readonly string UserId;
        /// <summary>
        /// This instance's version of the data model. This can change as new features are added that can be marked workbook.
        /// </summary>
        public readonly string? Version;
        /// <summary>
        /// Internally assigned unique id of the workbook definition.
        /// </summary>
        public readonly string WorkbookId;

        [OutputConstructor]
        private WorkbookPropertiesResponseResult(
            string category,

            string kind,

            string name,

            string serializedData,

            string? sourceResourceId,

            ImmutableArray<string> tags,

            string timeModified,

            string userId,

            string? version,

            string workbookId)
        {
            Category = category;
            Kind = kind;
            Name = name;
            SerializedData = serializedData;
            SourceResourceId = sourceResourceId;
            Tags = tags;
            TimeModified = timeModified;
            UserId = userId;
            Version = version;
            WorkbookId = workbookId;
        }
    }
}