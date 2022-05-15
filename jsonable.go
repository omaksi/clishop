package rdg

import "encoding/json"

/* Json serialization struct */
type Jsonable struct {
}

func (u *Jsonable) ToJson() string {
	b, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	return string(b)
}
