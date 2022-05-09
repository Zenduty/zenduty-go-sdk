package client

import "testing"

func TestClient_DeleteService(t *testing.T) {
	type args struct {
		team string
		id   string
	}
	tests := []struct {
		name    string
		c       *Client
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Services.DeleteService(tt.args.team, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Client.DeleteService() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
