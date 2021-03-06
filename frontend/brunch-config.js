'use strict';

exports.config = {
  paths: {
    watched: [
      'app',
      'node_modules/babel-brunch/node_modules/babel-core/browser-polyfill.min.js'
    ]
  },
  files: {
    javascripts: {
      joinTo: {
        'js/vendor.js': [
          'node_modules/babel-brunch/node_modules/babel-core/browser-polyfill.min.js',
          'bower_components/eventEmitter/EventEmitter.js',
          'bower_components/react/react.js',
          'bower_components/flux/dist/Flux.js',
          'bower_components/fetch/fetch.js',
          'bower_components/general-store/build/general-store.js',
          'bower_components/jquery/dist/jquery.js',
        ],
        'js/bower.js': /^bower_components/,
        'js/app.js': /^app/
      }
    },
    stylesheets: {
      joinTo: {
        'styles/app.css': /^(app|bower_components)/
      }
    }
  },
  onCompile: function() {
    require('fs').appendFile('public/js/app.js', '\n\nrequire(\'js/app\');');
  }
};
