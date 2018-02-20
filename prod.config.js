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
  watch: false,
  plugins: [
    new UglifyJsPlugin()
  ]
};