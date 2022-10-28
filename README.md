# spaceuid

A simple golang interface for identity oriented programming




## Installing

### using go get

```sh
go get github.com/gitsakos/spaceuid
```


## Usage Example

To turn your custom typed ID into a SpaceUID, simply implement the SpaceUID interface.


```go
type UserID int64

func (id UserID) New(idStr string) (spaceuid.SpaceUID, error) {
    i, err := strconv.ParseInt(idStr, 10, 64)
    return UserID(i), err
}

func (id UserID) Space() string  { return "user" }

func (id UserID) SUID() string   { return strconv.FormatInt(int64(id), 10) }

func (id UserID) String() string { return id.Space() + "_" + id.SUID() }`

```



Maintain a map of all your domain identities and convert easily beteen a string pair and SpaceUID.

```go
var Identities = spaceuid.SpaceUIDsMap{
    ReleaseID(0).Space(): ReleaseID(0),
    UserID(0).Space():    UserID(0),
}


id, err := Identities.GetID("user":1234)
```
