// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hashmap

import (
	"encoding/json"
	"github.com/mymmsc/gox/util"
	"github.com/mymmsc/gox/util/containers"
)

func assertSerializationImplementation() {
	var _ containers.JSONSerializer = (*maps.Map)(nil)
	var _ containers.JSONDeserializer = (*maps.Map)(nil)
}

// ToJSON outputs the JSON representation of the map.
func (m *maps.Map) ToJSON() ([]byte, error) {
	elements := make(map[string]interface{})
	for key, value := range m.m {
		elements[maps.ToString(key)] = value
	}
	return json.Marshal(&elements)
}

// FromJSON populates the map from the input JSON representation.
func (m *maps.Map) FromJSON(data []byte) error {
	elements := make(map[string]interface{})
	err := json.Unmarshal(data, &elements)
	if err == nil {
		m.Clear()
		for key, value := range elements {
			m.m[key] = value
		}
	}
	return err
}