const fs = require('fs');
const path = require('path');
const sizeOf = require('image-size');
const camelCase = require('camelcase');
const os = require('os');
const _ = require('lodash');
const moment = require('moment');
const getDirectories = source =>
    fs.readdirSync(source, {
        withFileTypes: true
    })
        .filter(dirent => dirent.isDirectory())
        .map(dirent => dirent.name);


function main() {
    let root_dir = "A:\\Gitrepo\\Unlight-Zwei\\data\\output_ph2";
    // let root_dir = "/Users/stephencheng/GitHub/Unlight-Zwei/data/output_ph2";
    let lst = getDirectories(root_dir);
    let skill_list = [];
    let counter = 5; 
    lst.forEach(CC => {
        if (CC.includes("cc")  && counter > 0) {
            let sk = skill_main_data(root_dir, CC);
            console.log("sk" + CC, sk[0]['id']);
            skill_list = skill_list.concat(sk);
            counter -- ;
        }
    });

    skill_list.sort((a, b) => a['id'] - b['id']);
    console.log(skill_list);
    let skill_name_list = _.map(skill_list, function (e) {
        return { name: e.name, feat_no: e.feat_no, file_index: e.file_index };
    })
    skill_name_list = _.uniqWith(skill_name_list, function (a, b) { return a.feat_no == b.feat_no && a.file_index == b.file_index });
    console.log(skill_name_list);
    if (!fs.existsSync(path.join(root_dir, "backend"))) {
        fs.mkdirSync(path.join(root_dir, "backend"));
    }
    var tmp = JSON.stringify(skill_list);
    fs.writeFileSync(path.join(root_dir, "backend", "skill_data.json"), tmp);
    let skilgp = _.groupBy(skill_name_list, function (e) {
        return e.file_index;
    });

    for (let [key, e] of Object.entries(skilgp)) {
        let skilFuncGen = [];
        e.forEach(eh => {
            skilFuncGen = skilFuncGen.concat(
                createSkilllTempJs(eh.name, eh.feat_no, eh.file_index)
            );
        })
        console.log(e);
        console.log(skilFuncGen);
        createCCSkillJsFile(root_dir, skilFuncGen, key);
    };

    createCCSkillIndexJs(root_dir, skilgp);

}
var words = ['Zer', 'One', 'Two', 'Thr', 'Fou', 'Fiv', 'Six', 'Sev', 'Eig', 'Nin'];
const umlautMap = {
    '\u00dc': 'UE',
    '\u00c4': 'AE',
    '\u00d6': 'OE',
    '\u00fc': 'ue',
    '\u00e4': 'ae',
    '\u00f6': 'oe',
    '\u00df': 'ss',
}

function replaceUmlaute(str) {
    return str
        .replace(/[\u00dc|\u00c4|\u00d6][a-z]/g, (a) => {
            const big = umlautMap[a.slice(0, 1)];
            return big.charAt(0) + big.charAt(1).toLowerCase() + a.slice(1);
        })
        .replace(new RegExp('[' + Object.keys(umlautMap).join('|') + ']', "g"),
            (a) => umlautMap[a]
        );
}
function wordConv(income) {
    let wordName = camelCase(income);
    wordName = replaceUmlaute(wordName);
    wordName = wordName.replace(/\W+/g, "");
    let newWord = "";
    [...wordName].forEach(c => {
        let k = c;
        if (!isNaN(c)) {
            k = words[c];
        }
        newWord += k
    });
    console.log(newWord);
    return newWord;
}

function skill_main_data(root_dir, CC) {
    let rawdata = fs.readFileSync(path.join(root_dir, CC, 'skill.json'));
    let jsonFile = JSON.parse(rawdata.toString());
    // console.log(jsonFile);
    let lst = []
    jsonFile.forEach(e => {
        lst.push({
            "id": e["id"],
            "feat_no": e["feat_no"],
            "pow": e["pow"],
            "dice_attribute": e["dice_attribute"],
            "condition": e["condition"],
            "created_at": e["created_at"],
            "file_index": CC,
            "name": wordConv(e['name']['en']),
        });
    });

    return lst;
}
function createSkilllTempJs(skill_feat_name, feat_no_i, binding_cc) {
    let func_name = skill_feat_name;
    var function_templ = `\n
    /**
    * Character binded  : ${binding_cc}
    * Skill feat_no     : ${feat_no_i} 
    * Skill feat_name   : ${skill_feat_name}
    */
    ${func_name} : ($total_val, $pow_val , $gm_dt) => {
        let gm_data = new msg.GameDT();
        return gm_data;
    }`;
    return function_templ.toString();
}

function createCCSkillJsFile(root_dir, skill_func_list, binding_cc) {

    // os.userInfo().username
    var file_templ = `
// #node ${binding_cc}_skill.js 
/** auto-gen from ulz/data/bin/adjust_skill.js 
* generate by : ${os.userInfo().username}
* generate at : ${moment().format('MMMM Do YYYY, h:mm:ss a')}
* !please manually modify the skill function
*/
const util = require(./util);
module.exports = {
${skill_func_list}
};
// export this module 
// !endl
`;
    fs.writeFileSync(path.join(root_dir, "backend", `${binding_cc}_skill.js`), file_templ);
}

function createCCSkillIndexJs(root_dir, skilgp) {
    let filePathRef = []; featTable = [];
    for (let [key, e] of Object.entries(skilgp)) {
        e.forEach(eh => {
            featTable.push(`
if(featFuncTable[${eh.feat_no}] == null){
    featFuncTable[${eh.feat_no}] = ${key}Sk.${eh.name};
}`);
        });
        filePathRef.push(`const ${key}Sk = require('./${key}_skill');`);
    };
    var file_templ = `// #node index.js 
/** auto-gen from ulz/data/bin/adjust_skill.js 
* generate by : ${os.userInfo().username}
* generate at : ${moment().format('MMMM Do YYYY, h:mm:ss a')}
* !please manually modify the skill function
*/
// -----
// local skill file import
const util = require('./util');
const dataStruct = require('../proto_js/Data_pb');

${filePathRef.join('\n')}
// export and local run setting
var featFuncTable = [];
${featTable.join('\n')}

var RunFeatFunc = function($request_feat, $request_set, $gm_dt){
    if(featFuncTable[$request_feat] != null){
        return featFuncTable[$request_feat]();
    }
    return null;
};
exports.RunFeatFunc = RunFeatFunc;
`;
    fs.writeFileSync(path.join(root_dir, "backend", `index.js`), file_templ);
}

main();


