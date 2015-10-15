var fs = require('fs');

function listFileContentSync (fileName) {
    var array = [];
    var contents = fs.readFileSync(fileName).toString();
    array.push(contents);
    return array;
}

var dataA = listFileContentSync(process.argv[2]);
var dataB = listFileContentSync(process.argv[3]);

console.log('> Data in A: ' + dataA.length + ' | Data in B: ' + dataB.length);

for (var i = 0; i < dataA.length; i++) {
    var value = dataA[i];
    if (dataB.indexOf(value) < 0) {
        console.log(value);
    }
}