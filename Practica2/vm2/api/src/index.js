const express = require("express");
const dbConn = require('./db.js');
const cors = require("cors");
const app = express();

app.use(express.json());
app.use(cors());
app.use(express.urlencoded({ extended: true }));

app.get("/", function (req, res) {
  res.send("Server on port 5000");
});

app.get("/info", function (req, res) {
  res.send("Laboratorio Sistemas Operativos 1");
});

app.get("/all", function (req, res) {
  dbConn.connect(function (err) {
    if (err) throw err;
    var sql = `SELECT * FROM registros WHERE id = 1;`;
    dbConn.query(sql, function (err, result) {
      if (err) throw err;
      res.send(result);
    });
  });
});


app.listen(5000, () => console.log("Server on port 5000"));