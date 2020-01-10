const fs = require('fs');
const path = require('path');
const sizeOf = require('image-size');

const getDirectories = source =>
    fs.readdirSync(source, { withFileTypes: true })
        .filter(dirent => dirent.isDirectory())
        .map(dirent => dirent.name);


function main() {
    let root_dir = "A:\\Gitrepo\\Unlight-Zwei\\data\\output_ph2";
    let lst = getDirectories(root_dir);
    lst.forEach(CC => {
        card_set_adj(root_dir, CC);
        skill_set_adj(root_dir, CC);

       
    });

}


function card_set_adj(root_dir, CC) {
    let rawdata = fs.readFileSync(path.join(root_dir, CC, 'card_set.json'));
    var jsonFile = JSON.parse(rawdata.toString());
    // console.log(jsonFile);

    (jsonFile.card_set).forEach(cs => {
        let y = cs["stand_image"];
        var dim = sizeOf(path.join(root_dir, CC, y));
        dim["name"] = y;
        cs["stand_image"] = dim;
        // console.log(cs['stand_image'])

        y = cs["chara_image"];
        var dim1 = {};
        try {
            dim1 = sizeOf(path.join(root_dir, CC, y));
            dim1["name"] = y;
        } catch (e) {
            let tempc = "cover_" + cs["level"] + ".png";
            dim1 = sizeOf(path.join(root_dir, CC, tempc))
            dim1["name"] = tempc;
        }
        cs["chara_image"] = dim1;
        // console.log(cs["chara_image"]);

        y = cs["artifact_image"];
        var dim2 = sizeOf(path.join(root_dir, CC, y));
        dim2["name"] = y;
        cs["artifact_image"] = dim2;
        // console.log(cs["artifact_image"]);

        y = cs["bg_image"];
        var dim3 = sizeOf(path.join(root_dir, CC, y));
        dim3["name"] = y;
        cs["bg_image"] = dim3;
        // console.log(cs["bg_image"]);
    });

    // jsonFile
    console.log(jsonFile);
    // write file
    var tmp = JSON.stringify(jsonFile);
    fs.writeFileSync(path.join(root_dir, CC, "card_set.json"), tmp);

    fs.writeFileSync(path.join(root_dir, CC, "card_set1.json"), tmp);
    fs.unlinkSync(path.join(root_dir, CC, 'card_set1.json'));
}


function skill_set_adj(root_dir, CC) {
    let rawdata = fs.readFileSync(path.join(root_dir, CC, 'skill.json'));
    let jsonFile = JSON.parse(rawdata.toString());

    jsonFile.forEach(skil => {
        let y = skil["effect_image"];
        if (y.match(/cutinchara\d\d_skl/gi)) {
            y = y.replace(/cutinchara\d\d_skl/gi, "");
            y = y.replace(/ex\.swf/, "");
            console.log(y);
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
        // console.log(skil['effect_image']);
    })

    // var tmp = JSON.stringify(jsonFile);

    // fs.writeFileSync(path.join(root_dir, CC, "skill.json"), tmp);
}
main();