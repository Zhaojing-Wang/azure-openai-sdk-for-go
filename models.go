//go:build go1.18
// +build go1.18

// Code generated by Microsoft (R) AutoRest Code Generator (autorest: 3.9.6, generator: @autorest/go@4.0.0-preview.50)
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// Code generated by @autorest/go. DO NOT EDIT.

package azureopenai

import "time"

type ChatCompletionChoiceCommon struct {
	FinishReason *string
	Index        *int32
}

// ChatCompletionsClientCreateOptions contains the optional parameters for the ChatCompletionsClient.Create method.
type ChatCompletionsClientCreateOptions struct {
	// endpoint - server parameter
	Endpoint *string
}

type ChatCompletionsRequestCommon struct {
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far,
	// decreasing the model's likelihood to repeat the same line verbatim.
	FrequencyPenalty *float32

	// Modify the likelihood of specified tokens appearing in the completion. Accepts a json object that maps tokens (specified
	// by their token ID in the tokenizer) to an associated bias value from -100 to
	// 100. Mathematically, the bias is added to the logits generated by the model prior to sampling. The exact effect will vary
	// per model, but values between -1 and 1 should decrease or increase likelihood
	// of selection; values like -100 or 100 should result in a ban or exclusive selection of the relevant token.
	LogitBias any

	// The maximum number of tokens allowed for the generated answer. By default, the number of tokens the model can return will
	// be (4096 - prompt tokens).
	MaxTokens *int32

	// Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far, increasing
	// the model's likelihood to talk about new topics.
	PresencePenalty *float32

	// Up to 4 sequences where the API will stop generating further tokens.
	Stop *ChatCompletionsRequestCommonStop

	// If set, partial message deltas will be sent, like in ChatGPT. Tokens will be sent as data-only server-sent events as they
	// become available, with the stream terminated by a data: [DONE] message.
	Stream *bool

	// What sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random, while lower
	// values like 0.2 will make it more focused and deterministic. We generally
	// recommend altering this or top_p but not both.
	Temperature *float32

	// An alternative to sampling with temperature, called nucleus sampling, where the model considers the results of the tokens
	// with top_p probability mass. So 0.1 means only the tokens comprising the top
	// 10% probability mass are considered. We generally recommend altering this or temperature but not both.
	TopP *float32

	// A unique identifier representing your end-user, which can help Azure OpenAI to monitor and detect abuse.
	User *string
}

// ChatCompletionsRequestCommonStop - Up to 4 sequences where the API will stop generating further tokens.
type ChatCompletionsRequestCommonStop struct {
}

type ChatCompletionsResponseCommon struct {
	// REQUIRED
	Created *time.Time

	// REQUIRED
	ID *string

	// REQUIRED
	Model *string

	// REQUIRED
	Object *string
	Usage  *ChatCompletionsResponseCommonUsage
}

type ChatCompletionsResponseCommonUsage struct {
	// REQUIRED
	CompletionTokens *int32

	// REQUIRED
	PromptTokens *int32

	// REQUIRED
	TotalTokens *int32
}

// CompletionsClientCreateOptions contains the optional parameters for the CompletionsClient.Create method.
type CompletionsClientCreateOptions struct {
	// endpoint - server parameter
	Endpoint *string
}

type ContentFilterResult struct {
	// REQUIRED
	Filtered *bool

	// REQUIRED
	Severity *ContentFilterSeverity
}

// ContentFilterResults - Information about the content filtering category (hate, sexual, violence, selfharm), if it has been
// detected, as well as the severity level (verylow, low, medium, high-scale that determines the
// intensity and risk level of harmful content) and if it has been filtered or not.
type ContentFilterResults struct {
	Error    *ErrorBase
	Hate     *ContentFilterResult
	SelfHarm *ContentFilterResult
	Sexual   *ContentFilterResult
	Violence *ContentFilterResult
}

// DataSource - The data source to be used for the Azure OpenAI on your data feature.
type DataSource struct {
	// REQUIRED; The data source type.
	Type *string

	// The parameters to be used for the data source in runtime.
	Parameters map[string]any
}

// EmbeddingsClientCreateOptions contains the optional parameters for the EmbeddingsClient.Create method.
type EmbeddingsClientCreateOptions struct {
	// endpoint - server parameter
	Endpoint *string
}

type Error struct {
	Code *string

	// Inner error with additional details.
	InnerError *InnerError
	Message    *string
	Param      *string
	Type       *string
}

type ErrorBase struct {
	Code    *string
	Message *string
}

type ErrorResponse struct {
	Error *Error
}

type ExtensionsChatCompletionChoice struct {
	FinishReason *string
	Index        *int32

	// The list of messages returned by the service.
	Messages []*Message
}

// ExtensionsChatCompletionsClientCreateOptions contains the optional parameters for the ExtensionsChatCompletionsClient.Create
// method.
type ExtensionsChatCompletionsClientCreateOptions struct {
	// endpoint - server parameter
	Endpoint *string
}

// ExtensionsChatCompletionsRequest - Request for the chat completions using extensions
type ExtensionsChatCompletionsRequest struct {
	// REQUIRED
	Messages []*Message

	// The data sources to be used for the Azure OpenAI on your data feature.
	DataSources []*DataSource

	// Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far,
	// decreasing the model's likelihood to repeat the same line verbatim.
	FrequencyPenalty *float32

	// Modify the likelihood of specified tokens appearing in the completion. Accepts a json object that maps tokens (specified
	// by their token ID in the tokenizer) to an associated bias value from -100 to
	// 100. Mathematically, the bias is added to the logits generated by the model prior to sampling. The exact effect will vary
	// per model, but values between -1 and 1 should decrease or increase likelihood
	// of selection; values like -100 or 100 should result in a ban or exclusive selection of the relevant token.
	LogitBias any

	// The maximum number of tokens allowed for the generated answer. By default, the number of tokens the model can return will
	// be (4096 - prompt tokens).
	MaxTokens *int32

	// Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far, increasing
	// the model's likelihood to talk about new topics.
	PresencePenalty *float32

	// Up to 4 sequences where the API will stop generating further tokens.
	Stop *ChatCompletionsRequestCommonStop

	// If set, partial message deltas will be sent, like in ChatGPT. Tokens will be sent as data-only server-sent events as they
	// become available, with the stream terminated by a data: [DONE] message.
	Stream *bool

	// What sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random, while lower
	// values like 0.2 will make it more focused and deterministic. We generally
	// recommend altering this or top_p but not both.
	Temperature *float32

	// An alternative to sampling with temperature, called nucleus sampling, where the model considers the results of the tokens
	// with top_p probability mass. So 0.1 means only the tokens comprising the top
	// 10% probability mass are considered. We generally recommend altering this or temperature but not both.
	TopP *float32

	// A unique identifier representing your end-user, which can help Azure OpenAI to monitor and detect abuse.
	User *string
}

// ExtensionsChatCompletionsResponse - The response of the extensions chat completions.
type ExtensionsChatCompletionsResponse struct {
	// REQUIRED
	Created *time.Time

	// REQUIRED
	ID *string

	// REQUIRED
	Model *string

	// REQUIRED
	Object  *string
	Choices []*ExtensionsChatCompletionChoice
	Usage   *ChatCompletionsResponseCommonUsage
}

// InnerError - Inner error with additional details.
type InnerError struct {
	// Error codes for the inner error object.
	Code *InnerErrorCode

	// Information about the content filtering category (hate, sexual, violence, selfharm), if it has been detected, as well as
	// the severity level (verylow, low, medium, high-scale that determines the
	// intensity and risk level of harmful content) and if it has been filtered or not.
	ContentFilterResults *ContentFilterResults
}

// Message - A chat message.
type Message struct {
	// REQUIRED; The contents of the message
	Content *string

	// REQUIRED; The role of the author of this message.
	Role *MessageRole

	// Whether the message ends the turn.
	EndTurn *bool

	// The index of the message in the conversation.
	Index *int32

	// The recipient of the message in the format of .. Present if and only if the recipient is tool.
	Recipient *string
}

type Paths13PiqocDeploymentsDeploymentIDEmbeddingsPostRequestbodyContentApplicationJSONSchema struct {
	// REQUIRED; Input text to get embeddings for, encoded as a string. To get embeddings for multiple inputs in a single request,
	// pass an array of strings. Each input must not exceed 2048 tokens in length. Unless you
	// are embedding code, we suggest replacing newlines (\n) in your input with a single space, as we have observed inferior
	// results when newlines are present.
	Input *PostContentSchemaInput

	// OPTIONAL; Contains additional key/value pairs not defined in the schema.
	AdditionalProperties map[string]any

	// input type of embedding search to use
	InputType *string

	// A unique identifier representing your end-user, which can help monitoring and detecting abuse.
	User *string
}

type Paths15Cw454DeploymentsDeploymentIDEmbeddingsPostResponses200ContentApplicationJSONSchema struct {
	// REQUIRED
	Data []*Paths1Xmf2L5DeploymentsDeploymentIDEmbeddingsPostResponses200ContentApplicationJSONSchemaPropertiesDataItems

	// REQUIRED
	Model *string

	// REQUIRED
	Object *string

	// REQUIRED
	Usage *Paths18Tiy9VDeploymentsDeploymentIDEmbeddingsPostResponses200ContentApplicationJSONSchemaPropertiesUsage
}

type Paths18Tiy9VDeploymentsDeploymentIDEmbeddingsPostResponses200ContentApplicationJSONSchemaPropertiesUsage struct {
	// REQUIRED
	PromptTokens *int32

	// REQUIRED
	TotalTokens *int32
}

type Paths1G4Sf52DeploymentsDeploymentIDChatCompletionsPostResponses200ContentApplicationJSONSchemaPropertiesChoicesItems struct {
	// Information about the content filtering category (hate, sexual, violence, selfharm), if it has been detected, as well as
	// the severity level (verylow, low, medium, high-scale that determines the
	// intensity and risk level of harmful content) and if it has been filtered or not.
	ContentFilterResults *ContentFilterResults
	FinishReason         *string
	Index                *int32
	Message              *PostResponses200ContentApplicationJSONSchemaChoicesItemMessage
}

type Paths1H0F83DeploymentsDeploymentIDChatCompletionsPostResponses200ContentApplicationJSONSchema struct {
	// REQUIRED
	Choices []*Paths1G4Sf52DeploymentsDeploymentIDChatCompletionsPostResponses200ContentApplicationJSONSchemaPropertiesChoicesItems

	// REQUIRED
	Created *time.Time

	// REQUIRED
	ID *string

	// REQUIRED
	Model *string

	// REQUIRED
	Object *string

	// Content filtering results for zero or more prompts in the request. In a streaming request, results for different prompts
	// may arrive at different times or in different orders.
	PromptFilterResults []*PromptFilterResult
	Usage               *Paths1XcvledDeploymentsDeploymentIDChatCompletionsPostResponses200ContentApplicationJSONSchemaPropertiesUsage
}

type Paths1L1E8YpDeploymentsDeploymentIDChatCompletionsPostRequestbodyContentApplicationJSONSchema struct {
	// REQUIRED; The messages to generate chat completions for, in the chat format.
	Messages []*PostContentSchemaMessagesItem

	// Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far,
	// decreasing the model's likelihood to repeat the same line verbatim.
	FrequencyPenalty *float32

	// Modify the likelihood of specified tokens appearing in the completion. Accepts a json object that maps tokens (specified
	// by their token ID in the tokenizer) to an associated bias value from -100 to
	// 100. Mathematically, the bias is added to the logits generated by the model prior to sampling. The exact effect will vary
	// per model, but values between -1 and 1 should decrease or increase likelihood
	// of selection; values like -100 or 100 should result in a ban or exclusive selection of the relevant token.
	LogitBias any

	// The maximum number of tokens allowed for the generated answer. By default, the number of tokens the model can return will
	// be (4096 - prompt tokens).
	MaxTokens *int32

	// How many chat completion choices to generate for each input message.
	N *int32

	// Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far, increasing
	// the model's likelihood to talk about new topics.
	PresencePenalty *float32

	// Up to 4 sequences where the API will stop generating further tokens.
	Stop *PathsAcghv9DeploymentsDeploymentIDChatCompletionsPostRequestbodyContentApplicationJSONSchemaPropertiesStop

	// If set, partial message deltas will be sent, like in ChatGPT. Tokens will be sent as data-only server-sent events as they
	// become available, with the stream terminated by a data: [DONE] message.
	Stream *bool

	// What sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random, while lower
	// values like 0.2 will make it more focused and deterministic. We generally
	// recommend altering this or top_p but not both.
	Temperature *float32

	// An alternative to sampling with temperature, called nucleus sampling, where the model considers the results of the tokens
	// with top_p probability mass. So 0.1 means only the tokens comprising the top
	// 10% probability mass are considered. We generally recommend altering this or temperature but not both.
	TopP *float32

	// A unique identifier representing your end-user, which can help Azure OpenAI to monitor and detect abuse.
	User *string
}

type Paths1Phq6SvDeploymentsDeploymentIDCompletionsPostResponses200ContentApplicationJSONSchemaPropertiesUsage struct {
	// REQUIRED
	CompletionTokens *int32

	// REQUIRED
	PromptTokens *int32

	// REQUIRED
	TotalTokens *int32
}

type Paths1Vtxb06DeploymentsDeploymentIDCompletionsPostRequestbodyContentApplicationJSONSchema struct {
	// Generates bestof completions server-side and returns the "best" (the one with the highest log probability per token). Results
	// cannot be streamed. When used with n, bestof controls the number of
	// candidate completions and n specifies how many to return - bestof must be greater than n. Note: Because this parameter
	// generates many completions, it can quickly consume your token quota. Use
	// carefully and ensure that you have reasonable settings for maxtokens and stop. Has maximum value of 128.
	BestOf *int32

	// can be used to disable any server-side caching, 0=no cache, 1=prompt prefix enabled, 2=full cache
	CacheLevel       *int32
	CompletionConfig *string

	// Echo back the prompt in addition to the completion
	Echo *bool

	// Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far,
	// decreasing the model's likelihood to repeat the same line verbatim.
	FrequencyPenalty *float32

	// Defaults to null. Modify the likelihood of specified tokens appearing in the completion. Accepts a json object that maps
	// tokens (specified by their token ID in the GPT tokenizer) to an associated bias
	// value from -100 to 100. You can use this tokenizer tool (which works for both GPT-2 and GPT-3) to convert text to token
	// IDs. Mathematically, the bias is added to the logits generated by the model
	// prior to sampling. The exact effect will vary per model, but values between -1 and 1 should decrease or increase likelihood
	// of selection; values like -100 or 100 should result in a ban or exclusive
	// selection of the relevant token. As an example, you can pass {"50256" : -100} to prevent the token from being generated.
	LogitBias any

	// Include the log probabilities on the logprobs most likely tokens, as well the chosen tokens. For example, if logprobs is
	// 5, the API will return a list of the 5 most likely tokens. The API will always
	// return the logprob of the sampled token, so there may be up to logprobs+1 elements in the response. Minimum of 0 and maximum
	// of 5 allowed.
	Logprobs *int32

	// The token count of your prompt plus max_tokens cannot exceed the model's context length. Most models have a context length
	// of 2048 tokens (except for the newest models, which support 4096). Has
	// minimum of 0.
	MaxTokens *int32

	// How many completions to generate for each prompt. Minimum of 1 and maximum of 128 allowed. Note: Because this parameter
	// generates many completions, it can quickly consume your token quota. Use
	// carefully and ensure that you have reasonable settings for max_tokens and stop.
	N *int32

	// Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far, increasing
	// the model's likelihood to talk about new topics.
	PresencePenalty *float32

	// The prompt(s) to generate completions for, encoded as a string or array of strings. Note that is the document separator
	// that the model sees during training, so if a prompt is not specified the model
	// will generate as if from the beginning of a new document. Maximum allowed size of string list is 2048.
	Prompt *PostContentSchemaPrompt

	// Up to 4 sequences where the API will stop generating further tokens. The returned text will not contain the stop sequence.
	Stop *PostContentSchemaStop

	// Whether to stream back partial progress. If set, tokens will be sent as data-only server-sent events as they become available,
	// with the stream terminated by a data: [DONE] message.
	Stream *bool

	// The suffix that comes after a completion of inserted text.
	Suffix *string

	// What sampling temperature to use. Higher values means the model will take more risks. Try 0.9 for more creative applications,
	// and 0 (argmax sampling) for ones with a well-defined answer. We generally
	// recommend altering this or top_p but not both.
	Temperature *float32

	// An alternative to sampling with temperature, called nucleus sampling, where the model considers the results of the tokens
	// with top_p probability mass. So 0.1 means only the tokens comprising the top
	// 10% probability mass are considered. We generally recommend altering this or temperature but not both.
	TopP *float32

	// A unique identifier representing your end-user, which can help monitoring and detecting abuse
	User *string
}

type Paths1XcvledDeploymentsDeploymentIDChatCompletionsPostResponses200ContentApplicationJSONSchemaPropertiesUsage struct {
	// REQUIRED
	CompletionTokens *int32

	// REQUIRED
	PromptTokens *int32

	// REQUIRED
	TotalTokens *int32
}

type Paths1Xmf2L5DeploymentsDeploymentIDEmbeddingsPostResponses200ContentApplicationJSONSchemaPropertiesDataItems struct {
	// REQUIRED
	Embedding []*float32

	// REQUIRED
	Index *int32

	// REQUIRED
	Object *string
}

// PathsAcghv9DeploymentsDeploymentIDChatCompletionsPostRequestbodyContentApplicationJSONSchemaPropertiesStop - Up to 4 sequences
// where the API will stop generating further tokens.
type PathsAcghv9DeploymentsDeploymentIDChatCompletionsPostRequestbodyContentApplicationJSONSchemaPropertiesStop struct {
}

type PathsMaorw9DeploymentsDeploymentIDCompletionsPostResponses200ContentApplicationJSONSchema struct {
	// REQUIRED
	Choices []*Post200ApplicationJSONPropertiesItemsItem

	// REQUIRED
	Created *int32

	// REQUIRED
	ID *string

	// REQUIRED
	Model *string

	// REQUIRED
	Object *string

	// Content filtering results for zero or more prompts in the request. In a streaming request, results for different prompts
	// may arrive at different times or in different orders.
	PromptFilterResults []*PromptFilterResult
	Usage               *Paths1Phq6SvDeploymentsDeploymentIDCompletionsPostResponses200ContentApplicationJSONSchemaPropertiesUsage
}

type Post200ApplicationJSONPropertiesItemsItem struct {
	// Information about the content filtering category (hate, sexual, violence, selfharm), if it has been detected, as well as
	// the severity level (verylow, low, medium, high-scale that determines the
	// intensity and risk level of harmful content) and if it has been filtered or not.
	ContentFilterResults *ContentFilterResults
	FinishReason         *string
	Index                *int32
	Logprobs             *PostResponses200ContentApplicationJSONSchemaChoicesItemLogprobs
	Text                 *string
}

// PostContentSchemaInput - Input text to get embeddings for, encoded as a string. To get embeddings for multiple inputs in
// a single request, pass an array of strings. Each input must not exceed 2048 tokens in length. Unless you
// are embedding code, we suggest replacing newlines (\n) in your input with a single space, as we have observed inferior
// results when newlines are present.
type PostContentSchemaInput struct {
}

type PostContentSchemaMessagesItem struct {
	// REQUIRED; The contents of the message
	Content *string

	// REQUIRED; The role of the author of this message.
	Role *PostRequestBodyContentApplicationJSONSchemaMessagesItemRole

	// The name of the user in a multi-user chat
	Name *string
}

// PostContentSchemaPrompt - The prompt(s) to generate completions for, encoded as a string or array of strings. Note that
// is the document separator that the model sees during training, so if a prompt is not specified the model
// will generate as if from the beginning of a new document. Maximum allowed size of string list is 2048.
type PostContentSchemaPrompt struct {
}

// PostContentSchemaStop - Up to 4 sequences where the API will stop generating further tokens. The returned text will not
// contain the stop sequence.
type PostContentSchemaStop struct {
}

type PostResponses200ContentApplicationJSONSchemaChoicesItemLogprobs struct {
	TextOffset    []*int32
	TokenLogprobs []*float32
	Tokens        []*string
	TopLogprobs   []map[string]*float32
}

type PostResponses200ContentApplicationJSONSchemaChoicesItemMessage struct {
	// REQUIRED; The contents of the message
	Content *string

	// REQUIRED; The role of the author of this message.
	Role *Post200ApplicationJSONPropertiesItemsMessageRole
}

// PromptFilterResult - Content filtering results for a single prompt in the request.
type PromptFilterResult struct {
	// Information about the content filtering category (hate, sexual, violence, selfharm), if it has been detected, as well as
	// the severity level (verylow, low, medium, high-scale that determines the
	// intensity and risk level of harmful content) and if it has been filtered or not.
	ContentFilterResults *ContentFilterResults
	PromptIndex          *int32
}
