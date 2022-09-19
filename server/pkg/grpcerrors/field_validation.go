package grpcerrors

import (
	"fmt"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func FieldValidationErorr(
	fieldName string,
	description string,
	fieldValue string,
) (*status.Status, error) {
	st := status.New(codes.InvalidArgument, fmt.Sprintf("invalid %s", fieldName))
	v := &errdetails.BadRequest_FieldViolation{
		Field:       fieldName,
		Description: description,
	}
	badReq := &errdetails.BadRequest{
		FieldViolations: []*errdetails.BadRequest_FieldViolation{v},
	}
	debugInfo := &errdetails.DebugInfo{
		StackEntries: []string{fmt.Sprintf("field raw value %s", fieldValue)},
		Detail:       "",
	}
	st, err := st.WithDetails(badReq)
	if err != nil {
		return nil, err
	}
	st, err = st.WithDetails(debugInfo)
	if err != nil {
		return nil, err
	}
	return st, nil

}
