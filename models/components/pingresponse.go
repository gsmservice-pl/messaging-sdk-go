// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// PingResponse - 'Ping' response object
type PingResponse struct {
	// API service status: OK - API available, ERR - API unavailable
	Status *string `json:"status,omitempty"`
	// Current API Version
	Version *string `json:"version,omitempty"`
	// Was the connection established with the test system (SANDBOX)?
	Sandbox *bool `json:"sandbox,omitempty"`
}

func (o *PingResponse) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *PingResponse) GetVersion() *string {
	if o == nil {
		return nil
	}
	return o.Version
}

func (o *PingResponse) GetSandbox() *bool {
	if o == nil {
		return nil
	}
	return o.Sandbox
}
