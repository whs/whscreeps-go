console.warn = (...args) => console.log(...args);
global.crypto = {
    getRandomValues: require('polyfill-crypto.getrandomvalues')
};
require('fast-text-encoding');
global.__go_wasm__ = {};
global.Game = Game;
global.RawMemory = RawMemory;
if(typeof InterShardMemory !== 'undefined'){
    global.InterShardMemory = InterShardMemory;
}

require('./wasm_exec');

console.log('Loading wasm');

const wasm = new WebAssembly.Module(require('main.wasm'));
const go = new Go();
let instance = new WebAssembly.Instance(wasm, go.importObject);
console.log('Wasm ready');
go.run(instance);

module.exports.loop = function() {
    if(typeof global.loop === 'undefined') {
        console.log('Skipping tick - code not loaded');
        return;
    }
    global.loop();
}
