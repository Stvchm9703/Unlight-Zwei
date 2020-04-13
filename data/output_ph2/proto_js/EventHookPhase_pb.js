/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

goog.exportSymbol('proto.ULZProto.EventHookPhase', null, global);
goog.exportSymbol('proto.ULZProto.EventHookType', null, global);
/**
 * @enum {number}
 */
proto.ULZProto.EventHookPhase = {
  GAMESET_START: 0,
  START_TURN_PHASE: 1,
  REFILL_ACTION_CARD_PHASE: 2,
  MOVE_CARD_DROP_PHASE: 3,
  DETERMINE_MOVE_PHASE: 4,
  FINISH_MOVE_PHASE: 5,
  CHARA_CHANGE_PHASE: 6,
  DETERMINE_CHARA_CHANGE_PHASE: 7,
  ATTACK_CARD_DROP_PHASE: 8,
  DEFENCE_CARD_DROP_PHASE: 9,
  DETERMINE_BATTLE_POINT_PHASE: 10,
  BATTLE_RESULT_PHASE: 11,
  DAMAGE_PHASE: 12,
  DEAD_CHARA_CHANGE_PHASE: 13,
  DETERMINE_DEAD_CHARA_CHANGE_PHASE: 14,
  CHANGE_INITIATIVE_PHASE: 15,
  FINISH_TURN_PHASE: 16,
  GAMESET_END: 17
};

/**
 * @enum {number}
 */
proto.ULZProto.EventHookType = {
  INSTANT: 0,
  BEFORE: 1,
  PROXY: 2,
  AFTER: 3
};

goog.object.extend(exports, proto.ULZProto);
