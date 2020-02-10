const fs = require('fs');
const path = require('path');
const sizeOf = require('image-size');

const getDirectories = source =>
    fs.readdirSync(source, {
        withFileTypes: true
    })
    .filter(dirent => dirent.isDirectory())
    .map(dirent => dirent.name);


function main() {
    // let root_dir = "A:\\Gitrepo\\Unlight-Zwei\\data\\output_ph2";
    let root_dir = "/Users/stephencheng/GitHub/Unlight-Zwei/data/output_ph2";
    let lst = getDirectories(root_dir);
    lst.forEach(CC => {
        skill_set_adj(root_dir, CC);
    });

}


function skill_set_adj(root_dir, CC) {
    let rawdata = fs.readFileSync(path.join(root_dir, CC, 'skill.json'));
    let jsonFile = JSON.parse(rawdata.toString());

    jsonFile.forEach(skil => {
        let y = skil["effect_image"];
        if (y.match(/cutinchara\d\d_skl/gi)) {
            y = y.replace(/cutinchara\d\d_skl/gi, "");
            y = y.replace(/ex\.swf/, "");
            let x = "skill_";
            switch (y) {
                case "00": x += "0"; break;
                case "01": x += "1"; break;
                case "02": x += "2"; break;
                case "03": x += "3"; break;
            }
            x += "_1.png";
            y = x;
        }
        var dim = sizeOf(path.join(root_dir, CC, y));
        dim["name"] = y;
        skil["effect_image"] = dim;
    })
     console.log(jsonFile);
     var tmp = JSON.stringify(jsonFile);
     fs.writeFileSync(path.join(root_dir, CC, "skill.json"), tmp);
     fs.writeFileSync(path.join(root_dir, CC, "skill1.json"), tmp);
     fs.unlinkSync(path.join(root_dir, CC, 'skill1.json'));
   
}
main();