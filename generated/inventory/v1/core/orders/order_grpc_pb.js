// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var inventory_v1_core_orders_order_pb = require('../../../../inventory/v1/core/orders/order_pb.js');
var inventory_v1_global_meta_meta_pb = require('../../../../inventory/v1/global/meta/meta_pb.js');

function serialize_CancelOrderRequest(arg) {
  if (!(arg instanceof inventory_v1_core_orders_order_pb.CancelOrderRequest)) {
    throw new Error('Expected argument of type CancelOrderRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_CancelOrderRequest(buffer_arg) {
  return inventory_v1_core_orders_order_pb.CancelOrderRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_CancelOrderResponse(arg) {
  if (!(arg instanceof inventory_v1_core_orders_order_pb.CancelOrderResponse)) {
    throw new Error('Expected argument of type CancelOrderResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_CancelOrderResponse(buffer_arg) {
  return inventory_v1_core_orders_order_pb.CancelOrderResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_CreateOrderRequest(arg) {
  if (!(arg instanceof inventory_v1_core_orders_order_pb.CreateOrderRequest)) {
    throw new Error('Expected argument of type CreateOrderRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_CreateOrderRequest(buffer_arg) {
  return inventory_v1_core_orders_order_pb.CreateOrderRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_CreateOrderResponse(arg) {
  if (!(arg instanceof inventory_v1_core_orders_order_pb.CreateOrderResponse)) {
    throw new Error('Expected argument of type CreateOrderResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_CreateOrderResponse(buffer_arg) {
  return inventory_v1_core_orders_order_pb.CreateOrderResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_GetOrderRequest(arg) {
  if (!(arg instanceof inventory_v1_core_orders_order_pb.GetOrderRequest)) {
    throw new Error('Expected argument of type GetOrderRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_GetOrderRequest(buffer_arg) {
  return inventory_v1_core_orders_order_pb.GetOrderRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_GetOrderResponse(arg) {
  if (!(arg instanceof inventory_v1_core_orders_order_pb.GetOrderResponse)) {
    throw new Error('Expected argument of type GetOrderResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_GetOrderResponse(buffer_arg) {
  return inventory_v1_core_orders_order_pb.GetOrderResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var OrderServiceService = exports.OrderServiceService = {
  createOrder: {
    path: '/OrderService/CreateOrder',
    requestStream: false,
    responseStream: false,
    requestType: inventory_v1_core_orders_order_pb.CreateOrderRequest,
    responseType: inventory_v1_core_orders_order_pb.CreateOrderResponse,
    requestSerialize: serialize_CreateOrderRequest,
    requestDeserialize: deserialize_CreateOrderRequest,
    responseSerialize: serialize_CreateOrderResponse,
    responseDeserialize: deserialize_CreateOrderResponse,
  },
  cancelOrder: {
    path: '/OrderService/CancelOrder',
    requestStream: false,
    responseStream: false,
    requestType: inventory_v1_core_orders_order_pb.CancelOrderRequest,
    responseType: inventory_v1_core_orders_order_pb.CancelOrderResponse,
    requestSerialize: serialize_CancelOrderRequest,
    requestDeserialize: deserialize_CancelOrderRequest,
    responseSerialize: serialize_CancelOrderResponse,
    responseDeserialize: deserialize_CancelOrderResponse,
  },
  getOrder: {
    path: '/OrderService/GetOrder',
    requestStream: false,
    responseStream: false,
    requestType: inventory_v1_core_orders_order_pb.GetOrderRequest,
    responseType: inventory_v1_core_orders_order_pb.GetOrderResponse,
    requestSerialize: serialize_GetOrderRequest,
    requestDeserialize: deserialize_GetOrderRequest,
    responseSerialize: serialize_GetOrderResponse,
    responseDeserialize: deserialize_GetOrderResponse,
  },
};

exports.OrderServiceClient = grpc.makeGenericClientConstructor(OrderServiceService, 'OrderService');
