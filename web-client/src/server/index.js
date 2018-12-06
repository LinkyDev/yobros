const express = require('express');
const path = require('path');

const apiRouter = require('./api');

console.log('Initializing server...');
const config = {
    port: process.env['PORT'] || 80,
    publicPath: process.env['PUBLIC_FOLDER'] || './public'
}

const app = express();

[ // Loading static folders
  // Route,    Folder
    ['/',      'html'],
    ['/style', 'css'],
    ['/js',    'js']
].forEach((paths) => {
    const folderPath = path.resolve(
        __dirname,
        config.publicPath,
        paths[1]
    );

    app.use(paths[0], express.static(folderPath));
});

// Load api router
app.use('/api', apiRouter);

// Start server
app.listen(
    config.port,
    () => {
        console.log(`Server started litening on port: '${config.port}'.`);
    }
);