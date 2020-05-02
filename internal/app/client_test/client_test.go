package client_test

import (
	"go-future/internal/app/client"
	"testing"
)

func TestClient(t *testing.T){
	t.Log(client.Sum(1,2,3,4))
}

