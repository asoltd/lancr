package server

import (
	"encoding/json"
	"os"

	"cloud.google.com/go/firestore"
	lancrv1 "github.com/asoltd/lancr/gen/go/lancr/v1"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func (b *Backend) snapshotToMessage(snapshot *firestore.DocumentSnapshot, pb *lancrv1.Hero) error {
	// Marshal snapshot data into JSON
	jsonBytes, err := json.Marshal(snapshot.Data())
	if err != nil {
		return err
	}

	// Unmarshal JSON into protobuf struct
	err = protojson.Unmarshal(jsonBytes, pb)
	if err != nil {
		return err
	}

	return nil
}

// messageToStruct is to be used for passing on the protobufs into firestore it
// is suboptimal and should be replaced with a more efficient solution, does
// Firestore support protobufs directly? maybe in the transport layer of the
// grpc package?
func messageToStruct(pb proto.Message) (map[string]any, error) {
	var res map[string]any

	// Marshal protobuf struct into JSON
	b, err := protojson.Marshal(pb)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON into map of any
	err = json.Unmarshal(b, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func UnmarshalJSONFileToMessage(filepath string, pb proto.Message) error {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	// Unmarshal the JSON data into the Protocol Buffers message
	if err := protojson.Unmarshal(data, pb); err != nil {
		return err
	}

	return nil
}
