// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var inventory_v1_core_auth_auth_pb = require('../../../../inventory/v1/core/auth/auth_pb.js');
var inventory_v1_global_meta_meta_pb = require('../../../../inventory/v1/global/meta/meta_pb.js');

function serialize_LoginRequest(arg) {
  if (!(arg instanceof inventory_v1_core_auth_auth_pb.LoginRequest)) {
    throw new Error('Expected argument of type LoginRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_LoginRequest(buffer_arg) {
  return inventory_v1_core_auth_auth_pb.LoginRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_LoginResponse(arg) {
  if (!(arg instanceof inventory_v1_core_auth_auth_pb.LoginResponse)) {
    throw new Error('Expected argument of type LoginResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_LoginResponse(buffer_arg) {
  return inventory_v1_core_auth_auth_pb.LoginResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var AuthServiceService = exports.AuthServiceService = {
  login: {
    path: '/AuthService/Login',
    requestStream: false,
    responseStream: false,
    requestType: inventory_v1_core_auth_auth_pb.LoginRequest,
    responseType: inventory_v1_core_auth_auth_pb.LoginResponse,
    requestSerialize: serialize_LoginRequest,
    requestDeserialize: deserialize_LoginRequest,
    responseSerialize: serialize_LoginResponse,
    responseDeserialize: deserialize_LoginResponse,
  },
};

exports.AuthServiceClient = grpc.makeGenericClientConstructor(AuthServiceService, 'AuthService');
