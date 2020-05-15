
// #node cc02_skill.js 
/** auto-gen from ulz/data/bin/adjust_skill.js 
* generate by : Steven Chm
* generate at : April 15th 2020, 3:53:56 pm
* !please manually modify the skill function
*/
const util = require('./util');
const _ = require('lodash');
const dataStruct = require('../proto_js/Data_pb');


/**
* Character binded  : cc02
* Skill feat_no     : 2 
* Skill feat_name   : rapidfire
*   @function {rapidfire}
*   @param {totalVal} $total_val
*   @param {number} $pow_val
*   @returns {totalVal, effect}
*/
var rapidfire = function ($total_val, $pow_val) {
    let _skill_id_ = 2;
    var totalVal = _.clone($total_val);
    var effectList = [];
    // * Code the skill below

    return {totalVal, effectList};
}
exports.rapidfire = rapidfire;

/**
* Character binded  : cc02
* Skill feat_no     : 10 
* Skill feat_name   : supersonicSword
*   @function {supersonicSword}
*   @param {totalVal} $total_val
*   @param {number} $pow_val
*   @returns {totalVal, effect}
*/
var supersonicSword = function ($total_val, $pow_val) {
    let _skill_id_ = 10;
    var totalVal = _.clone($total_val);
    var effectList = [];
    // * Code the skill below

    return {totalVal, effectList};
}
exports.supersonicSword = supersonicSword;

/**
* Character binded  : cc02
* Skill feat_no     : 11 
* Skill feat_name   : adrenalinAttack
*   @function {adrenalinAttack}
*   @param {totalVal} $total_val
*   @param {number} $pow_val
*   @returns {totalVal, effect}
*/
var adrenalinAttack = function ($total_val, $pow_val) {
    let _skill_id_ = 11;
    var totalVal = _.clone($total_val);
    var effectList = [];
    // * Code the skill below

    return {totalVal, effectList};
}
exports.adrenalinAttack = adrenalinAttack;

/**
* Character binded  : cc02
* Skill feat_no     : 39 
* Skill feat_name   : indomitableSpirit
*   @function {indomitableSpirit}
*   @param {totalVal} $total_val
*   @param {number} $pow_val
*   @returns {totalVal, effect}
*/
var indomitableSpirit = function ($total_val, $pow_val) {
    let _skill_id_ = 39;
    var totalVal = _.clone($total_val);
    var effectList = [];
    // * Code the skill below

    return {totalVal, effectList};
}
exports.indomitableSpirit = indomitableSpirit;

// export this module 
// !endl
