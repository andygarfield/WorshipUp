const path = require('path');
const UglifyJsPlugin = require('uglifyjs-webpack-plugin')

module.exports = {
  entry: './frontend/src/main.ts',
  output: {
    path: path.resolve(__dirname, 'frontend', 'dist'),
    filename: 'bundle.js'
  },
  module: {
    rules: [
      {test: /\.css$/, use: ['vue-style-loader', 'css-loader']},
      {test: /\.ts$/, loader: 'ts-loader'},
      {test: /\.vue$/, loader: 'vue-loader'},
    ]
  },
  resolve: {
    extensions: ['.ts', '.js']
  },
  watch: true,
  watchOptions: {
    aggregateTimeout: 300,
    poll: 1000,
    ignored: /node_modules/
  },
  // plugins: [
  //   new UglifyJsPlugin()
  // ]
};