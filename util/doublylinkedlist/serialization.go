// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package doublylinkedlist

import (
	"encoding/json"
	"github.com/mymmsc/gox/util"
	"github.com/mymmsc/gox/util/containers"
)

func assertSerializationImplementation() {
	var _ containers.JSONSerializer = (*util.List)(nil)
	var _ containers.JSONDeserializer = (*util.List)(nil)
}

// ToJSON outputs the JSON representation of list's elements.
func (list *util.List) ToJSON() ([]byte, error) {
	return json.Marshal(list.Values())
}

// FromJSON populates list's elements from the input JSON representation.
func (list *util.List) FromJSON(data []byte) error {
	elements := []interface{}{}
	err := json.Unmarshal(data, &elements)
	if err == nil {
		list.Clear()
		list.Add(elements...)
	}
	return err
}