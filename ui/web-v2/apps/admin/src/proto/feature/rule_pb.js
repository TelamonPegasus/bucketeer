// source: proto/feature/rule.proto
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
var global = (function() { return this || window || global || self || Function('return this')(); }).call(null);

var proto_feature_clause_pb = require('../../proto/feature/clause_pb.js');
goog.object.extend(proto, proto_feature_clause_pb);
var proto_feature_strategy_pb = require('../../proto/feature/strategy_pb.js');
goog.object.extend(proto, proto_feature_strategy_pb);
goog.exportSymbol('proto.bucketeer.feature.Rule', null, global);
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
proto.bucketeer.feature.Rule = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.bucketeer.feature.Rule.repeatedFields_, null);
};
goog.inherits(proto.bucketeer.feature.Rule, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.bucketeer.feature.Rule.displayName = 'proto.bucketeer.feature.Rule';
}

/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.bucketeer.feature.Rule.repeatedFields_ = [3];



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
proto.bucketeer.feature.Rule.prototype.toObject = function(opt_includeInstance) {
  return proto.bucketeer.feature.Rule.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.bucketeer.feature.Rule} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.bucketeer.feature.Rule.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, ""),
    strategy: (f = msg.getStrategy()) && proto_feature_strategy_pb.Strategy.toObject(includeInstance, f),
    clausesList: jspb.Message.toObjectList(msg.getClausesList(),
    proto_feature_clause_pb.Clause.toObject, includeInstance)
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
 * @return {!proto.bucketeer.feature.Rule}
 */
proto.bucketeer.feature.Rule.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.bucketeer.feature.Rule;
  return proto.bucketeer.feature.Rule.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.bucketeer.feature.Rule} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.bucketeer.feature.Rule}
 */
proto.bucketeer.feature.Rule.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setId(value);
      break;
    case 2:
      var value = new proto_feature_strategy_pb.Strategy;
      reader.readMessage(value,proto_feature_strategy_pb.Strategy.deserializeBinaryFromReader);
      msg.setStrategy(value);
      break;
    case 3:
      var value = new proto_feature_clause_pb.Clause;
      reader.readMessage(value,proto_feature_clause_pb.Clause.deserializeBinaryFromReader);
      msg.addClauses(value);
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
proto.bucketeer.feature.Rule.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.bucketeer.feature.Rule.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.bucketeer.feature.Rule} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.bucketeer.feature.Rule.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getStrategy();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto_feature_strategy_pb.Strategy.serializeBinaryToWriter
    );
  }
  f = message.getClausesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      3,
      f,
      proto_feature_clause_pb.Clause.serializeBinaryToWriter
    );
  }
};


/**
 * optional string id = 1;
 * @return {string}
 */
proto.bucketeer.feature.Rule.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.bucketeer.feature.Rule} returns this
 */
proto.bucketeer.feature.Rule.prototype.setId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional Strategy strategy = 2;
 * @return {?proto.bucketeer.feature.Strategy}
 */
proto.bucketeer.feature.Rule.prototype.getStrategy = function() {
  return /** @type{?proto.bucketeer.feature.Strategy} */ (
    jspb.Message.getWrapperField(this, proto_feature_strategy_pb.Strategy, 2));
};


/**
 * @param {?proto.bucketeer.feature.Strategy|undefined} value
 * @return {!proto.bucketeer.feature.Rule} returns this
*/
proto.bucketeer.feature.Rule.prototype.setStrategy = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.bucketeer.feature.Rule} returns this
 */
proto.bucketeer.feature.Rule.prototype.clearStrategy = function() {
  return this.setStrategy(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.bucketeer.feature.Rule.prototype.hasStrategy = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * repeated Clause clauses = 3;
 * @return {!Array<!proto.bucketeer.feature.Clause>}
 */
proto.bucketeer.feature.Rule.prototype.getClausesList = function() {
  return /** @type{!Array<!proto.bucketeer.feature.Clause>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto_feature_clause_pb.Clause, 3));
};


/**
 * @param {!Array<!proto.bucketeer.feature.Clause>} value
 * @return {!proto.bucketeer.feature.Rule} returns this
*/
proto.bucketeer.feature.Rule.prototype.setClausesList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 3, value);
};


/**
 * @param {!proto.bucketeer.feature.Clause=} opt_value
 * @param {number=} opt_index
 * @return {!proto.bucketeer.feature.Clause}
 */
proto.bucketeer.feature.Rule.prototype.addClauses = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 3, opt_value, proto.bucketeer.feature.Clause, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.bucketeer.feature.Rule} returns this
 */
proto.bucketeer.feature.Rule.prototype.clearClausesList = function() {
  return this.setClausesList([]);
};


goog.object.extend(exports, proto.bucketeer.feature);
