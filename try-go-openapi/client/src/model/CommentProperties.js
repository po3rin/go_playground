/**
 * 匿名掲示板API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * OpenAPI spec version: 1.0.0
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 *
 * OpenAPI Generator version: 3.3.4
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'));
  } else {
    // Browser globals (root is window)
    if (!root.Api) {
      root.Api = {};
    }
    root.Api.CommentProperties = factory(root.Api.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';



  /**
   * The CommentProperties model module.
   * @module model/CommentProperties
   * @version 1.0.0
   */

  /**
   * Constructs a new <code>CommentProperties</code>.
   * @alias module:model/CommentProperties
   * @class
   * @param id {Number} 
   * @param content {String} 
   * @param postId {Number} 
   * @param commentedAt {Date} 
   */
  var exports = function(id, content, postId, commentedAt) {
    var _this = this;

    _this['id'] = id;
    _this['content'] = content;
    _this['post_id'] = postId;
    _this['commented_at'] = commentedAt;
  };

  /**
   * Constructs a <code>CommentProperties</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/CommentProperties} obj Optional instance to populate.
   * @return {module:model/CommentProperties} The populated <code>CommentProperties</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();
      if (data.hasOwnProperty('id')) {
        obj['id'] = ApiClient.convertToType(data['id'], 'Number');
      }
      if (data.hasOwnProperty('content')) {
        obj['content'] = ApiClient.convertToType(data['content'], 'String');
      }
      if (data.hasOwnProperty('post_id')) {
        obj['post_id'] = ApiClient.convertToType(data['post_id'], 'Number');
      }
      if (data.hasOwnProperty('commented_at')) {
        obj['commented_at'] = ApiClient.convertToType(data['commented_at'], 'Date');
      }
    }
    return obj;
  }

  /**
   * @member {Number} id
   */
  exports.prototype['id'] = undefined;
  /**
   * @member {String} content
   */
  exports.prototype['content'] = undefined;
  /**
   * @member {Number} post_id
   */
  exports.prototype['post_id'] = undefined;
  /**
   * @member {Date} commented_at
   */
  exports.prototype['commented_at'] = undefined;



  return exports;
}));


