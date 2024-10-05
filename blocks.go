// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	time "time"
)

type BlocksListRequest struct {
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

type BlocksCreateRequest struct {
	CreateConversationBlockDto *CreateConversationBlockDto
	CreateToolCallBlockDto     *CreateToolCallBlockDto
	CreateWorkflowBlockDto     *CreateWorkflowBlockDto
}

func (b *BlocksCreateRequest) UnmarshalJSON(data []byte) error {
	valueCreateConversationBlockDto := new(CreateConversationBlockDto)
	if err := json.Unmarshal(data, &valueCreateConversationBlockDto); err == nil {
		b.CreateConversationBlockDto = valueCreateConversationBlockDto
		return nil
	}
	valueCreateToolCallBlockDto := new(CreateToolCallBlockDto)
	if err := json.Unmarshal(data, &valueCreateToolCallBlockDto); err == nil {
		b.CreateToolCallBlockDto = valueCreateToolCallBlockDto
		return nil
	}
	valueCreateWorkflowBlockDto := new(CreateWorkflowBlockDto)
	if err := json.Unmarshal(data, &valueCreateWorkflowBlockDto); err == nil {
		b.CreateWorkflowBlockDto = valueCreateWorkflowBlockDto
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, b)
}

func (b BlocksCreateRequest) MarshalJSON() ([]byte, error) {
	if b.CreateConversationBlockDto != nil {
		return json.Marshal(b.CreateConversationBlockDto)
	}
	if b.CreateToolCallBlockDto != nil {
		return json.Marshal(b.CreateToolCallBlockDto)
	}
	if b.CreateWorkflowBlockDto != nil {
		return json.Marshal(b.CreateWorkflowBlockDto)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", b)
}

type BlocksCreateRequestVisitor interface {
	VisitCreateConversationBlockDto(*CreateConversationBlockDto) error
	VisitCreateToolCallBlockDto(*CreateToolCallBlockDto) error
	VisitCreateWorkflowBlockDto(*CreateWorkflowBlockDto) error
}

func (b *BlocksCreateRequest) Accept(visitor BlocksCreateRequestVisitor) error {
	if b.CreateConversationBlockDto != nil {
		return visitor.VisitCreateConversationBlockDto(b.CreateConversationBlockDto)
	}
	if b.CreateToolCallBlockDto != nil {
		return visitor.VisitCreateToolCallBlockDto(b.CreateToolCallBlockDto)
	}
	if b.CreateWorkflowBlockDto != nil {
		return visitor.VisitCreateWorkflowBlockDto(b.CreateWorkflowBlockDto)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", b)
}

type BlocksCreateResponse struct {
	ConversationBlock *ConversationBlock
	ToolCallBlock     *ToolCallBlock
	WorkflowBlock     *WorkflowBlock
}

func (b *BlocksCreateResponse) UnmarshalJSON(data []byte) error {
	valueConversationBlock := new(ConversationBlock)
	if err := json.Unmarshal(data, &valueConversationBlock); err == nil {
		b.ConversationBlock = valueConversationBlock
		return nil
	}
	valueToolCallBlock := new(ToolCallBlock)
	if err := json.Unmarshal(data, &valueToolCallBlock); err == nil {
		b.ToolCallBlock = valueToolCallBlock
		return nil
	}
	valueWorkflowBlock := new(WorkflowBlock)
	if err := json.Unmarshal(data, &valueWorkflowBlock); err == nil {
		b.WorkflowBlock = valueWorkflowBlock
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, b)
}

func (b BlocksCreateResponse) MarshalJSON() ([]byte, error) {
	if b.ConversationBlock != nil {
		return json.Marshal(b.ConversationBlock)
	}
	if b.ToolCallBlock != nil {
		return json.Marshal(b.ToolCallBlock)
	}
	if b.WorkflowBlock != nil {
		return json.Marshal(b.WorkflowBlock)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", b)
}

type BlocksCreateResponseVisitor interface {
	VisitConversationBlock(*ConversationBlock) error
	VisitToolCallBlock(*ToolCallBlock) error
	VisitWorkflowBlock(*WorkflowBlock) error
}

func (b *BlocksCreateResponse) Accept(visitor BlocksCreateResponseVisitor) error {
	if b.ConversationBlock != nil {
		return visitor.VisitConversationBlock(b.ConversationBlock)
	}
	if b.ToolCallBlock != nil {
		return visitor.VisitToolCallBlock(b.ToolCallBlock)
	}
	if b.WorkflowBlock != nil {
		return visitor.VisitWorkflowBlock(b.WorkflowBlock)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", b)
}

type BlocksDeleteResponse struct {
	ConversationBlock *ConversationBlock
	ToolCallBlock     *ToolCallBlock
	WorkflowBlock     *WorkflowBlock
}

func (b *BlocksDeleteResponse) UnmarshalJSON(data []byte) error {
	valueConversationBlock := new(ConversationBlock)
	if err := json.Unmarshal(data, &valueConversationBlock); err == nil {
		b.ConversationBlock = valueConversationBlock
		return nil
	}
	valueToolCallBlock := new(ToolCallBlock)
	if err := json.Unmarshal(data, &valueToolCallBlock); err == nil {
		b.ToolCallBlock = valueToolCallBlock
		return nil
	}
	valueWorkflowBlock := new(WorkflowBlock)
	if err := json.Unmarshal(data, &valueWorkflowBlock); err == nil {
		b.WorkflowBlock = valueWorkflowBlock
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, b)
}

func (b BlocksDeleteResponse) MarshalJSON() ([]byte, error) {
	if b.ConversationBlock != nil {
		return json.Marshal(b.ConversationBlock)
	}
	if b.ToolCallBlock != nil {
		return json.Marshal(b.ToolCallBlock)
	}
	if b.WorkflowBlock != nil {
		return json.Marshal(b.WorkflowBlock)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", b)
}

type BlocksDeleteResponseVisitor interface {
	VisitConversationBlock(*ConversationBlock) error
	VisitToolCallBlock(*ToolCallBlock) error
	VisitWorkflowBlock(*WorkflowBlock) error
}

func (b *BlocksDeleteResponse) Accept(visitor BlocksDeleteResponseVisitor) error {
	if b.ConversationBlock != nil {
		return visitor.VisitConversationBlock(b.ConversationBlock)
	}
	if b.ToolCallBlock != nil {
		return visitor.VisitToolCallBlock(b.ToolCallBlock)
	}
	if b.WorkflowBlock != nil {
		return visitor.VisitWorkflowBlock(b.WorkflowBlock)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", b)
}

type BlocksGetResponse struct {
	ConversationBlock *ConversationBlock
	ToolCallBlock     *ToolCallBlock
	WorkflowBlock     *WorkflowBlock
}

func (b *BlocksGetResponse) UnmarshalJSON(data []byte) error {
	valueConversationBlock := new(ConversationBlock)
	if err := json.Unmarshal(data, &valueConversationBlock); err == nil {
		b.ConversationBlock = valueConversationBlock
		return nil
	}
	valueToolCallBlock := new(ToolCallBlock)
	if err := json.Unmarshal(data, &valueToolCallBlock); err == nil {
		b.ToolCallBlock = valueToolCallBlock
		return nil
	}
	valueWorkflowBlock := new(WorkflowBlock)
	if err := json.Unmarshal(data, &valueWorkflowBlock); err == nil {
		b.WorkflowBlock = valueWorkflowBlock
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, b)
}

func (b BlocksGetResponse) MarshalJSON() ([]byte, error) {
	if b.ConversationBlock != nil {
		return json.Marshal(b.ConversationBlock)
	}
	if b.ToolCallBlock != nil {
		return json.Marshal(b.ToolCallBlock)
	}
	if b.WorkflowBlock != nil {
		return json.Marshal(b.WorkflowBlock)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", b)
}

type BlocksGetResponseVisitor interface {
	VisitConversationBlock(*ConversationBlock) error
	VisitToolCallBlock(*ToolCallBlock) error
	VisitWorkflowBlock(*WorkflowBlock) error
}

func (b *BlocksGetResponse) Accept(visitor BlocksGetResponseVisitor) error {
	if b.ConversationBlock != nil {
		return visitor.VisitConversationBlock(b.ConversationBlock)
	}
	if b.ToolCallBlock != nil {
		return visitor.VisitToolCallBlock(b.ToolCallBlock)
	}
	if b.WorkflowBlock != nil {
		return visitor.VisitWorkflowBlock(b.WorkflowBlock)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", b)
}

type BlocksListResponseItem struct {
	ConversationBlock *ConversationBlock
	ToolCallBlock     *ToolCallBlock
	WorkflowBlock     *WorkflowBlock
}

func (b *BlocksListResponseItem) UnmarshalJSON(data []byte) error {
	valueConversationBlock := new(ConversationBlock)
	if err := json.Unmarshal(data, &valueConversationBlock); err == nil {
		b.ConversationBlock = valueConversationBlock
		return nil
	}
	valueToolCallBlock := new(ToolCallBlock)
	if err := json.Unmarshal(data, &valueToolCallBlock); err == nil {
		b.ToolCallBlock = valueToolCallBlock
		return nil
	}
	valueWorkflowBlock := new(WorkflowBlock)
	if err := json.Unmarshal(data, &valueWorkflowBlock); err == nil {
		b.WorkflowBlock = valueWorkflowBlock
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, b)
}

func (b BlocksListResponseItem) MarshalJSON() ([]byte, error) {
	if b.ConversationBlock != nil {
		return json.Marshal(b.ConversationBlock)
	}
	if b.ToolCallBlock != nil {
		return json.Marshal(b.ToolCallBlock)
	}
	if b.WorkflowBlock != nil {
		return json.Marshal(b.WorkflowBlock)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", b)
}

type BlocksListResponseItemVisitor interface {
	VisitConversationBlock(*ConversationBlock) error
	VisitToolCallBlock(*ToolCallBlock) error
	VisitWorkflowBlock(*WorkflowBlock) error
}

func (b *BlocksListResponseItem) Accept(visitor BlocksListResponseItemVisitor) error {
	if b.ConversationBlock != nil {
		return visitor.VisitConversationBlock(b.ConversationBlock)
	}
	if b.ToolCallBlock != nil {
		return visitor.VisitToolCallBlock(b.ToolCallBlock)
	}
	if b.WorkflowBlock != nil {
		return visitor.VisitWorkflowBlock(b.WorkflowBlock)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", b)
}

type BlocksUpdateResponse struct {
	ConversationBlock *ConversationBlock
	ToolCallBlock     *ToolCallBlock
	WorkflowBlock     *WorkflowBlock
}

func (b *BlocksUpdateResponse) UnmarshalJSON(data []byte) error {
	valueConversationBlock := new(ConversationBlock)
	if err := json.Unmarshal(data, &valueConversationBlock); err == nil {
		b.ConversationBlock = valueConversationBlock
		return nil
	}
	valueToolCallBlock := new(ToolCallBlock)
	if err := json.Unmarshal(data, &valueToolCallBlock); err == nil {
		b.ToolCallBlock = valueToolCallBlock
		return nil
	}
	valueWorkflowBlock := new(WorkflowBlock)
	if err := json.Unmarshal(data, &valueWorkflowBlock); err == nil {
		b.WorkflowBlock = valueWorkflowBlock
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, b)
}

func (b BlocksUpdateResponse) MarshalJSON() ([]byte, error) {
	if b.ConversationBlock != nil {
		return json.Marshal(b.ConversationBlock)
	}
	if b.ToolCallBlock != nil {
		return json.Marshal(b.ToolCallBlock)
	}
	if b.WorkflowBlock != nil {
		return json.Marshal(b.WorkflowBlock)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", b)
}

type BlocksUpdateResponseVisitor interface {
	VisitConversationBlock(*ConversationBlock) error
	VisitToolCallBlock(*ToolCallBlock) error
	VisitWorkflowBlock(*WorkflowBlock) error
}

func (b *BlocksUpdateResponse) Accept(visitor BlocksUpdateResponseVisitor) error {
	if b.ConversationBlock != nil {
		return visitor.VisitConversationBlock(b.ConversationBlock)
	}
	if b.ToolCallBlock != nil {
		return visitor.VisitToolCallBlock(b.ToolCallBlock)
	}
	if b.WorkflowBlock != nil {
		return visitor.VisitWorkflowBlock(b.WorkflowBlock)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", b)
}

type UpdateBlockDtoMessagesItem struct {
	BlockStartMessage    *BlockStartMessage
	BlockCompleteMessage *BlockCompleteMessage
}

func (u *UpdateBlockDtoMessagesItem) UnmarshalJSON(data []byte) error {
	valueBlockStartMessage := new(BlockStartMessage)
	if err := json.Unmarshal(data, &valueBlockStartMessage); err == nil {
		u.BlockStartMessage = valueBlockStartMessage
		return nil
	}
	valueBlockCompleteMessage := new(BlockCompleteMessage)
	if err := json.Unmarshal(data, &valueBlockCompleteMessage); err == nil {
		u.BlockCompleteMessage = valueBlockCompleteMessage
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, u)
}

func (u UpdateBlockDtoMessagesItem) MarshalJSON() ([]byte, error) {
	if u.BlockStartMessage != nil {
		return json.Marshal(u.BlockStartMessage)
	}
	if u.BlockCompleteMessage != nil {
		return json.Marshal(u.BlockCompleteMessage)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", u)
}

type UpdateBlockDtoMessagesItemVisitor interface {
	VisitBlockStartMessage(*BlockStartMessage) error
	VisitBlockCompleteMessage(*BlockCompleteMessage) error
}

func (u *UpdateBlockDtoMessagesItem) Accept(visitor UpdateBlockDtoMessagesItemVisitor) error {
	if u.BlockStartMessage != nil {
		return visitor.VisitBlockStartMessage(u.BlockStartMessage)
	}
	if u.BlockCompleteMessage != nil {
		return visitor.VisitBlockCompleteMessage(u.BlockCompleteMessage)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", u)
}

type UpdateBlockDtoStepsItem struct {
	HandoffStep  *HandoffStep
	CallbackStep *CallbackStep
}

func (u *UpdateBlockDtoStepsItem) UnmarshalJSON(data []byte) error {
	valueHandoffStep := new(HandoffStep)
	if err := json.Unmarshal(data, &valueHandoffStep); err == nil {
		u.HandoffStep = valueHandoffStep
		return nil
	}
	valueCallbackStep := new(CallbackStep)
	if err := json.Unmarshal(data, &valueCallbackStep); err == nil {
		u.CallbackStep = valueCallbackStep
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, u)
}

func (u UpdateBlockDtoStepsItem) MarshalJSON() ([]byte, error) {
	if u.HandoffStep != nil {
		return json.Marshal(u.HandoffStep)
	}
	if u.CallbackStep != nil {
		return json.Marshal(u.CallbackStep)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", u)
}

type UpdateBlockDtoStepsItemVisitor interface {
	VisitHandoffStep(*HandoffStep) error
	VisitCallbackStep(*CallbackStep) error
}

func (u *UpdateBlockDtoStepsItem) Accept(visitor UpdateBlockDtoStepsItemVisitor) error {
	if u.HandoffStep != nil {
		return visitor.VisitHandoffStep(u.HandoffStep)
	}
	if u.CallbackStep != nil {
		return visitor.VisitCallbackStep(u.CallbackStep)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", u)
}

// This is the tool that the block will call. To use an existing tool, use `toolId`.
type UpdateBlockDtoTool struct {
	CreateDtmfToolDto         *CreateDtmfToolDto
	CreateEndCallToolDto      *CreateEndCallToolDto
	CreateVoicemailToolDto    *CreateVoicemailToolDto
	CreateFunctionToolDto     *CreateFunctionToolDto
	CreateGhlToolDto          *CreateGhlToolDto
	CreateMakeToolDto         *CreateMakeToolDto
	CreateTransferCallToolDto *CreateTransferCallToolDto
}

func (u *UpdateBlockDtoTool) UnmarshalJSON(data []byte) error {
	valueCreateDtmfToolDto := new(CreateDtmfToolDto)
	if err := json.Unmarshal(data, &valueCreateDtmfToolDto); err == nil {
		u.CreateDtmfToolDto = valueCreateDtmfToolDto
		return nil
	}
	valueCreateEndCallToolDto := new(CreateEndCallToolDto)
	if err := json.Unmarshal(data, &valueCreateEndCallToolDto); err == nil {
		u.CreateEndCallToolDto = valueCreateEndCallToolDto
		return nil
	}
	valueCreateVoicemailToolDto := new(CreateVoicemailToolDto)
	if err := json.Unmarshal(data, &valueCreateVoicemailToolDto); err == nil {
		u.CreateVoicemailToolDto = valueCreateVoicemailToolDto
		return nil
	}
	valueCreateFunctionToolDto := new(CreateFunctionToolDto)
	if err := json.Unmarshal(data, &valueCreateFunctionToolDto); err == nil {
		u.CreateFunctionToolDto = valueCreateFunctionToolDto
		return nil
	}
	valueCreateGhlToolDto := new(CreateGhlToolDto)
	if err := json.Unmarshal(data, &valueCreateGhlToolDto); err == nil {
		u.CreateGhlToolDto = valueCreateGhlToolDto
		return nil
	}
	valueCreateMakeToolDto := new(CreateMakeToolDto)
	if err := json.Unmarshal(data, &valueCreateMakeToolDto); err == nil {
		u.CreateMakeToolDto = valueCreateMakeToolDto
		return nil
	}
	valueCreateTransferCallToolDto := new(CreateTransferCallToolDto)
	if err := json.Unmarshal(data, &valueCreateTransferCallToolDto); err == nil {
		u.CreateTransferCallToolDto = valueCreateTransferCallToolDto
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, u)
}

func (u UpdateBlockDtoTool) MarshalJSON() ([]byte, error) {
	if u.CreateDtmfToolDto != nil {
		return json.Marshal(u.CreateDtmfToolDto)
	}
	if u.CreateEndCallToolDto != nil {
		return json.Marshal(u.CreateEndCallToolDto)
	}
	if u.CreateVoicemailToolDto != nil {
		return json.Marshal(u.CreateVoicemailToolDto)
	}
	if u.CreateFunctionToolDto != nil {
		return json.Marshal(u.CreateFunctionToolDto)
	}
	if u.CreateGhlToolDto != nil {
		return json.Marshal(u.CreateGhlToolDto)
	}
	if u.CreateMakeToolDto != nil {
		return json.Marshal(u.CreateMakeToolDto)
	}
	if u.CreateTransferCallToolDto != nil {
		return json.Marshal(u.CreateTransferCallToolDto)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", u)
}

type UpdateBlockDtoToolVisitor interface {
	VisitCreateDtmfToolDto(*CreateDtmfToolDto) error
	VisitCreateEndCallToolDto(*CreateEndCallToolDto) error
	VisitCreateVoicemailToolDto(*CreateVoicemailToolDto) error
	VisitCreateFunctionToolDto(*CreateFunctionToolDto) error
	VisitCreateGhlToolDto(*CreateGhlToolDto) error
	VisitCreateMakeToolDto(*CreateMakeToolDto) error
	VisitCreateTransferCallToolDto(*CreateTransferCallToolDto) error
}

func (u *UpdateBlockDtoTool) Accept(visitor UpdateBlockDtoToolVisitor) error {
	if u.CreateDtmfToolDto != nil {
		return visitor.VisitCreateDtmfToolDto(u.CreateDtmfToolDto)
	}
	if u.CreateEndCallToolDto != nil {
		return visitor.VisitCreateEndCallToolDto(u.CreateEndCallToolDto)
	}
	if u.CreateVoicemailToolDto != nil {
		return visitor.VisitCreateVoicemailToolDto(u.CreateVoicemailToolDto)
	}
	if u.CreateFunctionToolDto != nil {
		return visitor.VisitCreateFunctionToolDto(u.CreateFunctionToolDto)
	}
	if u.CreateGhlToolDto != nil {
		return visitor.VisitCreateGhlToolDto(u.CreateGhlToolDto)
	}
	if u.CreateMakeToolDto != nil {
		return visitor.VisitCreateMakeToolDto(u.CreateMakeToolDto)
	}
	if u.CreateTransferCallToolDto != nil {
		return visitor.VisitCreateTransferCallToolDto(u.CreateTransferCallToolDto)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", u)
}

type UpdateBlockDto struct {
	// These are the pre-configured messages that will be spoken to the user while the block is running.
	Messages []*UpdateBlockDtoMessagesItem `json:"messages,omitempty" url:"-"`
	// This is the input schema for the block. This is the input the block needs to run. It's given to the block as `steps[0].input`
	//
	// These are accessible as variables:
	// - ({{input.propertyName}}) in context of the block execution (step)
	// - ({{stepName.input.propertyName}}) in context of the workflow
	InputSchema *JsonSchema `json:"inputSchema,omitempty" url:"-"`
	// This is the output schema for the block. This is the output the block will return to the workflow (`{{stepName.output}}`).
	//
	// These are accessible as variables:
	// - ({{output.propertyName}}) in context of the block execution (step)
	// - ({{stepName.output.propertyName}}) in context of the workflow (read caveat #1)
	// - ({{blockName.output.propertyName}}) in context of the workflow (read caveat #2)
	//
	// Caveats:
	// 1. a workflow can execute a step multiple times. example, if a loop is used in the graph. {{stepName.output.propertyName}} will reference the latest usage of the step.
	// 2. a workflow can execute a block multiple times. example, if a step is called multiple times or if a block is used in multiple steps. {{blockName.output.propertyName}} will reference the latest usage of the block. this liquid variable is just provided for convenience when creating blocks outside of a workflow with steps.
	OutputSchema *JsonSchema `json:"outputSchema,omitempty" url:"-"`
	// This is the tool that the block will call. To use an existing tool, use `toolId`.
	Tool *UpdateBlockDtoTool `json:"tool,omitempty" url:"-"`
	// These are the steps in the workflow.
	Steps []*UpdateBlockDtoStepsItem `json:"steps,omitempty" url:"-"`
	// This is the name of the block. This is just for your reference.
	Name *string `json:"name,omitempty" url:"-"`
	// This is the instruction to the model.
	//
	// You can reference any variable in the context of the current block execution (step):
	// - "{{input.your-property-name}}" for the current step's input
	// - "{{your-step-name.output.your-property-name}}" for another step's output (in the same workflow; read caveat #1)
	// - "{{your-step-name.input.your-property-name}}" for another step's input (in the same workflow; read caveat #1)
	// - "{{your-block-name.output.your-property-name}}" for another block's output (in the same workflow; read caveat #2)
	// - "{{your-block-name.input.your-property-name}}" for another block's input (in the same workflow; read caveat #2)
	// - "{{workflow.input.your-property-name}}" for the current workflow's input
	// - "{{global.your-property-name}}" for the global context
	//
	// This can be as simple or as complex as you want it to be.
	// - "say hello and ask the user about their day!"
	// - "collect the user's first and last name"
	// - "user is {{input.firstName}} {{input.lastName}}. their age is {{input.age}}. ask them about their salary and if they might be interested in buying a house. we offer {{input.offer}}"
	//
	// Caveats:
	// 1. a workflow can execute a step multiple times. example, if a loop is used in the graph. {{stepName.output/input.propertyName}} will reference the latest usage of the step.
	// 2. a workflow can execute a block multiple times. example, if a step is called multiple times or if a block is used in multiple steps. {{blockName.output/input.propertyName}} will reference the latest usage of the block. this liquid variable is just provided for convenience when creating blocks outside of a workflow with steps.
	Instruction *string `json:"instruction,omitempty" url:"-"`
	// This is the id of the tool that the block will call. To use a transient tool, use `tool`.
	ToolId *string `json:"toolId,omitempty" url:"-"`
}