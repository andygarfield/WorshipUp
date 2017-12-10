const path = require('path');

module.exports = {
  entry: './webapp/src/app.ts',
  output: {
    path: path.resolve(__dirname, 'webapp', 'dist'),
    filename: 'bundle.js'
  },
  module: {
      rules: [
        {test: /\.ts$/, loader: 'ts-loader'},
      ]
  },
  watch: true,
  watchOptions: {
    aggregateTimeout: 300,
    poll: 1000,
    ignored: /node_modules/
  },
};