
// #node cc05_skill.js 
/** auto-gen from ulz/data/bin/adjust_skill.js 
* generate by : Steven Chm
* generate at : April 15th 2020, 3:53:56 pm
* !please manually modify the skill function
*/
const util = require('./util');
const _ = require('lodash');
const dataStruct = require('../proto_js/Data_pb');


/**
* Character binded  : cc05
* Skill feat_no     : 13 
* Skill feat_name   : shadowStrike
*   @function {shadowStrike}
*   @param {totalVal} $total_val
*   @param {number} $pow_val
*   @returns {totalVal, effect}
*/
var shadowStrike = function ($total_val, $pow_val) {
    let _skill_id_ = 13;
    var totalVal = _.clone($total_val);
    var effectList = [];
    // * Code the skill below

    return {totalVal, effectList};
}
exports.shadowStrike = shadowStrike;

/**
* Character binded  : cc05
* Skill feat_no     : 14 
* Skill feat_name   : redFang
*   @function {redFang}
*   @param {totalVal} $total_val
*   @param {number} $pow_val
*   @returns {totalVal, effect}
*/
var redFang = function ($total_val, $pow_val) {
    let _skill_id_ = 14;
    var totalVal = _.clone($total_val);
    var effectList = [];
    // * Code the skill below

    return {totalVal, effectList};
}
exports.redFang = redFang;

/**
* Character binded  : cc05
* Skill feat_no     : 16 
* Skill feat_name   : rocketCounterAttack
*   @function {rocketCounterAttack}
*   @param {totalVal} $total_val
*   @param {number} $pow_val
*   @returns {totalVal, effect}
*/
var rocketCounterAttack = function ($total_val, $pow_val) {
    let _skill_id_ = 16;
    var totalVal = _.clone($total_val);
    var effectList = [];
    // * Code the skill below

    return {totalVal, effectList};
}
exports.rocketCounterAttack = rocketCounterAttack;

/**
* Character binded  : cc05
* Skill feat_no     : 41 
* Skill feat_name   : backstab
*   @function {backstab}
*   @param {totalVal} $total_val
*   @param {number} $pow_val
*   @returns {totalVal, effect}
*/
var backstab = function ($total_val, $pow_val) {
    let _skill_id_ = 41;
    var totalVal = _.clone($total_val);
    var effectList = [];
    // * Code the skill below

    return {totalVal, effectList};
}
exports.backstab = backstab;

// export this module 
// !endl
