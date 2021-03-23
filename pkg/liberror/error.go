package liberror

import (
	"fmt"
	"strings"
)

//TODO: ADD TEST CASE
type Error struct {
	kind      Kind
	operation Operation
	severity  Severity
	err       error
	data      interface{}
}

// RETURNS THE FIRST KIND IN THE STACK
func (e *Error) Kind() Kind {
	if e == nil {
		return ""
	}

	if len(e.kind) != 0 {
		return e.kind
	}

	if e.err == nil {
		return ""
	}

	t, ok := e.err.(*Error)
	if !ok {
		return ""
	}

	return t.Kind()
}

//TODO: REFACTOR
func (e *Error) Details() string {
	var encode func(d map[string]interface{}) string
	encode = func(d map[string]interface{}) string {

		b := new(strings.Builder)

		for k, v := range d {
			t, ok := v.(map[string]interface{})

			var vl string

			if ok {
				vl = encode(t)
			} else {
				vl = fmt.Sprintf("%s", v)
			}

			b.WriteString(fmt.Sprintf(" %s:%s ", k, vl))
		}

		return fmt.Sprintf("[%s]", strings.TrimSpace(b.String()))
	}

	return encode(e.detailsMap())
}

func (e *Error) detailsMap() map[string]interface{} {
	r := make(map[string]interface{})

	if len(e.kind) != 0 {
		r["kind"] = e.kind
	}

	if len(e.operation) != 0 {
		r["operation"] = e.operation
	}

	if len(e.severity) != 0 {
		r["severity"] = e.severity
	}

	if e.err != nil {
		t, ok := e.err.(*Error)
		if ok {
			r["cause"] = t.detailsMap()
		} else {
			r["cause"] = e.err
		}
	}

	if e.data != nil {
		r["data"] = e.data
	}

	return r
}

func (e *Error) Operation() Operation {
	return e.operation
}

func (e *Error) Operations() []Operation {
	var ops []Operation
	if e == nil || e.err == nil {
		return ops
	}

	ops = append(ops, e.operation)

	t, ok := e.err.(*Error)
	if ok {
		ops = append(ops, t.Operations()...)
	}

	return ops

}

func (e *Error) Severity() Severity {
	return e.severity
}

func (e *Error) Unwrap() error {
	return e.err
}

func (e *Error) Error() string {
	if e.err == nil {
		return ""
	}

	return e.err.Error()
}

func (e *Error) Data() interface{} {
	return e.data
}

func NewError(kind Kind, operation Operation, severity Severity, err error, data interface{}) *Error {
	return &Error{
		kind:      kind,
		operation: operation,
		severity:  severity,
		err:       err,
		data:      data,
	}
}

//TODO: REMOVE BUILDER PATTERN
type ErrorBuilder struct {
	kind      Kind
	operation Operation
	severity  Severity
	err       error
	data      interface{}
}

func Builder() *ErrorBuilder {
	return &ErrorBuilder{}
}

func (eb *ErrorBuilder) SetKind(kind Kind) *ErrorBuilder {
	eb.kind = kind
	return eb
}

func (eb *ErrorBuilder) SetSeverity(severity Severity) *ErrorBuilder {
	eb.severity = severity
	return eb
}

func (eb *ErrorBuilder) SetOperation(operation Operation) *ErrorBuilder {
	eb.operation = operation
	return eb
}

func (eb *ErrorBuilder) SetCause(err error) *ErrorBuilder {
	eb.err = err
	return eb
}

func (eb *ErrorBuilder) SetData(data interface{}) *ErrorBuilder {
	eb.data = data
	return eb
}

func (eb *ErrorBuilder) Build() *Error {
	return NewError(eb.kind, eb.operation, eb.severity, eb.err, eb.data)
}
