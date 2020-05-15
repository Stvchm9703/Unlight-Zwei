// #node index.js 
/** auto-gen from ulz/data/bin/adjust_skill.js 
* generate by : Steven Chm
* generate at : April 15th 2020, 3:53:56 pm
* !please manually modify the skill function
*/
// -----
// local skill file import
const util = require('./util');
const dataStruct = require('../proto_js/Data_pb');

const cc01Sk = require('./cc01_skill');
const cc02Sk = require('./cc02_skill');
const cc03Sk = require('./cc03_skill');
const cc04Sk = require('./cc04_skill');
const cc05Sk = require('./cc05_skill');
// export and local run setting
var featFuncTable = [];

if(featFuncTable[2] == null){
    featFuncTable[2] = cc01Sk.precisionMarksmanship;
}

if(featFuncTable[3] == null){
    featFuncTable[3] = cc01Sk.torpedoAttack;
}

if(featFuncTable[5] == null){
    featFuncTable[5] = cc01Sk.thornForest;
}

if(featFuncTable[35] == null){
    featFuncTable[35] = cc01Sk.intelligentTactics;
}

if(featFuncTable[2] == null){
    featFuncTable[2] = cc02Sk.rapidfire;
}

if(featFuncTable[10] == null){
    featFuncTable[10] = cc02Sk.supersonicSword;
}

if(featFuncTable[11] == null){
    featFuncTable[11] = cc02Sk.adrenalinAttack;
}

if(featFuncTable[39] == null){
    featFuncTable[39] = cc02Sk.indomitableSpirit;
}

if(featFuncTable[1] == null){
    featFuncTable[1] = cc03Sk.bash;
}

if(featFuncTable[12] == null){
    featFuncTable[12] = cc03Sk.killStance;
}

if(featFuncTable[15] == null){
    featFuncTable[15] = cc03Sk.bloodsBlessing;
}

if(featFuncTable[40] == null){
    featFuncTable[40] = cc03Sk.spiritDrain;
}

if(featFuncTable[1] == null){
    featFuncTable[1] = cc04Sk.slash;
}

if(featFuncTable[6] == null){
    featFuncTable[6] = cc04Sk.chargedThrust;
}

if(featFuncTable[4] == null){
    featFuncTable[4] = cc04Sk.swordDance;
}

if(featFuncTable[42] == null){
    featFuncTable[42] = cc04Sk.abandon;
}

if(featFuncTable[13] == null){
    featFuncTable[13] = cc05Sk.shadowStrike;
}

if(featFuncTable[14] == null){
    featFuncTable[14] = cc05Sk.redFang;
}

if(featFuncTable[16] == null){
    featFuncTable[16] = cc05Sk.rocketCounterAttack;
}

if(featFuncTable[41] == null){
    featFuncTable[41] = cc05Sk.backstab;
}

exports.RunFeatFunc = util.RunFeatFunc;
exports.GetSkillFunc = util.GetSkillFunc;
