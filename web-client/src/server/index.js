const express = require('express');
const path = require('path');


const app = express();

app.use('/res', express.static(
    path.resolve(__dirname, './public/html')
));


app.listen(
    port = process.env['PORT'] || 80,
    () => {
        console.log(`Server started litening on port: '${port}'.`);
    }
);