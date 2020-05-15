
// #node cc04_skill.js 
/** auto-gen from ulz/data/bin/adjust_skill.js 
* generate by : Steven Chm
* generate at : April 15th 2020, 3:53:56 pm
* !please manually modify the skill function
*/
const util = require('./util');
const _ = require('lodash');
const dataStruct = require('../proto_js/Data_pb');


/**
* Character binded  : cc04
* Skill feat_no     : 1 
* Skill feat_name   : slash
*   @function {slash}
*   @param {totalVal} $total_val
*   @param {number} $pow_val
*   @returns {totalVal, effect}
*/
var slash = function ($total_val, $pow_val) {
    let _skill_id_ = 1;
    var totalVal = _.clone($total_val);
    var effectList = [];
    // * Code the skill below

    return {totalVal, effectList};
}
exports.slash = slash;

/**
* Character binded  : cc04
* Skill feat_no     : 6 
* Skill feat_name   : chargedThrust
*   @function {chargedThrust}
*   @param {totalVal} $total_val
*   @param {number} $pow_val
*   @returns {totalVal, effect}
*/
var chargedThrust = function ($total_val, $pow_val) {
    let _skill_id_ = 6;
    var totalVal = _.clone($total_val);
    var effectList = [];
    // * Code the skill below

    return {totalVal, effectList};
}
exports.chargedThrust = chargedThrust;

/**
* Character binded  : cc04
* Skill feat_no     : 4 
* Skill feat_name   : swordDance
*   @function {swordDance}
*   @param {totalVal} $total_val
*   @param {number} $pow_val
*   @returns {totalVal, effect}
*/
var swordDance = function ($total_val, $pow_val) {
    let _skill_id_ = 4;
    var totalVal = _.clone($total_val);
    var effectList = [];
    // * Code the skill below

    return {totalVal, effectList};
}
exports.swordDance = swordDance;

/**
* Character binded  : cc04
* Skill feat_no     : 42 
* Skill feat_name   : abandon
*   @function {abandon}
*   @param {totalVal} $total_val
*   @param {number} $pow_val
*   @returns {totalVal, effect}
*/
var abandon = function ($total_val, $pow_val) {
    let _skill_id_ = 42;
    var totalVal = _.clone($total_val);
    var effectList = [];
    // * Code the skill below

    return {totalVal, effectList};
}
exports.abandon = abandon;

// export this module 
// !endl
