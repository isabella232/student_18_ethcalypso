const HelloWorld = artifacts.require("./Greeter.sol");

var fs = require('fs');

function getByteArray(filePath){
    let fileData = fs.readFileSync(filePath).toString('hex');
    let result = []
    for (var i = 0; i < fileData.length; i+=2)
      result.push('0x'+fileData[i]+''+fileData[i+1])
    return result;
}

result = getByteArray('./test.txt');

module.exports = function(deployer) {
	deployer.deploy(HelloWorld, "Bjorn", result);
};