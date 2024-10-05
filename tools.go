// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	time "time"
)

type ToolsListRequest struct {
	// This is the maximum number of items to return. Defaults to 100.
	Limit *float64 `json:"-" url:"limit,omitempty"`
	// This will return items where the createdAt is greater than the specified value.
	CreatedAtGt *time.Time `json:"-" url:"createdAtGt,omitempty"`
	// This will return items where the createdAt is less than the specified value.
	CreatedAtLt *time.Time `json:"-" url:"createdAtLt,omitempty"`
	// This will return items where the createdAt is greater than or equal to the specified value.
	CreatedAtGe *time.Time `json:"-" url:"createdAtGe,omitempty"`
	// This will return items where the createdAt is less than or equal to the specified value.
	CreatedAtLe *time.Time `json:"-" url:"createdAtLe,omitempty"`
	// This will return items where the updatedAt is greater than the specified value.
	UpdatedAtGt *time.Time `json:"-" url:"updatedAtGt,omitempty"`
	// This will return items where the updatedAt is less than the specified value.
	UpdatedAtLt *time.Time `json:"-" url:"updatedAtLt,omitempty"`
	// This will return items where the updatedAt is greater than or equal to the specified value.
	UpdatedAtGe *time.Time `json:"-" url:"updatedAtGe,omitempty"`
	// This will return items where the updatedAt is less than or equal to the specified value.
	UpdatedAtLe *time.Time `json:"-" url:"updatedAtLe,omitempty"`
}

type ToolsCreateRequest struct {
	CreateDtmfToolDto         *CreateDtmfToolDto
	CreateEndCallToolDto      *CreateEndCallToolDto
	CreateFunctionToolDto     *CreateFunctionToolDto
	CreateGhlToolDto          *CreateGhlToolDto
	CreateMakeToolDto         *CreateMakeToolDto
	CreateTransferCallToolDto *CreateTransferCallToolDto
	CreateOutputToolDto       *CreateOutputToolDto
}

func (t *ToolsCreateRequest) UnmarshalJSON(data []byte) error {
	valueCreateDtmfToolDto := new(CreateDtmfToolDto)
	if err := json.Unmarshal(data, &valueCreateDtmfToolDto); err == nil {
		t.CreateDtmfToolDto = valueCreateDtmfToolDto
		return nil
	}
	valueCreateEndCallToolDto := new(CreateEndCallToolDto)
	if err := json.Unmarshal(data, &valueCreateEndCallToolDto); err == nil {
		t.CreateEndCallToolDto = valueCreateEndCallToolDto
		return nil
	}
	valueCreateFunctionToolDto := new(CreateFunctionToolDto)
	if err := json.Unmarshal(data, &valueCreateFunctionToolDto); err == nil {
		t.CreateFunctionToolDto = valueCreateFunctionToolDto
		return nil
	}
	valueCreateGhlToolDto := new(CreateGhlToolDto)
	if err := json.Unmarshal(data, &valueCreateGhlToolDto); err == nil {
		t.CreateGhlToolDto = valueCreateGhlToolDto
		return nil
	}
	valueCreateMakeToolDto := new(CreateMakeToolDto)
	if err := json.Unmarshal(data, &valueCreateMakeToolDto); err == nil {
		t.CreateMakeToolDto = valueCreateMakeToolDto
		return nil
	}
	valueCreateTransferCallToolDto := new(CreateTransferCallToolDto)
	if err := json.Unmarshal(data, &valueCreateTransferCallToolDto); err == nil {
		t.CreateTransferCallToolDto = valueCreateTransferCallToolDto
		return nil
	}
	valueCreateOutputToolDto := new(CreateOutputToolDto)
	if err := json.Unmarshal(data, &valueCreateOutputToolDto); err == nil {
		t.CreateOutputToolDto = valueCreateOutputToolDto
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, t)
}

func (t ToolsCreateRequest) MarshalJSON() ([]byte, error) {
	if t.CreateDtmfToolDto != nil {
		return json.Marshal(t.CreateDtmfToolDto)
	}
	if t.CreateEndCallToolDto != nil {
		return json.Marshal(t.CreateEndCallToolDto)
	}
	if t.CreateFunctionToolDto != nil {
		return json.Marshal(t.CreateFunctionToolDto)
	}
	if t.CreateGhlToolDto != nil {
		return json.Marshal(t.CreateGhlToolDto)
	}
	if t.CreateMakeToolDto != nil {
		return json.Marshal(t.CreateMakeToolDto)
	}
	if t.CreateTransferCallToolDto != nil {
		return json.Marshal(t.CreateTransferCallToolDto)
	}
	if t.CreateOutputToolDto != nil {
		return json.Marshal(t.CreateOutputToolDto)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", t)
}

type ToolsCreateRequestVisitor interface {
	VisitCreateDtmfToolDto(*CreateDtmfToolDto) error
	VisitCreateEndCallToolDto(*CreateEndCallToolDto) error
	VisitCreateFunctionToolDto(*CreateFunctionToolDto) error
	VisitCreateGhlToolDto(*CreateGhlToolDto) error
	VisitCreateMakeToolDto(*CreateMakeToolDto) error
	VisitCreateTransferCallToolDto(*CreateTransferCallToolDto) error
	VisitCreateOutputToolDto(*CreateOutputToolDto) error
}

func (t *ToolsCreateRequest) Accept(visitor ToolsCreateRequestVisitor) error {
	if t.CreateDtmfToolDto != nil {
		return visitor.VisitCreateDtmfToolDto(t.CreateDtmfToolDto)
	}
	if t.CreateEndCallToolDto != nil {
		return visitor.VisitCreateEndCallToolDto(t.CreateEndCallToolDto)
	}
	if t.CreateFunctionToolDto != nil {
		return visitor.VisitCreateFunctionToolDto(t.CreateFunctionToolDto)
	}
	if t.CreateGhlToolDto != nil {
		return visitor.VisitCreateGhlToolDto(t.CreateGhlToolDto)
	}
	if t.CreateMakeToolDto != nil {
		return visitor.VisitCreateMakeToolDto(t.CreateMakeToolDto)
	}
	if t.CreateTransferCallToolDto != nil {
		return visitor.VisitCreateTransferCallToolDto(t.CreateTransferCallToolDto)
	}
	if t.CreateOutputToolDto != nil {
		return visitor.VisitCreateOutputToolDto(t.CreateOutputToolDto)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", t)
}

type ToolsCreateResponse struct {
	DtmfTool         *DtmfTool
	EndCallTool      *EndCallTool
	FunctionTool     *FunctionTool
	GhlTool          *GhlTool
	MakeTool         *MakeTool
	TransferCallTool *TransferCallTool
	OutputTool       *OutputTool
}

func (t *ToolsCreateResponse) UnmarshalJSON(data []byte) error {
	valueDtmfTool := new(DtmfTool)
	if err := json.Unmarshal(data, &valueDtmfTool); err == nil {
		t.DtmfTool = valueDtmfTool
		return nil
	}
	valueEndCallTool := new(EndCallTool)
	if err := json.Unmarshal(data, &valueEndCallTool); err == nil {
		t.EndCallTool = valueEndCallTool
		return nil
	}
	valueFunctionTool := new(FunctionTool)
	if err := json.Unmarshal(data, &valueFunctionTool); err == nil {
		t.FunctionTool = valueFunctionTool
		return nil
	}
	valueGhlTool := new(GhlTool)
	if err := json.Unmarshal(data, &valueGhlTool); err == nil {
		t.GhlTool = valueGhlTool
		return nil
	}
	valueMakeTool := new(MakeTool)
	if err := json.Unmarshal(data, &valueMakeTool); err == nil {
		t.MakeTool = valueMakeTool
		return nil
	}
	valueTransferCallTool := new(TransferCallTool)
	if err := json.Unmarshal(data, &valueTransferCallTool); err == nil {
		t.TransferCallTool = valueTransferCallTool
		return nil
	}
	valueOutputTool := new(OutputTool)
	if err := json.Unmarshal(data, &valueOutputTool); err == nil {
		t.OutputTool = valueOutputTool
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, t)
}

func (t ToolsCreateResponse) MarshalJSON() ([]byte, error) {
	if t.DtmfTool != nil {
		return json.Marshal(t.DtmfTool)
	}
	if t.EndCallTool != nil {
		return json.Marshal(t.EndCallTool)
	}
	if t.FunctionTool != nil {
		return json.Marshal(t.FunctionTool)
	}
	if t.GhlTool != nil {
		return json.Marshal(t.GhlTool)
	}
	if t.MakeTool != nil {
		return json.Marshal(t.MakeTool)
	}
	if t.TransferCallTool != nil {
		return json.Marshal(t.TransferCallTool)
	}
	if t.OutputTool != nil {
		return json.Marshal(t.OutputTool)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", t)
}

type ToolsCreateResponseVisitor interface {
	VisitDtmfTool(*DtmfTool) error
	VisitEndCallTool(*EndCallTool) error
	VisitFunctionTool(*FunctionTool) error
	VisitGhlTool(*GhlTool) error
	VisitMakeTool(*MakeTool) error
	VisitTransferCallTool(*TransferCallTool) error
	VisitOutputTool(*OutputTool) error
}

func (t *ToolsCreateResponse) Accept(visitor ToolsCreateResponseVisitor) error {
	if t.DtmfTool != nil {
		return visitor.VisitDtmfTool(t.DtmfTool)
	}
	if t.EndCallTool != nil {
		return visitor.VisitEndCallTool(t.EndCallTool)
	}
	if t.FunctionTool != nil {
		return visitor.VisitFunctionTool(t.FunctionTool)
	}
	if t.GhlTool != nil {
		return visitor.VisitGhlTool(t.GhlTool)
	}
	if t.MakeTool != nil {
		return visitor.VisitMakeTool(t.MakeTool)
	}
	if t.TransferCallTool != nil {
		return visitor.VisitTransferCallTool(t.TransferCallTool)
	}
	if t.OutputTool != nil {
		return visitor.VisitOutputTool(t.OutputTool)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", t)
}

type ToolsDeleteResponse struct {
	DtmfTool         *DtmfTool
	EndCallTool      *EndCallTool
	FunctionTool     *FunctionTool
	GhlTool          *GhlTool
	MakeTool         *MakeTool
	TransferCallTool *TransferCallTool
	OutputTool       *OutputTool
}

func (t *ToolsDeleteResponse) UnmarshalJSON(data []byte) error {
	valueDtmfTool := new(DtmfTool)
	if err := json.Unmarshal(data, &valueDtmfTool); err == nil {
		t.DtmfTool = valueDtmfTool
		return nil
	}
	valueEndCallTool := new(EndCallTool)
	if err := json.Unmarshal(data, &valueEndCallTool); err == nil {
		t.EndCallTool = valueEndCallTool
		return nil
	}
	valueFunctionTool := new(FunctionTool)
	if err := json.Unmarshal(data, &valueFunctionTool); err == nil {
		t.FunctionTool = valueFunctionTool
		return nil
	}
	valueGhlTool := new(GhlTool)
	if err := json.Unmarshal(data, &valueGhlTool); err == nil {
		t.GhlTool = valueGhlTool
		return nil
	}
	valueMakeTool := new(MakeTool)
	if err := json.Unmarshal(data, &valueMakeTool); err == nil {
		t.MakeTool = valueMakeTool
		return nil
	}
	valueTransferCallTool := new(TransferCallTool)
	if err := json.Unmarshal(data, &valueTransferCallTool); err == nil {
		t.TransferCallTool = valueTransferCallTool
		return nil
	}
	valueOutputTool := new(OutputTool)
	if err := json.Unmarshal(data, &valueOutputTool); err == nil {
		t.OutputTool = valueOutputTool
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, t)
}

func (t ToolsDeleteResponse) MarshalJSON() ([]byte, error) {
	if t.DtmfTool != nil {
		return json.Marshal(t.DtmfTool)
	}
	if t.EndCallTool != nil {
		return json.Marshal(t.EndCallTool)
	}
	if t.FunctionTool != nil {
		return json.Marshal(t.FunctionTool)
	}
	if t.GhlTool != nil {
		return json.Marshal(t.GhlTool)
	}
	if t.MakeTool != nil {
		return json.Marshal(t.MakeTool)
	}
	if t.TransferCallTool != nil {
		return json.Marshal(t.TransferCallTool)
	}
	if t.OutputTool != nil {
		return json.Marshal(t.OutputTool)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", t)
}

type ToolsDeleteResponseVisitor interface {
	VisitDtmfTool(*DtmfTool) error
	VisitEndCallTool(*EndCallTool) error
	VisitFunctionTool(*FunctionTool) error
	VisitGhlTool(*GhlTool) error
	VisitMakeTool(*MakeTool) error
	VisitTransferCallTool(*TransferCallTool) error
	VisitOutputTool(*OutputTool) error
}

func (t *ToolsDeleteResponse) Accept(visitor ToolsDeleteResponseVisitor) error {
	if t.DtmfTool != nil {
		return visitor.VisitDtmfTool(t.DtmfTool)
	}
	if t.EndCallTool != nil {
		return visitor.VisitEndCallTool(t.EndCallTool)
	}
	if t.FunctionTool != nil {
		return visitor.VisitFunctionTool(t.FunctionTool)
	}
	if t.GhlTool != nil {
		return visitor.VisitGhlTool(t.GhlTool)
	}
	if t.MakeTool != nil {
		return visitor.VisitMakeTool(t.MakeTool)
	}
	if t.TransferCallTool != nil {
		return visitor.VisitTransferCallTool(t.TransferCallTool)
	}
	if t.OutputTool != nil {
		return visitor.VisitOutputTool(t.OutputTool)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", t)
}

type ToolsGetResponse struct {
	DtmfTool         *DtmfTool
	EndCallTool      *EndCallTool
	FunctionTool     *FunctionTool
	GhlTool          *GhlTool
	MakeTool         *MakeTool
	TransferCallTool *TransferCallTool
	OutputTool       *OutputTool
}

func (t *ToolsGetResponse) UnmarshalJSON(data []byte) error {
	valueDtmfTool := new(DtmfTool)
	if err := json.Unmarshal(data, &valueDtmfTool); err == nil {
		t.DtmfTool = valueDtmfTool
		return nil
	}
	valueEndCallTool := new(EndCallTool)
	if err := json.Unmarshal(data, &valueEndCallTool); err == nil {
		t.EndCallTool = valueEndCallTool
		return nil
	}
	valueFunctionTool := new(FunctionTool)
	if err := json.Unmarshal(data, &valueFunctionTool); err == nil {
		t.FunctionTool = valueFunctionTool
		return nil
	}
	valueGhlTool := new(GhlTool)
	if err := json.Unmarshal(data, &valueGhlTool); err == nil {
		t.GhlTool = valueGhlTool
		return nil
	}
	valueMakeTool := new(MakeTool)
	if err := json.Unmarshal(data, &valueMakeTool); err == nil {
		t.MakeTool = valueMakeTool
		return nil
	}
	valueTransferCallTool := new(TransferCallTool)
	if err := json.Unmarshal(data, &valueTransferCallTool); err == nil {
		t.TransferCallTool = valueTransferCallTool
		return nil
	}
	valueOutputTool := new(OutputTool)
	if err := json.Unmarshal(data, &valueOutputTool); err == nil {
		t.OutputTool = valueOutputTool
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, t)
}

func (t ToolsGetResponse) MarshalJSON() ([]byte, error) {
	if t.DtmfTool != nil {
		return json.Marshal(t.DtmfTool)
	}
	if t.EndCallTool != nil {
		return json.Marshal(t.EndCallTool)
	}
	if t.FunctionTool != nil {
		return json.Marshal(t.FunctionTool)
	}
	if t.GhlTool != nil {
		return json.Marshal(t.GhlTool)
	}
	if t.MakeTool != nil {
		return json.Marshal(t.MakeTool)
	}
	if t.TransferCallTool != nil {
		return json.Marshal(t.TransferCallTool)
	}
	if t.OutputTool != nil {
		return json.Marshal(t.OutputTool)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", t)
}

type ToolsGetResponseVisitor interface {
	VisitDtmfTool(*DtmfTool) error
	VisitEndCallTool(*EndCallTool) error
	VisitFunctionTool(*FunctionTool) error
	VisitGhlTool(*GhlTool) error
	VisitMakeTool(*MakeTool) error
	VisitTransferCallTool(*TransferCallTool) error
	VisitOutputTool(*OutputTool) error
}

func (t *ToolsGetResponse) Accept(visitor ToolsGetResponseVisitor) error {
	if t.DtmfTool != nil {
		return visitor.VisitDtmfTool(t.DtmfTool)
	}
	if t.EndCallTool != nil {
		return visitor.VisitEndCallTool(t.EndCallTool)
	}
	if t.FunctionTool != nil {
		return visitor.VisitFunctionTool(t.FunctionTool)
	}
	if t.GhlTool != nil {
		return visitor.VisitGhlTool(t.GhlTool)
	}
	if t.MakeTool != nil {
		return visitor.VisitMakeTool(t.MakeTool)
	}
	if t.TransferCallTool != nil {
		return visitor.VisitTransferCallTool(t.TransferCallTool)
	}
	if t.OutputTool != nil {
		return visitor.VisitOutputTool(t.OutputTool)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", t)
}

type ToolsListResponseItem struct {
	DtmfTool         *DtmfTool
	EndCallTool      *EndCallTool
	FunctionTool     *FunctionTool
	GhlTool          *GhlTool
	MakeTool         *MakeTool
	TransferCallTool *TransferCallTool
	OutputTool       *OutputTool
}

func (t *ToolsListResponseItem) UnmarshalJSON(data []byte) error {
	valueDtmfTool := new(DtmfTool)
	if err := json.Unmarshal(data, &valueDtmfTool); err == nil {
		t.DtmfTool = valueDtmfTool
		return nil
	}
	valueEndCallTool := new(EndCallTool)
	if err := json.Unmarshal(data, &valueEndCallTool); err == nil {
		t.EndCallTool = valueEndCallTool
		return nil
	}
	valueFunctionTool := new(FunctionTool)
	if err := json.Unmarshal(data, &valueFunctionTool); err == nil {
		t.FunctionTool = valueFunctionTool
		return nil
	}
	valueGhlTool := new(GhlTool)
	if err := json.Unmarshal(data, &valueGhlTool); err == nil {
		t.GhlTool = valueGhlTool
		return nil
	}
	valueMakeTool := new(MakeTool)
	if err := json.Unmarshal(data, &valueMakeTool); err == nil {
		t.MakeTool = valueMakeTool
		return nil
	}
	valueTransferCallTool := new(TransferCallTool)
	if err := json.Unmarshal(data, &valueTransferCallTool); err == nil {
		t.TransferCallTool = valueTransferCallTool
		return nil
	}
	valueOutputTool := new(OutputTool)
	if err := json.Unmarshal(data, &valueOutputTool); err == nil {
		t.OutputTool = valueOutputTool
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, t)
}

func (t ToolsListResponseItem) MarshalJSON() ([]byte, error) {
	if t.DtmfTool != nil {
		return json.Marshal(t.DtmfTool)
	}
	if t.EndCallTool != nil {
		return json.Marshal(t.EndCallTool)
	}
	if t.FunctionTool != nil {
		return json.Marshal(t.FunctionTool)
	}
	if t.GhlTool != nil {
		return json.Marshal(t.GhlTool)
	}
	if t.MakeTool != nil {
		return json.Marshal(t.MakeTool)
	}
	if t.TransferCallTool != nil {
		return json.Marshal(t.TransferCallTool)
	}
	if t.OutputTool != nil {
		return json.Marshal(t.OutputTool)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", t)
}

type ToolsListResponseItemVisitor interface {
	VisitDtmfTool(*DtmfTool) error
	VisitEndCallTool(*EndCallTool) error
	VisitFunctionTool(*FunctionTool) error
	VisitGhlTool(*GhlTool) error
	VisitMakeTool(*MakeTool) error
	VisitTransferCallTool(*TransferCallTool) error
	VisitOutputTool(*OutputTool) error
}

func (t *ToolsListResponseItem) Accept(visitor ToolsListResponseItemVisitor) error {
	if t.DtmfTool != nil {
		return visitor.VisitDtmfTool(t.DtmfTool)
	}
	if t.EndCallTool != nil {
		return visitor.VisitEndCallTool(t.EndCallTool)
	}
	if t.FunctionTool != nil {
		return visitor.VisitFunctionTool(t.FunctionTool)
	}
	if t.GhlTool != nil {
		return visitor.VisitGhlTool(t.GhlTool)
	}
	if t.MakeTool != nil {
		return visitor.VisitMakeTool(t.MakeTool)
	}
	if t.TransferCallTool != nil {
		return visitor.VisitTransferCallTool(t.TransferCallTool)
	}
	if t.OutputTool != nil {
		return visitor.VisitOutputTool(t.OutputTool)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", t)
}

type ToolsUpdateResponse struct {
	DtmfTool         *DtmfTool
	EndCallTool      *EndCallTool
	FunctionTool     *FunctionTool
	GhlTool          *GhlTool
	MakeTool         *MakeTool
	TransferCallTool *TransferCallTool
	OutputTool       *OutputTool
}

func (t *ToolsUpdateResponse) UnmarshalJSON(data []byte) error {
	valueDtmfTool := new(DtmfTool)
	if err := json.Unmarshal(data, &valueDtmfTool); err == nil {
		t.DtmfTool = valueDtmfTool
		return nil
	}
	valueEndCallTool := new(EndCallTool)
	if err := json.Unmarshal(data, &valueEndCallTool); err == nil {
		t.EndCallTool = valueEndCallTool
		return nil
	}
	valueFunctionTool := new(FunctionTool)
	if err := json.Unmarshal(data, &valueFunctionTool); err == nil {
		t.FunctionTool = valueFunctionTool
		return nil
	}
	valueGhlTool := new(GhlTool)
	if err := json.Unmarshal(data, &valueGhlTool); err == nil {
		t.GhlTool = valueGhlTool
		return nil
	}
	valueMakeTool := new(MakeTool)
	if err := json.Unmarshal(data, &valueMakeTool); err == nil {
		t.MakeTool = valueMakeTool
		return nil
	}
	valueTransferCallTool := new(TransferCallTool)
	if err := json.Unmarshal(data, &valueTransferCallTool); err == nil {
		t.TransferCallTool = valueTransferCallTool
		return nil
	}
	valueOutputTool := new(OutputTool)
	if err := json.Unmarshal(data, &valueOutputTool); err == nil {
		t.OutputTool = valueOutputTool
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, t)
}

func (t ToolsUpdateResponse) MarshalJSON() ([]byte, error) {
	if t.DtmfTool != nil {
		return json.Marshal(t.DtmfTool)
	}
	if t.EndCallTool != nil {
		return json.Marshal(t.EndCallTool)
	}
	if t.FunctionTool != nil {
		return json.Marshal(t.FunctionTool)
	}
	if t.GhlTool != nil {
		return json.Marshal(t.GhlTool)
	}
	if t.MakeTool != nil {
		return json.Marshal(t.MakeTool)
	}
	if t.TransferCallTool != nil {
		return json.Marshal(t.TransferCallTool)
	}
	if t.OutputTool != nil {
		return json.Marshal(t.OutputTool)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", t)
}

type ToolsUpdateResponseVisitor interface {
	VisitDtmfTool(*DtmfTool) error
	VisitEndCallTool(*EndCallTool) error
	VisitFunctionTool(*FunctionTool) error
	VisitGhlTool(*GhlTool) error
	VisitMakeTool(*MakeTool) error
	VisitTransferCallTool(*TransferCallTool) error
	VisitOutputTool(*OutputTool) error
}

func (t *ToolsUpdateResponse) Accept(visitor ToolsUpdateResponseVisitor) error {
	if t.DtmfTool != nil {
		return visitor.VisitDtmfTool(t.DtmfTool)
	}
	if t.EndCallTool != nil {
		return visitor.VisitEndCallTool(t.EndCallTool)
	}
	if t.FunctionTool != nil {
		return visitor.VisitFunctionTool(t.FunctionTool)
	}
	if t.GhlTool != nil {
		return visitor.VisitGhlTool(t.GhlTool)
	}
	if t.MakeTool != nil {
		return visitor.VisitMakeTool(t.MakeTool)
	}
	if t.TransferCallTool != nil {
		return visitor.VisitTransferCallTool(t.TransferCallTool)
	}
	if t.OutputTool != nil {
		return visitor.VisitOutputTool(t.OutputTool)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", t)
}

type UpdateToolDtoMessagesItem struct {
	ToolMessageStart    *ToolMessageStart
	ToolMessageComplete *ToolMessageComplete
	ToolMessageFailed   *ToolMessageFailed
	ToolMessageDelayed  *ToolMessageDelayed
}

func (u *UpdateToolDtoMessagesItem) UnmarshalJSON(data []byte) error {
	valueToolMessageStart := new(ToolMessageStart)
	if err := json.Unmarshal(data, &valueToolMessageStart); err == nil {
		u.ToolMessageStart = valueToolMessageStart
		return nil
	}
	valueToolMessageComplete := new(ToolMessageComplete)
	if err := json.Unmarshal(data, &valueToolMessageComplete); err == nil {
		u.ToolMessageComplete = valueToolMessageComplete
		return nil
	}
	valueToolMessageFailed := new(ToolMessageFailed)
	if err := json.Unmarshal(data, &valueToolMessageFailed); err == nil {
		u.ToolMessageFailed = valueToolMessageFailed
		return nil
	}
	valueToolMessageDelayed := new(ToolMessageDelayed)
	if err := json.Unmarshal(data, &valueToolMessageDelayed); err == nil {
		u.ToolMessageDelayed = valueToolMessageDelayed
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, u)
}

func (u UpdateToolDtoMessagesItem) MarshalJSON() ([]byte, error) {
	if u.ToolMessageStart != nil {
		return json.Marshal(u.ToolMessageStart)
	}
	if u.ToolMessageComplete != nil {
		return json.Marshal(u.ToolMessageComplete)
	}
	if u.ToolMessageFailed != nil {
		return json.Marshal(u.ToolMessageFailed)
	}
	if u.ToolMessageDelayed != nil {
		return json.Marshal(u.ToolMessageDelayed)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", u)
}

type UpdateToolDtoMessagesItemVisitor interface {
	VisitToolMessageStart(*ToolMessageStart) error
	VisitToolMessageComplete(*ToolMessageComplete) error
	VisitToolMessageFailed(*ToolMessageFailed) error
	VisitToolMessageDelayed(*ToolMessageDelayed) error
}

func (u *UpdateToolDtoMessagesItem) Accept(visitor UpdateToolDtoMessagesItemVisitor) error {
	if u.ToolMessageStart != nil {
		return visitor.VisitToolMessageStart(u.ToolMessageStart)
	}
	if u.ToolMessageComplete != nil {
		return visitor.VisitToolMessageComplete(u.ToolMessageComplete)
	}
	if u.ToolMessageFailed != nil {
		return visitor.VisitToolMessageFailed(u.ToolMessageFailed)
	}
	if u.ToolMessageDelayed != nil {
		return visitor.VisitToolMessageDelayed(u.ToolMessageDelayed)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", u)
}

type UpdateToolDto struct {
	// This determines if the tool is async.
	//
	// If async, the assistant will move forward without waiting for your server to respond. This is useful if you just want to trigger something on your server.
	//
	// If sync, the assistant will wait for your server to respond. This is useful if want assistant to respond with the result from your server.
	//
	// Defaults to synchronous (`false`).
	Async *bool `json:"async,omitempty" url:"-"`
	// These are the messages that will be spoken to the user as the tool is running.
	//
	// For some tools, this is auto-filled based on special fields like `tool.destinations`. For others like the function tool, these can be custom configured.
	Messages []*UpdateToolDtoMessagesItem `json:"messages,omitempty" url:"-"`
	// This is the function definition of the tool.
	//
	// For `endCall`, `transferCall`, and `dtmf` tools, this is auto-filled based on tool-specific fields like `tool.destinations`. But, even in those cases, you can provide a custom function definition for advanced use cases.
	//
	// An example of an advanced use case is if you want to customize the message that's spoken for `endCall` tool. You can specify a function where it returns an argument "reason". Then, in `messages` array, you can have many "request-complete" messages. One of these messages will be triggered if the `messages[].conditions` matches the "reason" argument.
	Function *OpenAiFunction `json:"function,omitempty" url:"-"`
	// This is the server that will be hit when this tool is requested by the model.
	//
	// All requests will be sent with the call object among other things. You can find more details in the Server URL documentation.
	//
	// This overrides the serverUrl set on the org and the phoneNumber. Order of precedence: highest tool.server.url, then assistant.serverUrl, then phoneNumber.serverUrl, then org.serverUrl.
	Server *Server `json:"server,omitempty" url:"-"`
}
