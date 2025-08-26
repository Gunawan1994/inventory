// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var inventory_v1_core_inventory_stock_pb = require('../../../../inventory/v1/core/inventory/stock_pb.js');
var inventory_v1_global_meta_meta_pb = require('../../../../inventory/v1/global/meta/meta_pb.js');

function serialize_CheckStockRequest(arg) {
  if (!(arg instanceof inventory_v1_core_inventory_stock_pb.CheckStockRequest)) {
    throw new Error('Expected argument of type CheckStockRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_CheckStockRequest(buffer_arg) {
  return inventory_v1_core_inventory_stock_pb.CheckStockRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_CheckStockResponse(arg) {
  if (!(arg instanceof inventory_v1_core_inventory_stock_pb.CheckStockResponse)) {
    throw new Error('Expected argument of type CheckStockResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_CheckStockResponse(buffer_arg) {
  return inventory_v1_core_inventory_stock_pb.CheckStockResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_ReleaseStockRequest(arg) {
  if (!(arg instanceof inventory_v1_core_inventory_stock_pb.ReleaseStockRequest)) {
    throw new Error('Expected argument of type ReleaseStockRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ReleaseStockRequest(buffer_arg) {
  return inventory_v1_core_inventory_stock_pb.ReleaseStockRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_ReleaseStockResponse(arg) {
  if (!(arg instanceof inventory_v1_core_inventory_stock_pb.ReleaseStockResponse)) {
    throw new Error('Expected argument of type ReleaseStockResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ReleaseStockResponse(buffer_arg) {
  return inventory_v1_core_inventory_stock_pb.ReleaseStockResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_ReserveStockRequest(arg) {
  if (!(arg instanceof inventory_v1_core_inventory_stock_pb.ReserveStockRequest)) {
    throw new Error('Expected argument of type ReserveStockRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ReserveStockRequest(buffer_arg) {
  return inventory_v1_core_inventory_stock_pb.ReserveStockRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_ReserveStockResponse(arg) {
  if (!(arg instanceof inventory_v1_core_inventory_stock_pb.ReserveStockResponse)) {
    throw new Error('Expected argument of type ReserveStockResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ReserveStockResponse(buffer_arg) {
  return inventory_v1_core_inventory_stock_pb.ReserveStockResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var InventoryServiceService = exports.InventoryServiceService = {
  checkStock: {
    path: '/InventoryService/CheckStock',
    requestStream: false,
    responseStream: false,
    requestType: inventory_v1_core_inventory_stock_pb.CheckStockRequest,
    responseType: inventory_v1_core_inventory_stock_pb.CheckStockResponse,
    requestSerialize: serialize_CheckStockRequest,
    requestDeserialize: deserialize_CheckStockRequest,
    responseSerialize: serialize_CheckStockResponse,
    responseDeserialize: deserialize_CheckStockResponse,
  },
  reserveStock: {
    path: '/InventoryService/ReserveStock',
    requestStream: false,
    responseStream: false,
    requestType: inventory_v1_core_inventory_stock_pb.ReserveStockRequest,
    responseType: inventory_v1_core_inventory_stock_pb.ReserveStockResponse,
    requestSerialize: serialize_ReserveStockRequest,
    requestDeserialize: deserialize_ReserveStockRequest,
    responseSerialize: serialize_ReserveStockResponse,
    responseDeserialize: deserialize_ReserveStockResponse,
  },
  releaseStock: {
    path: '/InventoryService/ReleaseStock',
    requestStream: false,
    responseStream: false,
    requestType: inventory_v1_core_inventory_stock_pb.ReleaseStockRequest,
    responseType: inventory_v1_core_inventory_stock_pb.ReleaseStockResponse,
    requestSerialize: serialize_ReleaseStockRequest,
    requestDeserialize: deserialize_ReleaseStockRequest,
    responseSerialize: serialize_ReleaseStockResponse,
    responseDeserialize: deserialize_ReleaseStockResponse,
  },
};

exports.InventoryServiceClient = grpc.makeGenericClientConstructor(InventoryServiceService, 'InventoryService');
