
// #node cc03_skill.js 
/** auto-gen from ulz/data/bin/adjust_skill.js 
* generate by : Steven Chm
* generate at : April 15th 2020, 3:53:56 pm
* !please manually modify the skill function
*/
const util = require('./util');
const _ = require('lodash');
const dataStruct = require('../proto_js/Data_pb');


/**
* Character binded  : cc03
* Skill feat_no     : 1 
* Skill feat_name   : bash
*   @function {bash}
*   @param {totalVal} $total_val
*   @param {number} $pow_val
*   @returns {totalVal, effect}
*/
var bash = function ($total_val, $pow_val) {
    let _skill_id_ = 1;
    var totalVal = _.clone($total_val);
    var effectList = [];
    // * Code the skill below

    return {totalVal, effectList};
}
exports.bash = bash;

/**
* Character binded  : cc03
* Skill feat_no     : 12 
* Skill feat_name   : killStance
*   @function {killStance}
*   @param {totalVal} $total_val
*   @param {number} $pow_val
*   @returns {totalVal, effect}
*/
var killStance = function ($total_val, $pow_val) {
    let _skill_id_ = 12;
    var totalVal = _.clone($total_val);
    var effectList = [];
    // * Code the skill below

    return {totalVal, effectList};
}
exports.killStance = killStance;

/**
* Character binded  : cc03
* Skill feat_no     : 15 
* Skill feat_name   : bloodsBlessing
*   @function {bloodsBlessing}
*   @param {totalVal} $total_val
*   @param {number} $pow_val
*   @returns {totalVal, effect}
*/
var bloodsBlessing = function ($total_val, $pow_val) {
    let _skill_id_ = 15;
    var totalVal = _.clone($total_val);
    var effectList = [];
    // * Code the skill below

    return {totalVal, effectList};
}
exports.bloodsBlessing = bloodsBlessing;

/**
* Character binded  : cc03
* Skill feat_no     : 40 
* Skill feat_name   : spiritDrain
*   @function {spiritDrain}
*   @param {totalVal} $total_val
*   @param {number} $pow_val
*   @returns {totalVal, effect}
*/
var spiritDrain = function ($total_val, $pow_val) {
    let _skill_id_ = 40;
    var totalVal = _.clone($total_val);
    var effectList = [];
    // * Code the skill below

    return {totalVal, effectList};
}
exports.spiritDrain = spiritDrain;

// export this module 
// !endl
