const readline = require("readline").createInterface({
  input: process.stdin,
  output: process.stdout,
});

readline.question(`Input: `, (input) => {
  // Do something here

  readline.close();
});
