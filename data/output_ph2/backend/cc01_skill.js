
// #node cc01_skill.js 
/** auto-gen from ulz/data/bin/adjust_skill.js 
* generate by : Steven Chm
* generate at : April 15th 2020, 3:53:56 pm
* !please manually modify the skill function
*/
const util = require('./util');
const _ = require('lodash');
const dataStruct = require('../proto_js/Data_pb');


/**
* Character binded  : cc01
* Skill feat_no     : 2 
* Skill feat_name   : precisionMarksmanship
*   @function {precisionMarksmanship}
*   @param {totalVal} $total_val
*   @param {number} $pow_val
*   @returns {totalVal, effect}
*/
var precisionMarksmanship = function ($total_val, $pow_val) {
    let _skill_id_ = 2;
    var totalVal = _.clone($total_val);
    var effectList = [];
    // * Code the skill below

    return {totalVal, effectList};
}
exports.precisionMarksmanship = precisionMarksmanship;

/**
* Character binded  : cc01
* Skill feat_no     : 3 
* Skill feat_name   : torpedoAttack
*   @function {torpedoAttack}
*   @param {totalVal} $total_val
*   @param {number} $pow_val
*   @returns {totalVal, effect}
*/
var torpedoAttack = function ($total_val, $pow_val) {
    let _skill_id_ = 3;
    var totalVal = _.clone($total_val);
    var effectList = [];
    // * Code the skill below

    return {totalVal, effectList};
}
exports.torpedoAttack = torpedoAttack;

/**
* Character binded  : cc01
* Skill feat_no     : 5 
* Skill feat_name   : thornForest
*   @function {thornForest}
*   @param {totalVal} $total_val
*   @param {number} $pow_val
*   @returns {totalVal, effect}
*/
var thornForest = function ($total_val, $pow_val) {
    let _skill_id_ = 5;
    var totalVal = _.clone($total_val);
    var effectList = [];
    // * Code the skill below

    return {totalVal, effectList};
}
exports.thornForest = thornForest;

/**
* Character binded  : cc01
* Skill feat_no     : 35 
* Skill feat_name   : intelligentTactics
*   @function {intelligentTactics}
*   @param {totalVal} $total_val
*   @param {number} $pow_val
*   @returns {totalVal, effect}
*/
var intelligentTactics = function ($total_val, $pow_val) {
    let _skill_id_ = 35;
    var totalVal = _.clone($total_val);
    var effectList = [];
    // * Code the skill below

    return {totalVal, effectList};
}
exports.intelligentTactics = intelligentTactics;

// export this module 
// !endl
