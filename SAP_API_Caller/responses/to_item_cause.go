package responses

type ToItemCause struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			MaintenanceNotification        string `json:"MaintenanceNotification"`
			MaintenanceNotificationItem    string `json:"MaintenanceNotificationItem"`
			MaintenanceNotificationCause   string `json:"MaintenanceNotificationCause"`
			MaintNotifCauseText            string `json:"MaintNotifCauseText"`
			MaintNotifCauseCodeGroup       string `json:"MaintNotifCauseCodeGroup"`
			MaintNotifCauseCodeGroupName   string `json:"MaintNotifCauseCodeGroupName"`
			MaintNotificationCauseCode     string `json:"MaintNotificationCauseCode"`
			MaintNotificationCauseCodeName string `json:"MaintNotificationCauseCodeName"`
			MaintNotificationRootCause     string `json:"MaintNotificationRootCause"`
			MaintNotificationRootCauseText string `json:"MaintNotificationRootCauseText"`
			IsDeleted                      bool   `json:"IsDeleted"`
		} `json:"results"`
	} `json:"d"`
}
