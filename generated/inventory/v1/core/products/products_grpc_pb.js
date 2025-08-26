// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var inventory_v1_core_products_products_pb = require('../../../../inventory/v1/core/products/products_pb.js');
var inventory_v1_global_meta_meta_pb = require('../../../../inventory/v1/global/meta/meta_pb.js');

function serialize_CreateProductsRequest(arg) {
  if (!(arg instanceof inventory_v1_core_products_products_pb.CreateProductsRequest)) {
    throw new Error('Expected argument of type CreateProductsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_CreateProductsRequest(buffer_arg) {
  return inventory_v1_core_products_products_pb.CreateProductsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_CreateProductsResponse(arg) {
  if (!(arg instanceof inventory_v1_core_products_products_pb.CreateProductsResponse)) {
    throw new Error('Expected argument of type CreateProductsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_CreateProductsResponse(buffer_arg) {
  return inventory_v1_core_products_products_pb.CreateProductsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_DeleteProductsRequest(arg) {
  if (!(arg instanceof inventory_v1_core_products_products_pb.DeleteProductsRequest)) {
    throw new Error('Expected argument of type DeleteProductsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_DeleteProductsRequest(buffer_arg) {
  return inventory_v1_core_products_products_pb.DeleteProductsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_DeleteProductsResponse(arg) {
  if (!(arg instanceof inventory_v1_core_products_products_pb.DeleteProductsResponse)) {
    throw new Error('Expected argument of type DeleteProductsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_DeleteProductsResponse(buffer_arg) {
  return inventory_v1_core_products_products_pb.DeleteProductsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_GetProductsRequest(arg) {
  if (!(arg instanceof inventory_v1_core_products_products_pb.GetProductsRequest)) {
    throw new Error('Expected argument of type GetProductsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_GetProductsRequest(buffer_arg) {
  return inventory_v1_core_products_products_pb.GetProductsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_GetProductsResponse(arg) {
  if (!(arg instanceof inventory_v1_core_products_products_pb.GetProductsResponse)) {
    throw new Error('Expected argument of type GetProductsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_GetProductsResponse(buffer_arg) {
  return inventory_v1_core_products_products_pb.GetProductsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_ListProductssRequest(arg) {
  if (!(arg instanceof inventory_v1_core_products_products_pb.ListProductssRequest)) {
    throw new Error('Expected argument of type ListProductssRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ListProductssRequest(buffer_arg) {
  return inventory_v1_core_products_products_pb.ListProductssRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_ListProductssResponse(arg) {
  if (!(arg instanceof inventory_v1_core_products_products_pb.ListProductssResponse)) {
    throw new Error('Expected argument of type ListProductssResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ListProductssResponse(buffer_arg) {
  return inventory_v1_core_products_products_pb.ListProductssResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_UpdateProductsRequest(arg) {
  if (!(arg instanceof inventory_v1_core_products_products_pb.UpdateProductsRequest)) {
    throw new Error('Expected argument of type UpdateProductsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_UpdateProductsRequest(buffer_arg) {
  return inventory_v1_core_products_products_pb.UpdateProductsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_UpdateProductsResponse(arg) {
  if (!(arg instanceof inventory_v1_core_products_products_pb.UpdateProductsResponse)) {
    throw new Error('Expected argument of type UpdateProductsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_UpdateProductsResponse(buffer_arg) {
  return inventory_v1_core_products_products_pb.UpdateProductsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var ProductsServiceService = exports.ProductsServiceService = {
  createProducts: {
    path: '/ProductsService/CreateProducts',
    requestStream: false,
    responseStream: false,
    requestType: inventory_v1_core_products_products_pb.CreateProductsRequest,
    responseType: inventory_v1_core_products_products_pb.CreateProductsResponse,
    requestSerialize: serialize_CreateProductsRequest,
    requestDeserialize: deserialize_CreateProductsRequest,
    responseSerialize: serialize_CreateProductsResponse,
    responseDeserialize: deserialize_CreateProductsResponse,
  },
  getProducts: {
    path: '/ProductsService/GetProducts',
    requestStream: false,
    responseStream: false,
    requestType: inventory_v1_core_products_products_pb.GetProductsRequest,
    responseType: inventory_v1_core_products_products_pb.GetProductsResponse,
    requestSerialize: serialize_GetProductsRequest,
    requestDeserialize: deserialize_GetProductsRequest,
    responseSerialize: serialize_GetProductsResponse,
    responseDeserialize: deserialize_GetProductsResponse,
  },
  listProductss: {
    path: '/ProductsService/ListProductss',
    requestStream: false,
    responseStream: false,
    requestType: inventory_v1_core_products_products_pb.ListProductssRequest,
    responseType: inventory_v1_core_products_products_pb.ListProductssResponse,
    requestSerialize: serialize_ListProductssRequest,
    requestDeserialize: deserialize_ListProductssRequest,
    responseSerialize: serialize_ListProductssResponse,
    responseDeserialize: deserialize_ListProductssResponse,
  },
  updateProducts: {
    path: '/ProductsService/UpdateProducts',
    requestStream: false,
    responseStream: false,
    requestType: inventory_v1_core_products_products_pb.UpdateProductsRequest,
    responseType: inventory_v1_core_products_products_pb.UpdateProductsResponse,
    requestSerialize: serialize_UpdateProductsRequest,
    requestDeserialize: deserialize_UpdateProductsRequest,
    responseSerialize: serialize_UpdateProductsResponse,
    responseDeserialize: deserialize_UpdateProductsResponse,
  },
  deleteProducts: {
    path: '/ProductsService/DeleteProducts',
    requestStream: false,
    responseStream: false,
    requestType: inventory_v1_core_products_products_pb.DeleteProductsRequest,
    responseType: inventory_v1_core_products_products_pb.DeleteProductsResponse,
    requestSerialize: serialize_DeleteProductsRequest,
    requestDeserialize: deserialize_DeleteProductsRequest,
    responseSerialize: serialize_DeleteProductsResponse,
    responseDeserialize: deserialize_DeleteProductsResponse,
  },
};

exports.ProductsServiceClient = grpc.makeGenericClientConstructor(ProductsServiceService, 'ProductsService');
