﻿using GAPixel;
using GoogleAnalyticsTracker.Core;
using GoogleAnalyticsTracker.Core.TrackerParameters;
using GoogleAnalyticsTracker.Simple;
using Microsoft.Deployment.WindowsInstaller;
using System;
using System.Threading.Tasks;


namespace ActiveState
{
    public sealed class TrackerSingleton
	{
        private static readonly Lazy<TrackerSingleton> lazy = new Lazy<TrackerSingleton>(() => new TrackerSingleton());
        private static string GoogleAnalyticsUserAgent = "UA-118120158-2";

        private readonly SimpleTracker _tracker;
        private readonly string _cid;

        public static TrackerSingleton Instance {  get { return lazy.Value; } }

        public TrackerSingleton()
        {
            var simpleTrackerEnvironment = new SimpleTrackerEnvironment(Environment.OSVersion.Platform.ToString(),
                Environment.OSVersion.Version.ToString(),
                Environment.OSVersion.VersionString);
            this._tracker = new SimpleTracker(GoogleAnalyticsUserAgent, simpleTrackerEnvironment);
            this._cid = GetInfo.GetUniqueId();
        }

        public async Task<TrackingResult> TrackEventAsync(string category, string action, string label, string msiVersion, long value = 1)
        {
            var eventTrackingParameters = new EventTracking
            {
                Category = category,
                Action = action,
                Label = label,
                Value = value,
            };

            eventTrackingParameters.ClientId = this._cid;
            eventTrackingParameters.SetCustomDimensions(new System.Collections.Generic.Dictionary<int, string> { { 1, msiVersion } });

            return await this._tracker.TrackAsync(eventTrackingParameters);
        }

        /// <summary>
        /// Sends a GA event in background (fires and forgets)
        /// </summary>
        /// <description>
        /// The event can fail to be send if the main process gets cancelled before the task finishes.
        /// Use the synchronous version of this command in that case.
        /// </description>
        public void TrackEventInBackground(ActiveState.Logging log, string category, string action, string label, string langVersion, long value=1)
	{
            log.Log("Sending background event {0}/{1}/{2} for cid={3} (custom dimension 1: {4})", category, action, label, this._cid, langVersion);
            Task.Run(() => TrackEventAsync(category, action, label, langVersion, value));
	}

        /// <summary>
        /// Sends a GA event and waits for the request to complete.
        /// </summary>
        public void TrackEventSynchronously(ActiveState.Logging log, string category, string action, string label, string langVersion, long value=1)
        {
            log.Log("Sending event {0}/{1}/{2} for cid={3} (custom dimension 1: {4})", category, action, label, this._cid, langVersion);
            var t = Task.Run(() => TrackEventAsync(category, action, label, langVersion, value));
            t.Wait();
        }
    }
};