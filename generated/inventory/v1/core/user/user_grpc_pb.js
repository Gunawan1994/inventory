// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var inventory_v1_core_user_user_pb = require('../../../../inventory/v1/core/user/user_pb.js');
var inventory_v1_global_meta_meta_pb = require('../../../../inventory/v1/global/meta/meta_pb.js');

function serialize_CreateUserRequest(arg) {
  if (!(arg instanceof inventory_v1_core_user_user_pb.CreateUserRequest)) {
    throw new Error('Expected argument of type CreateUserRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_CreateUserRequest(buffer_arg) {
  return inventory_v1_core_user_user_pb.CreateUserRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_CreateUserResponse(arg) {
  if (!(arg instanceof inventory_v1_core_user_user_pb.CreateUserResponse)) {
    throw new Error('Expected argument of type CreateUserResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_CreateUserResponse(buffer_arg) {
  return inventory_v1_core_user_user_pb.CreateUserResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_DeleteUserRequest(arg) {
  if (!(arg instanceof inventory_v1_core_user_user_pb.DeleteUserRequest)) {
    throw new Error('Expected argument of type DeleteUserRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_DeleteUserRequest(buffer_arg) {
  return inventory_v1_core_user_user_pb.DeleteUserRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_DeleteUserResponse(arg) {
  if (!(arg instanceof inventory_v1_core_user_user_pb.DeleteUserResponse)) {
    throw new Error('Expected argument of type DeleteUserResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_DeleteUserResponse(buffer_arg) {
  return inventory_v1_core_user_user_pb.DeleteUserResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_GetUserRequest(arg) {
  if (!(arg instanceof inventory_v1_core_user_user_pb.GetUserRequest)) {
    throw new Error('Expected argument of type GetUserRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_GetUserRequest(buffer_arg) {
  return inventory_v1_core_user_user_pb.GetUserRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_GetUserResponse(arg) {
  if (!(arg instanceof inventory_v1_core_user_user_pb.GetUserResponse)) {
    throw new Error('Expected argument of type GetUserResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_GetUserResponse(buffer_arg) {
  return inventory_v1_core_user_user_pb.GetUserResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_ListUsersRequest(arg) {
  if (!(arg instanceof inventory_v1_core_user_user_pb.ListUsersRequest)) {
    throw new Error('Expected argument of type ListUsersRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ListUsersRequest(buffer_arg) {
  return inventory_v1_core_user_user_pb.ListUsersRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_ListUsersResponse(arg) {
  if (!(arg instanceof inventory_v1_core_user_user_pb.ListUsersResponse)) {
    throw new Error('Expected argument of type ListUsersResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ListUsersResponse(buffer_arg) {
  return inventory_v1_core_user_user_pb.ListUsersResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_UpdateUserRequest(arg) {
  if (!(arg instanceof inventory_v1_core_user_user_pb.UpdateUserRequest)) {
    throw new Error('Expected argument of type UpdateUserRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_UpdateUserRequest(buffer_arg) {
  return inventory_v1_core_user_user_pb.UpdateUserRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_UpdateUserResponse(arg) {
  if (!(arg instanceof inventory_v1_core_user_user_pb.UpdateUserResponse)) {
    throw new Error('Expected argument of type UpdateUserResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_UpdateUserResponse(buffer_arg) {
  return inventory_v1_core_user_user_pb.UpdateUserResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var UserServiceService = exports.UserServiceService = {
  createUser: {
    path: '/UserService/CreateUser',
    requestStream: false,
    responseStream: false,
    requestType: inventory_v1_core_user_user_pb.CreateUserRequest,
    responseType: inventory_v1_core_user_user_pb.CreateUserResponse,
    requestSerialize: serialize_CreateUserRequest,
    requestDeserialize: deserialize_CreateUserRequest,
    responseSerialize: serialize_CreateUserResponse,
    responseDeserialize: deserialize_CreateUserResponse,
  },
  getUser: {
    path: '/UserService/GetUser',
    requestStream: false,
    responseStream: false,
    requestType: inventory_v1_core_user_user_pb.GetUserRequest,
    responseType: inventory_v1_core_user_user_pb.GetUserResponse,
    requestSerialize: serialize_GetUserRequest,
    requestDeserialize: deserialize_GetUserRequest,
    responseSerialize: serialize_GetUserResponse,
    responseDeserialize: deserialize_GetUserResponse,
  },
  listUsers: {
    path: '/UserService/ListUsers',
    requestStream: false,
    responseStream: false,
    requestType: inventory_v1_core_user_user_pb.ListUsersRequest,
    responseType: inventory_v1_core_user_user_pb.ListUsersResponse,
    requestSerialize: serialize_ListUsersRequest,
    requestDeserialize: deserialize_ListUsersRequest,
    responseSerialize: serialize_ListUsersResponse,
    responseDeserialize: deserialize_ListUsersResponse,
  },
  updateArticle: {
    path: '/UserService/UpdateArticle',
    requestStream: false,
    responseStream: false,
    requestType: inventory_v1_core_user_user_pb.UpdateUserRequest,
    responseType: inventory_v1_core_user_user_pb.UpdateUserResponse,
    requestSerialize: serialize_UpdateUserRequest,
    requestDeserialize: deserialize_UpdateUserRequest,
    responseSerialize: serialize_UpdateUserResponse,
    responseDeserialize: deserialize_UpdateUserResponse,
  },
  deleteArticle: {
    path: '/UserService/DeleteArticle',
    requestStream: false,
    responseStream: false,
    requestType: inventory_v1_core_user_user_pb.DeleteUserRequest,
    responseType: inventory_v1_core_user_user_pb.DeleteUserResponse,
    requestSerialize: serialize_DeleteUserRequest,
    requestDeserialize: deserialize_DeleteUserRequest,
    responseSerialize: serialize_DeleteUserResponse,
    responseDeserialize: deserialize_DeleteUserResponse,
  },
};

exports.UserServiceClient = grpc.makeGenericClientConstructor(UserServiceService, 'UserService');
