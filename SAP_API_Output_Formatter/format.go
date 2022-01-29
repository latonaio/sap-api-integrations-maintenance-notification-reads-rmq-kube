package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-maintenance-notification-reads-rmq-kube/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		header = append(header, Header{
	MaintenanceNotification:        data.MaintenanceNotification,
	MaintNotifInternalID:           data.MaintNotifInternalID,
	NotificationText:               data.NotificationText,
	MaintPriority:                  data.MaintPriority,
	NotificationType:               data.NotificationType,
	NotifProcessingPhase:           data.NotifProcessingPhase,
	NotifProcessingPhaseDesc:       data.NotifProcessingPhaseDesc,
	MaintPriorityDesc:              data.MaintPriorityDesc,
	CreationDate:                   data.CreationDate,
	LastChangeTime:                 data.LastChangeTime,
	LastChangeDate:                 data.LastChangeDate,
	LastChangeDateTime:             data.LastChangeDateTime,
	CreationTime:                   data.CreationTime,
	ReportedByUser:                 data.ReportedByUser,
	ReporterFullName:               data.ReporterFullName,
	PersonResponsible:              data.PersonResponsible,
	MalfunctionEffect:              data.MalfunctionEffect,
	MalfunctionEffectText:          data.MalfunctionEffectText,
	MalfunctionStartDate:           data.MalfunctionStartDate,
	MalfunctionStartTime:           data.MalfunctionStartTime,
	MalfunctionEndDate:             data.MalfunctionEndDate,
	MalfunctionEndTime:             data.MalfunctionEndTime,
	MaintNotificationCatalog:       data.MaintNotificationCatalog,
	MaintNotificationCode:          data.MaintNotificationCode,
	MaintNotificationCodeGroup:     data.MaintNotificationCodeGroup,
	CatalogProfile:                 data.CatalogProfile,
	NotificationCreationDate:       data.NotificationCreationDate,
	NotificationCreationTime:       data.NotificationCreationTime,
	NotificationTimeZone:           data.NotificationTimeZone,
	RequiredStartDate:              data.RequiredStartDate,
	RequiredStartTime:              data.RequiredStartTime,
	RequiredEndDate:                data.RequiredEndDate,
	RequiredEndTime:                data.RequiredEndTime,
	LatestAcceptableCompletionDate: data.LatestAcceptableCompletionDate,
	MaintenanceObjectIsDown:        data.MaintenanceObjectIsDown,
	MaintNotificationLongText:      data.MaintNotificationLongText,
	MaintNotifLongTextForEdit:      data.MaintNotifLongTextForEdit,
	TechnicalObject:                data.TechnicalObject,
	TechObjIsEquipOrFuncnlLoc:      data.TechObjIsEquipOrFuncnlLoc,
	TechnicalObjectLabel:           data.TechnicalObjectLabel,
	MaintenancePlanningPlant:       data.MaintenancePlanningPlant,
	MaintenancePlannerGroup:        data.MaintenancePlannerGroup,
	SuperiorTechnicalObject:        data.SuperiorTechnicalObject,
	SuperiorTechnicalObjectName:    data.SuperiorTechnicalObjectName,
	SuperiorObjIsEquipOrFuncnlLoc:  data.SuperiorObjIsEquipOrFuncnlLoc,
	SuperiorTechnicalObjectLabel:   data.SuperiorTechnicalObjectLabel,
	ManufacturerPartTypeName:       data.ManufacturerPartTypeName,
	TechObjIsEquipOrFuncnlLocDesc:  data.TechObjIsEquipOrFuncnlLocDesc,
	FunctionalLocation:             data.FunctionalLocation,
	FunctionalLocationLabelName:    data.FunctionalLocationLabelName,
	TechnicalObjectDescription:     data.TechnicalObjectDescription,
	AssetLocation:                  data.AssetLocation,
	LocationName:                   data.LocationName,
	BusinessArea:                   data.BusinessArea,
	CompanyCode:                    data.CompanyCode,
	TechnicalObjectCategory:        data.TechnicalObjectCategory,
	TechnicalObjectType:            data.TechnicalObjectType,
	MainWorkCenterPlant:            data.MainWorkCenterPlant,
	MainWorkCenter:                 data.MainWorkCenter,
	PlantName:                      data.PlantName,
	MaintenancePlannerGroupName:    data.MaintenancePlannerGroupName,
	MaintenancePlant:               data.MaintenancePlant,
	LocationDescription:            data.LocationDescription,
	MainWorkCenterText:             data.MainWorkCenterText,
	MainWorkCenterPlantName:        data.MainWorkCenterPlantName,
	MaintenancePlantName:           data.MaintenancePlantName,
	PlantSectionPersonRespName:     data.PlantSectionPersonRespName,
	PersonResponsibleName:          data.PersonResponsibleName,
	MaintenanceOrder:               data.MaintenanceOrder,
	MaintenanceOrderType:           data.MaintenanceOrderType,
	MaintenanceActivityType:        data.MaintenanceActivityType,
	MaintObjDowntimeDurationUnit:   data.MaintObjDowntimeDurationUnit,
	MaintObjectDowntimeDuration:    data.MaintObjectDowntimeDuration,
	MaintenancePlan:                data.MaintenancePlan,
	MaintenanceItem:                data.MaintenanceItem,
	TaskListGroup:                  data.TaskListGroup,
	TaskListGroupCounter:           data.TaskListGroupCounter,
	MaintenancePlanCallNumber:      data.MaintenancePlanCallNumber,
	MaintenanceTaskListType:        data.MaintenanceTaskListType,
	NotificationReferenceDate:      data.NotificationReferenceDate,
	NotificationReferenceTime:      data.NotificationReferenceTime,
	NotificationCompletionDate:     data.NotificationCompletionDate,
	CompletionTime:                 data.CompletionTime,
	AssetRoom:                      data.AssetRoom,
	MaintNotifExtReferenceNumber:   data.MaintNotifExtReferenceNumber,
	MaintNotifRejectionReasonCode:  data.MaintNotifRejectionReasonCode,
	MaintNotifRejectionRsnCodeTxt:  data.MaintNotifRejectionRsnCodeTxt,
	MaintNotifDetectionCatalog:     data.MaintNotifDetectionCatalog,
	MaintNotifDetectionCode:        data.MaintNotifDetectionCode,
	MaintNotifDetectionCodeText:    data.MaintNotifDetectionCodeText,
	MaintNotifDetectionCodeGroup:   data.MaintNotifDetectionCodeGroup,
	MaintNotifDetectionCodeGrpTxt:  data.MaintNotifDetectionCodeGrpTxt,
	MaintNotifProcessPhaseCode:     data.MaintNotifProcessPhaseCode,
	MaintNotifProcessSubPhaseCode:  data.MaintNotifProcessSubPhaseCode,
	ToItem:                         data.ToItem.Deferred.URI,
		})
	}

	return header, nil
}

func ConvertToToItem(raw []byte, l *logger.Logger) ([]ToItem, error) {
	pm := &responses.ToItem{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItem. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toItem := make([]ToItem, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItem = append(toItem, ToItem{
	MaintenanceNotification:       data.MaintenanceNotification,
	MaintenanceNotificationItem:   data.MaintenanceNotificationItem,
	MaintNotifItemText:            data.MaintNotifItemText,
	MaintNotifDamageCodeGroup:     data.MaintNotifDamageCodeGroup,
	MaintNotifDamageCodeGroupName: data.MaintNotifDamageCodeGroupName,
	MaintNotificationDamageCode:   data.MaintNotificationDamageCode,
	MaintNotifDamageCodeName:      data.MaintNotifDamageCodeName,
	MaintNotifObjPrtCodeGroup:     data.MaintNotifObjPrtCodeGroup,
	MaintNotifObjPrtCodeGroupName: data.MaintNotifObjPrtCodeGroupName,
	MaintNotifObjPrtCode:          data.MaintNotifObjPrtCode,
	MaintNotifObjPrtCodeName:      data.MaintNotifObjPrtCodeName,
	IsDeleted:                     data.IsDeleted,
	ToItemCause:                   data.ToItemCause.Deferred.URI,
		})
	}

	return toItem, nil
}

func ConvertToToItemCause(raw []byte, l *logger.Logger) ([]ToItemCause, error) {
	pm := &responses.ToItemCause{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItemCause. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toItemCause := make([]ToItemCause, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItemCause = append(toItemCause, ToItemCause{
	MaintenanceNotification:        data.MaintenanceNotification,
	MaintenanceNotificationItem:    data.MaintenanceNotificationItem,
	MaintenanceNotificationCause:   data.MaintenanceNotificationCause,
	MaintNotifCauseText:            data.MaintNotifCauseText,
	MaintNotifCauseCodeGroup:       data.MaintNotifCauseCodeGroup,
	MaintNotifCauseCodeGroupName:   data.MaintNotifCauseCodeGroupName,
	MaintNotificationCauseCode:     data.MaintNotificationCauseCode,
	MaintNotificationCauseCodeName: data.MaintNotificationCauseCodeName,
	MaintNotificationRootCause:     data.MaintNotificationRootCause,
	MaintNotificationRootCauseText: data.MaintNotificationRootCauseText,
	IsDeleted:                      data.IsDeleted,
		})
	}

	return toItemCause, nil
}
