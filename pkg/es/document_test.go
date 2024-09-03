package es

import (
	"github.com/lbzss/nodeInfo/pkg/node"
	"testing"
)

func TestClient_CreateDocument(t *testing.T) {
	type args struct {
		indexName string
		address   []string
		username  string
		password  string
	}
	arg := args{
		indexName: "test",
		address:   []string{"http://10.89.64.76:9200"},
		username:  "elastic",
		password:  "P@ssw0rd",
	}

	t.Run("createDocument", func(t *testing.T) {
		es, err := NewEsClient(arg.address, arg.username, arg.password)
		if err != nil {
			t.Fatal(err)
		}
		data := node.NewNode()
		data.Collect()
		data.Complete()
		err = es.CreateDocument(arg.indexName, data)
		if err != nil {
			t.Errorf("CreateDocument() error = %v", err)
		}
	})

	t.Run("getDocument", func(t *testing.T) {
		es, err := NewEsClient(arg.address, arg.username, arg.password)
		if err != nil {
			t.Fatal(err)
		}
		data := node.NewNode()
		data.Collect()
		data.Complete()
		err = es.GetDocument(arg.indexName)
		if err != nil {
			t.Errorf("GetDocument() error = %v", err)
		}
	})
}
