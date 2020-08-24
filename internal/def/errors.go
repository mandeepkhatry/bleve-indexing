package def

import "errors"

var (
	ErrIndexRegistered                  = errors.New("Index Already Registered")
	ErrIndexMappingUnregistered         = errors.New("Unregistered Index Mapping")
	ErrQueryServiceUnRegistered         = errors.New("Unregistered Query Service")
	ErrIndexRegisterServiceUnRegistered = errors.New("Unregistered Index Register Service")
	ErrIndexServiceUnRegistered         = errors.New("Unregistered Index Service")
)
