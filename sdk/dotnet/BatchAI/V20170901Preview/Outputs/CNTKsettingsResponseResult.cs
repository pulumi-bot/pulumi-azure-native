// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.BatchAI.V20170901Preview.Outputs
{

    [OutputType]
    public sealed class CNTKsettingsResponseResult
    {
        public readonly string? CommandLineArgs;
        /// <summary>
        /// This property can be specified only if the languageType is 'BrainScript'.
        /// </summary>
        public readonly string? ConfigFilePath;
        /// <summary>
        /// Valid values are 'BrainScript' or 'Python'.
        /// </summary>
        public readonly string? LanguageType;
        /// <summary>
        /// The default value for this property is equal to nodeCount property
        /// </summary>
        public readonly int? ProcessCount;
        /// <summary>
        /// This property can be specified only if the languageType is 'Python'.
        /// </summary>
        public readonly string? PythonInterpreterPath;
        /// <summary>
        /// This property can be specified only if the languageType is 'Python'.
        /// </summary>
        public readonly string? PythonScriptFilePath;

        [OutputConstructor]
        private CNTKsettingsResponseResult(
            string? commandLineArgs,

            string? configFilePath,

            string? languageType,

            int? processCount,

            string? pythonInterpreterPath,

            string? pythonScriptFilePath)
        {
            CommandLineArgs = commandLineArgs;
            ConfigFilePath = configFilePath;
            LanguageType = languageType;
            ProcessCount = processCount;
            PythonInterpreterPath = pythonInterpreterPath;
            PythonScriptFilePath = pythonScriptFilePath;
        }
    }
}
