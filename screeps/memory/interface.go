package memory

import (
	"bytes"
	"encoding/ascii85"
	"encoding/gob"
	"encoding/json"
	"fmt"
)

type Memory interface {
	Get() string
	Set(val string)
}

// GetBytes85 return data stored in m with SetBytes85
func GetBytes85(m Memory) ([]byte, error) {
	data := []byte(m.Get())

	decoded := make([]byte, len(data))
	length, _, err := ascii85.Decode(decoded, data, true)
	if err != nil {
		return nil, fmt.Errorf("error decoding string: %w", err)
	}

	return decoded[:length], nil
}

// SetBytes85 set the data in m as ascii85 encoded string
func SetBytes85(m Memory, v []byte) error {
	encoded := make([]byte, ascii85.MaxEncodedLen(len(v)))
	length := ascii85.Encode(encoded, v)
	str := string(encoded[:length])
	m.Set(str)
	return nil
}

// GetGob retrieve the gob-encoded value from m
func Get(m Memory, v interface{}) error {
	data, err := GetBytes85(m)
	if err != nil {
		return err
	}

	err = gob.NewDecoder(bytes.NewBuffer(data)).Decode(v)
	return err
}

// SetGob serialize the given value using gob and store it in m
func Set(m Memory, val interface{}) error {
	var buf bytes.Buffer
	err := gob.NewEncoder(&buf).Encode(val)
	if err != nil {
		return err
	}

	return SetBytes85(m, buf.Bytes())
}

// GetJSON retrieve the data stored in m into v
func GetJSON(m Memory, v interface{}) error {
	return json.Unmarshal([]byte(m.Get()), v)
}

// SetJSON set m to val, which must be marshalable with encoding/json
func SetJSON(m Memory, val interface{}) error {
	v, err := json.Marshal(val)
	if err != nil {
		return err
	}
	m.Set(string(v))
	return nil
}
