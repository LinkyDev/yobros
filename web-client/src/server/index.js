const express = require('express');
const path = require('path');


const config = {
    port: process.env["PORT"] || 80,
    publicPath: process.env["PUBLIC_FOLDER"] || './public'
}

const app = express();

app.use('/res', express.static(
    path.resolve(__dirname, config.publicPath)
));


app.listen(
    config.port,
    () => {
        console.log(`Server started litening on port: '${config.port}'.`);
    }
);