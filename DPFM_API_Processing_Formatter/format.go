package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-operations-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
		Operations:                 data.Operations,
		OperationsText:      		data.OperationsText,
		OperationsStatus: 			data.OperationsStatus,
		ResponsiblePlannerGroup:    data.ResponsiblePlannerGroup,
		PlainLongText:              data.PlainLongText,
		ValidityStartDate:          data.ValidityStartDate,
		ValidityEndDate:            data.ValidityEndDate,
	}
}

func ConvertToItemUpdates(header dpfm_api_input_reader.Header, item dpfm_api_input_reader.Item) *ItemUpdates {
	dataHeader := header
	data := item

	return &ItemUpdates{
		Operations:             	dataHeader.Operations,
		OperationsItem:             data.OperationsItem,
		OperationsText:      		data.OperationsText,
		BillOfMaterial:     		data.BillOfMaterial,
		OperationsUnit:             data.OperationsUnit,
		StandardLotSizeQuantity:    data.StandardLotSizeQuantity,
		MinimumLotSizeQuantity:     data.MinimumLotSizeQuantity,
		MaximumLotSizeQuantity:     data.MaximumLotSizeQuantity,
		PlainLongText:              data.PlainLongText,
		WorkCenter:                 data.WorkCenter,
		ValidityStartDate:       	data.ValidityStartDate,
		ValidityEndDate:         	data.ValidityEndDate,
	}
}

func ConvertToItemOperationUpdates(header dpfm_api_input_reader.Header, item dpfm_api_input_reader.Item, itemOperation dpfm_api_input_reader.ItemOperation) *ItemOperationUpdates {
	dataHeader := header
	dataItem := item
	data := itemOperation

	return &ItemOperationUpdates{
		Operations:		         			dataHeader.Operations,
		OperationsItem:	         			dataItem.OperationsItem,
		OperationID:						data.OperationID,
		Sequence:							data.Sequence,
		SequenceText:						data.SequenceText,
		OperationText:						data.OperationText,
		OperationStatus:					data.OperationStatus,
		ResponsiblePlannerGroup:			data.ResponsiblePlannerGroup,
		OperationUnit:						data.OperationUnit,
		StandardLotSizeQuantity:			data.StandardLotSizeQuantity,
		MinimumLotSizeQuantity:				data.MinimumLotSizeQuantity,
		MaximumLotSizeQuantity:				data.MaximumLotSizeQuantity,
		PlainLongText:						data.PlainLongText,
		WorkCenter:							data.WorkCenter,
		OperationCostingRelevancyType:		data.OperationCostingRelevancyType,
		OperationSetupType:					data.OperationSetupType,
		OperationSetupGroupCategory:		data.OperationSetupGroupCategory,
		OperationSetupGroup:				data.OperationSetupGroup,
		OperationReferenceQuantity:			data.OperationReferenceQuantity,
		MaximumWaitDuration:				data.MaximumWaitDuration,
		StandardWaitDuration:				data.StandardWaitDuration,
		MinimumWaitDuration:				data.MinimumWaitDuration,
		WaitDurationUnit:					data.WaitDurationUnit,
		MaximumQueueDuration:				data.MaximumQueueDuration,
		StandardQueueDuration:				data.StandardQueueDuration,
		MinimumQueueDuration:				data.MinimumQueueDuration,
		QueueDurationUnit:					data.QueueDurationUnit,
		MaximumMoveDuration:				data.MaximumMoveDuration,
		StandardMoveDuration:				data.StandardMoveDuration,
		MinimumMoveDuration:				data.MinimumMoveDuration,
		MoveDurationUnit:					data.MoveDurationUnit,
		StandardDeliveryDuration:			data.StandardDeliveryDuration,
		StandardDeliveryDurationUnit:		data.StandardDeliveryDurationUnit,
		StandardOperationScrapPercent:		data.StandardOperationScrapPercent,
	}
}

func ConvertToItemOperationComponentUpdates(header dpfm_api_input_reader.Header, item dpfm_api_input_reader.Item, itemOperation dpfm_api_input_reader.ItemOperation, itemOperationComponent dpfm_api_input_reader.ItemOperationComponent) *ItemOperationComponentUpdates {
	dataHeader := header
	dataItem := item
	dataItemOperation := itemOperation
	data := itemOperationComponent

	return &ItemOperationComponentUpdates{
		Operations:		         			dataHeader.Operations,
		OperationsItem:	         			dataItem.OperationsItem,
		OperationID:						data.OperationID,
		BillOfMaterial:						data.BillOfMaterial,
		BillOfMaterialItem:					data.BillOfMaterialItem,
		ValidityStartDate:					data.ValidityStartDate,
		ValidityEndDate:					data.ValidityEndDate,
	}
}
