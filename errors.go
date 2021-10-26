package circleci

import "errors"

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrNotFound     = errors.New("not found")

	ErrRequiredEitherIDOrSlug                = errors.New("either organization ID or slug is required")
	ErrRequiredContextID                     = errors.New("context ID is required")
	ErrRequiredContextVariableName           = errors.New("environment variable name is required")
	ErrRequiredContextVariableValue          = errors.New("missing environment variable value")
	ErrRequiredProjectSlug                   = errors.New("project slug is required")
	ErrRequiredProjectCheckoutKeyType        = errors.New("project checkout key type is required")
	ErrRequiredProjectCheckoutKeyFingerprint = errors.New("project checkout key fingerprint is required")
	ErrRequiredProjectVariableName           = errors.New("project variable name is required")
	ErrRequiredProjectVariableValue          = errors.New("project variable value is required")
	ErrRequiredUsersUserID                   = errors.New("user id is required")
	ErrRequiredWorkflowsWorkflowID           = errors.New("workflow id is required")
	ErrRequiredWorkflowsApprovalRequestID    = errors.New("approval request id (the id of the job being approved) is required")
	ErrRequiredPipelineContinuationKey       = errors.New("pipeline continuation key is required")
	ErrRequiredPipelineConfiguration         = errors.New("pipeline configuration is required")
	ErrRequiredPipelinePipelineID            = errors.New("pipeline ID is required")
)
