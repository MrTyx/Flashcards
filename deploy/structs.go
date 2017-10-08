package main

type Data struct {
  LogoutURL string
  Email     string
  ID        string
  Order     int
  Due       int
  BucketURL string
  Progress  []PROGRESS
}

type STATS struct {
  LogoutURL string
  Email     string
  ID        string
  Users     int
  Studies   int
  SAverage  int
  Reviews   int
  RCount    int
  RAverage  int
  Hardest   []PROGRESS
  Easiest   []PROGRESS
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
  UID     string
  Code    string
  Due     int64
  Time    int64
  Reviews int
}
