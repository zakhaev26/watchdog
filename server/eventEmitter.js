const EventEmitter = require('events');

const eventEmitter = new EventEmitter();

function function1() {
  console.log('Function 1: Started');

  eventEmitter.emit('message', 'Hello from Function 1!');

  console.log('Function 1: Completed');
}

function function2() {
  console.log('Function 2: Started');

  eventEmitter.on('message', (message) => {
    console.log(`Function 2: Received message - ${message}`);
  });

  console.log('Function 2: Completed');
}

function1();
function2();
