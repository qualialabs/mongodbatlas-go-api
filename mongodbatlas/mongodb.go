package mongodbatlas

import (
  // auth "github.com/abbot/go-http-auth"
  // dac "github.com/xinsnake/go-http-digest-auth-client"
  // digest "github.com/bobziuchkovski/digest"
  "net/http"
  "log"
  "os/exec"
  "fmt"

  // "encoding/json"
  // "fmt"
  // "bytes"
  // "time"
  // "io/ioutil"
  // "strings"
  // "errors"

)

// ##############################
// Client is the object that handles talking to the Datadog API. This maintains
// state information for a particular application connection.
type Client struct {
  AtlasApiKey string
  AtlasUsername string
  AtlasGroupId string

  // URL to the Compose API to use
  URL string

  //The Http Client that is used to make requests
  HttpClient *http.Client
}

// NewClient returns a new Atlas.Client which can be used to access the API
// methods. The expected argument is the Atlas token.
func NewClient(atlas_api_key string, atlas_username string, atlas_group_id string) *Client {
  return &Client{
    AtlasApiKey: atlas_api_key,
    AtlasUsername: atlas_username,
    AtlasGroupId: atlas_group_id,
    URL: "https://cloud.mongodb.com/api/atlas/v1.0",
    HttpClient: http.DefaultClient,
  }
}

// ################################


type MongodbUser struct{
  Username string `json:"username,omitempty"`
  Password string `json:"password,omitempty"`
  DatabaseName string `json:"databasename,omitempty"`
}

// #############################




// this just create useer with readwirte role for db and read role for oplog
func (client *Client) CreateMongodbUser(mongodbuser *MongodbUser) error {
  db_user := mongodbuser.Username
  db_password := mongodbuser.Password
  db_name := mongodbuser.DatabaseName

  url := client.URL + "/groups/" + client.AtlasGroupId + "/databaseUsers"
  log.Println("[DEBUG] URL:", url)

  cmd := `curl -i -u "` + client.AtlasUsername+ `:`+client.AtlasApiKey +`" --digest -H "Content-Type: application/json" -X POST "`+ url +`" --data '
 {
   "databaseName" : "admin",
   "roles" : [ {
     "databaseName" : "` + db_name + `",
     "roleName" : "readWrite"
   }, {
     "databaseName" : "local",
     "roleName" : "read"
   } ],
   "username" : "` + db_user + `",
   "password" : "` + db_password + `"
 }'`

  out, err := exec.Command("sh","-c",cmd).Output()
  if err != nil {
    return err    
  }else {
    fmt.Printf("The response is %s\n", out)
    return nil
  }

}

func (client *Client) ReadMongodbUser(mongodbuser *MongodbUser) error {

  db_user := mongodbuser.Username
  // db_password := mongodbuser.Password
  // db_name := mongodbuser.DatabaseName
  
  url := client.URL + "/groups/" + client.AtlasGroupId + "/databaseUsers/admin/" + db_user
  log.Println("[DEBUG] URL:", url)


  cmd := `curl -i -u "` + client.AtlasUsername+ `:`+client.AtlasApiKey +`" --digest "`+ url +`"`



  out, err := exec.Command("sh","-c",cmd).Output()
  if err != nil {
    return err    
  }else {
    fmt.Printf("The response is %s\n", out)
    return nil
  }

  // # TODO: 
  // ## need to handle the result 
  // if !ok {
  //   d.SetId("")
  //   return nil
  // }

  // d.Set("address", obj.Address)
  // return nil
}

func (client *Client) UpdateMongodbUser(mongodbuser *MongodbUser) error {
  db_user := mongodbuser.Username
  db_password := mongodbuser.Password
  db_name := mongodbuser.DatabaseName
  
  url := client.URL + "/groups/" + client.AtlasGroupId + "/databaseUsers/admin/" + db_user
  log.Println("[DEBUG] URL:", url)


  cmd := `curl -i -u "` + client.AtlasUsername+ `:`+client.AtlasApiKey +`" --digest -H "Content-Type: application/json" -X PATCH "`+ url +`" --data '
 {
   "roles" : [ {
     "databaseName" : "` + db_name + `",
     "roleName" : "readWrite"
   }, {
     "databaseName" : "local",
     "roleName" : "read"
   } ],
   "password" : "` + db_password + `"
 }'`

  out, err := exec.Command("sh","-c",cmd).Output()
  if err != nil {
    return err    
  }else {
    fmt.Printf("The response is %s\n", out)
    return nil
  }
}

func (client *Client) DeleteMongodbUser(mongodbuser *MongodbUser) error {
  db_user := mongodbuser.Username
  // db_password := mongodbuser.Password
  // db_name := mongodbuser.DatabaseName

  url := client.URL + "/groups/" + client.AtlasGroupId + "/databaseUsers/admin/" + db_user
  log.Println("[DEBUG] URL:", url)

  cmd := `curl -i -u "` + client.AtlasUsername+ `:`+client.AtlasApiKey +`" --digest -X DELETE "`+ url +`"`

  out, err := exec.Command("sh","-c",cmd).Output()
  if err != nil {
    return err    
  }else {
    fmt.Printf("The response is %s\n", out)
    return nil
  }
}



