package sumologic

import (
    "encoding/json"
    "fmt"
)



func (s *Client) CreateUser(user User) (string, error) {
    data, err := s.Post("v1/users", user)
    if err != nil {
        return "", err
    }

    var createduser User
    err = json.Unmarshal(data, &createduser)
    if err != nil {
        return "", err
    }

    return createduser.ID, nil
}



func (s *Client) DeleteUser(id string) error {
    _, err := s.Delete(fmt.Sprintf("v1/users/%s", id))
    return err
}



func (s *Client) GetUser(id string) (*User, error) {
   data, _, err := s.Get(fmt.Sprintf("v1/users/%s", id))
   if err != nil {
       return nil, err
   }
   if data == nil {
       return nil, nil
   }

   var user User
   err = json.Unmarshal(data, &user)
   if err != nil {
       return nil, err
   }
   return &user, nil
}











func (s *Client) UpdateUser(user User) error {
    url := fmt.Sprintf("v1/users/%s", user.ID)

    user.ID = ""
      user.Email = ""

    _, err := s.Put(url, user)
    return err
}

// models
type User struct {
    ID        string   `json:"id,omitempty"`
    // First name of the user.
    FirstName  string `json:"firstName"`
    // Last name of the user.
    LastName  string `json:"lastName"`
    // Email address of the user.
    Email  string `json:"email"`
    // List of roleIds associated with the user.
    RoleIds  []string `json:"roleIds"`
    // True if the user is active.
    IsActive  bool `json:"isActive,omitempty"`
}
