package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-operations-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "data-platform-api-operations-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"data-platform-api-operations-creates-rmq-kube/sub_func_complementer"
	"encoding/json"

	"golang.org/x/xerrors"
)

func ConvertToHeaderCreates(subfuncSDC *sub_func_complementer.SDC) (*[]Header, error) {
	headers := make([]Header, 0)

	for _, data := range *subfuncSDC.Message.Header {
		header, err := TypeConverter[*Header](data)
		if err != nil {
			return nil, err
		}

		headers = append(headers, *header)
	}

	return &headers, nil
}

func ConvertToItemCreates(subfuncSDC *sub_func_complementer.SDC) (*[]Item, error) {
	items := make([]Item, 0)

	for _, data := range *subfuncSDC.Message.Item {
		item, err := TypeConverter[*Item](data)
		if err != nil {
			return nil, err
		}

		items = append(items, *item)
	}

	return &items, nil
}

func ConvertToItemOperationCreates(subfuncSDC *sub_func_complementer.SDC) (*[]ItemOperation, error) {
	itemOperations := make([]ItemOperation, 0)

	for _, data := range *subfuncSDC.Message.ItemOperation {
		itemOperation, err := TypeConverter[*ItemOperation](data)
		if err != nil {
			return nil, err
		}

		itemOperations = append(itemOperations, *itemOperation)
	}

	return &itemOperations, nil
}

func ConvertToItemOperationComponentCreates(subfuncSDC *sub_func_complementer.SDC) (*[]ItemOperationComponent, error) {
	itemOperationComponents := make([]ItemOperationComponent, 0)

	for _, data := range *subfuncSDC.Message.ItemOperationComponent {
		itemOperationComponent, err := TypeConverter[*ItemOperationComponent](data)
		if err != nil {
			return nil, err
		}

		itemOperationComponents = append(itemOperationComponents, *itemOperationComponent)
	}

	return &itemOperationComponents, nil
}

func ConvertToHeaderUpdates(headerData dpfm_api_input_reader.Header) (*[]Header, error) {
	headers := make([]Header, 0)
	data := headerData

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}
	headers = append(headers, *header)

	return &headers, nil
}

func ConvertToItemUpdates(itemUpdates *[]dpfm_api_processing_formatter.ItemUpdates) (*[]Item, error) {
	items := make([]Item, 0)

	for _, data := range *itemUpdates {
		item, err := TypeConverter[*Item](data)
		if err != nil {
			return nil, err
		}

		items = append(items, *item)
	}

	return &items, nil
}

func ConvertToItemOperationUpdates(itemOperationUpdates *[]dpfm_api_processing_formatter.ItemOperationUpdates) (*[]ItemOperation, error) {
	itemOperations := make([]ItemOperation, 0)

	for _, data := range *itemOperationUpdates {
		itemOperation, err := TypeConverter[*ItemOperation](data)
		if err != nil {
			return nil, err
		}

		itemOperations = append(itemOperations, *itemOperation)
	}

	return &itemOperations, nil
}

func ConvertToItemOperationComponentUpdates(itemOperationComponentUpdates *[]dpfm_api_processing_formatter.ItemOperationComponentUpdates) (*[]ItemOperationComponent, error) {
	itemOperationComponents := make([]ItemOperationComponent, 0)

	for _, data := range *itemOperationComponentUpdates {
		itemOperationComponent, err := TypeConverter[*ItemOperationComponent](data)
		if err != nil {
			return nil, err
		}

		itemOperationComponents = append(itemOperationComponents, *itemOperationComponent)
	}

	return &itemOperationComponents, nil
}


func TypeConverter[T any](data interface{}) (T, error) {
	var dist T
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}
