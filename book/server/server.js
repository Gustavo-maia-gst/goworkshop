const express = require('express');
const data = require('./personagens.json');

const app = express();

app.get('/', async (req, res) => {
  const delay = Math.random() * 1500 + 500;
  await new Promise(r => setTimeout(r, delay));
  const personagem = data[Math.floor(Math.random() * data.length)];
  res.json(personagem);
});

server = app.listen(3080, () => console.log('Server running on :3000'));

server.on('error', (err) => {
  console.error('Error listening on port:', err.message);
  process.exit(1);
});
