package responses

type ToItem struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			MaintenanceNotification       string `json:"MaintenanceNotification"`
			MaintenanceNotificationItem   string `json:"MaintenanceNotificationItem"`
			MaintNotifItemText            string `json:"MaintNotifItemText"`
			MaintNotifDamageCodeGroup     string `json:"MaintNotifDamageCodeGroup"`
			MaintNotifDamageCodeGroupName string `json:"MaintNotifDamageCodeGroupName"`
			MaintNotificationDamageCode   string `json:"MaintNotificationDamageCode"`
			MaintNotifDamageCodeName      string `json:"MaintNotifDamageCodeName"`
			MaintNotifObjPrtCodeGroup     string `json:"MaintNotifObjPrtCodeGroup"`
			MaintNotifObjPrtCodeGroupName string `json:"MaintNotifObjPrtCodeGroupName"`
			MaintNotifObjPrtCode          string `json:"MaintNotifObjPrtCode"`
			MaintNotifObjPrtCodeName      string `json:"MaintNotifObjPrtCodeName"`
			IsDeleted                     bool   `json:"IsDeleted"`
			ToItemCause struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_ItemCause"`
		} `json:"results"`
	} `json:"d"`
}
