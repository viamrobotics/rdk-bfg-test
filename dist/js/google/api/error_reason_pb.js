// source: google/api/error_reason.proto
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

goog.exportSymbol('proto.google.api.ErrorReason', null, global);
/**
 * @enum {number}
 */
proto.google.api.ErrorReason = {
  ERROR_REASON_UNSPECIFIED: 0,
  SERVICE_DISABLED: 1,
  BILLING_DISABLED: 2,
  API_KEY_INVALID: 3,
  API_KEY_SERVICE_BLOCKED: 4,
  API_KEY_HTTP_REFERRER_BLOCKED: 7,
  API_KEY_IP_ADDRESS_BLOCKED: 8,
  API_KEY_ANDROID_APP_BLOCKED: 9,
  API_KEY_IOS_APP_BLOCKED: 13,
  RATE_LIMIT_EXCEEDED: 5,
  RESOURCE_QUOTA_EXCEEDED: 6,
  LOCATION_TAX_POLICY_VIOLATED: 10,
  USER_PROJECT_DENIED: 11,
  CONSUMER_SUSPENDED: 12,
  CONSUMER_INVALID: 14,
  SECURITY_POLICY_VIOLATED: 15,
  ACCESS_TOKEN_EXPIRED: 16
};

goog.object.extend(exports, proto.google.api);
