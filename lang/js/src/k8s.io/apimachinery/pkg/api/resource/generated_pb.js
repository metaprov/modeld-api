// source: k8s.io/apimachinery/pkg/api/resource/generated.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

goog.exportSymbol('proto.k8s.io.apimachinery.pkg.api.resource.Quantity', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.k8s.io.apimachinery.pkg.api.resource.Quantity = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.k8s.io.apimachinery.pkg.api.resource.Quantity, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.k8s.io.apimachinery.pkg.api.resource.Quantity.displayName = 'proto.k8s.io.apimachinery.pkg.api.resource.Quantity';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.k8s.io.apimachinery.pkg.api.resource.Quantity.prototype.toObject = function(opt_includeInstance) {
  return proto.k8s.io.apimachinery.pkg.api.resource.Quantity.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.k8s.io.apimachinery.pkg.api.resource.Quantity} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.k8s.io.apimachinery.pkg.api.resource.Quantity.toObject = function(includeInstance, msg) {
  var f, obj = {
    string: (f = jspb.Message.getField(msg, 1)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.k8s.io.apimachinery.pkg.api.resource.Quantity}
 */
proto.k8s.io.apimachinery.pkg.api.resource.Quantity.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.k8s.io.apimachinery.pkg.api.resource.Quantity;
  return proto.k8s.io.apimachinery.pkg.api.resource.Quantity.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.k8s.io.apimachinery.pkg.api.resource.Quantity} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.k8s.io.apimachinery.pkg.api.resource.Quantity}
 */
proto.k8s.io.apimachinery.pkg.api.resource.Quantity.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setString(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.k8s.io.apimachinery.pkg.api.resource.Quantity.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.k8s.io.apimachinery.pkg.api.resource.Quantity.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.k8s.io.apimachinery.pkg.api.resource.Quantity} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.k8s.io.apimachinery.pkg.api.resource.Quantity.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = /** @type {string} */ (jspb.Message.getField(message, 1));
  if (f != null) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string string = 1;
 * @return {string}
 */
proto.k8s.io.apimachinery.pkg.api.resource.Quantity.prototype.getString = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.k8s.io.apimachinery.pkg.api.resource.Quantity} returns this
 */
proto.k8s.io.apimachinery.pkg.api.resource.Quantity.prototype.setString = function(value) {
  return jspb.Message.setField(this, 1, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.k8s.io.apimachinery.pkg.api.resource.Quantity} returns this
 */
proto.k8s.io.apimachinery.pkg.api.resource.Quantity.prototype.clearString = function() {
  return jspb.Message.setField(this, 1, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.k8s.io.apimachinery.pkg.api.resource.Quantity.prototype.hasString = function() {
  return jspb.Message.getField(this, 1) != null;
};


goog.object.extend(exports, proto.k8s.io.apimachinery.pkg.api.resource);
