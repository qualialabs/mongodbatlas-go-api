package mongodbatlas

import (
  // auth "github.com/abbot/go-http-auth"
  "net/http"
  "log"
  // "encoding/json"
  // "fmt"
  "bytes"
  // "time"
  "io/ioutil"
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

func (client *Client) CreateMongodbUser(mongodbuser *MongodbUser) error {

  url := client.URL + "/groups/" + client.AtlasGroupId + "/databaseUsers"
  log.Println("[DEBUG] URL:>", url)
  db_user := mongodbuser.Username
  db_password := mongodbuser.Password
  db_name := mongodbuser.DatabaseName

  jsonStr := []byte(`{ "databaseName" : "admin", "roles" : [ {"databaseName" : "`+ db_name +`", "roleName" : "readWrite"}, { "databaseName" : "local", "roleName" : "read"} ],"username": "`+ db_user +`", "password":"`+ db_password + `"}`)

  req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
  if err != nil {
    return err
  }
  req.Header.Set("Content-Type", "application/json")
  req.SetBasicAuth("username", "apiKey")
  
  resp, err := client.HttpClient.Do(req)

  if err != nil {
    return err
  } else {
    body, _ := ioutil.ReadAll(resp.Body)
    log.Println("[DEBUG] Get response from Atlas response: ", string(body))
    return nil 
  }
}

// ############ the api not work #############
// func (client *Client) ReadMongodbUser(mongodb *Mongodb, user *User) error {

//   url := client.URL + "/deployments/" + mongodb.Account + "/" + mongodb.Deployment + "/mongodb/" + mongodb.Name + "/users"
//   log.Println("[DEBUG] URL:>", url)
//   req, err := http.NewRequest("GET", url, nil)
//   if err != nil {
//     return err
//   }
//   req.Header.Set("Content-Type", "application/json")
//   req.Header.Set("Accept-Version", "2014-06")
//   req.Header.Set("Authorization", "Bearer " + client.ComposeioToken)

//   resp, err := client.HttpClient.Do(req)

//   if err != nil {
//     return err
//   } else {
//     body, _ := ioutil.ReadAll(resp.Body)
//     log.Println("[DEBUG] Get response from Atlas response: ", string(body))
//     if !strings.Contains(string(body), user.Username) {
//       // err := "error"
//       log.Println("[DEBUG] "+  user.Username + " not found ")
//       return errors.New(user.Username + " not found")
//     } else {
//       return nil       
//     }
//   }

// }

// func (client *Client) UpdateMongodbUser(mongodbuser *MongodbUser) error {

  // url := client.URL + "/deployments/" + mongodb.Account + "/" + mongodb.Deployment + "/mongodb/" + mongodb.Name + "/users/" + user.Username
  // log.Println("[DEBUG] URL:>", url)

  // password := user.Password
  // var readOnly = "false"
  // if user.ReadOnly {
  //   readOnly = "true"
  // } 
  // jsonStr := []byte(`{"password":"`+ password + `", "readOnly":`+ readOnly + `}`)
  // req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
  // if err != nil {
  //   return err
  // }
  // req.Header.Set("Content-Type", "application/json")
  // req.Header.Set("Accept-Version", "2014-06")
  // req.Header.Set("Authorization", "Bearer " + client.ComposeioToken)

  
  // resp, err := client.HttpClient.Do(req)

  // if err != nil {
  //   return err
  // } else {
  //   body, _ := ioutil.ReadAll(resp.Body)
  //   log.Println("[DEBUG] Get response from Atlas response: ", string(body))
  //   return nil 
  // }
// }

// func (client *Client) DeleteMongodbUser(mongodbuser *MongodbUser) error {

  // url := client.URL + "/deployments/" + mongodb.Account + "/" + mongodb.Deployment + "/mongodb/" + mongodb.Name + "/users/" + user.Username
  // log.Println("[DEBUG] URL:>", url)
  // req, err := http.NewRequest("DELETE", url, nil )
  // if err != nil {
  //   return err
  // }
  // req.Header.Set("Content-Type", "application/json")
  // req.Header.Set("Accept-Version", "2014-06")
  // req.Header.Set("Authorization", "Bearer " + client.ComposeioToken)

  
  // resp, err := client.HttpClient.Do(req)

  // if err != nil {
  //   return err
  // } else {
  //   body, _ := ioutil.ReadAll(resp.Body)
  //   log.Println("[DEBUG] Get response from Atlas response: ", string(body))
  //   return nil 
  // }
// }



