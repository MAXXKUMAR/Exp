syntax = "proto3";

package models;

import "google.golang.org/protobuf/proto";

message Mode {
  string name = 1;
  int32 count = 2;
}

