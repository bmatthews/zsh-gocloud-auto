package persister

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"sync"
)

type Persister interface {
	Save(path string, v interface{}) error
	Load(path string, v interface{}) error
}

type DiskPersister struct {
	lock *sync.Mutex
}

//NewDiskPersister saves and loads structs from disk
func NewDiskPersister() *DiskPersister {
	var lock sync.Mutex
	return &DiskPersister{lock: &lock}
}

//Save - saves struct to disk
func (d *DiskPersister) Save(path string, v interface{}) error {
	d.lock.Lock()
	defer d.lock.Unlock()
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	r, err := d.marshal(v)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, r)
	return err
}

//Load - loads struct to disk
func (d *DiskPersister) Load(path string, v interface{}) error {
	d.lock.Lock()
	defer d.lock.Unlock()
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return d.unmarshal(f, v)
}

func (d *DiskPersister) marshal(v interface{}) (io.Reader, error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func (d *DiskPersister) unmarshal(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}
