package main

type Data struct {
  LogoutURL string
  Email string
  ID string
  Order int
}

type USER struct {
  UID   string
  Order int
}

type FLAG struct {
  Order int
  Code  string
  Name  string
}

type PROGRESS struct {
  UID  string
  Code string
  Due  int64
  Time int64
}
