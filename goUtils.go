package goUtils
import "sync"

//import "time"
//import "log"
import "io"
import "os"
import "encoding/json"
import "bytes"
import "fmt"

var lock sync.Mutex

// Load loads the file at path into v.
// Use os.IsNotExist() to see if the returned error is due
// to the file being missing.
func Load(path string, v interface{}) error {
  lock.Lock()
  defer lock.Unlock()
  f, err := os.Open(path)
  if err != nil {
    return err
  }
  defer f.Close()
  return Unmarshal(f, &v)
}

// Unmarshal is a function that unmarshals the data from the
// reader into the specified value.
// By default, it uses the JSON unmarshaller.
var Unmarshal = func(r io.Reader, v interface{}) error {
  return json.NewDecoder(r).Decode(v)
}

// Save saves a representation of v to the file at path.
func Save(path string, v interface{}) error {
  lock.Lock()
  defer lock.Unlock()
  f, err := os.Create(path)
  if err != nil {
    return err
  }
  defer f.Close()
  r, err := Marshal(v)
  if err != nil {
    return err
  }
  _, err = io.Copy(f, r)
  return err
}

// Marshal is a function that marshals the object into an
// io.Reader.
// By default, it uses the JSON marshaller.
var Marshal = func(v interface{}) (io.Reader, error) {
  b, err := json.MarshalIndent(v, "", "\t")
  if err != nil {
    return nil, err
  }
  return bytes.NewReader(b), nil
}

func Describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
