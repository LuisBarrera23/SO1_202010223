const mysql = require('mysql2');

const connection = mysql.createConnection({
  host: 'db',
  user: 'root',
  password: '1234',
  database: 'logs',
  port:'3306'
});

connection.connect((err) => {
  if (err) {
    console.error('Error connecting: ' + err);
    return;
  }

  console.log('Connected to MySQL');
});

module.exports = connection;