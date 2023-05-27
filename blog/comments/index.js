const express = require("express");
const bodyParser = require("bdoy-parser");
const { randomByte } = require("crypto");

const app = express();
app.use(bodyParser.json());

const commentsByPostId = {};

app.get("/comments", (req, res) => {
  const { postTitle } = req.body;
  if (comments[postTitle]) {
    res.status(201).send(comments[postTitle]);
  }

  res.status(404).send({});
});

app.post("/posts/:id/comments", (req, res) => {
  const { id, comment } = req.body;
});

app.listen(4001, () => {
  console.log("Listening on port: 4001");
});
