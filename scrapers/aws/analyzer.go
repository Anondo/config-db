package aws

import (
	"fmt"
	"time"

	"github.com/flanksource/commons/console"
	"github.com/flanksource/commons/logger"
	v1 "github.com/flanksource/confighub/api/v1"
)

func EC2InstanceAnalyzer(configs []v1.ScrapeResult) v1.AnalysisResult {

	result := v1.AnalysisResult{
		Analyzer: "ec2-instance",
	}
	for _, config := range configs {
		switch config.Config.(type) {
		case Instance:
			host := config.Config.(Instance)
			state := ""
			if host.PatchState != nil {
				if host.PatchState.FailedCount > 0 {
					state += fmt.Sprintf(" failed=%d", console.Redf("%d", host.PatchState.FailedCount))
				}
				if (host.PatchState.InstalledCount) > 0 {
					state += fmt.Sprintf(" installed=%d", host.PatchState.InstalledCount)
				}
				if host.PatchState.MissingCount > 0 {
					state += fmt.Sprintf(" missing=%d", console.Redf("%d", host.PatchState.MissingCount))
				}
				if host.PatchState.OperationEndTime != nil {
					state += " end=" + time.Since(*host.PatchState.OperationEndTime).String()
				}
			} else {
				state += console.Redf(" no patch state")
			}
			logger.Infof("[%s/%s] os=%s %s", host.GetHostname(), host.GetId(), host.GetPlatform(), state)

			for name, compliance := range host.Compliance {
				if compliance.ComplianceType != "COMPLIANT" {
					result.Messages = append(result.Messages, fmt.Sprintf("[%s/%s] %s - %s: %s", host.GetHostname(), host.GetId(), name, compliance.ComplianceType, compliance.Annotation))
				}
			}
		}
	}

	return result
}
