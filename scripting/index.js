require('express')().get('/', (req, res) => res.send('Hello World!')).listen(3000, () => console.log(`Example app listening at http://localhost:` + 3000))
