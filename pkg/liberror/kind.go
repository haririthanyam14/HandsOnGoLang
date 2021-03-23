package liberror

type Kind string

const (
	InternalError         Kind = "internalError"
	ResourceNotFound      Kind = "resourceNotFound"
	ValidationError       Kind = "validationError"
	ResourceConflictError Kind = "resourceConflictError"

	BusinessError       Kind = "businessError"
	PaymentRefusedError Kind = "paymentRefusedError"
	DependencyError     Kind = "dependencyError"
)
